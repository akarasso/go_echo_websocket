package ft_websocket

import (
    "github.com/gorilla/websocket"
    "errors"
)

func (s *Websocket_s)NewAction(key string, f func(*Websocket_s, *websocket.Conn, map[string]interface{})error) error {
    if _, exist := s.handler[key]; exist {
        return errors.New("Action '" + key + "' already exist")
    }
    s.handler[key] = f
    return nil
}
