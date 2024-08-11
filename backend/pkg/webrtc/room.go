package webrtc

type Room struct {
	Peers *Peers
}

var Rooms map[string]*Room
