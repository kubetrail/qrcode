/*
Copyright © 2022 kubetrail.io authors

*/
package cmd

import (
	"github.com/kubetrail/qrcode/pkg/flags"
	"github.com/kubetrail/qrcode/pkg/run"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate qrcode",
	Long: `Generating qrcode for an input string either
by piping it in or passing as command line args

Run as:
    echo this is a test | qrcode gen
or:
    qrcode gen this is a test
`,
	RunE: run.Run,
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().String(flags.OutputFilename, "-", "Output filename")
}
