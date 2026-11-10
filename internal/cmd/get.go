package cmd

import (
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/folder"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/group"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/resource"
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd/user"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets a Passbolt Entity",
	Long:    `Gets a Passbolt Entity`,
	Aliases: []string{"read"},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().BoolP("json", "j", false, "Output JSON")
	getCmd.AddCommand(resource.ResourceGetCmd)
	getCmd.AddCommand(folder.FolderGetCmd)
	getCmd.AddCommand(group.GroupGetCmd)
	getCmd.AddCommand(user.UserGetCmd)

}
