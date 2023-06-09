package define

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var _ = os.Setenv("MailPassword", "VWUMHUFFPSZWKZEB")
var MailPassWord = os.Getenv("MailPassword")

// CodeLength 定义验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecreKey 腾讯云对象存储
var _ = os.Setenv("TencentSecretID", "AKIDscbfdOxSnANPHySo4qSs57xYyEh3JP1e")
var TencentSecretID = os.Getenv("TencentSecretID")
var _ = os.Setenv("TencentSecretKey", "jKrHXaUEm5B7tSoX7TbWIyUoxDJBtVAh")
var TencentSecretKey = os.Getenv("TencentSecretKey")

// 存储桶路径
var CosBucket = "https://tanyuyan-19990720-1318633606.cos.ap-nanjing.myqcloud.com"
