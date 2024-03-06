// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: main.proto

package toclassifier

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

// ToClassifierClient is the client API for ToClassifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToClassifierClient interface {
	TrainOnSavedDocuments(ctx context.Context, in *Username, opts ...grpc.CallOption) (*TrainResponse, error)
	TrainActions(ctx context.Context, opts ...grpc.CallOption) (ToClassifier_TrainActionsClient, error)
	TrainonSavedDocumentsJson(ctx context.Context, in *JsonData, opts ...grpc.CallOption) (*TrainResponse, error)
	QueryActions(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ActionFull, error)
	GiveAudioAnswerForQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error)
	GiveAudioAnswerOrJustTextAnswer(ctx context.Context, in *Query, opts ...grpc.CallOption) (*AudoWithText, error)
	SaveDocuments(ctx context.Context, in *Documents, opts ...grpc.CallOption) (*GeneralAnswer, error)
	GetGreetingMessage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error)
	// to classifier service
	ClassifyAndAnswer(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error)
	OpenChat(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ChatID, error)
	CloseChat(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*GeneralAnswer, error)
}

type toClassifierClient struct {
	cc grpc.ClientConnInterface
}

func NewToClassifierClient(cc grpc.ClientConnInterface) ToClassifierClient {
	return &toClassifierClient{cc}
}

func (c *toClassifierClient) TrainOnSavedDocuments(ctx context.Context, in *Username, opts ...grpc.CallOption) (*TrainResponse, error) {
	out := new(TrainResponse)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/TrainOnSavedDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) TrainActions(ctx context.Context, opts ...grpc.CallOption) (ToClassifier_TrainActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ToClassifier_ServiceDesc.Streams[0], "/toclassifier.ToClassifier/TrainActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &toClassifierTrainActionsClient{stream}
	return x, nil
}

type ToClassifier_TrainActionsClient interface {
	Send(*ActionFull) error
	CloseAndRecv() (*TrainResponse, error)
	grpc.ClientStream
}

type toClassifierTrainActionsClient struct {
	grpc.ClientStream
}

func (x *toClassifierTrainActionsClient) Send(m *ActionFull) error {
	return x.ClientStream.SendMsg(m)
}

func (x *toClassifierTrainActionsClient) CloseAndRecv() (*TrainResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TrainResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toClassifierClient) TrainonSavedDocumentsJson(ctx context.Context, in *JsonData, opts ...grpc.CallOption) (*TrainResponse, error) {
	out := new(TrainResponse)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/TrainonSavedDocumentsJson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) QueryActions(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ActionFull, error) {
	out := new(ActionFull)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/QueryActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) GiveAudioAnswerForQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error) {
	out := new(GeneralAnswer)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/GiveAudioAnswerForQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) GiveAudioAnswerOrJustTextAnswer(ctx context.Context, in *Query, opts ...grpc.CallOption) (*AudoWithText, error) {
	out := new(AudoWithText)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/GiveAudioAnswerOrJustTextAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) SaveDocuments(ctx context.Context, in *Documents, opts ...grpc.CallOption) (*GeneralAnswer, error) {
	out := new(GeneralAnswer)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/SaveDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) GetGreetingMessage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error) {
	out := new(GeneralAnswer)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/GetGreetingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) ClassifyAndAnswer(ctx context.Context, in *Query, opts ...grpc.CallOption) (*GeneralAnswer, error) {
	out := new(GeneralAnswer)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/ClassifyAndAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) OpenChat(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ChatID, error) {
	out := new(ChatID)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/OpenChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toClassifierClient) CloseChat(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*GeneralAnswer, error) {
	out := new(GeneralAnswer)
	err := c.cc.Invoke(ctx, "/toclassifier.ToClassifier/CloseChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToClassifierServer is the server API for ToClassifier service.
// All implementations must embed UnimplementedToClassifierServer
// for forward compatibility
type ToClassifierServer interface {
	TrainOnSavedDocuments(context.Context, *Username) (*TrainResponse, error)
	TrainActions(ToClassifier_TrainActionsServer) error
	TrainonSavedDocumentsJson(context.Context, *JsonData) (*TrainResponse, error)
	QueryActions(context.Context, *Query) (*ActionFull, error)
	GiveAudioAnswerForQuery(context.Context, *Query) (*GeneralAnswer, error)
	GiveAudioAnswerOrJustTextAnswer(context.Context, *Query) (*AudoWithText, error)
	SaveDocuments(context.Context, *Documents) (*GeneralAnswer, error)
	GetGreetingMessage(context.Context, *Query) (*GeneralAnswer, error)
	// to classifier service
	ClassifyAndAnswer(context.Context, *Query) (*GeneralAnswer, error)
	OpenChat(context.Context, *Query) (*ChatID, error)
	CloseChat(context.Context, *ChatID) (*GeneralAnswer, error)
	mustEmbedUnimplementedToClassifierServer()
}

// UnimplementedToClassifierServer must be embedded to have forward compatible implementations.
type UnimplementedToClassifierServer struct {
}

func (UnimplementedToClassifierServer) TrainOnSavedDocuments(context.Context, *Username) (*TrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainOnSavedDocuments not implemented")
}
func (UnimplementedToClassifierServer) TrainActions(ToClassifier_TrainActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method TrainActions not implemented")
}
func (UnimplementedToClassifierServer) TrainonSavedDocumentsJson(context.Context, *JsonData) (*TrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainonSavedDocumentsJson not implemented")
}
func (UnimplementedToClassifierServer) QueryActions(context.Context, *Query) (*ActionFull, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActions not implemented")
}
func (UnimplementedToClassifierServer) GiveAudioAnswerForQuery(context.Context, *Query) (*GeneralAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveAudioAnswerForQuery not implemented")
}
func (UnimplementedToClassifierServer) GiveAudioAnswerOrJustTextAnswer(context.Context, *Query) (*AudoWithText, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveAudioAnswerOrJustTextAnswer not implemented")
}
func (UnimplementedToClassifierServer) SaveDocuments(context.Context, *Documents) (*GeneralAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveDocuments not implemented")
}
func (UnimplementedToClassifierServer) GetGreetingMessage(context.Context, *Query) (*GeneralAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreetingMessage not implemented")
}
func (UnimplementedToClassifierServer) ClassifyAndAnswer(context.Context, *Query) (*GeneralAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassifyAndAnswer not implemented")
}
func (UnimplementedToClassifierServer) OpenChat(context.Context, *Query) (*ChatID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenChat not implemented")
}
func (UnimplementedToClassifierServer) CloseChat(context.Context, *ChatID) (*GeneralAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseChat not implemented")
}
func (UnimplementedToClassifierServer) mustEmbedUnimplementedToClassifierServer() {}

// UnsafeToClassifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToClassifierServer will
// result in compilation errors.
type UnsafeToClassifierServer interface {
	mustEmbedUnimplementedToClassifierServer()
}

func RegisterToClassifierServer(s grpc.ServiceRegistrar, srv ToClassifierServer) {
	s.RegisterService(&ToClassifier_ServiceDesc, srv)
}

func _ToClassifier_TrainOnSavedDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).TrainOnSavedDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/TrainOnSavedDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).TrainOnSavedDocuments(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_TrainActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ToClassifierServer).TrainActions(&toClassifierTrainActionsServer{stream})
}

type ToClassifier_TrainActionsServer interface {
	SendAndClose(*TrainResponse) error
	Recv() (*ActionFull, error)
	grpc.ServerStream
}

type toClassifierTrainActionsServer struct {
	grpc.ServerStream
}

func (x *toClassifierTrainActionsServer) SendAndClose(m *TrainResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *toClassifierTrainActionsServer) Recv() (*ActionFull, error) {
	m := new(ActionFull)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ToClassifier_TrainonSavedDocumentsJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).TrainonSavedDocumentsJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/TrainonSavedDocumentsJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).TrainonSavedDocumentsJson(ctx, req.(*JsonData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_QueryActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).QueryActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/QueryActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).QueryActions(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_GiveAudioAnswerForQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).GiveAudioAnswerForQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/GiveAudioAnswerForQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).GiveAudioAnswerForQuery(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_GiveAudioAnswerOrJustTextAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).GiveAudioAnswerOrJustTextAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/GiveAudioAnswerOrJustTextAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).GiveAudioAnswerOrJustTextAnswer(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_SaveDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Documents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).SaveDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/SaveDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).SaveDocuments(ctx, req.(*Documents))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_GetGreetingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).GetGreetingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/GetGreetingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).GetGreetingMessage(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_ClassifyAndAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).ClassifyAndAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/ClassifyAndAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).ClassifyAndAnswer(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_OpenChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).OpenChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/OpenChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).OpenChat(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToClassifier_CloseChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToClassifierServer).CloseChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toclassifier.ToClassifier/CloseChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToClassifierServer).CloseChat(ctx, req.(*ChatID))
	}
	return interceptor(ctx, in, info, handler)
}

// ToClassifier_ServiceDesc is the grpc.ServiceDesc for ToClassifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToClassifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "toclassifier.ToClassifier",
	HandlerType: (*ToClassifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainOnSavedDocuments",
			Handler:    _ToClassifier_TrainOnSavedDocuments_Handler,
		},
		{
			MethodName: "TrainonSavedDocumentsJson",
			Handler:    _ToClassifier_TrainonSavedDocumentsJson_Handler,
		},
		{
			MethodName: "QueryActions",
			Handler:    _ToClassifier_QueryActions_Handler,
		},
		{
			MethodName: "GiveAudioAnswerForQuery",
			Handler:    _ToClassifier_GiveAudioAnswerForQuery_Handler,
		},
		{
			MethodName: "GiveAudioAnswerOrJustTextAnswer",
			Handler:    _ToClassifier_GiveAudioAnswerOrJustTextAnswer_Handler,
		},
		{
			MethodName: "SaveDocuments",
			Handler:    _ToClassifier_SaveDocuments_Handler,
		},
		{
			MethodName: "GetGreetingMessage",
			Handler:    _ToClassifier_GetGreetingMessage_Handler,
		},
		{
			MethodName: "ClassifyAndAnswer",
			Handler:    _ToClassifier_ClassifyAndAnswer_Handler,
		},
		{
			MethodName: "OpenChat",
			Handler:    _ToClassifier_OpenChat_Handler,
		},
		{
			MethodName: "CloseChat",
			Handler:    _ToClassifier_CloseChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TrainActions",
			Handler:       _ToClassifier_TrainActions_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "main.proto",
}
