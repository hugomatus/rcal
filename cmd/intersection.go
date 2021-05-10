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

// intersectionCmd represents the intersection command
var intersectionCmd = &cobra.Command{
	Use:   "intersection",
	Short: "calculates the intersection between two rectangles if any exist",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		recA := parseCoords(rectangleA)
		recA.Describe()

		recB := parseCoords(rectangleB)
		recB.Describe()

		points := recA.Intersection(*recB)

		for k, v := range points {
			fmt.Printf("Points of Intersection in Block No. %d : %v \n", k, v)

		}
	},
}

func init() {
	rootCmd.AddCommand(intersectionCmd)

	flags := intersectionCmd.Flags()
	flags.StringVarP(&rectangleA, "rectangleA", "a", "", "rectangle A coords")
	_ = intersectionCmd.MarkFlagRequired("rectangleA")

	flags.StringVarP(&rectangleB, "rectangleB", "b", "", "rectangle B coords")
	_ = intersectionCmd.MarkFlagRequired("rectangleB")
}
