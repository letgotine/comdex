// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/vault/v1beta1/genesis.proto

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

type GenesisState struct {
	Vaults                      []Vault                                `protobuf:"bytes,1,rep,name=vaults,proto3" json:"vaults" yaml:"vaults"`
	StableMintVault             []StableMintVault                      `protobuf:"bytes,2,rep,name=stableMintVault,proto3" json:"stableMintVault" yaml:"stableMintVault"`
	AppExtendedPairVaultMapping []AppExtendedPairVaultMappingData      `protobuf:"bytes,3,rep,name=appExtendedPairVaultMapping,proto3" json:"appExtendedPairVaultMapping" yaml:"appExtendedPairVaultMapping"`
	UserVaultAssetMapping       []OwnerAppExtendedPairVaultMappingData `protobuf:"bytes,4,rep,name=userVaultAssetMapping,proto3" json:"userVaultAssetMapping" yaml:"userVaultAssetMapping"`
	LengthOfVaults              uint64                                 `protobuf:"varint,5,opt,name=lengthOfVaults,proto3" json:"lengthOfVaults,omitempty" yaml:"lengthOfVaults"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ed468466e13f8a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "comdex.vault.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("comdex/vault/v1beta1/genesis.proto", fileDescriptor_00ed468466e13f8a)
}

var fileDescriptor_00ed468466e13f8a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xdb, 0x0b, 0x97, 0x45, 0xef, 0x1f, 0x93, 0x06, 0x4c, 0x05, 0x33, 0x25, 0x13, 0x4d,
	0x88, 0x89, 0x6d, 0xd0, 0xb8, 0x61, 0x47, 0xa3, 0x71, 0x45, 0xd0, 0x62, 0x5c, 0xb8, 0x9b, 0xc2,
	0x50, 0x9a, 0x94, 0xb6, 0x61, 0x06, 0x84, 0xb7, 0x70, 0xe5, 0x0b, 0xb8, 0xf1, 0x51, 0x58, 0xb2,
	0xd3, 0x55, 0xa3, 0xe5, 0x0d, 0x78, 0x02, 0xd3, 0x99, 0x6a, 0x02, 0x19, 0x71, 0xd7, 0xd3, 0xf3,
	0x9d, 0xdf, 0xf7, 0x9d, 0xcc, 0x51, 0x60, 0x37, 0x1c, 0xf6, 0xf0, 0xd4, 0x9c, 0xa0, 0xb1, 0x4f,
	0xcd, 0x49, 0xdd, 0xc1, 0x14, 0xd5, 0x4d, 0x17, 0x07, 0x98, 0x78, 0xc4, 0x88, 0x46, 0x21, 0x0d,
	0xd5, 0x22, 0xd7, 0x18, 0x4c, 0x63, 0x64, 0x9a, 0x72, 0xd1, 0x0d, 0xdd, 0x90, 0x09, 0xcc, 0xf4,
	0x8b, 0x6b, 0xcb, 0x55, 0x21, 0x8f, 0x4f, 0x32, 0x05, 0x7c, 0xc9, 0x2b, 0x7f, 0x2f, 0x39, 0xbf,
	0x43, 0x11, 0xc5, 0xea, 0x8d, 0x52, 0x60, 0x7d, 0xa2, 0xc9, 0xd5, 0x5c, 0xed, 0xcf, 0x49, 0xc5,
	0x10, 0xf9, 0x19, 0xb7, 0x69, 0x65, 0xe9, 0xf3, 0x58, 0x97, 0x92, 0x58, 0x2f, 0xb0, 0x92, 0xac,
	0x62, 0xfd, 0xdf, 0x0c, 0x0d, 0xfd, 0x06, 0xe4, 0x08, 0x68, 0x67, 0x2c, 0x35, 0x54, 0x76, 0x08,
	0x45, 0x8e, 0x8f, 0x5b, 0x5e, 0x40, 0x99, 0x58, 0xfb, 0xc5, 0xf0, 0x87, 0x62, 0x7c, 0x67, 0x5d,
	0x6c, 0x81, 0xd4, 0x68, 0x15, 0xeb, 0xbb, 0x1c, 0xbf, 0xc1, 0x82, 0xf6, 0x26, 0x5d, 0x7d, 0x92,
	0x95, 0x0a, 0x8a, 0xa2, 0x8b, 0x29, 0xc5, 0x41, 0x0f, 0xf7, 0xae, 0x90, 0x37, 0x62, 0x8d, 0x16,
	0x8a, 0x22, 0x2f, 0x70, 0xb5, 0x1c, 0x73, 0x3f, 0x13, 0xbb, 0x37, 0xbf, 0x1f, 0x3c, 0x47, 0x14,
	0x59, 0x47, 0x59, 0x1a, 0xc8, 0xd3, 0x6c, 0xf1, 0x81, 0xf6, 0xb6, 0x14, 0xea, 0xa3, 0xac, 0x94,
	0xc6, 0x04, 0xf3, 0x9f, 0x4d, 0x42, 0xf0, 0x57, 0xbe, 0x3c, 0xcb, 0xd7, 0x10, 0xe7, 0x6b, 0xdf,
	0x07, 0x78, 0xf4, 0x53, 0xc8, 0x83, 0x2c, 0xe4, 0x3e, 0x0f, 0x29, 0xb4, 0x81, 0xb6, 0xd8, 0x5e,
	0x6d, 0x2a, 0xff, 0x7d, 0x1c, 0xb8, 0x74, 0xd0, 0xee, 0xf3, 0xa7, 0xd5, 0x7e, 0x57, 0xe5, 0x5a,
	0xde, 0xda, 0x5b, 0xc5, 0x7a, 0x89, 0x03, 0xd7, 0xfb, 0xd0, 0xde, 0x18, 0xb0, 0xae, 0xe7, 0xef,
	0x40, 0x7a, 0x4e, 0x80, 0x34, 0x4f, 0x80, 0xbc, 0x48, 0x80, 0xfc, 0x96, 0x00, 0xf9, 0x61, 0x09,
	0xa4, 0xc5, 0x12, 0x48, 0xaf, 0x4b, 0x20, 0xdd, 0x99, 0xae, 0x47, 0x07, 0x63, 0x27, 0xdd, 0xd1,
	0xe4, 0x7b, 0x1e, 0x87, 0xfd, 0xbe, 0xd7, 0xf5, 0x90, 0x9f, 0xd5, 0xe6, 0xe7, 0xe9, 0xd2, 0x59,
	0x84, 0x89, 0x53, 0x60, 0x37, 0x7b, 0xfa, 0x11, 0x00, 0x00, 0xff, 0xff, 0x43, 0xaa, 0xf1, 0xcc,
	0x27, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LengthOfVaults != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LengthOfVaults))
		i--
		dAtA[i] = 0x28
	}
	if len(m.UserVaultAssetMapping) > 0 {
		for iNdEx := len(m.UserVaultAssetMapping) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserVaultAssetMapping[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AppExtendedPairVaultMapping) > 0 {
		for iNdEx := len(m.AppExtendedPairVaultMapping) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AppExtendedPairVaultMapping[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StableMintVault) > 0 {
		for iNdEx := len(m.StableMintVault) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StableMintVault[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Vaults) > 0 {
		for iNdEx := len(m.Vaults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vaults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Vaults) > 0 {
		for _, e := range m.Vaults {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StableMintVault) > 0 {
		for _, e := range m.StableMintVault {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AppExtendedPairVaultMapping) > 0 {
		for _, e := range m.AppExtendedPairVaultMapping {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserVaultAssetMapping) > 0 {
		for _, e := range m.UserVaultAssetMapping {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LengthOfVaults != 0 {
		n += 1 + sovGenesis(uint64(m.LengthOfVaults))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vaults = append(m.Vaults, Vault{})
			if err := m.Vaults[len(m.Vaults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableMintVault", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StableMintVault = append(m.StableMintVault, StableMintVault{})
			if err := m.StableMintVault[len(m.StableMintVault)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppExtendedPairVaultMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppExtendedPairVaultMapping = append(m.AppExtendedPairVaultMapping, AppExtendedPairVaultMappingData{})
			if err := m.AppExtendedPairVaultMapping[len(m.AppExtendedPairVaultMapping)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserVaultAssetMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserVaultAssetMapping = append(m.UserVaultAssetMapping, OwnerAppExtendedPairVaultMappingData{})
			if err := m.UserVaultAssetMapping[len(m.UserVaultAssetMapping)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LengthOfVaults", wireType)
			}
			m.LengthOfVaults = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LengthOfVaults |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
