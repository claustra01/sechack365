package infrastructure

import (
	"sync"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/gorilla/websocket"
)

type WsHandler struct {
	Ws   map[string]*websocket.Conn
	lock sync.Mutex
}

func NewWsHandler(relays []*model.NostrRelay, logger model.ILogger) (model.IWsHandler, error) {
	urls := make([]string, 0, len(relays))
	for _, r := range relays {
		urls = append(urls, r.Url)
	}
	ws := make(map[string]*websocket.Conn)
	for _, url := range urls {
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
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
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ws *WsHandler) Receive() (string, error) {
	var msg string
	for _, conn := range ws.Ws {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			return "", err
		}
		msg = string(msgBytes)
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
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			ws.lock.Lock()
			for url, conn := range ws.Ws {
				err := conn.WriteMessage(websocket.PingMessage, nil)
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
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	logger.Info("reconnected websocket to " + url)
	return conn, nil
}
