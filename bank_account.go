package stripe

import "encoding/json"


type BankAccountParams struct {
	Params
	RoutingNumber string
	AccountNumber string
	CountryCode   string
}

type BankAccountStatus string

type BankAccount struct {
	ID          string            `json:"id"`
	Name        string            `json:"bank_name"`
	Last4       string            `json:"last4"`
	Country     string            `json:"country"`
	Currency    string            `json:"currency"`
	FingerPrint string            `json:"fingerprint"`
	Customer    string            `json:"customer"`
	Status      BankAccountStatus `json:"status"`
	Disabled    bool              `json:"disabled"`
	Validated   bool              `json:"validated"`
	Verified    bool              `json:"verified"`
}

func (b *BankAccount) UnmarshalJSON(data []byte) error {
	type bankaccount BankAccount
	var ba bankaccount
	err := json.Unmarshal(data, &ba)
	if err == nil {
		*b = BankAccount(ba)
	} else {
		b.ID = string(data[1 : len(data)-1])
	}
	return nil
}
