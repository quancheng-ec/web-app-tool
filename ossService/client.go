package ossService

import (
	"fmt"
	"os"

	"github.com/denverdino/aliyungo/oss"
)

type ServiceConfig struct {
	AccessKeyId     string     `required:"true" envconfig:"OSS_KEY"`
	AccessKeySecret string     `required:"true" envconfig:"OSS_SECRET"`
	Region          oss.Region `default:"oss-cn-hangzhou"`
	Bucket          string     `required:"true" envconfig:"OSS_BUCKET"`
}

var (
	Client *oss.Client
	Bucket *oss.Bucket
	Config *ServiceConfig
)

func Init(c *ServiceConfig) error {
	Config = c
	Client = oss.NewOSSClient(c.Region, false, c.AccessKeyId, c.AccessKeySecret, false)
	Bucket = Client.Bucket(c.Bucket)
	Client.SetDebug(true)

	return nil
}

func UploadToBucket(path string, f *os.File) error {
	fmt.Println(path)
	err := Bucket.PutFile("/test/"+path, f, "public-read", oss.Options{})
	if err != nil {
		return err
	}
	return nil
}
