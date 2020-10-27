package cmd

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/andersondalmina/golang-rpc/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [host] [port]",
	Short: "Initialize the TCP Server",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		persist.CreateDatabase()

		fmt.Println("TCP server starting")

		Server := new(server.Server)
		rpc.Register(Server)

		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", args[0]+":"+args[1])
		checkError(err)

		http.Serve(listener, nil)
	},
}
