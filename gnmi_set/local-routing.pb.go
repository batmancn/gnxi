// Code generated by protoc-gen-go. DO NOT EDIT.
// source: local-routing.proto

package proto_local_routing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type NextHopState struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	NextHop              string   `protobuf:"bytes,2,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Metric               uint32   `protobuf:"varint,3,opt,name=metric,proto3" json:"metric,omitempty"`
	Recurse              bool     `protobuf:"varint,4,opt,name=recurse,proto3" json:"recurse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextHopState) Reset()         { *m = NextHopState{} }
func (m *NextHopState) String() string { return proto.CompactTextString(m) }
func (*NextHopState) ProtoMessage()    {}
func (*NextHopState) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{0}
}

func (m *NextHopState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextHopState.Unmarshal(m, b)
}
func (m *NextHopState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextHopState.Marshal(b, m, deterministic)
}
func (m *NextHopState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextHopState.Merge(m, src)
}
func (m *NextHopState) XXX_Size() int {
	return xxx_messageInfo_NextHopState.Size(m)
}
func (m *NextHopState) XXX_DiscardUnknown() {
	xxx_messageInfo_NextHopState.DiscardUnknown(m)
}

var xxx_messageInfo_NextHopState proto.InternalMessageInfo

func (m *NextHopState) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NextHopState) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *NextHopState) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *NextHopState) GetRecurse() bool {
	if m != nil {
		return m.Recurse
	}
	return false
}

type NextHopConfig struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	NextHop              string   `protobuf:"bytes,2,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Metric               uint32   `protobuf:"varint,3,opt,name=metric,proto3" json:"metric,omitempty"`
	Recurse              bool     `protobuf:"varint,4,opt,name=recurse,proto3" json:"recurse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextHopConfig) Reset()         { *m = NextHopConfig{} }
func (m *NextHopConfig) String() string { return proto.CompactTextString(m) }
func (*NextHopConfig) ProtoMessage()    {}
func (*NextHopConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{1}
}

func (m *NextHopConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextHopConfig.Unmarshal(m, b)
}
func (m *NextHopConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextHopConfig.Marshal(b, m, deterministic)
}
func (m *NextHopConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextHopConfig.Merge(m, src)
}
func (m *NextHopConfig) XXX_Size() int {
	return xxx_messageInfo_NextHopConfig.Size(m)
}
func (m *NextHopConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NextHopConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NextHopConfig proto.InternalMessageInfo

func (m *NextHopConfig) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NextHopConfig) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *NextHopConfig) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *NextHopConfig) GetRecurse() bool {
	if m != nil {
		return m.Recurse
	}
	return false
}

type NextHopsNextHop struct {
	Config               *NextHopConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	State                *NextHopState  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NextHopsNextHop) Reset()         { *m = NextHopsNextHop{} }
func (m *NextHopsNextHop) String() string { return proto.CompactTextString(m) }
func (*NextHopsNextHop) ProtoMessage()    {}
func (*NextHopsNextHop) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{2}
}

func (m *NextHopsNextHop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextHopsNextHop.Unmarshal(m, b)
}
func (m *NextHopsNextHop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextHopsNextHop.Marshal(b, m, deterministic)
}
func (m *NextHopsNextHop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextHopsNextHop.Merge(m, src)
}
func (m *NextHopsNextHop) XXX_Size() int {
	return xxx_messageInfo_NextHopsNextHop.Size(m)
}
func (m *NextHopsNextHop) XXX_DiscardUnknown() {
	xxx_messageInfo_NextHopsNextHop.DiscardUnknown(m)
}

var xxx_messageInfo_NextHopsNextHop proto.InternalMessageInfo

func (m *NextHopsNextHop) GetConfig() *NextHopConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *NextHopsNextHop) GetState() *NextHopState {
	if m != nil {
		return m.State
	}
	return nil
}

type StaticNextHops struct {
	NextHop              []*NextHopsNextHop `protobuf:"bytes,1,rep,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StaticNextHops) Reset()         { *m = StaticNextHops{} }
func (m *StaticNextHops) String() string { return proto.CompactTextString(m) }
func (*StaticNextHops) ProtoMessage()    {}
func (*StaticNextHops) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{3}
}

func (m *StaticNextHops) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticNextHops.Unmarshal(m, b)
}
func (m *StaticNextHops) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticNextHops.Marshal(b, m, deterministic)
}
func (m *StaticNextHops) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticNextHops.Merge(m, src)
}
func (m *StaticNextHops) XXX_Size() int {
	return xxx_messageInfo_StaticNextHops.Size(m)
}
func (m *StaticNextHops) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticNextHops.DiscardUnknown(m)
}

var xxx_messageInfo_StaticNextHops proto.InternalMessageInfo

func (m *StaticNextHops) GetNextHop() []*NextHopsNextHop {
	if m != nil {
		return m.NextHop
	}
	return nil
}

type StaticRoutesStatic struct {
	Address              string          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	NextHops             *StaticNextHops `protobuf:"bytes,2,opt,name=next_hops,json=nextHops,proto3" json:"next_hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StaticRoutesStatic) Reset()         { *m = StaticRoutesStatic{} }
func (m *StaticRoutesStatic) String() string { return proto.CompactTextString(m) }
func (*StaticRoutesStatic) ProtoMessage()    {}
func (*StaticRoutesStatic) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{4}
}

func (m *StaticRoutesStatic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticRoutesStatic.Unmarshal(m, b)
}
func (m *StaticRoutesStatic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticRoutesStatic.Marshal(b, m, deterministic)
}
func (m *StaticRoutesStatic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticRoutesStatic.Merge(m, src)
}
func (m *StaticRoutesStatic) XXX_Size() int {
	return xxx_messageInfo_StaticRoutesStatic.Size(m)
}
func (m *StaticRoutesStatic) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticRoutesStatic.DiscardUnknown(m)
}

var xxx_messageInfo_StaticRoutesStatic proto.InternalMessageInfo

func (m *StaticRoutesStatic) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StaticRoutesStatic) GetNextHops() *StaticNextHops {
	if m != nil {
		return m.NextHops
	}
	return nil
}

type LocalRoutesStaticRoutes struct {
	Route                []*StaticRoutesStatic `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalRoutesStaticRoutes) Reset()         { *m = LocalRoutesStaticRoutes{} }
func (m *LocalRoutesStaticRoutes) String() string { return proto.CompactTextString(m) }
func (*LocalRoutesStaticRoutes) ProtoMessage()    {}
func (*LocalRoutesStaticRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{5}
}

func (m *LocalRoutesStaticRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalRoutesStaticRoutes.Unmarshal(m, b)
}
func (m *LocalRoutesStaticRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalRoutesStaticRoutes.Marshal(b, m, deterministic)
}
func (m *LocalRoutesStaticRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRoutesStaticRoutes.Merge(m, src)
}
func (m *LocalRoutesStaticRoutes) XXX_Size() int {
	return xxx_messageInfo_LocalRoutesStaticRoutes.Size(m)
}
func (m *LocalRoutesStaticRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRoutesStaticRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRoutesStaticRoutes proto.InternalMessageInfo

func (m *LocalRoutesStaticRoutes) GetRoute() []*StaticRoutesStatic {
	if m != nil {
		return m.Route
	}
	return nil
}

type LocalRoutes struct {
	StaticRoutes         *LocalRoutesStaticRoutes `protobuf:"bytes,1,opt,name=static_routes,json=staticRoutes,proto3" json:"static_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LocalRoutes) Reset()         { *m = LocalRoutes{} }
func (m *LocalRoutes) String() string { return proto.CompactTextString(m) }
func (*LocalRoutes) ProtoMessage()    {}
func (*LocalRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b55a02301b89ba, []int{6}
}

func (m *LocalRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalRoutes.Unmarshal(m, b)
}
func (m *LocalRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalRoutes.Marshal(b, m, deterministic)
}
func (m *LocalRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRoutes.Merge(m, src)
}
func (m *LocalRoutes) XXX_Size() int {
	return xxx_messageInfo_LocalRoutes.Size(m)
}
func (m *LocalRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRoutes proto.InternalMessageInfo

func (m *LocalRoutes) GetStaticRoutes() *LocalRoutesStaticRoutes {
	if m != nil {
		return m.StaticRoutes
	}
	return nil
}

func init() {
	proto.RegisterType((*NextHopState)(nil), "proto.local_routing.NextHopState")
	proto.RegisterType((*NextHopConfig)(nil), "proto.local_routing.NextHopConfig")
	proto.RegisterType((*NextHopsNextHop)(nil), "proto.local_routing.NextHopsNextHop")
	proto.RegisterType((*StaticNextHops)(nil), "proto.local_routing.StaticNextHops")
	proto.RegisterType((*StaticRoutesStatic)(nil), "proto.local_routing.StaticRoutesStatic")
	proto.RegisterType((*LocalRoutesStaticRoutes)(nil), "proto.local_routing.LocalRoutesStaticRoutes")
	proto.RegisterType((*LocalRoutes)(nil), "proto.local_routing.LocalRoutes")
}

func init() { proto.RegisterFile("local-routing.proto", fileDescriptor_98b55a02301b89ba) }

var fileDescriptor_98b55a02301b89ba = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x51, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x25, 0xdf, 0x7c, 0xf3, 0x77, 0x3b, 0xa3, 0x70, 0x47, 0xb4, 0xee, 0x6a, 0x14, 0xec, 0x42,
	0xbb, 0x18, 0x17, 0x82, 0x20, 0x0a, 0x6e, 0x5c, 0x88, 0x30, 0x71, 0xe3, 0x6e, 0x1c, 0x3b, 0x71,
	0x2c, 0x8c, 0x4d, 0x4d, 0x52, 0xe8, 0x13, 0xf8, 0xdc, 0xd2, 0x24, 0x95, 0x16, 0x6a, 0x97, 0xae,
	0x92, 0xc3, 0xcd, 0xf9, 0xb9, 0x27, 0x30, 0xdb, 0x8a, 0x78, 0xb5, 0x3d, 0x97, 0x22, 0xd7, 0x49,
	0xba, 0x89, 0x32, 0x29, 0xb4, 0xc0, 0x99, 0x39, 0x22, 0x33, 0x5a, 0xba, 0x11, 0xfd, 0x84, 0xc9,
	0x23, 0x2f, 0xf4, 0xbd, 0xc8, 0x9e, 0xf4, 0x4a, 0x73, 0xdc, 0x83, 0x7e, 0x92, 0xae, 0x79, 0xe1,
	0x93, 0x80, 0x84, 0x63, 0x66, 0x01, 0x1e, 0xc2, 0x28, 0xe5, 0x85, 0x5e, 0xbe, 0x8b, 0xcc, 0xff,
	0x67, 0x06, 0xc3, 0xd4, 0xb2, 0x70, 0x1f, 0x06, 0x1f, 0x5c, 0xcb, 0x24, 0xf6, 0x7b, 0x01, 0x09,
	0xa7, 0xcc, 0x21, 0xf4, 0x61, 0x28, 0x79, 0x9c, 0x4b, 0xc5, 0xfd, 0xff, 0x01, 0x09, 0x47, 0xac,
	0x82, 0x54, 0xc2, 0xd4, 0x59, 0xde, 0x89, 0xf4, 0x2d, 0xd9, 0xfc, 0x85, 0xe7, 0x17, 0x81, 0x5d,
	0x67, 0xaa, 0xdc, 0x89, 0x57, 0x30, 0x88, 0x4d, 0x00, 0xe3, 0xeb, 0xcd, 0x69, 0xd4, 0x52, 0x50,
	0xd4, 0x88, 0xca, 0x1c, 0x03, 0x2f, 0xa1, 0xaf, 0xca, 0xbe, 0x4c, 0x32, 0x6f, 0x7e, 0xd4, 0x45,
	0x35, 0xc5, 0x32, 0xfb, 0x9e, 0x2e, 0x60, 0xa7, 0xc4, 0x49, 0x5c, 0xa5, 0xc1, 0x9b, 0xda, 0x9e,
	0x24, 0xe8, 0x85, 0xde, 0xfc, 0xa4, 0x4b, 0xad, 0x8a, 0xff, 0xd3, 0x06, 0xcd, 0x00, 0xad, 0x24,
	0x13, 0xb9, 0xe6, 0xca, 0xde, 0xcb, 0x2e, 0x56, 0xeb, 0xb5, 0xe4, 0x4a, 0xb9, 0x5a, 0x2b, 0x88,
	0xb7, 0x30, 0xae, 0x0c, 0x95, 0xcb, 0x7f, 0xdc, 0xea, 0xd8, 0x0c, 0xca, 0x46, 0xce, 0x50, 0xd1,
	0x67, 0x38, 0x78, 0x28, 0x5f, 0xd6, 0x0d, 0xed, 0x1d, 0xaf, 0xa1, 0x5f, 0xd2, 0xb9, 0x5b, 0xe5,
	0xb4, 0x43, 0xb8, 0xc1, 0xb6, 0x2c, 0xfa, 0x02, 0x5e, 0x4d, 0x19, 0x17, 0x30, 0x55, 0x66, 0x6e,
	0xa8, 0x5c, 0xb9, 0x9f, 0x3a, 0x6b, 0x55, 0xfd, 0x25, 0x12, 0x9b, 0xa8, 0x1a, 0x7a, 0x1d, 0x18,
	0xea, 0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xbb, 0x2f, 0xb3, 0x23, 0x03, 0x00, 0x00,
}
