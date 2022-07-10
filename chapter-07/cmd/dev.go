package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// devcmd
var devCmd = cobra.Command{
	// desc
	Use: "dev",
	//Function Description
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// Execute 将所有子命令添加到dev命令并适当设置标志。会被 main.main() 调用一次。
func Execute() {
	if err := devCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
}
