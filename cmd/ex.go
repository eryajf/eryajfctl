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

	"github.com/eryajf/eryajfctl/api/ex"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// exCmd represents the jenkins command
var exCmd = &cobra.Command{
	Use:   "ex",
	Short: `这是一个示例,你可以参考帮助信息使用,见: https://github.com/eryajf/eryajfctl/blob/main/README.md`,
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PLUGIN")
	viper.BindEnv("ename") // 绑定环境变量 PLUGIN_ENAME
	viper.BindEnv("epass") // 绑定环境变量 PLUGIN_EPASS

	rootCmd.AddCommand(exCmd)
	exCmd.AddCommand(ex.GetConfigCmd)
	ex.GetConfigCmd.Flags().StringP("ename", "n", viper.GetString("ename"), "传入要打印的内容 [$PLUGIN_ENAME]")
	viper.BindPFlag("ename", ex.GetConfigCmd.Flags().Lookup("ename"))
	ex.GetConfigCmd.Flags().StringP("epass", "p", viper.GetString("epass"), "传入要打印的内容 [$PLUGIN_EPASS")
	viper.BindPFlag("epass", ex.GetConfigCmd.Flags().Lookup("epass"))
	// 遍历所有 flags，清除默认值占位符，避免在日志中打印
	ex.GetConfigCmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.DefValue = ""
	})
	ex.GetConfigCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if viper.GetString("ename") == "" {
			return fmt.Errorf("必须通过环境变量 PLUGIN_ENAME 或命令行参数 --ename 提供值")
		}
		if viper.GetString("epass") == "" {
			return fmt.Errorf("必须通过环境变量 PLUGIN_EPASS 或命令行参数 --epass 提供值")
		}
		return nil
	}
}
