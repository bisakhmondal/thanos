// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: info/infopb/rpc.proto

package infopb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type InfoRequest struct {
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{0}
}
func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

type InfoResponse struct {
	LabelSets     []labelpb.ZLabelSet `protobuf:"bytes,1,rep,name=label_sets,json=labelSets,proto3" json:"label_sets"`
	ComponentType string              `protobuf:"bytes,2,opt,name=ComponentType,proto3" json:"ComponentType,omitempty"`
	// StoreInfo holds the metadata related to Store API if exposed by the component otherwise it will be null.
	Store *StoreInfo `protobuf:"bytes,3,opt,name=store,proto3" json:"store,omitempty"`
	// RulesInfo holds the metadata related to Rules API if exposed by the component otherwise it will be null.
	Rules *RulesInfo `protobuf:"bytes,4,opt,name=rules,proto3" json:"rules,omitempty"`
	// MetricMetadataInfo holds the metadata related to Metadata API if exposed by the component otherwise it will be null.
	MetricMetadata *MetricMetadataInfo `protobuf:"bytes,5,opt,name=metric_metadata,json=metricMetadata,proto3" json:"metric_metadata,omitempty"`
	// TargetsInfo holds the metadata related to Targets API if exposed by the component otherwise it will be null.
	Targets *TargetsInfo `protobuf:"bytes,6,opt,name=targets,proto3" json:"targets,omitempty"`
	// ExemplarsInfo holds the metadata related to Exemplars API if exposed by the component otherwise it will be null.
	Exemplars *ExemplarsInfo `protobuf:"bytes,7,opt,name=exemplars,proto3" json:"exemplars,omitempty"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{1}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

// StoreInfo holds the metadata related to Store API exposed by the component.
type StoreInfo struct {
	MinTime int64 `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64 `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
}

func (m *StoreInfo) Reset()         { *m = StoreInfo{} }
func (m *StoreInfo) String() string { return proto.CompactTextString(m) }
func (*StoreInfo) ProtoMessage()    {}
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{2}
}
func (m *StoreInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreInfo.Merge(m, src)
}
func (m *StoreInfo) XXX_Size() int {
	return m.Size()
}
func (m *StoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StoreInfo proto.InternalMessageInfo

// RulesInfo holds the metadata related to Rules API exposed by the component.
type RulesInfo struct {
}

func (m *RulesInfo) Reset()         { *m = RulesInfo{} }
func (m *RulesInfo) String() string { return proto.CompactTextString(m) }
func (*RulesInfo) ProtoMessage()    {}
func (*RulesInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{3}
}
func (m *RulesInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RulesInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RulesInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RulesInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RulesInfo.Merge(m, src)
}
func (m *RulesInfo) XXX_Size() int {
	return m.Size()
}
func (m *RulesInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RulesInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RulesInfo proto.InternalMessageInfo

// MetricMetadataInfo holds the metadata related to Metadata API exposed by the component.
type MetricMetadataInfo struct {
}

func (m *MetricMetadataInfo) Reset()         { *m = MetricMetadataInfo{} }
func (m *MetricMetadataInfo) String() string { return proto.CompactTextString(m) }
func (*MetricMetadataInfo) ProtoMessage()    {}
func (*MetricMetadataInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{4}
}
func (m *MetricMetadataInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricMetadataInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricMetadataInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricMetadataInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricMetadataInfo.Merge(m, src)
}
func (m *MetricMetadataInfo) XXX_Size() int {
	return m.Size()
}
func (m *MetricMetadataInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricMetadataInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MetricMetadataInfo proto.InternalMessageInfo

// TargetsInfo holds the metadata related to Targets API exposed by the component.
type TargetsInfo struct {
}

func (m *TargetsInfo) Reset()         { *m = TargetsInfo{} }
func (m *TargetsInfo) String() string { return proto.CompactTextString(m) }
func (*TargetsInfo) ProtoMessage()    {}
func (*TargetsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{5}
}
func (m *TargetsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TargetsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TargetsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TargetsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetsInfo.Merge(m, src)
}
func (m *TargetsInfo) XXX_Size() int {
	return m.Size()
}
func (m *TargetsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TargetsInfo proto.InternalMessageInfo

// ExemplarsInfo holds the metadata related to Exemplars API exposed by the component.
type ExemplarsInfo struct {
	MinTime int64 `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64 `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
}

func (m *ExemplarsInfo) Reset()         { *m = ExemplarsInfo{} }
func (m *ExemplarsInfo) String() string { return proto.CompactTextString(m) }
func (*ExemplarsInfo) ProtoMessage()    {}
func (*ExemplarsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1214ec45d2bf952, []int{6}
}
func (m *ExemplarsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExemplarsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExemplarsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExemplarsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExemplarsInfo.Merge(m, src)
}
func (m *ExemplarsInfo) XXX_Size() int {
	return m.Size()
}
func (m *ExemplarsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExemplarsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExemplarsInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InfoRequest)(nil), "thanos.info.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "thanos.info.InfoResponse")
	proto.RegisterType((*StoreInfo)(nil), "thanos.info.StoreInfo")
	proto.RegisterType((*RulesInfo)(nil), "thanos.info.RulesInfo")
	proto.RegisterType((*MetricMetadataInfo)(nil), "thanos.info.MetricMetadataInfo")
	proto.RegisterType((*TargetsInfo)(nil), "thanos.info.TargetsInfo")
	proto.RegisterType((*ExemplarsInfo)(nil), "thanos.info.ExemplarsInfo")
}

func init() { proto.RegisterFile("info/infopb/rpc.proto", fileDescriptor_a1214ec45d2bf952) }

var fileDescriptor_a1214ec45d2bf952 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xed, 0x26, 0x4d, 0xf0, 0x98, 0x80, 0x58, 0x15, 0xb4, 0xc9, 0xc1, 0x8d, 0xac, 0x1e,
	0x72, 0x40, 0xb6, 0x14, 0x24, 0x84, 0xc4, 0x89, 0x56, 0x95, 0x40, 0xa2, 0x17, 0x37, 0xa7, 0x5e,
	0xa2, 0x4d, 0x99, 0x06, 0x4b, 0xde, 0x3f, 0x78, 0xb7, 0x52, 0x7a, 0xe3, 0x11, 0x78, 0xac, 0x1c,
	0x7b, 0xe4, 0x84, 0x20, 0x79, 0x11, 0xb4, 0xbb, 0x6e, 0x89, 0x45, 0x4f, 0xbd, 0xd8, 0xbb, 0xfb,
	0xfb, 0xbe, 0xd9, 0x99, 0xf1, 0x18, 0x5e, 0x96, 0xe2, 0x4a, 0xe6, 0xf6, 0xa1, 0x16, 0x79, 0xad,
	0x2e, 0x33, 0x55, 0x4b, 0x23, 0x49, 0x6c, 0xbe, 0x32, 0x21, 0x75, 0x66, 0xc1, 0x68, 0xa8, 0x8d,
	0xac, 0x31, 0xaf, 0xd8, 0x02, 0x2b, 0xb5, 0xc8, 0xcd, 0x8d, 0x42, 0xed, 0x75, 0xa3, 0x83, 0xa5,
	0x5c, 0x4a, 0xb7, 0xcc, 0xed, 0xca, 0x9f, 0xa6, 0x03, 0x88, 0x3f, 0x89, 0x2b, 0x59, 0xe0, 0xb7,
	0x6b, 0xd4, 0x26, 0xfd, 0xde, 0x81, 0xa7, 0x7e, 0xaf, 0x95, 0x14, 0x1a, 0xc9, 0x5b, 0x00, 0x17,
	0x6c, 0xae, 0xd1, 0x68, 0x1a, 0x8e, 0x3b, 0x93, 0x78, 0xfa, 0x22, 0x6b, 0xae, 0xbc, 0xf8, 0x6c,
	0xd1, 0x39, 0x9a, 0xe3, 0xee, 0xfa, 0xd7, 0x61, 0x50, 0x44, 0x55, 0xb3, 0xd7, 0xe4, 0x08, 0x06,
	0x27, 0x92, 0x2b, 0x29, 0x50, 0x98, 0xd9, 0x8d, 0x42, 0xba, 0x37, 0x0e, 0x27, 0x51, 0xd1, 0x3e,
	0x24, 0xaf, 0x61, 0xdf, 0x25, 0x4c, 0x3b, 0xe3, 0x70, 0x12, 0x4f, 0x5f, 0x65, 0x3b, 0xb5, 0x64,
	0xe7, 0x96, 0xb8, 0x64, 0xbc, 0xc8, 0xaa, 0xeb, 0xeb, 0x0a, 0x35, 0xed, 0x3e, 0xa0, 0x2e, 0x2c,
	0xf1, 0x6a, 0x27, 0x22, 0x1f, 0xe1, 0x39, 0x47, 0x53, 0x97, 0x97, 0x73, 0x8e, 0x86, 0x7d, 0x61,
	0x86, 0xd1, 0x7d, 0xe7, 0x3b, 0x6c, 0xf9, 0xce, 0x9c, 0xe6, 0xac, 0x91, 0xb8, 0x00, 0xcf, 0x78,
	0xeb, 0x8c, 0x4c, 0xa1, 0x6f, 0x58, 0xbd, 0xb4, 0x0d, 0xe8, 0xb9, 0x08, 0xb4, 0x15, 0x61, 0xe6,
	0x99, 0xb3, 0xde, 0x09, 0xc9, 0x3b, 0x88, 0x70, 0x85, 0x5c, 0x55, 0xac, 0xd6, 0xb4, 0xef, 0x5c,
	0xa3, 0x96, 0xeb, 0xf4, 0x8e, 0x3a, 0xdf, 0x3f, 0x71, 0xfa, 0x01, 0xa2, 0xfb, 0xca, 0xc9, 0x10,
	0x9e, 0xf0, 0x52, 0xcc, 0x4d, 0xc9, 0x91, 0x86, 0xe3, 0x70, 0xd2, 0x29, 0xfa, 0xbc, 0x14, 0xb3,
	0x92, 0xa3, 0x43, 0x6c, 0xe5, 0xd1, 0x5e, 0x83, 0xd8, 0xca, 0xa2, 0x34, 0x86, 0xe8, 0xbe, 0x1d,
	0xe9, 0x01, 0x90, 0xff, 0x6b, 0xb4, 0xdf, 0x7d, 0x27, 0xef, 0xf4, 0x14, 0x06, 0xad, 0x84, 0x1e,
	0x77, 0xf1, 0xf4, 0x04, 0xba, 0xce, 0xfd, 0xbe, 0x79, 0xb7, 0x1b, 0xb5, 0x33, 0x68, 0xa3, 0xe1,
	0x03, 0xc4, 0x8f, 0xdc, 0xf1, 0xd1, 0xfa, 0x4f, 0x12, 0xac, 0x37, 0x49, 0x78, 0xbb, 0x49, 0xc2,
	0xdf, 0x9b, 0x24, 0xfc, 0xb1, 0x4d, 0x82, 0xdb, 0x6d, 0x12, 0xfc, 0xdc, 0x26, 0xc1, 0x45, 0xcf,
	0xff, 0x00, 0x8b, 0x9e, 0x9b, 0xdf, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x9c, 0xd8,
	0x20, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoClient interface {
	// Info returns the metadata (Eg. LabelSets, Min/Max time) about all the APIs the component supports.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/thanos.info.Info/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
type InfoServer interface {
	// Info returns the metadata (Eg. LabelSets, Min/Max time) about all the APIs the component supports.
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
}

// UnimplementedInfoServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (*UnimplementedInfoServer) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thanos.info.Info/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thanos.info.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Info_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info/infopb/rpc.proto",
}

func (m *InfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *InfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exemplars != nil {
		{
			size, err := m.Exemplars.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Targets != nil {
		{
			size, err := m.Targets.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.MetricMetadata != nil {
		{
			size, err := m.MetricMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Rules != nil {
		{
			size, err := m.Rules.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Store != nil {
		{
			size, err := m.Store.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ComponentType) > 0 {
		i -= len(m.ComponentType)
		copy(dAtA[i:], m.ComponentType)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ComponentType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LabelSets) > 0 {
		for iNdEx := len(m.LabelSets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LabelSets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StoreInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTime != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTime != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RulesInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RulesInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RulesInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MetricMetadataInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricMetadataInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricMetadataInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *TargetsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TargetsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TargetsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ExemplarsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExemplarsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExemplarsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTime != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTime != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *InfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LabelSets) > 0 {
		for _, e := range m.LabelSets {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	l = len(m.ComponentType)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Store != nil {
		l = m.Store.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Rules != nil {
		l = m.Rules.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.MetricMetadata != nil {
		l = m.MetricMetadata.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Targets != nil {
		l = m.Targets.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Exemplars != nil {
		l = m.Exemplars.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *StoreInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTime != 0 {
		n += 1 + sovRpc(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + sovRpc(uint64(m.MaxTime))
	}
	return n
}

func (m *RulesInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MetricMetadataInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *TargetsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ExemplarsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTime != 0 {
		n += 1 + sovRpc(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + sovRpc(uint64(m.MaxTime))
	}
	return n
}

func sovRpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LabelSets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LabelSets = append(m.LabelSets, labelpb.ZLabelSet{})
			if err := m.LabelSets[len(m.LabelSets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComponentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComponentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Store == nil {
				m.Store = &StoreInfo{}
			}
			if err := m.Store.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = &RulesInfo{}
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetricMetadata == nil {
				m.MetricMetadata = &MetricMetadataInfo{}
			}
			if err := m.MetricMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Targets == nil {
				m.Targets = &TargetsInfo{}
			}
			if err := m.Targets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exemplars", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Exemplars == nil {
				m.Exemplars = &ExemplarsInfo{}
			}
			if err := m.Exemplars.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RulesInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RulesInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RulesInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricMetadataInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricMetadataInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricMetadataInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TargetsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TargetsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TargetsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExemplarsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExemplarsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExemplarsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpc = fmt.Errorf("proto: unexpected end of group")
)
