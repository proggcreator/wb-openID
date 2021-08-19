package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	ClientId     string
	ClientSecret string
	GrandType    string
}

func (h *Handler) userIdentity(c *gin.Context) {
	curUserId := UserAuth{
		ClientId:     "finance-test-client",
		ClientSecret: "testsecret",
	}
	//post request application/x-www-form-urlencoded
	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("client_id", curUserId.ClientId)
	params.Set("client_secret", curUserId.ClientSecret)
	buffer.WriteString(params.Encode())
	req, _ := http.NewRequest("POST", "https://payments.wildberries.ru/authtest/connect/token", buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	fmt.Fprintln(c.Writer, "Hello")
	fmt.Fprintln(c.Writer, req)

}
