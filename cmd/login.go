/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var Username string
var Password string

// var AQUA_URL string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LOGGING IN...")
		response := login(Username, Password, args[0])
		fmt.Println(string(response.Status))
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "Username (required if password is set)")
	loginCmd.Flags().StringVarP(&Password, "password", "p", "", "Password (required if username is set)")
	// loginCmd.Flags().StringVarP(&AQUA_URL, "host", "", "", "AQUA HOST (IP/NAME)")
	// loginCmd.MarkFlagsRequiredTogether("username", "password")

	rootCmd.AddCommand(loginCmd)
}

func login(username string, password string, aqua_url string) *http.Response {
	fmt.Println("Username: ", username)
	// fmt.Println("Password: ", password)
	fmt.Println("AQUA_URL: ", aqua_url)
	auth := fmt.Sprintf("%s:%s", username, password)
	authEncode := b64.StdEncoding.EncodeToString([]byte(auth))

	request, err := http.NewRequest(
		http.MethodGet,
		aqua_url,
		nil,
	)
	if err != nil {
		fmt.Println("Could not login. %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "NSTH CLI Tools")
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authEncode))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	return response
}
