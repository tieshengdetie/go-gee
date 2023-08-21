package main

import (
	"gee"
	"net/http"
)

func main() {

	router := gee.New()

	router.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1> this is index</h1>")
	})
	router.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	router.POST("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	router.Run(":9999")
}
