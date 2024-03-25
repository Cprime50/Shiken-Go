// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: quiz.proto

package quiz

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuizService_GetQuiz_FullMethodName          = "/quizpb.QuizService/GetQuiz"
	QuizService_GetLeaderBoard_FullMethodName   = "/quizpb.QuizService/GetLeaderBoard"
	QuizService_GetScore_FullMethodName         = "/quizpb.QuizService/GetScore"
	QuizService_GetResult_FullMethodName        = "/quizpb.QuizService/GetResult"
	QuizService_CreateUpdateQuiz_FullMethodName = "/quizpb.QuizService/CreateUpdateQuiz"
	QuizService_DeleteQuiz_FullMethodName       = "/quizpb.QuizService/DeleteQuiz"
	QuizService_GetAllQuizzes_FullMethodName    = "/quizpb.QuizService/GetAllQuizzes"
)

// QuizServiceClient is the client API for QuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizServiceClient interface {
	GetQuiz(ctx context.Context, in *GetQuizRequest, opts ...grpc.CallOption) (*GetQuizResponse, error)
	GetLeaderBoard(ctx context.Context, in *Empty, opts ...grpc.CallOption) (QuizService_GetLeaderBoardClient, error)
	GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
	CreateUpdateQuiz(ctx context.Context, in *CreateUpdateQuizRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteQuiz(ctx context.Context, in *DeleteQuizRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAllQuizzes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (QuizService_GetAllQuizzesClient, error)
}

type quizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizServiceClient(cc grpc.ClientConnInterface) QuizServiceClient {
	return &quizServiceClient{cc}
}

func (c *quizServiceClient) GetQuiz(ctx context.Context, in *GetQuizRequest, opts ...grpc.CallOption) (*GetQuizResponse, error) {
	out := new(GetQuizResponse)
	err := c.cc.Invoke(ctx, QuizService_GetQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetLeaderBoard(ctx context.Context, in *Empty, opts ...grpc.CallOption) (QuizService_GetLeaderBoardClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuizService_ServiceDesc.Streams[0], QuizService_GetLeaderBoard_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &quizServiceGetLeaderBoardClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuizService_GetLeaderBoardClient interface {
	Recv() (*LeaderBoard, error)
	grpc.ClientStream
}

type quizServiceGetLeaderBoardClient struct {
	grpc.ClientStream
}

func (x *quizServiceGetLeaderBoardClient) Recv() (*LeaderBoard, error) {
	m := new(LeaderBoard)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *quizServiceClient) GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error) {
	out := new(GetScoreResponse)
	err := c.cc.Invoke(ctx, QuizService_GetScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, QuizService_GetResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) CreateUpdateQuiz(ctx context.Context, in *CreateUpdateQuizRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuizService_CreateUpdateQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) DeleteQuiz(ctx context.Context, in *DeleteQuizRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuizService_DeleteQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetAllQuizzes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (QuizService_GetAllQuizzesClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuizService_ServiceDesc.Streams[1], QuizService_GetAllQuizzes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &quizServiceGetAllQuizzesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuizService_GetAllQuizzesClient interface {
	Recv() (*Quiz, error)
	grpc.ClientStream
}

type quizServiceGetAllQuizzesClient struct {
	grpc.ClientStream
}

func (x *quizServiceGetAllQuizzesClient) Recv() (*Quiz, error) {
	m := new(Quiz)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QuizServiceServer is the server API for QuizService service.
// All implementations must embed UnimplementedQuizServiceServer
// for forward compatibility
type QuizServiceServer interface {
	GetQuiz(context.Context, *GetQuizRequest) (*GetQuizResponse, error)
	GetLeaderBoard(*Empty, QuizService_GetLeaderBoardServer) error
	GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error)
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
	CreateUpdateQuiz(context.Context, *CreateUpdateQuizRequest) (*Empty, error)
	DeleteQuiz(context.Context, *DeleteQuizRequest) (*Empty, error)
	GetAllQuizzes(*Empty, QuizService_GetAllQuizzesServer) error
	mustEmbedUnimplementedQuizServiceServer()
}

// UnimplementedQuizServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuizServiceServer struct {
}

func (UnimplementedQuizServiceServer) GetQuiz(context.Context, *GetQuizRequest) (*GetQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuiz not implemented")
}
func (UnimplementedQuizServiceServer) GetLeaderBoard(*Empty, QuizService_GetLeaderBoardServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLeaderBoard not implemented")
}
func (UnimplementedQuizServiceServer) GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
func (UnimplementedQuizServiceServer) GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedQuizServiceServer) CreateUpdateQuiz(context.Context, *CreateUpdateQuizRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUpdateQuiz not implemented")
}
func (UnimplementedQuizServiceServer) DeleteQuiz(context.Context, *DeleteQuizRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuiz not implemented")
}
func (UnimplementedQuizServiceServer) GetAllQuizzes(*Empty, QuizService_GetAllQuizzesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllQuizzes not implemented")
}
func (UnimplementedQuizServiceServer) mustEmbedUnimplementedQuizServiceServer() {}

// UnsafeQuizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizServiceServer will
// result in compilation errors.
type UnsafeQuizServiceServer interface {
	mustEmbedUnimplementedQuizServiceServer()
}

func RegisterQuizServiceServer(s grpc.ServiceRegistrar, srv QuizServiceServer) {
	s.RegisterService(&QuizService_ServiceDesc, srv)
}

func _QuizService_GetQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetQuiz(ctx, req.(*GetQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetLeaderBoard_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuizServiceServer).GetLeaderBoard(m, &quizServiceGetLeaderBoardServer{stream})
}

type QuizService_GetLeaderBoardServer interface {
	Send(*LeaderBoard) error
	grpc.ServerStream
}

type quizServiceGetLeaderBoardServer struct {
	grpc.ServerStream
}

func (x *quizServiceGetLeaderBoardServer) Send(m *LeaderBoard) error {
	return x.ServerStream.SendMsg(m)
}

func _QuizService_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetScore(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_CreateUpdateQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).CreateUpdateQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_CreateUpdateQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).CreateUpdateQuiz(ctx, req.(*CreateUpdateQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_DeleteQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).DeleteQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_DeleteQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).DeleteQuiz(ctx, req.(*DeleteQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetAllQuizzes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuizServiceServer).GetAllQuizzes(m, &quizServiceGetAllQuizzesServer{stream})
}

type QuizService_GetAllQuizzesServer interface {
	Send(*Quiz) error
	grpc.ServerStream
}

type quizServiceGetAllQuizzesServer struct {
	grpc.ServerStream
}

func (x *quizServiceGetAllQuizzesServer) Send(m *Quiz) error {
	return x.ServerStream.SendMsg(m)
}

// QuizService_ServiceDesc is the grpc.ServiceDesc for QuizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quizpb.QuizService",
	HandlerType: (*QuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuiz",
			Handler:    _QuizService_GetQuiz_Handler,
		},
		{
			MethodName: "GetScore",
			Handler:    _QuizService_GetScore_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _QuizService_GetResult_Handler,
		},
		{
			MethodName: "CreateUpdateQuiz",
			Handler:    _QuizService_CreateUpdateQuiz_Handler,
		},
		{
			MethodName: "DeleteQuiz",
			Handler:    _QuizService_DeleteQuiz_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLeaderBoard",
			Handler:       _QuizService_GetLeaderBoard_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllQuizzes",
			Handler:       _QuizService_GetAllQuizzes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "quiz.proto",
}