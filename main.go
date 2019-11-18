package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func serveSocket(c echo.Context) error {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv" + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	//server.OnError("/", func(s socketio.Conn, e error) {
	//	fmt.Println("meet error:", e)
	//})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	fmt.Println("served")
	//defer server.Close()
	//
	return nil
}

func main() {

	e := echo.New()

	renderer := &TemplateRenderer{
		template: template.Must(template.ParseGlob("static/*.html")),
	}
	e.Renderer = renderer

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
	e.GET("/socket", serveSocket)

	e.Logger.Fatal(e.Start(":8060"))
}
