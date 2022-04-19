package udptest

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestEchoServerUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	msg := []byte("ping")                    // array med 4 (8bites)
	_, err = client.WriteTo(msg, serverAddr) // kobling mellom client-> server med msg.
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf) // må ha med addr for å lese udp endepunkt --> sjekke at addr er riktig
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != serverAddr.String() {
		t.Fatalf("recived from %q instead of %q", addr, serverAddr)

	}
	if !bytes.Equal(msg, buf[:n]) {
		t.Errorf(" expeceted reply %q; actual replay %q", msg, buf[:n])

	}
}
