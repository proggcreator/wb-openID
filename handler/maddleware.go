package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	NeedClaimes = "world"
	keyExp      = "exp"
)

type MyStruct struct {
	Id        int
	Name      string
	Last_name string
}
type MyToken struct {
	Access_token string
	Expires_in   int
	Token_type   string
}

type UserAuth struct {
	ClientId     string
	ClientSecret string
	GrandType    string
}

func ParseAccess(t MyToken) bool {
	//parse access token
	flagClaims := false
	flagLifetime := false
	tokenString := t.Access_token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("KeyString"), nil
	})
	if err != nil {

	}
	keyclaims := "client_claims"
	//for range over claims interface
	s := reflect.ValueOf(claims[keyclaims])
	for i := 0; i < s.Len(); i++ {
		//check needed claims
		if s.Index(i).Elem().String() == NeedClaimes {
			flagClaims = true
		}
	}
	//check a token lifetime

	var tm time.Time
	//convert exp to time.Time
	switch exp := claims[keyExp].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}
	if time.Now().Unix() <= tm.Unix() {
		flagLifetime = true
	}

	return flagClaims && flagLifetime
}

func (h *Handler) userIdentity(c *gin.Context) {

	curUserId := UserAuth{
		ClientId:     "finance-test-client",
		ClientSecret: "testsecret",
		GrandType:    "client_credentials",
	}
	client := &http.Client{}
	//post request application/x-www-form-urlencoded
	buffer := new(bytes.Buffer)
	params := url.Values{}

	params.Set("grant_type", curUserId.GrandType)
	params.Set("client_id", curUserId.ClientId)
	params.Set("client_secret", curUserId.ClientSecret)
	buffer.WriteString(params.Encode())
	req, _ := http.NewRequest("POST", "https://payments.wildberries.ru/authtest/connect/token", buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)

	//decode body json
	var t MyToken

	err := json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	//if neened claims is exist
	if ParseAccess(t) {
		return
	} else {
		logrus.Fatalf("Error token verification ")
	}

}
