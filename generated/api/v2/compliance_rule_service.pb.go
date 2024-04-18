// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v2/compliance_rule_service.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RuleRequest struct {
	RuleName             string    `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Query                *RawQuery `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RuleRequest) Reset()         { *m = RuleRequest{} }
func (m *RuleRequest) String() string { return proto.CompactTextString(m) }
func (*RuleRequest) ProtoMessage()    {}
func (*RuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d47f4ceced26463, []int{0}
}
func (m *RuleRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuleRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleRequest.Merge(m, src)
}
func (m *RuleRequest) XXX_Size() int {
	return m.Size()
}
func (m *RuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RuleRequest proto.InternalMessageInfo

func (m *RuleRequest) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *RuleRequest) GetQuery() *RawQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *RuleRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *RuleRequest) Clone() *RuleRequest {
	if m == nil {
		return nil
	}
	cloned := new(RuleRequest)
	*cloned = *m

	cloned.Query = m.Query.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*RuleRequest)(nil), "v2.RuleRequest")
}

func init() {
	proto.RegisterFile("api/v2/compliance_rule_service.proto", fileDescriptor_3d47f4ceced26463)
}

var fileDescriptor_3d47f4ceced26463 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xcd, 0x81, 0x62, 0x53, 0x41, 0x0c, 0x08, 0xb5, 0xca, 0x51, 0x8a, 0x60, 0x5d, 0x12,
	0x88, 0x83, 0x83, 0x9b, 0x0e, 0x6e, 0x05, 0xcf, 0x45, 0x5c, 0x4a, 0x3c, 0x3e, 0x6a, 0xf0, 0x72,
	0xb9, 0x26, 0xb9, 0x68, 0x11, 0x97, 0xfe, 0x05, 0x17, 0x7f, 0x92, 0xa3, 0xe0, 0x1f, 0x90, 0xea,
	0x0f, 0x91, 0x5c, 0x4a, 0xa5, 0xb8, 0xbe, 0xcf, 0xcb, 0xcb, 0xf3, 0x7d, 0xf8, 0x50, 0x54, 0x92,
	0x79, 0xce, 0x72, 0xad, 0xaa, 0x42, 0x8a, 0x32, 0x87, 0x91, 0xa9, 0x0b, 0x18, 0x59, 0x30, 0x5e,
	0xe6, 0x40, 0x2b, 0xa3, 0x9d, 0x26, 0x89, 0xe7, 0xdd, 0xf4, 0x7f, 0x33, 0xd7, 0x4a, 0xe9, 0x32,
	0x76, 0xba, 0x7b, 0x0b, 0x6e, 0x41, 0x98, 0xfc, 0x7e, 0x34, 0xa9, 0xc1, 0x4c, 0x17, 0xe8, 0x60,
	0xac, 0xf5, 0xb8, 0x00, 0x16, 0x1a, 0xa2, 0x2c, 0xb5, 0x13, 0x4e, 0xea, 0xd2, 0x46, 0xda, 0x1f,
	0xe2, 0x76, 0x56, 0x17, 0x90, 0xc1, 0xa4, 0x06, 0xeb, 0xc8, 0x3e, 0x6e, 0x35, 0x06, 0xa5, 0x50,
	0xd0, 0x41, 0x3d, 0x34, 0x68, 0x65, 0x9b, 0x21, 0x18, 0x0a, 0x05, 0xa4, 0x8f, 0xd7, 0x9b, 0xe1,
	0x4e, 0xd2, 0x43, 0x83, 0x36, 0xdf, 0xa2, 0x9e, 0xd3, 0x4c, 0x3c, 0x5e, 0x85, 0x2c, 0x8b, 0x88,
	0xcf, 0x10, 0xde, 0xbd, 0x58, 0x4a, 0x86, 0xe9, 0xeb, 0x78, 0x0c, 0x91, 0x78, 0xe7, 0x12, 0xdc,
	0x2a, 0x23, 0xdb, 0xcd, 0xc6, 0x9f, 0x40, 0x97, 0x84, 0x60, 0xb5, 0xd4, 0x67, 0xb3, 0xcf, 0x9f,
	0xd7, 0xe4, 0x98, 0x1c, 0xad, 0x7e, 0x80, 0x05, 0x31, 0x66, 0x6b, 0xa5, 0x84, 0x99, 0xb2, 0xe7,
	0xa5, 0xf7, 0xcb, 0xf9, 0xe9, 0xfb, 0x3c, 0x45, 0x1f, 0xf3, 0x14, 0x7d, 0xcd, 0x53, 0xf4, 0xf6,
	0x9d, 0xae, 0xe1, 0x8e, 0xd4, 0xd4, 0x3a, 0x91, 0x3f, 0x18, 0xfd, 0x14, 0x2f, 0xa7, 0xa2, 0x92,
	0xd4, 0xf3, 0xdb, 0x36, 0x65, 0xf1, 0x77, 0x67, 0x9e, 0xdf, 0x24, 0x77, 0x1b, 0x0d, 0x3c, 0xf9,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xab, 0x99, 0x5e, 0x02, 0x9b, 0x01, 0x00, 0x00,
}

func (m *RuleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuleRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuleRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintComplianceRuleService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RuleName) > 0 {
		i -= len(m.RuleName)
		copy(dAtA[i:], m.RuleName)
		i = encodeVarintComplianceRuleService(dAtA, i, uint64(len(m.RuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceRuleService(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceRuleService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RuleRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RuleName)
	if l > 0 {
		n += 1 + l + sovComplianceRuleService(uint64(l))
	}
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovComplianceRuleService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComplianceRuleService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceRuleService(x uint64) (n int) {
	return sovComplianceRuleService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RuleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceRuleService
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
			return fmt.Errorf("proto: RuleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceRuleService
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
				return ErrInvalidLengthComplianceRuleService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceRuleService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceRuleService
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
				return ErrInvalidLengthComplianceRuleService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceRuleService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Query == nil {
				m.Query = &RawQuery{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceRuleService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceRuleService
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
func skipComplianceRuleService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceRuleService
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
					return 0, ErrIntOverflowComplianceRuleService
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
					return 0, ErrIntOverflowComplianceRuleService
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
				return 0, ErrInvalidLengthComplianceRuleService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceRuleService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceRuleService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceRuleService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceRuleService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceRuleService = fmt.Errorf("proto: unexpected end of group")
)