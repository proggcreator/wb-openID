package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) showAll(c *gin.Context) {

	mystr := MyStruct{
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
	c.Writer.Write(jData)
}
