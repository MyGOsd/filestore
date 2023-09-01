package helper

import (
	"cloud_disk/core/define"
	"crypto/md5"
	"errors"
	"math/rand"
	"os"
	"path"
	"time"

	"fmt"
	"net/http"
	"net/smtp"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string, second int) (string, error) {
	//id
	//identity
	//name
	uc := define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// 发送邮件验证码
func MailSendCode(To, code string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "1623294652@qq.com"
	// 设置接收方的邮箱
	e.To = []string{To}
	//设置主题
	e.Subject = "登录验证码"
	//设置文件发送的内容
	bodyText := fmt.Sprintf("您的验证码是:%s， 有效期为5分钟", code)
	e.Text = []byte(bodyText)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1623294652@qq.com", "xsupdqqvaxapdhjj", "smtp.qq.com"))
	if err != nil {
		return err
	}
	return nil
}

// 生成随机数
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
func GetUUID() string {
	return uuid.NewV4().String()
}

// 上传文件到oss
func OssUpload(r *http.Request) (string, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", define.Oos_key, define.Oos_AccessKey_Secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(define.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传文件，并获取完整路径
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)

	// 将文件流上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject(key, file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return define.BucketName + "/" + key, err
}

// 分片上传文件到oss
func OssInitPart(r *http.Request) (string, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", define.Oos_key, define.Oos_AccessKey_Secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(define.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传文件，并获取完整路径
	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)
	// 将本地文件分片，且分片数量指定为3。
	fd, openError := fileHeader.Open()
	if openError != nil {
		return "", openError
	}
	chunks, err := SplitFileByPartNum(fileHeader.Size, 3)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}
	imur, err := bucket.InitiateMultipartUpload(key, options...)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		parts = append(parts, part)
	}
	// 指定Object的读写权限为私有，默认为继承Bucket的读写权限
	objectAcl := oss.ObjectACL(oss.ACLPrivate)
	// 步骤3：完成分片上传，指定文件可以预览。
	_, err = bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return define.BucketName + "/" + key, err
}

// 分片计算
func SplitFileByPartNum(size int64, chunkNum int) ([]oss.FileChunk, error) {
	if chunkNum <= 0 || chunkNum > 10000 {
		return nil, errors.New("chunkNum invalid")
	}
	if int64(chunkNum) > size {
		return nil, errors.New("oss: chunkNum invalid")
	}
	var chunks []oss.FileChunk
	var chunk = oss.FileChunk{}
	var chunkN = (int64)(chunkNum)
	for i := int64(0); i < chunkN; i++ {
		chunk.Number = int(i + 1)
		chunk.Offset = i * (size / chunkN)
		if i == chunkN-1 {
			chunk.Size = size/chunkN + size%chunkN
		} else {
			chunk.Size = size / chunkN
		}
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}
