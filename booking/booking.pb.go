// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: booking.proto

package booking

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookingInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookingId     int32                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	ResourceId    int32                  `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	StartAt       string                 `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt         string                 `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookingInfo) Reset() {
	*x = BookingInfo{}
	mi := &file_booking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingInfo) ProtoMessage() {}

func (x *BookingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingInfo.ProtoReflect.Descriptor instead.
func (*BookingInfo) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingInfo) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingInfo) GetResourceId() int32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *BookingInfo) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *BookingInfo) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

type CreateBookingRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	UserId           int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId       int32                  `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceCapacity int32                  `protobuf:"varint,3,opt,name=resource_capacity,json=resourceCapacity,proto3" json:"resource_capacity,omitempty"`
	StartAt          string                 `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt            string                 `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	mi := &file_booking_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateBookingRequest) GetResourceId() int32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *CreateBookingRequest) GetResourceCapacity() int32 {
	if x != nil {
		return x.ResourceCapacity
	}
	return 0
}

func (x *CreateBookingRequest) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *CreateBookingRequest) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

type CreateBookingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookingId     int32                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookingResponse) Reset() {
	*x = CreateBookingResponse{}
	mi := &file_booking_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingResponse) ProtoMessage() {}

func (x *CreateBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingResponse.ProtoReflect.Descriptor instead.
func (*CreateBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookingResponse) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type GetBookingsByResourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    []int32                `protobuf:"varint,1,rep,packed,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookingsByResourceRequest) Reset() {
	*x = GetBookingsByResourceRequest{}
	mi := &file_booking_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookingsByResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingsByResourceRequest) ProtoMessage() {}

func (x *GetBookingsByResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingsByResourceRequest.ProtoReflect.Descriptor instead.
func (*GetBookingsByResourceRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookingsByResourceRequest) GetResourceId() []int32 {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

type GetBookingsByResourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookings      []*BookingInfo         `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookingsByResourceResponse) Reset() {
	*x = GetBookingsByResourceResponse{}
	mi := &file_booking_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookingsByResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingsByResourceResponse) ProtoMessage() {}

func (x *GetBookingsByResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingsByResourceResponse.ProtoReflect.Descriptor instead.
func (*GetBookingsByResourceResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookingsByResourceResponse) GetBookings() []*BookingInfo {
	if x != nil {
		return x.Bookings
	}
	return nil
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x7f, 0x0a, 0x0b,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22, 0xaf, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22,
	0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x32,
	0xd1, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x68, 0x65, 0x63, 0x68, 0x65, 0x6e, 0x49, 0x74, 0x7a, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData []byte
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_booking_proto_rawDesc), len(file_booking_proto_rawDesc)))
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_booking_proto_goTypes = []any{
	(*BookingInfo)(nil),                   // 0: booking_api.BookingInfo
	(*CreateBookingRequest)(nil),          // 1: booking_api.CreateBookingRequest
	(*CreateBookingResponse)(nil),         // 2: booking_api.CreateBookingResponse
	(*GetBookingsByResourceRequest)(nil),  // 3: booking_api.GetBookingsByResourceRequest
	(*GetBookingsByResourceResponse)(nil), // 4: booking_api.GetBookingsByResourceResponse
}
var file_booking_proto_depIdxs = []int32{
	0, // 0: booking_api.GetBookingsByResourceResponse.bookings:type_name -> booking_api.BookingInfo
	1, // 1: booking_api.Booking.CreateBooking:input_type -> booking_api.CreateBookingRequest
	3, // 2: booking_api.Booking.GetBookingsByResource:input_type -> booking_api.GetBookingsByResourceRequest
	2, // 3: booking_api.Booking.CreateBooking:output_type -> booking_api.CreateBookingResponse
	4, // 4: booking_api.Booking.GetBookingsByResource:output_type -> booking_api.GetBookingsByResourceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_booking_proto_rawDesc), len(file_booking_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
