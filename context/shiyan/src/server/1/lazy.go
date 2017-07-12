package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const requestIDKey = "rid"

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	reqId := requestIDFromContext(req.Context())
	ranNum := rand.Intn(4)
	if ranNum <= 3 {
		time.Sleep(6 * time.Second) // simulate req hang.Client should cancle request.
		fmt.Fprintf(w, "slow response, %d  %v\n", ranNum, reqId)
		fmt.Printf("slow response, %d %v\n", ranNum, reqId)
		return
	}
	fmt.Fprintf(w, "quick response, %d %v\n", ranNum, reqId)
	fmt.Printf("quick response, %d %v\n", ranNum, reqId)
	return
}

func main() {
	http.Handle("/lazy", middleWare(http.HandlerFunc(lazyHandler)))
	fmt.Println("listen at :9200")
	http.ListenAndServe(":9200", nil)
}

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = "0"
	}
	return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

//use to get req id
func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
