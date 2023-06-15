package helpers

import "encoding/json"

type AuctionPayload struct {
	OwnerID             string  `json:"owner_id"`
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	BasePrice           float64 `json:"base_price"`
	ExpiresAfterMinutes int     `json:"expires_after_minutes"`
}

func MarshalUnmarshalArray(data interface{}) (u []map[string]interface{}, err error) {

	d, err := json.Marshal(data)

	if err != nil {
		ErrorLog(err)
	}

	err = json.Unmarshal(d, &u)
	if err != nil {
		ErrorLog(err)
	}
	return
}

type BidingPayload struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	AuctionID string  `json:"auctionId"`
	Amount    float64 `json:"amount"`
}
