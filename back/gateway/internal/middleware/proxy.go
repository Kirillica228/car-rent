package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	remote, err := url.Parse(target)
	if err != nil {
		log.Fatalf("bad proxy target: %v", err)
	}

	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host

			proxyPath := c.Param("proxyPath")
			if !strings.HasPrefix(proxyPath, "/") {
				proxyPath = "/" + proxyPath
			}

			req.URL.Path = path.Join(remote.Path, proxyPath)

			req.Host = remote.Host

			if userID, exists := c.Get("user_id"); exists {
				req.Header.Set("X-User-ID", toString(userID))
			}
			if role, exists := c.Get("role"); exists {
				req.Header.Set("X-User-Role", toString(role))
			}
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(fmt.Sprintf("%v", v), "\n", ""))
}
