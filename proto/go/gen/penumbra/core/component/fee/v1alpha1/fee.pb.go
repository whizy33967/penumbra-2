// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/core/component/fee/v1alpha1/fee.proto

package feev1alpha1

import (
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/asset/v1alpha1"
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

// Specifies fees paid by a transaction.
type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the token used to pay fees.
	Amount *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// If present, the asset ID of the token used to pay fees.
	// If absent, specifies the staking token implicitly.
	AssetId *v1alpha11.AssetId `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP(), []int{0}
}

func (x *Fee) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Fee) GetAssetId() *v1alpha11.AssetId {
	if x != nil {
		return x.AssetId
	}
	return nil
}

var File_penumbra_core_component_fee_v1alpha1_fee_proto protoreflect.FileDescriptor

var file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x42, 0xca, 0x02, 0x0a,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x46, 0x65, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x65,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x66, 0x65, 0x65, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x43, 0x43, 0x46, 0xaa, 0x02, 0x24,
	0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x24, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c,
	0x43, 0x6f, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5c, 0x46,
	0x65, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x30, 0x50, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5c, 0x46, 0x65, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x28, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x46, 0x65, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescOnce sync.Once
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData = file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc
)

func file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP() []byte {
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescOnce.Do(func() {
		file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData)
	})
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData
}

var file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes = []interface{}{
	(*Fee)(nil),               // 0: penumbra.core.component.fee.v1alpha1.Fee
	(*v1alpha1.Amount)(nil),   // 1: penumbra.core.num.v1alpha1.Amount
	(*v1alpha11.AssetId)(nil), // 2: penumbra.core.asset.v1alpha1.AssetId
}
var file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs = []int32{
	1, // 0: penumbra.core.component.fee.v1alpha1.Fee.amount:type_name -> penumbra.core.num.v1alpha1.Amount
	2, // 1: penumbra.core.component.fee.v1alpha1.Fee.asset_id:type_name -> penumbra.core.asset.v1alpha1.AssetId
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_penumbra_core_component_fee_v1alpha1_fee_proto_init() }
func file_penumbra_core_component_fee_v1alpha1_fee_proto_init() {
	if File_penumbra_core_component_fee_v1alpha1_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
			RawDescriptor: file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes,
		DependencyIndexes: file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs,
		MessageInfos:      file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes,
	}.Build()
	File_penumbra_core_component_fee_v1alpha1_fee_proto = out.File
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc = nil
	file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes = nil
	file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs = nil
}