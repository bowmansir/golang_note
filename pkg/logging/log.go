package logging

import "log"

type Level int

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}
