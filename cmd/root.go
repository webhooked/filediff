/*
Copyright © 2021 Jonathan Edwardsson (https://github.com/webhooked)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorBlue string = "\033[34m"
var colorWhite string = "\033[37m"

var rootCmd = &cobra.Command{
	Use:   "filediff",
	Short: "Takes two files and displays their differences",
	Long:  `FileDiff takes two files and displays their differences`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			filename1 := args[0]
			filename2 := args[1]

			file1 := readFileLines(filename1)
			file2 := readFileLines(filename2)

			if equal(file1, file2) {
				fmt.Println("File contents are the same.")
			} else {
				fmt.Println()
				fmt.Println(colorBlue, "Differences between", filename1, "and", filename2, colorReset)
				fmt.Println()

				if len(file1) >= len(file2) {
					checkDifferences(file1, file2)
				} else {
					checkDifferences(file2, file1)
				}

				fmt.Println()
			}
		} else {
			fmt.Println()
			fmt.Println(colorBlue, "--- FileDiff Usage Example ---", colorReset)
			fmt.Println()
			fmt.Print(colorWhite, "filediff ", colorReset)
			fmt.Print(colorGreen, "file1.css ", colorReset)
			fmt.Print(colorRed, "file2.css", colorReset)
			fmt.Println()
			fmt.Println()
		}

	},
}

func readFileLines(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	check(err)
	file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func checkDifferences(a, b []string) {
	var removedLines []string
	var addedLines []string

	for i, v := range a {
		if i < len(b) {
			if v != b[i] {
				removedLines = append(removedLines, v)
				addedLines = append(addedLines, b[i])
			} else {
				if len(removedLines) > 0 {
					printDifferences(removedLines, addedLines)
				} else {
					fmt.Println(colorWhite, "  "+v, colorReset)
				}
				removedLines = removedLines[:0]
				addedLines = addedLines[:0]
			}
		} else {
			removedLines = append(removedLines, v)
			addedLines = append(addedLines, "")
		}
	}

	if len(removedLines) > 0 {
		printDifferences(removedLines, addedLines)
		removedLines = removedLines[:0]
		addedLines = addedLines[:0]
	}
}

func printDifferences(removed, added []string) {
	for _, line := range removed {
		fmt.Println(colorRed, "- "+line, colorReset)
	}

	for _, line := range added {
		fmt.Println(colorGreen, "+ "+line, colorReset)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.filediff.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".filediff")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
