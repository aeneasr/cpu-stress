// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"github.com/pborman/uuid"
	"fmt"
	"sync"
	"runtime"
	"strconv"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cpus := int64(runtime.NumCPU())
		if len(args) > 0 {
			cpus, _ = strconv.ParseInt(args[0], 10, 64)
		}

		var wg sync.WaitGroup
		for i := 0; i <= int(cpus); i++ {
			wg.Add(1)
			fmt.Printf("Starting thread %d\n", i)
			go cpuStress()
		}
		wg.Wait()
	},
}

func cpuStress () {
	for {
		_, err := bcrypt.GenerateFromPassword([]byte(uuid.New()), 20)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	RootCmd.AddCommand(cpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
