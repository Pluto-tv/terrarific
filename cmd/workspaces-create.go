/*
Copyright © 2020 Iggy Jackson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// workspacesCreateCmd represents the workspacesCreate command
var workspacesCreateCmd = &cobra.Command{
	Use:   "create <organization> <workspace>",
	Short: "Create a new workspace",
	Long:  `Create a workspace and set it up according to requested arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspaces create not implemented yet")
	},
}

func init() {
	workspacesCmd.AddCommand(workspacesCreateCmd)

	workspacesCreateCmd.Flags().Bool("autoapply", false, "Whether to automatically apply after a successful plan")
	workspacesCreateCmd.Flags().StringArrayP("env", "e", []string{}, "Environment variables to set on the workspace. Can be passed multiple times. (-e KEY=Value)")
	workspacesCreateCmd.Flags().StringP("version", "v", "latest", "Version of terraform to run in the terraform cloud runners (ex. latest, 0.12.29, 0.11.14")
	workspacesCreateCmd.Flags().Bool("remote", true, "Whether to run commands in terraform cloud runners (remote) or locally on a workstation")
	workspacesCreateCmd.Flags().String("workingdir", "", "Directory to run terraform commands in")
}
