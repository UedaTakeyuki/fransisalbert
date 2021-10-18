// fransisalbert
//
// Copyright 2020 Aterier UEDA
// Author: Takeyuki UEDA

package fransisalbert

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

/*
 * getRequestBody
 */
func GetRequestBody(c *gin.Context) (body []byte) {
	//	buf := make([]byte, 4096)
	//	n, _ := c.Request.Body.Read(buf)
	//	body = buf[0:n]

	// https://stackoverflow.com/a/38874756/11073131
	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 65536))
	if err != nil {
		log.Println(err)
	}
	//	log.Println("body = ", string(body))
	return
}
