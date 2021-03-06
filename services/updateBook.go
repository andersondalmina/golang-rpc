package services

import (
	"errors"
	"strconv"

	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/manifoldco/promptui"
)

func UpdateBookMenu() (map[string]string, error) {
	prompt := promptui.Prompt{
		Label: "Book Code",
		Validate: func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("Invalid number")
			}

			return nil
		},
	}

	id, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Title",
	}

	title, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Author",
	}

	author, err := prompt.Run()

	prompt = promptui.Prompt{
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

	prompt = promptui.Prompt{
		Label: "Book Year",
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

	params := map[string]string{
		"id":      id,
		"title":   title,
		"author":  author,
		"edition": edition,
		"year":    year,
	}

	return params, err
}

func UpdateBook(p map[string]string) (persist.Book, error) {
	id, err := strconv.Atoi(p["id"])
	if err != nil {
		panic(err)
	}

	edition, err := strconv.Atoi(p["edition"])
	if err != nil {
		panic(err)
	}

	year, err := strconv.Atoi(p["year"])
	if err != nil {
		panic(err)
	}

	return persist.UpdateBook(int64(id), p["title"], p["author"], int64(edition), int64(year))
}
