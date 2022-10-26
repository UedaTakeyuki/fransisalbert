package fransisalbert

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func OkNgErrorOutWithMessage(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"message": errors.New(message),
	})
	//	c.Abort()
	return
}

func OkNgErrorOut(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"message": err,
	})
	//	c.Abort()
	return
}

func OkNgNotOwnerReturn() (err error) {
	err = errors.New("not owner")
	return
}
