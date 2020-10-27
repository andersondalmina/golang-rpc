package services

import (
	"errors"
	"strconv"

	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/manifoldco/promptui"
)

// SearchBookMenu open search book menu
func SearchBookMenu() (action string, params string, err error) {
	prompt := promptui.Select{
		Label: "Search a book",
		Items: []string{"By title", "By author", "By year", "By edition"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", "", err
	}

	return handleSearchBookMenu(result)
}

func handleSearchBookMenu(item string) (action string, param string, err error) {
	switch item {
	case "By title":
		prompt := promptui.Prompt{
			Label: "Title",
		}

		title, err := prompt.Run()

		return "SearchBookByTitle", title, err
	case "By author":
		prompt := promptui.Prompt{
			Label: "Author",
		}

		author, err := prompt.Run()

		return "SearchBookByAuthor", author, err
	case "By year":
		prompt := promptui.Prompt{
			Label: "Published Year",
			Validate: func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return errors.New("Invalid number")
				}

				if len(input) > 4 {
					return errors.New("Max length 4")
				}

				return nil
			},
		}

		year, err := prompt.Run()

		return "SearchBookByYear", year, err
	case "By edition":
		prompt := promptui.Prompt{
			Label: "Book Edition",
			Validate: func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return errors.New("Invalid number")
				}

				if len(input) > 1 {
					return errors.New("Max length 1")
				}

				return nil
			},
		}

		edition, err := prompt.Run()

		return "SearchBookByEdition", edition, err
	}

	return "", "", nil
}

func SearchBooksByTitle(title string) []persist.Book {
	return persist.SearchBookByTitle(title)
}

func SearchBooksByAuthor(author string) []persist.Book {
	return persist.SearchBookByAuthor(author)
}

func SearchBooksByYear(year string) []persist.Book {
	return persist.SearchBookByYear(year)
}

func SearchBooksByEdition(edition string) []persist.Book {
	return persist.SearchBooksByEdition(edition)
}
