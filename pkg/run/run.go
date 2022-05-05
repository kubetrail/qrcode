package run

import (
	"fmt"
	"strings"

	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/kubetrail/qrcode/pkg/flags"
	"github.com/mdp/qrterminal/v3"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.OutputFilename, cmd.Flag(flags.OutputFilename))
	outputFilename := viper.GetString(flags.OutputFilename)

	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	var input string

	if len(args) > 0 {
		input = strings.Join(args, " ")
		input = mnemonics.Tidy(input)
	} else {
		if prompt {
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), "Enter input: "); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}
		}
		input, err = mnemonics.Read(cmd.InOrStdin())
	}

	if outputFilename == "-" {
		qrterminal.Generate(input, qrterminal.H, cmd.OutOrStdout())
		return nil
	}

	if err := qrcode.WriteFile(input, qrcode.Highest, 256, outputFilename); err != nil {
		return fmt.Errorf("failed to generate qrcode as png file: %w", err)
	}

	return nil
}
