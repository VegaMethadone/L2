package structs

import (
	"time"
)

type Args struct {
	Host    string
	Port    string
	Timeout time.Duration
}
