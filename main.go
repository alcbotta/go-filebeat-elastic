package main

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger = logger.With(zap.String("app", "myapp")).With(zap.String("environment", "psm"))
	count := 0
	for {

		if rand.Float32() > 0.8 {
			logger.Error("oops...something is wrong",
				zap.Int("count", count),
				zap.Error(errors.New("error details")))
		} else {
			logger.Info("everything is fine",
				zap.Int("count", count))
		}
		count++
		time.Sleep(time.Second * 2)
	}
}
