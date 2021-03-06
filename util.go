package main

import (
	"fmt"
	"strconv"
)

const delim = "⛄️"

var index uint64 = 1

// UniqID returns and increments the current index
func UniqID() uint64 {
	res := index
	index++
	return res
}

// FormatID returns a new 'random' string id
func FormatID(id uint64) string {
	return fmt.Sprintf("%03s", strconv.FormatUint(id*46649%6125, 36))
}

// UniqIdf returns a new 'random' string id and increaes the index
func UniqIdf() string {
	return FormatID(UniqID())
}

// MakeWsErr builds a byte slice that the server can send back to the client
// over a websocket connections as an error
func MakeWsErr(err string) []byte {
	return []byte("yolo" + delim + err + "\n")
}
