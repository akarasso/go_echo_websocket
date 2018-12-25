package ft_websocket

import (
    "github.com/gorilla/websocket"
)

func (s *Websocket_s)RemoveClient(ws *websocket.Conn) {
	if user, exist := s.clients.byConn[ws]; exist{
		user.conn = filterConnection(user.conn, ws)
		if len(user.conn) == 0 {
            s.Event.Auth.Disconnect(s, ws)
		}
	} else {
        s.Event.UnAuth.Disconnect(s, ws)
    }
}
