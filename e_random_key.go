package main

import (
	"fmt"
    "crypto/rand"
	"github.com/piotrnar/gocoin/lib/btc"
)

func main() {

    privateKey := make([]byte, 32)
    
	for {
	
	    rand.Read(privateKey)
	    publicKey := btc.PublicFromPrivate(privateKey, false)
	    address := btc.NewAddrFromPubkey(publicKey, 0x00).String()
	    fmt.Printf("%x %34s\n", privateKey, address)
	
	}
}	