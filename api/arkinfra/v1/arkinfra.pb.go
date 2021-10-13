// TODO: 学习谷歌api设计规范，重新定义一下api与命名
// TODO: 为什么要有UserID？不放在元数据里吗？寻找微服务间调用的最佳实践

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: arkinfra.proto

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

// 一般来讲应该是调用user服务的pb文件
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateInfraReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateInfraReply) Reset() {
	*x = CreateInfraReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInfraReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInfraReply) ProtoMessage() {}

func (x *CreateInfraReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInfraReply.ProtoReflect.Descriptor instead.
func (*CreateInfraReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{1}
}

// 首页获取基建产出数目、问题数目的返回值
type GetInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductAmount int32 `protobuf:"varint,1,opt,name=product_amount,json=productAmount,proto3" json:"product_amount,omitempty"`
	ErrorAmount   int32 `protobuf:"varint,2,opt,name=error_amount,json=errorAmount,proto3" json:"error_amount,omitempty"`
}

func (x *GetInfoReply) Reset() {
	*x = GetInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReply) ProtoMessage() {}

func (x *GetInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoReply.ProtoReflect.Descriptor instead.
func (*GetInfoReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoReply) GetProductAmount() int32 {
	if x != nil {
		return x.ProductAmount
	}
	return 0
}

func (x *GetInfoReply) GetErrorAmount() int32 {
	if x != nil {
		return x.ErrorAmount
	}
	return 0
}

// 房间详细信息
type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId       int32    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomType     int32    `protobuf:"varint,2,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	RoomLevel    int32    `protobuf:"varint,3,opt,name=room_level,json=roomLevel,proto3" json:"room_level,omitempty"`
	CharacterId  []string `protobuf:"bytes,4,rep,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	CharacterFtg []int32  `protobuf:"varint,5,rep,packed,name=character_ftg,json=characterFtg,proto3" json:"character_ftg,omitempty"`
	Storage      int32    `protobuf:"varint,6,opt,name=storage,proto3" json:"storage,omitempty"`
	StorageUsed  int32    `protobuf:"varint,7,opt,name=storage_used,json=storageUsed,proto3" json:"storage_used,omitempty"`
	ItemId       int32    `protobuf:"varint,8,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemProgress int32    `protobuf:"varint,9,opt,name=item_progress,json=itemProgress,proto3" json:"item_progress,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{3}
}

func (x *Room) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *Room) GetRoomType() int32 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

func (x *Room) GetRoomLevel() int32 {
	if x != nil {
		return x.RoomLevel
	}
	return 0
}

func (x *Room) GetCharacterId() []string {
	if x != nil {
		return x.CharacterId
	}
	return nil
}

func (x *Room) GetCharacterFtg() []int32 {
	if x != nil {
		return x.CharacterFtg
	}
	return nil
}

func (x *Room) GetStorage() int32 {
	if x != nil {
		return x.Storage
	}
	return 0
}

func (x *Room) GetStorageUsed() int32 {
	if x != nil {
		return x.StorageUsed
	}
	return 0
}

func (x *Room) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Room) GetItemProgress() int32 {
	if x != nil {
		return x.ItemProgress
	}
	return 0
}

// 每个房间的详细信息
type GetDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetDetailReply) Reset() {
	*x = GetDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailReply) ProtoMessage() {}

func (x *GetDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailReply.ProtoReflect.Descriptor instead.
func (*GetDetailReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{4}
}

func (x *GetDetailReply) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type GetProductionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RoomId []int32 `protobuf:"varint,2,rep,packed,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *GetProductionReq) Reset() {
	*x = GetProductionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionReq) ProtoMessage() {}

func (x *GetProductionReq) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionReq.ProtoReflect.Descriptor instead.
func (*GetProductionReq) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductionReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetProductionReq) GetRoomId() []int32 {
	if x != nil {
		return x.RoomId
	}
	return nil
}

type GetProductionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Productions []*Production `protobuf:"bytes,1,rep,name=productions,proto3" json:"productions,omitempty"`
}

func (x *GetProductionReply) Reset() {
	*x = GetProductionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionReply) ProtoMessage() {}

func (x *GetProductionReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionReply.ProtoReflect.Descriptor instead.
func (*GetProductionReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductionReply) GetProductions() []*Production {
	if x != nil {
		return x.Productions
	}
	return nil
}

type Production struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId     int32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemAmount int32 `protobuf:"varint,2,opt,name=item_amount,json=itemAmount,proto3" json:"item_amount,omitempty"`
}

func (x *Production) Reset() {
	*x = Production{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Production) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Production) ProtoMessage() {}

func (x *Production) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Production.ProtoReflect.Descriptor instead.
func (*Production) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{7}
}

func (x *Production) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Production) GetItemAmount() int32 {
	if x != nil {
		return x.ItemAmount
	}
	return 0
}

type UpdateInfraReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RoomId int32 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *UpdateInfraReq) Reset() {
	*x = UpdateInfraReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfraReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfraReq) ProtoMessage() {}

func (x *UpdateInfraReq) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfraReq.ProtoReflect.Descriptor instead.
func (*UpdateInfraReq) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateInfraReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateInfraReq) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type UpdateInfraReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInfraReply) Reset() {
	*x = UpdateInfraReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfraReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfraReply) ProtoMessage() {}

func (x *UpdateInfraReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfraReply.ProtoReflect.Descriptor instead.
func (*UpdateInfraReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{9}
}

type ChangeProductionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RoomId     int32 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	ItemId     int32 `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemAmount int32 `protobuf:"varint,4,opt,name=item_amount,json=itemAmount,proto3" json:"item_amount,omitempty"`
}

func (x *ChangeProductionReq) Reset() {
	*x = ChangeProductionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeProductionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeProductionReq) ProtoMessage() {}

func (x *ChangeProductionReq) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeProductionReq.ProtoReflect.Descriptor instead.
func (*ChangeProductionReq) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{10}
}

func (x *ChangeProductionReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ChangeProductionReq) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ChangeProductionReq) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ChangeProductionReq) GetItemAmount() int32 {
	if x != nil {
		return x.ItemAmount
	}
	return 0
}

type ChangeProductionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修改物品前生产完毕的物品和数量
	ItemProductedId     int32 `protobuf:"varint,1,opt,name=item_producted_id,json=itemProductedId,proto3" json:"item_producted_id,omitempty"`
	ItemProductedAmount int32 `protobuf:"varint,2,opt,name=item_producted_amount,json=itemProductedAmount,proto3" json:"item_producted_amount,omitempty"`
}

func (x *ChangeProductionReply) Reset() {
	*x = ChangeProductionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeProductionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeProductionReply) ProtoMessage() {}

func (x *ChangeProductionReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeProductionReply.ProtoReflect.Descriptor instead.
func (*ChangeProductionReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{11}
}

func (x *ChangeProductionReply) GetItemProductedId() int32 {
	if x != nil {
		return x.ItemProductedId
	}
	return 0
}

func (x *ChangeProductionReply) GetItemProductedAmount() int32 {
	if x != nil {
		return x.ItemProductedAmount
	}
	return 0
}

type ChangeWorkerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RoomId      int32   `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	CharacterId []int32 `protobuf:"varint,3,rep,packed,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
}

func (x *ChangeWorkerReq) Reset() {
	*x = ChangeWorkerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeWorkerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeWorkerReq) ProtoMessage() {}

func (x *ChangeWorkerReq) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeWorkerReq.ProtoReflect.Descriptor instead.
func (*ChangeWorkerReq) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{12}
}

func (x *ChangeWorkerReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ChangeWorkerReq) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ChangeWorkerReq) GetCharacterId() []int32 {
	if x != nil {
		return x.CharacterId
	}
	return nil
}

type ChangeWorkerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeWorkerReply) Reset() {
	*x = ChangeWorkerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arkinfra_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeWorkerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeWorkerReply) ProtoMessage() {}

func (x *ChangeWorkerReply) ProtoReflect() protoreflect.Message {
	mi := &file_arkinfra_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeWorkerReply.ProtoReflect.Descriptor instead.
func (*ChangeWorkerReply) Descriptor() ([]byte, []int) {
	return file_arkinfra_proto_rawDescGZIP(), []int{13}
}

var File_arkinfra_proto protoreflect.FileDescriptor

var file_arkinfra_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x1f, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x58, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x02, 0x0a,
	0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x74, 0x67, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x46,
	0x74, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x39, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x27, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x46, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x72, 0x61, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x8f, 0x01, 0x0a, 0x13,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x77, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x65, 0x64,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x13, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x89, 0x04, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11,
	0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x19, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x61,
	0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x1b, 0x2e, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x72, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x4c, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x1c, 0x2e, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arkinfra_proto_rawDescOnce sync.Once
	file_arkinfra_proto_rawDescData = file_arkinfra_proto_rawDesc
)

func file_arkinfra_proto_rawDescGZIP() []byte {
	file_arkinfra_proto_rawDescOnce.Do(func() {
		file_arkinfra_proto_rawDescData = protoimpl.X.CompressGZIP(file_arkinfra_proto_rawDescData)
	})
	return file_arkinfra_proto_rawDescData
}

var file_arkinfra_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_arkinfra_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: arkinfra.v1.User
	(*CreateInfraReply)(nil),      // 1: arkinfra.v1.CreateInfraReply
	(*GetInfoReply)(nil),          // 2: arkinfra.v1.GetInfoReply
	(*Room)(nil),                  // 3: arkinfra.v1.Room
	(*GetDetailReply)(nil),        // 4: arkinfra.v1.GetDetailReply
	(*GetProductionReq)(nil),      // 5: arkinfra.v1.GetProductionReq
	(*GetProductionReply)(nil),    // 6: arkinfra.v1.GetProductionReply
	(*Production)(nil),            // 7: arkinfra.v1.Production
	(*UpdateInfraReq)(nil),        // 8: arkinfra.v1.UpdateInfraReq
	(*UpdateInfraReply)(nil),      // 9: arkinfra.v1.UpdateInfraReply
	(*ChangeProductionReq)(nil),   // 10: arkinfra.v1.ChangeProductionReq
	(*ChangeProductionReply)(nil), // 11: arkinfra.v1.ChangeProductionReply
	(*ChangeWorkerReq)(nil),       // 12: arkinfra.v1.ChangeWorkerReq
	(*ChangeWorkerReply)(nil),     // 13: arkinfra.v1.ChangeWorkerReply
}
var file_arkinfra_proto_depIdxs = []int32{
	3,  // 0: arkinfra.v1.GetDetailReply.rooms:type_name -> arkinfra.v1.Room
	0,  // 1: arkinfra.v1.GetProductionReq.user:type_name -> arkinfra.v1.User
	7,  // 2: arkinfra.v1.GetProductionReply.productions:type_name -> arkinfra.v1.Production
	0,  // 3: arkinfra.v1.UpdateInfraReq.user:type_name -> arkinfra.v1.User
	0,  // 4: arkinfra.v1.ChangeProductionReq.user:type_name -> arkinfra.v1.User
	0,  // 5: arkinfra.v1.ChangeWorkerReq.user:type_name -> arkinfra.v1.User
	0,  // 6: arkinfra.v1.InfraService.CreateInfra:input_type -> arkinfra.v1.User
	0,  // 7: arkinfra.v1.InfraService.GetInfo:input_type -> arkinfra.v1.User
	0,  // 8: arkinfra.v1.InfraService.GetDetail:input_type -> arkinfra.v1.User
	5,  // 9: arkinfra.v1.InfraService.GetProduction:input_type -> arkinfra.v1.GetProductionReq
	8,  // 10: arkinfra.v1.InfraService.UpdateInfra:input_type -> arkinfra.v1.UpdateInfraReq
	10, // 11: arkinfra.v1.InfraService.ChangeProduction:input_type -> arkinfra.v1.ChangeProductionReq
	12, // 12: arkinfra.v1.InfraService.ChangeWorker:input_type -> arkinfra.v1.ChangeWorkerReq
	1,  // 13: arkinfra.v1.InfraService.CreateInfra:output_type -> arkinfra.v1.CreateInfraReply
	2,  // 14: arkinfra.v1.InfraService.GetInfo:output_type -> arkinfra.v1.GetInfoReply
	4,  // 15: arkinfra.v1.InfraService.GetDetail:output_type -> arkinfra.v1.GetDetailReply
	6,  // 16: arkinfra.v1.InfraService.GetProduction:output_type -> arkinfra.v1.GetProductionReply
	9,  // 17: arkinfra.v1.InfraService.UpdateInfra:output_type -> arkinfra.v1.UpdateInfraReply
	11, // 18: arkinfra.v1.InfraService.ChangeProduction:output_type -> arkinfra.v1.ChangeProductionReply
	13, // 19: arkinfra.v1.InfraService.ChangeWorker:output_type -> arkinfra.v1.ChangeWorkerReply
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_arkinfra_proto_init() }
func file_arkinfra_proto_init() {
	if File_arkinfra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arkinfra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_arkinfra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInfraReply); i {
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
		file_arkinfra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoReply); i {
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
		file_arkinfra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_arkinfra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailReply); i {
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
		file_arkinfra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductionReq); i {
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
		file_arkinfra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductionReply); i {
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
		file_arkinfra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Production); i {
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
		file_arkinfra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfraReq); i {
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
		file_arkinfra_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfraReply); i {
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
		file_arkinfra_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeProductionReq); i {
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
		file_arkinfra_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeProductionReply); i {
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
		file_arkinfra_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeWorkerReq); i {
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
		file_arkinfra_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeWorkerReply); i {
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
			RawDescriptor: file_arkinfra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arkinfra_proto_goTypes,
		DependencyIndexes: file_arkinfra_proto_depIdxs,
		MessageInfos:      file_arkinfra_proto_msgTypes,
	}.Build()
	File_arkinfra_proto = out.File
	file_arkinfra_proto_rawDesc = nil
	file_arkinfra_proto_goTypes = nil
	file_arkinfra_proto_depIdxs = nil
}
