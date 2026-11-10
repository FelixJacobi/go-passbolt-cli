package cmd

import (
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/folder"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/group"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/resource"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/settings"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/user"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Updates a Passbolt Entity",
	Long:    `Updates a Passbolt Entity`,
	Aliases: []string{"change"},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(resource.ResourceUpdateCmd)
	updateCmd.AddCommand(folder.FolderUpdateCmd)
	updateCmd.AddCommand(group.GroupUpdateCmd)
	updateCmd.AddCommand(user.UserUpdateCmd)
	updateCmd.AddCommand(settings.NewUpdateCommand())
}
