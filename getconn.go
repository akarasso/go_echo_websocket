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
        ids := strconv.FormatFloat(by.(float64), 'E', -1, 64)
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case int64:
        ids := strconv.FormatInt(by.(int64), 16)
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case int:
        ids := strconv.Itoa(by.(int))
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case string:
		if conn, exist := s.clients.byID[by.(string)]; exist {
			return conn, nil
		}
	}
    return nil, nil
}
