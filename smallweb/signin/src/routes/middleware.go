package routes

import (
	"fmt"
	"net/http"
	"strconv"
)

//参考: http://studygolang.com/articles/2500
//参考: https://github.com/gin-gonic/gin/blob/e31cbdf241b1ff161e3bc5eb4af1c9601fbb7639/logger.go
//中间件，可以不要
//打印请求的链接,方法，时间，以及状态码
//TODO :400 会打印两次
//TODO:不存在的链接服务端能提示
func ReqURLPrt(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		codestr1 := w.Header().Get("codelog")
		code, err := strconv.Atoi(codestr1)
		if err != nil {
			return
			//	logl.Error(errors.New("no status code"))
			//	panic(err)
		}
		var codeStr string
		var methodStr string
		switch {
		case 100 <= code && code < 300:
			codeStr = fmt.Sprintf("%c[1;42;37m%d%c[0m", 0x1B, code, 0x1B)
		case 300 <= code && code < 400:
			codeStr = fmt.Sprintf("%c[1;40;37m%d%c[0m", 0x1B, code, 0x1B)
		case 400 <= code && code < 500:
			codeStr = fmt.Sprintf("%c[1;43;37m%d%c[0m", 0x1B, code, 0x1B)
		case 500 <= code && code < 600:
			codeStr = fmt.Sprintf("%c[1;41;37m%d%c[0m", 0x1B, code, 0x1B)
		default:
			codeStr = fmt.Sprintf("%c[1;46;37m%d%c[0m", 0x1B, code, 0x1B)
		}

		switch r.Method {
		case "GET":
			methodStr = fmt.Sprintf(" %c[1;40;32m%s:%c[0m", 0x1B, r.Method, 0x1B)
		case "POST":
			methodStr = fmt.Sprintf(" %c[1;40;31m%s:%c[0m", 0x1B, r.Method, 0x1B)
		case "UPDATE":
			methodStr = fmt.Sprintf(" %c[1;40;34m%s:%c[0m", 0x1B, r.Method, 0x1B)
		case "DELETE":
			methodStr = fmt.Sprintf(" %c[1;40;33m%s:%c[0m", 0x1B, r.Method, 0x1B)
		default:
			methodStr = fmt.Sprintf(" %c[1;40;36m%s:%c[0m", 0x1B, r.Method, 0x1B)
		}

		logl.Info("request from", codeStr, fmt.Sprintf(" [%s]%s", methodStr, r.Host+r.RequestURI))
	})
}
