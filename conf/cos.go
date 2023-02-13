package conf

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Client *cos.Client

func InitCos() {
	u, err := url.Parse(CosUrl)
	if err != nil {
		panic(err)
	}
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecreteKey,
		},
	})
}
