package worder

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/viper"
)

type RemoteDictionary struct {
	Name string
	Url  string
	Path string
	Dest string
}

type Config struct {
	AppName            string             `mapstructure:"app_name"`
	AppVersion         string             `mapstructure:"version"`
	DebugMode          bool               `mapstructure:"debug_mode"`
	RunDir             string             `mapstructure:"run_dir"`
	RemoteDictionaries []RemoteDictionary `mapstructure:"dictionaries"`
}

func Initalize() Config {

	// initial application config
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/worder/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.worder") // call multiple times to add many search paths
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error reading config file:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("failed to unmarshal configuration: %s", err)
	}
	log.Println("configuration successfully loaded -> ", config)

	// Setup local working directory
	if _, err := os.Stat(config.RunDir); os.IsNotExist(err) {
		err := os.Mkdir(config.RunDir, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating: ", config.RunDir, err)
		}
	}

	return config
}

func (cfg Config) LoadRemoteFiles() []Dictionary {

	log.Println("LoadRemoteFiles ")
	dictionaries := []Dictionary{}

	// Define the maximum number of goroutines to run in parallel
	maxParallel := 5
	semaphore := make(chan struct{}, maxParallel)
	var wg sync.WaitGroup

	// Iterate over the file paths and start a goroutine for each file
	for _, rd := range cfg.RemoteDictionaries {
		log.Println(rd)
		// 	// Add a goroutine to the wait group and get semaphore
		wg.Add(1)
		semaphore <- struct{}{}

		// 	// Start a goroutine to load the remotedictionaries
		go func(name, url, dest string) {
			// Helper to return semaphore and mark as done
			defer func() {
				<-semaphore
				wg.Done()
			}()

			LoadRemoteFile(url, dest)

			// this append is not thread safe ???
			dictionaries = append(dictionaries, NewDictionary(name, dest))

		}(rd.Name, rd.Url, rd.Dest)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All files loaded!")

	return dictionaries
}

func LoadRemoteFile(fileUrl, fileDestination string) {

	_, err := os.Stat(fileDestination)

	if !os.IsNotExist(err) {
		log.Println(fileDestination, " file already exists.")
		return
	}

	log.Println("Fetching file ", fileUrl)
	log.Println("Destination file: ", fileDestination)

	// mkdir for any missing dirs ?

	// Create a new getter client
	client := &getter.Client{
		Src:  fileUrl,
		Dst:  fileDestination,
		Mode: getter.ClientModeFile,
	}

	// Pull the file
	if err := client.Get(); err != nil {

		log.Fatal("Error pulling file:", err)
	}

	log.Println("Remote file pulled successfully")
}
