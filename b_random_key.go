package main

import (
	"fmt"
	"log"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func generateKeyPair() (*btcec.PublicKey, *btcec.PrivateKey) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}
	return priv.PubKey(), priv
}

func main() {
   for {
	    public, privkey := generateKeyPair()
		wif, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, true)
        wifu, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, false)

		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		fmt.Println(caddr.EncodeAddress(), wif.String())
		fmt.Println(uaddr.EncodeAddress(), wifu.String())
	}
}