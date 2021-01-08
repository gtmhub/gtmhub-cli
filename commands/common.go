package commands

import (
	"gtmhub-cli/gtmhub_client"

	"github.com/dgrijalva/jwt-go"
)

var (
	client gtmhub_client.GtmhubHttpClient
)

var(
	logoutMsg = "You have succesfully logged out. Be quick and log back again!"
	authorizationRequestMsgFmt = "Please open your browser and navigate to: %s then enter the following code %s"
	nothingToUpdateCongratsMsg = `:tada::tada: Congrats. You are up to date, All your metrics are recently updated! :tada::tada:

If you still would like to update your krs you can either go through all of them using:
gtmhub show krs

or you can use one of your handy lists using:
gtmhub lists

Or maybe just pat yourself on the back for a job well done :beer:`
	listQueryNoItemsMsg = "no lists found matching the specified parameters"
	tooMuchListsMsg = ":confused: The specified parameter yealds more than one list. You will have to be more specific :confused:"
	tooMuchListsEnumMsgFmt = "ID: %s   Name: %s"
	metricUpdatedMsg = ":beer: Sweet! Your metric has just been updated. Keep going! :beer:"
)

func getClaimFromToken(idtoken string, lKey string) string {
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(idtoken, claims, nil)

	// ... error handling

	for key, val := range claims {
		if key == lKey {
			return val.(string)
		}
	}

	return ""
}
