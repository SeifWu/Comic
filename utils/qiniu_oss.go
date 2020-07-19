package utils

import (
	"context"
	"log"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
)

// QiNiuUpToken 上传 Token
func QiNiuUpToken() string {
	accessKey := viper.GetString("qiniu.accessKey")
	secretKey := viper.GetString("qiniu.secretKey")
	bucket := viper.GetString("qiniu.bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	log.Println(upToken)
	return upToken
}

// QiNiuUpload 上传文件
func QiNiuUpload(key string, localFile string) map[string]string {
	result := make(map[string]string)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, QiNiuUpToken(), key, localFile, &putExtra)
	if err != nil {
		log.Println(err)
		result["key"] = "err"
		result["hash"] = "exist"
		return result
	}
	result["key"] = ret.Key
	result["hash"] = ret.Hash
	return result
}
