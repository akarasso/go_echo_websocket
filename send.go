package ft_websocket

import (
	"log"
)

func (s *Websocket_s)Send(id interface{}, data interface{}) {
	user, err := s.GetConn(id)
	if err != nil {
		log.Println(err)
		return
	}
	if user != nil {
		log.Println("Send message to ", id)
		for _, conn := range user.Conn {
			err := conn.WriteJSON(data)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Println("Send:", id, "user is dc")
	}
}
