package core

import "testing"

func TestBridgeJS(t *testing.T) {
	t.Log(BridgeJS("http://localhost:4242"))
}
