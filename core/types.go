package core

import (
	"time"
)

type Data []Metric

type Metric struct {
	Key       string
	Value     string
	Timestamp time.Time
}

type Reader interface {
	Configure(settings ReaderSettings) error
	Read() (*Data, error)
}

type Writer interface {
	Configure(settings WriterSettings) error
	Write(data *Data) error
}
