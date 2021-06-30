// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: api/watcher.proto

package api

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Watcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Keywords []string `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *Watcher) Reset() {
	*x = Watcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watcher) ProtoMessage() {}

func (x *Watcher) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watcher.ProtoReflect.Descriptor instead.
func (*Watcher) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *Watcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Watcher) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type CreateWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watcher *Watcher `protobuf:"bytes,1,opt,name=watcher,proto3" json:"watcher,omitempty"`
}

func (x *CreateWatcherRequest) Reset() {
	*x = CreateWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWatcherRequest) ProtoMessage() {}

func (x *CreateWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWatcherRequest.ProtoReflect.Descriptor instead.
func (*CreateWatcherRequest) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWatcherRequest) GetWatcher() *Watcher {
	if x != nil {
		return x.Watcher
	}
	return nil
}

type DeleteWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteWatcherRequest) Reset() {
	*x = DeleteWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWatcherRequest) ProtoMessage() {}

func (x *DeleteWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWatcherRequest.ProtoReflect.Descriptor instead.
func (*DeleteWatcherRequest) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteWatcherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetWatcherRequest) Reset() {
	*x = GetWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWatcherRequest) ProtoMessage() {}

func (x *GetWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWatcherRequest.ProtoReflect.Descriptor instead.
func (*GetWatcherRequest) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{3}
}

func (x *GetWatcherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListWatcherRequest) Reset() {
	*x = ListWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWatcherRequest) ProtoMessage() {}

func (x *ListWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWatcherRequest.ProtoReflect.Descriptor instead.
func (*ListWatcherRequest) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{4}
}

func (x *ListWatcherRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListWatcherRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListWatcherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextPageToken string     `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	Watchers      []*Watcher `protobuf:"bytes,2,rep,name=watchers,proto3" json:"watchers,omitempty"`
}

func (x *ListWatcherResponse) Reset() {
	*x = ListWatcherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWatcherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWatcherResponse) ProtoMessage() {}

func (x *ListWatcherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWatcherResponse.ProtoReflect.Descriptor instead.
func (*ListWatcherResponse) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{5}
}

func (x *ListWatcherResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListWatcherResponse) GetWatchers() []*Watcher {
	if x != nil {
		return x.Watchers
	}
	return nil
}

type UpdateWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Watcher    *Watcher              `protobuf:"bytes,2,opt,name=watcher,proto3" json:"watcher,omitempty"`
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateWatcherRequest) Reset() {
	*x = UpdateWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_watcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWatcherRequest) ProtoMessage() {}

func (x *UpdateWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_watcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWatcherRequest.ProtoReflect.Descriptor instead.
func (*UpdateWatcherRequest) Descriptor() ([]byte, []int) {
	return file_api_watcher_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateWatcherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateWatcherRequest) GetWatcher() *Watcher {
	if x != nil {
		return x.Watcher
	}
	return nil
}

func (x *UpdateWatcherRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_api_watcher_proto protoreflect.FileDescriptor

var file_api_watcher_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x59, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f,
	0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x07, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x82, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x43, 0x0a, 0x08, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x41, 0x0a, 0x07, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73,
	0x6b, 0x32, 0xb1, 0x04, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73,
	0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x6f,
	0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73,
	0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x31, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x76, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x6e,
	0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f,
	0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e,
	0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x67, 0x6f,
	0x73, 0x6c, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_watcher_proto_rawDescOnce sync.Once
	file_api_watcher_proto_rawDescData = file_api_watcher_proto_rawDesc
)

func file_api_watcher_proto_rawDescGZIP() []byte {
	file_api_watcher_proto_rawDescOnce.Do(func() {
		file_api_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_watcher_proto_rawDescData)
	})
	return file_api_watcher_proto_rawDescData
}

var file_api_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_watcher_proto_goTypes = []interface{}{
	(*Watcher)(nil),              // 0: nokamoto.github.com.egosla.api.Watcher
	(*CreateWatcherRequest)(nil), // 1: nokamoto.github.com.egosla.api.CreateWatcherRequest
	(*DeleteWatcherRequest)(nil), // 2: nokamoto.github.com.egosla.api.DeleteWatcherRequest
	(*GetWatcherRequest)(nil),    // 3: nokamoto.github.com.egosla.api.GetWatcherRequest
	(*ListWatcherRequest)(nil),   // 4: nokamoto.github.com.egosla.api.ListWatcherRequest
	(*ListWatcherResponse)(nil),  // 5: nokamoto.github.com.egosla.api.ListWatcherResponse
	(*UpdateWatcherRequest)(nil), // 6: nokamoto.github.com.egosla.api.UpdateWatcherRequest
	(*field_mask.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*empty.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_api_watcher_proto_depIdxs = []int32{
	0, // 0: nokamoto.github.com.egosla.api.CreateWatcherRequest.watcher:type_name -> nokamoto.github.com.egosla.api.Watcher
	0, // 1: nokamoto.github.com.egosla.api.ListWatcherResponse.watchers:type_name -> nokamoto.github.com.egosla.api.Watcher
	0, // 2: nokamoto.github.com.egosla.api.UpdateWatcherRequest.watcher:type_name -> nokamoto.github.com.egosla.api.Watcher
	7, // 3: nokamoto.github.com.egosla.api.UpdateWatcherRequest.update_mask:type_name -> google.protobuf.FieldMask
	1, // 4: nokamoto.github.com.egosla.api.WatcherService.CreateWatcher:input_type -> nokamoto.github.com.egosla.api.CreateWatcherRequest
	2, // 5: nokamoto.github.com.egosla.api.WatcherService.DeleteWatcher:input_type -> nokamoto.github.com.egosla.api.DeleteWatcherRequest
	3, // 6: nokamoto.github.com.egosla.api.WatcherService.GetWatcher:input_type -> nokamoto.github.com.egosla.api.GetWatcherRequest
	4, // 7: nokamoto.github.com.egosla.api.WatcherService.ListWatcher:input_type -> nokamoto.github.com.egosla.api.ListWatcherRequest
	6, // 8: nokamoto.github.com.egosla.api.WatcherService.UpdateWatcher:input_type -> nokamoto.github.com.egosla.api.UpdateWatcherRequest
	0, // 9: nokamoto.github.com.egosla.api.WatcherService.CreateWatcher:output_type -> nokamoto.github.com.egosla.api.Watcher
	8, // 10: nokamoto.github.com.egosla.api.WatcherService.DeleteWatcher:output_type -> google.protobuf.Empty
	0, // 11: nokamoto.github.com.egosla.api.WatcherService.GetWatcher:output_type -> nokamoto.github.com.egosla.api.Watcher
	5, // 12: nokamoto.github.com.egosla.api.WatcherService.ListWatcher:output_type -> nokamoto.github.com.egosla.api.ListWatcherResponse
	0, // 13: nokamoto.github.com.egosla.api.WatcherService.UpdateWatcher:output_type -> nokamoto.github.com.egosla.api.Watcher
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_watcher_proto_init() }
func file_api_watcher_proto_init() {
	if File_api_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watcher); i {
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
		file_api_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWatcherRequest); i {
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
		file_api_watcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWatcherRequest); i {
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
		file_api_watcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWatcherRequest); i {
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
		file_api_watcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWatcherRequest); i {
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
		file_api_watcher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWatcherResponse); i {
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
		file_api_watcher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWatcherRequest); i {
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
			RawDescriptor: file_api_watcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_watcher_proto_goTypes,
		DependencyIndexes: file_api_watcher_proto_depIdxs,
		MessageInfos:      file_api_watcher_proto_msgTypes,
	}.Build()
	File_api_watcher_proto = out.File
	file_api_watcher_proto_rawDesc = nil
	file_api_watcher_proto_goTypes = nil
	file_api_watcher_proto_depIdxs = nil
}
