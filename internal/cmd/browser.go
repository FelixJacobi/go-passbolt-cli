package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/passbolt/go-passbolt-cli/internal/agent"
	"github.com/passbolt/go-passbolt-cli/internal/browserauth"
	"github.com/spf13/cobra"
)

func newBrowserCommand() *cobra.Command {
	command := &cobra.Command{Use: "browser", Short: "Install browser authorization integration"}
	var executable string
	install := &cobra.Command{Use: "install", Args: cobra.NoArgs, RunE: func(_ *cobra.Command, _ []string) error {
		if executable == "" {
			self, err := os.Executable()
			if err != nil {
				return err
			}
			executable = filepath.Join(filepath.Dir(self), "passbolt-browser-host")
			if filepath.Ext(self) == ".exe" {
				executable += ".exe"
			}
		}
		return browserauth.Install(executable)
	}}
	install.Flags().StringVar(&executable, "host-path", "", "Absolute path to the passbolt-browser-host executable")
	command.AddCommand(install, &cobra.Command{Use: "uninstall", Args: cobra.NoArgs, RunE: func(_ *cobra.Command, _ []string) error { return browserauth.Uninstall() }})
	return command
}

// RunBrowserHost handles only native messages and never loads CLI configuration.
// Its stdout belongs exclusively to the browser native messaging protocol.
func RunBrowserHost(in io.Reader, out io.Writer) error {
	for {
		var request struct {
			Action   string               `json:"action"`
			Envelope browserauth.Envelope `json:"envelope"`
		}
		if err := browserauth.ReadNative(in, &request); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		response, err := relayBrowserMessage(request.Action, request.Envelope)
		if err != nil {
			response = map[string]any{"error": "Browser authorization unavailable. Start a CLI command using --agent --agent-auth=browser and try again."}
		}
		if err = browserauth.WriteNative(out, response); err != nil {
			return err
		}
	}
}
func relayBrowserMessage(action string, envelope browserauth.Envelope) (any, error) {
	if action != "pending" && action != "submit" {
		return nil, fmt.Errorf("unknown native browser action")
	}
	endpoint, err := agent.Endpoint()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := agent.Dial(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()
	_ = conn.SetDeadline(time.Now().Add(5 * time.Second))
	req := agent.Request{Version: agent.ProtocolVersion, Action: "browser-" + action}
	if action == "submit" {
		req.BrowserData, err = json.Marshal(envelope)
		if err != nil {
			return nil, err
		}
	}
	if err = agent.WriteFrame(conn, req); err != nil {
		return nil, err
	}
	var result agent.Message
	if err = agent.ReadFrame(conn, &result); err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, fmt.Errorf("browser request failed")
	}
	if action == "submit" {
		return map[string]bool{"submitted": true}, nil
	}
	var pending browserauth.Request
	if err = json.Unmarshal(result.Data, &pending); err != nil {
		return nil, err
	}
	return map[string]any{"request": pending}, nil
}
