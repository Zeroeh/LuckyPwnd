package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//send the encrypted packets payload here
func decryptPacket(d string) string {
	x := Encrypted{}
	b := new(bytes.Buffer)
	b.Write([]byte(d))
	json.NewDecoder(b).Decode(&x)
	daa, _ := base64.StdEncoding.DecodeString(x.Payload)
	da := string(daa)
	db := decrypt([]byte(decryptKey), []byte(decryptKey), []byte(da))
	return db
}

func encryptPacket(d string) string {
	da := encrypt([]byte(encryptKey), []byte(encryptKey), []byte(d))
	db := base64.StdEncoding.EncodeToString(da)
	//kids, stay away from IIS servers
	dc := padNewLines(db)
	dd := padSlashes(dc)
	de := new(bytes.Buffer)
	//fmt.Println([]byte(dd))
	fmt.Fprintf(de, "{\"Encrypt\":\"%s\"}", dd[:len(dd)-4]) //weird hack fix... overall this is scuffed and should be redone at some point
	return string(de.Bytes())
}

func padNewLines(s string) string {
	ls := s
	strlen := len(ls)
	counts := (strlen / 64)
	bugs := make([]byte, strlen+(counts*4))
	offset := 0
	for i := 0; i < strlen; i++ {
		if i%64 == 0 && i != 0 {
			bugs[i+offset] = byte("\\"[0])
			bugs[i+1+offset] = byte("r"[0])
			bugs[i+2+offset] = byte("\\"[0])
			bugs[i+3+offset] = byte("n"[0])
			bugs[i+4+offset] = ls[i]
			offset += 4
		} else {
			bugs[i+offset] = ls[i]
		}
	}
	return string(bugs)
}

func padSlashes(s string) string {
	return strings.Replace(s, "/", "\\/", -1)
}

func padString(s string) string {
	ls := s
	//next we will count 64 characters in the string and insert "\r\n"
	strlen := len(ls)
	counts := (strlen / 64) + 1
	fmt.Println(counts)
	bugs := make([]byte, strlen+(counts*4))
	fmt.Println("Old | New", len(ls), len(bugs))
	for i := 0; i < strlen; i++ {
		if i%64 == 0 {
			if i == 0 {
				bugs[i] = ls[i]
			} else {
				//bugs[i] = ls[i]
				bugs[i] = byte("\\"[0])
				bugs[i+1] = byte("r"[0])
				bugs[i+2] = byte("\\"[0])
				bugs[i+3] = byte("n"[0])
				//i += 3
			}

		} else {
			bugs[i] = ls[i]
		}

	}
	bigger := string(bugs)
	bigString := strings.Replace(bigger, "/", "\\/", -1)
	fmt.Println(bigString)
	return bigString
}

//Encrypted is the encrypted json payload
type Encrypted struct {
	Payload string `json:"Encrypt"`
}

type RegisterRequest struct {
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	UserDevice Device `json:"Device"`
}

type RegisterResponse struct {
	Created string  `json:"CreatedDate"`
	Error   *string `json:"Error"`
	Message *string `json:"Message"`
}

type Device struct {
	DeviceToken   string `json:"DeviceToken"`
	OperatingSys  int    `json:"OperatingSystem"`
	DeviceVersion string `json:"DeviceVersion"`
}

type Location struct {
	City    string `json:"City"`
	Country string `json:"Country"`
}

type LoginRequest struct {
	Email        string   `json:"Email"`
	Password     string   `json:"Password"`
	GameVersion  string   `json:"Version"`
	MyLocation   Location `json:"Location"`
	MyDevice     Device   `json:"Device"`
	Notification string   `json:"NotificationToken"`
}

type PlayScratcherRequest struct {
	IsSpecialScratcher bool `json:"IsSpecialScratcher"`
	ResultID           int  `json:"ResultId"`
	AreaPlayed         int  `json:"AreaPlayed"`
	VariantDailyReward int  `json:"VariantDailyReward"`
}

type PlayScratcherResponse struct {
	NeedUpdate bool              `json:"IsNeedToUpdate"`
	Model      DailyBonusModel   `json:"DailyBonusModel"`
	CheckPoint ScratchCheckPoint `json:"ScratchCheckpoint"`
	Card       CardInfo          `json:"CardInfo"`
	Wallet     ActualWallet      `json:"ActualWallet"`
	Error      *string           `json:"Error"`
	Message    *string           `json:"Message"`
}

type DailyBonusModel struct {
	CanPlayBonusGame bool `json:"CanPlayBonusGame"`
	Bonus            int  `json:"Bonus"`
	DaysInApp        int  `json:"DaysInApp"`
	Coeff            int  `json:"Coefficient"`
}

//LoginResponse is returned after sending a loginrequest (successfully)
// assume for now error field is populated if bad login. Set user cookie on this response
type LoginResponse struct {
	Error         *string  `json:"Error"`
	ValidEmail    bool     `json:"IsEmailValid"`
	TutorialState Tutorial `json:"Tutorial"`
}

//UpdateHomePageRequest is first sent when reloading the app while we actually have a session cookie
type UpdateHomePageRequest struct {
	OnBoardVariantScratcher   int `json:"OnBoardingVariantScratcher"`
	OnBoardVariantDailyReward int `json:"OnBoardingVariantDailyReward"`
}

//UpdateHomePageResponse BIG BOI
type UpdateHomePageResponse struct {
	TimeToNextDay     int64            `json:"TimeTillNextDay"`
	GameSection       GameSectionModel `json:"GamesSectionModel"`
	Wallet            ActualWallet     `json:"ActualWallet"`
	Error             *string          `json:"Error"`
	ServerDate        string           `json:"ServerDate"`
	Cards             CardInfo         `json:"CardInfo"`
	ExpireTimeSeconds int              `json:"ExpiresTime"`
	ReferralModel     interface{}      `json:"ReferralModel"`
	MyData            UserData         `json:"UserData"`
	ShowWelcomeBack   bool             `json:"ShowWelcomeBackMenu"`
}

type UserData struct {
	ID         int     `json:"Id"`
	FirstName  string  `json:"FirstName"`
	LastName   string  `json:"LastName"`
	AvatarURL  *string `json:"AvatarUrl"`
	PromoCode  string  `json:"PromoCode"`
	CreatedOn  string  `json:"CreatedDate"`
	EmailValid bool    `json:"IsEmailValid"`
	Token      string  `json:"Token"`
}

type CardInfo struct {
	CardsPlayed    int    `json:"CardsPlayed"`
	LastCardPlayed string `json:"LastCardPlayed"`
}

type RedeemRewardRequest struct {
	RewardID int    `json:"RewardId"`
	Email    string `json:"Email"`
}

type RedeemRewardResponse struct {
	OrderID *int         `json:"OrderId"` //assuming int
	Wallet  ActualWallet `json:"ActualWallet"`
	Error   ErrorPacket  `json:"Error"`
	Message *string      `json:"Message"`
}

type ActualWallet struct {
	CurrentChips        float32     `json:"WinChips"`
	CurrentCash         float32     `json:"CashWallet"`
	ChipsUpdate         float32     `json:"ChipsUpdate"`
	CashUpdate          float32     `json:"CashUpdate"`
	WalletFlowStatement interface{} `json:"WalletFlowStatement"`
	WalletType          interface{} `json:"WalletType"`
}

//GameSectionModel is the state of all the games
type GameSectionModel struct {
	MyCashCars       CashCarsScratcher  `json:"CashCarsScratchers"`
	MyRaffles        []Raffles          `json:"Raffles"`
	MyScratchers     []Scratchers       `json:"AllScratchers"`
	MyLotto          Lotto              `json:"Lotto"`
	SpecialScratcher []SpecialScratcher `json:"SpecialScratcher"`
	CheckPoint       ScratchCheckPoint  `json:"ScratchCheckpoint"`
}

type SpecialScratcher struct {
	Expiration int64 `json:"ExpirationTime"`
}

type ScratchCheckPoint struct {
	PlayedCards   int         `json:"PlayedCards"`
	RequiredCards int         `json:"RequiredCards"`
	Prize         interface{} `json:"Prize"`
}

type Lotto struct {
	RevealAvailable   bool `json:"IsRevealAvailable"`
	TimeToNextDrawing int  `json:"TimeTillNextDrawing"`
	LottoAvailable    bool `json:"IsLottoAvailable"`
}

type Scratchers struct {
	ID             int        `json:"Id"`
	Name           string     `json:"Name"`
	PrizeType      int        `json:"PrizeType"`
	PrizeAmount    float32    `json:"PrizeAmount"`
	BonusPrizeType int        `json:"BonusPrizeType"`
	ImageType      int        `json:"TypeOfImage"`
	Result         GameResult `json:"GameResult"`
	ScratchOrder   int        `json:"ScratchOrderType"`
	Error          *string    `json:"Error"`
}

//Raffles
type Raffles struct {
	GameID              int     `json:"GameId"`
	Title               string  `json:"Title"`
	Price               int     `json:"Price"`
	PrizeAmount         float32 `json:"PrizeAmount"`
	PrizeType           int     `json:"PrizeType"`
	Description         string  `json:"Descriptions"`
	TimeToNextGame      int     `json:"TimeTillNextGame"`
	UserTicketsCount    int     `json:"UserTicketsCount"`
	TotalTicketsCount   int     `json:"TotalTicketsCount"`
	RaffleType          int     `json:"RaffleType"`
	IsWinner            bool    `json:"IsWinner"`
	FreeTicketAvailable bool    `json:"IsAvailableFreeTicket"`
}

type RaffleTicketOddsResponse struct {
	Amount int     `json:"Amount"`
	Error  *string `json:"Error"`
}

type CashCarsScratcher struct {
	DaysLeft  int           `json:"DaysLeft"`
	IsLocked  bool          `json:"IsLocked"`
	ImageType int           `json:"TitleImageType"`
	Results   []GameResults `json:"GameResults"`
}

type GameResults struct {
	GameID           int        `json:"Id"`
	Name             string     `json:"Name"`
	PrizeType        int        `json:"PrizeType"`
	PrizeAmount      float32    `json:"PrizeAmount"`
	BonusPrizeType   int        `json:"BonusPrizeType"`
	TypeOfImage      int        `json:"TypeOfImage"`
	Result           GameResult `json:"GameResult"`
	ScratchOrderType int        `json:"ScratchOrderType"`
	Error            *string    `json:"Error"`
}

type GameResult struct {
	Rows      []int   `json:"Rows"`
	Bonus     float32 `json:"Bonus"`
	BonusType *int    `json:"BonusType"`
	ResultID  int     `json:"ResultId"`
}

//IsRateDisplayedResponse is sent after sending IsRateDisplayed http request. No request version exists
type IsRateDisplayedResponse struct {
	RateDisplayed bool    `json:"IsRateDisplayed"`
	Error         *string `json:"Error"`
}

type Tutorial struct {
	ShowBlackJackTut            bool `json:"ShowBlackJackTutorial"`
	ShowScratcherTut            bool `json:"ShowScratcherTutorial"`
	ShowChallengeTut            bool `json:"ShowChallengeTutorial"`
	ShowNewUserTut              bool `json:"ShowNewUserTutorial"`
	ShowInstantTut              bool `json:"ShowInstantTutorial"`
	ShowRaffleTut               bool `json:"ShowRaffleTutorial"`
	NeedsToPlayWelcomeScratcher bool `json:"IsNeedToPlayWelcomeScratcher"`
}

type SaveLuckyCodeRequest struct {
	Code string `json:"LuckyCode"`
}

type SaveLuckyCodeResponse struct {
	Error   *string `json:"Error"`
	Message *string `json:"Message"`
}

type StartLottoGameRequest struct {
	VariantDailyReward int   `json:"VariantDailyReward"`
	Numbers            []int `json:"Numbers"`
	LuckyNumber        int   `json:"LuckyNumber"`
}

type StartLottoGameResponse struct {
	NeedUpdate bool            `json:"IsNeedToUpdate"`
	Card       CardInfo        `json:"CardInfo"`
	BonusModel DailyBonusModel `json:"DailyBonusModel"`
	Wallet     ActualWallet    `json:"ActualWallet"`
	Error      *string         `json:"Error"`
}

type LottoDetailsResponse struct {
	Jackpot           float32             `json:"JackPot"`
	TimeToNextDrawing int                 `json:"TimeTillNextDrawing"`
	WinInfo           interface{}         `json:"WinInfo"`
	WinningNumbers    interface{}         `json:"WinningNumbers"` //probably []int
	PreviousNumbers   PreviousGameNumbers `json:"PreviousGameNumbers"`
}

type LottoDetailsRequest struct {
	GetPreviousNumbers bool `json:"GetPreviousGameNumber"`
}

type PreviousGameNumbers struct {
	ID          int     `json:"Id"`
	Numbers     []int   `json:"Numbers"`
	LuckyNumber int     `json:"LuckyNumber"`
	UserNumbers []int   `json:"UserNumbers"`
	Description *string `json:"DescriptionTimeOfReveal"`
}

type ErrorPacket struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
	Show      bool   `json:"Show"`
}
