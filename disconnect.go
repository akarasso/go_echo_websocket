package ft_websocket

import (
    "github.com/gorilla/websocket"
)

func (s *Websocket_s)RemoveClient(ws *websocket.Conn) {
	if user, exist := s.clients.byConn[ws]; exist{
		user.Conn = filterConnection(user.Conn, ws)
		if len(user.Conn) == 0 {
            if s.Event.Auth.Disconnect != nil {
                s.Event.Auth.Disconnect(s, ws)
            }
		}
	} else {
        if s.Event.UnAuth.Disconnect != nil {
            s.Event.UnAuth.Disconnect(s, ws)
        }
    }
}
