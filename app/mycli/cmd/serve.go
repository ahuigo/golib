/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var username string
var sqls []string
var serveCmd = &cobra.Command{
	Use:   "serve [flags] sql1 sql2 sql3 ...",
	Short: "short msg for serve",
	Long:  `mycli serve is used to start the server`,
	Args:  cobra.MinimumNArgs(0), // 允许 0 个或多个参数
	Run: func(cmd *cobra.Command, args []string) {
		sqls = args // 将命令行参数赋值给 sqls
		fmt.Println("serve called, with args:", args)
		fmt.Println("config:", cfgFile)
		fmt.Println("user:", username)
		fmt.Println("SQLs:", sqls)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.LocalFlags().StringVarP(&username, "user", "u", "anonymouse", "用户名")
}
