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
	viper.BindEnv("word") // 绑定环境变量 PLUGIN_WORD

	rootCmd.AddCommand(exCmd)
	exCmd.AddCommand(ex.GetConfigCmd)
	ex.GetConfigCmd.Flags().StringP("word", "w", viper.GetString("word"), "传入要打印的内容 [$PLUGIN_WORD]")
	viper.BindPFlag("word", ex.GetConfigCmd.Flags().Lookup("word"))
	ex.GetConfigCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if viper.GetString("word") == "" {
			return fmt.Errorf("必须通过环境变量 PLUGIN_WORD 或命令行参数 --word 提供值")
		}
		return nil
	}
}
