package config

import (
	"fmt"
	"os"
	"strconv"
)

// GetPortForService retrieves the port number for the specified service from the environment variables.
func GetPortForService(serviceName string) (int, error) {
	portStr := os.Getenv("SERVICE_PORT")
	if portStr == "" {
		return 0, fmt.Errorf("no port found for service: %s", serviceName)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("invalid port format for service: %s", serviceName)
	}

	return port, nil
}
