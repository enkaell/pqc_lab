// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: server.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Main communication models
type Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Versions []*Algorithm_Version `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *Algorithm) Reset() {
	*x = Algorithm{}
	mi := &file_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm) ProtoMessage() {}

func (x *Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm.ProtoReflect.Descriptor instead.
func (*Algorithm) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Algorithm) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Algorithm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Algorithm) GetVersions() []*Algorithm_Version {
	if x != nil {
		return x.Versions
	}
	return nil
}

// Request message
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response message for listing all algorithms
type AlgorithmList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithms []*Algorithm `protobuf:"bytes,1,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
}

func (x *AlgorithmList) Reset() {
	*x = AlgorithmList{}
	mi := &file_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlgorithmList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmList) ProtoMessage() {}

func (x *AlgorithmList) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmList.ProtoReflect.Descriptor instead.
func (*AlgorithmList) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *AlgorithmList) GetAlgorithms() []*Algorithm {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

// Response message for listing all versions of the algorithm
type AlgorithmVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []*Algorithm_Version `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *AlgorithmVersionList) Reset() {
	*x = AlgorithmVersionList{}
	mi := &file_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlgorithmVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmVersionList) ProtoMessage() {}

func (x *AlgorithmVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmVersionList.ProtoReflect.Descriptor instead.
func (*AlgorithmVersionList) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *AlgorithmVersionList) GetVersions() []*Algorithm_Version {
	if x != nil {
		return x.Versions
	}
	return nil
}

// Response message for listing all runs of the algorithm version
type AlgorithmVersionRunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runs []*Algorithm_Version_Run `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *AlgorithmVersionRunList) Reset() {
	*x = AlgorithmVersionRunList{}
	mi := &file_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlgorithmVersionRunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmVersionRunList) ProtoMessage() {}

func (x *AlgorithmVersionRunList) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmVersionRunList.ProtoReflect.Descriptor instead.
func (*AlgorithmVersionRunList) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *AlgorithmVersionRunList) GetRuns() []*Algorithm_Version_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

type Algorithm_Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Runs []*Algorithm_Version_Run `protobuf:"bytes,3,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *Algorithm_Version) Reset() {
	*x = Algorithm_Version{}
	mi := &file_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Algorithm_Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm_Version) ProtoMessage() {}

func (x *Algorithm_Version) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm_Version.ProtoReflect.Descriptor instead.
func (*Algorithm_Version) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Algorithm_Version) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Algorithm_Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Algorithm_Version) GetRuns() []*Algorithm_Version_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

type Algorithm_Version_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Keygen         string `protobuf:"bytes,3,opt,name=keygen,proto3" json:"keygen,omitempty"`
	Sign           string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
	Verify         string `protobuf:"bytes,5,opt,name=verify,proto3" json:"verify,omitempty"`
	PrivateKeySize uint32 `protobuf:"varint,6,opt,name=private_key_size,json=privateKeySize,proto3" json:"private_key_size,omitempty"`
	PublicKeySize  uint32 `protobuf:"varint,7,opt,name=public_key_size,json=publicKeySize,proto3" json:"public_key_size,omitempty"`
	SignatureSize  uint32 `protobuf:"varint,8,opt,name=signature_size,json=signatureSize,proto3" json:"signature_size,omitempty"`
}

func (x *Algorithm_Version_Run) Reset() {
	*x = Algorithm_Version_Run{}
	mi := &file_server_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Algorithm_Version_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm_Version_Run) ProtoMessage() {}

func (x *Algorithm_Version_Run) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm_Version_Run.ProtoReflect.Descriptor instead.
func (*Algorithm_Version_Run) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Algorithm_Version_Run) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Algorithm_Version_Run) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Algorithm_Version_Run) GetKeygen() string {
	if x != nil {
		return x.Keygen
	}
	return ""
}

func (x *Algorithm_Version_Run) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *Algorithm_Version_Run) GetVerify() string {
	if x != nil {
		return x.Verify
	}
	return ""
}

func (x *Algorithm_Version_Run) GetPrivateKeySize() uint32 {
	if x != nil {
		return x.PrivateKeySize
	}
	return 0
}

func (x *Algorithm_Version_Run) GetPublicKeySize() uint32 {
	if x != nil {
		return x.PublicKeySize
	}
	return 0
}

func (x *Algorithm_Version_Run) GetSignatureSize() uint32 {
	if x != nil {
		return x.SignatureSize
	}
	return 0
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x09,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xc2, 0x02,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x1a, 0xe6, 0x01, 0x0a, 0x03, 0x52, 0x75,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x22, 0x46,
	0x0a, 0x14, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x45, 0x0a, 0x17, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x32, 0xfd, 0x01,
	0x0a, 0x10, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e,
	0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x41, 0x6c, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x67, 0x52, 0x75, 0x6e, 0x73, 0x42, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x52, 0x75,
	0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x0e, 0x2e, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_server_proto_goTypes = []any{
	(*Algorithm)(nil),               // 0: Algorithm
	(*Request)(nil),                 // 1: Request
	(*AlgorithmList)(nil),           // 2: AlgorithmList
	(*AlgorithmVersionList)(nil),    // 3: AlgorithmVersionList
	(*AlgorithmVersionRunList)(nil), // 4: AlgorithmVersionRunList
	(*Algorithm_Version)(nil),       // 5: Algorithm.Version
	(*Algorithm_Version_Run)(nil),   // 6: Algorithm.Version.Run
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_server_proto_depIdxs = []int32{
	5, // 0: Algorithm.versions:type_name -> Algorithm.Version
	0, // 1: AlgorithmList.algorithms:type_name -> Algorithm
	5, // 2: AlgorithmVersionList.versions:type_name -> Algorithm.Version
	6, // 3: AlgorithmVersionRunList.runs:type_name -> Algorithm.Version.Run
	6, // 4: Algorithm.Version.runs:type_name -> Algorithm.Version.Run
	7, // 5: AlgorithmService.GetAllAlgorithms:input_type -> google.protobuf.Empty
	1, // 6: AlgorithmService.GetAllAlgVersionByAlgName:input_type -> Request
	1, // 7: AlgorithmService.GetAllAlgRunsByVersion:input_type -> Request
	2, // 8: AlgorithmService.RunAlgorithms:input_type -> AlgorithmList
	2, // 9: AlgorithmService.GetAllAlgorithms:output_type -> AlgorithmList
	3, // 10: AlgorithmService.GetAllAlgVersionByAlgName:output_type -> AlgorithmVersionList
	4, // 11: AlgorithmService.GetAllAlgRunsByVersion:output_type -> AlgorithmVersionRunList
	2, // 12: AlgorithmService.RunAlgorithms:output_type -> AlgorithmList
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
