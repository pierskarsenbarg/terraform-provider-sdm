// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: secret_store_types.proto

// This file was generated by protogen. DO NOT EDIT.

package v1

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

// A SecretStore is a server where resource secrets (passwords, keys) are stored.
// Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
type SecretStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SecretStore:
	//	*SecretStore_Aws
	//	*SecretStore_VaultTls
	//	*SecretStore_VaultToken
	SecretStore isSecretStore_SecretStore `protobuf_oneof:"secret_store"`
}

func (x *SecretStore) Reset() {
	*x = SecretStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_store_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretStore) ProtoMessage() {}

func (x *SecretStore) ProtoReflect() protoreflect.Message {
	mi := &file_secret_store_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretStore.ProtoReflect.Descriptor instead.
func (*SecretStore) Descriptor() ([]byte, []int) {
	return file_secret_store_types_proto_rawDescGZIP(), []int{0}
}

func (m *SecretStore) GetSecretStore() isSecretStore_SecretStore {
	if m != nil {
		return m.SecretStore
	}
	return nil
}

func (x *SecretStore) GetAws() *AWSStore {
	if x, ok := x.GetSecretStore().(*SecretStore_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *SecretStore) GetVaultTls() *VaultTLSStore {
	if x, ok := x.GetSecretStore().(*SecretStore_VaultTls); ok {
		return x.VaultTls
	}
	return nil
}

func (x *SecretStore) GetVaultToken() *VaultTokenStore {
	if x, ok := x.GetSecretStore().(*SecretStore_VaultToken); ok {
		return x.VaultToken
	}
	return nil
}

type isSecretStore_SecretStore interface {
	isSecretStore_SecretStore()
}

type SecretStore_Aws struct {
	Aws *AWSStore `protobuf:"bytes,3,opt,name=aws,proto3,oneof"`
}

type SecretStore_VaultTls struct {
	VaultTls *VaultTLSStore `protobuf:"bytes,1,opt,name=vault_tls,json=vaultTls,proto3,oneof"`
}

type SecretStore_VaultToken struct {
	VaultToken *VaultTokenStore `protobuf:"bytes,2,opt,name=vault_token,json=vaultToken,proto3,oneof"`
}

func (*SecretStore_Aws) isSecretStore_SecretStore() {}

func (*SecretStore_VaultTls) isSecretStore_SecretStore() {}

func (*SecretStore_VaultToken) isSecretStore_SecretStore() {}

type AWSStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the SecretStore.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique human-readable name of the SecretStore.
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// Tags is a map of key, value pairs.
	Tags *Tags `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *AWSStore) Reset() {
	*x = AWSStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_store_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSStore) ProtoMessage() {}

func (x *AWSStore) ProtoReflect() protoreflect.Message {
	mi := &file_secret_store_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSStore.ProtoReflect.Descriptor instead.
func (*AWSStore) Descriptor() ([]byte, []int) {
	return file_secret_store_types_proto_rawDescGZIP(), []int{1}
}

func (x *AWSStore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AWSStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AWSStore) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AWSStore) GetTags() *Tags {
	if x != nil {
		return x.Tags
	}
	return nil
}

type VaultTLSStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the SecretStore.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique human-readable name of the SecretStore.
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServerAddress  string `protobuf:"bytes,3,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	CACertPath     string `protobuf:"bytes,4,opt,name=CA_cert_path,json=CACertPath,proto3" json:"CA_cert_path,omitempty"`
	ClientCertPath string `protobuf:"bytes,5,opt,name=client_cert_path,json=clientCertPath,proto3" json:"client_cert_path,omitempty"`
	ClientKeyPath  string `protobuf:"bytes,6,opt,name=client_key_path,json=clientKeyPath,proto3" json:"client_key_path,omitempty"`
	// Tags is a map of key, value pairs.
	Tags *Tags `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *VaultTLSStore) Reset() {
	*x = VaultTLSStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_store_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultTLSStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultTLSStore) ProtoMessage() {}

func (x *VaultTLSStore) ProtoReflect() protoreflect.Message {
	mi := &file_secret_store_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultTLSStore.ProtoReflect.Descriptor instead.
func (*VaultTLSStore) Descriptor() ([]byte, []int) {
	return file_secret_store_types_proto_rawDescGZIP(), []int{2}
}

func (x *VaultTLSStore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VaultTLSStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VaultTLSStore) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *VaultTLSStore) GetCACertPath() string {
	if x != nil {
		return x.CACertPath
	}
	return ""
}

func (x *VaultTLSStore) GetClientCertPath() string {
	if x != nil {
		return x.ClientCertPath
	}
	return ""
}

func (x *VaultTLSStore) GetClientKeyPath() string {
	if x != nil {
		return x.ClientKeyPath
	}
	return ""
}

func (x *VaultTLSStore) GetTags() *Tags {
	if x != nil {
		return x.Tags
	}
	return nil
}

type VaultTokenStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the SecretStore.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique human-readable name of the SecretStore.
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServerAddress string `protobuf:"bytes,3,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	// Tags is a map of key, value pairs.
	Tags *Tags `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *VaultTokenStore) Reset() {
	*x = VaultTokenStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_store_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultTokenStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultTokenStore) ProtoMessage() {}

func (x *VaultTokenStore) ProtoReflect() protoreflect.Message {
	mi := &file_secret_store_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultTokenStore.ProtoReflect.Descriptor instead.
func (*VaultTokenStore) Descriptor() ([]byte, []int) {
	return file_secret_store_types_proto_rawDescGZIP(), []int{3}
}

func (x *VaultTokenStore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VaultTokenStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VaultTokenStore) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *VaultTokenStore) GetTags() *Tags {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_secret_store_types_proto protoreflect.FileDescriptor

var file_secret_store_types_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x74,
	0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x61, 0x77, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x42, 0x0d, 0xf2, 0xf8, 0xb3, 0x07, 0x08, 0x8a, 0xf4, 0xb3, 0x07, 0x03,
	0x61, 0x77, 0x73, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x4c, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x42, 0x12, 0xf2, 0xf8, 0xb3, 0x07, 0x0d, 0x8a, 0xf4, 0xb3, 0x07, 0x08, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x4c, 0x53, 0x48, 0x00, 0x52, 0x08, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6c, 0x73,
	0x12, 0x4c, 0x0a, 0x0b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x14, 0xf2, 0xf8, 0xb3, 0x07,
	0x0f, 0x8a, 0xf4, 0xb3, 0x07, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x48, 0x00, 0x52, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x0a,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x42, 0x41, 0x0a, 0x0c, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0xaa, 0xf8, 0xb3, 0x07,
	0x10, 0xa2, 0xf8, 0xb3, 0x07, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0xaa, 0xf8, 0xb3, 0x07, 0x09, 0xaa, 0xf8, 0xb3, 0x07, 0x04, 0x74, 0x61, 0x67, 0x73, 0xaa,
	0xf8, 0xb3, 0x07, 0x09, 0xaa, 0xf8, 0xb3, 0x07, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x02,
	0x0a, 0x08, 0x41, 0x57, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xf8, 0xb3, 0x07, 0x1c, 0xa2, 0xf3, 0xb3,
	0x07, 0x02, 0x49, 0x44, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xca, 0xf3, 0xb3, 0x07, 0x0b, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xf2, 0xf8, 0xb3,
	0x07, 0x13, 0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xf2, 0xf8, 0xb3,
	0x07, 0x25, 0xa2, 0xf3, 0xb3, 0x07, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3, 0x07, 0x01, 0x8a, 0xf4, 0xb3, 0x07,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x31, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x42, 0x13, 0xf2, 0xf8, 0xb3, 0x07, 0x0e, 0xa2, 0xf3,
	0xb3, 0x07, 0x04, 0x54, 0x61, 0x67, 0x73, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x2d, 0xa2, 0xf3, 0xb3, 0x07, 0x0b, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xda, 0xf3,
	0xb3, 0x07, 0x03, 0x61, 0x77, 0x73, 0xe2, 0xf3, 0xb3, 0x07, 0x03, 0x61, 0x77, 0x73, 0xea, 0xf3,
	0xb3, 0x07, 0x03, 0x61, 0x77, 0x73, 0x22, 0xe0, 0x04, 0x0a, 0x0d, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x4c, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xf8, 0xb3, 0x07, 0x1c, 0xa2, 0xf3, 0xb3, 0x07, 0x02,
	0x49, 0x44, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xca, 0xf3, 0xb3, 0x07, 0x0b, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xf2, 0xf8, 0xb3, 0x07, 0x13,
	0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3,
	0xb3, 0x07, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x38, 0xf2, 0xf8, 0xb3, 0x07, 0x33, 0xa2, 0xf3, 0xb3, 0x07, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0,
	0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3, 0x07, 0x01, 0x8a, 0xf4, 0xb3, 0x07, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x41,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2d, 0xf2, 0xf8, 0xb3, 0x07, 0x28, 0xa2, 0xf3, 0xb3, 0x07, 0x0a, 0x43, 0x41, 0x43, 0x65,
	0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3, 0x07, 0x01,
	0x8a, 0xf4, 0xb3, 0x07, 0x0a, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x0a, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x64, 0x0a, 0x10, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xf2, 0xf8, 0xb3, 0x07, 0x35, 0xa2, 0xf3, 0xb3, 0x07,
	0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3, 0x07, 0x01, 0x8a, 0xf4,
	0xb3, 0x07, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x60, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xf2, 0xf8, 0xb3, 0x07,
	0x33, 0xa2, 0xf3, 0xb3, 0x07, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x50,
	0x61, 0x74, 0x68, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3,
	0x07, 0x01, 0x8a, 0xf4, 0xb3, 0x07, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x42, 0x13, 0xf2, 0xf8, 0xb3,
	0x07, 0x0e, 0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x54, 0x61, 0x67, 0x73, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x41, 0xfa, 0xf8, 0xb3, 0x07, 0x3c, 0xa2, 0xf3, 0xb3,
	0x07, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xa8, 0xf3, 0xb3,
	0x07, 0x01, 0xda, 0xf3, 0xb3, 0x07, 0x08, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x4c, 0x53, 0xe2,
	0xf3, 0xb3, 0x07, 0x08, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x4c, 0x53, 0xea, 0xf3, 0xb3, 0x07,
	0x08, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x4c, 0x53, 0x22, 0xcf, 0x02, 0x0a, 0x0f, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xf8, 0xb3, 0x07, 0x1c,
	0xa2, 0xf3, 0xb3, 0x07, 0x02, 0x49, 0x44, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xca, 0xf3, 0xb3, 0x07,
	0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xf2, 0xf8, 0xb3, 0x07, 0x13, 0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0xb0, 0xf3,
	0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5f,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xf2, 0xf8, 0xb3, 0x07, 0x33, 0xa2, 0xf3, 0xb3,
	0x07, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0xd0, 0xf3, 0xb3, 0x07, 0x01, 0x8a, 0xf4,
	0xb3, 0x07, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x31, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x42, 0x13, 0xf2, 0xf8, 0xb3, 0x07, 0x0e, 0xa2, 0xf3,
	0xb3, 0x07, 0x04, 0x54, 0x61, 0x67, 0x73, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x3a, 0x47, 0xfa, 0xf8, 0xb3, 0x07, 0x42, 0xa2, 0xf3, 0xb3, 0x07, 0x0b, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xda, 0xf3,
	0xb3, 0x07, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0xe2, 0xf3, 0xb3,
	0x07, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0xea, 0xf3, 0xb3, 0x07,
	0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6d, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x19, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x6c,
	0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f,
	0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_secret_store_types_proto_rawDescOnce sync.Once
	file_secret_store_types_proto_rawDescData = file_secret_store_types_proto_rawDesc
)

func file_secret_store_types_proto_rawDescGZIP() []byte {
	file_secret_store_types_proto_rawDescOnce.Do(func() {
		file_secret_store_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_secret_store_types_proto_rawDescData)
	})
	return file_secret_store_types_proto_rawDescData
}

var file_secret_store_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_secret_store_types_proto_goTypes = []interface{}{
	(*SecretStore)(nil),     // 0: v1.SecretStore
	(*AWSStore)(nil),        // 1: v1.AWSStore
	(*VaultTLSStore)(nil),   // 2: v1.VaultTLSStore
	(*VaultTokenStore)(nil), // 3: v1.VaultTokenStore
	(*Tags)(nil),            // 4: v1.Tags
}
var file_secret_store_types_proto_depIdxs = []int32{
	1, // 0: v1.SecretStore.aws:type_name -> v1.AWSStore
	2, // 1: v1.SecretStore.vault_tls:type_name -> v1.VaultTLSStore
	3, // 2: v1.SecretStore.vault_token:type_name -> v1.VaultTokenStore
	4, // 3: v1.AWSStore.tags:type_name -> v1.Tags
	4, // 4: v1.VaultTLSStore.tags:type_name -> v1.Tags
	4, // 5: v1.VaultTokenStore.tags:type_name -> v1.Tags
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_secret_store_types_proto_init() }
func file_secret_store_types_proto_init() {
	if File_secret_store_types_proto != nil {
		return
	}
	file_options_proto_init()
	file_tags_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_secret_store_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretStore); i {
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
		file_secret_store_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSStore); i {
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
		file_secret_store_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultTLSStore); i {
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
		file_secret_store_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultTokenStore); i {
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
	file_secret_store_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SecretStore_Aws)(nil),
		(*SecretStore_VaultTls)(nil),
		(*SecretStore_VaultToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_secret_store_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_secret_store_types_proto_goTypes,
		DependencyIndexes: file_secret_store_types_proto_depIdxs,
		MessageInfos:      file_secret_store_types_proto_msgTypes,
	}.Build()
	File_secret_store_types_proto = out.File
	file_secret_store_types_proto_rawDesc = nil
	file_secret_store_types_proto_goTypes = nil
	file_secret_store_types_proto_depIdxs = nil
}
