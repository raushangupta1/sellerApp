package actions

import (
	"encoding/json"
	"io"
	"net/http"
	"seller_app_go/controller"
	"seller_app_go/helpers"

	"github.com/gobuffalo/buffalo"
)

func AddAuction(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		helpers.ErrorLog("ERROR in BODY", err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"message": "Invalid request body"}))
	}

	var data helpers.AuctionPayload
	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.ErrorLog(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"message": "Invalid request body"}))
	}

	err = controller.AddAuction(&data)
	if err != nil {
		helpers.ErrorLog(err, "failed to add auction", data)
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"message": "Request Failed"}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Auction Added Successfully"}))
}

func GetAllAuctions(c buffalo.Context) error {

	data, err := controller.GetAllOpenAuctions()
	if err != nil {
		helpers.ErrorLog(err, "failed to get auctions", data)
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"message": "Request Failed"}))
	}
	return c.Render(http.StatusOK, r.JSON(data))
}
