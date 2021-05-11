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

// containmentCmd represents the containment command
var containmentCmd = &cobra.Command{
	Use:   "containment",
	Short: "calculates containment between two rectangles.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		recA := parseCoords(rectangleA)

		if recA == nil  {
			return
		}

		recB := parseCoords(rectangleB)

		if recB == nil  {
			return
		}

		recA.Describe()
		recB.Describe()

		if recA.Containment(*recB) {
			fmt.Printf("\n\n*** Containment between Rectangle A and Rectangle B: Found ***\n\n" )
		} else {
			fmt.Printf("\n\n*** Containment between Rectangle A and Rectangle B: NOT Found ***\n\n" )
		}
	},
}

func init() {
	rootCmd.AddCommand(containmentCmd)

	flags := containmentCmd.Flags()
	flags.StringVarP(&rectangleA, "rectangleA", "a", "", "rectangle A coords BottomLeft (x,y) and Top Right (x,y) as -a 6,4,14,10 -b  4,2,7,6")
	_ = intersectionCmd.MarkFlagRequired("rectangleA")

	flags.StringVarP(&rectangleB, "rectangleB", "b", "", "rectangle B coords BottomLeft (x,y) and Top Right (x,y) as -a 6,4,14,10 -b  4,2,7,6")
	_ = intersectionCmd.MarkFlagRequired("rectangleB")

}
