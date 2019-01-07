package ft_websocket

import (
    "github.com/gorilla/websocket"
    "strconv"
    "log"
)

func (s *Websocket_s) GetConn(by interface{}) (*s_client, error) {
	switch by.(type) {
    case *websocket.Conn:
        if conn, exist := s.clients.byConn[by.(*websocket.Conn)]; exist {
			return conn, nil
		}
	case float64:
        ids := strconv.Itoa(int(by.(float64)))
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case int64:
        ids := strconv.FormatInt(by.(int64), 10)
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
    default:
        log.Printf("Try to get conn with non accepted type:%T\n", by)
	}
    return nil, nil
}
