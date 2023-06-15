package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const BID = "bid"

type Bid struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	AuctionId primitive.ObjectID `json:"auction_id" bson:"auction_id"`
	Amount    float64            `json:"amount" bson:"amount"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func AddBid(bid *Bid) (insertedID primitive.ObjectID, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := MongoDB.Collection(BID).InsertOne(ctx, &bid)
	if err != nil {
		fmt.Println(err, "Error in AddBid", bid.Email, err)
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return
}
