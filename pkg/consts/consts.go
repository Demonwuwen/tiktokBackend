package consts

const (
	VideoTableName     = "videos"
	UserTableName      = "users"
	UserLoginTableName = "user_logins"
	SecretKey          = "secret key"
	IdentityKey        = "id"
	Total              = "total"
	Notes              = "notes"
	ApiServiceName     = "api"
	VideoServiceName   = "video"
	UserServiceName    = "user"
	MySQLDefaultDSN    = "demon:123456@tcp(192.168.14.194:3307)/tiktok?charset=utf8&parseTime=True&loc=Local"
	//MySQLDefaultDSN  = "root:Yww=Hxq13@tcp(localhost:3306)/tiktok?charset=utf8&parseTime=True&loc=Local"
	TCP              = "tcp"
	UserServiceAddr  = ":9000"
	VideoServiceAddr = ":10000"
	ExportEndpoint   = ":4317"
	ETCDAddress      = "127.0.0.1:2379"
	DefaultLimit     = 10
	RandNameLen      = 8
	SuccessCode      = 0
	FailedCode       = 110
)
