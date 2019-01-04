package ft_websocket

import (
    "github.com/gorilla/websocket"
    "strconv"
    "log"
)

func (s *Websocket_s) GetConn(by interface{}) (*s_client, error) {
	switch by.(type) {
    case *websocket.Conn:
        log.Println("Try to get conn by ws")
        if conn, exist := s.clients.byConn[by.(*websocket.Conn)]; exist {
			return conn, nil
		}
	case float64:
        ids := strconv.Itoa(int(by.(float64)))
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case int64:
        log.Println("Try to get conn by int64")
        ids := strconv.FormatInt(by.(int64), 16)
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case int:
        log.Println("Try to get conn by int")
        ids := strconv.Itoa(by.(int))
		if conn, exist := s.clients.byID[ids]; exist {
			return conn, nil
		}
	case string:
        log.Println("Try to get conn by string")
		if conn, exist := s.clients.byID[by.(string)]; exist {
			return conn, nil
		}
    default:
        log.Println("Try to get conn with non accepted type")
	}
    return nil, nil
}
