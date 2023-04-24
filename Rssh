
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gliderlabs/ssh"
	ssh2 "golang.org/x/crypto/ssh"
)

func ReadPrivateKeyFromFile(path string) (ssh.Signer, error) {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := ssh2.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func main() {

	// Create a new server instance
	s := &ssh.Server{
		Addr:            "127.0.0.1:2222",
		Handler:         handleSession,
		PasswordHandler: handleAuthentication,
		IdleTimeout:     60 * time.Second,
	}

	key, err := ReadPrivateKeyFromFile("key.rsa")
	if err != nil {
		panic(err)
	}
	s.AddHostKey(key)

	log.Printf("[+]Starting SSH server on address: %v\n", s.Addr)

	log.Fatal(s.ListenAndServe())

}

// Called if a new ssh session was created
func handleSession(s ssh.Session) {
	s.Write([]byte("Hello World!"))
	s.Close()
}

// Return true to accept password and false to deny
func handleAuthentication(ctx ssh.Context, passwd string) bool {

	if ctx.User() != "root" || passwd != "PA$$W0RD" {
                // Deny
		return false
	}

	fmt.Printf("User: %s,Password: %s, Address: %s", ctx.User(), passwd, ctx.RemoteAddr().String())

        // Accept
	return true

}
