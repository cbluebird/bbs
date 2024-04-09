package constants

import "bbs/pkg/config/viper"

const (
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	ApiServiceName          = "bbs_api"
	UserServiceName         = "bbs_user"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)

var (
	EtcdAddress = viper.Config.GetString("etcd")
	IP          = viper.Config.GetString("ip")
	UserPort    = viper.Config.GetString("port.user")
)
