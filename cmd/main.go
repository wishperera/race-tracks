package main

import (
	"flag"
	defaultLog "log"

	"github.com/wishperera/race-tracks/internal/config"
	"github.com/wishperera/race-tracks/internal/input"
	"github.com/wishperera/race-tracks/internal/models"
	"github.com/wishperera/race-tracks/internal/output"
	"github.com/wishperera/race-tracks/internal/pkg/log"
	"github.com/wishperera/race-tracks/internal/service"
	"github.com/wishperera/race-tracks/internal/util"
)

func main() {
	cfg := config.AppConfig{}

	flag.StringVar(&cfg.LogLevel, "l", config.DefaultLogLevel, "-l <log_level>  log level can be one of trace,info,debug,warn,"+
		"error,fatal")
	flag.StringVar(&cfg.InputFilePath, "i", config.DefaultInputFilePath, "-i <file_path> relative file path "+
		"to input file")
	flag.StringVar(&cfg.OutputFilePath, "o", config.DefaultOutPutFilePath, "-o <file_path> relative file path"+
		"to output file")
	flag.Parse()

	logLevel, err := log.ParseLevelFromString(cfg.LogLevel)
	if err != nil {
		defaultLog.Fatalf("failed to initialize logger due: %s", err)
	}

	logger := log.NewLog(logLevel)

	inputFile, err := util.OpenInputFile(cfg.InputFilePath)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer inputFile.Close()

	outputFile, err := util.OpenOutputFile(cfg.OutputFilePath)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer outputFile.Close()

	inputProvider, err := input.NewReader(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	outputProvider, err := output.NewWriter(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	data, err := inputProvider.ReadInput(inputFile)
	if err != nil {
		logger.Fatal(err.Error())
	}

	results := make([]int, 0)
	for _, v := range data {
		results = append(results, service.FindMinimumHops(v.Start, v.Target, models.NewGrid(v.GridLength, v.GridWidth, v.Obstacles)))
	}

	err = outputProvider.WriteOutput(results, outputFile)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
