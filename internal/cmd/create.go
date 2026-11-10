package cmd

import (
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/folder"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/group"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/resource"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/user"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates a Passbolt Entity",
	Long:    `Creates a Passbolt Entity`,
	Aliases: []string{"new"},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().BoolP("json", "j", false, "Output JSON")
	createCmd.AddCommand(resource.ResourceCreateCmd)
	createCmd.AddCommand(folder.FolderCreateCmd)
	createCmd.AddCommand(group.GroupCreateCmd)
	createCmd.AddCommand(user.UserCreateCmd)
}
