// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: voter/poll.proto

package types

import (
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Poll struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Options string `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Poll) Reset()         { *m = Poll{} }
func (m *Poll) String() string { return proto.CompactTextString(m) }
func (*Poll) ProtoMessage()    {}
func (*Poll) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16e8e3b131544e5, []int{0}
}
func (m *Poll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Poll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Poll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Poll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Poll.Merge(m, src)
}
func (m *Poll) XXX_Size() int {
	return m.Size()
}
func (m *Poll) XXX_DiscardUnknown() {
	xxx_messageInfo_Poll.DiscardUnknown(m)
}

var xxx_messageInfo_Poll proto.InternalMessageInfo

func (m *Poll) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Poll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Poll) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *Poll) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Poll)(nil), "iroan.voter.voter.Poll")
}

func init() { proto.RegisterFile("voter/poll.proto", fileDescriptor_c16e8e3b131544e5) }

var fileDescriptor_c16e8e3b131544e5 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xcb, 0x2f, 0x49,
	0x2d, 0xd2, 0x2f, 0xc8, 0xcf, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x2c,
	0xca, 0x4f, 0xcc, 0xd3, 0x03, 0x8b, 0x43, 0x48, 0xa5, 0x04, 0x2e, 0x96, 0x80, 0xfc, 0x9c, 0x1c,
	0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14,
	0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x82, 0x19, 0x2c,
	0x0e, 0xe3, 0x82, 0x64, 0x92, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x58, 0x20, 0x32, 0x50,
	0xae, 0x93, 0xdd, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa4, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x5d, 0xa6, 0x0f, 0x71, 0x71, 0x05,
	0x94, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x08, 0xce, 0xbb, 0x6f, 0xcf, 0x00, 0x00, 0x00,
}

func (m *Poll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Poll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Poll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Options) > 0 {
		i -= len(m.Options)
		copy(dAtA[i:], m.Options)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Options)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPoll(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoll(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoll(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Poll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPoll(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Options)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	return n
}

func sovPoll(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoll(x uint64) (n int) {
	return sovPoll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Poll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: Poll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Poll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func skipPoll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
				return 0, ErrInvalidLengthPoll
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoll
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoll
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoll        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoll          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoll = fmt.Errorf("proto: unexpected end of group")
)
