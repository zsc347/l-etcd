// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package etcdserverpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResponseHeader struct {
	// cluster_id is the ID of the cluster which sent the response.
	ClusterId uint64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// member_id is the ID of the member which sent the response.
	MemberId uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	// revision is the key-value store revision when the request was applied.
	// For watch progress responses, the header.revision indicates progress.
	// All future events received in this stream are guaranteed to have a
	// higner revision number than the header.revision number
	Revision int64 `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// raft_term is the raft term when the request was applied.
	RaftTerm             uint64   `protobuf:"varint,4,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetClusterId() uint64 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *ResponseHeader) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *ResponseHeader) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ResponseHeader) GetRaftTerm() uint64 {
	if m != nil {
		return m.RaftTerm
	}
	return 0
}

type LeaseTimeToLiveRequest struct {
	// ID is the lease ID for the lease.
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// keys is true to query all the keys attached to this lease.
	Keys                 bool     `protobuf:"varint,2,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseTimeToLiveRequest) Reset()         { *m = LeaseTimeToLiveRequest{} }
func (m *LeaseTimeToLiveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseTimeToLiveRequest) ProtoMessage()    {}
func (*LeaseTimeToLiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}
func (m *LeaseTimeToLiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseTimeToLiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseTimeToLiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseTimeToLiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseTimeToLiveRequest.Merge(m, src)
}
func (m *LeaseTimeToLiveRequest) XXX_Size() int {
	return m.Size()
}
func (m *LeaseTimeToLiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseTimeToLiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseTimeToLiveRequest proto.InternalMessageInfo

func (m *LeaseTimeToLiveRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseTimeToLiveRequest) GetKeys() bool {
	if m != nil {
		return m.Keys
	}
	return false
}

type LeaseTimeToLiveResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// ID is the lease ID from the keep alive request.
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	// TTL is the remaining TTL in seconds for the lease;
	// the lease will expire in under TTL+1 seconds.
	TTL int64 `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	// GrantedTTL is the initial granted time in seconds upon lease
	// creation/renewal.
	GrantedTTL int64 `protobuf:"varint,4,opt,name=grantedTTL,proto3" json:"grantedTTL,omitempty"`
	// Keys is the list of keys attached to this lease.
	Keys                 [][]byte `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseTimeToLiveResponse) Reset()         { *m = LeaseTimeToLiveResponse{} }
func (m *LeaseTimeToLiveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseTimeToLiveResponse) ProtoMessage()    {}
func (*LeaseTimeToLiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}
func (m *LeaseTimeToLiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseTimeToLiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseTimeToLiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseTimeToLiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseTimeToLiveResponse.Merge(m, src)
}
func (m *LeaseTimeToLiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *LeaseTimeToLiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseTimeToLiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseTimeToLiveResponse proto.InternalMessageInfo

func (m *LeaseTimeToLiveResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LeaseTimeToLiveResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetGrantedTTL() int64 {
	if m != nil {
		return m.GrantedTTL
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseHeader)(nil), "etcdserverpb.ResponseHeader")
	proto.RegisterType((*LeaseTimeToLiveRequest)(nil), "etcdserverpb.LeaseTimeToLiveRequest")
	proto.RegisterType((*LeaseTimeToLiveResponse)(nil), "etcdserverpb.LeaseTimeToLiveResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x75, 0x92, 0x58, 0xda, 0x6f, 0x29, 0x65, 0x10, 0x2d, 0x55, 0x43, 0xe9, 0xaa, 0xab, 0x0a,
	0xea, 0xd2, 0x95, 0x74, 0x61, 0x20, 0xab, 0x21, 0xfb, 0x92, 0x64, 0xbe, 0x31, 0x68, 0x32, 0x71,
	0x66, 0x1a, 0xf0, 0x00, 0xde, 0xc1, 0x0b, 0x78, 0x17, 0x97, 0x1e, 0x41, 0xe2, 0x45, 0x24, 0x33,
	0x21, 0x54, 0xdc, 0xfd, 0x79, 0xef, 0xbf, 0x79, 0xef, 0xcd, 0xc0, 0x48, 0x56, 0xe9, 0xba, 0x92,
	0x42, 0x0b, 0x3a, 0x46, 0x9d, 0x72, 0x85, 0xb2, 0x46, 0x59, 0x25, 0xf3, 0xe3, 0x4c, 0x64, 0xc2,
	0x10, 0x97, 0xed, 0x64, 0x77, 0x96, 0x6f, 0x04, 0x26, 0x0c, 0x55, 0x25, 0x4a, 0x85, 0xf7, 0x18,
	0x73, 0x94, 0xf4, 0x02, 0x20, 0x7d, 0xde, 0x29, 0x8d, 0x72, 0x9b, 0xf3, 0x19, 0x59, 0x90, 0x95,
	0xc7, 0x46, 0x1d, 0x12, 0x70, 0x7a, 0x06, 0xa3, 0x02, 0x8b, 0xc4, 0xb2, 0x8e, 0x61, 0x87, 0x16,
	0x08, 0x38, 0x9d, 0xc3, 0x50, 0x62, 0x9d, 0xab, 0x5c, 0x94, 0x33, 0x77, 0x41, 0x56, 0x2e, 0xeb,
	0xcf, 0xad, 0x50, 0xc6, 0x0f, 0x7a, 0xab, 0x51, 0x16, 0x33, 0xcf, 0x0a, 0x5b, 0x20, 0x42, 0x59,
	0x2c, 0x6f, 0xe1, 0x24, 0xc4, 0x58, 0x61, 0x94, 0x17, 0x18, 0x89, 0x30, 0xaf, 0x91, 0xe1, 0xcb,
	0x0e, 0x95, 0xa6, 0x13, 0x70, 0x82, 0x8d, 0x89, 0xe1, 0x32, 0x27, 0xd8, 0x50, 0x0a, 0xde, 0x13,
	0xbe, 0x2a, 0x63, 0x3d, 0x64, 0x66, 0x5e, 0x7e, 0x10, 0x38, 0xfd, 0x27, 0xb7, 0xa5, 0xe8, 0x0d,
	0x0c, 0x1e, 0x4d, 0x31, 0x73, 0xc7, 0xd1, 0xd5, 0xf9, 0x7a, 0xff, 0x59, 0xd6, 0x7f, 0xcb, 0xb3,
	0x6e, 0xb7, 0x73, 0x75, 0x7a, 0xd7, 0x29, 0xb8, 0x51, 0x14, 0x76, 0x9d, 0xda, 0x91, 0xfa, 0x00,
	0x99, 0x8c, 0x4b, 0x8d, 0xbc, 0x25, 0x3c, 0x43, 0xec, 0x21, 0x7d, 0xce, 0xc3, 0x85, 0xbb, 0x1a,
	0xdb, 0x9c, 0x77, 0xd3, 0xcf, 0xc6, 0x27, 0x5f, 0x8d, 0x4f, 0xbe, 0x1b, 0x9f, 0xbc, 0xff, 0xf8,
	0x07, 0xc9, 0xc0, 0x7c, 0xc3, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x68, 0xa7, 0x19,
	0xb7, 0x01, 0x00, 0x00,
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RaftTerm != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftTerm))
		i--
		dAtA[i] = 0x20
	}
	if m.Revision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
	if m.MemberId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberId))
		i--
		dAtA[i] = 0x10
	}
	if m.ClusterId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ClusterId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseTimeToLiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseTimeToLiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseTimeToLiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Keys {
		i--
		if m.Keys {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseTimeToLiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseTimeToLiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseTimeToLiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.GrantedTTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.GrantedTTL))
		i--
		dAtA[i] = 0x20
	}
	if m.TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResponseHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterId != 0 {
		n += 1 + sovRpc(uint64(m.ClusterId))
	}
	if m.MemberId != 0 {
		n += 1 + sovRpc(uint64(m.MemberId))
	}
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.RaftTerm != 0 {
		n += 1 + sovRpc(uint64(m.RaftTerm))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseTimeToLiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.Keys {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseTimeToLiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.GrantedTTL != 0 {
		n += 1 + sovRpc(uint64(m.GrantedTTL))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			m.ClusterId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberId", wireType)
			}
			m.MemberId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftTerm", wireType)
			}
			m.RaftTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftTerm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseTimeToLiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Keys = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseTimeToLiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantedTTL", wireType)
			}
			m.GrantedTTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrantedTTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpc = fmt.Errorf("proto: unexpected end of group")
)
