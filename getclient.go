package ft_websocket

import (
    "github.com/gorilla/websocket"
    "strconv"
)

func (s *Websocket_s) GetConn(by interface{}) (*s_client, error) {
	switch by.(type) {
    case *websocket.Conn:
        if conn, exist := s.clients.byConn[by.(*websocket.Conn)]; exist {
			return conn, nil
		}
	case float64:
		if conn, exist := s.clients.byID[int(by.(float64))]; exist {
			return conn, nil
		}
	case int64:
		if conn, exist := s.clients.byID[int(by.(int64))]; exist {
			return conn, nil
		}
	case int:
		if conn, exist := s.clients.byID[by.(int)]; exist {
			return conn, nil
		}
	case string:
		x, err := strconv.Atoi(by.(string))
        if err != nil {
            return nil, err
        }
		if conn, exist := s.clients.byID[x]; exist {
			return conn, nil
		}
	}
    return nil, nil
}
