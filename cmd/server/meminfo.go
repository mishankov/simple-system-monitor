package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
)

const MEMINFO_PATH = "/proc/meminfo"

type MemInfo struct {
	MemTotal     int `json:"mem_total"`
	MemFree      int `json:"mem_free"`
	MemAvailable int `json:"mem_available"`
}

func GetMemInfo() (*MemInfo, error) {
	file, err := os.Open(MEMINFO_PATH)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	memInfo := &MemInfo{}
	for _, lineByte := range bytes.Split(data, []byte{10}) {
		if len(lineByte) == 0 {
			continue
		}

		line := string(lineByte)
		lineSplited := strings.Split(line, ":")
		key := lineSplited[0]
		value, err := strconv.Atoi(strings.Split(strings.TrimSpace(lineSplited[1]), " ")[0])
		if err != nil {
			return nil, err
		}

		switch key {
		case "MemTotal":
			memInfo.MemTotal = value
		case "MemFree":
			memInfo.MemFree = value
		case "MemAvailable":
			memInfo.MemAvailable = value
		}
	}

	return memInfo, nil
}
