package ft_websocket

import (
    "github.com/gorilla/websocket"
    "errors"
    "strconv"
)

func (s *Websocket_s) AddClient(id interface{}, ws *websocket.Conn) error {
    switch id.(type) {
    case string:
        if user, exist := s.clients.byID[id]; exist {
			user.conn = append(user.conn, ws)
            i, err := strconv.Atoi(id.(string))
            if err != nil {
                return err
            }
			s.clients.byID[i] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
            i, err := strconv.Atoi(id.(string))
            if err != nil {
                return err
            }
            n.id = i
			n.conn = append(n.conn, ws)
			s.clients.byID[i] = n
			s.clients.byConn[ws] = n
		}
        return nil
    case int:
        if user, exist := s.clients.byID[id]; exist {
			user.conn = append(user.conn, ws)
			s.clients.byID[id.(int)] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.id = id.(int)
			n.conn = append(n.conn, ws)
			s.clients.byID[n.id] = n
			s.clients.byConn[ws] = n
		}
        return nil
	case float64:
		if user, exist := s.clients.byID[id]; exist {
			user.conn = append(user.conn, ws)
			s.clients.byID[int(id.(float64))] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.id = int(id.(float64))
			n.conn = append(n.conn, ws)
			s.clients.byID[n.id] = n
			s.clients.byConn[ws] = n
		}
        return nil
    default:
        return errors.New("Fail to add client, index accepted string/float64/int")
	}
}
