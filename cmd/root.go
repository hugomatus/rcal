/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/hugomatus/rcal/rectangle"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string

var rectangleA string

var rectangleB string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rcal",
	Short: "rcal cli calculates intersection, containment and adjacency between two rectangles.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rcal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rcal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rcal")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// parseCoords
func parseCoords(coords string) *rectangle.Rectangle {

	tmp_ := strings.Split(coords, ",")

	if len(tmp_) != 4 {
		fmt.Println("Coordinates for Rectangle A and B are required \ni.e. BottomLeft (x,y) and Top Right (x,y) as -a 6,4,14,10 -b  4,2,7,6")

		return nil
	}

	var coordsAsInt []int
	var result []rectangle.Point

	for _, v := range tmp_ {

		val, _ := strconv.Atoi(v)
		coordsAsInt = append(coordsAsInt, val)
	}

	result = append(result, rectangle.Point{
		X: coordsAsInt[0],
		Y: coordsAsInt[1],
	}, rectangle.Point{
		X: coordsAsInt[2],
		Y: coordsAsInt[3],
	})

	rec := rectangle.New(result[1], result[0])

	return rec
}
