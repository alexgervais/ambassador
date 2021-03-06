// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/transport_socket/tap/v3alpha/tap.proto

package envoy_config_transport_socket_tap_v3alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/go/apis/envoy/api/v3alpha/core"
	v3alpha "github.com/datawire/ambassador/go/apis/envoy/config/common/tap/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

// Configuration for tap transport socket. This wraps another transport socket, providing the
// ability to interpose and record in plain text any traffic that is surfaced to Envoy.
type Tap struct {
	// Common configuration for the tap transport socket.
	CommonConfig *v3alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// The underlying transport socket being wrapped.
	TransportSocket      *core.TransportSocket `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_e77d5f362221a43f, []int{0}
}
func (m *Tap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return m.Size()
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v3alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.transport_socket.tap.v3alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/tap/v3alpha/tap.proto", fileDescriptor_e77d5f362221a43f)
}

var fileDescriptor_e77d5f362221a43f = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33,
	0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0xc1,
	0x9a, 0xf4, 0x20, 0x9a, 0xf4, 0xd0, 0x35, 0xe9, 0x81, 0x14, 0x42, 0x35, 0x49, 0xe9, 0xa0, 0x98,
	0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0x87, 0x62, 0x2a, 0x44, 0x08, 0x62, 0xb0, 0x94, 0x22, 0x44,
	0x75, 0x62, 0x41, 0x26, 0x92, 0x7c, 0x51, 0xaa, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0x54, 0x89, 0x78,
	0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa, 0x3e, 0x8c, 0x01, 0x91, 0x50, 0x3a, 0xc7, 0xc8,
	0xc5, 0x1c, 0x92, 0x58, 0x20, 0x94, 0xce, 0xc5, 0x0b, 0x31, 0x33, 0x1e, 0x62, 0xa9, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x99, 0x1e, 0x8a, 0xa3, 0xa1, 0xd6, 0x22, 0x39, 0x55, 0xcf, 0x19,
	0x2c, 0xe4, 0x5a, 0x51, 0x92, 0x9a, 0x57, 0x9c, 0x99, 0x9f, 0xe7, 0x0c, 0x56, 0xe8, 0xc4, 0xb5,
	0xeb, 0xe5, 0x01, 0x66, 0xd6, 0x2e, 0x46, 0x26, 0x01, 0xc6, 0x20, 0x1e, 0x88, 0x2e, 0x88, 0x8c,
	0x50, 0x1c, 0x97, 0x00, 0xba, 0xd7, 0x25, 0x98, 0xc0, 0x76, 0xa9, 0x43, 0xed, 0x4a, 0x2c, 0xc8,
	0x84, 0x9b, 0x0e, 0xf2, 0x87, 0x5e, 0x08, 0x4c, 0x7d, 0x30, 0x58, 0x39, 0x8a, 0xe1, 0xfc, 0x25,
	0x68, 0x92, 0x81, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x23,
	0x97, 0x79, 0x66, 0x3e, 0xc4, 0xd4, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0xa2, 0x63, 0xc0, 0x89,
	0x23, 0x24, 0xb1, 0x20, 0x00, 0x14, 0x42, 0x01, 0x8c, 0x49, 0x6c, 0xe0, 0xa0, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0xce, 0xbe, 0x87, 0xf6, 0x01, 0x00, 0x00,
}

func (m *Tap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TransportSocket != nil {
		{
			size, err := m.TransportSocket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovTap(uint64(l))
	}
	if m.TransportSocket != nil {
		l = m.TransportSocket.Size()
		n += 1 + l + sovTap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTap(x uint64) (n int) {
	return sovTap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTap
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
			return fmt.Errorf("proto: Tap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTap
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
				return ErrInvalidLengthTap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &v3alpha.CommonExtensionConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransportSocket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTap
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
				return ErrInvalidLengthTap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransportSocket == nil {
				m.TransportSocket = &core.TransportSocket{}
			}
			if err := m.TransportSocket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTap
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
func skipTap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
				return 0, ErrInvalidLengthTap
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTap
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTap
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
				next, err := skipTap(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTap
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
	ErrInvalidLengthTap = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTap   = fmt.Errorf("proto: integer overflow")
)
