package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {
	router := httprouter.New()
	router.GET("/", Logging(index, "index"))
	router.GET("/todos", Logging(TodoIndex, "todo-index"))
	router.GET("/todos/:todoId", Logging(TodoShow, "todo-show"))
	router.POST("/todos", Logging(TodoCreate, "todo-create"))
	router.DELETE("/todos/:todoId", Logging(TodoDelete, "todo-delete"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
