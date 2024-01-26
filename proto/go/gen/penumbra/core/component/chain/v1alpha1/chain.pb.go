// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: penumbra/core/component/chain/v1alpha1/chain.proto

package chainv1alpha1

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

// Global chain configuration data, such as chain ID, epoch duration, etc.
type ChainParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the chain.
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// The duration of each epoch, in number of blocks.
	EpochDuration uint64 `protobuf:"varint,2,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
}

func (x *ChainParameters) Reset() {
	*x = ChainParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainParameters) ProtoMessage() {}

func (x *ChainParameters) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainParameters.ProtoReflect.Descriptor instead.
func (*ChainParameters) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{0}
}

func (x *ChainParameters) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ChainParameters) GetEpochDuration() uint64 {
	if x != nil {
		return x.EpochDuration
	}
	return 0
}

// The ratio between two numbers, used in governance to describe vote thresholds and quorums.
type Ratio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The numerator.
	Numerator uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// The denominator.
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

func (x *Ratio) Reset() {
	*x = Ratio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ratio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ratio) ProtoMessage() {}

func (x *Ratio) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ratio.ProtoReflect.Descriptor instead.
func (*Ratio) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{1}
}

func (x *Ratio) GetNumerator() uint64 {
	if x != nil {
		return x.Numerator
	}
	return 0
}

func (x *Ratio) GetDenominator() uint64 {
	if x != nil {
		return x.Denominator
	}
	return 0
}

// Chain-specific genesis content.
type GenesisContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ChainParameters present at genesis.
	ChainParams *ChainParameters `protobuf:"bytes,1,opt,name=chain_params,json=chainParams,proto3" json:"chain_params,omitempty"`
}

func (x *GenesisContent) Reset() {
	*x = GenesisContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisContent) ProtoMessage() {}

func (x *GenesisContent) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisContent.ProtoReflect.Descriptor instead.
func (*GenesisContent) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{2}
}

func (x *GenesisContent) GetChainParams() *ChainParameters {
	if x != nil {
		return x.ChainParams
	}
	return nil
}

type Epoch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	StartHeight uint64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
}

func (x *Epoch) Reset() {
	*x = Epoch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Epoch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Epoch) ProtoMessage() {}

func (x *Epoch) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Epoch.ProtoReflect.Descriptor instead.
func (*Epoch) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{3}
}

func (x *Epoch) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Epoch) GetStartHeight() uint64 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

type EpochByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *EpochByHeightRequest) Reset() {
	*x = EpochByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochByHeightRequest) ProtoMessage() {}

func (x *EpochByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochByHeightRequest.ProtoReflect.Descriptor instead.
func (*EpochByHeightRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{4}
}

func (x *EpochByHeightRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type EpochByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch *Epoch `protobuf:"bytes,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (x *EpochByHeightResponse) Reset() {
	*x = EpochByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochByHeightResponse) ProtoMessage() {}

func (x *EpochByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochByHeightResponse.ProtoReflect.Descriptor instead.
func (*EpochByHeightResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP(), []int{5}
}

func (x *EpochByHeightResponse) GetEpoch() *Epoch {
	if x != nil {
		return x.Epoch
	}
	return nil
}

var File_penumbra_core_component_chain_v1alpha1_chain_proto protoreflect.FileDescriptor

var file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x53, 0x0a, 0x0f,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x47, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x0e, 0x47, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x0c,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x40, 0x0a, 0x05, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5c, 0x0a, 0x15, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x32, 0x9d, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x0d, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3c, 0x2e, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xda, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x43, 0x43,
	0x43, 0xaa, 0x02, 0x26, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x26, 0x50, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43,
	0x6f, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5c, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x50, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescOnce sync.Once
	file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescData = file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDesc
)

func file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescGZIP() []byte {
	file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescOnce.Do(func() {
		file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescData)
	})
	return file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDescData
}

var file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_penumbra_core_component_chain_v1alpha1_chain_proto_goTypes = []interface{}{
	(*ChainParameters)(nil),       // 0: penumbra.core.component.chain.v1alpha1.ChainParameters
	(*Ratio)(nil),                 // 1: penumbra.core.component.chain.v1alpha1.Ratio
	(*GenesisContent)(nil),        // 2: penumbra.core.component.chain.v1alpha1.GenesisContent
	(*Epoch)(nil),                 // 3: penumbra.core.component.chain.v1alpha1.Epoch
	(*EpochByHeightRequest)(nil),  // 4: penumbra.core.component.chain.v1alpha1.EpochByHeightRequest
	(*EpochByHeightResponse)(nil), // 5: penumbra.core.component.chain.v1alpha1.EpochByHeightResponse
}
var file_penumbra_core_component_chain_v1alpha1_chain_proto_depIdxs = []int32{
	0, // 0: penumbra.core.component.chain.v1alpha1.GenesisContent.chain_params:type_name -> penumbra.core.component.chain.v1alpha1.ChainParameters
	3, // 1: penumbra.core.component.chain.v1alpha1.EpochByHeightResponse.epoch:type_name -> penumbra.core.component.chain.v1alpha1.Epoch
	4, // 2: penumbra.core.component.chain.v1alpha1.QueryService.EpochByHeight:input_type -> penumbra.core.component.chain.v1alpha1.EpochByHeightRequest
	5, // 3: penumbra.core.component.chain.v1alpha1.QueryService.EpochByHeight:output_type -> penumbra.core.component.chain.v1alpha1.EpochByHeightResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_penumbra_core_component_chain_v1alpha1_chain_proto_init() }
func file_penumbra_core_component_chain_v1alpha1_chain_proto_init() {
	if File_penumbra_core_component_chain_v1alpha1_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainParameters); i {
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
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ratio); i {
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
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisContent); i {
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
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Epoch); i {
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
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochByHeightRequest); i {
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
		file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochByHeightResponse); i {
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
			RawDescriptor: file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_penumbra_core_component_chain_v1alpha1_chain_proto_goTypes,
		DependencyIndexes: file_penumbra_core_component_chain_v1alpha1_chain_proto_depIdxs,
		MessageInfos:      file_penumbra_core_component_chain_v1alpha1_chain_proto_msgTypes,
	}.Build()
	File_penumbra_core_component_chain_v1alpha1_chain_proto = out.File
	file_penumbra_core_component_chain_v1alpha1_chain_proto_rawDesc = nil
	file_penumbra_core_component_chain_v1alpha1_chain_proto_goTypes = nil
	file_penumbra_core_component_chain_v1alpha1_chain_proto_depIdxs = nil
}
