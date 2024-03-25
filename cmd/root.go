/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/learnselfs/whs"
	"github.com/learnselfs/ws/config"
	"github.com/learnselfs/ws/routes"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ws",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: New,
}

func New(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")
	staticPath, _ := cmd.Flags().GetString("staticPath")
	staticRoute, _ := cmd.Flags().GetString("staticRoute")
	htmlPath, _ := cmd.Flags().GetString("htmlPath")
	config.Server = whs.New(host, port)

	// html css js
	Route(staticRoute, staticPath, htmlPath)

	config.Server.Start()
}

func Route(staticRoute, staticPath, htmlPath string) {
	routes.Routes()
	routes.SetStatic(staticRoute, staticPath, htmlPath)

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ws.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("host", "H", "localhost", "ws server bond address")
	rootCmd.Flags().IntP("port", "P", 80, "ws server bond port")
	rootCmd.Flags().StringP("staticRoute", "r", "/static/", "ws server bond address")
	rootCmd.Flags().StringP("staticPath", "t", "front/static", "ws server bond address")
	rootCmd.Flags().StringP("htmlPath", "s", "front/template/**/*", "ws server bond address")
}
