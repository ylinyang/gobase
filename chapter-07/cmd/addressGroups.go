package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylinyang/gobase/chapter-07/third_party/huawei"
	"log"
)

var addressgroups string

var getCmd = &cobra.Command{
	Use: "get",
	Run: listGroups,
}

func listGroups(cmd *cobra.Command, args []string) {
	if err := huawei.ListAddressGroups(); err != nil {
		log.Println(err)
		return
	}
}
func init() {
	devCmd.AddCommand(getCmd)

	// 获取用户参数 此时args为空 直接将用户值赋给addressgroups  当前子命令有效不继承
	getCmd.Flags().StringVarP(&addressgroups, "addressgroups", "a", "", "get addressgroups info")
}
