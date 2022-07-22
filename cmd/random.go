package cmd

import (
	"fmt"

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

}

func getRandomFact(){
	fmt.Println("random fact")
}