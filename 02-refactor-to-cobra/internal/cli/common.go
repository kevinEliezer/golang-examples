package cli

import "github.com/spf13/cobra"

const idFlag = "id"

type CobraFn func(cmd *cobra.Command, args []string)
