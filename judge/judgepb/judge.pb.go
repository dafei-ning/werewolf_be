// Code generated by protoc-gen-go. DO NOT EDIT.
// source: judge/judgepb/judge.proto

package judgepb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NewGameInfo_GamePattern int32

const (
	NewGameInfo_CHARMWOLF_GUARD    NewGameInfo_GamePattern = 0
	NewGameInfo_BLACKMARKETER_BROS NewGameInfo_GamePattern = 1
	NewGameInfo_BLOOD_WITCHER      NewGameInfo_GamePattern = 2
	NewGameInfo_CUPID_ROBBER       NewGameInfo_GamePattern = 3
	NewGameInfo_NIGHTMARE_GUARD    NewGameInfo_GamePattern = 4
	NewGameInfo_SHADOW_AVENGER     NewGameInfo_GamePattern = 5
	NewGameInfo_SWHI               NewGameInfo_GamePattern = 6
	NewGameInfo_SWGHI3             NewGameInfo_GamePattern = 7
)

var NewGameInfo_GamePattern_name = map[int32]string{
	0: "CHARMWOLF_GUARD",
	1: "BLACKMARKETER_BROS",
	2: "BLOOD_WITCHER",
	3: "CUPID_ROBBER",
	4: "NIGHTMARE_GUARD",
	5: "SHADOW_AVENGER",
	6: "SWHI",
	7: "SWGHI3",
}

var NewGameInfo_GamePattern_value = map[string]int32{
	"CHARMWOLF_GUARD":    0,
	"BLACKMARKETER_BROS": 1,
	"BLOOD_WITCHER":      2,
	"CUPID_ROBBER":       3,
	"NIGHTMARE_GUARD":    4,
	"SHADOW_AVENGER":     5,
	"SWHI":               6,
	"SWGHI3":             7,
}

func (x NewGameInfo_GamePattern) String() string {
	return proto.EnumName(NewGameInfo_GamePattern_name, int32(x))
}

func (NewGameInfo_GamePattern) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d265219a8a7b46e8, []int{0, 0}
}

type NewGameInfo struct {
	PlayerAmount         int32                   `protobuf:"varint,1,opt,name=playerAmount,proto3" json:"playerAmount,omitempty"`
	GamePattern          NewGameInfo_GamePattern `protobuf:"varint,2,opt,name=gamePattern,proto3,enum=judge.NewGameInfo_GamePattern" json:"gamePattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NewGameInfo) Reset()         { *m = NewGameInfo{} }
func (m *NewGameInfo) String() string { return proto.CompactTextString(m) }
func (*NewGameInfo) ProtoMessage()    {}
func (*NewGameInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d265219a8a7b46e8, []int{0}
}

func (m *NewGameInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewGameInfo.Unmarshal(m, b)
}
func (m *NewGameInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewGameInfo.Marshal(b, m, deterministic)
}
func (m *NewGameInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewGameInfo.Merge(m, src)
}
func (m *NewGameInfo) XXX_Size() int {
	return xxx_messageInfo_NewGameInfo.Size(m)
}
func (m *NewGameInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewGameInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NewGameInfo proto.InternalMessageInfo

func (m *NewGameInfo) GetPlayerAmount() int32 {
	if m != nil {
		return m.PlayerAmount
	}
	return 0
}

func (m *NewGameInfo) GetGamePattern() NewGameInfo_GamePattern {
	if m != nil {
		return m.GamePattern
	}
	return NewGameInfo_CHARMWOLF_GUARD
}

type StartGameRequest struct {
	NewGameInfo          *NewGameInfo `protobuf:"bytes,1,opt,name=new_game_info,json=newGameInfo,proto3" json:"new_game_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StartGameRequest) Reset()         { *m = StartGameRequest{} }
func (m *StartGameRequest) String() string { return proto.CompactTextString(m) }
func (*StartGameRequest) ProtoMessage()    {}
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d265219a8a7b46e8, []int{1}
}

func (m *StartGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameRequest.Unmarshal(m, b)
}
func (m *StartGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameRequest.Marshal(b, m, deterministic)
}
func (m *StartGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameRequest.Merge(m, src)
}
func (m *StartGameRequest) XXX_Size() int {
	return xxx_messageInfo_StartGameRequest.Size(m)
}
func (m *StartGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameRequest proto.InternalMessageInfo

func (m *StartGameRequest) GetNewGameInfo() *NewGameInfo {
	if m != nil {
		return m.NewGameInfo
	}
	return nil
}

type StartGameResponse struct {
	ResponseStartgame    string   `protobuf:"bytes,1,opt,name=response_startgame,json=responseStartgame,proto3" json:"response_startgame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartGameResponse) Reset()         { *m = StartGameResponse{} }
func (m *StartGameResponse) String() string { return proto.CompactTextString(m) }
func (*StartGameResponse) ProtoMessage()    {}
func (*StartGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d265219a8a7b46e8, []int{2}
}

func (m *StartGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameResponse.Unmarshal(m, b)
}
func (m *StartGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameResponse.Marshal(b, m, deterministic)
}
func (m *StartGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameResponse.Merge(m, src)
}
func (m *StartGameResponse) XXX_Size() int {
	return xxx_messageInfo_StartGameResponse.Size(m)
}
func (m *StartGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameResponse proto.InternalMessageInfo

func (m *StartGameResponse) GetResponseStartgame() string {
	if m != nil {
		return m.ResponseStartgame
	}
	return ""
}

func init() {
	proto.RegisterEnum("judge.NewGameInfo_GamePattern", NewGameInfo_GamePattern_name, NewGameInfo_GamePattern_value)
	proto.RegisterType((*NewGameInfo)(nil), "judge.NewGameInfo")
	proto.RegisterType((*StartGameRequest)(nil), "judge.StartGameRequest")
	proto.RegisterType((*StartGameResponse)(nil), "judge.StartGameResponse")
}

func init() { proto.RegisterFile("judge/judgepb/judge.proto", fileDescriptor_d265219a8a7b46e8) }

var fileDescriptor_d265219a8a7b46e8 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0x86, 0x6f, 0xf1, 0xc2, 0x95, 0x53, 0xc0, 0xe1, 0x98, 0x28, 0xba, 0x30, 0xa4, 0x2b, 0x36,
	0x62, 0x02, 0x89, 0x6b, 0xa6, 0x1f, 0xb6, 0x15, 0x68, 0xc9, 0x14, 0x6c, 0xe2, 0x66, 0x52, 0x74,
	0x20, 0x18, 0xfb, 0x61, 0x5b, 0x24, 0xfe, 0x0f, 0xd7, 0xfe, 0x56, 0xd3, 0x52, 0xa4, 0x5e, 0x36,
	0xed, 0x9c, 0x79, 0x66, 0x9e, 0xbc, 0x67, 0x72, 0xe0, 0xd5, 0xb7, 0xe3, 0xd7, 0xbd, 0x78, 0x57,
	0x7e, 0x93, 0xed, 0xf9, 0x3f, 0x4e, 0xd2, 0x38, 0x8f, 0xb1, 0x59, 0x16, 0xca, 0xef, 0x06, 0xc8,
	0x8e, 0x38, 0x99, 0x41, 0x28, 0xec, 0x68, 0x17, 0xa3, 0x02, 0x9d, 0xe4, 0x7b, 0xf0, 0x4b, 0xa4,
	0x34, 0x8c, 0x8f, 0x51, 0x3e, 0x90, 0x86, 0xd2, 0xa8, 0xc9, 0xfe, 0xdb, 0xc3, 0x19, 0xc8, 0xfb,
	0x20, 0x14, 0xab, 0x20, 0xcf, 0x45, 0x1a, 0x0d, 0x1a, 0x43, 0x69, 0xd4, 0x9b, 0xbc, 0x19, 0x9f,
	0xed, 0x35, 0xd9, 0xd8, 0xbc, 0x9e, 0x62, 0xf5, 0x2b, 0xca, 0x1f, 0x09, 0xe4, 0x1a, 0xc4, 0xe7,
	0xf0, 0x4c, 0xb3, 0x28, 0x5b, 0xfa, 0xee, 0xe2, 0x03, 0x37, 0x37, 0x94, 0xe9, 0xe4, 0x0e, 0x5f,
	0x00, 0xaa, 0x0b, 0xaa, 0xcd, 0x97, 0x94, 0xcd, 0x8d, 0xb5, 0xc1, 0xb8, 0xca, 0x5c, 0x8f, 0x48,
	0xd8, 0x87, 0xae, 0xba, 0x70, 0x5d, 0x9d, 0xfb, 0xf6, 0x5a, 0xb3, 0x0c, 0x46, 0x1a, 0x48, 0xa0,
	0xa3, 0x6d, 0x56, 0xb6, 0xce, 0x99, 0xab, 0xaa, 0x06, 0x23, 0x4f, 0x0a, 0xa3, 0x63, 0x9b, 0xd6,
	0x7a, 0x49, 0x99, 0x51, 0x19, 0xef, 0x11, 0xa1, 0xe7, 0x59, 0x54, 0x77, 0x7d, 0x4e, 0x3f, 0x19,
	0x8e, 0x69, 0x30, 0xd2, 0xc4, 0xa7, 0x70, 0xef, 0xf9, 0x96, 0x4d, 0x5a, 0x08, 0xd0, 0xf2, 0x7c,
	0xd3, 0xb2, 0xa7, 0xe4, 0x41, 0xf9, 0x08, 0xc4, 0xcb, 0x83, 0x34, 0x2f, 0x42, 0x32, 0xf1, 0xe3,
	0x28, 0xb2, 0x1c, 0xdf, 0x43, 0x37, 0x12, 0x27, 0x5e, 0xf4, 0xc1, 0x0f, 0xd1, 0x2e, 0x2e, 0xdf,
	0x46, 0x9e, 0xe0, 0x6d, 0xe3, 0x4c, 0x8e, 0xae, 0x85, 0xa2, 0x42, 0xbf, 0xe6, 0xca, 0x92, 0x38,
	0xca, 0x04, 0xbe, 0x05, 0x4c, 0xab, 0x35, 0xcf, 0x0a, 0x5a, 0x68, 0x4b, 0x63, 0x9b, 0xf5, 0x2f,
	0xc4, 0xbb, 0x80, 0x89, 0x0b, 0xb2, 0x27, 0xd2, 0x9f, 0x87, 0x2f, 0xc2, 0x09, 0x42, 0x81, 0x33,
	0x68, 0xff, 0x53, 0xe2, 0xcb, 0x2a, 0xc0, 0xe3, 0xc0, 0xaf, 0x07, 0xb7, 0xe0, 0xec, 0x55, 0xee,
	0xd4, 0xf6, 0xe7, 0x87, 0x6a, 0x2a, 0xb6, 0xad, 0x72, 0x20, 0xa6, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0xa6, 0x4a, 0xdc, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceNameClient is the client API for ServiceName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceNameClient interface {
	// Unary 开始游戏
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error)
}

type serviceNameClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceNameClient(cc grpc.ClientConnInterface) ServiceNameClient {
	return &serviceNameClient{cc}
}

func (c *serviceNameClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error) {
	out := new(StartGameResponse)
	err := c.cc.Invoke(ctx, "/judge.ServiceName/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceNameServer is the server API for ServiceName service.
type ServiceNameServer interface {
	// Unary 开始游戏
	StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error)
}

// UnimplementedServiceNameServer can be embedded to have forward compatible implementations.
type UnimplementedServiceNameServer struct {
}

func (*UnimplementedServiceNameServer) StartGame(ctx context.Context, req *StartGameRequest) (*StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}

func RegisterServiceNameServer(s *grpc.Server, srv ServiceNameServer) {
	s.RegisterService(&_ServiceName_serviceDesc, srv)
}

func _ServiceName_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.ServiceName/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceName_serviceDesc = grpc.ServiceDesc{
	ServiceName: "judge.ServiceName",
	HandlerType: (*ServiceNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartGame",
			Handler:    _ServiceName_StartGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge/judgepb/judge.proto",
}