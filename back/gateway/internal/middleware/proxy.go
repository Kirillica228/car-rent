package middleware

import (
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

			// убираем дублирование путей
			proxyPath := c.Param("proxyPath")
			if !strings.HasPrefix(proxyPath, "/") {
				proxyPath = "/" + proxyPath
			}

			// правильно формируем финальный путь
			req.URL.Path = path.Join(remote.Path, proxyPath) 

			// пробрасываем заголовки
			req.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
