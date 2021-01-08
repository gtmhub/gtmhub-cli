package commands

import (
	"fmt"
	"gtmhub-cli/model"
	"gtmhub-cli/output"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

var (
	ListCommand = &cli.Command{
		Name:  "lists",
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
	var selectedList *model.ListResponse
	if len(id) > 0 {
		response, err = client.GetListsByID(id)
	} else if len(listName) > 0 {
		response, err = client.GetListsByName(listName)
	} else {
		model := interactiveListSelector{}
		p := tea.NewProgram(&model)

		if errP := p.Start(); errP != nil {
			fmt.Printf(errP.Error())
			os.Exit(1)
		}

		if model.selectedList == nil {
			err = fmt.Errorf("no list selected")
		} else {
			selectedList = model.selectedList
		}
	}

	if err != nil {
		return err
	}
	if selectedList == nil {
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

		selectedList = &lists[0]
	}



	result, err := client.LoadList(*selectedList)
	if err != nil {
		return err
	}

	output.PrintKrsFromList((*selectedList).Columns, result)

	return nil
}
