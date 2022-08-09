package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	BasicAuthUsername string `json:",optional"`
	BasicAuthPassword string `json:",optional"`
}
