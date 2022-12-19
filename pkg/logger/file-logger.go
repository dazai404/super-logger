package logger

import (
	"os"

	"go.uber.org/zap"
)

type FileLogger struct {
	Logger *zap.Logger
}

func NewFileLogger(filename string) (*FileLogger, error) {
    os.Chdir("logs")
    os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
    c := zap.NewProductionConfig()
    c.OutputPaths = []string{"stdout", filename}
    logger, err := c.Build()
    if err != nil {
        return nil, err
    }
    return &FileLogger{
        Logger: logger,
    }, nil
}
