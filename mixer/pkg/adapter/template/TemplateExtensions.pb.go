// Code generated by protoc-gen-gogo.
// source: mixer/pkg/adapter/template/TemplateExtensions.proto
// DO NOT EDIT!

/*
Package istio_mixer_v1_config_template is a generated protocol buffer package.

It is generated from these files:
	mixer/pkg/adapter/template/TemplateExtensions.proto

It has these top-level messages:
	Expr
*/
package istio_mixer_v1_config_template

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TemplateVariety int32

const (
	TEMPLATE_VARIETY_CHECK  TemplateVariety = 0
	TEMPLATE_VARIETY_REPORT TemplateVariety = 1
	TEMPLATE_VARIETY_QUOTA  TemplateVariety = 2
)

var TemplateVariety_name = map[int32]string{
	0: "TEMPLATE_VARIETY_CHECK",
	1: "TEMPLATE_VARIETY_REPORT",
	2: "TEMPLATE_VARIETY_QUOTA",
}
var TemplateVariety_value = map[string]int32{
	"TEMPLATE_VARIETY_CHECK":  0,
	"TEMPLATE_VARIETY_REPORT": 1,
	"TEMPLATE_VARIETY_QUOTA":  2,
}

func (TemplateVariety) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTemplateExtensions, []int{0}
}

type Expr struct {
}

func (m *Expr) Reset()                    { *m = Expr{} }
func (*Expr) ProtoMessage()               {}
func (*Expr) Descriptor() ([]byte, []int) { return fileDescriptorTemplateExtensions, []int{0} }

var E_TemplateVariety = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*TemplateVariety)(nil),
	Field:         72295727,
	Name:          "istio.mixer.v1.config.template.template_variety",
	Tag:           "varint,72295727,opt,name=template_variety,json=templateVariety,enum=istio.mixer.v1.config.template.TemplateVariety",
	Filename:      "mixer/pkg/adapter/template/TemplateExtensions.proto",
}

func init() {
	proto.RegisterType((*Expr)(nil), "istio.mixer.v1.config.template.Expr")
	proto.RegisterEnum("istio.mixer.v1.config.template.TemplateVariety", TemplateVariety_name, TemplateVariety_value)
	proto.RegisterExtension(E_TemplateVariety)
}
func (x TemplateVariety) String() string {
	s, ok := TemplateVariety_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Expr) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Expr)
	if !ok {
		that2, ok := that.(Expr)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *Expr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_config_template.Expr{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTemplateExtensions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Expr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Expr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64TemplateExtensions(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32TemplateExtensions(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTemplateExtensions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Expr) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovTemplateExtensions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateExtensions(x uint64) (n int) {
	return sovTemplateExtensions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Expr) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Expr{`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateExtensions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Expr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateExtensions
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Expr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Expr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateExtensions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateExtensions
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
func skipTemplateExtensions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateExtensions
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
					return 0, ErrIntOverflowTemplateExtensions
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
					return 0, ErrIntOverflowTemplateExtensions
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTemplateExtensions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateExtensions
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
				next, err := skipTemplateExtensions(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthTemplateExtensions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateExtensions   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/pkg/adapter/template/TemplateExtensions.proto", fileDescriptorTemplateExtensions)
}

var fileDescriptorTemplateExtensions = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xce, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0x4c, 0x49, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0x2f, 0x49,
	0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x0f, 0x81, 0x32, 0x5c, 0x2b, 0x4a, 0x52, 0xf3, 0x8a,
	0x33, 0xf3, 0xf3, 0x8a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x32, 0x8b, 0x4b, 0x32,
	0xf3, 0xf5, 0xc0, 0x5a, 0xf5, 0xca, 0x0c, 0xf5, 0x92, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0xf5, 0x60,
	0x1a, 0xa5, 0x14, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xaa, 0x93, 0x4a, 0xd3, 0xf4,
	0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x20, 0x26, 0x28, 0xb1, 0x71, 0xb1,
	0xb8, 0x56, 0x14, 0x14, 0x69, 0xa5, 0x71, 0xf1, 0xc3, 0x6c, 0x09, 0x4b, 0x2c, 0xca, 0x4c, 0x2d,
	0xa9, 0x14, 0x92, 0xe2, 0x12, 0x0b, 0x71, 0xf5, 0x0d, 0xf0, 0x71, 0x0c, 0x71, 0x8d, 0x0f, 0x73,
	0x0c, 0xf2, 0x74, 0x0d, 0x89, 0x8c, 0x77, 0xf6, 0x70, 0x75, 0xf6, 0x16, 0x60, 0x10, 0x92, 0xe6,
	0x12, 0xc7, 0x90, 0x0b, 0x72, 0x0d, 0xf0, 0x0f, 0x0a, 0x11, 0x60, 0xc4, 0xaa, 0x31, 0x30, 0xd4,
	0x3f, 0xc4, 0x51, 0x80, 0xc9, 0xaa, 0x9a, 0x4b, 0x00, 0xe6, 0xba, 0xf8, 0x32, 0xa8, 0x45, 0x32,
	0x7a, 0x10, 0x67, 0xea, 0xc1, 0x9c, 0xa9, 0xe7, 0x96, 0x99, 0x93, 0xea, 0x5f, 0x50, 0x02, 0xf2,
	0xa9, 0xc4, 0xfa, 0x53, 0x7b, 0x94, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0xf4, 0xf5, 0xf0, 0xfb, 0x56,
	0x0f, 0xcd, 0x03, 0x41, 0xfc, 0x25, 0xa8, 0x02, 0x4e, 0x3a, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78,
	0x28, 0xc7, 0xf0, 0xe1, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18,
	0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x1c, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x32, 0xa5, 0xd8, 0x77, 0x9a, 0x01, 0x00, 0x00,
}
