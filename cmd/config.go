package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)
type ConfigVersion struct {
	Global   string        `json:"global"`
	Versions []VersionInfo `json:"versions"`
}
type VersionInfo struct {
	Version string `json:"version"`
	Path    string `json:"path"`
}

var versionName string
var versionPath string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure versions",
	Long:  `Configure the versions and paths for the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionName == "" {
			fmt.Print("Enter version name: ")
			versionName = readUserInput()
		}

		if versionPath == "" {
			fmt.Print("Enter version path: ")
			versionPath = readUserInput()
		}

		config, err := readVersionConfig()
		if err != nil {
			fmt.Println("Error reading config:", err)
			return
		}

		config.Versions = append(config.Versions, VersionInfo{Version: versionName, Path: versionPath})

		if err := writeVersionConfig(config); err != nil {
			fmt.Println("Error writing config:", err)
			return
		}

		fmt.Println("Version and path added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&versionName, "version", "v", "", "Version name")
	configCmd.Flags().StringVarP(&versionPath, "path", "p", "", "Version path")
}

func readVersionConfig() (ConfigVersion, error) {
	var config ConfigVersion
	file, err := os.ReadFile("./config/config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	return config, err
}

func writeVersionConfig(config ConfigVersion) error {
	file, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("./config/config.json", file, 0644)
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
