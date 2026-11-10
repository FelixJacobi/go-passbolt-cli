package cmd

import (
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/folder"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/group"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/resource"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/settings"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/user"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = makedeleteCmd()

func makedeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "delete",
		Short:   "Deletes a Passbolt Entity",
		Long:    `Deletes a Passbolt Entity`,
		Aliases: []string{"remove"},
	}
}

func configureDelete() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(resource.ResourceDeleteCmd)
	deleteCmd.AddCommand(folder.FolderDeleteCmd)
	deleteCmd.AddCommand(group.GroupDeleteCmd)
	deleteCmd.AddCommand(user.UserDeleteCmd)
	deleteCmd.AddCommand(settings.NewDeleteCommand())

	deleteCmd.PersistentFlags().String("id", "", "ID of the Entity to Delete")
}

func init() { configureDelete() }
