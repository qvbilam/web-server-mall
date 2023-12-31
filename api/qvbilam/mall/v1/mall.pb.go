// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: mall.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "mall/api/qvbilam/page/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PayType int32

const (
	PayType_None  PayType = 0
	PayType_Money PayType = 1
	PayType_Coin  PayType = 2
)

// Enum value maps for PayType.
var (
	PayType_name = map[int32]string{
		0: "None",
		1: "Money",
		2: "Coin",
	}
	PayType_value = map[string]int32{
		"None":  0,
		"Money": 1,
		"Coin":  2,
	}
)

func (x PayType) Enum() *PayType {
	p := new(PayType)
	*p = x
	return p
}

func (x PayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayType) Descriptor() protoreflect.EnumDescriptor {
	return file_mall_proto_enumTypes[0].Descriptor()
}

func (PayType) Type() protoreflect.EnumType {
	return &file_mall_proto_enumTypes[0]
}

func (x PayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayType.Descriptor instead.
func (PayType) EnumDescriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{0}
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid  int64  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tag  string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryResponse) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *CategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*CategoryResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CategoryListResponse) Reset() {
	*x = CategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListResponse) ProtoMessage() {}

func (x *CategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListResponse.ProtoReflect.Descriptor instead.
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryListResponse) GetList() []*CategoryResponse {
	if x != nil {
		return x.List
	}
	return nil
}

type CategoryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *CategoryListRequest) Reset() {
	*x = CategoryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListRequest) ProtoMessage() {}

func (x *CategoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListRequest.ProtoReflect.Descriptor instead.
func (*CategoryListRequest) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryListRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Price     int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Introduce string `protobuf:"bytes,5,opt,name=introduce,proto3" json:"introduce,omitempty"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Tag       string `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{3}
}

func (x *ProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ProductResponse) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *ProductResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProductResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GoodsProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GoodsId   int64            `protobuf:"varint,2,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	ProductId int64            `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
	Count     int64            `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Product   *ProductResponse `protobuf:"bytes,5,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GoodsProductResponse) Reset() {
	*x = GoodsProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsProductResponse) ProtoMessage() {}

func (x *GoodsProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsProductResponse.ProtoReflect.Descriptor instead.
func (*GoodsProductResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{4}
}

func (x *GoodsProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsProductResponse) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GoodsProductResponse) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *GoodsProductResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GoodsProductResponse) GetProduct() *ProductResponse {
	if x != nil {
		return x.Product
	}
	return nil
}

type GoodsDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId    int64                   `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	Name          string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduce     string                  `protobuf:"bytes,4,opt,name=introduce,proto3" json:"introduce,omitempty"`
	Icon          string                  `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	PayType       string                  `protobuf:"bytes,6,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	Price         int64                   `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	OriginalPrice int64                   `protobuf:"varint,8,opt,name=originalPrice,proto3" json:"originalPrice,omitempty"`
	Stocks        int64                   `protobuf:"varint,9,opt,name=stocks,proto3" json:"stocks,omitempty"`
	SoldCount     int64                   `protobuf:"varint,10,opt,name=soldCount,proto3" json:"soldCount,omitempty"`
	IsHot         bool                    `protobuf:"varint,11,opt,name=isHot,proto3" json:"isHot,omitempty"`
	IsUnlimited   bool                    `protobuf:"varint,12,opt,name=isUnlimited,proto3" json:"isUnlimited,omitempty"`
	OnSale        bool                    `protobuf:"varint,13,opt,name=onSale,proto3" json:"onSale,omitempty"`
	Products      []*GoodsProductResponse `protobuf:"bytes,14,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GoodsDetailResponse) Reset() {
	*x = GoodsDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsDetailResponse) ProtoMessage() {}

func (x *GoodsDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsDetailResponse.ProtoReflect.Descriptor instead.
func (*GoodsDetailResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{5}
}

func (x *GoodsDetailResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsDetailResponse) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GoodsDetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsDetailResponse) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *GoodsDetailResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *GoodsDetailResponse) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *GoodsDetailResponse) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsDetailResponse) GetOriginalPrice() int64 {
	if x != nil {
		return x.OriginalPrice
	}
	return 0
}

func (x *GoodsDetailResponse) GetStocks() int64 {
	if x != nil {
		return x.Stocks
	}
	return 0
}

func (x *GoodsDetailResponse) GetSoldCount() int64 {
	if x != nil {
		return x.SoldCount
	}
	return 0
}

func (x *GoodsDetailResponse) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

func (x *GoodsDetailResponse) GetIsUnlimited() bool {
	if x != nil {
		return x.IsUnlimited
	}
	return false
}

func (x *GoodsDetailResponse) GetOnSale() bool {
	if x != nil {
		return x.OnSale
	}
	return false
}

func (x *GoodsDetailResponse) GetProducts() []*GoodsProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type GoodsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*GoodsDetailResponse `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GoodsListResponse) Reset() {
	*x = GoodsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListResponse) ProtoMessage() {}

func (x *GoodsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListResponse.ProtoReflect.Descriptor instead.
func (*GoodsListResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{6}
}

func (x *GoodsListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GoodsListResponse) GetList() []*GoodsDetailResponse {
	if x != nil {
		return x.List
	}
	return nil
}

type GoodsDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GoodsDetailRequest) Reset() {
	*x = GoodsDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsDetailRequest) ProtoMessage() {}

func (x *GoodsDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsDetailRequest.ProtoReflect.Descriptor instead.
func (*GoodsDetailRequest) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{7}
}

func (x *GoodsDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids         []int64         `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CategoryId  int64           `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	Page        *v1.PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
	OnSale      bool            `protobuf:"varint,4,opt,name=OnSale,proto3" json:"OnSale,omitempty"`
	NeedProduct bool            `protobuf:"varint,5,opt,name=NeedProduct,proto3" json:"NeedProduct,omitempty"`
}

func (x *GoodsListRequest) Reset() {
	*x = GoodsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListRequest) ProtoMessage() {}

func (x *GoodsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListRequest.ProtoReflect.Descriptor instead.
func (*GoodsListRequest) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{8}
}

func (x *GoodsListRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GoodsListRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GoodsListRequest) GetPage() *v1.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *GoodsListRequest) GetOnSale() bool {
	if x != nil {
		return x.OnSale
	}
	return false
}

func (x *GoodsListRequest) GetNeedProduct() bool {
	if x != nil {
		return x.NeedProduct
	}
	return false
}

type SellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	UserId int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *SellRequest) Reset() {
	*x = SellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellRequest) ProtoMessage() {}

func (x *SellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellRequest.ProtoReflect.Descriptor instead.
func (*SellRequest) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{9}
}

func (x *SellRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SellRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SellRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayType PayType `protobuf:"varint,1,opt,name=payType,proto3,enum=mallPb.v1.PayType" json:"payType,omitempty"`
	OrderSn string  `protobuf:"bytes,2,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
}

func (x *SellResponse) Reset() {
	*x = SellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellResponse) ProtoMessage() {}

func (x *SellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mall_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellResponse.ProtoReflect.Descriptor instead.
func (*SellResponse) Descriptor() ([]byte, []int) {
	return file_mall_proto_rawDescGZIP(), []int{10}
}

func (x *SellResponse) GetPayType() PayType {
	if x != nil {
		return x.PayType
	}
	return PayType_None
}

func (x *SellResponse) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

var File_mall_proto protoreflect.FileDescriptor

var file_mall_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61,
	0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61,
	0x6d, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x22, 0x47, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0xa5, 0x03, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x55, 0x6e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x55,
	0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x53, 0x61,
	0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x53, 0x61, 0x6c, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x5d, 0x0a,
	0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x12,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x50, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x6e, 0x53, 0x61, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4f, 0x6e, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x4e, 0x65, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x4e, 0x65, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x4b, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x0c,
	0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x6e, 0x2a, 0x28, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x02, 0x32, 0xe7,
	0x02, 0x0a, 0x04, 0x4d, 0x61, 0x6c, 0x6c, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x6d, 0x61,
	0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08,
	0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x50,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x24, 0x5a, 0x22, 0x6d, 0x61, 0x6c, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x6d, 0x61, 0x6c,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x62, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mall_proto_rawDescOnce sync.Once
	file_mall_proto_rawDescData = file_mall_proto_rawDesc
)

func file_mall_proto_rawDescGZIP() []byte {
	file_mall_proto_rawDescOnce.Do(func() {
		file_mall_proto_rawDescData = protoimpl.X.CompressGZIP(file_mall_proto_rawDescData)
	})
	return file_mall_proto_rawDescData
}

var file_mall_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mall_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mall_proto_goTypes = []interface{}{
	(PayType)(0),                 // 0: mallPb.v1.PayType
	(*CategoryResponse)(nil),     // 1: mallPb.v1.CategoryResponse
	(*CategoryListResponse)(nil), // 2: mallPb.v1.CategoryListResponse
	(*CategoryListRequest)(nil),  // 3: mallPb.v1.CategoryListRequest
	(*ProductResponse)(nil),      // 4: mallPb.v1.ProductResponse
	(*GoodsProductResponse)(nil), // 5: mallPb.v1.GoodsProductResponse
	(*GoodsDetailResponse)(nil),  // 6: mallPb.v1.GoodsDetailResponse
	(*GoodsListResponse)(nil),    // 7: mallPb.v1.GoodsListResponse
	(*GoodsDetailRequest)(nil),   // 8: mallPb.v1.GoodsDetailRequest
	(*GoodsListRequest)(nil),     // 9: mallPb.v1.GoodsListRequest
	(*SellRequest)(nil),          // 10: mallPb.v1.SellRequest
	(*SellResponse)(nil),         // 11: mallPb.v1.SellResponse
	(*v1.PageRequest)(nil),       // 12: pagePb.v1.PageRequest
	(*emptypb.Empty)(nil),        // 13: google.protobuf.Empty
}
var file_mall_proto_depIdxs = []int32{
	1,  // 0: mallPb.v1.CategoryListResponse.list:type_name -> mallPb.v1.CategoryResponse
	4,  // 1: mallPb.v1.GoodsProductResponse.product:type_name -> mallPb.v1.ProductResponse
	5,  // 2: mallPb.v1.GoodsDetailResponse.products:type_name -> mallPb.v1.GoodsProductResponse
	6,  // 3: mallPb.v1.GoodsListResponse.list:type_name -> mallPb.v1.GoodsDetailResponse
	12, // 4: mallPb.v1.GoodsListRequest.page:type_name -> pagePb.v1.PageRequest
	0,  // 5: mallPb.v1.SellResponse.payType:type_name -> mallPb.v1.PayType
	3,  // 6: mallPb.v1.Mall.GetCategory:input_type -> mallPb.v1.CategoryListRequest
	9,  // 7: mallPb.v1.Mall.GetGoodsList:input_type -> mallPb.v1.GoodsListRequest
	8,  // 8: mallPb.v1.Mall.GetGoodsDetail:input_type -> mallPb.v1.GoodsDetailRequest
	10, // 9: mallPb.v1.Mall.Sell:input_type -> mallPb.v1.SellRequest
	10, // 10: mallPb.v1.Mall.Rollback:input_type -> mallPb.v1.SellRequest
	2,  // 11: mallPb.v1.Mall.GetCategory:output_type -> mallPb.v1.CategoryListResponse
	7,  // 12: mallPb.v1.Mall.GetGoodsList:output_type -> mallPb.v1.GoodsListResponse
	6,  // 13: mallPb.v1.Mall.GetGoodsDetail:output_type -> mallPb.v1.GoodsDetailResponse
	11, // 14: mallPb.v1.Mall.Sell:output_type -> mallPb.v1.SellResponse
	13, // 15: mallPb.v1.Mall.Rollback:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mall_proto_init() }
func file_mall_proto_init() {
	if File_mall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_mall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListResponse); i {
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
		file_mall_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListRequest); i {
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
		file_mall_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_mall_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsProductResponse); i {
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
		file_mall_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsDetailResponse); i {
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
		file_mall_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListResponse); i {
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
		file_mall_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsDetailRequest); i {
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
		file_mall_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListRequest); i {
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
		file_mall_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellRequest); i {
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
		file_mall_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellResponse); i {
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
			RawDescriptor: file_mall_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mall_proto_goTypes,
		DependencyIndexes: file_mall_proto_depIdxs,
		EnumInfos:         file_mall_proto_enumTypes,
		MessageInfos:      file_mall_proto_msgTypes,
	}.Build()
	File_mall_proto = out.File
	file_mall_proto_rawDesc = nil
	file_mall_proto_goTypes = nil
	file_mall_proto_depIdxs = nil
}
