package cmd

import (
	"fineC/runner"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Command() {
	var rootCmd = &cobra.Command{Use: "finc"}

	var inputModule string
	var outputModule string

	var cmdFinc = &cobra.Command{
		Use:   "finc",
		Short: "Finc cmd",
		Long:  `Finc is a tool for transferring data from one module to another.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Input Module: %s\n", inputModule)
			fmt.Printf("Output Module: %s\n", outputModule)
			runner := runner.NewRunner(inputModule, outputModule, "kubernetes/kubernetes")
			runner.Run()
			// 여기에 실제 로직을 구현하세요.
		},
	}

	cmdFinc.Flags().StringVarP(&inputModule, "input-module", "i", "", "Input module")
	cmdFinc.Flags().StringVarP(&outputModule, "output-module", "o", "", "Output module")

	rootCmd.AddCommand(cmdFinc)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
