package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	//"github.com/umirode/golang-echo-socket.io"
	//"net/http"
)

//func serveSocket(c echo.Context) error {
//	server, err := socketio.NewServer(nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		s.SetContext("")
//		fmt.Println("connected:", s.ID())
//		return nil
//	})
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
//		fmt.Println("notice:", msg)
//		s.Emit("reply", "have "+msg)
//	})
//
//	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
//		s.SetContext(msg)
//		return "recv" + msg
//	})
//
//	server.OnEvent("/", "bye", func(s socketio.Conn) string {
//		last := s.Context().(string)
//		s.Emit("bye", last)
//		s.Close()
//		return last
//	})
//
//	//server.OnError("/", func(s socketio.Conn, e error) {
//	//	fmt.Println("meet error:", e)
//	//})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//	})
//
//	go server.Serve()
//	defer server.Close()
//	fmt.Println("served")
//
//	return nil
//}



func main() {

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "static/")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("test")
		s.Join("testRoom")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) string {
		fmt.Println("notice:", msg)

		server.BroadcastToRoom("testRoom", "reply", msg)
		return msg
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv" + msg
	})

	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		s.Close()
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	//e.GET("/", func(c echo.Context) error {
	//	return c.Render(http.StatusOK, "index", nil)
	//})

	e.GET("/", func(c echo.Context) error {
		return c.File("static/")
	})
	e.Any("/socket.io", echo.WrapHandler(server))

	e.Logger.Fatal(e.Start(":8060"))
}





//func socketIOWrapper() *golang_echo_socket_io.Wrapper {
//	wrapper, err := golang_echo_socket_io.NewWrapper(nil)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	wrapper.OnConnect("/", func(context echo.Context, conn socketio.Conn) error {
//		conn.SetContext("")
//		fmt.Println("connected:", conn.ID())
//		return nil
//	})
//	wrapper.OnError("/", func(context echo.Context, e error) {
//		fmt.Println("meet error:", e)
//	})
//	wrapper.OnDisconnect("/", func(context echo.Context, conn socketio.Conn, msg string) {
//		fmt.Println("closed", msg)
//	})
//
//	wrapper.OnEvent("/chat", "msg", func(context echo.Context, conn socketio.Conn, msg string) {
//		conn.SetContext(msg)
//		fmt.Println("notice:", msg)
//		conn.Emit("test", msg)
//	})
//
//	go wrapper.Serve()
//
//	return wrapper
//}
//
//func main() {
//
//	e := echo.New()
//
//	renderer := &TemplateRenderer{
//		template: template.Must(template.ParseGlob("static/*.html")),
//	}
//	e.Renderer = renderer
//
//	e.Pre(middleware.RemoveTrailingSlash())
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	e.GET("/", func(c echo.Context) error {
//		return c.Render(http.StatusOK, "index", nil)
//	})
//	e.Any("/socket.io", socketIOWrapper().HandlerFunc)
//
//	e.Logger.Fatal(e.Start(":8060"))
//}
