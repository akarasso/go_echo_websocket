package ft_websocket

import (
    "github.com/gorilla/websocket"
    "errors"
    "strconv"
    "log"
)

func (s *Websocket_s) AddClient(id interface{}, ws *websocket.Conn) error {
    switch id.(type) {
    case string:
        i, err := strconv.Atoi(id.(string))
        if err != nil {
            log.Println(err)
            return err
        }
        if user, exist := s.clients.byID[i]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[i] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
            n.ID = i
			n.Conn = append(n.Conn, ws)
			s.clients.byID[i] = n
			s.clients.byConn[ws] = n
		}
        return nil
    case int:
        if user, exist := s.clients.byID[id.(int)]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[id.(int)] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.ID = id.(int)
			n.Conn = append(n.Conn, ws)
			s.clients.byID[n.ID] = n
			s.clients.byConn[ws] = n
		}
        return nil
	case float64:
		if user, exist := s.clients.byID[int(id.(float64))]; exist {
			user.Conn = append(user.Conn, ws)
			s.clients.byID[int(id.(float64))] = user
			s.clients.byConn[ws] = user
		} else {
			n := new(s_client)
			n.ID = int(id.(float64))
			n.Conn = append(n.Conn, ws)
			s.clients.byID[n.ID] = n
			s.clients.byConn[ws] = n
		}
        return nil
    default:
        return errors.New("Fail to add client, index accepted string/float64/int")
	}
}
