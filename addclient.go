package ft_websocket

import (
    "github.com/gorilla/websocket"
    "errors"
    "strconv"
    // "log"
)

func (s *Websocket_s) AddClient(id interface{}, ws *websocket.Conn) error {
    switch id.(type) {
    case string:
        if user, exist := s.clients.byID[id.(string)]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[id.(string)] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
            n.ID = id.(string)
			n.Conn = append(n.Conn, ws)
			s.clients.byID[id.(string)] = n
			s.clients.byConn[ws] = n
		}
        return nil
    case int:
        ids := strconv.Itoa(id.(int))
        if user, exist := s.clients.byID[ids]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[ids] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.ID = ids
			n.Conn = append(n.Conn, ws)
			s.clients.byID[n.ID] = n
			s.clients.byConn[ws] = n
		}
        return nil
	case float64:
        ids := strconv.FormatFloat(id.(float64), 'E', -1, 64)
        if user, exist := s.clients.byID[ids]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[ids] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.ID = ids
			n.Conn = append(n.Conn, ws)
			s.clients.byID[n.ID] = n
			s.clients.byConn[ws] = n
		}
        return nil
    default:
        return errors.New("Fail to add client, index accepted string/float64/int")
	}
}
