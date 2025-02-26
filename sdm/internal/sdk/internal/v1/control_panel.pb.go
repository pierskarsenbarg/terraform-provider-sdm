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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: control_panel.proto

package v1

import (
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

// ControlPanelGetSSHCAPublicKeyRequest represents a request for an
// organization's SSH CA public key.
type ControlPanelGetSSHCAPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ControlPanelGetSSHCAPublicKeyRequest) Reset() {
	*x = ControlPanelGetSSHCAPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_panel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPanelGetSSHCAPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPanelGetSSHCAPublicKeyRequest) ProtoMessage() {}

func (x *ControlPanelGetSSHCAPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_panel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPanelGetSSHCAPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ControlPanelGetSSHCAPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_control_panel_proto_rawDescGZIP(), []int{0}
}

func (x *ControlPanelGetSSHCAPublicKeyRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

// ControlPanelGetSSHCAPublicKeyResponse represents a request for an
// organization's SSH Certificate Authority public key.
type ControlPanelGetSSHCAPublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The public key of the SSH Certificate Authority, in OpenSSH RSA public
	// key format.
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ControlPanelGetSSHCAPublicKeyResponse) Reset() {
	*x = ControlPanelGetSSHCAPublicKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_panel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPanelGetSSHCAPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPanelGetSSHCAPublicKeyResponse) ProtoMessage() {}

func (x *ControlPanelGetSSHCAPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_panel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPanelGetSSHCAPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*ControlPanelGetSSHCAPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_control_panel_proto_rawDescGZIP(), []int{1}
}

func (x *ControlPanelGetSSHCAPublicKeyResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ControlPanelGetSSHCAPublicKeyResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ControlPanelGetSSHCAPublicKeyResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// ControlPanelVerifyJWTRequest represents a request for x-sdm-token validation.
type ControlPanelVerifyJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The x-sdm-token string.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ControlPanelVerifyJWTRequest) Reset() {
	*x = ControlPanelVerifyJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_panel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPanelVerifyJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPanelVerifyJWTRequest) ProtoMessage() {}

func (x *ControlPanelVerifyJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_panel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPanelVerifyJWTRequest.ProtoReflect.Descriptor instead.
func (*ControlPanelVerifyJWTRequest) Descriptor() ([]byte, []int) {
	return file_control_panel_proto_rawDescGZIP(), []int{2}
}

func (x *ControlPanelVerifyJWTRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ControlPanelVerifyJWTRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// ControlPanelVerifyJWTResponse reports whether x-sdm-token is valid.
type ControlPanelVerifyJWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Reports if the given token is valid.
	Valid bool `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ControlPanelVerifyJWTResponse) Reset() {
	*x = ControlPanelVerifyJWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_panel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPanelVerifyJWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPanelVerifyJWTResponse) ProtoMessage() {}

func (x *ControlPanelVerifyJWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_panel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPanelVerifyJWTResponse.ProtoReflect.Descriptor instead.
func (*ControlPanelVerifyJWTResponse) Descriptor() ([]byte, []int) {
	return file_control_panel_proto_rawDescGZIP(), []int{3}
}

func (x *ControlPanelVerifyJWTResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ControlPanelVerifyJWTResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ControlPanelVerifyJWTResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

var File_control_panel_proto protoreflect.FileDescriptor

var file_control_panel_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x24, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x43, 0x41, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xdf, 0x02, 0x0a, 0x25, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x43,
	0x41, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0xc5, 0x01, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x8e, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2,
	0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x0e, 0xb2,
	0xf4, 0xb3, 0x07, 0x09, 0x21, 0x6a, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0xf2, 0xf8, 0xb3,
	0x07, 0x16, 0xb2, 0xf4, 0xb3, 0x07, 0x11, 0x21, 0x6a, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3,
	0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xf2,
	0xf8, 0xb3, 0x07, 0x1a, 0xb2, 0xf4, 0xb3, 0x07, 0x15, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0xf2, 0xf8,
	0xb3, 0x07, 0x10, 0xb2, 0xf4, 0xb3, 0x07, 0x0b, 0x21, 0x74, 0x79, 0x70, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0a,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x6c, 0x0a, 0x1c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xce, 0x02, 0x0a, 0x1d, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a,
	0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0xc5, 0x01, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x8e, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8,
	0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x0e, 0xb2, 0xf4,
	0xb3, 0x07, 0x09, 0x21, 0x6a, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0xf2, 0xf8, 0xb3, 0x07,
	0x16, 0xb2, 0xf4, 0xb3, 0x07, 0x11, 0x21, 0x6a, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07,
	0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xf2, 0xf8,
	0xb3, 0x07, 0x1a, 0xb2, 0xf4, 0xb3, 0x07, 0x15, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0xf2, 0xf8, 0xb3,
	0x07, 0x10, 0xb2, 0xf4, 0xb3, 0x07, 0x0b, 0x21, 0x74, 0x79, 0x70, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0a, 0xfa,
	0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x32, 0xb4, 0x02, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x99, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x53, 0x48, 0x43, 0x41, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x43, 0x41, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x53,
	0x53, 0x48, 0x43, 0x41, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07,
	0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x1d, 0xaa, 0xf3, 0xb3, 0x07, 0x18, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2f,
	0x73, 0x73, 0x68, 0x2f, 0x63, 0x61, 0x12, 0x87, 0x01, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4a, 0x57, 0x54, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a, 0x57, 0x54, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a, 0x57,
	0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xf9, 0xb3, 0x07, 0x09,
	0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x22, 0xaa, 0xf3,
	0xb3, 0x07, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x42, 0x6b, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67,
	0x42, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6c,
	0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f,
	0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_panel_proto_rawDescOnce sync.Once
	file_control_panel_proto_rawDescData = file_control_panel_proto_rawDesc
)

func file_control_panel_proto_rawDescGZIP() []byte {
	file_control_panel_proto_rawDescOnce.Do(func() {
		file_control_panel_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_panel_proto_rawDescData)
	})
	return file_control_panel_proto_rawDescData
}

var file_control_panel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_control_panel_proto_goTypes = []interface{}{
	(*ControlPanelGetSSHCAPublicKeyRequest)(nil),  // 0: v1.ControlPanelGetSSHCAPublicKeyRequest
	(*ControlPanelGetSSHCAPublicKeyResponse)(nil), // 1: v1.ControlPanelGetSSHCAPublicKeyResponse
	(*ControlPanelVerifyJWTRequest)(nil),          // 2: v1.ControlPanelVerifyJWTRequest
	(*ControlPanelVerifyJWTResponse)(nil),         // 3: v1.ControlPanelVerifyJWTResponse
	(*GetRequestMetadata)(nil),                    // 4: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),                   // 5: v1.GetResponseMetadata
	(*RateLimitMetadata)(nil),                     // 6: v1.RateLimitMetadata
}
var file_control_panel_proto_depIdxs = []int32{
	4, // 0: v1.ControlPanelGetSSHCAPublicKeyRequest.meta:type_name -> v1.GetRequestMetadata
	5, // 1: v1.ControlPanelGetSSHCAPublicKeyResponse.meta:type_name -> v1.GetResponseMetadata
	6, // 2: v1.ControlPanelGetSSHCAPublicKeyResponse.rate_limit:type_name -> v1.RateLimitMetadata
	4, // 3: v1.ControlPanelVerifyJWTRequest.meta:type_name -> v1.GetRequestMetadata
	5, // 4: v1.ControlPanelVerifyJWTResponse.meta:type_name -> v1.GetResponseMetadata
	6, // 5: v1.ControlPanelVerifyJWTResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0, // 6: v1.ControlPanel.GetSSHCAPublicKey:input_type -> v1.ControlPanelGetSSHCAPublicKeyRequest
	2, // 7: v1.ControlPanel.VerifyJWT:input_type -> v1.ControlPanelVerifyJWTRequest
	1, // 8: v1.ControlPanel.GetSSHCAPublicKey:output_type -> v1.ControlPanelGetSSHCAPublicKeyResponse
	3, // 9: v1.ControlPanel.VerifyJWT:output_type -> v1.ControlPanelVerifyJWTResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_control_panel_proto_init() }
func file_control_panel_proto_init() {
	if File_control_panel_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_control_panel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlPanelGetSSHCAPublicKeyRequest); i {
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
		file_control_panel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlPanelGetSSHCAPublicKeyResponse); i {
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
		file_control_panel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlPanelVerifyJWTRequest); i {
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
		file_control_panel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlPanelVerifyJWTResponse); i {
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
			RawDescriptor: file_control_panel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_control_panel_proto_goTypes,
		DependencyIndexes: file_control_panel_proto_depIdxs,
		MessageInfos:      file_control_panel_proto_msgTypes,
	}.Build()
	File_control_panel_proto = out.File
	file_control_panel_proto_rawDesc = nil
	file_control_panel_proto_goTypes = nil
	file_control_panel_proto_depIdxs = nil
}
