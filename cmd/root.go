package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flag Values
var globalFlag bool
var setValue string
var proenvCmd string

var rootCmd = &cobra.Command{
	Use:   "provm",
	Short: "Progress ABL Version Manager",
	Long:  `How to manage multiple versions of Progress ABL within one terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if globalFlag {
			config, err := readConfig()
			if err != nil {
				fmt.Println("Error reading config:", err)
				return
			}
			fmt.Println("Global version:", config.Global)
		}

		if setValue != "" {
			fmt.Printf("The global version called is : %s", setValue)
		}

		if proenvCmd != "" {
			fmt.Printf("The Proenv called is : %s", proenvCmd)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&globalFlag, "global", "g", false, "Global flag")
	rootCmd.PersistentFlags().StringVarP(&setValue, "set-global", "s", "", "Optional value for the global flag")
	rootCmd.PersistentFlags().StringVarP(&proenvCmd, "call", "c", "", "Optional value for the global flag")
}

type Config struct {
	Global   string `json:"global"`
	Versions []struct {
		Version string `json:"version"`
		Path    string `json:"path"`
	} `json:"versions"`
}

func readConfig() (Config, error) {
	var config Config

	// Replace 'config.json' with the path to your JSON file
	file, err := os.ReadFile("./config/config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
