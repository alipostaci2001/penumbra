// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/core/asset/v1alpha1/asset.proto

package assetv1alpha1

import (
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/num/v1alpha1"
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

type BalanceCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner []byte `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *BalanceCommitment) Reset() {
	*x = BalanceCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceCommitment) ProtoMessage() {}

func (x *BalanceCommitment) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceCommitment.ProtoReflect.Descriptor instead.
func (*BalanceCommitment) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceCommitment) GetInner() []byte {
	if x != nil {
		return x.Inner
	}
	return nil
}

// A Penumbra asset ID.
type AssetId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of the asset ID.
	Inner []byte `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	// Alternatively, a Bech32m-encoded string representation of the `inner`
	// bytes.
	//
	// NOTE: implementations are not required to support parsing this field.
	// Implementations should prefer to encode the `inner` bytes in all messages they
	// produce. Implementations must not accept messages with both `inner` and
	// `alt_bech32m` set.  This field exists for convenience of RPC users.
	AltBech32M string `protobuf:"bytes,2,opt,name=alt_bech32m,json=altBech32m,proto3" json:"alt_bech32m,omitempty"`
	// Alternatively, a base denomination string which should be hashed to obtain the asset ID.
	//
	// NOTE: implementations are not required to support parsing this field.
	// Implementations should prefer to encode the bytes in all messages they
	// produce. Implementations must not accept messages with both `inner` and
	// `alt_base_denom` set.  This field exists for convenience of RPC users.
	AltBaseDenom string `protobuf:"bytes,3,opt,name=alt_base_denom,json=altBaseDenom,proto3" json:"alt_base_denom,omitempty"`
}

func (x *AssetId) Reset() {
	*x = AssetId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetId) ProtoMessage() {}

func (x *AssetId) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetId.ProtoReflect.Descriptor instead.
func (*AssetId) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{1}
}

func (x *AssetId) GetInner() []byte {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *AssetId) GetAltBech32M() string {
	if x != nil {
		return x.AltBech32M
	}
	return ""
}

func (x *AssetId) GetAltBaseDenom() string {
	if x != nil {
		return x.AltBaseDenom
	}
	return ""
}

type Denom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (x *Denom) Reset() {
	*x = Denom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Denom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Denom) ProtoMessage() {}

func (x *Denom) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Denom.ProtoReflect.Descriptor instead.
func (*Denom) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{2}
}

func (x *Denom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

// DenomMetadata represents a struct that describes a basic token.
type DenomMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// denom_units represents the list of DenomUnit's for a given coin
	DenomUnits []*DenomUnit `protobuf:"bytes,2,rep,name=denom_units,json=denomUnits,proto3" json:"denom_units,omitempty"`
	// base represents the base denom (should be the DenomUnit with exponent = 0).
	Base string `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	// display indicates the suggested denom that should be
	// displayed in clients.
	Display string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
	// name defines the name of the token (eg: Cosmos Atom)
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
	// be the same as the display.
	Symbol string `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// URI to a document (on or off-chain) that contains additional information. Optional.
	Uri string `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`
	// URIHash is a sha256 hash of a document pointed by URI. It's used to verify that
	// the document didn't change. Optional.
	UriHash string `protobuf:"bytes,8,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	// the asset ID on Penumbra for this denomination.
	PenumbraAssetId *AssetId `protobuf:"bytes,1984,opt,name=penumbra_asset_id,json=penumbraAssetId,proto3" json:"penumbra_asset_id,omitempty"`
}

func (x *DenomMetadata) Reset() {
	*x = DenomMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenomMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenomMetadata) ProtoMessage() {}

func (x *DenomMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenomMetadata.ProtoReflect.Descriptor instead.
func (*DenomMetadata) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{3}
}

func (x *DenomMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DenomMetadata) GetDenomUnits() []*DenomUnit {
	if x != nil {
		return x.DenomUnits
	}
	return nil
}

func (x *DenomMetadata) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *DenomMetadata) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *DenomMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DenomMetadata) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *DenomMetadata) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *DenomMetadata) GetUriHash() string {
	if x != nil {
		return x.UriHash
	}
	return ""
}

func (x *DenomMetadata) GetPenumbraAssetId() *AssetId {
	if x != nil {
		return x.PenumbraAssetId
	}
	return nil
}

// DenomUnit represents a struct that describes a given denomination unit of the basic token.
type DenomUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom represents the string name of the given denom unit (e.g uatom).
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// exponent represents power of 10 exponent that one must
	// raise the base_denom to in order to equal the given DenomUnit's denom
	// 1 denom = 10^exponent base_denom
	// (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
	// exponent = 6, thus: 1 atom = 10^6 uatom).
	Exponent uint32 `protobuf:"varint,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// aliases is a list of string aliases for the given denom
	Aliases []string `protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *DenomUnit) Reset() {
	*x = DenomUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenomUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenomUnit) ProtoMessage() {}

func (x *DenomUnit) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenomUnit.ProtoReflect.Descriptor instead.
func (*DenomUnit) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{4}
}

func (x *DenomUnit) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *DenomUnit) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

func (x *DenomUnit) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AssetId *AssetId         `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{5}
}

func (x *Value) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Value) GetAssetId() *AssetId {
	if x != nil {
		return x.AssetId
	}
	return nil
}

// Represents a value of a known or unknown denomination.
//
// Note: unlike some other View types, we don't just store the underlying
// `Value` message together with an additional `Denom`.  Instead, we record
// either an `Amount` and `Denom` (only) or an `Amount` and `AssetId`.  This is
// because we don't want to allow a situation where the supplied `Denom` doesn't
// match the `AssetId`, and a consumer of the API that doesn't check is tricked.
// This way, the `Denom` will always match, because the consumer is forced to
// recompute it themselves if they want it.
type ValueView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ValueView:
	//
	//	*ValueView_KnownDenom_
	//	*ValueView_UnknownDenom_
	ValueView isValueView_ValueView `protobuf_oneof:"value_view"`
}

func (x *ValueView) Reset() {
	*x = ValueView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueView) ProtoMessage() {}

func (x *ValueView) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueView.ProtoReflect.Descriptor instead.
func (*ValueView) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{6}
}

func (m *ValueView) GetValueView() isValueView_ValueView {
	if m != nil {
		return m.ValueView
	}
	return nil
}

func (x *ValueView) GetKnownDenom() *ValueView_KnownDenom {
	if x, ok := x.GetValueView().(*ValueView_KnownDenom_); ok {
		return x.KnownDenom
	}
	return nil
}

func (x *ValueView) GetUnknownDenom() *ValueView_UnknownDenom {
	if x, ok := x.GetValueView().(*ValueView_UnknownDenom_); ok {
		return x.UnknownDenom
	}
	return nil
}

type isValueView_ValueView interface {
	isValueView_ValueView()
}

type ValueView_KnownDenom_ struct {
	KnownDenom *ValueView_KnownDenom `protobuf:"bytes,1,opt,name=known_denom,json=knownDenom,proto3,oneof"`
}

type ValueView_UnknownDenom_ struct {
	UnknownDenom *ValueView_UnknownDenom `protobuf:"bytes,2,opt,name=unknown_denom,json=unknownDenom,proto3,oneof"`
}

func (*ValueView_KnownDenom_) isValueView_ValueView() {}

func (*ValueView_UnknownDenom_) isValueView_ValueView() {}

// A value whose asset ID has a known denomination.
type ValueView_KnownDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom  *DenomMetadata   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (x *ValueView_KnownDenom) Reset() {
	*x = ValueView_KnownDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueView_KnownDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueView_KnownDenom) ProtoMessage() {}

func (x *ValueView_KnownDenom) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueView_KnownDenom.ProtoReflect.Descriptor instead.
func (*ValueView_KnownDenom) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ValueView_KnownDenom) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ValueView_KnownDenom) GetDenom() *DenomMetadata {
	if x != nil {
		return x.Denom
	}
	return nil
}

type ValueView_UnknownDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AssetId *AssetId         `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *ValueView_UnknownDenom) Reset() {
	*x = ValueView_UnknownDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueView_UnknownDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueView_UnknownDenom) ProtoMessage() {}

func (x *ValueView_UnknownDenom) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueView_UnknownDenom.ProtoReflect.Descriptor instead.
func (*ValueView_UnknownDenom) Descriptor() ([]byte, []int) {
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP(), []int{6, 1}
}

func (x *ValueView_UnknownDenom) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ValueView_UnknownDenom) GetAssetId() *AssetId {
	if x != nil {
		return x.AssetId
	}
	return nil
}

var File_penumbra_core_asset_v1alpha1_asset_proto protoreflect.FileDescriptor

var file_penumbra_core_asset_v1alpha1_asset_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29,
	0x0a, 0x11, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x07, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c,
	0x74, 0x5f, 0x62, 0x65, 0x63, 0x68, 0x33, 0x32, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x6c, 0x74, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x61,
	0x6c, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x74, 0x42, 0x61, 0x73, 0x65, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x22, 0x1d, 0x0a, 0x05, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x22, 0xd6, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72,
	0x69, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72,
	0x69, 0x48, 0x61, 0x73, 0x68, 0x12, 0x52, 0x0a, 0x11, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xc0, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x52, 0x0f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x09, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x65, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0xea, 0x03, 0x0a, 0x09, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x55, 0x0a, 0x0b, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x48, 0x00, 0x52, 0x0a, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x5b, 0x0a, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x0c,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x1a, 0x8b, 0x01, 0x0a,
	0x0a, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x1a, 0x8c, 0x01, 0x0a, 0x0c, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x42, 0x9c, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x43, 0x41, 0xaa, 0x02, 0x1c, 0x50, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x50, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x28, 0x50, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1f, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x3a,
	0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_core_asset_v1alpha1_asset_proto_rawDescOnce sync.Once
	file_penumbra_core_asset_v1alpha1_asset_proto_rawDescData = file_penumbra_core_asset_v1alpha1_asset_proto_rawDesc
)

func file_penumbra_core_asset_v1alpha1_asset_proto_rawDescGZIP() []byte {
	file_penumbra_core_asset_v1alpha1_asset_proto_rawDescOnce.Do(func() {
		file_penumbra_core_asset_v1alpha1_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_asset_v1alpha1_asset_proto_rawDescData)
	})
	return file_penumbra_core_asset_v1alpha1_asset_proto_rawDescData
}

var file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_penumbra_core_asset_v1alpha1_asset_proto_goTypes = []interface{}{
	(*BalanceCommitment)(nil),      // 0: penumbra.core.asset.v1alpha1.BalanceCommitment
	(*AssetId)(nil),                // 1: penumbra.core.asset.v1alpha1.AssetId
	(*Denom)(nil),                  // 2: penumbra.core.asset.v1alpha1.Denom
	(*DenomMetadata)(nil),          // 3: penumbra.core.asset.v1alpha1.DenomMetadata
	(*DenomUnit)(nil),              // 4: penumbra.core.asset.v1alpha1.DenomUnit
	(*Value)(nil),                  // 5: penumbra.core.asset.v1alpha1.Value
	(*ValueView)(nil),              // 6: penumbra.core.asset.v1alpha1.ValueView
	(*ValueView_KnownDenom)(nil),   // 7: penumbra.core.asset.v1alpha1.ValueView.KnownDenom
	(*ValueView_UnknownDenom)(nil), // 8: penumbra.core.asset.v1alpha1.ValueView.UnknownDenom
	(*v1alpha1.Amount)(nil),        // 9: penumbra.core.num.v1alpha1.Amount
}
var file_penumbra_core_asset_v1alpha1_asset_proto_depIdxs = []int32{
	4,  // 0: penumbra.core.asset.v1alpha1.DenomMetadata.denom_units:type_name -> penumbra.core.asset.v1alpha1.DenomUnit
	1,  // 1: penumbra.core.asset.v1alpha1.DenomMetadata.penumbra_asset_id:type_name -> penumbra.core.asset.v1alpha1.AssetId
	9,  // 2: penumbra.core.asset.v1alpha1.Value.amount:type_name -> penumbra.core.num.v1alpha1.Amount
	1,  // 3: penumbra.core.asset.v1alpha1.Value.asset_id:type_name -> penumbra.core.asset.v1alpha1.AssetId
	7,  // 4: penumbra.core.asset.v1alpha1.ValueView.known_denom:type_name -> penumbra.core.asset.v1alpha1.ValueView.KnownDenom
	8,  // 5: penumbra.core.asset.v1alpha1.ValueView.unknown_denom:type_name -> penumbra.core.asset.v1alpha1.ValueView.UnknownDenom
	9,  // 6: penumbra.core.asset.v1alpha1.ValueView.KnownDenom.amount:type_name -> penumbra.core.num.v1alpha1.Amount
	3,  // 7: penumbra.core.asset.v1alpha1.ValueView.KnownDenom.denom:type_name -> penumbra.core.asset.v1alpha1.DenomMetadata
	9,  // 8: penumbra.core.asset.v1alpha1.ValueView.UnknownDenom.amount:type_name -> penumbra.core.num.v1alpha1.Amount
	1,  // 9: penumbra.core.asset.v1alpha1.ValueView.UnknownDenom.asset_id:type_name -> penumbra.core.asset.v1alpha1.AssetId
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_penumbra_core_asset_v1alpha1_asset_proto_init() }
func file_penumbra_core_asset_v1alpha1_asset_proto_init() {
	if File_penumbra_core_asset_v1alpha1_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceCommitment); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetId); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Denom); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenomMetadata); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenomUnit); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueView); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueView_KnownDenom); i {
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
		file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueView_UnknownDenom); i {
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
	file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ValueView_KnownDenom_)(nil),
		(*ValueView_UnknownDenom_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_core_asset_v1alpha1_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_penumbra_core_asset_v1alpha1_asset_proto_goTypes,
		DependencyIndexes: file_penumbra_core_asset_v1alpha1_asset_proto_depIdxs,
		MessageInfos:      file_penumbra_core_asset_v1alpha1_asset_proto_msgTypes,
	}.Build()
	File_penumbra_core_asset_v1alpha1_asset_proto = out.File
	file_penumbra_core_asset_v1alpha1_asset_proto_rawDesc = nil
	file_penumbra_core_asset_v1alpha1_asset_proto_goTypes = nil
	file_penumbra_core_asset_v1alpha1_asset_proto_depIdxs = nil
}
