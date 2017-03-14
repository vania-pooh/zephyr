package main

import (
	"fmt"
	. "github.com/aerokube/zephyr/core"
	"github.com/aerokube/zephyr/reader"
	"github.com/aerokube/zephyr/writer"
)

var (
	readers = make(map[string]Reader)
	writers = make(map[string]Writer)
)

func init() {
	//Available readers
	AddReader("selenoid", &reader.SelenoidReader{})

	//Available writers
	AddWriter("graphite", &writer.GraphiteWriter{})
}

func AddReader(name string, reader Reader) {
	readers[name] = reader
}

func GetReader(name string) (Reader, error) {
	if rd, ok := readers[name]; ok {
		return rd, nil
	}
	return nil, fmt.Errorf("Reader %s does not exist", name)
}

func AddWriter(name string, writer Writer) {
	writers[name] = writer
}

func GetWriter(name string) (Writer, error) {
	if wr, ok := writers[name]; ok {
		return wr, nil
	}
	return nil, fmt.Errorf("Writer %s does not exist", name)
}
