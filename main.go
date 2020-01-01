package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "")
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("username中的内容为：", username, password)
	if username == "123" && password == "456" {
		t := template.Must(template.ParseFiles("index.html"))
		// 执行
		t.Execute(w, "账号密码正确！")
	} else {
		t := template.Must(template.ParseFiles("index.html"))
		// 执行
		t.Execute(w, "账号密码错误！")
	}
}

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", nil)
}
