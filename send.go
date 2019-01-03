package ft_websocket

import (
	"log"
)

func (s *Websocket_s)Send(id interface {}, data interface{}) {
	user, err := s.GetConn(id)
	if err != nil {
		log.Println(err)
		return
	}
	if user != nil {
		for _, conn := range user.Conn {
			err := conn.WriteJSON(data)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
