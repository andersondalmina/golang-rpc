package services

import (
	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/manifoldco/promptui"
)

// DeleteBookMenu open delete book menu
func DeleteBookMenu() (string, error) {
	prompt := promptui.Prompt{
		Label: "Book Title",
	}

	title, err := prompt.Run()

	return title, err
}

// DeleteBooksByTitle delete a book by the title given
func DeleteBooksByTitle(title string) error {
	return persist.DeleteBooksByTitle(title)
}
