// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: doctors/doctors.proto

package doctors

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Doctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Surname    string `protobuf:"bytes,3,opt,name=Surname,json=surname,proto3" json:"Surname,omitempty"`
	Position   string `protobuf:"bytes,4,opt,name=Position,json=position,proto3" json:"Position,omitempty"`
	Age        uint32 `protobuf:"varint,5,opt,name=Age,json=age,proto3" json:"Age,omitempty"`
	Experience int32  `protobuf:"varint,6,opt,name=Experience,json=experience,proto3" json:"Experience,omitempty"`
}

func (x *Doctor) Reset() {
	*x = Doctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doctor) ProtoMessage() {}

func (x *Doctor) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doctor.ProtoReflect.Descriptor instead.
func (*Doctor) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{0}
}

func (x *Doctor) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Doctor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doctor) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Doctor) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Doctor) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Doctor) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type CreateDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *CreateDoctorRequest) Reset() {
	*x = CreateDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoctorRequest) ProtoMessage() {}

func (x *CreateDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoctorRequest.ProtoReflect.Descriptor instead.
func (*CreateDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDoctorRequest) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type CreateDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *CreateDoctorResponse) Reset() {
	*x = CreateDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoctorResponse) ProtoMessage() {}

func (x *CreateDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoctorResponse.ProtoReflect.Descriptor instead.
func (*CreateDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type GetDoctorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDoctorsRequest) Reset() {
	*x = GetDoctorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorsRequest) ProtoMessage() {}

func (x *GetDoctorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorsRequest.ProtoReflect.Descriptor instead.
func (*GetDoctorsRequest) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{3}
}

type GetDoctorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctors *Doctor `protobuf:"bytes,1,opt,name=doctors,proto3" json:"doctors,omitempty"`
}

func (x *GetDoctorsResponse) Reset() {
	*x = GetDoctorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorsResponse) ProtoMessage() {}

func (x *GetDoctorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorsResponse.ProtoReflect.Descriptor instead.
func (*GetDoctorsResponse) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{4}
}

func (x *GetDoctorsResponse) GetDoctors() *Doctor {
	if x != nil {
		return x.Doctors
	}
	return nil
}

type GetDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDoctorRequest) Reset() {
	*x = GetDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorRequest) ProtoMessage() {}

func (x *GetDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorRequest.ProtoReflect.Descriptor instead.
func (*GetDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{5}
}

func (x *GetDoctorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *GetDoctorResponse) Reset() {
	*x = GetDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorResponse) ProtoMessage() {}

func (x *GetDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorResponse.ProtoReflect.Descriptor instead.
func (*GetDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{6}
}

func (x *GetDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type UpdateDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Doctor *Doctor `protobuf:"bytes,2,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *UpdateDoctorRequest) Reset() {
	*x = UpdateDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDoctorRequest) ProtoMessage() {}

func (x *UpdateDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDoctorRequest.ProtoReflect.Descriptor instead.
func (*UpdateDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDoctorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDoctorRequest) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type UpdateDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *UpdateDoctorResponse) Reset() {
	*x = UpdateDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDoctorResponse) ProtoMessage() {}

func (x *UpdateDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDoctorResponse.ProtoReflect.Descriptor instead.
func (*UpdateDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type DeleteDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDoctorRequest) Reset() {
	*x = DeleteDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDoctorRequest) ProtoMessage() {}

func (x *DeleteDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDoctorRequest.ProtoReflect.Descriptor instead.
func (*DeleteDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteDoctorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *DeleteDoctorResponse) Reset() {
	*x = DeleteDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_doctors_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDoctorResponse) ProtoMessage() {}

func (x *DeleteDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_doctors_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDoctorResponse.ProtoReflect.Descriptor instead.
func (*DeleteDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctors_doctors_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

var File_doctors_doctors_proto protoreflect.FileDescriptor

var file_doctors_doctors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x22, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x4e,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x3f,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x32, 0xf9, 0x03, 0x0a, 0x07, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x63, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x68, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x79, 0x65, 0x72, 0x7a, 0x68, 0x61, 0x6e, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x3b, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doctors_doctors_proto_rawDescOnce sync.Once
	file_doctors_doctors_proto_rawDescData = file_doctors_doctors_proto_rawDesc
)

func file_doctors_doctors_proto_rawDescGZIP() []byte {
	file_doctors_doctors_proto_rawDescOnce.Do(func() {
		file_doctors_doctors_proto_rawDescData = protoimpl.X.CompressGZIP(file_doctors_doctors_proto_rawDescData)
	})
	return file_doctors_doctors_proto_rawDescData
}

var file_doctors_doctors_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_doctors_doctors_proto_goTypes = []interface{}{
	(*Doctor)(nil),               // 0: doctors.Doctor
	(*CreateDoctorRequest)(nil),  // 1: doctors.CreateDoctorRequest
	(*CreateDoctorResponse)(nil), // 2: doctors.CreateDoctorResponse
	(*GetDoctorsRequest)(nil),    // 3: doctors.GetDoctorsRequest
	(*GetDoctorsResponse)(nil),   // 4: doctors.GetDoctorsResponse
	(*GetDoctorRequest)(nil),     // 5: doctors.GetDoctorRequest
	(*GetDoctorResponse)(nil),    // 6: doctors.GetDoctorResponse
	(*UpdateDoctorRequest)(nil),  // 7: doctors.UpdateDoctorRequest
	(*UpdateDoctorResponse)(nil), // 8: doctors.UpdateDoctorResponse
	(*DeleteDoctorRequest)(nil),  // 9: doctors.DeleteDoctorRequest
	(*DeleteDoctorResponse)(nil), // 10: doctors.DeleteDoctorResponse
}
var file_doctors_doctors_proto_depIdxs = []int32{
	0,  // 0: doctors.CreateDoctorRequest.doctor:type_name -> doctors.Doctor
	0,  // 1: doctors.CreateDoctorResponse.doctor:type_name -> doctors.Doctor
	0,  // 2: doctors.GetDoctorsResponse.doctors:type_name -> doctors.Doctor
	0,  // 3: doctors.GetDoctorResponse.doctor:type_name -> doctors.Doctor
	0,  // 4: doctors.UpdateDoctorRequest.doctor:type_name -> doctors.Doctor
	0,  // 5: doctors.UpdateDoctorResponse.doctor:type_name -> doctors.Doctor
	0,  // 6: doctors.DeleteDoctorResponse.doctor:type_name -> doctors.Doctor
	1,  // 7: doctors.Doctors.CreateDoctor:input_type -> doctors.CreateDoctorRequest
	3,  // 8: doctors.Doctors.GetDoctors:input_type -> doctors.GetDoctorsRequest
	5,  // 9: doctors.Doctors.GetDoctor:input_type -> doctors.GetDoctorRequest
	7,  // 10: doctors.Doctors.UpdateDoctor:input_type -> doctors.UpdateDoctorRequest
	9,  // 11: doctors.Doctors.DeleteDoctor:input_type -> doctors.DeleteDoctorRequest
	2,  // 12: doctors.Doctors.CreateDoctor:output_type -> doctors.CreateDoctorResponse
	4,  // 13: doctors.Doctors.GetDoctors:output_type -> doctors.GetDoctorsResponse
	6,  // 14: doctors.Doctors.GetDoctor:output_type -> doctors.GetDoctorResponse
	8,  // 15: doctors.Doctors.UpdateDoctor:output_type -> doctors.UpdateDoctorResponse
	10, // 16: doctors.Doctors.DeleteDoctor:output_type -> doctors.DeleteDoctorResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_doctors_doctors_proto_init() }
func file_doctors_doctors_proto_init() {
	if File_doctors_doctors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doctors_doctors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doctor); i {
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
		file_doctors_doctors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoctorRequest); i {
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
		file_doctors_doctors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoctorResponse); i {
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
		file_doctors_doctors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorsRequest); i {
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
		file_doctors_doctors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorsResponse); i {
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
		file_doctors_doctors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorRequest); i {
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
		file_doctors_doctors_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorResponse); i {
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
		file_doctors_doctors_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDoctorRequest); i {
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
		file_doctors_doctors_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDoctorResponse); i {
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
		file_doctors_doctors_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDoctorRequest); i {
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
		file_doctors_doctors_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDoctorResponse); i {
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
			RawDescriptor: file_doctors_doctors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doctors_doctors_proto_goTypes,
		DependencyIndexes: file_doctors_doctors_proto_depIdxs,
		MessageInfos:      file_doctors_doctors_proto_msgTypes,
	}.Build()
	File_doctors_doctors_proto = out.File
	file_doctors_doctors_proto_rawDesc = nil
	file_doctors_doctors_proto_goTypes = nil
	file_doctors_doctors_proto_depIdxs = nil
}
