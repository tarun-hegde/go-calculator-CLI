/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
)
func m(args []string) int{
	mul:=1
	for _,arg:=range args{
	k,err:=strconv.Atoi(arg)
	if err!=nil{
		fmt.Println("Error")
	}
	mul=mul*k
}
return mul
}
func s(args []string) int{
	sum:=0
	for _,arg:=range args{
	k,err:=strconv.Atoi(arg)
	if err!=nil{
		fmt.Println("Error")
	}
	sum=sum+k
}
return sum
}
func divide(a,b int) int{
	if b==0{
		fmt.Println("Error")
	}
	return a/b
}
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi, I am root command.")
	 },
}
//sum command
var sCmd = &cobra.Command{
	Use:   "add  num[1] num[2] ...",
	Short: "A brief description of your application",
	Aliases: []string{"sum"},
	Args: cobra.MinimumNArgs(2),
Run: func(cmd *cobra.Command, args []string ) {
	res:=s(args)
	fmt.Println(res)
},
}
//divide command
var divideCmd = &cobra.Command{
	Use:   "d [num1] [num2]",
	Short: "Divides two numbers",
	Args: cobra.ExactArgs(2),
	Aliases: []string{"div"},
	Run: func(cmd *cobra.Command, args []string) {
		a,err:=strconv.Atoi(args[0]);
		b,err:=strconv.Atoi(args[1]);
		if err!=nil{
			fmt.Println("Error")
		}
		res:=divide(a,b)
		fmt.Println(res)
	 },
}
//command to perform multiplication
var productCmd = &cobra.Command{
	Use:   "p num[1] num[2] ...",
	Short: "A brief description of your application",
	Aliases: []string{"p"},
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res:=m(args)
		fmt.Println(res)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
rootCmd.AddCommand(productCmd)
rootCmd.AddCommand(sCmd)
rootCmd.AddCommand(divideCmd)
}


