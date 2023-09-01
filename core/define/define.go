package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// 验证码长度
var CodeLength = 6

// 验证码过期时间(s)
var CodeExpire = 300

// oss的环境变量配置
// yourBucketName填写存储空间名称。
var BucketName = "mycloud-disk"

// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
var objectName = "cloud-disk/"

var Oos_key = "LTAI5t9NbaLgyNYsCSACyWAZ"

var Oos_AccessKey_Secret = "F8V57WuoDn6DgGEWATlJFXEsHVjRhW"

var bucketName = "mycloud-disk"

// 分页默认参数
var PageSize = 20

var Datatime = "2006-01-02 15:04:05"

var TokenExpire = 3600

var RefreshTokenExpire = 36000
