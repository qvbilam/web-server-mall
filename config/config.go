package config

type ServerConfig struct {
	Name                string              `mapstructure:"name" json:"name"`
	Host                string              `mapstructure:"host" json:"host"`
	Port                int64               `mapstructure:"port" json:"port"`
	Tags                []string            `mapstructure:"tags" json:"tags"`
	UserServerConfig    UserServerConfig    `mapstructure:"user-server" json:"user-server"`
	MessageServerConfig MessageServerConfig `mapstructure:"message-server" json:"message-server"`
	MallServerConfig    MallServerConfig    `mapstructure:"mall-server" json:"mall-server"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type MessageServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type MallServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
