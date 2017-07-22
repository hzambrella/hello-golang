package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func ParseTime(timeStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	return t, err
}

// 读取form值
func FormValue(req *http.Request, key string) string {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req.FormValue(key)
}

func DumpRequest(req *http.Request) {
	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
