package core

import (
	"time"
	"fmt"
)

var (
	readers = make(map[string] Reader)
	writers = make(map[string] Writer)
)

type Data []Metric

type Metric struct {
	Key string
	Value string
	Timestamp time.Time
}

type Configurable interface {
	Configure(settings Settings) error
}

type Reader interface {
	Configurable
	Read() (*Data, error)
}

type Writer interface {
	Configurable
	Write(data *Data) error
}

func AddReader(name string, reader Reader) {
	readers[name] = reader
}

func GetReader(name string) (Reader, error) {
	if reader, ok := readers[name]; ok {
		return reader, nil
	}
	return nil, fmt.Errorf("Reader %s does not exist", name)
}

func AddWriter(name string, writer Writer) {
	writers[name] = writer
}

func GetWriter(name string) (Writer, error) {
	if writer, ok := writers[name]; ok {
		return writer, nil
	}
	return nil, fmt.Errorf("Writer %s does not exist", name)
}
