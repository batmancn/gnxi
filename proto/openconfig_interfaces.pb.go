// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openconfig_interfaces.proto

package mt_proto

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

type AddressesAddress struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressesAddress) Reset()         { *m = AddressesAddress{} }
func (m *AddressesAddress) String() string { return proto.CompactTextString(m) }
func (*AddressesAddress) ProtoMessage()    {}
func (*AddressesAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_347be543b845363f, []int{0}
}

func (m *AddressesAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressesAddress.Unmarshal(m, b)
}
func (m *AddressesAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressesAddress.Marshal(b, m, deterministic)
}
func (m *AddressesAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressesAddress.Merge(m, src)
}
func (m *AddressesAddress) XXX_Size() int {
	return xxx_messageInfo_AddressesAddress.Size(m)
}
func (m *AddressesAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressesAddress.DiscardUnknown(m)
}

var xxx_messageInfo_AddressesAddress proto.InternalMessageInfo

func (m *AddressesAddress) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Ipv4Addresses struct {
	Addresses            []*AddressesAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ipv4Addresses) Reset()         { *m = Ipv4Addresses{} }
func (m *Ipv4Addresses) String() string { return proto.CompactTextString(m) }
func (*Ipv4Addresses) ProtoMessage()    {}
func (*Ipv4Addresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_347be543b845363f, []int{1}
}

func (m *Ipv4Addresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4Addresses.Unmarshal(m, b)
}
func (m *Ipv4Addresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4Addresses.Marshal(b, m, deterministic)
}
func (m *Ipv4Addresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4Addresses.Merge(m, src)
}
func (m *Ipv4Addresses) XXX_Size() int {
	return xxx_messageInfo_Ipv4Addresses.Size(m)
}
func (m *Ipv4Addresses) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4Addresses.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4Addresses proto.InternalMessageInfo

func (m *Ipv4Addresses) GetAddresses() []*AddressesAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressesAddress)(nil), "mt.proto.AddressesAddress")
	proto.RegisterType((*Ipv4Addresses)(nil), "mt.proto.Ipv4Addresses")
}

func init() { proto.RegisterFile("openconfig_interfaces.proto", fileDescriptor_347be543b845363f) }

var fileDescriptor_347be543b845363f = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x2f, 0x48, 0xcd,
	0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x8f, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x2d, 0x81, 0xb0, 0x94, 0xd4, 0xb8,
	0x04, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x53, 0x8b, 0xa1, 0x0c, 0x21, 0x21, 0x2e, 0x96,
	0xc4, 0x94, 0x94, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x93, 0x8b,
	0xd7, 0xb3, 0xa0, 0xcc, 0x04, 0xae, 0x56, 0xc8, 0x82, 0x8b, 0x33, 0x11, 0xc6, 0x91, 0x60, 0x54,
	0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0x19, 0xab, 0x87, 0x6e, 0x66, 0x10, 0x42, 0x71, 0x12,
	0x1b, 0x58, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xdb, 0x7c, 0x9b, 0xa2, 0x00, 0x00,
	0x00,
}
