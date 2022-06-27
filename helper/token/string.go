package token

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber() string {
	mrand.Seed(time.Now().Unix())
	max := big.NewInt(99999999)
	n, _ := rand.Int(rand.Reader, max)
	return strconv.Itoa(int(n.Int64()))
}
