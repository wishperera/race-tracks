package config

const (
	DefaultLogLevel       = "info"
	DefaultInputFilePath  = "input.txt"
	DefaultOutPutFilePath = "output.txt"
)

// AppConfig : holds the application configuration
type AppConfig struct {
	// LogLevel : log level for the logger, one of trace,debug,info,error,warn,fatal
	LogLevel string
	// InputFilePath : relative path of the input file
	InputFilePath string
	// OutputFilePath : relative path of the output file
	OutputFilePath string
}
