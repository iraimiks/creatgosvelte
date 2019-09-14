package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const (
	PORT = "8080"
)

func main() {
	router := getRouter()
	fmt.Println("starting server at : http://localhost:" + PORT)
	http.ListenAndServe(":"+PORT, router)
}

// не используйте такой код в прошакшене
// ошибка должна всегда явно обрабатываться
func __err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
