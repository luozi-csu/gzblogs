package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type GithubAuth struct {
	client *http.Client
	config *oauth2.Config
}

func NewGithubAuth(clientID, clientSecret string) *GithubAuth {
	auth := &GithubAuth{
		client: defaultHttpClient,
		config: &oauth2.Config{
			Scopes: []string{"user"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://github.com/login/oauth/authorize",
				TokenURL: "https://github.com/login/oauth/access_token",
			},
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}
	return auth
}

type GithubToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// response schema from GitHub REST API https://api.github.com/user
type GitHubUserInfo struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeId                  string      `json:"node_id"`
	AvatarUrl               string      `json:"avatar_url"`
	GravatarId              string      `json:"gravatar_id"`
	Url                     string      `json:"url"`
	HtmlUrl                 string      `json:"html_url"`
	FollowersUrl            string      `json:"followers_url"`
	FollowingUrl            string      `json:"following_url"`
	GistsUrl                string      `json:"gists_url"`
	StarredUrl              string      `json:"starred_url"`
	SubscriptionsUrl        string      `json:"subscriptions_url"`
	OrganizationsUrl        string      `json:"organizations_url"`
	ReposUrl                string      `json:"repos_url"`
	EventsUrl               string      `json:"events_url"`
	ReceivedEventsUrl       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   string      `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         interface{} `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

func (g *GithubAuth) GetToken(code string) (*oauth2.Token, error) {
	if len(g.config.ClientID) == 0 || len(g.config.ClientSecret) == 0 {
		return nil, errors.New("github oauth client id or client secret is empty")
	}

	body := &struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}{g.config.ClientID, g.config.ClientSecret, code}

	s, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "marshal json body failed")
	}

	req, err := http.NewRequest("POST", g.config.Endpoint.TokenURL, strings.NewReader(string(s)))
	if err != nil {
		return nil, errors.Wrap(err, "create http request failed")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client do request failed")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body failed")
	}

	githubToken := new(GithubToken)
	if err = json.Unmarshal(data, githubToken); err != nil {
		return nil, errors.Wrap(err, "parse response body failed")
	}

	return &oauth2.Token{
		AccessToken: githubToken.AccessToken,
		TokenType:   "Bearer",
	}, nil
}

func (g *GithubAuth) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	if token == nil {
		return nil, errors.New("emtpy oauth token")
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, errors.Wrap(err, "create http request failed")
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client do request failed")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body failed")
	}

	githubUserInfo := new(GitHubUserInfo)
	if err := json.Unmarshal(data, githubUserInfo); err != nil {
		return nil, errors.Wrap(err, "parse response body failed")
	}

	return &UserInfo{
		ID:          githubUserInfo.ID,
		Url:         githubUserInfo.Url,
		AuthType:    GithubAuthType,
		Username:    githubUserInfo.Login,
		DisplayName: githubUserInfo.Name,
		Email:       githubUserInfo.Email,
		AvatarUrl:   githubUserInfo.AvatarUrl,
	}, nil
}
