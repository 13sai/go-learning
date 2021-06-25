package main

import (
	"net/http"
)

func main()  {
	srv := http.Server{
		Addr: ":8088",
		Handler: http.HandlerFunc(defaultHttp),
	}
	srv.ListenAndServe()
}

// 默认http处理
func defaultHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello 13sai!"))
}