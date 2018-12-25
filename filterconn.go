package ft_websocket

import (
    "github.com/gorilla/websocket"

)

func filterConnection(list []*websocket.Conn, elem *websocket.Conn) []*websocket.Conn {
	ret := make([]*websocket.Conn, 0)
	for _, v := range list {
		if v != elem {
            ret = append(ret, v)
        }
	}
	return ret
}
