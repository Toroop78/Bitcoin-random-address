package main

import (
	"os"
    "time"
	"fmt"
	"math/big"
	"strings"
	"crypto/rand"
	"github.com/btcsuite/btcutil"
    "github.com/btcsuite/btcd/btcec"
    "github.com/btcsuite/btcd/chaincfg"
)

func generatePrivateKey() []byte {
	bytes := make([]byte, 32)
	if len(os.Args) < 2 {
		rand.Read(bytes)
	} else {
		arg := os.Args[1]
		startnum, ok := new(big.Int).SetString(arg,10)
		if ok != true { startnum, ok = new(big.Int).SetString(strings.TrimPrefix(arg,"0x"),16) }
		copy(bytes[32-len(startnum.Bytes()):], startnum.Bytes())
	}
	return bytes
}

func printAddressesForever() {
    padded := generatePrivateKey()
    count, one := new(big.Int).SetBytes(padded), big.NewInt(1)
    for {
	
		count.Add(count, one)
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

func main() {
    for i := 1; i < 180; i++ {
        go printAddressesForever()
    }
    for {
        time.Sleep(3000 * time.Millisecond)
    }
}