package leasehttp

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/l-etcd/etcdserver/etcdserverpb"
	"github.com/l-etcd/lease"
	"github.com/l-etcd/lease/leasepb"
)

var (
	LeasePrefix         = "/leases"
	LeaseInternalPrefix = "/leases/internal"
	applyTimeout        = time.Second
	ErrLeaseHTTPTimeout = errors.New("waiting for node to catch up its applied index has timed out")
)

// NewHandler returns an http Handler for lease renewals
func NewHandler(l lease.Lessor, waitch func() <-chan struct{}) http.Handler {
	return &leaseHandler{l, waitch}
}

type leaseHandler struct {
	l      lease.Lessor
	waitch func() <-chan struct{}
}

func (h *leaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}

	var v []byte
	switch r.URL.Path {
	case LeasePrefix:
		lreq := etcdserverpb.LeaseTimeToLiveRequest{}
		if uerr := lreq.Unmarshal(b); uerr != nil {
			http.Error(w, "error unmarshalling request", http.StatusBadRequest)
			return
		}
		select {
		case <-h.waitch():
		case <-time.After(applyTimeout):
			http.Error(w, ErrLeaseHTTPTimeout.Error(), http.StatusRequestTimeout)
			return
		}

		ttl, rerr := h.l.Renew(lease.LeaseID(lreq.ID))
		if rerr != nil {
			if rerr == lease.ErrLeaseNotFound {
				http.Error(w, rerr.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		// TODO: fill out ResponseHeader
		resp := &etcdserverpb.LeaseKeepAliveResponse{ID: lreq.ID, TTL: ttl}
		v, err = resp.Marshal()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case LeaseInternalPrefix:
		lreq := leasepb.LeaseInternalRequest{}
		if lerr := lreq.Unmarshal(b); lerr != nil {
			http.Error(w, "err unmarshalling request", http.StatusBadRequest)
			return
		}
		select {
		case <-h.waitch():
		case <-time.After(applyTimeout):
			http.Error(w, ErrLeaseHTTPTimeout.Error(), http.StatusRequestTimeout)
			return
		}

		l := h.l.Lookup(lease.LeaseID(lreq.LeaseTimeToLiveRequest.ID))
		if l == nil {
			http.Error(w, lease.ErrLeaseNotFound.Error(), http.StatusNotFound)
			return
		}

		// TODO: fill out ResponseHeader
		resp := &leasepb.LeaseInternalResponse{
			LeaseTimeToLiveResponse: &etcdserverpb.LeaseTimeToLiveResponse{
				Header:     &etcdserverpb.ResponseHeader{},
				ID:         lreq.LeaseTimeToLiveRequest.ID,
				TTL:        int64(l.Remaining().Seconds()),
				GrantedTTL: l.TTL(),
			},
		}

		if lreq.LeaseTimeToLiveRequest.Keys {
			ks := l.Keys()
			kbs := make([][]byte, len(ks))
			for i := range ks {
				kbs[i] = []byte(ks[i])
			}
			resp.LeaseTimeToLiveResponse.Keys = kbs
		}

		v, err = resp.Marshal()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:

	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(v)
}
