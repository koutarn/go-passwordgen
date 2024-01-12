/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	generate "github.com/koutarn/go-passwordgen/pkg"
	"github.com/spf13/cobra"
)

var passwordLength = 10
var lowcaseFlag = true
var uppercaseFlag = true
var specialCharsFlag = true
var numbersFlag = true

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-passwordgen",
	Short: "パスワードをいい感じに作成してくれます",
	Run: func(cmd *cobra.Command, args []string) { 
		var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()0123456789"
		const ignoreLowcase = "abcdefghijklmnopqrstuvwxyz"
		const ignoreUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		const ignoreSpecialChars = "!@#$%^&*()"
		const ignoreNumbers = "0123456789"

		if lowcaseFlag {
			charset = strings.ReplaceAll(charset, ignoreLowcase, "")
		}

		if uppercaseFlag {
			charset = strings.ReplaceAll(charset, ignoreUppercase, "")
		}

		if specialCharsFlag {
			charset = strings.ReplaceAll(charset, ignoreSpecialChars, "")
		}

		if numbersFlag {
			charset = strings.ReplaceAll(charset, ignoreNumbers, "")
		}

		if charset == "" {
			fmt.Println("Error generating password")
			return			
		}

		password, err := generate.GeneratePassword(passwordLength, charset)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		fmt.Println(password)

	},
}



// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&passwordLength,"length","l",10,"パスワードの長さを指定します")
	rootCmd.Flags().BoolVarP(&lowcaseFlag,"lower","c",false,"小文字を除外")
	rootCmd.Flags().BoolVarP(&uppercaseFlag,"upper","C",false,"大文字を除外")
	rootCmd.Flags().BoolVarP(&specialCharsFlag,"special","s",false,"数字を除外")
	rootCmd.Flags().BoolVarP(&numbersFlag,"number","n",false,"数字を除外")
}


