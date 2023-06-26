package global

import (
	ut "github.com/go-playground/universal-translator"
	proto "mall/api/qvbilam/mall/v1"
	messageProto "mall/api/qvbilam/message/v1"
	userProto "mall/api/qvbilam/user/v1"
	"mall/config"
)

var (
	Trans               ut.Translator // 表单验证
	ServerConfig        *config.ServerConfig
	MessageServerClient messageProto.MessageClient
	UserServerClient    userProto.UserClient
	MallServerClient    proto.MallClient
)
