package models

import (
	"context"
	"fmt"
	"seller_app_go/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AUCTION = "auction"

type Auction struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OwnerId     string             `json:"owner_id" bson:"owner_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	BasePrice   float64            `json:"base_price" bson:"base_price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	EndTime     time.Time          `json:"end_time" bson:"end_time"`
	Winner      Winner             `json:"winner" bson:"winner"`
}

type Winner struct {
	BidderId  primitive.ObjectID `json:"bidder_id" bson:"bidder_id"`
	BidAmount float64            `json:"bid_amount" bson:"bid_amount"`
}

func AddAuction(auction *Auction) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = MongoDB.Collection(AUCTION).InsertOne(ctx, &auction)
	if err != nil {
		fmt.Println(err, "Error in AddAuction", auction.Title, err)
	}
	return
}

func GetAllOpenAuctions() (res []Auction, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"end_time": bson.M{
			"$gt": time.Now().UTC(),
		},
		"status": "open",
	}

	cursor, err := MongoDB.Collection(AUCTION).Find(ctx, filter, options.Find())

	if err != nil {
		helpers.ErrorLog("Error in GetAllOpenAuctions", err)
		return
	}
	err = cursor.All(ctx, &res)
	if err != nil {
		helpers.ErrorLog("Error in decoding GetAllOpenAuctions", err)
	}
	return
}

func GetAuctionById(id primitive.ObjectID) (res Auction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	err = MongoDB.Collection(AUCTION).FindOne(ctx, filter).Decode(&res)

	if err != nil {
		helpers.ErrorLog("Error in GetAuctionById", err)
	}

	return
}

func UpdateWinner(auction *Auction) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": auction.Id,
	}

	update := bson.M{
		"$set": map[string]interface{}{
			"winner":     auction.Winner,
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := MongoDB.Collection(AUCTION).UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		helpers.ErrorLog("Error occured during UpdateWinner", err)
	}

	return
}

func UpdateAuction(auction *Auction) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": auction.Id,
	}

	update := bson.M{
		"$set": auction,
	}

	result, err := MongoDB.Collection(AUCTION).UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		helpers.ErrorLog("Error occured during UpdateAuction", err)
	}

	return
}

func GetAllOpenAuctionsWithExpiry() (res []Auction, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"status": "open",
	}

	cursor, err := MongoDB.Collection(AUCTION).Find(ctx, filter, options.Find())

	if err != nil {
		helpers.ErrorLog("Error in GetAllOpenAuctionsWithExpiry", err)
		return
	}
	err = cursor.All(ctx, &res)
	if err != nil {
		helpers.ErrorLog("Error in decoding GetAllOpenAuctionsWithExpiry", err)
	}
	return
}
