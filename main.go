package main

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type Data struct {
	Room    string
	Message string
}

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
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "room_in", func(s socketio.Conn, room string) {
		s.Join(room)
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, d string) string {
		var data Data
		err := json.Unmarshal([]byte(d), &data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("received message: ", data.Message, " from room: ", data.Room)

		server.BroadcastToRoom(data.Room, "reply", data.Message)
		return data.Message
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

	e.GET("/", func(c echo.Context) error {
		return c.File("static/")
	})
	e.Any("/socket.io", echo.WrapHandler(server))

	e.Logger.Fatal(e.StartTLS(":8060", "cert.pem", "key.pem"))
}
