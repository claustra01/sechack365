package infrastructure

import (
	"sync"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"golang.org/x/net/websocket"
)

type WsHandler struct {
	Ws   map[string]*websocket.Conn
	lock sync.Mutex
}

func NewWsHandler(urls []string, logger model.ILogger) (model.IWsHandler, error) {
	ws := make(map[string]*websocket.Conn)
	for _, url := range urls {
		conn, err := websocket.Dial(url, "", "http://localhost")
		if err != nil {
			for _, c := range ws {
				c.Close()
			}
			return nil, err
		}
		logger.Info("connected websocket to " + url)
		ws[url] = conn
	}
	wsHandler := &WsHandler{Ws: ws}
	wsHandler.monitor(logger)
	return wsHandler, nil
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
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for _, conn := range ws.Ws {
		conn.Close()
	}
	return nil
}

func (ws *WsHandler) monitor(logger model.ILogger) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			ws.lock.Lock()
			for url, conn := range ws.Ws {
				_, err := conn.Write([]byte("PING"))
				if err != nil {
					logger.Error("connection broken: " + err.Error())
					conn.Close()
					conn, err = reconnect(url, logger)
					if err != nil {
						logger.Error("reconnection error: " + err.Error())
					}
					ws.Ws[url] = conn
				}
			}
			ws.lock.Unlock()
		}
	}()
}

func reconnect(url string, logger model.ILogger) (*websocket.Conn, error) {
	conn, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		return nil, err
	}
	logger.Info("reconnected websocket to " + url)
	return conn, nil
}
