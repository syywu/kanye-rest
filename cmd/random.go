package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate a random Kayne's quotes",
	Long:  `This command fetches a random quote from Kayne Rest API`,
	Run: func(cmd *cobra.Command, args []string) {
		getKayneQuote()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Quote struct {
Quote string `json:"quote"`
}

func getKayneQuote(){
	url:= "https://api.kanye.rest"
	responseBytes := getData(url)
	quote := Quote{}
	if err := json.Unmarshal(responseBytes, &quote); err != nil{
		log.Printf("unable to get unmarshal response %v", err)
	}
	fmt.Println(string(quote.Quote))
}

func getData(API string) []byte{
	request, err := http.NewRequest(
		http.MethodGet,
		API,
		nil,
	)
	if err!= nil{
		log.Printf("Unable to get data %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Kayne's quotes CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil{
		log.Println("Cannot make a request %v", err)
	}
	
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err!= nil{
		log.Println("Cannot read body %v", err)
	}

	return responseBytes
}