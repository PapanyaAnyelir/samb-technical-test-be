package utils

import (
	"fmt"
	"time"
)

func GenerateInboundNumber() string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("IN%s", timestamp)
}

func GenerateOutboundNumber() string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("OUT%s", timestamp)
}
