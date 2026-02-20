package httpLogger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type HttpLogger struct {
	Logger       *log.Logger
	fatalLogger  *log.Logger
	errorLogger  *log.Logger
	warnLogger   *log.Logger
	infoLogger   *log.Logger
	debugLogger  *log.Logger
	traceLogger  *log.Logger
	logFile      *os.File
	baseDir      string
	baseFilename string
	logLevel     int
}

func NewHttpLogger(outputPath string, logLevel string) *HttpLogger {
	fileName := filepath.Base(outputPath)
	dirName := filepath.Dir(outputPath)

	if err := os.MkdirAll(dirName, 0644); err != nil {
		log.Fatal(err)
	}
	log_file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, log_file)

	logger := log.New(multiWriter, "LOG: ", log.Ldate|log.Ltime)
	fatalLogger := log.New(multiWriter, "FATAL: ", log.Ldate|log.Ltime)
	errorLogger := log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime)
	warnLogger := log.New(multiWriter, "WARN:  ", log.Ldate|log.Ltime)
	infoLogger := log.New(multiWriter, "INFO:  ", log.Ldate|log.Ltime)
	debugLogger := log.New(multiWriter, "DEBUG: ", log.Ldate|log.Ltime)
	traceLogger := log.New(multiWriter, "TRACE: ", log.Ldate|log.Ltime)

	var logLevelInt int

	logLevel = strings.ToUpper(logLevel)

	switch logLevel {
	case FATAL:
		logLevelInt = 0
	case ERROR:
		logLevelInt = 1
	case WARN:
		logLevelInt = 2
	case INFO:
		logLevelInt = 3
	case DEBUG:
		logLevelInt = 4
	case TRACE:
		logLevelInt = 5
	}

	infoLogger.Printf("Logger started with log_level='%s'", logLevel)

	return &HttpLogger{
		Logger:       logger,
		fatalLogger:  fatalLogger,
		errorLogger:  errorLogger,
		warnLogger:   warnLogger,
		infoLogger:   infoLogger,
		debugLogger:  debugLogger,
		traceLogger:  traceLogger,
		logFile:      log_file,
		baseDir:      dirName,
		baseFilename: fileName,
		logLevel:     logLevelInt,
	}
}

func (l *HttpLogger) Close() error {
	if err := l.logFile.Close(); err != nil {
		return err
	}
	return nil
}
