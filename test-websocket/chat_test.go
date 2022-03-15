package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func hendler() http.HandlerFunc {
	hub := newHub()
	go hub.run()
	return  func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}
}

func newWSServer(t *testing.T) (*httptest.Server, *websocket.Conn) {
	server := httptest.NewServer(hendler())

	wsURL, err := url.Parse(server.URL)
	assert.NoError(t, err)
	wsURL.Scheme = "ws"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	assert.NoError(t, err)

	return server, ws
}

func TestChat(t *testing.T) {
	server, ws := newWSServer(t)
	defer server.Close()
	defer ws.Close()

	message := "test message"

	err := ws.WriteMessage(websocket.BinaryMessage, []byte(message))
	assert.NoError(t, err)

	_, readMessage, err := ws.ReadMessage()
	fmt.Println(string(readMessage))
	assert.NoError(t, err)

	assert.Equal(t, message, string(readMessage))
}