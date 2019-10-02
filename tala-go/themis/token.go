package themis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type ClientToken struct {
	host         string
	clientID     string
	clientSecret string
	http         HttpDoer

	token          *AccessToken
	lastFetchToken time.Time
	muToken        sync.RWMutex
}

var _ = TokenProvider(&ClientToken{})

func NewClientToken(host, clientID, clientSecret string, httpClient HttpDoer) *ClientToken {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &ClientToken{
		host:         host,
		clientID:     clientID,
		clientSecret: clientID,
		http:         httpClient,
	}
}

func (c *ClientToken) FetchToken(ctx context.Context) (*AccessToken, error) {
	params := url.Values{}
	params.Set("client_id", c.clientID)
	params.Set("client_secret", c.clientSecret)
	params.Set("grant_type", "client_credentials")
	params.Set("scope", "tiki.api")

	url := fmt.Sprintf("%s/oauth2/token", c.host)
	req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req = req.WithContext(ctx)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := &struct {
		*AccessToken
		*ThemisError
	}{}
	if err = json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, err
	}
	if data.Error != nil {
		return nil, data.ThemisError
	}
	return data.AccessToken, nil
}

func (c *ClientToken) AccessToken(ctx context.Context) (string, error) {
	c.muToken.Lock()
	defer c.muToken.Unlock()

	if c.token == nil || c.token.IsExpires(c.lastFetchToken) {
		token, err := c.FetchToken(ctx)
		if err != nil {
			return "", err
		}

		c.token = token
		c.lastFetchToken = time.Now()
	}

	return fmt.Sprintf("%s %s", c.token.TokenType, c.token.AccessToken), nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (t AccessToken) IsExpires(tm time.Time) bool {
	delta := time.Now().Sub(tm.Add(time.Second * time.Duration(t.ExpiresIn)))
	return delta < time.Second*2
}

type ThemisError struct {
	ErrorType        string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorHint        string `json:"error_hint"`
	StatusCode       int    `json:"status_code"`
}

func (e ThemisError) Error() string {
	return fmt.Sprintf("Themis(%s): %s", e.ErrorType, e.ErrorDescription)
}
