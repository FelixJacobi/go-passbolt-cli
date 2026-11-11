package cmd

import (
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/folder"
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/resource"
	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = makemoveCmd()

func makemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "move",
		Short: "Moves a Passbolt Entity",
		Long:  `Moves a Passbolt Entity`,
	}
}

func configureMove() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.AddCommand(resource.ResourceMoveCmd)
	moveCmd.AddCommand(folder.FolderMoveCmd)
}

func init() { configureMove() }
