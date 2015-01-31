package stripe

import (
	"encoding/json"
	"net/url"
)

// RecipientType is the list of allowed values for the recipient's type.
// Allowed values are "individual", "corporation".
type RecipientType string


// RecipientParams is the set of parameters that can be used when creating or updating recipients.
// For more details see https://stripe.com/docs/api#create_recipient and https://stripe.com/docs/api#update_recipient.
type RecipientParams struct {
	Params
	Name                      string
	Type                      RecipientType
	TaxID, Token, Email, Desc string
	Bank                      *BankAccountParams
	Card                      *CardParams
	DefaultCard               string
}

// RecipientListParams is the set of parameters that can be used when listing recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
type RecipientListParams struct {
	ListParams
	Verified bool
}


// Recipient is the resource representing a Stripe recipient.
// For more details see https://stripe.com/docs/api#recipients.
type Recipient struct {
	ID          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Created     int64             `json:"created"`
	Type        RecipientType     `json:"type"`
	Bank        *BankAccount      `json:"active_account"`
	Desc        string            `json:"description"`
	Email       string            `json:"email"`
	Meta        map[string]string `json:"metadata"`
	Name        string            `json:"name"`
	Cards       *CardList         `json:"cards"`
	DefaultCard *Card             `json:"default_card"`
}


// AppendDetails adds the bank account's details to the query string values.
func (b *BankAccountParams) AppendDetails(values *url.Values) {
	values.Add("bank_account[country]", b.CountryCode)
	values.Add("bank_account[routing_number]", b.RoutingNumber)
	values.Add("bank_account[account_number]", b.AccountNumber)
}

// UnmarshalJSON handles deserialization of a Recipient.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Recipient) UnmarshalJSON(data []byte) error {
	type recipient Recipient
	var rr recipient
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Recipient(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
