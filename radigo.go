package radigo

import (
	"os"
	"time"

	"github.com/yyoshiki41/go-radiko"
)

const (
	envRardigoHome = "RADIGO_HOME"

	tz             = "Asia/Tokyo"
	datetimeLayout = "20060102150405"
)

var (
	outputDir = "output"

	currentAreaID string
	location      *time.Location
)

func init() {
	var err error

	// If the environment variable RADIGO_HOME is set,
	// override working directory path.
	if e := os.Getenv(envRardigoHome); e != "" {
		outputDir = e
	}

	currentAreaID, err = radiko.AreaID()
	if err != nil {
		panic(err)
	}

	location, err = time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
}

const version = "v0.6.0"

// Version returns the app version.
func Version() string {
	return version
}
