// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/compliance_integration.proto

package storage

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

// Represents the status of compliance operator
type COStatus int32

const (
	COStatus_HEALTHY   COStatus = 0
	COStatus_UNHEALTHY COStatus = 1
)

var COStatus_name = map[int32]string{
	0: "HEALTHY",
	1: "UNHEALTHY",
}

var COStatus_value = map[string]int32{
	"HEALTHY":   0,
	"UNHEALTHY": 1,
}

func (x COStatus) String() string {
	return proto.EnumName(COStatus_name, int32(x))
}

func (COStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14e26a23cbdbee2c, []int{0}
}

// Next Tag: 8
type ComplianceIntegration struct {
	Id                  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Compliance Operator Integration ID,hidden,store" sql:"pk,type(uuid)"`                                // @gotags: search:"Compliance Operator Integration ID,hidden,store" sql:"pk,type(uuid)"
	Version             string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" search:"Compliance Operator Version,hidden,store"`                      // @gotags: search:"Compliance Operator Version,hidden,store"
	ClusterId           string `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid),index=category:unique;name:compliance_unique_indicator"` // @gotags: search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid),index=category:unique;name:compliance_unique_indicator"
	ComplianceNamespace string `protobuf:"bytes,4,opt,name=compliance_namespace,json=complianceNamespace,proto3" json:"compliance_namespace,omitempty"`
	// Collection of errors that occurred while trying to obtain compliance health info.
	StatusErrors         []string `protobuf:"bytes,5,rep,name=status_errors,json=statusErrors,proto3" json:"status_errors,omitempty"`
	OperatorInstalled    bool     `protobuf:"varint,6,opt,name=operator_installed,json=operatorInstalled,proto3" json:"operator_installed,omitempty" search:"Compliance Operator Installed,hidden"`              // @gotags: search:"Compliance Operator Installed,hidden"
	OperatorStatus       COStatus `protobuf:"varint,7,opt,name=operator_status,json=operatorStatus,proto3,enum=storage.COStatus" json:"operator_status,omitempty" search:"Compliance Operator Status"` //@gotags: search:"Compliance Operator Status"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplianceIntegration) Reset()         { *m = ComplianceIntegration{} }
func (m *ComplianceIntegration) String() string { return proto.CompactTextString(m) }
func (*ComplianceIntegration) ProtoMessage()    {}
func (*ComplianceIntegration) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e26a23cbdbee2c, []int{0}
}
func (m *ComplianceIntegration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceIntegration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceIntegration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceIntegration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceIntegration.Merge(m, src)
}
func (m *ComplianceIntegration) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceIntegration) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceIntegration.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceIntegration proto.InternalMessageInfo

func (m *ComplianceIntegration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ComplianceIntegration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ComplianceIntegration) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ComplianceIntegration) GetComplianceNamespace() string {
	if m != nil {
		return m.ComplianceNamespace
	}
	return ""
}

func (m *ComplianceIntegration) GetStatusErrors() []string {
	if m != nil {
		return m.StatusErrors
	}
	return nil
}

func (m *ComplianceIntegration) GetOperatorInstalled() bool {
	if m != nil {
		return m.OperatorInstalled
	}
	return false
}

func (m *ComplianceIntegration) GetOperatorStatus() COStatus {
	if m != nil {
		return m.OperatorStatus
	}
	return COStatus_HEALTHY
}

func (m *ComplianceIntegration) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ComplianceIntegration) Clone() *ComplianceIntegration {
	if m == nil {
		return nil
	}
	cloned := new(ComplianceIntegration)
	*cloned = *m

	if m.StatusErrors != nil {
		cloned.StatusErrors = make([]string, len(m.StatusErrors))
		copy(cloned.StatusErrors, m.StatusErrors)
	}
	return cloned
}

func init() {
	proto.RegisterEnum("storage.COStatus", COStatus_name, COStatus_value)
	proto.RegisterType((*ComplianceIntegration)(nil), "storage.ComplianceIntegration")
}

func init() {
	proto.RegisterFile("storage/compliance_integration.proto", fileDescriptor_14e26a23cbdbee2c)
}

var fileDescriptor_14e26a23cbdbee2c = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x97, 0x4e, 0xd7, 0xf5, 0xd3, 0x4d, 0x17, 0x15, 0xe2, 0xc1, 0x52, 0x54, 0xa4, 0x08,
	0x76, 0xa8, 0x37, 0x05, 0x41, 0xc7, 0x60, 0x03, 0x99, 0x50, 0xf5, 0xa0, 0x97, 0x12, 0xdb, 0x30,
	0x82, 0x5d, 0x53, 0x92, 0x4c, 0x7c, 0x14, 0xef, 0xbe, 0x8c, 0x47, 0x1f, 0x41, 0xea, 0x8b, 0x08,
	0x6d, 0xb3, 0x79, 0x4a, 0xfe, 0xff, 0xdf, 0x8f, 0xe4, 0xe3, 0x83, 0x43, 0xa5, 0x85, 0xa4, 0x53,
	0xd6, 0x8f, 0xc5, 0x2c, 0x4f, 0x39, 0xcd, 0x62, 0x16, 0xf1, 0x4c, 0xb3, 0xa9, 0xa4, 0x9a, 0x8b,
	0x2c, 0xc8, 0xa5, 0xd0, 0x02, 0xdb, 0xb5, 0xb5, 0xff, 0x69, 0xc1, 0xce, 0x60, 0x61, 0x8e, 0x97,
	0x22, 0xee, 0x82, 0xc5, 0x13, 0x82, 0x3c, 0xe4, 0x3b, 0xa1, 0xc5, 0x13, 0x4c, 0xc0, 0x7e, 0x63,
	0x52, 0x71, 0x91, 0x11, 0xab, 0x2c, 0x4d, 0xc4, 0x7b, 0x00, 0x71, 0x3a, 0x57, 0x9a, 0xc9, 0x88,
	0x27, 0xa4, 0x59, 0x42, 0xa7, 0x6e, 0xc6, 0x09, 0x3e, 0x85, 0xed, 0x7f, 0xb3, 0x64, 0x74, 0xc6,
	0x54, 0x4e, 0x63, 0x46, 0x56, 0x4a, 0x71, 0x6b, 0xc9, 0x26, 0x06, 0xe1, 0x03, 0xe8, 0x28, 0x4d,
	0xf5, 0x5c, 0x45, 0x4c, 0x4a, 0x21, 0x15, 0x59, 0xf5, 0x9a, 0xbe, 0x13, 0xae, 0x57, 0xe5, 0xb0,
	0xec, 0xf0, 0x09, 0x60, 0x91, 0x33, 0x49, 0xb5, 0x90, 0x11, 0xcf, 0x94, 0xa6, 0x69, 0xca, 0x12,
	0xd2, 0xf2, 0x90, 0xdf, 0x0e, 0x7b, 0x86, 0x8c, 0x0d, 0xc0, 0x17, 0xb0, 0xb1, 0xd0, 0xab, 0x77,
	0x88, 0xed, 0x21, 0xbf, 0x7b, 0xd6, 0x0b, 0xea, 0x65, 0x04, 0x83, 0xbb, 0xfb, 0x12, 0x84, 0x5d,
	0x63, 0x56, 0xf9, 0xf8, 0x08, 0xda, 0x86, 0xe1, 0x35, 0xb0, 0x47, 0xc3, 0xeb, 0xdb, 0x87, 0xd1,
	0xd3, 0x66, 0x03, 0x77, 0xc0, 0x79, 0x9c, 0x98, 0x88, 0x6e, 0xae, 0xbe, 0x0a, 0x17, 0x7d, 0x17,
	0x2e, 0xfa, 0x29, 0x5c, 0xf4, 0xf1, 0xeb, 0x36, 0x60, 0x97, 0x8b, 0x40, 0x69, 0x1a, 0xbf, 0x4a,
	0xf1, 0x5e, 0xed, 0xde, 0xfc, 0xf6, 0xdc, 0x0b, 0xfa, 0xf5, 0xf5, 0xb2, 0x3e, 0x5f, 0x5a, 0xa5,
	0x71, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x6d, 0x1f, 0x7d, 0xc5, 0x01, 0x00, 0x00,
}

func (m *ComplianceIntegration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceIntegration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceIntegration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OperatorStatus != 0 {
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(m.OperatorStatus))
		i--
		dAtA[i] = 0x38
	}
	if m.OperatorInstalled {
		i--
		if m.OperatorInstalled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.StatusErrors) > 0 {
		for iNdEx := len(m.StatusErrors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StatusErrors[iNdEx])
			copy(dAtA[i:], m.StatusErrors[iNdEx])
			i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.StatusErrors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ComplianceNamespace) > 0 {
		i -= len(m.ComplianceNamespace)
		copy(dAtA[i:], m.ComplianceNamespace)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.ComplianceNamespace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceIntegration(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceIntegration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceIntegration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.ComplianceNamespace)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	if len(m.StatusErrors) > 0 {
		for _, s := range m.StatusErrors {
			l = len(s)
			n += 1 + l + sovComplianceIntegration(uint64(l))
		}
	}
	if m.OperatorInstalled {
		n += 2
	}
	if m.OperatorStatus != 0 {
		n += 1 + sovComplianceIntegration(uint64(m.OperatorStatus))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComplianceIntegration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceIntegration(x uint64) (n int) {
	return sovComplianceIntegration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceIntegration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceIntegration
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
			return fmt.Errorf("proto: ComplianceIntegration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceIntegration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComplianceNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComplianceNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusErrors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusErrors = append(m.StatusErrors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorInstalled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
			m.OperatorInstalled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorStatus", wireType)
			}
			m.OperatorStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OperatorStatus |= COStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceIntegration
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
func skipComplianceIntegration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceIntegration
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
					return 0, ErrIntOverflowComplianceIntegration
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
					return 0, ErrIntOverflowComplianceIntegration
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
				return 0, ErrInvalidLengthComplianceIntegration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceIntegration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceIntegration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceIntegration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceIntegration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceIntegration = fmt.Errorf("proto: unexpected end of group")
)
