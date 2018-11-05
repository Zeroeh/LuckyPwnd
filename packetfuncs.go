package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func (b *Bot) updateGameBoard() string {
	r := UpdateHomePageRequest{}
	r.OnBoardVariantDailyReward = 0
	r.OnBoardVariantScratcher = 1
	bb := new(bytes.Buffer)
	json.NewEncoder(bb).Encode(&r)
	return string(bb.Bytes())
}

//the id is the "resultid" obtained from the big boi packet (updatehomepageresponse)
func (b *Bot) playScratcherRequest(resultid int) string {
	ps := PlayScratcherRequest{}
	ps.AreaPlayed = 1
	ps.IsSpecialScratcher = false
	ps.VariantDailyReward = 0
	ps.ResultID = resultid
	bb := new(bytes.Buffer)
	json.NewEncoder(bb).Encode(&ps)
	return string(bb.Bytes()[:bb.Len()-1]) //some sort of hidden line feed byte at the end causes it to bug out (ascii decimal code 10)
}

func (b *Bot) enterLuckyCodeRequest(code string) string {
	lc := SaveLuckyCodeRequest{}
	lc.Code = code
	bb := new(bytes.Buffer)
	json.NewEncoder(bb).Encode(&lc)
	return string(bb.Bytes())
}

func (b *Bot) registerRequest(email, password, devicev, token, first, last string) string {
	p := RegisterRequest{}
	p.Email = email
	p.UserDevice.DeviceToken = genTotalDeviceToken()
	p.UserDevice.DeviceVersion = "iPhone 6s Plus"
	p.UserDevice.OperatingSys = 0
	p.FirstName = getRandString(10)
	p.LastName = getRandString(10)
	p.Password = "Password123"
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(&p)
	if err != nil {
		fmt.Printf("error encoding config file: %s\n", err)
		os.Exit(1)
	}
	return string(buf.Bytes())
}

func (b *Bot) loginRequest() string {
	lr := LoginRequest{}
	lr.Email = b.Account.Email
	lr.Password = b.Account.Password
	lr.MyDevice = b.Account.MyDevice
	loc := Location{}
	loc.City = ""
	loc.Country = ""
	lr.MyLocation = loc
	lr.GameVersion = gameVersion
	lr.Notification = ""
	bb := new(bytes.Buffer)
	json.NewEncoder(bb).Encode(&lr)
	return string(bb.Bytes())
}
