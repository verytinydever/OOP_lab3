// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/asset/nft/v1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// EventClassIssued is emitted on MsgIssueClass.
type EventClassIssued struct {
	ID          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Issuer      string         `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Symbol      string         `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name        string         `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string         `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	URI         string         `protobuf:"bytes,6,opt,name=uri,proto3" json:"uri,omitempty"`
	URIHash     string         `protobuf:"bytes,7,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	Features    []ClassFeature `protobuf:"varint,8,rep,packed,name=features,proto3,enum=coreum.asset.nft.v1.ClassFeature" json:"features,omitempty"`
}

func (m *EventClassIssued) Reset()         { *m = EventClassIssued{} }
func (m *EventClassIssued) String() string { return proto.CompactTextString(m) }
func (*EventClassIssued) ProtoMessage()    {}
func (*EventClassIssued) Descriptor() ([]byte, []int) {
	return fileDescriptor_fef75aa7da633196, []int{0}
}
func (m *EventClassIssued) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClassIssued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClassIssued.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClassIssued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClassIssued.Merge(m, src)
}
func (m *EventClassIssued) XXX_Size() int {
	return m.Size()
}
func (m *EventClassIssued) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClassIssued.DiscardUnknown(m)
}

var xxx_messageInfo_EventClassIssued proto.InternalMessageInfo

func (m *EventClassIssued) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EventClassIssued) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *EventClassIssued) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EventClassIssued) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventClassIssued) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventClassIssued) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *EventClassIssued) GetURIHash() string {
	if m != nil {
		return m.URIHash
	}
	return ""
}

func (m *EventClassIssued) GetFeatures() []ClassFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterType((*EventClassIssued)(nil), "coreum.asset.nft.v1.EventClassIssued")
}

func init() { proto.RegisterFile("coreum/asset/nft/v1/event.proto", fileDescriptor_fef75aa7da633196) }

var fileDescriptor_fef75aa7da633196 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4e, 0xea, 0x40,
	0x18, 0xc5, 0x69, 0xcb, 0x6d, 0xb9, 0x43, 0x72, 0x73, 0x33, 0x1a, 0x33, 0x92, 0xd8, 0xa2, 0x0b,
	0xc3, 0xaa, 0x0d, 0xea, 0xd6, 0x0d, 0x28, 0xb1, 0x1b, 0x17, 0x93, 0xb0, 0x71, 0x63, 0x06, 0x3a,
	0xd0, 0x49, 0x68, 0x87, 0xcc, 0x1f, 0x22, 0x6f, 0x61, 0x7c, 0x2a, 0x97, 0x2c, 0x5d, 0x11, 0x53,
	0x5e, 0xc4, 0xcc, 0x94, 0x18, 0x16, 0xec, 0xbe, 0xef, 0x77, 0xce, 0x77, 0xda, 0x39, 0x20, 0x9a,
	0x72, 0x41, 0x75, 0x91, 0x10, 0x29, 0xa9, 0x4a, 0xca, 0x99, 0x4a, 0x56, 0xfd, 0x84, 0xae, 0x68,
	0xa9, 0xe2, 0xa5, 0xe0, 0x8a, 0xc3, 0x93, 0xda, 0x10, 0x5b, 0x43, 0x5c, 0xce, 0x54, 0xbc, 0xea,
	0x77, 0x4e, 0xe7, 0x7c, 0xce, 0xad, 0x9e, 0x98, 0xa9, 0xb6, 0x76, 0x2e, 0x8e, 0x65, 0x99, 0x0b,
	0x2b, 0x5f, 0x7d, 0xb8, 0xe0, 0xff, 0xa3, 0x49, 0x1e, 0x2e, 0x88, 0x94, 0xa9, 0x94, 0x9a, 0x66,
	0xf0, 0x0c, 0xb8, 0x2c, 0x43, 0x4e, 0xd7, 0xe9, 0xfd, 0x1d, 0xf8, 0xd5, 0x36, 0x72, 0xd3, 0x07,
	0xec, 0x32, 0xc3, 0x7d, 0x66, 0x1c, 0x02, 0xb9, 0x46, 0xc3, 0xfb, 0xcd, 0x70, 0xb9, 0x2e, 0x26,
	0x7c, 0x81, 0xbc, 0x9a, 0xd7, 0x1b, 0x84, 0xa0, 0x59, 0x92, 0x82, 0xa2, 0xa6, 0xa5, 0x76, 0x86,
	0x5d, 0xd0, 0xce, 0xa8, 0x9c, 0x0a, 0xb6, 0x54, 0x8c, 0x97, 0xe8, 0x8f, 0x95, 0x0e, 0x11, 0x3c,
	0x07, 0x9e, 0x16, 0x0c, 0xf9, 0xf6, 0xf3, 0x41, 0xb5, 0x8d, 0xbc, 0x31, 0x4e, 0xb1, 0x61, 0xf0,
	0x1a, 0xb4, 0xb4, 0x60, 0xaf, 0x39, 0x91, 0x39, 0x0a, 0xac, 0xde, 0xae, 0xb6, 0x51, 0x30, 0xc6,
	0xe9, 0x13, 0x91, 0x39, 0x0e, 0xb4, 0x60, 0x66, 0x80, 0xf7, 0xa0, 0x35, 0xa3, 0x44, 0x69, 0x41,
	0x25, 0x6a, 0x75, 0xbd, 0xde, 0xbf, 0x9b, 0xcb, 0xf8, 0x48, 0x65, 0xb1, 0x7d, 0xf4, 0xa8, 0x76,
	0xe2, 0xdf, 0x93, 0xc1, 0xf3, 0x67, 0x15, 0x3a, 0x9b, 0x2a, 0x74, 0xbe, 0xab, 0xd0, 0x79, 0xdf,
	0x85, 0x8d, 0xcd, 0x2e, 0x6c, 0x7c, 0xed, 0xc2, 0xc6, 0xcb, 0xdd, 0x9c, 0xa9, 0x5c, 0x4f, 0xe2,
	0x29, 0x2f, 0x92, 0xa1, 0x0d, 0x1c, 0x71, 0x5d, 0x66, 0xc4, 0xfc, 0x78, 0xb2, 0x6f, 0xfa, 0xed,
	0xa0, 0x6b, 0xb5, 0x5e, 0x52, 0x39, 0xf1, 0x6d, 0xd7, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0x6c, 0x63, 0x4e, 0xd8, 0x01, 0x00, 0x00,
}

func (m *EventClassIssued) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClassIssued) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClassIssued) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Features) > 0 {
		dAtA2 := make([]byte, len(m.Features)*10)
		var j1 int
		for _, num := range m.Features {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvent(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x42
	}
	if len(m.URIHash) > 0 {
		i -= len(m.URIHash)
		copy(dAtA[i:], m.URIHash)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.URIHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventClassIssued) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.URIHash)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Features) > 0 {
		l = 0
		for _, e := range m.Features {
			l += sovEvent(uint64(e))
		}
		n += 1 + sovEvent(uint64(l)) + l
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventClassIssued) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClassIssued: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClassIssued: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URIHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URIHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v ClassFeature
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ClassFeature(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Features = append(m.Features, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvent
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvent
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Features) == 0 {
					m.Features = make([]ClassFeature, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ClassFeature
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ClassFeature(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Features = append(m.Features, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
