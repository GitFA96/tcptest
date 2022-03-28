package main

import (
	"io"
	"net"
	"testing"
)

func TestListener(t *testing.T) { //Listen er lyttende metode for server modifisering i GOlang
	listener, err := net.Listen("tcp", "127.0.0.1:8080") // åpner server på lokale pcn dedikert lokalt nettverk.
	if err != nil {                                      // håndterer evt error
		t.Fatal(err) // can endre err etter (listener, err -> listener, _ å droppe err mld)
	}

	/**
	Mtp på ports typ - 127.0.0.1: kan vi gjøre forskjellige ting. blant annet definiere port selv, :8080
	eller calle på adresse func. --- listener.adress)
	eller søke på nettstat --- i linux -> lsof ---

	sjekke dail (web test) -- kan se bort ifra --
	curl -v telnet://127.0.0.1:8080  /// mac/lix ----   nc  ---- (netcat) /// teste ip på nettleser
	*/

	defer listener.Close() // lukker porten

	for { // for løkke som skal behandle accept request fra client.
		conn, err := listener.Accept()
		if err != nil { // error behandling mtp avslutte løkke,
			t.Fatal(err)
		}

		go func() { // anonym funksjon, kan brukes som parameter/variabel
			io.Copy(conn, conn) // ekko sendes tilbake til client -- kopierer
			conn.Close()
		}() // syntaks () for å kjøre programmet av en anonymfunc blant annet.
	}

	t.Logf("bundet til %q", listener.Addr())
}

// server "listen" test.
// i powershell  kan ta med -v for å få med verbal(tekst)
// (lokal port) 127.0.0.1: (53299) kommer fra opperativsystemet nytt -- kan velge selv port man da må man lukke egen kode
// https://pkg.go.dev/net -- nett ressurser
