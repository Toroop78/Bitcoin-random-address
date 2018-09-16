package main

import (
	"fmt"
	"time"
	"math/big"
	"math/rand"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

const privatekeyConstant = "115792089237316195423570985008687907853269984665640564039457584007913129639936"

func main() {

    padded := make([]byte, 32)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxprivatekeys := new(big.Int)
	maxprivatekeys.SetString(privatekeyConstant, 10)
   
	for {
		count := new(big.Int)
		count.Rand(rng, maxprivatekeys)
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		privkey, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)
		wif, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, true)
		wifu, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, false)

		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		fmt.Println(caddr.EncodeAddress(), wif.String())
		fmt.Println(uaddr.EncodeAddress(), wifu.String())
	}
}	