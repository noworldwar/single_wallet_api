package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckAppSecret(operatorID, appSecret string) bool {
	if operatorID == "14e5sfbg" && appSecret == "uLtBNRdp8qAlw9EHaFQSl_daPHmCl3zCCmmwfnb0ShE=" {
		return false
	} else {
		return true
	}
}

func HasPostFormEmpty(c *gin.Context, keys ...string) string {
	for _, v := range keys {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}
