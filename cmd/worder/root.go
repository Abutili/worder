package worder

import (
	"log"

	"abutili.com/worder/pkg/worder"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "worder",
	Short: "worder - a simple CLI to work with words and word fragments",
	Long: `worder can load multiple dictionary sources and provide both basic
	word search as well as graph manipulation and searching for letter 
	combinations to form words for contraint satisfaction problems.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Running root worder command.")

		// initial application config
		var config worder.Config
		config = worder.Initalize()

		log.Println(config)

	},
}

func Execute() {

	log.Println("worder/cmd/worder.Execute")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Whoops. There was an error while executing your CLI '%s'", err)
	}
}
