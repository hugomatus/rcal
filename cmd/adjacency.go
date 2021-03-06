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
	"github.com/spf13/cobra"
)

// adjacencyCmd represents the adjacency command
var adjacencyCmd = &cobra.Command{
	Use:   "adjacency",
	Short: "calculates the adjacency between two rectangles",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		recA := parseCoords(rectangleA)
		recA.Describe()

		recB := parseCoords(rectangleB)
		recB.Describe()

		adjacencyType := recA.Adjacency(*recB)

		fmt.Printf("\n\n*** Adjacency between Rectangle A and Rectangle B : %v ***\n\n", adjacencyType)
	},
}

func init() {
	rootCmd.AddCommand(adjacencyCmd)

	flags := adjacencyCmd.Flags()
	flags.StringVarP(&rectangleA, "rectangleA", "a", "", "rectangle A coords")
	_ = intersectionCmd.MarkFlagRequired("rectangleA")

	flags.StringVarP(&rectangleB, "rectangleB", "b", "", "rectangle B coords")
	_ = intersectionCmd.MarkFlagRequired("rectangleB")
}
