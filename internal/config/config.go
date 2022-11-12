package config

const (
	DefaultLogLevel       = "info"
	DefaultInputFilePath  = "input.txt"
	DefaultOutPutFilePath = "output.txt"
)

type AppConfig struct {
	LogLevel       string
	InputFilePath  string
	OutputFilePath string
}
