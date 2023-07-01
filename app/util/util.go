package util

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/common"
)

func Echo(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"status":  common.STATUS_FAILURE,
			"message": err.Error(),
			"data":    data,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  common.STATUS_SUCCESS,
			"message": common.MESSAGE_SUCCESS,
			"data":    data,
		})
	}
}
