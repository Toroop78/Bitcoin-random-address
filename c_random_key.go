package main

import (
	"fmt"
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {

        padded := make([]byte, 32)
	
	for {
        
		rand.Read(padded)
		
		privkey, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)
		wif, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, true)
                wifu, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, false)

		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)
		
		fmt.Println(caddr.EncodeAddress(), wif.String())
		fmt.Println(uaddr.EncodeAddress(), wifu.String())	
	}
}
