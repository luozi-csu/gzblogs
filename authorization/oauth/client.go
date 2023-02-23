package oauth

import (
	"net"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const (
	GithubAuthType string = "github"
)

type OAuthClient interface {
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}

var (
	defaultHttpClient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
		Timeout: 10 * time.Second,
	}
)

type UserInfo struct {
	ID          int
	Url         string
	AuthType    string
	Username    string
	DisplayName string
	Email       string
	AvatarUrl   string
}
