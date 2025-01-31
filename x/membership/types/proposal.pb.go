// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: membershipmodule/membership/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// DirectDemocracyUpdateProposal updates the guardians and the total voting weight
type DirectDemocracyUpdateProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Guardians to add
	GuardiansToAdd []string `protobuf:"bytes,3,rep,name=guardians_to_add,json=guardiansToAdd,proto3" json:"guardians_to_add,omitempty"`
	// Guardians to remove
	GuardiansToRemove []string `protobuf:"bytes,4,rep,name=guardians_to_remove,json=guardiansToRemove,proto3" json:"guardians_to_remove,omitempty"`
	// Total voting weight
	TotalVotingWeight *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=total_voting_weight,json=totalVotingWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_voting_weight,omitempty"`
}

func (m *DirectDemocracyUpdateProposal) Reset()      { *m = DirectDemocracyUpdateProposal{} }
func (*DirectDemocracyUpdateProposal) ProtoMessage() {}
func (*DirectDemocracyUpdateProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d31b03cbd2c6725, []int{0}
}
func (m *DirectDemocracyUpdateProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DirectDemocracyUpdateProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DirectDemocracyUpdateProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DirectDemocracyUpdateProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectDemocracyUpdateProposal.Merge(m, src)
}
func (m *DirectDemocracyUpdateProposal) XXX_Size() int {
	return m.Size()
}
func (m *DirectDemocracyUpdateProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectDemocracyUpdateProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DirectDemocracyUpdateProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DirectDemocracyUpdateProposal)(nil), "membershipmodule.membership.DirectDemocracyUpdateProposal")
}

func init() {
	proto.RegisterFile("membershipmodule/membership/proposal.proto", fileDescriptor_5d31b03cbd2c6725)
}

var fileDescriptor_5d31b03cbd2c6725 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x93, 0x5a, 0x4b, 0x4d, 0xa1, 0xb4, 0xd1, 0x43, 0xb0, 0x98, 0x88, 0xd0, 0x22, 0xa5,
	0x26, 0x87, 0x9e, 0xda, 0x5b, 0xc5, 0x63, 0xa1, 0xc5, 0xd6, 0x16, 0x7a, 0x09, 0x63, 0x66, 0x88,
	0x43, 0x33, 0x79, 0xc3, 0xcc, 0xd3, 0xd6, 0xd3, 0x1e, 0x77, 0x8f, 0x7b, 0xdc, 0xa3, 0x1f, 0xc7,
	0xa3, 0xc7, 0x65, 0x0f, 0xb2, 0xe8, 0x6d, 0x3f, 0xc5, 0x92, 0x51, 0xd6, 0xac, 0x78, 0x4a, 0xf2,
	0x7e, 0xff, 0xfc, 0xe6, 0x0d, 0xef, 0x39, 0xef, 0x05, 0x13, 0x63, 0xa6, 0xf4, 0x84, 0x4b, 0x01,
	0x74, 0x9a, 0xb1, 0xe8, 0x50, 0x88, 0xa4, 0x02, 0x09, 0x9a, 0x64, 0xa1, 0x54, 0x80, 0xe0, 0xbe,
	0x39, 0xce, 0x86, 0x87, 0x42, 0xb3, 0x91, 0x42, 0x0a, 0x26, 0x17, 0x15, 0x6f, 0xbb, 0x5f, 0x3a,
	0xe7, 0x15, 0xa7, 0x35, 0xe0, 0x8a, 0x25, 0x38, 0x60, 0x02, 0x12, 0x45, 0x92, 0xf9, 0x48, 0x52,
	0x82, 0xec, 0xfb, 0x5e, 0xed, 0x36, 0x9c, 0x2a, 0x72, 0xcc, 0x98, 0x67, 0xb7, 0xed, 0x6e, 0x6d,
	0xb8, 0xfb, 0x70, 0xdb, 0xce, 0x0b, 0xca, 0x74, 0xa2, 0xb8, 0x44, 0x0e, 0xb9, 0xf7, 0xc4, 0xb0,
	0x72, 0xc9, 0xfd, 0xea, 0xbc, 0x4a, 0xa7, 0x44, 0x51, 0x4e, 0x72, 0x1d, 0x23, 0xc4, 0x84, 0x52,
	0xaf, 0xd2, 0xae, 0x74, 0x6b, 0xfd, 0xce, 0x72, 0x1d, 0xd8, 0x77, 0xeb, 0xa0, 0x79, 0xcc, 0x3f,
	0x80, 0xe0, 0xc8, 0x84, 0xc4, 0xf9, 0xf0, 0xe5, 0x03, 0xfb, 0x09, 0x5f, 0x28, 0x75, 0x47, 0x4e,
	0xfd, 0x51, 0x5a, 0x31, 0x01, 0x33, 0xe6, 0x3d, 0x35, 0xc2, 0xb7, 0x7b, 0x61, 0xeb, 0x44, 0xa4,
	0xe4, 0x7c, 0x5d, 0x72, 0x0e, 0x0d, 0x74, 0xcf, 0x9c, 0x3a, 0x02, 0x92, 0x2c, 0x9e, 0x01, 0xf2,
	0x3c, 0x8d, 0xff, 0x31, 0x9e, 0x4e, 0xd0, 0xab, 0x16, 0xd7, 0xe9, 0x7f, 0x2b, 0xb4, 0x37, 0xeb,
	0xe0, 0x5d, 0xca, 0x71, 0x32, 0x1d, 0x87, 0x09, 0x88, 0x28, 0x01, 0x2d, 0x40, 0xef, 0x1f, 0x3d,
	0x4d, 0xff, 0x46, 0x38, 0x97, 0x4c, 0x87, 0x03, 0x96, 0x14, 0x0d, 0x9c, 0x90, 0x95, 0x1b, 0x30,
	0xf8, 0x97, 0xa1, 0xbf, 0x0d, 0xfc, 0xfc, 0xfc, 0x62, 0x11, 0x58, 0x57, 0x8b, 0xc0, 0xea, 0xff,
	0x58, 0x6e, 0x7c, 0x7b, 0xb5, 0xf1, 0xed, 0xdb, 0x8d, 0x6f, 0x5f, 0x6e, 0x7d, 0x6b, 0xb5, 0xf5,
	0xad, 0xeb, 0xad, 0x6f, 0xfd, 0xf9, 0x54, 0x3a, 0x3f, 0x07, 0xc5, 0x49, 0x2f, 0x67, 0x18, 0xed,
	0x26, 0xdc, 0x2b, 0x6d, 0xc3, 0xff, 0xf2, 0x6a, 0x98, 0xb6, 0xc6, 0xcf, 0xcc, 0x94, 0x3f, 0xde,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xc5, 0x65, 0x67, 0x46, 0x02, 0x00, 0x00,
}

func (m *DirectDemocracyUpdateProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectDemocracyUpdateProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DirectDemocracyUpdateProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalVotingWeight != nil {
		{
			size := m.TotalVotingWeight.Size()
			i -= size
			if _, err := m.TotalVotingWeight.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GuardiansToRemove) > 0 {
		for iNdEx := len(m.GuardiansToRemove) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GuardiansToRemove[iNdEx])
			copy(dAtA[i:], m.GuardiansToRemove[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.GuardiansToRemove[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GuardiansToAdd) > 0 {
		for iNdEx := len(m.GuardiansToAdd) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GuardiansToAdd[iNdEx])
			copy(dAtA[i:], m.GuardiansToAdd[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.GuardiansToAdd[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DirectDemocracyUpdateProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.GuardiansToAdd) > 0 {
		for _, s := range m.GuardiansToAdd {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.GuardiansToRemove) > 0 {
		for _, s := range m.GuardiansToRemove {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if m.TotalVotingWeight != nil {
		l = m.TotalVotingWeight.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DirectDemocracyUpdateProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: DirectDemocracyUpdateProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectDemocracyUpdateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardiansToAdd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardiansToAdd = append(m.GuardiansToAdd, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardiansToRemove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardiansToRemove = append(m.GuardiansToRemove, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.TotalVotingWeight = &v
			if err := m.TotalVotingWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
