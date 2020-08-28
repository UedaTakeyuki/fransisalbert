// fransisalbert
//
// Copyright 2020 Aterier UEDA
// Author: Takeyuki UEDA

package fransisalbert


import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
 * getRequestBody
 */
func GetRequestBody(c *gin.Context) (body []byte) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	//		body := string(buf[0:n])
	body = buf[0:n]
	log.Println("body = ", string(body))
	return
}
