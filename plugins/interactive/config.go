package interactive

import (
	"fmt"

	schematypes "github.com/taskcluster/go-schematypes"
)

type config struct {
	ArtifactPrefix             string `json:"artifactPrefix"`
	ForbidCustomArtifactPrefix bool   `json:"forbidCustomArtifactPrefix"`
	AlwaysEnabled              bool   `json:"alwaysEnabled"`
	DisableShell               bool   `json:"disableShell"`
	DisableDisplay             bool   `json:"disableDisplay"`
	ShellToolURL               string `json:"shellToolUrl"`
	DisplayToolURL             string `json:"displayToolUrl"`
}

var configSchema = schematypes.Object{
	MetaData: schematypes.MetaData{
		Title: "Interactive Plugin",
		Description: `Configuration for the 'interactive' plugin that allows user
      to configure tasks that expose an interactive shell or noVNC sessions.`,
	},
	Properties: schematypes.Properties{
		"artifactPrefix": schematypes.String{
			MetaData: schematypes.MetaData{
				Title: "Artifact Prefix",
				Description: "Prefix that the `sockets.json`, `display.html` and " +
					"`shell.html` should be created under. Defaults to " +
					fmt.Sprintf("`%s`.", defaultArtifactPrefix),
			},
			Pattern:       `^[\x20-.0-\x7e][\x20-\x7e]*/$`,
			MaximumLength: 255,
		},
		"forbidCustomArtifactPrefix": schematypes.Boolean{
			MetaData: schematypes.MetaData{
				Title: "Forbid Custom ArtifactPrefix",
				Description: "Prevent tasks from specifying a custom `artifactPrefix`" +
					" , by default tasks are allowed to overwrite the global setting.",
			},
		},
		"alwaysEnabled": schematypes.Boolean{
			MetaData: schematypes.MetaData{
				Title:       "Always Enabled",
				Description: "If set the interactive plugin will be abled for all tasks.",
			},
		},
		"disableShell": schematypes.Boolean{
			MetaData: schematypes.MetaData{
				Title:       "Disable Shell",
				Description: "If set the interactive shell will be disabled.",
			},
		},
		"disableDisplay": schematypes.Boolean{
			MetaData: schematypes.MetaData{
				Title:       "Disable Display",
				Description: "If set the interactive display will be disabled.",
			},
		},
		"shellToolUrl": schematypes.URI{
			MetaData: schematypes.MetaData{
				Title: "Shell Tool URL",
				Description: `URL to a tool that can take shell socket URL and display
					an interactive shell session. The URL will be given the querystring
					options: 'v=2', 'socketUrl', 'taskId', 'runId'.`,
			},
		},
		"displayToolUrl": schematypes.URI{
			MetaData: schematypes.MetaData{
				Title: "Display Tool URL",
				Description: `URL to a tool that can take display socket, list
					displays and render noVNC session. The URL will be given the
					querystring options: 'v=1', 'socketUrl', 'displaysUrl', 'taskId' and
					'runId'.`,
			},
		},
	},
}
