/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io"
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
		response, err := login(Username, Password, args[0])
		if err != nil {
			log.Printf("ERR. %v", err)
		}
		log.Println(string(response))
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
	loginCmd.MarkFlagsRequiredTogether("username", "password")

	rootCmd.AddCommand(loginCmd)
}

func login(usr string, pwd string, aqua_url string) ([]byte, error) {

	loginBody := map[string]string{
		"id":       usr,
		"password": pwd,
	}

	loginBodyM, err := json.Marshal(loginBody)
	if err != nil {
		log.Printf("ERR. %v", err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		aqua_url,
		bytes.NewReader(loginBodyM),
	)
	if err != nil {
		log.Printf("ERR. %v", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "NSTH CLI Tools")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("ERR. %v", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERR. %v", err)
	}

	return body, err
}
