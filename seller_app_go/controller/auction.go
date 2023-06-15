package controller

import (
	"seller_app_go/helpers"
	"seller_app_go/models"
	"time"
)

func AddAuction(input *helpers.AuctionPayload) (err error) {

	data := models.Auction{
		OwnerId:     input.OwnerID,
		Title:       input.Title,
		Description: input.Description,
		Status:      "open",
		BasePrice:   input.BasePrice,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		EndTime:     time.Now().UTC().Add(time.Duration(input.ExpiresAfterMinutes) * time.Minute),
	}

	return models.AddAuction(&data)
}

func GetAllOpenAuctions() (data []map[string]interface{}, err error) {

	auctions, err := models.GetAllOpenAuctions()
	if err != nil {
		return
	}
	data, err = helpers.MarshalUnmarshalArray(&auctions)
	return
}
