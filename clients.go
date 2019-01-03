package ft_websocket

import (
    "github.com/gorilla/websocket"
)

type s_clients struct {
	byID	map[interface{}]*s_client
	byConn	map[*websocket.Conn]*s_client
}

type s_client struct {
	ID		interface{}
	Conn	[]*websocket.Conn
    log     bool
}
