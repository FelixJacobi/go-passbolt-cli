package cmd

import (
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/folder"
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/group"
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/resource"
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/settings"
	"github.com/haspr-dev/go-haspr-cli/internal/cmd/user"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = makeupdateCmd()

func makeupdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "update",
		Short:   "Updates a Passbolt Entity",
		Long:    `Updates a Passbolt Entity`,
		Aliases: []string{"change"},
	}
}

func configureUpdate() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(resource.ResourceUpdateCmd)
	updateCmd.AddCommand(folder.FolderUpdateCmd)
	updateCmd.AddCommand(group.GroupUpdateCmd)
	updateCmd.AddCommand(user.UserUpdateCmd)
	updateCmd.AddCommand(settings.NewUpdateCommand())
}

func init() { configureUpdate() }
