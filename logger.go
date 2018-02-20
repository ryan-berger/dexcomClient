package dexcomClient

import "fmt"

type logger interface {
	Debug(string)
	Dev(string)
	Log(string)
}

type defaultLogger struct {
	config *Config
}

func (logger *defaultLogger) Log(string string) {
	if logger.config.Logging {
		fmt.Println("Log: ", string)
	}
}

func (logger *defaultLogger) Debug(string string) {
	if logger.config.IsDebug {
		fmt.Println("Debug: ", string)
	}
}

func (logger *defaultLogger) Dev(string string) {
	if logger.config.IsDev {
		fmt.Println("Dev: ", string)
	}
}
