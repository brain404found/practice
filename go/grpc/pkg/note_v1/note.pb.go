// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: note_v1/note.proto

package note_v1

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

type Note struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Note) Reset() {
	*x = Note{}
	mi := &file_note_v1_note_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Note) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Note) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateIn) Reset() {
	*x = CreateIn{}
	mi := &file_note_v1_note_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIn) ProtoMessage() {}

func (x *CreateIn) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIn.ProtoReflect.Descriptor instead.
func (*CreateIn) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIn) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateIn) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOut) Reset() {
	*x = CreateOut{}
	mi := &file_note_v1_note_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOut) ProtoMessage() {}

func (x *CreateOut) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOut.ProtoReflect.Descriptor instead.
func (*CreateOut) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOut) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type GetIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIn) Reset() {
	*x = GetIn{}
	mi := &file_note_v1_note_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIn) ProtoMessage() {}

func (x *GetIn) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIn.ProtoReflect.Descriptor instead.
func (*GetIn) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{3}
}

func (x *GetIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOut) Reset() {
	*x = GetOut{}
	mi := &file_note_v1_note_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOut) ProtoMessage() {}

func (x *GetOut) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOut.ProtoReflect.Descriptor instead.
func (*GetOut) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{4}
}

func (x *GetOut) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type ListIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListIn) Reset() {
	*x = ListIn{}
	mi := &file_note_v1_note_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIn) ProtoMessage() {}

func (x *ListIn) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIn.ProtoReflect.Descriptor instead.
func (*ListIn) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{5}
}

type ListOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notes         []*Note                `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOut) Reset() {
	*x = ListOut{}
	mi := &file_note_v1_note_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOut) ProtoMessage() {}

func (x *ListOut) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOut.ProtoReflect.Descriptor instead.
func (*ListOut) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{6}
}

func (x *ListOut) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type UpdateIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateIn) Reset() {
	*x = UpdateIn{}
	mi := &file_note_v1_note_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIn) ProtoMessage() {}

func (x *UpdateIn) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIn.ProtoReflect.Descriptor instead.
func (*UpdateIn) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateIn) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateIn) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOut) Reset() {
	*x = UpdateOut{}
	mi := &file_note_v1_note_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOut) ProtoMessage() {}

func (x *UpdateOut) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOut.ProtoReflect.Descriptor instead.
func (*UpdateOut) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOut) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type DeleteIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteIn) Reset() {
	*x = DeleteIn{}
	mi := &file_note_v1_note_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIn) ProtoMessage() {}

func (x *DeleteIn) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIn.ProtoReflect.Descriptor instead.
func (*DeleteIn) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOut) Reset() {
	*x = DeleteOut{}
	mi := &file_note_v1_note_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOut) ProtoMessage() {}

func (x *DeleteOut) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_note_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOut.ProtoReflect.Descriptor instead.
func (*DeleteOut) Descriptor() ([]byte, []int) {
	return file_note_v1_note_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_note_v1_note_proto protoreflect.FileDescriptor

const file_note_v1_note_proto_rawDesc = "" +
	"\n" +
	"\x12note_v1/note.proto\x12\anote_v1\"\x84\x01\n" +
	"\x04Note\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\tR\tupdatedAt\":\n" +
	"\bCreateIn\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\".\n" +
	"\tCreateOut\x12!\n" +
	"\x04note\x18\x01 \x01(\v2\r.note_v1.NoteR\x04note\"\x17\n" +
	"\x05GetIn\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"+\n" +
	"\x06GetOut\x12!\n" +
	"\x04note\x18\x01 \x01(\v2\r.note_v1.NoteR\x04note\"\b\n" +
	"\x06ListIn\".\n" +
	"\aListOut\x12#\n" +
	"\x05notes\x18\x01 \x03(\v2\r.note_v1.NoteR\x05notes\"J\n" +
	"\bUpdateIn\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\".\n" +
	"\tUpdateOut\x12!\n" +
	"\x04note\x18\x01 \x01(\v2\r.note_v1.NoteR\x04note\"\x1a\n" +
	"\bDeleteIn\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"%\n" +
	"\tDeleteOut\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xf3\x01\n" +
	"\vNoteService\x12/\n" +
	"\x06Create\x12\x11.note_v1.CreateIn\x1a\x12.note_v1.CreateOut\x12&\n" +
	"\x03Get\x12\x0e.note_v1.GetIn\x1a\x0f.note_v1.GetOut\x12)\n" +
	"\x04List\x12\x0f.note_v1.ListIn\x1a\x10.note_v1.ListOut\x12/\n" +
	"\x06Update\x12\x11.note_v1.UpdateIn\x1a\x12.note_v1.UpdateOut\x12/\n" +
	"\x06Delete\x12\x11.note_v1.DeleteIn\x1a\x12.note_v1.DeleteOutB\x19Z\x17simple-grpc/pkg;note_v1b\x06proto3"

var (
	file_note_v1_note_proto_rawDescOnce sync.Once
	file_note_v1_note_proto_rawDescData []byte
)

func file_note_v1_note_proto_rawDescGZIP() []byte {
	file_note_v1_note_proto_rawDescOnce.Do(func() {
		file_note_v1_note_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_note_v1_note_proto_rawDesc), len(file_note_v1_note_proto_rawDesc)))
	})
	return file_note_v1_note_proto_rawDescData
}

var file_note_v1_note_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_note_v1_note_proto_goTypes = []any{
	(*Note)(nil),      // 0: note_v1.Note
	(*CreateIn)(nil),  // 1: note_v1.CreateIn
	(*CreateOut)(nil), // 2: note_v1.CreateOut
	(*GetIn)(nil),     // 3: note_v1.GetIn
	(*GetOut)(nil),    // 4: note_v1.GetOut
	(*ListIn)(nil),    // 5: note_v1.ListIn
	(*ListOut)(nil),   // 6: note_v1.ListOut
	(*UpdateIn)(nil),  // 7: note_v1.UpdateIn
	(*UpdateOut)(nil), // 8: note_v1.UpdateOut
	(*DeleteIn)(nil),  // 9: note_v1.DeleteIn
	(*DeleteOut)(nil), // 10: note_v1.DeleteOut
}
var file_note_v1_note_proto_depIdxs = []int32{
	0,  // 0: note_v1.CreateOut.note:type_name -> note_v1.Note
	0,  // 1: note_v1.GetOut.note:type_name -> note_v1.Note
	0,  // 2: note_v1.ListOut.notes:type_name -> note_v1.Note
	0,  // 3: note_v1.UpdateOut.note:type_name -> note_v1.Note
	1,  // 4: note_v1.NoteService.Create:input_type -> note_v1.CreateIn
	3,  // 5: note_v1.NoteService.Get:input_type -> note_v1.GetIn
	5,  // 6: note_v1.NoteService.List:input_type -> note_v1.ListIn
	7,  // 7: note_v1.NoteService.Update:input_type -> note_v1.UpdateIn
	9,  // 8: note_v1.NoteService.Delete:input_type -> note_v1.DeleteIn
	2,  // 9: note_v1.NoteService.Create:output_type -> note_v1.CreateOut
	4,  // 10: note_v1.NoteService.Get:output_type -> note_v1.GetOut
	6,  // 11: note_v1.NoteService.List:output_type -> note_v1.ListOut
	8,  // 12: note_v1.NoteService.Update:output_type -> note_v1.UpdateOut
	10, // 13: note_v1.NoteService.Delete:output_type -> note_v1.DeleteOut
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_note_v1_note_proto_init() }
func file_note_v1_note_proto_init() {
	if File_note_v1_note_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_note_v1_note_proto_rawDesc), len(file_note_v1_note_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_note_v1_note_proto_goTypes,
		DependencyIndexes: file_note_v1_note_proto_depIdxs,
		MessageInfos:      file_note_v1_note_proto_msgTypes,
	}.Build()
	File_note_v1_note_proto = out.File
	file_note_v1_note_proto_goTypes = nil
	file_note_v1_note_proto_depIdxs = nil
}
