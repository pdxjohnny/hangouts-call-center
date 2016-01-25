package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/hangouts-call-center/call"
	"github.com/pdxjohnny/hangouts-call-center/caller"
	"github.com/pdxjohnny/hangouts-call-center/end"
	"github.com/pdxjohnny/hangouts-call-center/http"
)

// Commands are the commands that can be run
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "call",
		Short: "Make a call",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			call.Run()
		},
	},
	&cobra.Command{
		Use:   "end",
		Short: "End a call",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			end.Run()
		},
	},
	&cobra.Command{
		Use:   "caller",
		Short: "Communicate with the call center as a caller",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			caller.Run()
		},
	},
	&cobra.Command{
		Use:   "http",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			http.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
