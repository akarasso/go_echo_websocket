package ft_websocket

import (
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
)

var (
	upgrader = websocket.Upgrader{}
)

func (s *Websocket_s)Handler(c echo.Context) error {
	var data map[string]interface{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	ws.SetCloseHandler(func (code int, text string) error {
		s.RemoveClient(ws)
		return nil
	})
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if err := json.Unmarshal(msg, &data); err != nil {
			continue
		}
		action, exist := data["action"]
		if !exist {
			log.Println("No action define")
			continue
		}
		switch action.(type) {
		case string:
			f, existHandler := s.handler[action.(string)]
			if !existHandler {
				log.Println("Try to execute unknow action '" + action.(string) + "'")
				continue
			}
			ret := f(s, ws, data)
			if ret != nil {
				return ret
			}
		default:
			log.Println("action must be a string")
		}
	}
	return nil
}
