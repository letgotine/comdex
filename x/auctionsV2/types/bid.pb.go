// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/auctionsV2/v1beta1/bid.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Bid struct {
	BiddingId             uint64                                  `protobuf:"varint,1,opt,name=bidding_id,json=biddingId,proto3" json:"bidding_id,omitempty" yaml:"bidding_id"`
	AuctionId             uint64                                  `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	CollateralTokenAmount github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=collateral_token_amount,json=collateralTokenAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateral_token_amount" yaml:"outflow_token_amount"`
	DebtTokenAmount       github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=debt_token_amount,json=debtTokenAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"debt_token_amount" yaml:"inflow_token_amount"`
	BidderAddress         string                                  `protobuf:"bytes,5,opt,name=bidder_address,json=bidderAddress,proto3" json:"bidder_address,omitempty" yaml:"bidder"`
	BiddingTimestamp      time.Time                               `protobuf:"bytes,6,opt,name=bidding_timestamp,json=biddingTimestamp,proto3,stdtime" json:"bidding_timestamp" yaml:"bidding_timestamp"`
	AppId                 uint64                                  `protobuf:"varint,7,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	BidType               string                                  `protobuf:"bytes,8,opt,name=bid_type,json=bidType,proto3" json:"bid_type,omitempty" yaml:"bid_type"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f6db8f3a6a396ec, []int{0}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetBiddingId() uint64 {
	if m != nil {
		return m.BiddingId
	}
	return 0
}

func (m *Bid) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *Bid) GetCollateralTokenAmount() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.CollateralTokenAmount
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Bid) GetDebtTokenAmount() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.DebtTokenAmount
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Bid) GetBidderAddress() string {
	if m != nil {
		return m.BidderAddress
	}
	return ""
}

func (m *Bid) GetBiddingTimestamp() time.Time {
	if m != nil {
		return m.BiddingTimestamp
	}
	return time.Time{}
}

func (m *Bid) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Bid) GetBidType() string {
	if m != nil {
		return m.BidType
	}
	return ""
}

type LimitOrderBid struct {
	LimitOrderBiddingId uint64                                  `protobuf:"varint,1,opt,name=limit_order_bidding_id,json=limitOrderBiddingId,proto3" json:"limit_order_bidding_id,omitempty" yaml:"limit_order_bidding_id"`
	BidderAddress       string                                  `protobuf:"bytes,2,opt,name=bidder_address,json=bidderAddress,proto3" json:"bidder_address,omitempty" yaml:"bidder"`
	CollateralToken     github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=collateral_token,json=collateralToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateral_token" yaml:"outflow_token"`
	DebtToken           github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=debt_token,json=debtToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"debt_token" yaml:"debt_token"`
	BiddingId           []uint64                                `protobuf:"varint,5,rep,packed,name=bidding_id,json=biddingId,proto3" json:"bidding_id,omitempty" yaml:"bidding_id"`
	PremiumDiscount     github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,6,opt,name=premium_discount,json=premiumDiscount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"premium_discount" yaml:"premium_discount"`
}

func (m *LimitOrderBid) Reset()         { *m = LimitOrderBid{} }
func (m *LimitOrderBid) String() string { return proto.CompactTextString(m) }
func (*LimitOrderBid) ProtoMessage()    {}
func (*LimitOrderBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f6db8f3a6a396ec, []int{1}
}
func (m *LimitOrderBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderBid.Merge(m, src)
}
func (m *LimitOrderBid) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderBid) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderBid.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderBid proto.InternalMessageInfo

func (m *LimitOrderBid) GetLimitOrderBiddingId() uint64 {
	if m != nil {
		return m.LimitOrderBiddingId
	}
	return 0
}

func (m *LimitOrderBid) GetBidderAddress() string {
	if m != nil {
		return m.BidderAddress
	}
	return ""
}

func (m *LimitOrderBid) GetCollateralToken() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.CollateralToken
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *LimitOrderBid) GetDebtToken() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.DebtToken
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *LimitOrderBid) GetBiddingId() []uint64 {
	if m != nil {
		return m.BiddingId
	}
	return nil
}

type AuctionParams struct {
	AuctionDurationSeconds uint64                                 `protobuf:"varint,1,opt,name=auction_duration_seconds,json=auctionDurationSeconds,proto3" json:"auction_duration_seconds,omitempty" yaml:"auction_duration_seconds"`
	Step                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=step,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"step" yaml:"step"`
	WithdrawalFee          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=withdrawal_fee,json=withdrawalFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"withdrawal_fee" yaml:"withdrawal_fee"`
	ClosingFee             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=closing_fee,json=closingFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"closing_fee" yaml:"closing_fee"`
}

func (m *AuctionParams) Reset()         { *m = AuctionParams{} }
func (m *AuctionParams) String() string { return proto.CompactTextString(m) }
func (*AuctionParams) ProtoMessage()    {}
func (*AuctionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f6db8f3a6a396ec, []int{2}
}
func (m *AuctionParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionParams.Merge(m, src)
}
func (m *AuctionParams) XXX_Size() int {
	return m.Size()
}
func (m *AuctionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionParams.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionParams proto.InternalMessageInfo

func (m *AuctionParams) GetAuctionDurationSeconds() uint64 {
	if m != nil {
		return m.AuctionDurationSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Bid)(nil), "comdex.auctionsV2.v1beta1.Bid")
	proto.RegisterType((*LimitOrderBid)(nil), "comdex.auctionsV2.v1beta1.LimitOrderBid")
	proto.RegisterType((*AuctionParams)(nil), "comdex.auctionsV2.v1beta1.AuctionParams")
}

func init() {
	proto.RegisterFile("comdex/auctionsV2/v1beta1/bid.proto", fileDescriptor_6f6db8f3a6a396ec)
}

var fileDescriptor_6f6db8f3a6a396ec = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0x8d, 0xbf, 0xfc, 0x7c, 0xcd, 0x54, 0x69, 0x1b, 0xf7, 0x2f, 0x0d, 0x22, 0x0e, 0x2e, 0x82,
	0x6c, 0x6a, 0xab, 0xa5, 0x0b, 0x84, 0xc4, 0xa2, 0x21, 0x2a, 0x8a, 0x04, 0x05, 0x4c, 0xd5, 0x05,
	0x12, 0xb2, 0xc6, 0x9e, 0x49, 0x3a, 0xaa, 0xed, 0xb1, 0xec, 0x09, 0xa5, 0x1b, 0xde, 0x00, 0xa9,
	0x0b, 0x16, 0x6c, 0x79, 0x02, 0x5e, 0xa3, 0xcb, 0x2e, 0x11, 0x12, 0x06, 0xb5, 0x6f, 0x90, 0x25,
	0x2b, 0x34, 0x3f, 0x89, 0x93, 0xb4, 0x52, 0x1b, 0xb1, 0xca, 0xcc, 0xbd, 0x77, 0xce, 0x39, 0x73,
	0xe7, 0xe4, 0x1a, 0xec, 0xfb, 0x34, 0x44, 0xf8, 0x47, 0x1b, 0x8e, 0x7c, 0x46, 0x68, 0x94, 0x5e,
	0x1c, 0xd9, 0x3f, 0x1c, 0x7a, 0x98, 0xc1, 0x43, 0xdb, 0x23, 0xc8, 0x8a, 0x13, 0xca, 0xa8, 0xbe,
	0x27, 0x8b, 0xac, 0xbc, 0xc8, 0x52, 0x45, 0xcd, 0xad, 0x21, 0x1d, 0x52, 0x51, 0x65, 0xf3, 0x95,
	0x3c, 0xd0, 0x34, 0x86, 0x94, 0x0e, 0x03, 0x6c, 0x8b, 0x9d, 0x37, 0x1a, 0xd8, 0x8c, 0x84, 0x38,
	0x65, 0x30, 0x8c, 0x55, 0x41, 0xcb, 0xa7, 0x69, 0x48, 0x53, 0xdb, 0x83, 0x29, 0x9e, 0x12, 0xfa,
	0x94, 0x44, 0x32, 0x6f, 0xfe, 0x5e, 0x06, 0xc5, 0x2e, 0x41, 0xfa, 0x31, 0x00, 0x1e, 0x41, 0x88,
	0x44, 0x43, 0x97, 0xa0, 0x86, 0xd6, 0xd6, 0x3a, 0xa5, 0xee, 0xf6, 0x38, 0x33, 0xea, 0x37, 0x30,
	0x0c, 0x3e, 0x31, 0xf3, 0x9c, 0xe9, 0x54, 0xd5, 0xa6, 0x2f, 0x4e, 0x29, 0xa9, 0xfc, 0xd4, 0x9b,
	0xc5, 0x53, 0x79, 0xce, 0x74, 0xaa, 0x6a, 0xd3, 0x47, 0xfa, 0x6f, 0x1a, 0xd8, 0xf5, 0x69, 0x10,
	0x40, 0x86, 0x13, 0x18, 0xb8, 0x8c, 0x5e, 0xe1, 0xc8, 0x85, 0x21, 0x1d, 0x45, 0xac, 0x51, 0x6c,
	0x6b, 0x9d, 0xd5, 0xa3, 0x3d, 0x4b, 0xca, 0xb6, 0xb8, 0xec, 0x49, 0x0b, 0xac, 0xcf, 0x28, 0x89,
	0xba, 0x67, 0x77, 0x99, 0x51, 0x18, 0x67, 0xc6, 0x3b, 0x92, 0x82, 0x8e, 0xd8, 0x20, 0xa0, 0xd7,
	0x73, 0x20, 0xe6, 0xbf, 0x99, 0xf1, 0xe1, 0x90, 0xb0, 0xcb, 0x91, 0x67, 0xf9, 0x34, 0xb4, 0x55,
	0x0b, 0xe4, 0xcf, 0x41, 0x8a, 0xae, 0x6c, 0x76, 0x13, 0xe3, 0x54, 0xe0, 0x39, 0xdb, 0xb9, 0x92,
	0x73, 0x8e, 0x71, 0x22, 0x20, 0xf4, 0x5f, 0x34, 0x50, 0x47, 0xd8, 0x63, 0xf3, 0xea, 0x4a, 0x2f,
	0xa9, 0xfb, 0x52, 0xa9, 0x6b, 0x4a, 0x75, 0x24, 0xfa, 0x7f, 0xe2, 0xd6, 0xb9, 0x84, 0x59, 0x59,
	0x1f, 0x83, 0x35, 0xde, 0x7d, 0x9c, 0xb8, 0x10, 0xa1, 0x04, 0xa7, 0x69, 0xa3, 0xdc, 0xd6, 0x3a,
	0xd5, 0x6e, 0x7d, 0x9c, 0x19, 0xb5, 0xfc, 0xa9, 0x70, 0x62, 0x3a, 0x35, 0xb9, 0x38, 0x91, 0x75,
	0x7a, 0x08, 0xea, 0x93, 0x47, 0x9c, 0x7a, 0xa4, 0x51, 0x11, 0xf7, 0x69, 0x5a, 0xd2, 0x45, 0xd6,
	0xc4, 0x45, 0xd6, 0xf9, 0xa4, 0xa2, 0xfb, 0xbe, 0xba, 0x50, 0x63, 0xde, 0x07, 0x53, 0x08, 0xf3,
	0xf6, 0x6f, 0x43, 0x73, 0x36, 0x54, 0x7c, 0x7a, 0x4e, 0xef, 0x80, 0x0a, 0x8c, 0x63, 0xee, 0x8a,
	0xb7, 0xc2, 0x15, 0x33, 0x02, 0x65, 0xdc, 0x74, 0xca, 0x30, 0x8e, 0xfb, 0x48, 0xb7, 0xc0, 0x8a,
	0x47, 0x90, 0xcb, 0x6f, 0xdd, 0x58, 0x11, 0x97, 0xd9, 0x1c, 0x67, 0xc6, 0xfa, 0x94, 0x4f, 0x64,
	0x4c, 0xe7, 0xad, 0x47, 0xd0, 0x39, 0x5f, 0xfd, 0x55, 0x02, 0xb5, 0x2f, 0x48, 0x48, 0xd8, 0x57,
	0x09, 0xc2, 0x09, 0xf7, 0xee, 0x05, 0xd8, 0x09, 0x78, 0xc0, 0xa5, 0x3c, 0xe2, 0x3e, 0xf1, 0xf1,
	0x7b, 0xe3, 0xcc, 0x78, 0x57, 0xe2, 0x3d, 0x5f, 0x67, 0x3a, 0x9b, 0xc1, 0x2c, 0xa2, 0x72, 0xf7,
	0xd3, 0x66, 0xbf, 0x79, 0x65, 0xb3, 0x7f, 0xd6, 0xc0, 0xc6, 0xa2, 0xc3, 0x5f, 0xb6, 0xf6, 0xe7,
	0xaa, 0xd7, 0x5b, 0xcf, 0x58, 0x7b, 0x39, 0xdb, 0x2c, 0x78, 0x5a, 0xff, 0x09, 0x80, 0xdc, 0xcc,
	0x2f, 0xbb, 0xb8, 0xa7, 0x84, 0xa8, 0xbf, 0x71, 0x7e, 0x74, 0x29, 0x15, 0xd5, 0xa9, 0x79, 0x17,
	0xa6, 0x4b, 0xb9, 0x5d, 0x7c, 0xd5, 0x74, 0x61, 0x60, 0x23, 0x4e, 0x70, 0x48, 0x46, 0xa1, 0x8b,
	0x48, 0xea, 0x8b, 0x7f, 0x60, 0x45, 0xbc, 0x40, 0x9f, 0x0b, 0xfc, 0x33, 0x33, 0x3e, 0x78, 0x85,
	0x96, 0x1e, 0xf6, 0xc7, 0x99, 0xb1, 0x2b, 0x99, 0x16, 0xf1, 0x4c, 0x67, 0x5d, 0x85, 0x7a, 0x93,
	0xc8, 0xaf, 0x45, 0x50, 0x3b, 0x91, 0xb3, 0xea, 0x6b, 0x98, 0xc0, 0x30, 0xd5, 0xbf, 0x07, 0x8d,
	0xc9, 0x24, 0x43, 0xa3, 0x04, 0x8a, 0x45, 0x8a, 0x7d, 0x1a, 0xa1, 0x54, 0x39, 0x6c, 0x7f, 0x9c,
	0x19, 0xc6, 0xfc, 0xcc, 0x5b, 0xac, 0x34, 0x9d, 0x1d, 0x95, 0xea, 0xa9, 0xcc, 0xb7, 0x32, 0xa1,
	0x7f, 0x03, 0x4a, 0x29, 0xc3, 0xb1, 0x32, 0xd7, 0xa7, 0x4b, 0x5f, 0x6d, 0x55, 0x12, 0x73, 0x0c,
	0xd3, 0x11, 0x50, 0x7a, 0x04, 0xd6, 0xae, 0x09, 0xbb, 0x44, 0x09, 0xbc, 0x86, 0x81, 0x3b, 0xc0,
	0x58, 0x98, 0xaf, 0x2a, 0x1d, 0xb6, 0x14, 0xf8, 0xb6, 0x04, 0x9f, 0x47, 0x33, 0x9d, 0x5a, 0x1e,
	0x38, 0xc5, 0x58, 0xc7, 0x60, 0xd5, 0x0f, 0x68, 0xca, 0xdf, 0x90, 0x93, 0x95, 0x04, 0x59, 0x6f,
	0x69, 0x32, 0x5d, 0x92, 0xcd, 0x40, 0x99, 0x0e, 0x50, 0xbb, 0x53, 0x8c, 0xbb, 0x67, 0x77, 0x0f,
	0x2d, 0xed, 0xfe, 0xa1, 0xa5, 0xfd, 0xf3, 0xd0, 0xd2, 0x6e, 0x1f, 0x5b, 0x85, 0xfb, 0xc7, 0x56,
	0xe1, 0x8f, 0xc7, 0x56, 0xe1, 0xbb, 0xe3, 0x39, 0x0e, 0xfe, 0x0d, 0x3d, 0xa0, 0x83, 0x01, 0xf1,
	0x09, 0x0c, 0xd4, 0xde, 0x9e, 0xfb, 0xf4, 0x0a, 0x56, 0xaf, 0x22, 0x06, 0xde, 0x47, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x73, 0x38, 0xf1, 0x86, 0x9c, 0x07, 0x00, 0x00,
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidType) > 0 {
		i -= len(m.BidType)
		copy(dAtA[i:], m.BidType)
		i = encodeVarintBid(dAtA, i, uint64(len(m.BidType)))
		i--
		dAtA[i] = 0x42
	}
	if m.AppId != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x38
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BiddingTimestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBid(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if len(m.BidderAddress) > 0 {
		i -= len(m.BidderAddress)
		copy(dAtA[i:], m.BidderAddress)
		i = encodeVarintBid(dAtA, i, uint64(len(m.BidderAddress)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.DebtTokenAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.CollateralTokenAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AuctionId != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if m.BiddingId != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.BiddingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LimitOrderBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PremiumDiscount.Size()
		i -= size
		if _, err := m.PremiumDiscount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.BiddingId) > 0 {
		dAtA5 := make([]byte, len(m.BiddingId)*10)
		var j4 int
		for _, num := range m.BiddingId {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintBid(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.DebtToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.CollateralToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.BidderAddress) > 0 {
		i -= len(m.BidderAddress)
		copy(dAtA[i:], m.BidderAddress)
		i = encodeVarintBid(dAtA, i, uint64(len(m.BidderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.LimitOrderBiddingId != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.LimitOrderBiddingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClosingFee.Size()
		i -= size
		if _, err := m.ClosingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.WithdrawalFee.Size()
		i -= size
		if _, err := m.WithdrawalFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Step.Size()
		i -= size
		if _, err := m.Step.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AuctionDurationSeconds != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.AuctionDurationSeconds))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBid(dAtA []byte, offset int, v uint64) int {
	offset -= sovBid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BiddingId != 0 {
		n += 1 + sovBid(uint64(m.BiddingId))
	}
	if m.AuctionId != 0 {
		n += 1 + sovBid(uint64(m.AuctionId))
	}
	l = m.CollateralTokenAmount.Size()
	n += 1 + l + sovBid(uint64(l))
	l = m.DebtTokenAmount.Size()
	n += 1 + l + sovBid(uint64(l))
	l = len(m.BidderAddress)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BiddingTimestamp)
	n += 1 + l + sovBid(uint64(l))
	if m.AppId != 0 {
		n += 1 + sovBid(uint64(m.AppId))
	}
	l = len(m.BidType)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	return n
}

func (m *LimitOrderBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LimitOrderBiddingId != 0 {
		n += 1 + sovBid(uint64(m.LimitOrderBiddingId))
	}
	l = len(m.BidderAddress)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	l = m.CollateralToken.Size()
	n += 1 + l + sovBid(uint64(l))
	l = m.DebtToken.Size()
	n += 1 + l + sovBid(uint64(l))
	if len(m.BiddingId) > 0 {
		l = 0
		for _, e := range m.BiddingId {
			l += sovBid(uint64(e))
		}
		n += 1 + sovBid(uint64(l)) + l
	}
	l = m.PremiumDiscount.Size()
	n += 1 + l + sovBid(uint64(l))
	return n
}

func (m *AuctionParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionDurationSeconds != 0 {
		n += 1 + sovBid(uint64(m.AuctionDurationSeconds))
	}
	l = m.Step.Size()
	n += 1 + l + sovBid(uint64(l))
	l = m.WithdrawalFee.Size()
	n += 1 + l + sovBid(uint64(l))
	l = m.ClosingFee.Size()
	n += 1 + l + sovBid(uint64(l))
	return n
}

func sovBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBid(x uint64) (n int) {
	return sovBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingId", wireType)
			}
			m.BiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralTokenAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralTokenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtTokenAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtTokenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BiddingTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBid
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
func (m *LimitOrderBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
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
			return fmt.Errorf("proto: LimitOrderBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderBiddingId", wireType)
			}
			m.LimitOrderBiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitOrderBiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBid
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BiddingId = append(m.BiddingId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBid
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
					return ErrInvalidLengthBid
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBid
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BiddingId) == 0 {
					m.BiddingId = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBid
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BiddingId = append(m.BiddingId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingId", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PremiumDiscount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PremiumDiscount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBid
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
func (m *AuctionParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
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
			return fmt.Errorf("proto: AuctionParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDurationSeconds", wireType)
			}
			m.AuctionDurationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionDurationSeconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Step.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WithdrawalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBid
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
func skipBid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
				return 0, ErrInvalidLengthBid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBid = fmt.Errorf("proto: unexpected end of group")
)
