package webrtc

import "sync"

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	// TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	// PeerConnection *webrtc.PeerConnection
	Websocket *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	// Conn  *websocket.Conn
	Mutex sync.Mutex
}
