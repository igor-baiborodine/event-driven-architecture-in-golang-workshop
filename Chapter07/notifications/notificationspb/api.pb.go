// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: notificationspb/api.proto

package notificationspb

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

type NotifyOrderCreatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderCreatedRequest) Reset() {
	*x = NotifyOrderCreatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderCreatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCreatedRequest) ProtoMessage() {}

func (x *NotifyOrderCreatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCreatedRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderCreatedRequest) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyOrderCreatedRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderCreatedRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderCreatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderCreatedResponse) Reset() {
	*x = NotifyOrderCreatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderCreatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCreatedResponse) ProtoMessage() {}

func (x *NotifyOrderCreatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCreatedResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderCreatedResponse) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{1}
}

type NotifyOrderCanceledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderCanceledRequest) Reset() {
	*x = NotifyOrderCanceledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderCanceledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCanceledRequest) ProtoMessage() {}

func (x *NotifyOrderCanceledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCanceledRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderCanceledRequest) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyOrderCanceledRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderCanceledRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderCanceledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderCanceledResponse) Reset() {
	*x = NotifyOrderCanceledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderCanceledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCanceledResponse) ProtoMessage() {}

func (x *NotifyOrderCanceledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCanceledResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderCanceledResponse) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{3}
}

type NotifyOrderReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderReadyRequest) Reset() {
	*x = NotifyOrderReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderReadyRequest) ProtoMessage() {}

func (x *NotifyOrderReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderReadyRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderReadyRequest) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyOrderReadyRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderReadyRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderReadyResponse) Reset() {
	*x = NotifyOrderReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationspb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOrderReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderReadyResponse) ProtoMessage() {}

func (x *NotifyOrderReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationspb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderReadyResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderReadyResponse) Descriptor() ([]byte, []int) {
	return file_notificationspb_api_proto_rawDescGZIP(), []int{5}
}

var File_notificationspb_api_proto protoreflect.FileDescriptor

var file_notificationspb_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x1a, 0x1e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x19,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1d, 0x0a,
	0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x17,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xe6, 0x02, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x13, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x12, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a,
	0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x12, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb8, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62,
	0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x65, 0x64,
	0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa,
	0x02, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70,
	0x62, 0xca, 0x02, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0x62, 0xe2, 0x02, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notificationspb_api_proto_rawDescOnce sync.Once
	file_notificationspb_api_proto_rawDescData = file_notificationspb_api_proto_rawDesc
)

func file_notificationspb_api_proto_rawDescGZIP() []byte {
	file_notificationspb_api_proto_rawDescOnce.Do(func() {
		file_notificationspb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_notificationspb_api_proto_rawDescData)
	})
	return file_notificationspb_api_proto_rawDescData
}

var file_notificationspb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notificationspb_api_proto_goTypes = []interface{}{
	(*NotifyOrderCreatedRequest)(nil),   // 0: notificationspb.NotifyOrderCreatedRequest
	(*NotifyOrderCreatedResponse)(nil),  // 1: notificationspb.NotifyOrderCreatedResponse
	(*NotifyOrderCanceledRequest)(nil),  // 2: notificationspb.NotifyOrderCanceledRequest
	(*NotifyOrderCanceledResponse)(nil), // 3: notificationspb.NotifyOrderCanceledResponse
	(*NotifyOrderReadyRequest)(nil),     // 4: notificationspb.NotifyOrderReadyRequest
	(*NotifyOrderReadyResponse)(nil),    // 5: notificationspb.NotifyOrderReadyResponse
}
var file_notificationspb_api_proto_depIdxs = []int32{
	0, // 0: notificationspb.NotificationsService.NotifyOrderCreated:input_type -> notificationspb.NotifyOrderCreatedRequest
	2, // 1: notificationspb.NotificationsService.NotifyOrderCanceled:input_type -> notificationspb.NotifyOrderCanceledRequest
	4, // 2: notificationspb.NotificationsService.NotifyOrderReady:input_type -> notificationspb.NotifyOrderReadyRequest
	1, // 3: notificationspb.NotificationsService.NotifyOrderCreated:output_type -> notificationspb.NotifyOrderCreatedResponse
	3, // 4: notificationspb.NotificationsService.NotifyOrderCanceled:output_type -> notificationspb.NotifyOrderCanceledResponse
	5, // 5: notificationspb.NotificationsService.NotifyOrderReady:output_type -> notificationspb.NotifyOrderReadyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notificationspb_api_proto_init() }
func file_notificationspb_api_proto_init() {
	if File_notificationspb_api_proto != nil {
		return
	}
	file_notificationspb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notificationspb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderCreatedRequest); i {
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
		file_notificationspb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderCreatedResponse); i {
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
		file_notificationspb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderCanceledRequest); i {
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
		file_notificationspb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderCanceledResponse); i {
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
		file_notificationspb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderReadyRequest); i {
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
		file_notificationspb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOrderReadyResponse); i {
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
			RawDescriptor: file_notificationspb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notificationspb_api_proto_goTypes,
		DependencyIndexes: file_notificationspb_api_proto_depIdxs,
		MessageInfos:      file_notificationspb_api_proto_msgTypes,
	}.Build()
	File_notificationspb_api_proto = out.File
	file_notificationspb_api_proto_rawDesc = nil
	file_notificationspb_api_proto_goTypes = nil
	file_notificationspb_api_proto_depIdxs = nil
}
