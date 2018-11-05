package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//Bot defines our "player" and all of the session data
type Bot struct {
	Account             *StoredSession
	AWSALB              string //Supposedly "invalid" after 60 seconds, this is used for amazon load balancing but the idiots could use it for tracking too
	StoredPacket        string //The json payload to be sent in the next request
	Choice              string //The option that the dev took
	Tokens              float32
	Cash                float32
	AvailableGames      UpdateHomePageResponse //the last obtained game state
	AvailableScratchers map[int]int            //map[scratcherid]resultid
}

//StoredSession represents the data that is stored on disk (email, password, session cookie, ect)
type StoredSession struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Cookie   string `json:"Cookie"`
	MyDevice Device `json:"Device"`
}

//Start will grab a bots session cookie if it does not have one and perform the requested behavior
//todo: redo this so that its global and not tied to a bot
func (b *Bot) Start() {
	if b.Account.Cookie == "" {
		b.GetLoginCookie()
	}
	b.AvailableScratchers = make(map[int]int)
	switch b.Choice {
	case "2": //deprecated
	case "3": //enter lucky code
		b.EnterLuckyCode()
	case "4": //Play scratchers
		b.PlayScratchers()
	case "5": //Logout accounts
	case "6": //Cash out
	case "0": //Exit
		saveAccounts()
		os.Exit(0)
	default:
		saveAccounts()
		os.Exit(0)
	}
}

func (b *Bot) PlayScratchers() {
	b.GetGameBoard()
	b.MapScratchers()
	//fmt.Println("Available Scratchers:", b.AvailableScratchers)
	//fmt.Println("Game:", b.AvailableGames)
	idex := 0
	//todo: maybe check if our scratcher len is > 0 incase of errors
	for i := range b.AvailableScratchers {
		fmt.Printf("%s: Playing scratcher %d\n", b.Account.Email, idex+1)
		ps := b.playScratcherRequest(b.AvailableScratchers[i])
		//fmt.Println(ps)
		//fmt.Println([]byte(ps))
		eps := encryptPacket(ps)
		//fmt.Println(eps)
		//fmt.Println([]byte(eps))
		resp, _ := b.SendPostRequest(postScratcherPlay, eps, -1)
		if strings.Contains(resp, checkExists) == true {
			dp := decryptPacket(resp)
			bug := new(bytes.Buffer)
			bug.Write([]byte(dp))
			f := PlayScratcherResponse{}
			err := json.NewDecoder(bug).Decode(&f)
			if err != nil {
				fmt.Println("Error decoding scratcher response:", err)
			}
			//fmt.Println(dp)
			b.Cash += f.Wallet.CashUpdate
			b.Tokens += f.Wallet.ChipsUpdate
			fmt.Printf("Won %f tokens and %f cash\n", f.Wallet.ChipsUpdate, f.Wallet.CashUpdate)
		} else {
			fmt.Println(resp)
		}
		idex++
		time.Sleep(time.Second * 10)
	}
	fmt.Printf("New totals for %s:\n Tokens: %f\n Cash: %f\n", b.Account.Email, b.Tokens, b.Cash)
}

func (b *Bot) MapScratchers() {
	for i := 0; i < len(b.AvailableGames.GameSection.MyScratchers); i++ {
		b.AvailableScratchers[b.AvailableGames.GameSection.MyScratchers[i].ID] = b.AvailableGames.GameSection.MyScratchers[i].Result.ResultID
	}
}

func (b *Bot) GetGameBoard() {
	r := b.updateGameBoard()
	er := encryptPacket(r)
	resp, _ := b.SendPostRequest(postUpdateHomePage, er, -1)
	if strings.Contains(resp, checkExists) == true {
		bf := new(bytes.Buffer)
		x := UpdateHomePageResponse{}
		s := decryptPacket(resp)
		bf.Write([]byte(s))
		err := json.NewDecoder(bf).Decode(&x)
		if err != nil {
			fmt.Println("Error decoding:", err)
		}
		//todo: get types from decompiled apk
		b.Cash = x.Wallet.CurrentCash
		b.Tokens = x.Wallet.CurrentChips
		fmt.Printf("%s: Tokens: %f | Cash: %f\n", b.Account.Email, b.Tokens, b.Cash)
		b.AvailableGames = x
	} else {
		fmt.Println(resp)
	}

}

func (b *Bot) EnterLuckyCode() {
	var code string
	//fmt.Scanln(&code)
	// if globalCode == {
	// }
	//code = myCode
	rcode := strings.ToUpper(code)
	payload := b.enterLuckyCodeRequest(rcode)
	epayload := encryptPacket(payload)
	resp, httpr := b.SendPostRequest(postSaveLuckyCode, epayload, -1)
	if strings.Contains(resp, checkExists) == true {
		fmt.Println("Success!")
		_ = httpr
	} else {
		fmt.Println("Error:", resp)
	}
}

//GetLoginCookie retrieves the ASPX login session cookie from the server
func (b *Bot) GetLoginCookie() {
	payload := b.loginRequest()
	epayload := encryptPacket(payload)
	resp, httpr := b.SendPostRequest(postLogin, epayload, -1)
	if strings.Contains(resp, checkExists) == true {
		for i, v := range httpr.Header["Set-Cookie"] {
			fmt.Println(httpr.Header["Set-Cookie"][i])
			if strings.Contains(v, ".ASPXAUTH") {
				mess := httpr.Header["Set-Cookie"][i]
				idex := strings.Index(mess, ";")
				cookie := mess[:idex]
				b.Account.Cookie = cookie
			}
		}
		if b.Account.Cookie == "" {
			fmt.Println("Was not able to collect session cookie:", b.Account.Email)
		}
	} else {
		fmt.Println(resp)
	}
}

//RegisterAccounts makes the supplied number of accounts
func RegisterAccounts(amt int) {
	for i := 0; i < amt; i++ {
		b := Bot{}
		email := genRandomEmail(10)
		password := "SomePassword123"
		devToken := genTotalDeviceToken()
		first := getRandString(10)
		last := getRandString(10)
		devType := deviceVersion2
		payload := b.registerRequest(email, password, devType, devToken, first, last)
		epayload := encryptPacket(payload)
		res, _ := b.SendPostRequest(postRegisterAccount, epayload, -1)
		if strings.Contains(res, checkExists) {
			me := StoredSession{}
			me.Email = email
			me.Password = password
			device := Device{}
			device.DeviceToken = devToken
			device.DeviceVersion = devType
			device.OperatingSys = 0
			me.MyDevice = device
			accounts = append(accounts, &me)
		} else {
			fmt.Println("Error:", res)
		}
		time.Sleep(350 * time.Millisecond)
	}
	saveAccounts()
}

//SendRequest is prototype for sending out requests
func (b *Bot) SendPostRequest(head string, p string, pIndex int) (string, *http.Response) {
	bug := bytes.NewBuffer([]byte(p))
	r, err := http.NewRequest(http.MethodPost, hostName+head, bug)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	r.Header.Add("User-Agent", deviceHeader2)
	r.Header.Add("api-version", "3") //no way to get this to be lowercase which is what the app does
	r.Header.Add("Connection", "close")
	r.Header.Set("api-version", "3")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Proxy-Connection", "keep-alive")
	r.Header.Add("Accept", "*/*")
	r.Header.Add("Content-Type", "application/json") //server will freak out without this
	r.Header.Add("Accept-Language", "en-US;q=1, zh-Hant-US;q=0.9, ja-US;q=0.8, ko-US;q=0.7")
	if b.Account.Cookie != "" { //pretty much all urls require this except for logging in
		r.Header.Add("Cookie", b.Account.Cookie)
	}
	if useProxies == true {
		var err error
		proxyURL, err := url.Parse(proxyList[getRandProxy()])
		if err != nil {
			fmt.Println("Error forming a proxy:", err)
		}
		myClient := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
		// myClient := &http.Client{}
		// trans := &http.Transport{}
		// trans.Proxy = http.ProxyURL(proxyURL)
		// myClient.Transport = trans
		ex, err := myClient.Do(r)
		if err != nil {
			fmt.Println(err)
		}
		body, _ := ioutil.ReadAll(ex.Body) //don't think we'll error if we got to here
		ex.Body.Close()
		return string(body), ex
	}
	h := new(http.Client)
	resp, err := h.Do(r)
	if err != nil {
		fmt.Println("Error in response:", err)
		return "", resp
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ioutil:", err)
	}
	resp.Body.Close()
	return string(body), resp
}

func registerAccounts() {
	for i := 0; i < 1; i++ {
		p := RegisterRequest{}
		p.Email = strings.ToLower(genRandomEmail(10))
		p.UserDevice.DeviceToken = genTotalDeviceToken()
		p.UserDevice.DeviceVersion = deviceVersion2
		p.UserDevice.OperatingSys = 0
		p.FirstName = getRandString(10)
		p.LastName = getRandString(10)
		p.Password = "SomePassword123"

		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		err := enc.Encode(&p)
		if err != nil {
			fmt.Printf("error encoding config file: %s\n", err)
			os.Exit(1)
		}
		payload := encryptPacket(string(buf.Bytes()))
		fmt.Println("Payload:", payload)
		h := new(http.Client)
		_ = h
		bug := bytes.NewBuffer([]byte(payload))
		r, err := http.NewRequest(http.MethodPost, hostName+postRegisterAccount, bug)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
		r.Header.Add("User-Agent", deviceHeader2)
		r.Header.Add("api-version", "3")
		r.Header.Add("Connection", "close")
		r.Header.Set("api-version", "3")
		r.Header.Add("Accept-Encoding", "gzip, deflate")
		r.Header.Add("Proxy-Connection", "keep-alive")
		r.Header.Add("Accept", "*/*")
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Accept-Language", "en-US;q=1, zh-Hant-US;q=0.9, ja-US;q=0.8, ko-US;q=0.7")

		resp, err := h.Do(r)
		if err != nil {
			fmt.Println("Error in response:", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ioutil:", err)
		}
		resp.Body.Close()
		res := string(body)
		if strings.Contains(res, "Encrypt") {
			me := StoredSession{}
			me.Email = p.Email
			me.Password = p.Password
			me.MyDevice = p.UserDevice
			accounts = append(accounts, &me)
		} else {
			fmt.Println("Error:", res)
		}
		fmt.Println(res)
		time.Sleep(1000 * time.Millisecond)
	}
	saveAccounts()
}
