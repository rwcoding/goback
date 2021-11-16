package logger

import (
	"fmt"
	"github.com/rwcoding/goback/pkg/config"
	"go.uber.org/zap"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var sugar *zap.SugaredLogger

func newConfig() zap.Config {
	outputPaths := []string{"stderr"}
	if config.IsDev() {
		outputPaths = []string{"my://"}
	}
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}
}

type mySink struct {
	url     *url.URL
	ymd     int
	writers sync.Map
}

func (s *mySink) Sync() error {
	s.writers.Range(func(key, value interface{}) bool {
		value.(*os.File).Sync()
		return true
	})
	return nil
}

func (s *mySink) Write(p []byte) (n int, err error) {
	//time.LoadLocation("Asia/Shanghai")
	date := time.Now().Format("20060102")
	f := strings.ReplaceAll(config.Log(), "{ymd}", date)
	ymd, _ := strconv.Atoi(date)
	if ymd != s.ymd {
		s.gc()
	}
	v, ok := s.writers.Load(ymd)
	if !ok {
		v, _ = os.OpenFile(f, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		s.writers.Store(ymd, v)
	}
	return v.(io.Writer).Write(p)
}

func (s *mySink) Close() error {
	fmt.Println("close writer")
	s.writers.Range(func(key, value interface{}) bool {
		value.(io.Closer).Close()
		s.writers.Delete(key)
		return true
	})
	return nil
}

func (s *mySink) gc() {
	s.writers.Range(func(key, value interface{}) bool {
		if key.(int) < s.ymd {
			value.(io.Closer).Close()
			s.writers.Delete(key)
		}
		return true
	})
}

func newZap() {
	if err := zap.RegisterSink("my", func(url *url.URL) (zap.Sink, error) {
		return &mySink{url: url}, nil
	}); err != nil {
		log.Fatal(err)
	}
	logger, _ := newConfig().Build()
	defer logger.Sync()
	sugar = logger.Sugar()
}
