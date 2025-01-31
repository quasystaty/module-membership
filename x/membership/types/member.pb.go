// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: membershipmodule/membership/member.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// MembershipStatus enumerates the valid membership states for a citizen of The
// Denom
type MembershipStatus int32

const (
	// MEMBERSHIP_STATUS_UNSPECIFIED defines a no-op status
	MembershipStatus_MemberStatusEmpty MembershipStatus = 0
	// MEMBERSHIP_STATUS_ELECTORATE defines this member as being an active citizen
	MembershipStatus_MemberElectorate MembershipStatus = 1
	// MEMBERSHIP_STATUS_INACTIVE defines this member as being an inactive citizen
	MembershipStatus_MemberInactive MembershipStatus = 2
	// MEMBERSHIP_STATUS_RECALLED defines this member as being recalled
	MembershipStatus_MemberRecalled MembershipStatus = 3
	// MEMBERSHIP_STATUS_EXPULSED defines this member as being expulsed
	MembershipStatus_MemberExpulsed MembershipStatus = 4
)

var MembershipStatus_name = map[int32]string{
	0: "MEMBERSHIP_STATUS_UNSPECIFIED",
	1: "MEMBERSHIP_STATUS_ELECTORATE",
	2: "MEMBERSHIP_STATUS_INACTIVE",
	3: "MEMBERSHIP_STATUS_RECALLED",
	4: "MEMBERSHIP_STATUS_EXPULSED",
}

var MembershipStatus_value = map[string]int32{
	"MEMBERSHIP_STATUS_UNSPECIFIED": 0,
	"MEMBERSHIP_STATUS_ELECTORATE":  1,
	"MEMBERSHIP_STATUS_INACTIVE":    2,
	"MEMBERSHIP_STATUS_RECALLED":    3,
	"MEMBERSHIP_STATUS_EXPULSED":    4,
}

func (x MembershipStatus) String() string {
	return proto.EnumName(MembershipStatus_name, int32(x))
}

func (MembershipStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ba2fc082642cdf4, []int{0}
}

// Member is a specialisation of BaseAccount that adds Member Status and
// Nickname
type Member struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	// status defines the membership status of this member
	Status MembershipStatus `protobuf:"varint,2,opt,name=status,proto3,enum=membershipmodule.membership.MembershipStatus" json:"status,omitempty"`
	// nickname defines the nickname of this member
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// is_guardian defines whether this member is a guardian
	IsGuardian bool `protobuf:"varint,4,opt,name=is_guardian,json=isGuardian,proto3" json:"is_guardian,omitempty"`
}

func (m *Member) Reset()      { *m = Member{} }
func (*Member) ProtoMessage() {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba2fc082642cdf4, []int{0}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("membershipmodule.membership.MembershipStatus", MembershipStatus_name, MembershipStatus_value)
	proto.RegisterType((*Member)(nil), "membershipmodule.membership.Member")
}

func init() {
	proto.RegisterFile("membershipmodule/membership/member.proto", fileDescriptor_5ba2fc082642cdf4)
}

var fileDescriptor_5ba2fc082642cdf4 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0xd3, 0x4e,
	0x18, 0xc6, 0x33, 0xdd, 0x52, 0xfa, 0x9f, 0xfe, 0x59, 0xe2, 0xb0, 0x42, 0x89, 0x9a, 0x06, 0x4f,
	0x41, 0x68, 0xc2, 0x56, 0x10, 0xf5, 0x96, 0xb6, 0xa3, 0x06, 0xda, 0xb5, 0x24, 0xa9, 0x88, 0x97,
	0x32, 0x49, 0x87, 0x36, 0x98, 0x64, 0x42, 0x66, 0xb2, 0xec, 0x7e, 0x03, 0xe9, 0xc9, 0xa3, 0x97,
	0x42, 0x3f, 0x8e, 0xc7, 0x1e, 0x3d, 0x89, 0xb6, 0x5f, 0x44, 0x4c, 0xc2, 0xb6, 0xec, 0xaa, 0xb7,
	0xe7, 0x7d, 0x78, 0x7f, 0xcf, 0x3b, 0x30, 0x0f, 0xd4, 0x63, 0x1a, 0xfb, 0x34, 0xe3, 0xcb, 0x30,
	0x8d, 0xd9, 0x3c, 0x8f, 0xa8, 0x79, 0x30, 0x2a, 0x69, 0xa4, 0x19, 0x13, 0x0c, 0x3d, 0xb8, 0xbd,
	0x69, 0x1c, 0x0c, 0x45, 0x0d, 0x18, 0x8f, 0x19, 0x37, 0x49, 0x2e, 0x96, 0xe6, 0xe5, 0xb9, 0x4f,
	0x05, 0x39, 0x2f, 0x86, 0x12, 0x56, 0xce, 0x16, 0x6c, 0xc1, 0x0a, 0x69, 0xfe, 0x56, 0xa5, 0xfb,
	0xf8, 0x27, 0x80, 0x8d, 0x71, 0x11, 0x82, 0x6c, 0xf8, 0xbf, 0x4f, 0x38, 0x9d, 0x91, 0x20, 0x60,
	0x79, 0x22, 0xda, 0x40, 0x03, 0x7a, 0xab, 0xa7, 0x19, 0x65, 0xae, 0x51, 0x44, 0x55, 0xb9, 0x46,
	0x9f, 0x70, 0x6a, 0x95, 0x7b, 0xfd, 0xfa, 0xf6, 0x7b, 0x07, 0x38, 0x2d, 0xff, 0x60, 0x21, 0x0c,
	0x1b, 0x5c, 0x10, 0x91, 0xf3, 0x76, 0x4d, 0x03, 0xfa, 0x69, 0xaf, 0x6b, 0xfc, 0xe3, 0xe5, 0xc6,
	0xf8, 0x46, 0xba, 0x05, 0xe4, 0x54, 0x30, 0x52, 0x60, 0x33, 0x09, 0x83, 0x8f, 0x09, 0x89, 0x69,
	0xfb, 0x44, 0x03, 0xfa, 0x7f, 0xce, 0xcd, 0x8c, 0x3a, 0xb0, 0x15, 0xf2, 0xd9, 0x22, 0x27, 0xd9,
	0x3c, 0x24, 0x49, 0xbb, 0xae, 0x01, 0xbd, 0xe9, 0xc0, 0x90, 0xbf, 0xae, 0x9c, 0x97, 0xcd, 0x4f,
	0x9b, 0x8e, 0xf4, 0x65, 0xd3, 0x91, 0x9e, 0x6c, 0x6a, 0x50, 0xbe, 0x7d, 0x03, 0x3d, 0x87, 0x8f,
	0xc6, 0x78, 0xdc, 0xc7, 0x8e, 0xfb, 0xc6, 0x9e, 0xcc, 0x5c, 0xcf, 0xf2, 0xa6, 0xee, 0x6c, 0x7a,
	0xe1, 0x4e, 0xf0, 0xc0, 0x7e, 0x65, 0xe3, 0xa1, 0x2c, 0x29, 0xf7, 0x57, 0x6b, 0xed, 0x5e, 0x09,
	0x96, 0x10, 0x8e, 0x53, 0x71, 0x8d, 0x9e, 0xc1, 0x87, 0x77, 0x49, 0x3c, 0xc2, 0x03, 0xef, 0xad,
	0x63, 0x79, 0x58, 0x06, 0xca, 0xd9, 0x6a, 0xad, 0x55, 0x17, 0x71, 0x44, 0x03, 0xc1, 0x32, 0x22,
	0x28, 0xea, 0x41, 0xe5, 0x2e, 0x67, 0x5f, 0x58, 0x03, 0xcf, 0x7e, 0x87, 0xe5, 0x9a, 0x82, 0x56,
	0x6b, 0xed, 0xb4, 0xa4, 0xec, 0x84, 0x04, 0x22, 0xbc, 0xfc, 0x0b, 0xe3, 0xe0, 0x81, 0x35, 0x1a,
	0xe1, 0xa1, 0x7c, 0x72, 0xcc, 0x38, 0x34, 0x20, 0x51, 0x44, 0xe7, 0x7f, 0x66, 0xf0, 0xfb, 0xc9,
	0x74, 0xe4, 0xe2, 0xa1, 0x5c, 0x3f, 0x66, 0xf0, 0x55, 0x9a, 0x47, 0x9c, 0xce, 0xfb, 0xee, 0xd7,
	0x9d, 0x0a, 0xb6, 0x3b, 0x15, 0xfc, 0xd8, 0xa9, 0xe0, 0xf3, 0x5e, 0x95, 0xb6, 0x7b, 0x55, 0xfa,
	0xb6, 0x57, 0xa5, 0x0f, 0x2f, 0x16, 0xa1, 0x58, 0xe6, 0xbe, 0x11, 0xb0, 0xd8, 0x4c, 0x58, 0x16,
	0x92, 0x6e, 0x42, 0x85, 0x59, 0x7e, 0x62, 0xf7, 0xa8, 0xa8, 0x57, 0xc7, 0xad, 0x15, 0xd7, 0x29,
	0xe5, 0x7e, 0xa3, 0xa8, 0xd8, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x09, 0xcd, 0x70,
	0xe1, 0x02, 0x00, 0x00,
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsGuardian {
		i--
		if m.IsGuardian {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintMember(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintMember(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMember(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMember(dAtA []byte, offset int, v uint64) int {
	offset -= sovMember(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMember(uint64(m.Status))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	if m.IsGuardian {
		n += 2
	}
	return n
}

func sovMember(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMember(x uint64) (n int) {
	return sovMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MembershipStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGuardian", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
			m.IsGuardian = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMember
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
func skipMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
				return 0, ErrInvalidLengthMember
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMember
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMember
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMember        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMember          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMember = fmt.Errorf("proto: unexpected end of group")
)
