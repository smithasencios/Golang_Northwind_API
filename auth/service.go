package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Service interface {
	GetPermissions(ctx context.Context, request *getPermissionsByUserId) (*PermissionList, error)
}
type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetPermissions(ctx context.Context, request *getPermissionsByUserId) (*PermissionList, error) {
	url := os.Getenv("AUTHO_URL") + "/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=" + os.Getenv("AUTHO_CLIENT_ID") + "&client_secret=" + os.Getenv("AUTHO_CLIENT_SECRET") + "&audience=" + os.Getenv("AUTHO_URL") + "/api/v2/")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", body)
	if err != nil {
		panic(err.Error())
	}

	var accessTokenItem *AccessTokenItem
	err = json.Unmarshal(body, &accessTokenItem)

	if err != nil {
		panic(err.Error())
	}
	return GetPermissionsByUserId(request.UserID, accessTokenItem.Access_Token)
}

func GetPermissionsByUserId(userId string, accessToken string) (*PermissionList, error) {
	url := os.Getenv("AUTHO_URL") + "/api/v2/users/" + userId + "/permissions"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Add("authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var permissions []*Permission
	err = json.Unmarshal(body, &permissions)

	if err != nil {
		panic(err.Error())
	}

	return &PermissionList{Data: permissions}, err
}
