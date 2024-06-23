package utils

import (
	"fmt"
	"streamer/internal/constants"
	"strings"
)

func ReadArgs(args []string) map[string]string {
	argMap := make(map[string]string)

	for _, arg := range args {
		key, value, err := parseArg(arg)

		if nil != err {
			panic(err)
		}

		argMap[key] = value
	}

	return argMap
}

func parseArg(arg string) (string, string, error) {
	splitArg := strings.Split(arg, constants.DELIMITER)
	if len(splitArg) < 2 {
		return "", "", constants.NewArgParseError(fmt.Sprintf("Invalid argument format: %v", arg))
	}

	value := splitArg[1]
	keyList := strings.Split(splitArg[0], constants.PREFIX)
	if len(keyList) < 2 || len(keyList[1]) == 0 {
		return "", "", constants.NewArgParseError(fmt.Sprintf("Invalid argument format: %v", arg))
	}

	return keyList[1], value, nil
}
