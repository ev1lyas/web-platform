package main

import (
	"fmt"
	//"fmt"
	"platform/config"
	"platform/logging"
)

func writeMessage(logger logging.Logger) {
	logger.Info("Hello, Platform")
}
func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
	config := config.DefaultConfig{
		ConfigData: map[string]interface{}{
			"database": map[string]interface{}{
				"host": "localhost",
				"port": 5432,
				"credentials": map[string]interface{}{
					"username": "user",
					"password": "pass",
				},
			},
		},
	}
	res, found := config.Get("database:credentials:username")
	fmt.Println(res, found)

}
