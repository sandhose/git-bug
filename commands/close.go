package commands

import (
	"errors"
	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/bug/operations"
	"github.com/spf13/cobra"
)

func runCloseBug(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("Only closing one bug at a time is supported")
	}

	if len(args) == 0 {
		return errors.New("You must provide a bug id")
	}

	prefix := args[0]

	b, err := bug.FindBug(repo, prefix)
	if err != nil {
		return err
	}

	author, err := bug.GetUser(repo)
	if err != nil {
		return err
	}

	op := operations.NewSetStatusOp(author, bug.ClosedStatus)

	b.Append(op)

	err = b.Commit(repo)

	return err
}

var closeCmd = &cobra.Command{
	Use:   "close <id>",
	Short: "Mark the bug as closed",
	RunE:  runCloseBug,
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
