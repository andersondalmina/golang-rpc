package cmd

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/andersondalmina/golang-rpc/services"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var rpcClient *rpc.Client

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client [host] [port]",
	Short: "Initialize the TCP Client",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		for {
			var err error
			rpcClient, err = rpc.DialHTTP("tcp", args[0]+":"+args[1])
			checkError(err)

			openMenu()
		}
	},
}

func openMenu() {
	prompt := promptui.Select{
		Label: "Select your action",
		Items: []string{"Create a book", "Update a Book", "Search a book", "Delete a book", "Exit"},
	}

	_, result, err := prompt.Run()
	checkError(err)

	handleMenu(result)
}

func handleMenu(item string) {
	switch item {
	case "Create a book":
		params, err := services.CreateBookMenu()
		checkError(err)

		var reply persist.Book
		err = rpcClient.Call("Server.CreateBook", params, &reply)
		checkError(err)

		fmt.Println("Book created successfully")

	case "Update a Book":
		params, err := services.UpdateBookMenu()
		checkError(err)

		var reply persist.Book
		err = rpcClient.Call("Server.UpdateBook", params, &reply)
		checkError(err)

		fmt.Println("Book updated successfully")

	case "Search a book":
		action, param, err := services.SearchBookMenu()
		checkError(err)

		var reply []persist.Book
		err = rpcClient.Call("Server."+action, param, &reply)
		checkError(err)

		services.DisplayBooks(reply)

	case "Delete a book":
		title, err := services.DeleteBookMenu()
		checkError(err)

		var reply persist.Book
		err = rpcClient.Call("Server.DeleteBookByTitle", title, &reply)
		checkError(err)

		fmt.Println("Book removed successfully")

	case "Exit":
		rpcClient.Close()
		os.Exit(0)
	}

	fmt.Println()
	rpcClient.Close()
}
