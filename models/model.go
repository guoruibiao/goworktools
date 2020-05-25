package models

// 系统环境变量使用
type App struct {
	Server Server `toml:"server"`

	RedisTest RedisItem `toml:"redistest"`
	RedisBNS RedisBNS `toml:"redisbns"`
}

type Server struct {
	Port int `toml:"port"`
}
type RedisBNS struct {
	Names []string
}

type RedisItem struct {
	Host string
	Port int
}

// 内部变量定义
type BNSItem struct {
	Host string
	Port int
}

type InstanceContainer struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data []Instance `json:"data"`
}
type Instance struct {
	HostName string `json:"hostName"`
	Port string `json:"port"`
}