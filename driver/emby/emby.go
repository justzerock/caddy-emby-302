package emby

import "github.com/go-resty/resty/v2"

type DirverEmby struct {
	Url         string
	Token       string
	RestyClient *resty.Client
}
