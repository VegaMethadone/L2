package arguments

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"telnet/structs"
	"time"
	"unicode"
)

func GetArgs() (*structs.Args, error) {
	args := &structs.Args{
		Host:    "",
		Port:    "",
		Timeout: time.Duration(10) * time.Second,
	}

	if len(os.Args) < 2 {
		return nil, errors.New("not enough arguments")
	}

	for _, value := range os.Args {
		if strings.Contains(value, "--timeout=") {
			newStr := ""
			for _, char := range value {
				if unicode.IsDigit(char) {
					newStr += string(char)
				}
			}
			newStr += string(value[len(value)-1])
			tmpTime, err := parseTimeout(newStr)

			if err != nil {
				return nil, err
			}
			args.Timeout = tmpTime
		} else if len(value) < 6 {
			args.Port = value
		} else {
			args.Host = value
		}
	}

	if args.Host == "" || args.Port == "" {
		return nil, errors.New("host and port are required")
	}

	port, err := strconv.Atoi(args.Port)
	if err != nil {
		return nil, errors.New("invalid port format")
	}
	if port < 1024 || port > 65535 {
		return nil, errors.New("port must be between 1024 and 65535")
	}
	return args, nil
}

func parseTimeout(timeoutStr string) (time.Duration, error) {
	timeNum := []rune{}
	measurement := ""

	for _, char := range timeoutStr {
		if unicode.IsDigit(char) {
			timeNum = append(timeNum, char)
		} else {
			measurement = string(char)
			break
		}
	}

	duration, err := strconv.Atoi(string(timeNum))
	if err != nil {
		return 0, errors.New("incorrect time format: use 3s | 3m")
	}

	switch measurement {
	case "s":
		return time.Second * time.Duration(duration), nil
	case "m":
		return time.Minute * time.Duration(duration), nil
	case "h":
		return time.Hour * time.Duration(duration), nil
	default:
		return 0, errors.New("unsupported time measurement unit: use s, m, or h")
	}
}
