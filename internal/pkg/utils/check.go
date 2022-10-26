package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	if operatorID == viper.GetString("operator_id") && appSecret == viper.GetString("app_secret") {
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
