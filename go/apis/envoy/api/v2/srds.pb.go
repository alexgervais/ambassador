// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/srds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies a routing scope, which associates a
// :ref:`Key<envoy_api_msg_ScopedRouteConfiguration.Key>` to a
// :ref:`envoy_api_msg_RouteConfiguration` (identified by its resource name).
//
// The HTTP connection manager builds up a table consisting of these Key to
// RouteConfiguration mappings, and looks up the RouteConfiguration to use per
// request according to the algorithm specified in the
// :ref:`scope_key_builder<envoy_api_field_config.filter.network.http_connection_manager.v2.ScopedRoutes.scope_key_builder>`
// assigned to the HttpConnectionManager.
//
// For example, with the following configurations (in YAML):
//
// HttpConnectionManager config:
//
// .. code::
//
//   ...
//   scoped_routes:
//     name: foo-scoped-routes
//     scope_key_builder:
//       fragments:
//         - header_value_extractor:
//             name: X-Route-Selector
//             element_separator: ,
//             element:
//               separator: =
//               key: vip
//
// ScopedRouteConfiguration resources (specified statically via
// :ref:`scoped_route_configurations_list<envoy_api_field_config.filter.network.http_connection_manager.v2.ScopedRoutes.scoped_route_configurations_list>`
// or obtained dynamically via SRDS):
//
// .. code::
//
//  (1)
//   name: route-scope1
//   route_configuration_name: route-config1
//   key:
//      fragments:
//        - string_key: 172.10.10.20
//
//  (2)
//   name: route-scope2
//   route_configuration_name: route-config2
//   key:
//     fragments:
//       - string_key: 172.20.20.30
//
// A request from a client such as:
//
// .. code::
//
//     GET / HTTP/1.1
//     Host: foo.com
//     X-Route-Selector: vip=172.10.10.20
//
// would result in the routing table defined by the `route-config1`
// RouteConfiguration being assigned to the HTTP request/stream.
//
// [#comment:next free field: 4]
// [#proto-status: experimental]
type ScopedRouteConfiguration struct {
	// The name assigned to the routing scope.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name to use for a :ref:`envoy_api_msg_DiscoveryRequest` to an
	// RDS server to fetch the :ref:`envoy_api_msg_RouteConfiguration` associated
	// with this scope.
	RouteConfigurationName string `protobuf:"bytes,2,opt,name=route_configuration_name,json=routeConfigurationName,proto3" json:"route_configuration_name,omitempty"`
	// The key to match against.
	Key                  *ScopedRouteConfiguration_Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ScopedRouteConfiguration) Reset()         { *m = ScopedRouteConfiguration{} }
func (m *ScopedRouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration) ProtoMessage()    {}
func (*ScopedRouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_92f394721ede65e9, []int{0}
}
func (m *ScopedRouteConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration.Merge(m, src)
}
func (m *ScopedRouteConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration proto.InternalMessageInfo

func (m *ScopedRouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScopedRouteConfiguration) GetRouteConfigurationName() string {
	if m != nil {
		return m.RouteConfigurationName
	}
	return ""
}

func (m *ScopedRouteConfiguration) GetKey() *ScopedRouteConfiguration_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// Specifies a key which is matched against the output of the
// :ref:`scope_key_builder<envoy_api_field_config.filter.network.http_connection_manager.v2.ScopedRoutes.scope_key_builder>`
// specified in the HttpConnectionManager. The matching is done per HTTP
// request and is dependent on the order of the fragments contained in the
// Key.
type ScopedRouteConfiguration_Key struct {
	// The ordered set of fragments to match against. The order must match the
	// fragments in the corresponding
	// :ref:`scope_key_builder<envoy_api_field_config.filter.network.http_connection_manager.v2.ScopedRoutes.scope_key_builder>`.
	Fragments            []*ScopedRouteConfiguration_Key_Fragment `protobuf:"bytes,1,rep,name=fragments,proto3" json:"fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ScopedRouteConfiguration_Key) Reset()         { *m = ScopedRouteConfiguration_Key{} }
func (m *ScopedRouteConfiguration_Key) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration_Key) ProtoMessage()    {}
func (*ScopedRouteConfiguration_Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_92f394721ede65e9, []int{0, 0}
}
func (m *ScopedRouteConfiguration_Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration_Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration_Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration_Key.Merge(m, src)
}
func (m *ScopedRouteConfiguration_Key) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration_Key) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration_Key.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration_Key proto.InternalMessageInfo

func (m *ScopedRouteConfiguration_Key) GetFragments() []*ScopedRouteConfiguration_Key_Fragment {
	if m != nil {
		return m.Fragments
	}
	return nil
}

type ScopedRouteConfiguration_Key_Fragment struct {
	// Types that are valid to be assigned to Type:
	//	*ScopedRouteConfiguration_Key_Fragment_StringKey
	Type                 isScopedRouteConfiguration_Key_Fragment_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *ScopedRouteConfiguration_Key_Fragment) Reset()         { *m = ScopedRouteConfiguration_Key_Fragment{} }
func (m *ScopedRouteConfiguration_Key_Fragment) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration_Key_Fragment) ProtoMessage()    {}
func (*ScopedRouteConfiguration_Key_Fragment) Descriptor() ([]byte, []int) {
	return fileDescriptor_92f394721ede65e9, []int{0, 0, 0}
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.Merge(m, src)
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment proto.InternalMessageInfo

type isScopedRouteConfiguration_Key_Fragment_Type interface {
	isScopedRouteConfiguration_Key_Fragment_Type()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ScopedRouteConfiguration_Key_Fragment_StringKey struct {
	StringKey string `protobuf:"bytes,1,opt,name=string_key,json=stringKey,proto3,oneof"`
}

func (*ScopedRouteConfiguration_Key_Fragment_StringKey) isScopedRouteConfiguration_Key_Fragment_Type() {
}

func (m *ScopedRouteConfiguration_Key_Fragment) GetType() isScopedRouteConfiguration_Key_Fragment_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) GetStringKey() string {
	if x, ok := m.GetType().(*ScopedRouteConfiguration_Key_Fragment_StringKey); ok {
		return x.StringKey
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ScopedRouteConfiguration_Key_Fragment) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ScopedRouteConfiguration_Key_Fragment_OneofMarshaler, _ScopedRouteConfiguration_Key_Fragment_OneofUnmarshaler, _ScopedRouteConfiguration_Key_Fragment_OneofSizer, []interface{}{
		(*ScopedRouteConfiguration_Key_Fragment_StringKey)(nil),
	}
}

func _ScopedRouteConfiguration_Key_Fragment_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ScopedRouteConfiguration_Key_Fragment)
	// type
	switch x := m.Type.(type) {
	case *ScopedRouteConfiguration_Key_Fragment_StringKey:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.StringKey)
	case nil:
	default:
		return fmt.Errorf("ScopedRouteConfiguration_Key_Fragment.Type has unexpected type %T", x)
	}
	return nil
}

func _ScopedRouteConfiguration_Key_Fragment_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ScopedRouteConfiguration_Key_Fragment)
	switch tag {
	case 1: // type.string_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &ScopedRouteConfiguration_Key_Fragment_StringKey{x}
		return true, err
	default:
		return false, nil
	}
}

func _ScopedRouteConfiguration_Key_Fragment_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ScopedRouteConfiguration_Key_Fragment)
	// type
	switch x := m.Type.(type) {
	case *ScopedRouteConfiguration_Key_Fragment_StringKey:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringKey)))
		n += len(x.StringKey)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ScopedRouteConfiguration)(nil), "envoy.api.v2.ScopedRouteConfiguration")
	proto.RegisterType((*ScopedRouteConfiguration_Key)(nil), "envoy.api.v2.ScopedRouteConfiguration.Key")
	proto.RegisterType((*ScopedRouteConfiguration_Key_Fragment)(nil), "envoy.api.v2.ScopedRouteConfiguration.Key.Fragment")
}

func init() { proto.RegisterFile("envoy/api/v2/srds.proto", fileDescriptor_92f394721ede65e9) }

var fileDescriptor_92f394721ede65e9 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0xfb, 0x92, 0xad, 0x34, 0x53, 0x05, 0x3b, 0x07, 0x1b, 0xe2, 0xba, 0x0d, 0xab, 0xc8,
	0x52, 0x30, 0x2b, 0xe9, 0x45, 0xf6, 0xb8, 0x2d, 0xa5, 0x50, 0x90, 0x92, 0x3d, 0x8a, 0x2c, 0x63,
	0xf2, 0x1a, 0x83, 0xbb, 0x33, 0x71, 0x66, 0x36, 0x18, 0xf0, 0xe4, 0x49, 0x3c, 0xea, 0x17, 0xf0,
	0x63, 0x88, 0xa7, 0x1e, 0x3d, 0x0a, 0x7e, 0x00, 0x65, 0xf1, 0x52, 0xfc, 0x12, 0x92, 0xd9, 0x5d,
	0xdd, 0xb4, 0xac, 0x78, 0xe8, 0x2d, 0xcc, 0xfb, 0xff, 0x7f, 0xff, 0xf7, 0x32, 0xf3, 0xc8, 0x36,
	0xf2, 0x42, 0x94, 0x5d, 0x96, 0x67, 0xdd, 0x22, 0xec, 0x2a, 0x99, 0xa8, 0x20, 0x97, 0x42, 0x0b,
	0x7a, 0xdd, 0x14, 0x02, 0x96, 0x67, 0x41, 0x11, 0x7a, 0xcd, 0x9a, 0x2c, 0xc9, 0x54, 0x2c, 0x0a,
	0x94, 0xe5, 0x4c, 0xeb, 0x35, 0x53, 0x21, 0xd2, 0x11, 0x9a, 0x32, 0xe3, 0x5c, 0x68, 0xa6, 0x33,
	0xc1, 0xe7, 0x24, 0x6f, 0xbb, 0x60, 0xa3, 0x2c, 0x61, 0x1a, 0xbb, 0x8b, 0x8f, 0x59, 0xa1, 0xfd,
	0xcb, 0x22, 0xee, 0x20, 0x16, 0x39, 0x26, 0x91, 0x98, 0x68, 0xdc, 0x17, 0xfc, 0x34, 0x4b, 0x27,
	0xd2, 0x98, 0xe9, 0x1d, 0xd2, 0xe0, 0x6c, 0x8c, 0x2e, 0xf8, 0xd0, 0x71, 0xfa, 0xce, 0xe7, 0xf3,
	0x33, 0xbb, 0x21, 0x2d, 0x1f, 0x22, 0x73, 0x4c, 0xf7, 0x89, 0x2b, 0x2b, 0xd3, 0x30, 0x5e, 0x76,
	0x0d, 0x8d, 0xc5, 0xba, 0x68, 0xb9, 0x25, 0x2f, 0xf1, 0x1f, 0x57, 0x90, 0x23, 0x62, 0xbf, 0xc0,
	0xd2, 0xb5, 0x7d, 0xe8, 0x6c, 0x86, 0xbb, 0xc1, 0xf2, 0xc4, 0xc1, 0xaa, 0xc6, 0x82, 0x63, 0x2c,
	0xfb, 0xa4, 0x62, 0xaf, 0xbf, 0x03, 0xeb, 0x26, 0x44, 0x15, 0xc2, 0xfb, 0x08, 0xc4, 0x3e, 0xc6,
	0x92, 0x3e, 0x25, 0xce, 0xa9, 0x64, 0xe9, 0x18, 0xb9, 0x56, 0x2e, 0xf8, 0x76, 0x67, 0x33, 0xdc,
	0xfb, 0x7f, 0x6e, 0x70, 0x38, 0xf7, 0xce, 0x03, 0xde, 0x83, 0xb5, 0x01, 0xd1, 0x5f, 0xa2, 0xd7,
	0x23, 0x1b, 0x0b, 0x09, 0xdd, 0x21, 0x44, 0x69, 0x99, 0xf1, 0x74, 0x58, 0xcd, 0x60, 0x7e, 0xd3,
	0xd1, 0x5a, 0xe4, 0xcc, 0xce, 0xaa, 0x26, 0x6f, 0x90, 0x86, 0x2e, 0x73, 0xa4, 0xeb, 0x9f, 0xce,
	0xcf, 0x6c, 0x08, 0xbf, 0x5b, 0xa4, 0xb9, 0x14, 0xae, 0x0e, 0x16, 0x97, 0x38, 0x40, 0x59, 0x64,
	0x31, 0xd2, 0x27, 0x84, 0x0e, 0xb4, 0x44, 0x36, 0x5e, 0x56, 0xd1, 0x56, 0xbd, 0xfd, 0x3f, 0xae,
	0x08, 0x5f, 0x4e, 0x50, 0x69, 0x6f, 0x67, 0x65, 0x5d, 0xe5, 0x82, 0x2b, 0x6c, 0xaf, 0x75, 0xe0,
	0x21, 0xd0, 0x84, 0x6c, 0x1d, 0xe0, 0x48, 0xb3, 0x1a, 0xfb, 0xee, 0x05, 0x6f, 0x25, 0xb8, 0x14,
	0x70, 0xef, 0xdf, 0xa2, 0x5a, 0xca, 0x6b, 0xb2, 0x75, 0x88, 0x3a, 0x7e, 0x7e, 0xb5, 0x13, 0xdc,
	0x7f, 0xf3, 0xed, 0xe7, 0x07, 0xcb, 0x6f, 0xdf, 0xae, 0xbd, 0xfd, 0x9e, 0x32, 0x21, 0x0f, 0xcc,
	0xdb, 0x52, 0x3d, 0xd8, 0xed, 0x3f, 0xfa, 0x32, 0x6d, 0xc1, 0xd7, 0x69, 0x0b, 0x7e, 0x4c, 0x5b,
	0x40, 0xbc, 0x4c, 0xcc, 0xc0, 0xb9, 0x14, 0xaf, 0xca, 0x5a, 0x46, 0xdf, 0x19, 0xc8, 0x44, 0x9d,
	0x54, 0x4b, 0x70, 0x02, 0x6f, 0x01, 0x9e, 0x5d, 0x33, 0x0b, 0xb1, 0xf7, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xda, 0xad, 0x3f, 0x95, 0x8e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScopedRoutesDiscoveryServiceClient is the client API for ScopedRoutesDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScopedRoutesDiscoveryServiceClient interface {
	StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error)
	DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error)
	FetchScopedRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type scopedRoutesDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewScopedRoutesDiscoveryServiceClient(cc *grpc.ClientConn) ScopedRoutesDiscoveryServiceClient {
	return &scopedRoutesDiscoveryServiceClient{cc}
}

func (c *scopedRoutesDiscoveryServiceClient) StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ScopedRoutesDiscoveryService/StreamScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceStreamScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ScopedRoutesDiscoveryService/DeltaScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceDeltaScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) FetchScopedRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ScopedRoutesDiscoveryService/FetchScopedRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopedRoutesDiscoveryServiceServer is the server API for ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceServer interface {
	StreamScopedRoutes(ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error
	DeltaScopedRoutes(ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error
	FetchScopedRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedScopedRoutesDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScopedRoutesDiscoveryServiceServer struct {
}

func (*UnimplementedScopedRoutesDiscoveryServiceServer) StreamScopedRoutes(srv ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) DeltaScopedRoutes(srv ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) FetchScopedRoutes(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchScopedRoutes not implemented")
}

func RegisterScopedRoutesDiscoveryServiceServer(s *grpc.Server, srv ScopedRoutesDiscoveryServiceServer) {
	s.RegisterService(&_ScopedRoutesDiscoveryService_serviceDesc, srv)
}

func _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).StreamScopedRoutes(&scopedRoutesDiscoveryServiceStreamScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).DeltaScopedRoutes(&scopedRoutesDiscoveryServiceDeltaScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ScopedRoutesDiscoveryService/FetchScopedRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopedRoutesDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ScopedRoutesDiscoveryService",
	HandlerType: (*ScopedRoutesDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchScopedRoutes",
			Handler:    _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/srds.proto",
}

func (m *ScopedRouteConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSrds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RouteConfigurationName) > 0 {
		i -= len(m.RouteConfigurationName)
		copy(dAtA[i:], m.RouteConfigurationName)
		i = encodeVarintSrds(dAtA, i, uint64(len(m.RouteConfigurationName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSrds(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration_Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration_Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Fragments) > 0 {
		for iNdEx := len(m.Fragments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fragments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSrds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration_Key_Fragment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != nil {
		{
			size := m.Type.Size()
			i -= size
			if _, err := m.Type.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.StringKey)
	copy(dAtA[i:], m.StringKey)
	i = encodeVarintSrds(dAtA, i, uint64(len(m.StringKey)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func encodeVarintSrds(dAtA []byte, offset int, v uint64) int {
	offset -= sovSrds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScopedRouteConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSrds(uint64(l))
	}
	l = len(m.RouteConfigurationName)
	if l > 0 {
		n += 1 + l + sovSrds(uint64(l))
	}
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovSrds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fragments) > 0 {
		for _, e := range m.Fragments {
			l = e.Size()
			n += 1 + l + sovSrds(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key_Fragment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += m.Type.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringKey)
	n += 1 + l + sovSrds(uint64(l))
	return n
}

func sovSrds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSrds(x uint64) (n int) {
	return sovSrds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScopedRouteConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrds
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
			return fmt.Errorf("proto: ScopedRouteConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScopedRouteConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrds
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
				return ErrInvalidLengthSrds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSrds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteConfigurationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrds
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
				return ErrInvalidLengthSrds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSrds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteConfigurationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrds
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
				return ErrInvalidLengthSrds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSrds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &ScopedRouteConfiguration_Key{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScopedRouteConfiguration_Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrds
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrds
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
				return ErrInvalidLengthSrds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSrds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fragments = append(m.Fragments, &ScopedRouteConfiguration_Key_Fragment{})
			if err := m.Fragments[len(m.Fragments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScopedRouteConfiguration_Key_Fragment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrds
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
			return fmt.Errorf("proto: Fragment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fragment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrds
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
				return ErrInvalidLengthSrds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSrds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = &ScopedRouteConfiguration_Key_Fragment_StringKey{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSrds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSrds
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
					return 0, ErrIntOverflowSrds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSrds
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
				return 0, ErrInvalidLengthSrds
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSrds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSrds
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSrds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSrds
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSrds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSrds   = fmt.Errorf("proto: integer overflow")
)
