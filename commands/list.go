package commands

import (
	"fmt"
	"gtmhub-cli/model"
	"gtmhub-cli/output"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	ListCommand = &cli.Command{
		Name:  "list",
		Usage: "displays a kr list defined in gtmhub.",
		Action: ListAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases: []string{"n"},
				Usage:    "Specifies the name of the list to use. Defaults to the first kr list if none selected.",
			},
			&cli.StringFlag{
				Name: "id",
				Aliases: []string{"i"},
				Usage: "Specifies the id of the list you want to use. Takes precedents over -n.",
			},
		},
	}
)

func ListAction(c *cli.Context) error {

	var response model.FullListResponse
	var err error
	listName := c.String("name")
	id := c.String("id")
	if len(id) > 0 {
		response, err = client.GetListsByID(id)
	} else if len(listName) > 0 {
		response, err = client.GetListsByName(listName)
	}

	if err != nil {
		return err
	}

	lists := response.Items

	if len(lists) == 0 {
		output.Print(listQueryNoItemsMsg)
		os.Exit(0)
	}

	if len(lists) > 1 {
		output.Print(tooMuchListsMsg)
		fmt.Println()
		for _, list := range lists {
			output.Print(fmt.Sprintf(tooMuchListsEnumMsgFmt, list.ID, list.Title))
		}

		os.Exit(0)
	}

	result, err := client.LoadList(lists[0])
	if err != nil {
		return err
	}

	output.PrintKrsFromList(lists[0].Columns, result)

	return nil
	//if err != nil {
	//	return err
	//}
	//
	//for _, res := range result.Items {
	//	log.Println(fmt.Sprintf("ID:  %s  |  Name:  %s  |  Attainmenet:  %d", res.ID, res.Name, res.Attainmenet))
	//}

	//return fmt.Errorf("not implemented")
}
