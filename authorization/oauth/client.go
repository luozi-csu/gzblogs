package oauth

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/model"
	"golang.org/x/oauth2"
)

const (
	GithubAuthType = "github"
	EmptyAuthType  = ""
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
	ID          string
	Url         string
	AuthType    string
	Username    string
	DisplayName string
	Email       string
	AvatarUrl   string
}

func (ui *UserInfo) GetUser() *model.User {
	return &model.User{
		Name:   ui.Username,
		Email:  ui.Email,
		Avatar: ui.AvatarUrl,
		AuthInfos: []model.AuthInfo{
			{
				AuthType: ui.AuthType,
				AuthID:   ui.ID,
				Url:      ui.Url,
			},
		},
	}
}

type OAuthManager struct {
	conf *config.OAuthConfig
}

func NewOAuthManager(conf *config.OAuthConfig) *OAuthManager {
	return &OAuthManager{conf: conf}
}

func (om *OAuthManager) GetOAuthClient(authType string) (OAuthClient, error) {
	switch authType {
	case GithubAuthType:
		return NewGithubAuth(om.conf.Github.ClientID, om.conf.Github.ClientSecret), nil
	}

	return nil, fmt.Errorf("unknown auth type=%s", authType)
}
