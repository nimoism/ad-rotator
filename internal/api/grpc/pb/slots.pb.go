// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: api/slots.proto

package pb

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

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{0}
}

func (x *Slot) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Slot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AllSlotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllSlotsRequest) Reset() {
	*x = AllSlotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSlotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSlotsRequest) ProtoMessage() {}

func (x *AllSlotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSlotsRequest.ProtoReflect.Descriptor instead.
func (*AllSlotsRequest) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{1}
}

type AllSlotsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slots []*Slot `protobuf:"bytes,1,rep,name=slots,proto3" json:"slots,omitempty"`
}

func (x *AllSlotsResult) Reset() {
	*x = AllSlotsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSlotsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSlotsResult) ProtoMessage() {}

func (x *AllSlotsResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSlotsResult.ProtoReflect.Descriptor instead.
func (*AllSlotsResult) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{2}
}

func (x *AllSlotsResult) GetSlots() []*Slot {
	if x != nil {
		return x.Slots
	}
	return nil
}

type CreateSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot *CreateSlotRequest_NewSlot `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *CreateSlotRequest) Reset() {
	*x = CreateSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlotRequest) ProtoMessage() {}

func (x *CreateSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlotRequest.ProtoReflect.Descriptor instead.
func (*CreateSlotRequest) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSlotRequest) GetSlot() *CreateSlotRequest_NewSlot {
	if x != nil {
		return x.Slot
	}
	return nil
}

type CreateSlotResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot *Slot `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *CreateSlotResult) Reset() {
	*x = CreateSlotResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlotResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlotResult) ProtoMessage() {}

func (x *CreateSlotResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlotResult.ProtoReflect.Descriptor instead.
func (*CreateSlotResult) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSlotResult) GetSlot() *Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

type UpdateSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot *Slot `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *UpdateSlotRequest) Reset() {
	*x = UpdateSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSlotRequest) ProtoMessage() {}

func (x *UpdateSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSlotRequest.ProtoReflect.Descriptor instead.
func (*UpdateSlotRequest) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSlotRequest) GetSlot() *Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

type UpdateSlotResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot *Slot `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *UpdateSlotResult) Reset() {
	*x = UpdateSlotResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSlotResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSlotResult) ProtoMessage() {}

func (x *UpdateSlotResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSlotResult.ProtoReflect.Descriptor instead.
func (*UpdateSlotResult) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSlotResult) GetSlot() *Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

type DeleteSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSlotRequest) Reset() {
	*x = DeleteSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSlotRequest) ProtoMessage() {}

func (x *DeleteSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSlotRequest.ProtoReflect.Descriptor instead.
func (*DeleteSlotRequest) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSlotRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSlotResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSlotResult) Reset() {
	*x = DeleteSlotResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSlotResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSlotResult) ProtoMessage() {}

func (x *DeleteSlotResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSlotResult.ProtoReflect.Descriptor instead.
func (*DeleteSlotResult) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{8}
}

type CreateSlotRequest_NewSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSlotRequest_NewSlot) Reset() {
	*x = CreateSlotRequest_NewSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slots_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlotRequest_NewSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlotRequest_NewSlot) ProtoMessage() {}

func (x *CreateSlotRequest_NewSlot) ProtoReflect() protoreflect.Message {
	mi := &file_api_slots_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlotRequest_NewSlot.ProtoReflect.Descriptor instead.
func (*CreateSlotRequest_NewSlot) Descriptor() ([]byte, []int) {
	return file_api_slots_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CreateSlotRequest_NewSlot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_slots_proto protoreflect.FileDescriptor

var file_api_slots_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f,
	0x74, 0x1a, 0x1d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x31, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x31, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0xfa, 0x01, 0x0a, 0x05, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x34, 0x0a,
	0x05, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x6c,
	0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74,
	0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_slots_proto_rawDescOnce sync.Once
	file_api_slots_proto_rawDescData = file_api_slots_proto_rawDesc
)

func file_api_slots_proto_rawDescGZIP() []byte {
	file_api_slots_proto_rawDescOnce.Do(func() {
		file_api_slots_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_slots_proto_rawDescData)
	})
	return file_api_slots_proto_rawDescData
}

var file_api_slots_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_slots_proto_goTypes = []interface{}{
	(*Slot)(nil),                      // 0: api.Slot
	(*AllSlotsRequest)(nil),           // 1: api.AllSlotsRequest
	(*AllSlotsResult)(nil),            // 2: api.AllSlotsResult
	(*CreateSlotRequest)(nil),         // 3: api.CreateSlotRequest
	(*CreateSlotResult)(nil),          // 4: api.CreateSlotResult
	(*UpdateSlotRequest)(nil),         // 5: api.UpdateSlotRequest
	(*UpdateSlotResult)(nil),          // 6: api.UpdateSlotResult
	(*DeleteSlotRequest)(nil),         // 7: api.DeleteSlotRequest
	(*DeleteSlotResult)(nil),          // 8: api.DeleteSlotResult
	(*CreateSlotRequest_NewSlot)(nil), // 9: api.CreateSlotRequest.NewSlot
}
var file_api_slots_proto_depIdxs = []int32{
	0, // 0: api.AllSlotsResult.slots:type_name -> api.Slot
	9, // 1: api.CreateSlotRequest.slot:type_name -> api.CreateSlotRequest.NewSlot
	0, // 2: api.CreateSlotResult.slot:type_name -> api.Slot
	0, // 3: api.UpdateSlotRequest.slot:type_name -> api.Slot
	0, // 4: api.UpdateSlotResult.slot:type_name -> api.Slot
	1, // 5: api.Slots.Slots:input_type -> api.AllSlotsRequest
	3, // 6: api.Slots.CreateSlot:input_type -> api.CreateSlotRequest
	5, // 7: api.Slots.UpdateSlot:input_type -> api.UpdateSlotRequest
	7, // 8: api.Slots.DeleteSlot:input_type -> api.DeleteSlotRequest
	2, // 9: api.Slots.Slots:output_type -> api.AllSlotsResult
	4, // 10: api.Slots.CreateSlot:output_type -> api.CreateSlotResult
	6, // 11: api.Slots.UpdateSlot:output_type -> api.UpdateSlotResult
	8, // 12: api.Slots.DeleteSlot:output_type -> api.DeleteSlotResult
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_slots_proto_init() }
func file_api_slots_proto_init() {
	if File_api_slots_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_slots_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
		file_api_slots_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSlotsRequest); i {
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
		file_api_slots_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSlotsResult); i {
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
		file_api_slots_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlotRequest); i {
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
		file_api_slots_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlotResult); i {
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
		file_api_slots_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSlotRequest); i {
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
		file_api_slots_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSlotResult); i {
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
		file_api_slots_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSlotRequest); i {
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
		file_api_slots_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSlotResult); i {
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
		file_api_slots_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlotRequest_NewSlot); i {
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
			RawDescriptor: file_api_slots_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_slots_proto_goTypes,
		DependencyIndexes: file_api_slots_proto_depIdxs,
		MessageInfos:      file_api_slots_proto_msgTypes,
	}.Build()
	File_api_slots_proto = out.File
	file_api_slots_proto_rawDesc = nil
	file_api_slots_proto_goTypes = nil
	file_api_slots_proto_depIdxs = nil
}
