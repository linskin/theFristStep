package feature

import (
	"context"
	"example.com/m/v2/conf"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
)

var Client *cos.Client

func InitCos() {
	u, err := url.Parse(conf.CosUrl)
	if err != nil {
		panic(err)
	}
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecreteKey,
		},
	})
}
func Upload(file *multipart.FileHeader, saveName string) error {
	fd, err := file.Open()
	if err != nil {
		return err
	}
	if _, err = Client.Object.Put(context.Background(), saveName, fd, nil); err != nil {
		return err
	}
	return nil
}
