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
	Short: "Generate a random fact",
	Long:  `This command fetches a random fact from FACTS API`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomFact()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Fact struct {
Fact string 
}

func getRandomFact(){
	url:= "http://numbersapi.com/random/trivia"
	responseBytes := getData(url)
	fact := Fact{}
	if err := json.Unmarshal(responseBytes, &fact); err != nil{
		log.Printf("unable to get unmarshal response %v", err)
	}
	fmt.Println(string(fact.Fact))
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
	request.Header.Add("User-Agent", "Random Facts CLI")

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