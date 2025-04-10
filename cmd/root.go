/*
Copyright © 2021 eryajf

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
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// var cfgFile string
var (
	Version   string
	GitCommit string
	BuildTime string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eryajfctl",
	Short: "eryajf goscript",
	Long:  `🦄 利用cobra制作运维日常工具箱的框架。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 检查是否有 -v 参数
		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			fmt.Println(cmd.VersionTemplate())
			return
		}
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&MarkdownDocs, "md-docs", "m", false, "gen Markdown docs")
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate(fmt.Sprintf(`🍉 {{with .Name}}{{printf "%%s version information: " .}}{{end}}
  {{printf "Version:    %%s" .Version}}
  Git Commit: %s
  Go version: %s
  OS/Arch:    %s/%s
  Build Time: %s

`, GitCommit, runtime.Version(), runtime.GOOS, runtime.GOARCH, BuildTime))
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")
}

var MarkdownDocs bool

func GenDocs() {
	if MarkdownDocs {
		if err := doc.GenMarkdownTree(rootCmd, "./docs"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// 生成配置文件
func GenConfig() {

}
