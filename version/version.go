package version

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var (
	MinClusterVersion = "1.0.0"
	Version           = "1.0.0"
	APIVersion        = "unknown"
)

func init() {
	ver, err := semver.NewVersion(Version)
	if err == nil {
		APIVersion = fmt.Sprintf("%d.%d", ver.Major, ver.Minor)
	}
}

type Versions struct {
	Server  string `json:"etcdserver"`
	Cluster string `json:"etcdcluster"`
}

// Cluster only keeps the major.minor of version
func Cluster(v string) string {
	vs := strings.Split(v, ".")
	if len(vs) <= 2 {
		return v
	}
	return fmt.Sprintf("%s.%s", vs[0], vs[1])
}
