package infrastructure

import (
	"github.com/claustra01/sechack365/pkg/model"
	"golang.org/x/net/websocket"
)

type WsHandler struct {
	Ws []*websocket.Conn
}

func monitorConnection(ws *[]websocket.Conn) {
	for _, conn := range *ws {
		_, err := conn.Write([]byte("PING"))
		if err != nil {
			conn.Close()
		}
	}
}

func NewWsHandler(urls []string) (model.IWsHandler, error) {
	var ws []*websocket.Conn
	for _, url := range urls {
		conn, err := websocket.Dial(url, "", "http://localhost")
		if err != nil {
			for _, c := range ws {
				c.Close()
			}
			return nil, err
		}
		ws = append(ws, conn)
	}
	return &WsHandler{Ws: ws}, nil
}

func (ws *WsHandler) Send(msg string) error {
	for _, conn := range ws.Ws {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ws *WsHandler) Receive() (string, error) {
	var msg string
	for _, conn := range ws.Ws {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			return "", err
		}
		if msg != "" {
			return msg, nil
		}
	}
	return "", nil
}

func (ws *WsHandler) Close() error {
	for _, conn := range ws.Ws {
		conn.Close()
	}
	return nil
}
