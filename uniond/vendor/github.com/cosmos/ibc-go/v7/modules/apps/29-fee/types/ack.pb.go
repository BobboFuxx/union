// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/ack.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// IncentivizedAcknowledgement is the acknowledgement format to be used by applications wrapped in the fee middleware
type IncentivizedAcknowledgement struct {
	// the underlying app acknowledgement bytes
	AppAcknowledgement []byte `protobuf:"bytes,1,opt,name=app_acknowledgement,json=appAcknowledgement,proto3" json:"app_acknowledgement,omitempty"`
	// the relayer address which submits the recv packet message
	ForwardRelayerAddress string `protobuf:"bytes,2,opt,name=forward_relayer_address,json=forwardRelayerAddress,proto3" json:"forward_relayer_address,omitempty"`
	// success flag of the base application callback
	UnderlyingAppSuccess bool `protobuf:"varint,3,opt,name=underlying_app_success,json=underlyingAppSuccess,proto3" json:"underlying_app_success,omitempty"`
}

func (m *IncentivizedAcknowledgement) Reset()         { *m = IncentivizedAcknowledgement{} }
func (m *IncentivizedAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*IncentivizedAcknowledgement) ProtoMessage()    {}
func (*IncentivizedAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2834946fb65ea4, []int{0}
}
func (m *IncentivizedAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentivizedAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentivizedAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentivizedAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentivizedAcknowledgement.Merge(m, src)
}
func (m *IncentivizedAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *IncentivizedAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentivizedAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_IncentivizedAcknowledgement proto.InternalMessageInfo

func (m *IncentivizedAcknowledgement) GetAppAcknowledgement() []byte {
	if m != nil {
		return m.AppAcknowledgement
	}
	return nil
}

func (m *IncentivizedAcknowledgement) GetForwardRelayerAddress() string {
	if m != nil {
		return m.ForwardRelayerAddress
	}
	return ""
}

func (m *IncentivizedAcknowledgement) GetUnderlyingAppSuccess() bool {
	if m != nil {
		return m.UnderlyingAppSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*IncentivizedAcknowledgement)(nil), "ibc.applications.fee.v1.IncentivizedAcknowledgement")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/ack.proto", fileDescriptor_ab2834946fb65ea4) }

var fileDescriptor_ab2834946fb65ea4 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xcb, 0x4a, 0xf4, 0x30,
	0x18, 0xc6, 0xf1, 0xe6, 0xfb, 0x40, 0xb4, 0xb8, 0xaa, 0x87, 0x19, 0x10, 0x42, 0x75, 0xd5, 0xcd,
	0x34, 0x8c, 0x47, 0x5c, 0xd6, 0x9d, 0x2b, 0xa1, 0xee, 0xdc, 0x94, 0x34, 0x79, 0x5b, 0xc3, 0xb4,
	0x49, 0x48, 0xd2, 0x0e, 0xf5, 0x2a, 0xbc, 0x20, 0x2f, 0xc0, 0xe5, 0x2c, 0x5d, 0x4a, 0x7b, 0x23,
	0x52, 0x2b, 0x78, 0xd8, 0xe6, 0x97, 0x87, 0x17, 0xfe, 0xfe, 0xb1, 0xc8, 0x19, 0xa1, 0x5a, 0x57,
	0x82, 0x51, 0x27, 0x94, 0xb4, 0xa4, 0x00, 0x20, 0xed, 0x92, 0x50, 0xb6, 0x8a, 0xb5, 0x51, 0x4e,
	0x05, 0x33, 0x91, 0xb3, 0xf8, 0xe7, 0x97, 0xb8, 0x00, 0x88, 0xdb, 0xe5, 0xc9, 0x0b, 0xf2, 0x8f,
	0x6e, 0x25, 0x03, 0xe9, 0x44, 0x2b, 0x9e, 0x80, 0x27, 0x6c, 0x25, 0xd5, 0xba, 0x02, 0x5e, 0x42,
	0x0d, 0xd2, 0x05, 0xc4, 0xdf, 0xa3, 0x5a, 0x67, 0xf4, 0xf7, 0xf3, 0x1c, 0x85, 0x28, 0xda, 0x4d,
	0x03, 0xaa, 0xf5, 0xdf, 0xc1, 0xa5, 0x3f, 0x2b, 0x94, 0x59, 0x53, 0xc3, 0x33, 0x03, 0x15, 0xed,
	0xc0, 0x64, 0x94, 0x73, 0x03, 0xd6, 0xce, 0xff, 0x85, 0x28, 0xda, 0x49, 0x0f, 0xbe, 0x38, 0x9d,
	0x34, 0x99, 0x30, 0x38, 0xf7, 0x0f, 0x1b, 0xc9, 0xc1, 0x54, 0x9d, 0x90, 0x65, 0x36, 0xde, 0xb4,
	0x0d, 0x63, 0xe3, 0xec, 0x7f, 0x88, 0xa2, 0xed, 0x74, 0xff, 0x5b, 0x13, 0xad, 0xef, 0x27, 0xbb,
	0xb9, 0x7b, 0xed, 0x31, 0xda, 0xf4, 0x18, 0xbd, 0xf7, 0x18, 0x3d, 0x0f, 0xd8, 0xdb, 0x0c, 0xd8,
	0x7b, 0x1b, 0xb0, 0xf7, 0x70, 0x51, 0x0a, 0xf7, 0xd8, 0xe4, 0x31, 0x53, 0x35, 0x61, 0xca, 0xd6,
	0xca, 0x12, 0x91, 0xb3, 0x45, 0xa9, 0x48, 0x7b, 0x45, 0x6a, 0xc5, 0x9b, 0x0a, 0xec, 0x18, 0xcd,
	0x92, 0xd3, 0xeb, 0xc5, 0xd8, 0xcb, 0x75, 0x1a, 0x6c, 0xbe, 0xf5, 0xd9, 0xeb, 0xec, 0x23, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x47, 0xb1, 0xa5, 0x54, 0x01, 0x00, 0x00,
}

func (m *IncentivizedAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentivizedAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentivizedAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnderlyingAppSuccess {
		i--
		if m.UnderlyingAppSuccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ForwardRelayerAddress) > 0 {
		i -= len(m.ForwardRelayerAddress)
		copy(dAtA[i:], m.ForwardRelayerAddress)
		i = encodeVarintAck(dAtA, i, uint64(len(m.ForwardRelayerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AppAcknowledgement) > 0 {
		i -= len(m.AppAcknowledgement)
		copy(dAtA[i:], m.AppAcknowledgement)
		i = encodeVarintAck(dAtA, i, uint64(len(m.AppAcknowledgement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAck(dAtA []byte, offset int, v uint64) int {
	offset -= sovAck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncentivizedAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppAcknowledgement)
	if l > 0 {
		n += 1 + l + sovAck(uint64(l))
	}
	l = len(m.ForwardRelayerAddress)
	if l > 0 {
		n += 1 + l + sovAck(uint64(l))
	}
	if m.UnderlyingAppSuccess {
		n += 2
	}
	return n
}

func sovAck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAck(x uint64) (n int) {
	return sovAck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncentivizedAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAck
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
			return fmt.Errorf("proto: IncentivizedAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentivizedAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAcknowledgement", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAck
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAcknowledgement = append(m.AppAcknowledgement[:0], dAtA[iNdEx:postIndex]...)
			if m.AppAcknowledgement == nil {
				m.AppAcknowledgement = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardRelayerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
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
				return ErrInvalidLengthAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForwardRelayerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderlyingAppSuccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UnderlyingAppSuccess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAck
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
func skipAck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAck
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
					return 0, ErrIntOverflowAck
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
					return 0, ErrIntOverflowAck
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
				return 0, ErrInvalidLengthAck
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAck
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAck
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAck        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAck          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAck = fmt.Errorf("proto: unexpected end of group")
)