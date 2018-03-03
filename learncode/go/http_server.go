package main

import (
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello\n"))
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("handle\n"))
}

func readwrite(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	n := req.FormValue("times")
	times, _ := strconv.Atoi(n)

	re := ""
	for i := 0; i < times; i++ {
		re += "OK,"
	}
	re += "\n"
	w.Write([]byte(re))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/times/", readwrite)
	http.Handle("/handle", http.HandlerFunc(handle))
	http.ListenAndServe(":8001", nil)
	select {} //阻塞进程
}
