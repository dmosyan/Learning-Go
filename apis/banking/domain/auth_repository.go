package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:3001", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
