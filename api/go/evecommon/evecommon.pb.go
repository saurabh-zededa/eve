// Copyright(c) 2017-2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.6.1
// source: evecommon/evecommon.proto

package evecommon

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HashAlgorithm int32

const (
	HashAlgorithm_HASH_ALGORITHM_INVALID        HashAlgorithm = 0
	HashAlgorithm_HASH_ALGORITHM_SHA256_16BYTES HashAlgorithm = 1 // hash with sha256, the 1st 16 bytes of result
	HashAlgorithm_HASH_ALGORITHM_SHA256_32BYTES HashAlgorithm = 2 // hash with sha256, the whole 32 bytes of result
)

// Enum value maps for HashAlgorithm.
var (
	HashAlgorithm_name = map[int32]string{
		0: "HASH_ALGORITHM_INVALID",
		1: "HASH_ALGORITHM_SHA256_16BYTES",
		2: "HASH_ALGORITHM_SHA256_32BYTES",
	}
	HashAlgorithm_value = map[string]int32{
		"HASH_ALGORITHM_INVALID":        0,
		"HASH_ALGORITHM_SHA256_16BYTES": 1,
		"HASH_ALGORITHM_SHA256_32BYTES": 2,
	}
)

func (x HashAlgorithm) Enum() *HashAlgorithm {
	p := new(HashAlgorithm)
	*p = x
	return p
}

func (x HashAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_evecommon_proto_enumTypes[0].Descriptor()
}

func (HashAlgorithm) Type() protoreflect.EnumType {
	return &file_evecommon_evecommon_proto_enumTypes[0]
}

func (x HashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgorithm.Descriptor instead.
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{0}
}

type ZCertType int32

const (
	ZCertType_Z_CERT_TYPE_INVALID ZCertType = 0
	// controller generated certificates
	ZCertType_Z_CERT_TYPE_CONTROLLER_SIGNING       ZCertType = 1 //to sign payload envelopes
	ZCertType_Z_CERT_TYPE_CONTROLLER_INTERMEDIATE  ZCertType = 2 //intermediate certs used to validate the certificates
	ZCertType_Z_CERT_TYPE_CONTROLLER_ECDH_EXCHANGE ZCertType = 3 //to share symmetric key using ECDH
	// device generated certificates
	ZCertType_Z_CERT_TYPE_DEVICE_ONBOARDING         ZCertType = 10 //for identifying the device
	ZCertType_Z_CERT_TYPE_DEVICE_RESTRICTED_SIGNING ZCertType = 11 //node for attestation
	ZCertType_Z_CERT_TYPE_DEVICE_ENDORSEMENT_RSA    ZCertType = 12 //endorsement key certificate with RSASSA signing algorithm
	ZCertType_Z_CERT_TYPE_DEVICE_ECDH_EXCHANGE      ZCertType = 13 //to share symmetric key using ECDH
)

// Enum value maps for ZCertType.
var (
	ZCertType_name = map[int32]string{
		0:  "Z_CERT_TYPE_INVALID",
		1:  "Z_CERT_TYPE_CONTROLLER_SIGNING",
		2:  "Z_CERT_TYPE_CONTROLLER_INTERMEDIATE",
		3:  "Z_CERT_TYPE_CONTROLLER_ECDH_EXCHANGE",
		10: "Z_CERT_TYPE_DEVICE_ONBOARDING",
		11: "Z_CERT_TYPE_DEVICE_RESTRICTED_SIGNING",
		12: "Z_CERT_TYPE_DEVICE_ENDORSEMENT_RSA",
		13: "Z_CERT_TYPE_DEVICE_ECDH_EXCHANGE",
	}
	ZCertType_value = map[string]int32{
		"Z_CERT_TYPE_INVALID":                   0,
		"Z_CERT_TYPE_CONTROLLER_SIGNING":        1,
		"Z_CERT_TYPE_CONTROLLER_INTERMEDIATE":   2,
		"Z_CERT_TYPE_CONTROLLER_ECDH_EXCHANGE":  3,
		"Z_CERT_TYPE_DEVICE_ONBOARDING":         10,
		"Z_CERT_TYPE_DEVICE_RESTRICTED_SIGNING": 11,
		"Z_CERT_TYPE_DEVICE_ENDORSEMENT_RSA":    12,
		"Z_CERT_TYPE_DEVICE_ECDH_EXCHANGE":      13,
	}
)

func (x ZCertType) Enum() *ZCertType {
	p := new(ZCertType)
	*p = x
	return p
}

func (x ZCertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZCertType) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_evecommon_proto_enumTypes[1].Descriptor()
}

func (ZCertType) Type() protoreflect.EnumType {
	return &file_evecommon_evecommon_proto_enumTypes[1]
}

func (x ZCertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZCertType.Descriptor instead.
func (ZCertType) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{1}
}

type ZCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashAlgo   HashAlgorithm `protobuf:"varint,1,opt,name=hash_algo,json=hashAlgo,proto3,enum=org.lfedge.eve.common.HashAlgorithm" json:"hash_algo,omitempty"` //hash method used to arrive at certHash
	Type       ZCertType     `protobuf:"varint,2,opt,name=type,proto3,enum=org.lfedge.eve.common.ZCertType" json:"type,omitempty"`                             // certificate type
	Hash       []byte        `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`                                                                   // certificate hash, defined by hash algorithm
	Cert       []byte        `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`                                                                   //X509 cert in .PEM format
	Attributes *ZCertAttr    `protobuf:"bytes,5,opt,name=attributes,proto3" json:"attributes,omitempty"`                                                       //properties of this certificate
}

func (x *ZCert) Reset() {
	*x = ZCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evecommon_evecommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZCert) ProtoMessage() {}

func (x *ZCert) ProtoReflect() protoreflect.Message {
	mi := &file_evecommon_evecommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZCert.ProtoReflect.Descriptor instead.
func (*ZCert) Descriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{0}
}

func (x *ZCert) GetHashAlgo() HashAlgorithm {
	if x != nil {
		return x.HashAlgo
	}
	return HashAlgorithm_HASH_ALGORITHM_INVALID
}

func (x *ZCert) GetType() ZCertType {
	if x != nil {
		return x.Type
	}
	return ZCertType_Z_CERT_TYPE_INVALID
}

func (x *ZCert) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ZCert) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *ZCert) GetAttributes() *ZCertAttr {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ZCertAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMutable bool `protobuf:"varint,1,opt,name=is_mutable,json=isMutable,proto3" json:"is_mutable,omitempty"` //set to false for immutable certificates
}

func (x *ZCertAttr) Reset() {
	*x = ZCertAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evecommon_evecommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZCertAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZCertAttr) ProtoMessage() {}

func (x *ZCertAttr) ProtoReflect() protoreflect.Message {
	mi := &file_evecommon_evecommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZCertAttr.ProtoReflect.Descriptor instead.
func (*ZCertAttr) Descriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{1}
}

func (x *ZCertAttr) GetIsMutable() bool {
	if x != nil {
		return x.IsMutable
	}
	return false
}

var File_evecommon_evecommon_proto protoreflect.FileDescriptor

var file_evecommon_evecommon_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0xea, 0x01, 0x0a, 0x05, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x12, 0x41, 0x0a, 0x09,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12,
	0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x40, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0x2a, 0x0a, 0x09, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2a, 0x71, 0x0a, 0x0d, 0x48,
	0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x16,
	0x48, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x48, 0x41, 0x53, 0x48,
	0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35,
	0x36, 0x5f, 0x31, 0x36, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x48,
	0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48,
	0x41, 0x32, 0x35, 0x36, 0x5f, 0x33, 0x32, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x02, 0x2a, 0xb7,
	0x02, 0x0a, 0x09, 0x5a, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f,
	0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x5a, 0x5f, 0x43,
	0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x4c, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x45, 0x43, 0x44,
	0x48, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d,
	0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12,
	0x29, 0x0a, 0x25, 0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44,
	0x5f, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x26, 0x0a, 0x22, 0x5a, 0x5f,
	0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x53, 0x41,
	0x10, 0x0c, 0x12, 0x24, 0x0a, 0x20, 0x5a, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x43, 0x44, 0x48, 0x5f, 0x45, 0x58,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x42, 0x4d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x42, 0x09, 0x45, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64,
	0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evecommon_evecommon_proto_rawDescOnce sync.Once
	file_evecommon_evecommon_proto_rawDescData = file_evecommon_evecommon_proto_rawDesc
)

func file_evecommon_evecommon_proto_rawDescGZIP() []byte {
	file_evecommon_evecommon_proto_rawDescOnce.Do(func() {
		file_evecommon_evecommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_evecommon_evecommon_proto_rawDescData)
	})
	return file_evecommon_evecommon_proto_rawDescData
}

var file_evecommon_evecommon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_evecommon_evecommon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_evecommon_evecommon_proto_goTypes = []interface{}{
	(HashAlgorithm)(0), // 0: org.lfedge.eve.common.HashAlgorithm
	(ZCertType)(0),     // 1: org.lfedge.eve.common.ZCertType
	(*ZCert)(nil),      // 2: org.lfedge.eve.common.ZCert
	(*ZCertAttr)(nil),  // 3: org.lfedge.eve.common.ZCertAttr
}
var file_evecommon_evecommon_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.common.ZCert.hash_algo:type_name -> org.lfedge.eve.common.HashAlgorithm
	1, // 1: org.lfedge.eve.common.ZCert.type:type_name -> org.lfedge.eve.common.ZCertType
	3, // 2: org.lfedge.eve.common.ZCert.attributes:type_name -> org.lfedge.eve.common.ZCertAttr
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_evecommon_evecommon_proto_init() }
func file_evecommon_evecommon_proto_init() {
	if File_evecommon_evecommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evecommon_evecommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZCert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_evecommon_evecommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZCertAttr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_evecommon_evecommon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_evecommon_evecommon_proto_goTypes,
		DependencyIndexes: file_evecommon_evecommon_proto_depIdxs,
		EnumInfos:         file_evecommon_evecommon_proto_enumTypes,
		MessageInfos:      file_evecommon_evecommon_proto_msgTypes,
	}.Build()
	File_evecommon_evecommon_proto = out.File
	file_evecommon_evecommon_proto_rawDesc = nil
	file_evecommon_evecommon_proto_goTypes = nil
	file_evecommon_evecommon_proto_depIdxs = nil
}
