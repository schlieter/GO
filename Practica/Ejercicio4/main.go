package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	var buffer []string
	//////////SERVIDOR DE WEB SOCKET
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//sockets

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("nuevo usuario conectado")
		so.Join("chat_room")
		for _, msg := range buffer {
			so.Emit("chat mensaje", msg)
		}
		return nil
	})

	server.OnEvent("/", "chat mensaje", func(so socketio.Conn, msg string) {
		log.Println(msg)
		buffer = append(buffer, msg)
		server.BroadcastToRoom("/", "chat_room", "chat mensaje", msg)
	})

	go server.Serve()
	defer server.Close()

	//otros eventos

	///////////SERVIDOR DE HTTP

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server en el puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
