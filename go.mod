module github.com/torusresearch/torus-common

go 1.13

require (
	filippo.io/edwards25519 v1.0.0
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2
	github.com/bwesterb/go-ristretto v1.2.0
	github.com/ethereum/go-ethereum v1.13.1
	github.com/prometheus/client_golang v1.12.0
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
	github.com/torusresearch/bijson v0.0.0-20200227065959-98da85656344
	golang.org/x/crypto v0.12.0
)

replace github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.23.2
