package main

import "log"

func logInfo(message string) {
	log.Print(message)
}

func logError(error error, functionName string, traceData string) {
	log.Print("[" + functionName + "] " + " [" + traceData + "] " + error.Error())
}