// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: state.proto

/*
Package consensuspb is a generated protocol buffer package.
It is generated from these files:
	state.proto
It has these top-level messages:
	ConsensusRoot
*/
package consensuspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConsensusRoot struct {
	Timestamp   int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Proposer    []byte `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	DynastyRoot []byte `protobuf:"bytes,3,opt,name=dynasty_root,json=dynastyRoot,proto3" json:"dynasty_root,omitempty"`
}

func (m *ConsensusRoot) Reset()                    { *m = ConsensusRoot{} }
func (m *ConsensusRoot) String() string            { return proto.CompactTextString(m) }
func (*ConsensusRoot) ProtoMessage()               {}
func (*ConsensusRoot) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

func (m *ConsensusRoot) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ConsensusRoot) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *ConsensusRoot) GetDynastyRoot() []byte {
	if m != nil {
		return m.DynastyRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusRoot)(nil), "consensuspb.ConsensusRoot")
}

func init() { proto.RegisterFile("state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0x49, 0x2c,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xce, 0xcf, 0x2b, 0x4e, 0xcd, 0x2b,
	0x2e, 0x2d, 0x2e, 0x48, 0x52, 0xca, 0xe1, 0xe2, 0x75, 0x86, 0x71, 0x83, 0xf2, 0xf3, 0x4b, 0x84,
	0x64, 0xb8, 0x38, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0x10, 0x02, 0x42, 0x52, 0x5c, 0x1c, 0x05, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0xa9,
	0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x70, 0xbe, 0x90, 0x22, 0x17, 0x4f, 0x4a, 0x65,
	0x5e, 0x62, 0x71, 0x49, 0x65, 0x7c, 0x51, 0x7e, 0x7e, 0x89, 0x04, 0x33, 0x58, 0x9e, 0x1b, 0x2a,
	0x06, 0x32, 0x3c, 0x89, 0x0d, 0xec, 0x02, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xd5,
	0xcf, 0x41, 0x90, 0x00, 0x00, 0x00,
}