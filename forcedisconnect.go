package ft_websocket

import (
    "github.com/gorilla/websocket"
	"log"
)

func (s *Websocket_s)ForceDisconnect(ws *websocket.Conn) {
	if user, exist := s.clients.byConn[ws]; exist{
		delete(s.clients.byConn, ws)
		delete(s.clients.byID, user.ID)
	} else {
		log.Println("DC unexist user")
	}
}
