package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"lucas": {
		AuthToken: "50304",
		Username:  "lucas",
	},
	"zilas": {
		AuthToken: "00107",
		Username:  "zilas",
	},
	"sophia": {
		AuthToken: "67609",
		Username:  "sophia",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"lucas": {
		Coins:    100,
		Username: "lucas",
	},
	"zilas": {
		Coins:    200,
		Username: "zilas",
	},
	"sophia": {
		Coins:    500,
		Username: "sophia",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
