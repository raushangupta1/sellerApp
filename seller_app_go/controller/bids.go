package controller

import (
	"errors"
	"seller_app_go/helpers"
	"seller_app_go/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBid(input *helpers.BidingPayload) (err error) {

	auctionId, _ := primitive.ObjectIDFromHex(input.AuctionID)
	auction, _ := models.GetAuctionById(auctionId)
	if (auction.Id == primitive.ObjectID{}) || auction.Status != "open" || auction.EndTime.Before(time.Now().UTC()) {
		helpers.ErrorLog("Auction not allowed", auction)
		return errors.New("auction not allowed")
	}
	data := models.Bid{
		Name:      input.Name,
		Email:     input.Email,
		AuctionId: auction.Id,
		Amount:    float64(input.Amount),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	bidId, err := models.AddBid(&data)
	if err != nil {
		return
	}
	if data.Amount > auction.Winner.BidAmount {
		auction.Winner.BidderId = bidId
		auction.Winner.BidAmount = data.Amount
		return models.UpdateWinner(&auction)
	}
	return
}
