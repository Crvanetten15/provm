package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
			config, err := readConfig()
			if err != nil {
				fmt.Println("Error reading config:", err)
				return
			}

			var versionExists bool
			for _, v := range config.Versions {
				if v.Version == setValue {
					versionExists = true
					break
				}
			}

			if versionExists {
				config.Global = setValue
				if err := writeConfig(config); err != nil {
					fmt.Println("Error writing config:", err)
					return
				}
				fmt.Printf("Global version updated to: %s\n", setValue)
			} else {
				fmt.Printf("Version %s not found in the configuration\n", setValue)
			}
		}

		if proenvCmd != "" {
			config, err := readConfig()
			if err != nil {
				fmt.Println("Error reading config:", err)
				return
			}

			var pathFound bool
			for _, v := range config.Versions {
				if v.Version == config.Global {
					updatedPath := v.Path + proenvCmd
					fmt.Printf("Updated path for global version %s: %s\n", config.Global, updatedPath)
					pathFound = true
					break
				}
			}

			if !pathFound {
				fmt.Println("Global version not found in configuration")
			}
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

func writeConfig(config Config) error {
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("./config/config.json", file, 0644)
}
