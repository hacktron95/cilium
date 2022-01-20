// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: membership.proto

package membershippb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

// RaftAttributes represents the raft related attributes of an etcd member.
type RaftAttributes struct {
	// peerURLs is the list of peers in the raft cluster.
	PeerUrls []string `protobuf:"bytes,1,rep,name=peer_urls,json=peerUrls,proto3" json:"peer_urls,omitempty"`
	// isLearner indicates if the member is raft learner.
	IsLearner            bool     `protobuf:"varint,2,opt,name=is_learner,json=isLearner,proto3" json:"is_learner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaftAttributes) Reset()         { *m = RaftAttributes{} }
func (m *RaftAttributes) String() string { return proto.CompactTextString(m) }
func (*RaftAttributes) ProtoMessage()    {}
func (*RaftAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{0}
}
func (m *RaftAttributes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftAttributes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftAttributes.Merge(m, src)
}
func (m *RaftAttributes) XXX_Size() int {
	return m.Size()
}
func (m *RaftAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_RaftAttributes proto.InternalMessageInfo

// Attributes represents all the non-raft related attributes of an etcd member.
type Attributes struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClientUrls           []string `protobuf:"bytes,2,rep,name=client_urls,json=clientUrls,proto3" json:"client_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}
func (*Attributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{1}
}
func (m *Attributes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attributes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes.Merge(m, src)
}
func (m *Attributes) XXX_Size() int {
	return m.Size()
}
func (m *Attributes) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes proto.InternalMessageInfo

type Member struct {
	ID                   uint64          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	RaftAttributes       *RaftAttributes `protobuf:"bytes,2,opt,name=raft_attributes,json=raftAttributes,proto3" json:"raft_attributes,omitempty"`
	MemberAttributes     *Attributes     `protobuf:"bytes,3,opt,name=member_attributes,json=memberAttributes,proto3" json:"member_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{2}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

type ClusterVersionSetRequest struct {
	Ver                  string   `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterVersionSetRequest) Reset()         { *m = ClusterVersionSetRequest{} }
func (m *ClusterVersionSetRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterVersionSetRequest) ProtoMessage()    {}
func (*ClusterVersionSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{3}
}
func (m *ClusterVersionSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterVersionSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterVersionSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterVersionSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterVersionSetRequest.Merge(m, src)
}
func (m *ClusterVersionSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClusterVersionSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterVersionSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterVersionSetRequest proto.InternalMessageInfo

type ClusterMemberAttrSetRequest struct {
	Member_ID            uint64      `protobuf:"varint,1,opt,name=member_ID,json=memberID,proto3" json:"member_ID,omitempty"`
	MemberAttributes     *Attributes `protobuf:"bytes,2,opt,name=member_attributes,json=memberAttributes,proto3" json:"member_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClusterMemberAttrSetRequest) Reset()         { *m = ClusterMemberAttrSetRequest{} }
func (m *ClusterMemberAttrSetRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterMemberAttrSetRequest) ProtoMessage()    {}
func (*ClusterMemberAttrSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{4}
}
func (m *ClusterMemberAttrSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterMemberAttrSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterMemberAttrSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterMemberAttrSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMemberAttrSetRequest.Merge(m, src)
}
func (m *ClusterMemberAttrSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClusterMemberAttrSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMemberAttrSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMemberAttrSetRequest proto.InternalMessageInfo

type DowngradeInfoSetRequest struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Ver                  string   `protobuf:"bytes,2,opt,name=ver,proto3" json:"ver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DowngradeInfoSetRequest) Reset()         { *m = DowngradeInfoSetRequest{} }
func (m *DowngradeInfoSetRequest) String() string { return proto.CompactTextString(m) }
func (*DowngradeInfoSetRequest) ProtoMessage()    {}
func (*DowngradeInfoSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_949fe0d019050ef5, []int{5}
}
func (m *DowngradeInfoSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DowngradeInfoSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DowngradeInfoSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DowngradeInfoSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DowngradeInfoSetRequest.Merge(m, src)
}
func (m *DowngradeInfoSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *DowngradeInfoSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DowngradeInfoSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DowngradeInfoSetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RaftAttributes)(nil), "membershippb.RaftAttributes")
	proto.RegisterType((*Attributes)(nil), "membershippb.Attributes")
	proto.RegisterType((*Member)(nil), "membershippb.Member")
	proto.RegisterType((*ClusterVersionSetRequest)(nil), "membershippb.ClusterVersionSetRequest")
	proto.RegisterType((*ClusterMemberAttrSetRequest)(nil), "membershippb.ClusterMemberAttrSetRequest")
	proto.RegisterType((*DowngradeInfoSetRequest)(nil), "membershippb.DowngradeInfoSetRequest")
}

func init() { proto.RegisterFile("membership.proto", fileDescriptor_949fe0d019050ef5) }

var fileDescriptor_949fe0d019050ef5 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x99, 0x42, 0xf8, 0xdb, 0xcb, 0x1f, 0xc4, 0x09, 0x89, 0x8d, 0x68, 0x25, 0x5d, 0xb1,
	0x30, 0x98, 0xe8, 0x13, 0xa0, 0xb0, 0x20, 0x81, 0xcd, 0x18, 0xdd, 0x92, 0x56, 0x2e, 0xd8, 0xa4,
	0x74, 0xea, 0xcc, 0x54, 0xd7, 0xbe, 0x85, 0x4f, 0xe0, 0xb3, 0xb0, 0xf4, 0x11, 0x14, 0x5f, 0xc4,
	0x74, 0x5a, 0x4a, 0x49, 0xdc, 0xb8, 0xbb, 0x3d, 0xbd, 0xf7, 0x9c, 0xf3, 0x35, 0x85, 0xd6, 0x0a,
	0x57, 0x3e, 0x0a, 0xf9, 0x18, 0xc4, 0xfd, 0x58, 0x70, 0xc5, 0xe9, 0xff, 0x9d, 0x12, 0xfb, 0xc7,
	0xed, 0x25, 0x5f, 0x72, 0xfd, 0xe2, 0x22, 0x9d, 0xb2, 0x1d, 0x77, 0x02, 0x4d, 0xe6, 0x2d, 0xd4,
	0x40, 0x29, 0x11, 0xf8, 0x89, 0x42, 0x49, 0x3b, 0x60, 0xc5, 0x88, 0x62, 0x96, 0x88, 0x50, 0xda,
	0xa4, 0x5b, 0xed, 0x59, 0xcc, 0x4c, 0x85, 0x3b, 0x11, 0x4a, 0x7a, 0x0a, 0x10, 0xc8, 0x59, 0x88,
	0x9e, 0x88, 0x50, 0xd8, 0x46, 0x97, 0xf4, 0x4c, 0x66, 0x05, 0x72, 0x92, 0x09, 0xee, 0x00, 0xa0,
	0xe4, 0x44, 0xa1, 0x16, 0x79, 0x2b, 0xb4, 0x49, 0x97, 0xf4, 0x2c, 0xa6, 0x67, 0x7a, 0x06, 0x8d,
	0x87, 0x30, 0xc0, 0x48, 0x65, 0xfe, 0x86, 0xf6, 0x87, 0x4c, 0x4a, 0x13, 0xdc, 0x77, 0x02, 0xf5,
	0xa9, 0xee, 0x4d, 0x9b, 0x60, 0x8c, 0x87, 0xfa, 0xba, 0xc6, 0x8c, 0xf1, 0x90, 0x8e, 0xe0, 0x40,
	0x78, 0x0b, 0x35, 0xf3, 0x8a, 0x08, 0xdd, 0xa0, 0x71, 0x79, 0xd2, 0x2f, 0x93, 0xf6, 0xf7, 0x81,
	0x58, 0x53, 0xec, 0x03, 0x8e, 0xe0, 0x30, 0x5b, 0x2f, 0x1b, 0x55, 0xb5, 0x91, 0xbd, 0x6f, 0x54,
	0x32, 0xc9, 0xbf, 0xee, 0x4e, 0x71, 0xcf, 0xc1, 0xbe, 0x09, 0x13, 0xa9, 0x50, 0xdc, 0xa3, 0x90,
	0x01, 0x8f, 0x6e, 0x51, 0x31, 0x7c, 0x4a, 0x50, 0x2a, 0xda, 0x82, 0xea, 0x33, 0x8a, 0x1c, 0x3c,
	0x1d, 0xdd, 0x57, 0x02, 0x9d, 0x7c, 0x7d, 0x5a, 0x38, 0x95, 0x2e, 0x3a, 0x60, 0xe5, 0xa5, 0x0a,
	0x64, 0x33, 0x13, 0x34, 0xf8, 0x2f, 0x8d, 0x8d, 0x3f, 0x37, 0x1e, 0xc1, 0xd1, 0x90, 0xbf, 0x44,
	0x4b, 0xe1, 0xcd, 0x71, 0x1c, 0x2d, 0x78, 0x29, 0xde, 0x86, 0x7f, 0x18, 0x79, 0x7e, 0x88, 0x73,
	0x1d, 0x6e, 0xb2, 0xed, 0xe3, 0x16, 0xc5, 0x28, 0x50, 0xae, 0xdb, 0xeb, 0x2f, 0xa7, 0xb2, 0xde,
	0x38, 0xe4, 0x63, 0xe3, 0x90, 0xcf, 0x8d, 0x43, 0xde, 0xbe, 0x9d, 0x8a, 0x5f, 0xd7, 0xff, 0xd3,
	0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x93, 0x7d, 0x0b, 0x87, 0x02, 0x00, 0x00,
}

func (m *RaftAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftAttributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftAttributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsLearner {
		i--
		if m.IsLearner {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.PeerUrls) > 0 {
		for iNdEx := len(m.PeerUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeerUrls[iNdEx])
			copy(dAtA[i:], m.PeerUrls[iNdEx])
			i = encodeVarintMembership(dAtA, i, uint64(len(m.PeerUrls[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Attributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientUrls) > 0 {
		for iNdEx := len(m.ClientUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientUrls[iNdEx])
			copy(dAtA[i:], m.ClientUrls[iNdEx])
			i = encodeVarintMembership(dAtA, i, uint64(len(m.ClientUrls[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MemberAttributes != nil {
		{
			size, err := m.MemberAttributes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMembership(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RaftAttributes != nil {
		{
			size, err := m.RaftAttributes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMembership(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintMembership(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClusterVersionSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterVersionSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterVersionSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ver) > 0 {
		i -= len(m.Ver)
		copy(dAtA[i:], m.Ver)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Ver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterMemberAttrSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterMemberAttrSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterMemberAttrSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MemberAttributes != nil {
		{
			size, err := m.MemberAttributes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMembership(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Member_ID != 0 {
		i = encodeVarintMembership(dAtA, i, uint64(m.Member_ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DowngradeInfoSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DowngradeInfoSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DowngradeInfoSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ver) > 0 {
		i -= len(m.Ver)
		copy(dAtA[i:], m.Ver)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Ver)))
		i--
		dAtA[i] = 0x12
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMembership(dAtA []byte, offset int, v uint64) int {
	offset -= sovMembership(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RaftAttributes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PeerUrls) > 0 {
		for _, s := range m.PeerUrls {
			l = len(s)
			n += 1 + l + sovMembership(uint64(l))
		}
	}
	if m.IsLearner {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Attributes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if len(m.ClientUrls) > 0 {
		for _, s := range m.ClientUrls {
			l = len(s)
			n += 1 + l + sovMembership(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovMembership(uint64(m.ID))
	}
	if m.RaftAttributes != nil {
		l = m.RaftAttributes.Size()
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.MemberAttributes != nil {
		l = m.MemberAttributes.Size()
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterVersionSetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ver)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterMemberAttrSetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Member_ID != 0 {
		n += 1 + sovMembership(uint64(m.Member_ID))
	}
	if m.MemberAttributes != nil {
		l = m.MemberAttributes.Size()
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DowngradeInfoSetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.Ver)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMembership(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMembership(x uint64) (n int) {
	return sovMembership(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftAttributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: RaftAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerUrls = append(m.PeerUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLearner", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
			m.IsLearner = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *Attributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: Attributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientUrls = append(m.ClientUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RaftAttributes == nil {
				m.RaftAttributes = &RaftAttributes{}
			}
			if err := m.RaftAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MemberAttributes == nil {
				m.MemberAttributes = &Attributes{}
			}
			if err := m.MemberAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *ClusterVersionSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: ClusterVersionSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterVersionSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *ClusterMemberAttrSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: ClusterMemberAttrSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterMemberAttrSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member_ID", wireType)
			}
			m.Member_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Member_ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MemberAttributes == nil {
				m.MemberAttributes = &Attributes{}
			}
			if err := m.MemberAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *DowngradeInfoSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: DowngradeInfoSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DowngradeInfoSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func skipMembership(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMembership
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
					return 0, ErrIntOverflowMembership
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
					return 0, ErrIntOverflowMembership
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
				return 0, ErrInvalidLengthMembership
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMembership
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMembership
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMembership        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMembership          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMembership = fmt.Errorf("proto: unexpected end of group")
)