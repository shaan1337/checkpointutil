package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type appConfig struct {
	fileName   string
	readMode   bool
	newValue   uint64
	forceWrite bool
}

var config appConfig

const (
	filenameUsage   = "path of the checkpoint file"
	writeUsage      = "specify this to write to the checkpoint file (requires newValue)"
	newValueUsage   = "the new value of the checkpoint in decimal notation"
	forceWriteUsage = "do not prompt before overwriting a file"
)

func init() {
	flag.StringVar(&config.fileName, "file", "writer.chk", filenameUsage)
	flag.StringVar(&config.fileName, "f", "writer.chk", filenameUsage+" (shorthand)")

	flag.BoolVar(&config.readMode, "write", false, writeUsage)
	flag.BoolVar(&config.readMode, "w", false, writeUsage+" (shorthand)")

	flag.Uint64Var(&config.newValue, "value", 0, newValueUsage)
	flag.Uint64Var(&config.newValue, "v", 0, newValueUsage+" (shorthand)")

	flag.BoolVar(&config.forceWrite, "force", false, forceWriteUsage)
	flag.BoolVar(&config.forceWrite, "y", false, forceWriteUsage+" (shorthand)")
}

func main() {
	flag.Parse()

	checkpointData, err := ioutil.ReadFile(config.fileName)
	if err != nil {
		log.Fatalf("Cannot read file %s - %s", config.fileName, err)
	}

	var checkpointValue uint64

	err = binary.Read(bytes.NewReader(checkpointData), binary.LittleEndian, &checkpointValue)
	if err != nil {
		log.Fatalf("Cannot read a uint64 from the checkpoint data - %s", err)
	}

	fmt.Printf("Checkpoint Value: %d\n", checkpointValue)
}
