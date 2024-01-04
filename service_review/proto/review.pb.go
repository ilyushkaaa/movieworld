// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: review.proto

package review

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

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type FilmID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *FilmID) Reset() {
	*x = FilmID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmID) ProtoMessage() {}

func (x *FilmID) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmID.ProtoReflect.Descriptor instead.
func (*FilmID) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *FilmID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReviewID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReviewID) Reset() {
	*x = ReviewID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewID) ProtoMessage() {}

func (x *ReviewID) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewID.ProtoReflect.Descriptor instead.
func (*ReviewID) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type IsDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDeleted bool `protobuf:"varint,1,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
}

func (x *IsDeleted) Reset() {
	*x = IsDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsDeleted) ProtoMessage() {}

func (x *IsDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsDeleted.ProtoReflect.Descriptor instead.
func (*IsDeleted) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *IsDeleted) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      *ReviewID `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Mark    uint32    `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	Comment string    `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{4}
}

func (x *Review) GetID() *ReviewID {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Review) GetMark() uint32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *Review) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Reviews struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *Reviews) Reset() {
	*x = Reviews{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reviews) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reviews) ProtoMessage() {}

func (x *Reviews) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reviews.ProtoReflect.Descriptor instead.
func (*Reviews) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{5}
}

func (x *Reviews) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type NewReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Review *Review `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	FilmID *FilmID `protobuf:"bytes,2,opt,name=filmID,proto3" json:"filmID,omitempty"`
	UserID *UserID `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *NewReviewData) Reset() {
	*x = NewReviewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReviewData) ProtoMessage() {}

func (x *NewReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReviewData.ProtoReflect.Descriptor instead.
func (*NewReviewData) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{6}
}

func (x *NewReviewData) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

func (x *NewReviewData) GetFilmID() *FilmID {
	if x != nil {
		return x.FilmID
	}
	return nil
}

func (x *NewReviewData) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

type DeleteReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewID *ReviewID `protobuf:"bytes,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	UserID   *UserID   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *DeleteReviewData) Reset() {
	*x = DeleteReviewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewData) ProtoMessage() {}

func (x *DeleteReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewData.ProtoReflect.Descriptor instead.
func (*DeleteReviewData) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteReviewData) GetReviewID() *ReviewID {
	if x != nil {
		return x.ReviewID
	}
	return nil
}

func (x *DeleteReviewData) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

type UpdateReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Review *Review `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	UserID *UserID `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UpdateReviewData) Reset() {
	*x = UpdateReviewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewData) ProtoMessage() {}

func (x *UpdateReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewData.ProtoReflect.Descriptor instead.
func (*UpdateReviewData) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateReviewData) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

func (x *UpdateReviewData) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

var File_review_proto protoreflect.FileDescriptor

var file_review_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x18, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1a, 0x0a, 0x08, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x22, 0x58, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x07, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x22, 0x87, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x44, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x62, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x26, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x32, 0xeb, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x6d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x0e, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x18, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x18, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData = file_review_proto_rawDesc
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_proto_rawDescData)
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_review_proto_goTypes = []interface{}{
	(*UserID)(nil),           // 0: review.UserID
	(*FilmID)(nil),           // 1: review.FilmID
	(*ReviewID)(nil),         // 2: review.ReviewID
	(*IsDeleted)(nil),        // 3: review.IsDeleted
	(*Review)(nil),           // 4: review.Review
	(*Reviews)(nil),          // 5: review.Reviews
	(*NewReviewData)(nil),    // 6: review.NewReviewData
	(*DeleteReviewData)(nil), // 7: review.DeleteReviewData
	(*UpdateReviewData)(nil), // 8: review.UpdateReviewData
}
var file_review_proto_depIdxs = []int32{
	2,  // 0: review.Review.ID:type_name -> review.ReviewID
	4,  // 1: review.Reviews.reviews:type_name -> review.Review
	4,  // 2: review.NewReviewData.review:type_name -> review.Review
	1,  // 3: review.NewReviewData.filmID:type_name -> review.FilmID
	0,  // 4: review.NewReviewData.userID:type_name -> review.UserID
	2,  // 5: review.DeleteReviewData.reviewID:type_name -> review.ReviewID
	0,  // 6: review.DeleteReviewData.userID:type_name -> review.UserID
	4,  // 7: review.UpdateReviewData.review:type_name -> review.Review
	0,  // 8: review.UpdateReviewData.userID:type_name -> review.UserID
	1,  // 9: review.ReviewMaker.GetFilmReviews:input_type -> review.FilmID
	6,  // 10: review.ReviewMaker.NewReview:input_type -> review.NewReviewData
	7,  // 11: review.ReviewMaker.DeleteReview:input_type -> review.DeleteReviewData
	8,  // 12: review.ReviewMaker.UpdateReview:input_type -> review.UpdateReviewData
	5,  // 13: review.ReviewMaker.GetFilmReviews:output_type -> review.Reviews
	4,  // 14: review.ReviewMaker.NewReview:output_type -> review.Review
	3,  // 15: review.ReviewMaker.DeleteReview:output_type -> review.IsDeleted
	4,  // 16: review.ReviewMaker.UpdateReview:output_type -> review.Review
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmID); i {
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
		file_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewID); i {
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
		file_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsDeleted); i {
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
		file_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reviews); i {
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
		file_review_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReviewData); i {
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
		file_review_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReviewData); i {
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
		file_review_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReviewData); i {
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
			RawDescriptor: file_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_rawDesc = nil
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}
