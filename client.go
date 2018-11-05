package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"
	"time"
)

const (
	//DES keys for packets. Correct as of 31/10/2018
	encryptKey = "E5QRecZA"
	decryptKey = "DvNw3mJT"

	//random stuff i found
	hostName           = "http://api54.luckydayapp.com" //54.85.149.45
	authDatabaseKey    = "lfSlYETq1A8H4TMXFb9U5nLu0Oi6LpDoflC7mx6yMl6ydHmvaG6sXi3ZsolbRDcM" //mmm juicy
	consumerKey        = "jW6ZOhMVPOEmb7PETWlLeEJ7M"
	consumerSecret     = "ny6Lo59VewKEHLliM3ZhgZN93h1CD46NSAFBx6NE20bFmf9zkA"
	databaseVersion    = 43
	fyberSecurityToken = "ea647275d1259a87145b9d95e4bbf647"
	segmentWriteKey    = "wql4awz243OfM5AW9g1QXZ7gw5IaJeiU"
	tenjinKey          = "AK72SXJLAU8DJVDQGMNY93RDF5SSA7NX"
	versionCode        = 115030007
	versionString      = "5.4.0"
	gameVersion        = "5.4"
	myToken            = "<REDACTED>"
	checkExists        = "Encrypt"

	emailDomain = "@gmail.com"

	//account based stuff
	accountAuth    = ".ASPXAUTH=<REDACTED>" //my device auth token
	deviceHeader1  = "LuckyDay/5.3 (iPod touch; iOS 10.4; Scale/2.00)"
	deviceHeader2  = "LuckyDay/5.3 (iPhone; iOS 11.4; Scale/3.00)"
	deviceVersion1 = "iPod Touch 6G"
	deviceVersion2 = "iPhone 6s Plus"
	menu           = "Loaded up %d accounts and using %d proxies\n" +
		"  1. Register accounts\n" +
		"  2. Login accounts\n" +
		"  3. Enter lucky code\n" +
		"  4. Play scratchers\n" +
		"  5. Logout accounts\n" +
		"  6. Cash out\n" +
		"  0. Exit App\n"
)

var (
	proxyList  []string
	randSrc    rand.Source
	accounts   []*StoredSession
	bots       []Bot
	accLen     int
	useProxies = false
)

func main() {
	randSrc = rand.NewSource(time.Now().UnixNano())
	readAccounts()
	readProxies()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go sigHandler(signals)

	fmt.Printf(menu, len(accounts), len(proxyList))
	fmt.Printf("Select an option from above...\n")
	var choice string
	fmt.Scanln(&choice)
	if choice == "1" {
		registerAccounts()
		return
	}
	for i := 0; i < len(accounts); i++ {
		b := Bot{}
		b.Account = accounts[i]
		b.Choice = choice
		if useProxies == true {
			go b.Start()
		} else {
			b.Start()
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func readProxies() {
	f, err := os.Open("list.proxies")
	if err != nil {
		fmt.Printf("error opening proxies file: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		proxyList = append(proxyList, strings.Split(scanner.Text(), "\n")[0])
	}
}

func sigHandler(c chan os.Signal) {
	signal := <-c
	_ = signal
	fmt.Printf("\nGot signal. Shutting down...\n")
	pprof.StopCPUProfile()
	saveAccounts()
	time.Sleep(200 * time.Millisecond)
	os.Exit(0)
}
