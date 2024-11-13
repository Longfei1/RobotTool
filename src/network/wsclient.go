package network

import (
	"errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
	"time"
)

type IMsgHandle interface {
	HeartbeatInterval() time.Duration
	OnHeartbeat()
	HandleMessage(data []byte)
	OnConnect()
	OnClose()
}

type WsClient struct {
	wsConn *websocket.Conn
	send   chan []byte
	close  bool
	lock   sync.Mutex

	msgHandler IMsgHandle
}

func NewWsClient(heart IMsgHandle) *WsClient {
	return &WsClient{
		send:       make(chan []byte, 1000),
		msgHandler: heart,
	}
}

func (w *WsClient) Connect(url string) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.wsConn != nil {
		return errors.New("already connect")
	}

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	w.wsConn = ws

	go w.read()
	go w.write()

	if w.msgHandler != nil {
		w.msgHandler.OnConnect()
	}

	return nil
}

func (w *WsClient) SendData(data []byte) error {
	w.lock.Lock()
	if w.close {
		w.lock.Unlock()
		return errors.New("already close")
	}
	w.lock.Unlock()

	select {
	case w.send <- data:
		return nil
	default:
		log.Error("send chan full")
		return errors.New("send chan full")
	}
}

func (w *WsClient) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.close {
		return nil
	}

	w.close = true
	close(w.send)

	if w.msgHandler != nil {
		w.msgHandler.OnClose()
	}

	return nil
}

func (w *WsClient) read() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("read panic ", string(debug.Stack()))
		}
	}()

	defer w.Close()

	for {
		t, message, err := w.wsConn.ReadMessage()
		if err != nil {
			log.Info(err)
			return
		}

		if t == websocket.BinaryMessage {
			if w.msgHandler != nil {
				w.msgHandler.HandleMessage(message)
			}
		} else {
			log.Infof("receive message type: %v msg:%v", t, message)
		}
	}
}

func (w *WsClient) write() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("write panic ", string(debug.Stack()))
		}
	}()

	defer w.wsConn.Close()

	interval := time.Hour
	if w.msgHandler != nil {
		interval = w.msgHandler.HeartbeatInterval()
	}

	tick := time.NewTicker(interval)
	defer tick.Stop()

	for {
		select {
		case data, ok := <-w.send:
			if !ok {
				log.Info("send chan closed")
				return
			}
			err := w.wsConn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Error(err)
			}
		case <-tick.C:
			if w.msgHandler != nil {
				w.msgHandler.OnHeartbeat()
			}
		}
	}
}
