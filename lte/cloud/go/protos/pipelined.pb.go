// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/pipelined.proto

package protos // import "magma/lte/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protos "magma/orc8r/cloud/go/protos"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RuleModResult_Result int32

const (
	RuleModResult_SUCCESS         RuleModResult_Result = 0
	RuleModResult_PARTIAL_SUCCESS RuleModResult_Result = 1
	RuleModResult_FAILURE         RuleModResult_Result = 2
)

var RuleModResult_Result_name = map[int32]string{
	0: "SUCCESS",
	1: "PARTIAL_SUCCESS",
	2: "FAILURE",
}
var RuleModResult_Result_value = map[string]int32{
	"SUCCESS":         0,
	"PARTIAL_SUCCESS": 1,
	"FAILURE":         2,
}

func (x RuleModResult_Result) String() string {
	return proto.EnumName(RuleModResult_Result_name, int32(x))
}
func (RuleModResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{2, 0}
}

type ActivateFlowsRequest struct {
	Sid *SubscriberID `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	// Subscriber session ipv4 address
	IpAddr string `protobuf:"bytes,2,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	// List of static rules obtained from PCRF
	RuleIds []string `protobuf:"bytes,3,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids,omitempty"`
	// List of dynamic rules obtained from PCRF
	DynamicRules         []*PolicyRule `protobuf:"bytes,4,rep,name=dynamic_rules,json=dynamicRules,proto3" json:"dynamic_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ActivateFlowsRequest) Reset()         { *m = ActivateFlowsRequest{} }
func (m *ActivateFlowsRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateFlowsRequest) ProtoMessage()    {}
func (*ActivateFlowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{0}
}
func (m *ActivateFlowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateFlowsRequest.Unmarshal(m, b)
}
func (m *ActivateFlowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateFlowsRequest.Marshal(b, m, deterministic)
}
func (dst *ActivateFlowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateFlowsRequest.Merge(dst, src)
}
func (m *ActivateFlowsRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateFlowsRequest.Size(m)
}
func (m *ActivateFlowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateFlowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateFlowsRequest proto.InternalMessageInfo

func (m *ActivateFlowsRequest) GetSid() *SubscriberID {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *ActivateFlowsRequest) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

func (m *ActivateFlowsRequest) GetRuleIds() []string {
	if m != nil {
		return m.RuleIds
	}
	return nil
}

func (m *ActivateFlowsRequest) GetDynamicRules() []*PolicyRule {
	if m != nil {
		return m.DynamicRules
	}
	return nil
}

// DeactivateFlowsRequest can be used to deactivate all flows for a subscriber,
// all flows for some rules, or particular rules for a subscriber, depending on
// which parameters are passed. Rule IDs can apply to static rules or dynamic
// rules
// If no rule ids are given, all flows are deactivated
type DeactivateFlowsRequest struct {
	Sid                  *SubscriberID `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	RuleIds              []string      `protobuf:"bytes,2,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeactivateFlowsRequest) Reset()         { *m = DeactivateFlowsRequest{} }
func (m *DeactivateFlowsRequest) String() string { return proto.CompactTextString(m) }
func (*DeactivateFlowsRequest) ProtoMessage()    {}
func (*DeactivateFlowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{1}
}
func (m *DeactivateFlowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateFlowsRequest.Unmarshal(m, b)
}
func (m *DeactivateFlowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateFlowsRequest.Marshal(b, m, deterministic)
}
func (dst *DeactivateFlowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateFlowsRequest.Merge(dst, src)
}
func (m *DeactivateFlowsRequest) XXX_Size() int {
	return xxx_messageInfo_DeactivateFlowsRequest.Size(m)
}
func (m *DeactivateFlowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateFlowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateFlowsRequest proto.InternalMessageInfo

func (m *DeactivateFlowsRequest) GetSid() *SubscriberID {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *DeactivateFlowsRequest) GetRuleIds() []string {
	if m != nil {
		return m.RuleIds
	}
	return nil
}

type RuleModResult struct {
	RuleId               string               `protobuf:"bytes,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	Result               RuleModResult_Result `protobuf:"varint,2,opt,name=result,proto3,enum=magma.lte.RuleModResult_Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RuleModResult) Reset()         { *m = RuleModResult{} }
func (m *RuleModResult) String() string { return proto.CompactTextString(m) }
func (*RuleModResult) ProtoMessage()    {}
func (*RuleModResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{2}
}
func (m *RuleModResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleModResult.Unmarshal(m, b)
}
func (m *RuleModResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleModResult.Marshal(b, m, deterministic)
}
func (dst *RuleModResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleModResult.Merge(dst, src)
}
func (m *RuleModResult) XXX_Size() int {
	return xxx_messageInfo_RuleModResult.Size(m)
}
func (m *RuleModResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleModResult.DiscardUnknown(m)
}

var xxx_messageInfo_RuleModResult proto.InternalMessageInfo

func (m *RuleModResult) GetRuleId() string {
	if m != nil {
		return m.RuleId
	}
	return ""
}

func (m *RuleModResult) GetResult() RuleModResult_Result {
	if m != nil {
		return m.Result
	}
	return RuleModResult_SUCCESS
}

type ActivateFlowsResult struct {
	StaticRuleResults    []*RuleModResult `protobuf:"bytes,1,rep,name=static_rule_results,json=staticRuleResults,proto3" json:"static_rule_results,omitempty"`
	DynamicRuleResults   []*RuleModResult `protobuf:"bytes,2,rep,name=dynamic_rule_results,json=dynamicRuleResults,proto3" json:"dynamic_rule_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ActivateFlowsResult) Reset()         { *m = ActivateFlowsResult{} }
func (m *ActivateFlowsResult) String() string { return proto.CompactTextString(m) }
func (*ActivateFlowsResult) ProtoMessage()    {}
func (*ActivateFlowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{3}
}
func (m *ActivateFlowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateFlowsResult.Unmarshal(m, b)
}
func (m *ActivateFlowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateFlowsResult.Marshal(b, m, deterministic)
}
func (dst *ActivateFlowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateFlowsResult.Merge(dst, src)
}
func (m *ActivateFlowsResult) XXX_Size() int {
	return xxx_messageInfo_ActivateFlowsResult.Size(m)
}
func (m *ActivateFlowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateFlowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateFlowsResult proto.InternalMessageInfo

func (m *ActivateFlowsResult) GetStaticRuleResults() []*RuleModResult {
	if m != nil {
		return m.StaticRuleResults
	}
	return nil
}

func (m *ActivateFlowsResult) GetDynamicRuleResults() []*RuleModResult {
	if m != nil {
		return m.DynamicRuleResults
	}
	return nil
}

type DeactivateFlowsResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeactivateFlowsResult) Reset()         { *m = DeactivateFlowsResult{} }
func (m *DeactivateFlowsResult) String() string { return proto.CompactTextString(m) }
func (*DeactivateFlowsResult) ProtoMessage()    {}
func (*DeactivateFlowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{4}
}
func (m *DeactivateFlowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateFlowsResult.Unmarshal(m, b)
}
func (m *DeactivateFlowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateFlowsResult.Marshal(b, m, deterministic)
}
func (dst *DeactivateFlowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateFlowsResult.Merge(dst, src)
}
func (m *DeactivateFlowsResult) XXX_Size() int {
	return xxx_messageInfo_DeactivateFlowsResult.Size(m)
}
func (m *DeactivateFlowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateFlowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateFlowsResult proto.InternalMessageInfo

type FlowRequest struct {
	Match                *FlowMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	AppName              string     `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	BytesRx              uint64     `protobuf:"varint,3,opt,name=bytes_rx,json=bytesRx,proto3" json:"bytes_rx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FlowRequest) Reset()         { *m = FlowRequest{} }
func (m *FlowRequest) String() string { return proto.CompactTextString(m) }
func (*FlowRequest) ProtoMessage()    {}
func (*FlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{5}
}
func (m *FlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRequest.Unmarshal(m, b)
}
func (m *FlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRequest.Marshal(b, m, deterministic)
}
func (dst *FlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRequest.Merge(dst, src)
}
func (m *FlowRequest) XXX_Size() int {
	return xxx_messageInfo_FlowRequest.Size(m)
}
func (m *FlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRequest proto.InternalMessageInfo

func (m *FlowRequest) GetMatch() *FlowMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *FlowRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *FlowRequest) GetBytesRx() uint64 {
	if m != nil {
		return m.BytesRx
	}
	return 0
}

type FlowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowResponse) Reset()         { *m = FlowResponse{} }
func (m *FlowResponse) String() string { return proto.CompactTextString(m) }
func (*FlowResponse) ProtoMessage()    {}
func (*FlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{6}
}
func (m *FlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowResponse.Unmarshal(m, b)
}
func (m *FlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowResponse.Marshal(b, m, deterministic)
}
func (dst *FlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowResponse.Merge(dst, src)
}
func (m *FlowResponse) XXX_Size() int {
	return xxx_messageInfo_FlowResponse.Size(m)
}
func (m *FlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlowResponse proto.InternalMessageInfo

// UEMacFlowRequest is used to link a subscriber ID to a MAC address.
// This is used for Carrier WiFi data session establishment
type UEMacFlowRequest struct {
	Sid *SubscriberID `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	// UE MAC address
	MacAddr              string   `protobuf:"bytes,2,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UEMacFlowRequest) Reset()         { *m = UEMacFlowRequest{} }
func (m *UEMacFlowRequest) String() string { return proto.CompactTextString(m) }
func (*UEMacFlowRequest) ProtoMessage()    {}
func (*UEMacFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{7}
}
func (m *UEMacFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UEMacFlowRequest.Unmarshal(m, b)
}
func (m *UEMacFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UEMacFlowRequest.Marshal(b, m, deterministic)
}
func (dst *UEMacFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UEMacFlowRequest.Merge(dst, src)
}
func (m *UEMacFlowRequest) XXX_Size() int {
	return xxx_messageInfo_UEMacFlowRequest.Size(m)
}
func (m *UEMacFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UEMacFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UEMacFlowRequest proto.InternalMessageInfo

func (m *UEMacFlowRequest) GetSid() *SubscriberID {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *UEMacFlowRequest) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

type TableAssignment struct {
	AppName              string   `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	MainTable            uint64   `protobuf:"varint,2,opt,name=main_table,json=mainTable,proto3" json:"main_table,omitempty"`
	ScratchTables        []uint64 `protobuf:"varint,3,rep,packed,name=scratch_tables,json=scratchTables,proto3" json:"scratch_tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableAssignment) Reset()         { *m = TableAssignment{} }
func (m *TableAssignment) String() string { return proto.CompactTextString(m) }
func (*TableAssignment) ProtoMessage()    {}
func (*TableAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{8}
}
func (m *TableAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableAssignment.Unmarshal(m, b)
}
func (m *TableAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableAssignment.Marshal(b, m, deterministic)
}
func (dst *TableAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableAssignment.Merge(dst, src)
}
func (m *TableAssignment) XXX_Size() int {
	return xxx_messageInfo_TableAssignment.Size(m)
}
func (m *TableAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_TableAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_TableAssignment proto.InternalMessageInfo

func (m *TableAssignment) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *TableAssignment) GetMainTable() uint64 {
	if m != nil {
		return m.MainTable
	}
	return 0
}

func (m *TableAssignment) GetScratchTables() []uint64 {
	if m != nil {
		return m.ScratchTables
	}
	return nil
}

type AllTableAssignments struct {
	TableAssignments     []*TableAssignment `protobuf:"bytes,1,rep,name=table_assignments,json=tableAssignments,proto3" json:"table_assignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AllTableAssignments) Reset()         { *m = AllTableAssignments{} }
func (m *AllTableAssignments) String() string { return proto.CompactTextString(m) }
func (*AllTableAssignments) ProtoMessage()    {}
func (*AllTableAssignments) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipelined_d7288e4fcedc4f14, []int{9}
}
func (m *AllTableAssignments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTableAssignments.Unmarshal(m, b)
}
func (m *AllTableAssignments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTableAssignments.Marshal(b, m, deterministic)
}
func (dst *AllTableAssignments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTableAssignments.Merge(dst, src)
}
func (m *AllTableAssignments) XXX_Size() int {
	return xxx_messageInfo_AllTableAssignments.Size(m)
}
func (m *AllTableAssignments) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTableAssignments.DiscardUnknown(m)
}

var xxx_messageInfo_AllTableAssignments proto.InternalMessageInfo

func (m *AllTableAssignments) GetTableAssignments() []*TableAssignment {
	if m != nil {
		return m.TableAssignments
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivateFlowsRequest)(nil), "magma.lte.ActivateFlowsRequest")
	proto.RegisterType((*DeactivateFlowsRequest)(nil), "magma.lte.DeactivateFlowsRequest")
	proto.RegisterType((*RuleModResult)(nil), "magma.lte.RuleModResult")
	proto.RegisterType((*ActivateFlowsResult)(nil), "magma.lte.ActivateFlowsResult")
	proto.RegisterType((*DeactivateFlowsResult)(nil), "magma.lte.DeactivateFlowsResult")
	proto.RegisterType((*FlowRequest)(nil), "magma.lte.FlowRequest")
	proto.RegisterType((*FlowResponse)(nil), "magma.lte.FlowResponse")
	proto.RegisterType((*UEMacFlowRequest)(nil), "magma.lte.UEMacFlowRequest")
	proto.RegisterType((*TableAssignment)(nil), "magma.lte.TableAssignment")
	proto.RegisterType((*AllTableAssignments)(nil), "magma.lte.AllTableAssignments")
	proto.RegisterEnum("magma.lte.RuleModResult_Result", RuleModResult_Result_name, RuleModResult_Result_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelinedClient is the client API for Pipelined service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelinedClient interface {
	// Get subscriber metering flows
	GetSubscriberMeteringFlows(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*FlowTable, error)
	// Activate flows for a subscriber based on predefined flow templates
	ActivateFlows(ctx context.Context, in *ActivateFlowsRequest, opts ...grpc.CallOption) (*ActivateFlowsResult, error)
	// Deactivate flows for a subscriber
	DeactivateFlows(ctx context.Context, in *DeactivateFlowsRequest, opts ...grpc.CallOption) (*DeactivateFlowsResult, error)
	// Get policy usage stats
	GetPolicyUsage(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*RuleRecordTable, error)
	// Add new dpi flow
	CreateFlow(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*FlowResponse, error)
	// Update flow stats
	UpdateFlowStats(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*FlowResponse, error)
	// Add a flow for a subscriber by matching the provided UE MAC address
	AddUEMacFlow(ctx context.Context, in *UEMacFlowRequest, opts ...grpc.CallOption) (*FlowResponse, error)
	// Get the flow table assignment for all apps ordered by main table number
	// and name
	GetAllTableAssignments(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllTableAssignments, error)
}

type pipelinedClient struct {
	cc *grpc.ClientConn
}

func NewPipelinedClient(cc *grpc.ClientConn) PipelinedClient {
	return &pipelinedClient{cc}
}

func (c *pipelinedClient) GetSubscriberMeteringFlows(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*FlowTable, error) {
	out := new(FlowTable)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/GetSubscriberMeteringFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) ActivateFlows(ctx context.Context, in *ActivateFlowsRequest, opts ...grpc.CallOption) (*ActivateFlowsResult, error) {
	out := new(ActivateFlowsResult)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/ActivateFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) DeactivateFlows(ctx context.Context, in *DeactivateFlowsRequest, opts ...grpc.CallOption) (*DeactivateFlowsResult, error) {
	out := new(DeactivateFlowsResult)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/DeactivateFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) GetPolicyUsage(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*RuleRecordTable, error) {
	out := new(RuleRecordTable)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/GetPolicyUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) CreateFlow(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*FlowResponse, error) {
	out := new(FlowResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/CreateFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) UpdateFlowStats(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*FlowResponse, error) {
	out := new(FlowResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/UpdateFlowStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) AddUEMacFlow(ctx context.Context, in *UEMacFlowRequest, opts ...grpc.CallOption) (*FlowResponse, error) {
	out := new(FlowResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/AddUEMacFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinedClient) GetAllTableAssignments(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllTableAssignments, error) {
	out := new(AllTableAssignments)
	err := c.cc.Invoke(ctx, "/magma.lte.Pipelined/GetAllTableAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelinedServer is the server API for Pipelined service.
type PipelinedServer interface {
	// Get subscriber metering flows
	GetSubscriberMeteringFlows(context.Context, *protos.Void) (*FlowTable, error)
	// Activate flows for a subscriber based on predefined flow templates
	ActivateFlows(context.Context, *ActivateFlowsRequest) (*ActivateFlowsResult, error)
	// Deactivate flows for a subscriber
	DeactivateFlows(context.Context, *DeactivateFlowsRequest) (*DeactivateFlowsResult, error)
	// Get policy usage stats
	GetPolicyUsage(context.Context, *protos.Void) (*RuleRecordTable, error)
	// Add new dpi flow
	CreateFlow(context.Context, *FlowRequest) (*FlowResponse, error)
	// Update flow stats
	UpdateFlowStats(context.Context, *FlowRequest) (*FlowResponse, error)
	// Add a flow for a subscriber by matching the provided UE MAC address
	AddUEMacFlow(context.Context, *UEMacFlowRequest) (*FlowResponse, error)
	// Get the flow table assignment for all apps ordered by main table number
	// and name
	GetAllTableAssignments(context.Context, *protos.Void) (*AllTableAssignments, error)
}

func RegisterPipelinedServer(s *grpc.Server, srv PipelinedServer) {
	s.RegisterService(&_Pipelined_serviceDesc, srv)
}

func _Pipelined_GetSubscriberMeteringFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).GetSubscriberMeteringFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/GetSubscriberMeteringFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).GetSubscriberMeteringFlows(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_ActivateFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateFlowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).ActivateFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/ActivateFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).ActivateFlows(ctx, req.(*ActivateFlowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_DeactivateFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateFlowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).DeactivateFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/DeactivateFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).DeactivateFlows(ctx, req.(*DeactivateFlowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_GetPolicyUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).GetPolicyUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/GetPolicyUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).GetPolicyUsage(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_CreateFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).CreateFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/CreateFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).CreateFlow(ctx, req.(*FlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_UpdateFlowStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).UpdateFlowStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/UpdateFlowStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).UpdateFlowStats(ctx, req.(*FlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_AddUEMacFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEMacFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).AddUEMacFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/AddUEMacFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).AddUEMacFlow(ctx, req.(*UEMacFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelined_GetAllTableAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinedServer).GetAllTableAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Pipelined/GetAllTableAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinedServer).GetAllTableAssignments(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pipelined_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.Pipelined",
	HandlerType: (*PipelinedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriberMeteringFlows",
			Handler:    _Pipelined_GetSubscriberMeteringFlows_Handler,
		},
		{
			MethodName: "ActivateFlows",
			Handler:    _Pipelined_ActivateFlows_Handler,
		},
		{
			MethodName: "DeactivateFlows",
			Handler:    _Pipelined_DeactivateFlows_Handler,
		},
		{
			MethodName: "GetPolicyUsage",
			Handler:    _Pipelined_GetPolicyUsage_Handler,
		},
		{
			MethodName: "CreateFlow",
			Handler:    _Pipelined_CreateFlow_Handler,
		},
		{
			MethodName: "UpdateFlowStats",
			Handler:    _Pipelined_UpdateFlowStats_Handler,
		},
		{
			MethodName: "AddUEMacFlow",
			Handler:    _Pipelined_AddUEMacFlow_Handler,
		},
		{
			MethodName: "GetAllTableAssignments",
			Handler:    _Pipelined_GetAllTableAssignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/pipelined.proto",
}

func init() {
	proto.RegisterFile("lte/protos/pipelined.proto", fileDescriptor_pipelined_d7288e4fcedc4f14)
}

var fileDescriptor_pipelined_d7288e4fcedc4f14 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x8e, 0x23, 0x35,
	0x14, 0xed, 0xea, 0x84, 0x84, 0xdc, 0xee, 0x3c, 0xda, 0xe9, 0xe9, 0x24, 0x35, 0x1a, 0x08, 0x25,
	0x21, 0x05, 0x16, 0x89, 0x14, 0x16, 0x83, 0xd8, 0xa0, 0xa2, 0x1f, 0x21, 0x30, 0x41, 0x2d, 0x67,
	0x82, 0x46, 0x2c, 0x28, 0x39, 0x65, 0x2b, 0x94, 0x54, 0x2f, 0x6c, 0x07, 0xa6, 0x7f, 0x85, 0x5f,
	0x60, 0xc5, 0x27, 0xf1, 0x27, 0xc8, 0x2e, 0x27, 0xe3, 0x54, 0x67, 0x06, 0xc1, 0xac, 0x4a, 0xbe,
	0x3e, 0xf7, 0xdc, 0xd7, 0xf1, 0x2d, 0x70, 0x63, 0xc9, 0x26, 0x39, 0xcf, 0x64, 0x26, 0x26, 0x79,
	0x94, 0xb3, 0x38, 0x4a, 0x19, 0x1d, 0x6b, 0x03, 0x6a, 0x24, 0x64, 0x93, 0x90, 0x71, 0x2c, 0x99,
	0x3b, 0xc8, 0x78, 0xf8, 0x25, 0xdf, 0x01, 0xc3, 0x2c, 0x49, 0xb2, 0xb4, 0x40, 0xb9, 0x36, 0x43,
	0xc2, 0x24, 0xe3, 0x51, 0xba, 0x31, 0x0c, 0xee, 0xc0, 0x66, 0xcf, 0xe2, 0x28, 0x7c, 0xa0, 0x6b,
	0x73, 0x35, 0xb4, 0xae, 0x04, 0x13, 0x22, 0xca, 0xd2, 0x20, 0x21, 0x29, 0xd9, 0x30, 0x6e, 0x10,
	0xcf, 0x6c, 0xc4, 0x76, 0x2d, 0x42, 0x1e, 0xad, 0x19, 0xdf, 0x11, 0x78, 0x7f, 0x39, 0x70, 0xe9,
	0x87, 0x32, 0xfa, 0x8d, 0x48, 0x76, 0x17, 0x67, 0xbf, 0x0b, 0xcc, 0x7e, 0xdd, 0x32, 0x21, 0xd1,
	0x67, 0x50, 0x11, 0x11, 0xed, 0x3b, 0x43, 0x67, 0x74, 0x36, 0xed, 0x8d, 0xf7, 0x45, 0x8c, 0x97,
	0x7b, 0x92, 0xf9, 0x0d, 0x56, 0x18, 0xd4, 0x83, 0x7a, 0x94, 0x07, 0x84, 0x52, 0xde, 0x3f, 0x1d,
	0x3a, 0xa3, 0x06, 0xae, 0x45, 0xb9, 0x4f, 0x29, 0x47, 0x03, 0xf8, 0x90, 0x6f, 0x63, 0x16, 0x44,
	0x54, 0xf4, 0x2b, 0xc3, 0xca, 0xa8, 0x81, 0xeb, 0xea, 0x3c, 0xa7, 0x02, 0x7d, 0x05, 0x4d, 0xfa,
	0x90, 0x92, 0x24, 0x0a, 0x03, 0x65, 0x12, 0xfd, 0xea, 0xb0, 0x32, 0x3a, 0x9b, 0x3e, 0xb1, 0x02,
	0xdd, 0xeb, 0x52, 0xf1, 0x36, 0x66, 0xf8, 0xdc, 0x60, 0xd5, 0x41, 0x78, 0x3f, 0xc3, 0xd5, 0x0d,
	0x23, 0xef, 0x99, 0xb4, 0x9d, 0xdb, 0xe9, 0x41, 0x6e, 0xde, 0x1f, 0x0e, 0x34, 0x55, 0xa4, 0x45,
	0x46, 0x31, 0x13, 0xdb, 0x58, 0xaa, 0x0a, 0x0d, 0x58, 0x73, 0x37, 0x70, 0xad, 0xc0, 0xa2, 0xe7,
	0x50, 0xe3, 0x1a, 0xa2, 0x2b, 0x6f, 0x4d, 0x3f, 0xb6, 0x62, 0x1e, 0x50, 0x8c, 0x8b, 0x0f, 0x36,
	0x70, 0xef, 0x39, 0xd4, 0x0c, 0xf7, 0x19, 0xd4, 0x97, 0xab, 0xeb, 0xeb, 0xdb, 0xe5, 0xb2, 0x73,
	0x82, 0xba, 0xd0, 0xbe, 0xf7, 0xf1, 0xcb, 0xb9, 0xff, 0x22, 0xd8, 0x19, 0x1d, 0x85, 0xb8, 0xf3,
	0xe7, 0x2f, 0x56, 0xf8, 0xb6, 0x73, 0xea, 0xfd, 0xe9, 0x40, 0xb7, 0x34, 0x30, 0x4d, 0xf3, 0x2d,
	0x74, 0x85, 0x24, 0xd2, 0xf4, 0x33, 0x28, 0xc2, 0x88, 0xbe, 0xa3, 0xdb, 0xda, 0x7f, 0x5b, 0x5a,
	0xf8, 0xa2, 0x70, 0xd2, 0x5d, 0x2e, 0x5c, 0xd0, 0x77, 0x70, 0x69, 0x8f, 0x66, 0x4f, 0x75, 0xfa,
	0x2f, 0x54, 0xc8, 0x1a, 0x92, 0xe1, 0xf2, 0x7a, 0xf0, 0xe4, 0xd1, 0xa8, 0x74, 0xfd, 0x19, 0x9c,
	0xa9, 0xe3, 0x6e, 0x70, 0x9f, 0xc3, 0x07, 0x09, 0x91, 0xe1, 0x2f, 0x66, 0x74, 0x97, 0x56, 0x10,
	0x05, 0x5b, 0xa8, 0x3b, 0x5c, 0x40, 0xd4, 0xe4, 0x48, 0x9e, 0x07, 0x29, 0x49, 0x98, 0xd1, 0x5b,
	0x9d, 0xe4, 0xf9, 0x0f, 0x24, 0x61, 0xea, 0x6a, 0xfd, 0x20, 0x99, 0x08, 0xf8, 0xeb, 0x7e, 0x65,
	0xe8, 0x8c, 0xaa, 0xb8, 0xae, 0xcf, 0xf8, 0xb5, 0xd7, 0x82, 0xf3, 0x22, 0xa0, 0xc8, 0xb3, 0x54,
	0x30, 0xef, 0x15, 0x74, 0x56, 0xb7, 0x0b, 0x12, 0xda, 0x59, 0xfc, 0x37, 0xf9, 0x24, 0x24, 0xb4,
	0x45, 0x5f, 0x4f, 0x48, 0xa8, 0x54, 0xef, 0x71, 0x68, 0xbf, 0x24, 0xeb, 0x98, 0xf9, 0x42, 0x44,
	0x9b, 0x34, 0x61, 0xa9, 0x3c, 0x48, 0xd9, 0x39, 0x4c, 0xf9, 0x19, 0x40, 0x42, 0xa2, 0x34, 0x90,
	0xca, 0x45, 0x53, 0x55, 0x71, 0x43, 0x59, 0x34, 0x07, 0xfa, 0x14, 0x5a, 0x22, 0xe4, 0xaa, 0xee,
	0x02, 0x51, 0x3c, 0xa4, 0x2a, 0x6e, 0x1a, 0xab, 0x46, 0xa9, 0x27, 0xd1, 0xf5, 0xe3, 0xb8, 0x14,
	0x56, 0xa0, 0x19, 0x5c, 0x68, 0xaf, 0x80, 0xbc, 0x31, 0x1a, 0x49, 0xb8, 0x56, 0x79, 0x25, 0x3f,
	0xdc, 0x91, 0x25, 0xa2, 0xe9, 0xdf, 0x55, 0x68, 0xdc, 0xef, 0x16, 0x1b, 0x9a, 0x81, 0x3b, 0x63,
	0xf2, 0x4d, 0x53, 0x16, 0x66, 0x61, 0xe9, 0x01, 0xa3, 0x0b, 0xc3, 0xac, 0x97, 0xdd, 0xf8, 0xc7,
	0x2c, 0xa2, 0x6e, 0x79, 0x9e, 0x3a, 0xa0, 0x77, 0x82, 0x30, 0x34, 0x0f, 0xb4, 0x8c, 0xec, 0xf7,
	0x73, 0x6c, 0x2d, 0xb9, 0x1f, 0xbd, 0x1d, 0xa0, 0x75, 0x75, 0x82, 0x5e, 0x41, 0xbb, 0x24, 0x39,
	0xf4, 0x89, 0xe5, 0x74, 0x7c, 0x73, 0xb8, 0xc3, 0x77, 0x41, 0x0c, 0xb3, 0x0f, 0xad, 0x19, 0x93,
	0xc5, 0x5a, 0x5a, 0x09, 0xb2, 0x61, 0xc7, 0x4a, 0x75, 0x4b, 0xef, 0x03, 0xb3, 0x30, 0xe3, 0x74,
	0x57, 0xf0, 0xd7, 0x00, 0xd7, 0x9c, 0x19, 0x66, 0x74, 0x55, 0x6a, 0xcb, 0x2e, 0x99, 0xde, 0x23,
	0xbb, 0x11, 0xed, 0x09, 0xba, 0x81, 0xf6, 0x2a, 0xa7, 0x86, 0x60, 0x29, 0x89, 0x14, 0xff, 0x87,
	0xe5, 0x0e, 0xce, 0x7d, 0x4a, 0xf7, 0xfa, 0x47, 0x4f, 0x2d, 0x68, 0xf9, 0x55, 0xbc, 0x8b, 0xe7,
	0x7b, 0xb8, 0x9a, 0x31, 0x79, 0x4c, 0x79, 0x47, 0x3a, 0x73, 0x30, 0xba, 0xc7, 0x2e, 0xdf, 0x3c,
	0xfd, 0x69, 0xa0, 0x01, 0x13, 0xf5, 0xc7, 0x0a, 0xe3, 0x6c, 0x4b, 0x27, 0x9b, 0xcc, 0xfc, 0xba,
	0xd6, 0x35, 0xfd, 0xfd, 0xe2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x87, 0xd3, 0xac, 0x6a,
	0x07, 0x00, 0x00,
}
