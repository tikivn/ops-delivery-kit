package themis

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type TokenProvider interface {
	AccessToken(context.Context) (string, error)
}

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type User struct {
	Trn   string `json:"trn"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Authorize struct {
	Token    string `json:"token"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type Paging struct {
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
	From        int `json:"from"`
	To          int `json:"to"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
}

type Policy struct {
	Trn         string `json:"trn"`
	Key         string `json:"key"`
	Description string `json:"description"`
}

type groupResponse struct {
	Data []string `json:"data"`
}

type policiesResponse struct {
	Data []*Policy `json:"data"`
}

type userResponse struct {
	Data   []*User `json:"data"`
	Paging Paging  `json:"paging"`
}

type authorizeResponse struct {
	Allowed bool `json:"allowed"`
}

type Client interface {
	User(ctx context.Context, trn string) (*User, error)
	Users(ctx context.Context, name string, email string, page int, size int) ([]*User, int, error)
	UserGroups(ctx context.Context, trn string) ([]string, error)
	GroupPolicies(ctx context.Context, trn string) ([]*Policy, error)
	UsersInGroup(ctx context.Context, group string) ([]*User, error)
	AddUsersInGroup(ctx context.Context, group string, trns []string) error
	RemoveUsersInGroup(ctx context.Context, group string, trns []string) error

	AuthorizeSubject(ctx context.Context, resource, action, subject string) (bool, error)
	AuthorizeToken(ctx context.Context, resource, action, token string) (bool, error)
}

type client struct {
	host       string
	httpClient HttpDoer
	token      TokenProvider
}

func NewClient(host string, httpClient HttpDoer, token TokenProvider) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &client{
		host:       host,
		httpClient: httpClient,
		token:      token,
	}
}

func (t *client) User(ctx context.Context, trn string) (*User, error) {
	var userResponse User
	path := fmt.Sprintf("users/%s", trn)

	err := t.call(ctx, path, "GET", nil, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (t *client) Users(ctx context.Context, name string, email string, page int, size int) ([]*User, int, error) {
	var userResponse userResponse
	params := url.Values{}
	if name != "" {
		params.Add("name", name)
	}

	if email != "" {
		params.Add("email", email)
	}

	if page != 0 {
		params.Add("page", strconv.Itoa(page))
	}

	if size != 0 {
		params.Add("limit", strconv.Itoa(size))
	}

	path := fmt.Sprintf("users?%s", params.Encode())

	err := t.call(ctx, path, "GET", nil, &userResponse)
	if err != nil {
		return nil, 0, err
	}

	return userResponse.Data, userResponse.Paging.Total, nil
}

func (t *client) UserGroups(ctx context.Context, trn string) ([]string, error) {
	var result groupResponse
	path := fmt.Sprintf("users/%s/groups", trn)

	err := t.call(ctx, path, "GET", nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (t *client) GroupPolicies(ctx context.Context, trn string) ([]*Policy, error) {
	var result policiesResponse
	path := fmt.Sprintf("groups/%s/policies", trn)

	err := t.call(ctx, path, "GET", nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (t *client) UsersInGroup(ctx context.Context, group string) ([]*User, error) {
	var userResponse userResponse
	path := fmt.Sprintf("groups/%s/users", group)

	err := t.call(ctx, path, "GET", nil, &userResponse)
	if err != nil {
		return nil, err
	}

	return userResponse.Data, nil
}

func (t *client) AddUsersInGroup(ctx context.Context, group string, IDs []string) error {
	var userResponse userResponse
	path := fmt.Sprintf("groups/%s/users", group)

	err := t.call(ctx, path, "POST", IDs, &userResponse)
	if err != nil {
		return err
	}

	return nil
}

func (t *client) RemoveUsersInGroup(ctx context.Context, group string, IDs []string) error {
	var userResponse userResponse
	path := fmt.Sprintf("groups/%s/users", group)

	err := t.call(ctx, path, "DELETE", IDs, &userResponse)
	if err != nil {
		return err
	}

	return nil
}

func (t *client) AuthorizeSubject(ctx context.Context, resource, action, subject string) (bool, error) {
	var authorizeResponse authorizeResponse
	path := fmt.Sprintf("subjects/authorize?resource=%s&action=%s&subject=%s",
		resource, action, subject,
	)

	err := t.call(ctx, path, "GET", nil, &authorizeResponse)
	if err != nil {
		return false, err
	}

	return authorizeResponse.Allowed, nil
}

func (t *client) AuthorizeToken(ctx context.Context, resource, action, token string) (bool, error) {
	var authorizeResponse authorizeResponse
	path := fmt.Sprintf("access-tokens/authorize?resource=%s&action=%s&token=%s",
		resource, action, token,
	)

	err := t.call(ctx, path, "POST", Authorize{
		Token:    token,
		Resource: resource,
		Action:   action,
	}, &authorizeResponse)
	if err != nil {
		return false, err
	}

	return authorizeResponse.Allowed, nil
}

func (t *client) call(ctx context.Context, path string, method string, data interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", t.host, path)
	var request *http.Request
	var err error

	switch method {
	case "GET":
		request, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

	case "POST", "DELETE":
		var body []byte
		if data != nil {
			body, err = json.Marshal(data)

			if err != nil {
				return err
			}
		}

		request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		request.Header.Add("Content-Type", "application/json")

		if err != nil {
			return err
		}

	default:
		return errors.New("Unknown Method")
	}

	request = request.WithContext(ctx)

	if t.token != nil {
		token, err := t.token.AccessToken(ctx)
		if err == nil && token != "" {
			request.Header.Add("Authorization", token)
		}
	}

	resp, err := t.httpClient.Do(request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Themis(%d), url = %s", resp.StatusCode, path)
	}

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
