package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	NeedClaimes = "world"
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

func ParseAccess(t MyToken) bool {
	//parse access token
	flagClaims := false
	flagLifetime := false
	tokenString := t.Access_token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Hello"), nil
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
	keyExp := "exp"
	exp := reflect.ValueOf(claims[keyExp]).Elem().Int()
	if time.Now().Unix() < exp*1000 {
		flagLifetime = true
	}

	return flagClaims
}

func (h *Handler) showAll(c *gin.Context) {

	/*mystr := MyStruct{
		Id:        1,
		Name:      "john",
		Last_name: "wick",
	}
	//to JSON
	jData, err := json.Marshal(mystr)
	if err != nil {
		logrus.Printf("Error marshaling")
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(jData)*/
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
	fmt.Fprintln(c.Writer, params)
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
		fmt.Println("YYYYYYYYYYYYYYYy")
		return
	} else {
		fmt.Println("NNNNNNNNNNNNNNNNN")
	}

	//fmt.Printf(string(tokenn))
}
