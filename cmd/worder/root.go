package worder

import (
	"log"
	"sort"

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
		ds := config.LoadRemoteFiles()

		for _, cur_d := range ds {
			log.Println(cur_d.Name, " ", len(cur_d.Words))
		}

		d := ds[0]

		testing_word := "butter"
		if d.IsValidWord(testing_word) {
			log.Println(testing_word, " valid")
		} else {
			log.Println(testing_word, "not a valid word in ", d.Name)
		}

		// build map of sorted arrays by word
		// letters["A"] = { Word{ Word: "Add"} }

		// Create an empty slice to hold the strings
		sortedList := []string{}

		// Function to insert a new string into the sorted list
		insertString := func(newString string) {
			// Use binary search to find the index to insert the new string
			index := sort.Search(len(sortedList), func(i int) bool {
				return sortedList[i] >= newString
			})

			// Append the new string at the appropriate index
			sortedList = append(sortedList, "")
			copy(sortedList[index+1:], sortedList[index:])
			sortedList[index] = newString
		}

		// Insert new strings into the sorted list as they are discovered
		for _, w := range d.Words {
			insertString(w.Word)
		}
		//log.Println("Sorted list of strings:", sortedList) // Output: Sorted list of strings: [apple banana grape orange]

		wordMap := make(map[byte][]string)
		for _, w := range sortedList {
			//log.Println(w)

			if wl, exists := wordMap[w[0]]; exists {
				//log.Printf("The map contains the key '%s', and the value is %d.\n", w, value)
				wordMap[w[0]] = append(wl, w)
			} else {
				//log.Printf("The map does not contain the key '%s'.\n", w)
				wordMap[w[0]] = []string{w}
			}

		}
		log.Println(wordMap)

	},
}

func Execute() {

	log.Println("worder/cmd/worder.Execute")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Whoops. There was an error while executing your CLI '%s'", err)
	}
}
