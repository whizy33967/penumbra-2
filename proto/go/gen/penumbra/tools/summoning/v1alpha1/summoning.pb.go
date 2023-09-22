// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/tools/summoning/v1alpha1/summoning.proto

package summoningv1alpha1

import (
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/keys/v1alpha1"
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/num/v1alpha1"
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

type ParticipateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*ParticipateRequest_Identify_
	//	*ParticipateRequest_Contribution_
	Msg isParticipateRequest_Msg `protobuf_oneof:"msg"`
}

func (x *ParticipateRequest) Reset() {
	*x = ParticipateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateRequest) ProtoMessage() {}

func (x *ParticipateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateRequest.ProtoReflect.Descriptor instead.
func (*ParticipateRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{0}
}

func (m *ParticipateRequest) GetMsg() isParticipateRequest_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ParticipateRequest) GetIdentify() *ParticipateRequest_Identify {
	if x, ok := x.GetMsg().(*ParticipateRequest_Identify_); ok {
		return x.Identify
	}
	return nil
}

func (x *ParticipateRequest) GetContribution() *ParticipateRequest_Contribution {
	if x, ok := x.GetMsg().(*ParticipateRequest_Contribution_); ok {
		return x.Contribution
	}
	return nil
}

type isParticipateRequest_Msg interface {
	isParticipateRequest_Msg()
}

type ParticipateRequest_Identify_ struct {
	Identify *ParticipateRequest_Identify `protobuf:"bytes,1,opt,name=identify,proto3,oneof"`
}

type ParticipateRequest_Contribution_ struct {
	Contribution *ParticipateRequest_Contribution `protobuf:"bytes,2,opt,name=contribution,proto3,oneof"`
}

func (*ParticipateRequest_Identify_) isParticipateRequest_Msg() {}

func (*ParticipateRequest_Contribution_) isParticipateRequest_Msg() {}

type CeremonyCrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spend                 []byte `protobuf:"bytes,100,opt,name=spend,proto3" json:"spend,omitempty"`
	Output                []byte `protobuf:"bytes,101,opt,name=output,proto3" json:"output,omitempty"`
	DelegatorVote         []byte `protobuf:"bytes,102,opt,name=delegator_vote,json=delegatorVote,proto3" json:"delegator_vote,omitempty"`
	UndelegateClaim       []byte `protobuf:"bytes,103,opt,name=undelegate_claim,json=undelegateClaim,proto3" json:"undelegate_claim,omitempty"`
	Swap                  []byte `protobuf:"bytes,104,opt,name=swap,proto3" json:"swap,omitempty"`
	SwapClaim             []byte `protobuf:"bytes,105,opt,name=swap_claim,json=swapClaim,proto3" json:"swap_claim,omitempty"`
	NulliferDerivationCrs []byte `protobuf:"bytes,106,opt,name=nullifer_derivation_crs,json=nulliferDerivationCrs,proto3" json:"nullifer_derivation_crs,omitempty"`
}

func (x *CeremonyCrs) Reset() {
	*x = CeremonyCrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CeremonyCrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CeremonyCrs) ProtoMessage() {}

func (x *CeremonyCrs) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CeremonyCrs.ProtoReflect.Descriptor instead.
func (*CeremonyCrs) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{1}
}

func (x *CeremonyCrs) GetSpend() []byte {
	if x != nil {
		return x.Spend
	}
	return nil
}

func (x *CeremonyCrs) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *CeremonyCrs) GetDelegatorVote() []byte {
	if x != nil {
		return x.DelegatorVote
	}
	return nil
}

func (x *CeremonyCrs) GetUndelegateClaim() []byte {
	if x != nil {
		return x.UndelegateClaim
	}
	return nil
}

func (x *CeremonyCrs) GetSwap() []byte {
	if x != nil {
		return x.Swap
	}
	return nil
}

func (x *CeremonyCrs) GetSwapClaim() []byte {
	if x != nil {
		return x.SwapClaim
	}
	return nil
}

func (x *CeremonyCrs) GetNulliferDerivationCrs() []byte {
	if x != nil {
		return x.NulliferDerivationCrs
	}
	return nil
}

type ParticipateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*ParticipateResponse_Position_
	//	*ParticipateResponse_ContributeNow_
	//	*ParticipateResponse_Confirm_
	Msg isParticipateResponse_Msg `protobuf_oneof:"msg"`
}

func (x *ParticipateResponse) Reset() {
	*x = ParticipateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateResponse) ProtoMessage() {}

func (x *ParticipateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateResponse.ProtoReflect.Descriptor instead.
func (*ParticipateResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{2}
}

func (m *ParticipateResponse) GetMsg() isParticipateResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ParticipateResponse) GetPosition() *ParticipateResponse_Position {
	if x, ok := x.GetMsg().(*ParticipateResponse_Position_); ok {
		return x.Position
	}
	return nil
}

func (x *ParticipateResponse) GetContributeNow() *ParticipateResponse_ContributeNow {
	if x, ok := x.GetMsg().(*ParticipateResponse_ContributeNow_); ok {
		return x.ContributeNow
	}
	return nil
}

func (x *ParticipateResponse) GetConfirm() *ParticipateResponse_Confirm {
	if x, ok := x.GetMsg().(*ParticipateResponse_Confirm_); ok {
		return x.Confirm
	}
	return nil
}

type isParticipateResponse_Msg interface {
	isParticipateResponse_Msg()
}

type ParticipateResponse_Position_ struct {
	Position *ParticipateResponse_Position `protobuf:"bytes,1,opt,name=position,proto3,oneof"`
}

type ParticipateResponse_ContributeNow_ struct {
	ContributeNow *ParticipateResponse_ContributeNow `protobuf:"bytes,2,opt,name=contribute_now,json=contributeNow,proto3,oneof"`
}

type ParticipateResponse_Confirm_ struct {
	Confirm *ParticipateResponse_Confirm `protobuf:"bytes,3,opt,name=confirm,proto3,oneof"`
}

func (*ParticipateResponse_Position_) isParticipateResponse_Msg() {}

func (*ParticipateResponse_ContributeNow_) isParticipateResponse_Msg() {}

func (*ParticipateResponse_Confirm_) isParticipateResponse_Msg() {}

// Sent at the beginning of the connection to identify the participant.
type ParticipateRequest_Identify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *v1alpha1.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ParticipateRequest_Identify) Reset() {
	*x = ParticipateRequest_Identify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateRequest_Identify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateRequest_Identify) ProtoMessage() {}

func (x *ParticipateRequest_Identify) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateRequest_Identify.ProtoReflect.Descriptor instead.
func (*ParticipateRequest_Identify) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ParticipateRequest_Identify) GetAddress() *v1alpha1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// Sent by the participant after getting a `ContributeNow` message.
type ParticipateRequest_Contribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated     *CeremonyCrs `protobuf:"bytes,1,opt,name=updated,proto3" json:"updated,omitempty"`
	UpdateProof []byte       `protobuf:"bytes,2,opt,name=update_proof,json=updateProof,proto3" json:"update_proof,omitempty"`
}

func (x *ParticipateRequest_Contribution) Reset() {
	*x = ParticipateRequest_Contribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateRequest_Contribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateRequest_Contribution) ProtoMessage() {}

func (x *ParticipateRequest_Contribution) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateRequest_Contribution.ProtoReflect.Descriptor instead.
func (*ParticipateRequest_Contribution) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ParticipateRequest_Contribution) GetUpdated() *CeremonyCrs {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *ParticipateRequest_Contribution) GetUpdateProof() []byte {
	if x != nil {
		return x.UpdateProof
	}
	return nil
}

// Streamed to the participant to inform them of their position in the queue.
type ParticipateResponse_Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The position of the participant in the queue.
	Position uint32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	// The total number of participants in the queue.
	ConnectedParticipants uint32 `protobuf:"varint,2,opt,name=connected_participants,json=connectedParticipants,proto3" json:"connected_participants,omitempty"`
	// The bid for the most recently executed contribution slot.
	LastSlotBid *v1alpha11.Amount `protobuf:"bytes,3,opt,name=last_slot_bid,json=lastSlotBid,proto3" json:"last_slot_bid,omitempty"`
	// The participant's current bid.
	YourBid *v1alpha11.Amount `protobuf:"bytes,4,opt,name=your_bid,json=yourBid,proto3" json:"your_bid,omitempty"`
}

func (x *ParticipateResponse_Position) Reset() {
	*x = ParticipateResponse_Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateResponse_Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateResponse_Position) ProtoMessage() {}

func (x *ParticipateResponse_Position) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateResponse_Position.ProtoReflect.Descriptor instead.
func (*ParticipateResponse_Position) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ParticipateResponse_Position) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ParticipateResponse_Position) GetConnectedParticipants() uint32 {
	if x != nil {
		return x.ConnectedParticipants
	}
	return 0
}

func (x *ParticipateResponse_Position) GetLastSlotBid() *v1alpha11.Amount {
	if x != nil {
		return x.LastSlotBid
	}
	return nil
}

func (x *ParticipateResponse_Position) GetYourBid() *v1alpha11.Amount {
	if x != nil {
		return x.YourBid
	}
	return nil
}

// Sent to the participant to inform them that they should contribute now.
type ParticipateResponse_ContributeNow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent CRS, which the participant should update.
	Parent *CeremonyCrs `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ParticipateResponse_ContributeNow) Reset() {
	*x = ParticipateResponse_ContributeNow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateResponse_ContributeNow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateResponse_ContributeNow) ProtoMessage() {}

func (x *ParticipateResponse_ContributeNow) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateResponse_ContributeNow.ProtoReflect.Descriptor instead.
func (*ParticipateResponse_ContributeNow) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ParticipateResponse_ContributeNow) GetParent() *CeremonyCrs {
	if x != nil {
		return x.Parent
	}
	return nil
}

// Sent to the participant to confim their contribution was accepted.
type ParticipateResponse_Confirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *ParticipateResponse_Confirm) Reset() {
	*x = ParticipateResponse_Confirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateResponse_Confirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateResponse_Confirm) ProtoMessage() {}

func (x *ParticipateResponse_Confirm) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateResponse_Confirm.ProtoReflect.Descriptor instead.
func (*ParticipateResponse_Confirm) Descriptor() ([]byte, []int) {
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP(), []int{2, 2}
}

func (x *ParticipateResponse_Confirm) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

var File_penumbra_tools_summoning_v1alpha1_summoning_proto protoreflect.FileDescriptor

var file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x26, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x75,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x12, 0x68, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x4a, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x3e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a,
	0x7b, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x48, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x43, 0x72, 0x73,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x05, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0xf8, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79,
	0x43, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x76,
	0x6f, 0x74, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x6e, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x75, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x70, 0x18, 0x68, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x73, 0x77, 0x61, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x5f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x77, 0x61,
	0x70, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x36, 0x0a, 0x17, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72,
	0x73, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x65,
	0x72, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x73, 0x22, 0xa5,
	0x05, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4e, 0x6f, 0x77, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x4e, 0x6f, 0x77, 0x12, 0x5a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x1a, 0xe4, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x16, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73,
	0x12, 0x46, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x62, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x6c, 0x6f, 0x74, 0x42, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x79, 0x6f, 0x75, 0x72,
	0x5f, 0x62, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x79, 0x6f, 0x75, 0x72, 0x42, 0x69, 0x64, 0x1a, 0x57, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x77, 0x12, 0x46, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x65, 0x72,
	0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x43, 0x72, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x1a, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x9f, 0x01, 0x0a, 0x1a, 0x43, 0x65, 0x72, 0x65, 0x6d,
	0x6f, 0x6e, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0xc2, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x54, 0x53, 0xaa, 0x02,
	0x21, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x21, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x54, 0x6f,
	0x6f, 0x6c, 0x73, 0x5c, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x5c, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x5c, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x3a, 0x3a, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x3a, 0x3a, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescOnce sync.Once
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescData = file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDesc
)

func file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescGZIP() []byte {
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescOnce.Do(func() {
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescData)
	})
	return file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDescData
}

var file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_penumbra_tools_summoning_v1alpha1_summoning_proto_goTypes = []interface{}{
	(*ParticipateRequest)(nil),                // 0: penumbra.tools.summoning.v1alpha1.ParticipateRequest
	(*CeremonyCrs)(nil),                       // 1: penumbra.tools.summoning.v1alpha1.CeremonyCrs
	(*ParticipateResponse)(nil),               // 2: penumbra.tools.summoning.v1alpha1.ParticipateResponse
	(*ParticipateRequest_Identify)(nil),       // 3: penumbra.tools.summoning.v1alpha1.ParticipateRequest.Identify
	(*ParticipateRequest_Contribution)(nil),   // 4: penumbra.tools.summoning.v1alpha1.ParticipateRequest.Contribution
	(*ParticipateResponse_Position)(nil),      // 5: penumbra.tools.summoning.v1alpha1.ParticipateResponse.Position
	(*ParticipateResponse_ContributeNow)(nil), // 6: penumbra.tools.summoning.v1alpha1.ParticipateResponse.ContributeNow
	(*ParticipateResponse_Confirm)(nil),       // 7: penumbra.tools.summoning.v1alpha1.ParticipateResponse.Confirm
	(*v1alpha1.Address)(nil),                  // 8: penumbra.core.keys.v1alpha1.Address
	(*v1alpha11.Amount)(nil),                  // 9: penumbra.core.num.v1alpha1.Amount
}
var file_penumbra_tools_summoning_v1alpha1_summoning_proto_depIdxs = []int32{
	3,  // 0: penumbra.tools.summoning.v1alpha1.ParticipateRequest.identify:type_name -> penumbra.tools.summoning.v1alpha1.ParticipateRequest.Identify
	4,  // 1: penumbra.tools.summoning.v1alpha1.ParticipateRequest.contribution:type_name -> penumbra.tools.summoning.v1alpha1.ParticipateRequest.Contribution
	5,  // 2: penumbra.tools.summoning.v1alpha1.ParticipateResponse.position:type_name -> penumbra.tools.summoning.v1alpha1.ParticipateResponse.Position
	6,  // 3: penumbra.tools.summoning.v1alpha1.ParticipateResponse.contribute_now:type_name -> penumbra.tools.summoning.v1alpha1.ParticipateResponse.ContributeNow
	7,  // 4: penumbra.tools.summoning.v1alpha1.ParticipateResponse.confirm:type_name -> penumbra.tools.summoning.v1alpha1.ParticipateResponse.Confirm
	8,  // 5: penumbra.tools.summoning.v1alpha1.ParticipateRequest.Identify.address:type_name -> penumbra.core.keys.v1alpha1.Address
	1,  // 6: penumbra.tools.summoning.v1alpha1.ParticipateRequest.Contribution.updated:type_name -> penumbra.tools.summoning.v1alpha1.CeremonyCrs
	9,  // 7: penumbra.tools.summoning.v1alpha1.ParticipateResponse.Position.last_slot_bid:type_name -> penumbra.core.num.v1alpha1.Amount
	9,  // 8: penumbra.tools.summoning.v1alpha1.ParticipateResponse.Position.your_bid:type_name -> penumbra.core.num.v1alpha1.Amount
	1,  // 9: penumbra.tools.summoning.v1alpha1.ParticipateResponse.ContributeNow.parent:type_name -> penumbra.tools.summoning.v1alpha1.CeremonyCrs
	0,  // 10: penumbra.tools.summoning.v1alpha1.CeremonyCoordinatorService.Participate:input_type -> penumbra.tools.summoning.v1alpha1.ParticipateRequest
	2,  // 11: penumbra.tools.summoning.v1alpha1.CeremonyCoordinatorService.Participate:output_type -> penumbra.tools.summoning.v1alpha1.ParticipateResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_penumbra_tools_summoning_v1alpha1_summoning_proto_init() }
func file_penumbra_tools_summoning_v1alpha1_summoning_proto_init() {
	if File_penumbra_tools_summoning_v1alpha1_summoning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateRequest); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CeremonyCrs); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateResponse); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateRequest_Identify); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateRequest_Contribution); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateResponse_Position); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateResponse_ContributeNow); i {
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
		file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateResponse_Confirm); i {
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
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ParticipateRequest_Identify_)(nil),
		(*ParticipateRequest_Contribution_)(nil),
	}
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ParticipateResponse_Position_)(nil),
		(*ParticipateResponse_ContributeNow_)(nil),
		(*ParticipateResponse_Confirm_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_penumbra_tools_summoning_v1alpha1_summoning_proto_goTypes,
		DependencyIndexes: file_penumbra_tools_summoning_v1alpha1_summoning_proto_depIdxs,
		MessageInfos:      file_penumbra_tools_summoning_v1alpha1_summoning_proto_msgTypes,
	}.Build()
	File_penumbra_tools_summoning_v1alpha1_summoning_proto = out.File
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_rawDesc = nil
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_goTypes = nil
	file_penumbra_tools_summoning_v1alpha1_summoning_proto_depIdxs = nil
}
