package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	letterByteBits = "ABCDEF0123456789"
	integerBytes   = "0123456789"
	letterOnlyByte = "ABCDEF"
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits  = 6                    // 6 bits to represent a letter index
	letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax   = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func getRandProxy() int {
	return int(rand.Int31n(int32(len(proxyList))))
}

func saveAccounts() {
	file, err := os.OpenFile("accounts.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Unable to open %s\n", err)
	}
	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ") //prevent the formatting from going out the window
	err = enc.Encode(&accounts)
	if err != nil {
		fmt.Printf("error encoding config file: %s\n", err)
		os.Exit(1)
	}
	file.Close()
	log.Println("Saved!")
}

func readAccounts() {
	f, err := os.Open("accounts.json")
	if err != nil {
		fmt.Printf("error opening accounts file: %s\n", err)
		os.Exit(1)
	}
	if err := json.NewDecoder(f).Decode(&accounts); err != nil {
		fmt.Printf("error decoding config file: %s\n", err)
		os.Exit(1)
	}
	f.Close()
	accLen = len(accounts)
}

func genRandomEmail(n int) string {
	return getRandString(n) + emailDomain
}

func genTotalDeviceToken() string {
	return genDeviceTokenBoth(8) + "-" + genDeviceTokenBoth(4) + "-" + genDeviceTokenInt(4) + "-" + genDeviceTokenInt(4) + "-" + genDeviceTokenBoth(12)
}

func genDeviceTokenAlpha(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterOnlyByte) {
			b[i] = letterOnlyByte[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func genDeviceTokenInt(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(integerBytes) {
			b[i] = integerBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func genDeviceTokenBoth(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterByteBits) {
			b[i] = letterByteBits[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func getRandString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
