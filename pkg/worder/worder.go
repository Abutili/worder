package worder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	//"sync"

	"github.com/hashicorp/go-getter"
	//"github.com/spf13/viper"
)

// type RemoteDictionary struct {
// 	Name string
// 	Url  string
// 	Path string
// 	Dest string
// }

// type WorderCfg struct {
// 	RunDir string
// 	Dicts  []RemoteDictionary
// }

// https://go.dev/blog/strings
type LetterGraphNode struct {
	Letter string
}

type Dictionary struct {
	Name  string
	Words map[string]bool
}

// func (cfg WorderCfg) Initialize() {
// 	if _, err := os.Stat(cfg.RunDir); os.IsNotExist(err) {
// 		err := os.Mkdir(cfg.RunDir, os.ModePerm)
// 		if err != nil {
// 			log.Fatal("Error creating: ", cfg.RunDir, err)
// 		}
// 	}
// }

// func (cfg WorderCfg) LoadRemoteFiles() []Dictionary {

// 	dictionaries := []Dictionary{}

// 	// Define the maximum number of goroutines to run in parallel
// 	//maxParallel := 5
// 	//semaphore := make(chan struct{}, maxParallel)
// 	var wg sync.WaitGroup

// 	rds := viper.Get("dictionaries")
// 	log.Println("viper[dictionaries] -> ", rds)
// 	log.Printf("viper[dictionaries] Type -> %T\n", rds)

// 	// Iterate over the file paths and start a goroutine for each file
// 	// for _, rd := range rds {
// 	// 	// Add a goroutine to the wait group and get semaphore
// 	// 	wg.Add(1)
// 	// 	semaphore <- struct{}{}

// 	// 	// Start a goroutine to load the remotedictionaries
// 	// 	go func(name, url, path, dest string) {
// 	// 		// Helper to return semaphore and mark as done
// 	// 		defer func() {
// 	// 			<-semaphore
// 	// 			wg.Done()
// 	// 		}()

// 	// 		LoadRemoteFile(name, url, dest)

// 	// 		// this append is not thread safe ???
// 	// 		dictionaries = append(dictionaries, NewDictionary(name, dest))

// 	// 	}(rd["name"], rd["Url"], rd["Path"], rd["Dest"])
// 	// }

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	fmt.Println("All files loaded!")

// 	return dictionaries
// }

func LoadRemoteFile(fileUrl, filePath, fileDestination string) {

	_, err := os.Stat(fileDestination)

	if !os.IsNotExist(err) {
		log.Println(fileDestination, " file already exists.")
		return
	}

	// TODO: Specific construct the URL of the raw file on GitHub
	rawUrl := fileUrl
	if strings.Contains(fileUrl, "github") {
		rawUrl = fmt.Sprintf("%s/raw/master/%s", fileUrl, filePath)
	}

	log.Println("Fetching file ", rawUrl)
	log.Println(" Destination file: ", fileDestination)

	// mkdir for any missing dirs ?

	// Create a new getter client
	client := &getter.Client{
		Src:  rawUrl,
		Dst:  fileDestination,
		Mode: getter.ClientModeFile,
	}

	// Pull the file
	if err := client.Get(); err != nil {

		log.Fatal("Error pulling file:", err)
	}

	log.Println("Remote file pulled successfully")
	return
}

func NewDictionary(name, path string) (d Dictionary) {

	words := map[string]bool{}

	// load words in file rd.Path
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words[word] = true
		// build word graph here?
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	log.Println(path, " File read successfully!")

	d = Dictionary{
		Name:  name,
		Words: words,
	}
	return d
}
