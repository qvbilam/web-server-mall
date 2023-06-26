package initialize

import (
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	proto "mall/api/qvbilam/mall/v1"
	messageProto "mall/api/qvbilam/message/v1"
	userProto "mall/api/qvbilam/user/v1"
	"mall/global"
	"time"
)

type dialConfig struct {
	host string
	port int64
}

type serverClientConfig struct {
	mallDialConfig    *dialConfig
	messageDialConfig *dialConfig
	userDialConfig    *dialConfig
}

func InitServer() {
	s := serverClientConfig{
		mallDialConfig: &dialConfig{
			host: global.ServerConfig.MallServerConfig.Host,
			port: global.ServerConfig.MallServerConfig.Port,
		},
		messageDialConfig: &dialConfig{
			host: global.ServerConfig.MessageServerConfig.Host,
			port: global.ServerConfig.MessageServerConfig.Port,
		},
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
		},
	}

	s.initMallServer()
	s.initMessageServer()
	s.initUserServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initMallServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.mallDialConfig.host, s.mallDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.MallServerConfig.Name, err)
	}

	client := proto.NewMallClient(conn)
	global.MallServerClient = client
}

func (s *serverClientConfig) initMessageServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.messageDialConfig.host, s.messageDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.MessageServerConfig.Name, err)
	}

	client := messageProto.NewMessageClient(conn)
	global.MessageServerClient = client
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.UserServerConfig.Name, err)
	}

	client := userProto.NewUserClient(conn)
	global.UserServerClient = client
}
