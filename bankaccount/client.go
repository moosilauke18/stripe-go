package bankaccount

import (
	"net/url"

	stripe "github.com/moosilauke18/stripe-go"
)

func New(customer_id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().New(customer_id, params)
}
type Client struct {
	B stripe.Backend
	Key string
}

func (c Client) New(customer_id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	var body *url.Values
	var commonParams *stripe.Params
	if params != nil {
		body = &url.Values{
			"bank_account[routing_number]": {params.RoutingNumber},
			"bank_account[account_number]": {params.AccountNumber},
			"bank_account[country_code]": {params.CountryCode},
		}

	}
	bankAccount := &stripe.BankAccount{}
	commonParams = &params.Params
	err := c.B.Call("POST", "/customers/"+customer_id+"/bank_accounts", c.Key, body, commonParams, bankAccount)
	if err != nil {
		return nil, err
	}
	return bankAccount, nil
}
func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
