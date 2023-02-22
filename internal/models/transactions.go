package models

type Rates struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		PublicKey string `json:"publicKey"`
		Signature string `json:"signature"`
		Rates     struct {
			SellUSD struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN"`
			BuyUSD struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN_"`
		} `json:"rates"`
	} `json:"data"`
}
