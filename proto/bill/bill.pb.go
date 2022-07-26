// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/bill/bill.proto

package proto

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

type CreateBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Amount    int64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	BuyAt     int64   `protobuf:"varint,5,opt,name=buyAt,proto3" json:"buyAt,omitempty"`
}

func (x *CreateBillsRequest) Reset() {
	*x = CreateBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBillsRequest) ProtoMessage() {}

func (x *CreateBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBillsRequest.ProtoReflect.Descriptor instead.
func (*CreateBillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBillsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateBillsRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateBillsRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateBillsRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateBillsRequest) GetBuyAt() int64 {
	if x != nil {
		return x.BuyAt
	}
	return 0
}

type CreateBillsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBillsReply) Reset() {
	*x = CreateBillsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBillsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBillsReply) ProtoMessage() {}

func (x *CreateBillsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBillsReply.ProtoReflect.Descriptor instead.
func (*CreateBillsReply) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBillsReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBillsRequest) Reset() {
	*x = UpdateBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillsRequest) ProtoMessage() {}

func (x *UpdateBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillsRequest.ProtoReflect.Descriptor instead.
func (*UpdateBillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{2}
}

type UpdateBillsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBillsReply) Reset() {
	*x = UpdateBillsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillsReply) ProtoMessage() {}

func (x *UpdateBillsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillsReply.ProtoReflect.Descriptor instead.
func (*UpdateBillsReply) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{3}
}

type DeleteBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBillsRequest) Reset() {
	*x = DeleteBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillsRequest) ProtoMessage() {}

func (x *DeleteBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBillsRequest.ProtoReflect.Descriptor instead.
func (*DeleteBillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{4}
}

type DeleteBillsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBillsReply) Reset() {
	*x = DeleteBillsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillsReply) ProtoMessage() {}

func (x *DeleteBillsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBillsReply.ProtoReflect.Descriptor instead.
func (*DeleteBillsReply) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{5}
}

type GetBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetBillsRequest) Reset() {
	*x = GetBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillsRequest) ProtoMessage() {}

func (x *GetBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillsRequest.ProtoReflect.Descriptor instead.
func (*GetBillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{6}
}

func (x *GetBillsRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetBillsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBillsReply) Reset() {
	*x = GetBillsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBillsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillsReply) ProtoMessage() {}

func (x *GetBillsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillsReply.ProtoReflect.Descriptor instead.
func (*GetBillsReply) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{7}
}

type ListBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBillsRequest) Reset() {
	*x = ListBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillsRequest) ProtoMessage() {}

func (x *ListBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillsRequest.ProtoReflect.Descriptor instead.
func (*ListBillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{8}
}

type ListBillsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBillsReply) Reset() {
	*x = ListBillsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bill_bill_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillsReply) ProtoMessage() {}

func (x *ListBillsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bill_bill_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillsReply.ProtoReflect.Descriptor instead.
func (*ListBillsReply) Descriptor() ([]byte, []int) {
	return file_proto_bill_bill_proto_rawDescGZIP(), []int{9}
}

var File_proto_bill_bill_proto protoreflect.FileDescriptor

var file_proto_bill_bill_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x79,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x75, 0x79, 0x41, 0x74, 0x22,
	0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc7, 0x02, 0x0a, 0x05, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12,
	0x41, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c,
	0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x69, 0x6c, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42,
	0x69, 0x6c, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bill_bill_proto_rawDescOnce sync.Once
	file_proto_bill_bill_proto_rawDescData = file_proto_bill_bill_proto_rawDesc
)

func file_proto_bill_bill_proto_rawDescGZIP() []byte {
	file_proto_bill_bill_proto_rawDescOnce.Do(func() {
		file_proto_bill_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bill_bill_proto_rawDescData)
	})
	return file_proto_bill_bill_proto_rawDescData
}

var file_proto_bill_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_bill_bill_proto_goTypes = []interface{}{
	(*CreateBillsRequest)(nil), // 0: proto.CreateBillsRequest
	(*CreateBillsReply)(nil),   // 1: proto.CreateBillsReply
	(*UpdateBillsRequest)(nil), // 2: proto.UpdateBillsRequest
	(*UpdateBillsReply)(nil),   // 3: proto.UpdateBillsReply
	(*DeleteBillsRequest)(nil), // 4: proto.DeleteBillsRequest
	(*DeleteBillsReply)(nil),   // 5: proto.DeleteBillsReply
	(*GetBillsRequest)(nil),    // 6: proto.GetBillsRequest
	(*GetBillsReply)(nil),      // 7: proto.GetBillsReply
	(*ListBillsRequest)(nil),   // 8: proto.ListBillsRequest
	(*ListBillsReply)(nil),     // 9: proto.ListBillsReply
}
var file_proto_bill_bill_proto_depIdxs = []int32{
	0, // 0: proto.Bills.CreateBills:input_type -> proto.CreateBillsRequest
	2, // 1: proto.Bills.UpdateBills:input_type -> proto.UpdateBillsRequest
	4, // 2: proto.Bills.DeleteBills:input_type -> proto.DeleteBillsRequest
	6, // 3: proto.Bills.GetBills:input_type -> proto.GetBillsRequest
	8, // 4: proto.Bills.ListBills:input_type -> proto.ListBillsRequest
	1, // 5: proto.Bills.CreateBills:output_type -> proto.CreateBillsReply
	3, // 6: proto.Bills.UpdateBills:output_type -> proto.UpdateBillsReply
	5, // 7: proto.Bills.DeleteBills:output_type -> proto.DeleteBillsReply
	7, // 8: proto.Bills.GetBills:output_type -> proto.GetBillsReply
	9, // 9: proto.Bills.ListBills:output_type -> proto.ListBillsReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_bill_bill_proto_init() }
func file_proto_bill_bill_proto_init() {
	if File_proto_bill_bill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bill_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBillsRequest); i {
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
		file_proto_bill_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBillsReply); i {
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
		file_proto_bill_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillsRequest); i {
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
		file_proto_bill_bill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillsReply); i {
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
		file_proto_bill_bill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBillsRequest); i {
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
		file_proto_bill_bill_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBillsReply); i {
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
		file_proto_bill_bill_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBillsRequest); i {
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
		file_proto_bill_bill_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBillsReply); i {
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
		file_proto_bill_bill_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillsRequest); i {
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
		file_proto_bill_bill_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillsReply); i {
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
			RawDescriptor: file_proto_bill_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bill_bill_proto_goTypes,
		DependencyIndexes: file_proto_bill_bill_proto_depIdxs,
		MessageInfos:      file_proto_bill_bill_proto_msgTypes,
	}.Build()
	File_proto_bill_bill_proto = out.File
	file_proto_bill_bill_proto_rawDesc = nil
	file_proto_bill_bill_proto_goTypes = nil
	file_proto_bill_bill_proto_depIdxs = nil
}
