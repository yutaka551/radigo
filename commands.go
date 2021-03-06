package radigo

import (
	"os"

	"github.com/mitchellh/cli"
)

var Ui cli.Ui

const (
	infoPrefix  = "INFO: "
	warnPrefix  = "WARN: "
	errorPrefix = "ERROR: "
)

func init() {
	Ui = &cli.PrefixedUi{
		InfoPrefix:  infoPrefix,
		WarnPrefix:  warnPrefix,
		ErrorPrefix: errorPrefix,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
}

func AreaCommandFactory() (cli.Command, error) {
	return &areaCommand{
		ui: Ui,
	}, nil
}

func RecCommandFactory() (cli.Command, error) {
	return &recCommand{
		ui: Ui,
	}, nil
}

func RecLiveCommandFactory() (cli.Command, error) {
	return &recLiveCommand{
		ui: Ui,
	}, nil
}

func BrowseCommandFactory() (cli.Command, error) {
	return &browseCommand{
		ui: Ui,
	}, nil
}

func BrowseLiveCommandFactory() (cli.Command, error) {
	return &browseLiveCommand{
		ui: Ui,
	}, nil
}
