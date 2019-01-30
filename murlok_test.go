package murlok

import (
	"context"
	"net/http"
	"testing"
)

func TestRunLocalServer(t *testing.T) {
	serv := &http.Server{Addr: ":0"}
	defer serv.Shutdown(context.Background())

	port, err := runLocalServer(serv)
	if err != nil {
		t.Fatal(err)
	}

	if port <= 0 {
		t.Fatal("bad port:", port)
	}

	t.Log("server port:", port)
}
