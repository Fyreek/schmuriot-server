package util

import (
	"fmt"

	"github.com/schmonk.io/schmuriot-server/config"
)

// LogToConsole writes provided data to console if server is in debug mode
func LogToConsole(a ...interface{}) {
	if config.Config.Server.Debug {
		fmt.Println(a...)
	}
}
