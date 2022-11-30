package uploader

import (
	"context"
	"fmt"
	"github.com/iseki0/cos-uploader/env"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"path/filepath"
)

var client *cos.Client

func InitClient() {
	if env.CosBucketUrl == nil || env.CosServiceUrl == nil {
		panic("init")
	}
	baseUrl := &cos.BaseURL{
		BucketURL:  env.CosBucketUrl,
		ServiceURL: env.CosServiceUrl,
	}

	client = cos.NewClient(baseUrl, &http.Client{Transport: &cos.AuthorizationTransport{
		SecretID:  env.CosSecretId,
		SecretKey: env.CosSecretKey,
	}})

}

func UploadFile(file string, target string) error {
	fn := filepath.Base(file)
	result, response, e := client.Object.Upload(context.TODO(), filepath.Join(target, fn), file, nil)
	if e != nil {
		return e
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("file: %s, status code %d", file, response.StatusCode)
	}
	fmt.Printf("upload, location: %s", result.Location)
	return nil
}
