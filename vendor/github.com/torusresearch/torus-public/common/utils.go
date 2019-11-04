package common

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func BigIntToPoint(x, y *big.Int) Point {
	return Point{X: *x, Y: *y}
}

func HexToBigInt(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

// converts common point to ETH address format borrowint ethcrypto functions
func PointToEthAddress(point Point) *common.Address {
	addr := crypto.PubkeyToAddress(*PointToECDSAPublicKey(point))
	return &addr
}

func PointToECDSAPublicKey(point Point) *ecdsa.PublicKey {
	return &ecdsa.PublicKey{
		Curve: crypto.S256(),
		X:     &point.X,
		Y:     &point.Y,
	}
}

func BigIntToECDSAPrivateKey(x big.Int) *ecdsa.PrivateKey {
	return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: crypto.S256(),
		},
		D: &x,
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func StringChanWriteTimeout(channel chan<- string, msg string, timeout time.Duration) error {
	timeoutCh := make(chan bool)
	go func() {
		time.Sleep(timeout)
		<-timeoutCh
	}()
	select {
	case channel <- msg:
		return nil
	case timeoutCh <- true:
		return fmt.Errorf("Channel %v with msg %v timed out after %v", channel, msg, timeout)
	}
}

func IntChanWriteTimeout(channel chan<- int, msg int, timeout time.Duration) error {
	timeoutCh := make(chan bool)
	go func() {
		time.Sleep(timeout)
		<-timeoutCh
	}()
	select {
	case channel <- msg:
		return nil
	case timeoutCh <- true:
		return fmt.Errorf("Channel %v with msg %v timed out after %v", channel, msg, timeout)
	}
}
