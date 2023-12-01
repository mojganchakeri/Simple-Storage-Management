package internal

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(c *gin.Context, targetLocation string, targetPath string) {
	director := func(req *http.Request) {
		r := c.Request
		req.URL.Scheme = "http"
		req.URL.Host = targetLocation
		req.URL.Path = targetPath
		req.Header = r.Header
		req.Body = c.Request.Body
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
