// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/subscription.proto

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

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Watcher string `protobuf:"bytes,2,opt,name=watcher,proto3" json:"watcher,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subscription) GetWatcher() string {
	if x != nil {
		return x.Watcher
	}
	return ""
}

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubscriptionRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type DeleteSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSubscriptionRequest) Reset() {
	*x = DeleteSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriptionRequest) ProtoMessage() {}

func (x *DeleteSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteSubscriptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubscriptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListSubscriptionRequest) Reset() {
	*x = ListSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionRequest) ProtoMessage() {}

func (x *ListSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubscriptionRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListSubscriptionRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextPageToken string          `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	Watchers      []*Subscription `protobuf:"bytes,2,rep,name=watchers,proto3" json:"watchers,omitempty"`
}

func (x *ListSubscriptionResponse) Reset() {
	*x = ListSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionResponse) ProtoMessage() {}

func (x *ListSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *ListSubscriptionResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListSubscriptionResponse) GetWatchers() []*Subscription {
	if x != nil {
		return x.Watchers
	}
	return nil
}

type UpdateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subscription *Subscription         `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	UpdateMask   *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSubscriptionRequest) Reset() {
	*x = UpdateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionRequest) ProtoMessage() {}

func (x *UpdateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSubscriptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSubscriptionRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *UpdateSubscriptionRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_api_subscription_proto protoreflect.FileDescriptor

var file_api_subscription_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67,
	0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x50, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x48, 0x0a, 0x08, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xfd, 0x04, 0x0a, 0x13, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x67, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67,
	0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x77, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e,
	0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73,
	0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e,
	0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x67, 0x6f, 0x73, 0x6c, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_subscription_proto_rawDescOnce sync.Once
	file_api_subscription_proto_rawDescData = file_api_subscription_proto_rawDesc
)

func file_api_subscription_proto_rawDescGZIP() []byte {
	file_api_subscription_proto_rawDescOnce.Do(func() {
		file_api_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_subscription_proto_rawDescData)
	})
	return file_api_subscription_proto_rawDescData
}

var file_api_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_subscription_proto_goTypes = []interface{}{
	(*Subscription)(nil),              // 0: nokamoto.github.com.egosla.api.Subscription
	(*CreateSubscriptionRequest)(nil), // 1: nokamoto.github.com.egosla.api.CreateSubscriptionRequest
	(*DeleteSubscriptionRequest)(nil), // 2: nokamoto.github.com.egosla.api.DeleteSubscriptionRequest
	(*GetSubscriptionRequest)(nil),    // 3: nokamoto.github.com.egosla.api.GetSubscriptionRequest
	(*ListSubscriptionRequest)(nil),   // 4: nokamoto.github.com.egosla.api.ListSubscriptionRequest
	(*ListSubscriptionResponse)(nil),  // 5: nokamoto.github.com.egosla.api.ListSubscriptionResponse
	(*UpdateSubscriptionRequest)(nil), // 6: nokamoto.github.com.egosla.api.UpdateSubscriptionRequest
	(*field_mask.FieldMask)(nil),      // 7: google.protobuf.FieldMask
	(*empty.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_api_subscription_proto_depIdxs = []int32{
	0, // 0: nokamoto.github.com.egosla.api.CreateSubscriptionRequest.subscription:type_name -> nokamoto.github.com.egosla.api.Subscription
	0, // 1: nokamoto.github.com.egosla.api.ListSubscriptionResponse.watchers:type_name -> nokamoto.github.com.egosla.api.Subscription
	0, // 2: nokamoto.github.com.egosla.api.UpdateSubscriptionRequest.subscription:type_name -> nokamoto.github.com.egosla.api.Subscription
	7, // 3: nokamoto.github.com.egosla.api.UpdateSubscriptionRequest.update_mask:type_name -> google.protobuf.FieldMask
	1, // 4: nokamoto.github.com.egosla.api.SubscriptionService.CreateSubscription:input_type -> nokamoto.github.com.egosla.api.CreateSubscriptionRequest
	2, // 5: nokamoto.github.com.egosla.api.SubscriptionService.DeleteSubscription:input_type -> nokamoto.github.com.egosla.api.DeleteSubscriptionRequest
	3, // 6: nokamoto.github.com.egosla.api.SubscriptionService.GetSubscription:input_type -> nokamoto.github.com.egosla.api.GetSubscriptionRequest
	4, // 7: nokamoto.github.com.egosla.api.SubscriptionService.ListSubscription:input_type -> nokamoto.github.com.egosla.api.ListSubscriptionRequest
	6, // 8: nokamoto.github.com.egosla.api.SubscriptionService.UpdateSubscription:input_type -> nokamoto.github.com.egosla.api.UpdateSubscriptionRequest
	0, // 9: nokamoto.github.com.egosla.api.SubscriptionService.CreateSubscription:output_type -> nokamoto.github.com.egosla.api.Subscription
	8, // 10: nokamoto.github.com.egosla.api.SubscriptionService.DeleteSubscription:output_type -> google.protobuf.Empty
	0, // 11: nokamoto.github.com.egosla.api.SubscriptionService.GetSubscription:output_type -> nokamoto.github.com.egosla.api.Subscription
	5, // 12: nokamoto.github.com.egosla.api.SubscriptionService.ListSubscription:output_type -> nokamoto.github.com.egosla.api.ListSubscriptionResponse
	0, // 13: nokamoto.github.com.egosla.api.SubscriptionService.UpdateSubscription:output_type -> nokamoto.github.com.egosla.api.Subscription
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_subscription_proto_init() }
func file_api_subscription_proto_init() {
	if File_api_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_api_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubscriptionRequest); i {
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
		file_api_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSubscriptionRequest); i {
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
		file_api_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_api_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionRequest); i {
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
		file_api_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionResponse); i {
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
		file_api_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubscriptionRequest); i {
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
			RawDescriptor: file_api_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_subscription_proto_goTypes,
		DependencyIndexes: file_api_subscription_proto_depIdxs,
		MessageInfos:      file_api_subscription_proto_msgTypes,
	}.Build()
	File_api_subscription_proto = out.File
	file_api_subscription_proto_rawDesc = nil
	file_api_subscription_proto_goTypes = nil
	file_api_subscription_proto_depIdxs = nil
}
