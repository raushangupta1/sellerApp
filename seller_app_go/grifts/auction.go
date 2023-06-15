package grifts

import (
	"fmt"
	"seller_app_go/helpers"
	"seller_app_go/models"
	"time"

	"github.com/gobuffalo/grift/grift"
)

var _ = grift.Desc("close_expired_auction", "Close expired auctions")
var _ = grift.Add("close_expired_auction", func(c *grift.Context) (err error) {

	auctions, err := models.GetAllOpenAuctionsWithExpiry()
	if err != nil {
		helpers.ErrorLog("Error in GetAllOpenAuctions", err)
		return
	}
	fmt.Println(auctions)
	for _, auction := range auctions {
		if auction.EndTime.Before(time.Now().UTC()) {
			auction.Status = "closed"
			auction.UpdatedAt = time.Now().UTC()
			err = models.UpdateAuction(&auction)
			if err != nil {
				helpers.ErrorLog("Error in UpdateAuction", err)
			}
		}
	}
	return
})
