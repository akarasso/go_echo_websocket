package ft_websocket

import (
    "github.com/gorilla/websocket"
)

type event_list_s struct {
    Disconnect func(*Websocket_s, string)
}

type event_s struct {
    Auth        event_list_s
    UnAuth      event_list_s
}

type Websocket_s struct {
    handler     map[string]func(*Websocket_s, *websocket.Conn, map[string]interface{})error
    clients     s_clients
    Event       event_s
}

var webServices map[string]Websocket_s = map[string]Websocket_s{}

func Get(name string) Websocket_s {
    return webServices[name];
}

func New(name string) Websocket_s {
    r := Websocket_s{
        handler : make(map[string]func(*Websocket_s, *websocket.Conn, map[string]interface{})error, 0),
        clients : s_clients{
            make(map[interface{}]*s_client),
            make(map[*websocket.Conn]*s_client),
        },
    }
    webServices[name] = r
    return r
}
