package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/pkg/utils"
	"github.com/sirupsen/logrus"
)

func CheckHasPrefix(path string, vals ...string) bool {
	for _, v := range vals {
		if strings.HasPrefix(path, v) {
			return true
		}
	}

	return false
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		// Do not print log
		if CheckHasPrefix(path, "/doc", "/static", "/favicon.ico", "/test", "/live", "/ready") {
			c.Next()
			return
		}

		byteBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))
		body := string(byteBody)

		c.Next()

		// header log
		header := ""
		for k, v := range c.Request.Header {
			header += k + ": " + fmt.Sprint(v) + "\r\n"
		}

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		errorMsg := c.Keys["ErrorMsg"]

		fmt.Println("")

		logrus.WithField("HTTPResponse", LogResponse{
			Time:   time.Now().Format("2006/01/02 15:04:05"),
			IP:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   path,
			Header: header,
			Body:   body,
			Status: c.Writer.Status(),
			Error:  errorMsg,
		}).Println(c.Request.Method, path, c.Writer.Status())

		msg := "Time: " + time.Now().Format("2006/01/02 15:04:05") + "\n" +
			"IP: " + c.ClientIP() + "\n" +
			"Method: " + c.Request.Method + "\n" +
			"Path: " + path + "\n" +
			"Header: " + header + "\n" +
			"Body: " + body + "\n" +
			"Status: " + fmt.Sprintf("%d", c.Writer.Status()) + "\n\n"

		utils.WriteLog(msg)
	}
}

type LogResponse struct {
	Time   string
	IP     string
	Method string
	Path   string
	Header string
	Body   string
	Status int
	Params string
	Error  interface{} `json:"FailMsg,omitempty"`
}
