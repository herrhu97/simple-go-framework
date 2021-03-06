package main

import (
	"fmt"
	"gin"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "herrhu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gin.H{
			"title": "gin",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		names := []string{"herrhu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
