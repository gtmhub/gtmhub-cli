package commands

import (
	"fmt"
	"gtmhub-cli/auth"
	"gtmhub-cli/config"
	"gtmhub-cli/output"

	"github.com/urfave/cli/v2"
)

var (
	LoginCommand = &cli.Command{
		Name:   "login",
		Usage:  "logs you in gtmhub",
		Action: LoginAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "dc",
				Aliases:     []string{"c"},
				Usage:       "specifies a data center to login to. Valid options are us and eu.",
				Value:       "eu",
				DefaultText: "eu",
				Required:    false,
			},
		},
	}

	payloadAccountIdKey = "https://gtmhub.com/app_metadata/accountId"

	baseUrls = map[string]string{
		"eu": "https://app.gtmhub.com",
		"us": "https://app.us.gtmhub.com",
	}
)

func LoginAction(c *cli.Context) error {
	// default is eu
	dc := c.String("dc")

	baseUrl, found := baseUrls[dc]
	if found == false {
		return fmt.Errorf("unrecognized dc specified. see gtmhub login -h for help")
	}

	config.SetGtmhubUrl(baseUrl)
	config.SetGtmhubDC(dc)

	authClient, err := auth.GetClient()
	if err != nil {
		return err
	}

	deviceCode, err := authClient.InitAuth()
	if err != nil {
		return err
	}
	output.Print(fmt.Sprintf(authorizationRequestMsgFmt, deviceCode.VerificationURL, deviceCode.UserCode))

	accessCodeResponse, err := authClient.PoolForToken(deviceCode)
	if err != nil {
		return err
	}

	config.SetRefreshToken(accessCodeResponse.RefreshToken)
	config.SetToken(accessCodeResponse.AccessToken)
	accId := getClaimFromToken(accessCodeResponse.IDToken, payloadAccountIdKey)
	config.SetAccountId(accId)

	auth0Id := getClaimFromToken(accessCodeResponse.AccessToken, "sub")

	domain, err := client.GetAccountDomain()
	if err != nil {
		return fmt.Errorf("could not resolve your account domain")
	}

	userId, err := client.GetUserID(domain, auth0Id)
	if err != nil {
		return err
	}

	config.SetUserID(userId)

	output.Print("Loggin successfull. Happy OKR-ing! :beer:", output.Green)
	return nil
}

//,
