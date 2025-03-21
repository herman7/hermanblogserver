package hermanblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cfgFile string

func NewHermanBlogCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "hermanblog",
		Short: "Herman Blog",
		Long: `Herman Blog is a blog system for Herman written in Go.

Find more hermanblog information at:
        https://github.com/herman7/hermanblogserver#readme`,
		SilenceUsage: true, // 当命令执行失败时，不打印错误信息
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	command.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the hermanblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	command.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return command
}

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println("Hello HermanBlog!")
	return nil
}
