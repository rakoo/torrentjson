package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zeebo/bencode"
)

type torrent struct {
	Announce map[string]interface{} `bencode:"announce"`
	Info     map[string]interface{} `bencode:"info"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a file to work with")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Can't open %s\n", os.Args[0])
		os.Exit(1)
	}
	defer f.Close()

	var t map[string]interface{}
	err = bencode.NewDecoder(f).Decode(&t)
	if err != nil {
		fmt.Printf("Can't decode file (%v). Is it a torrent file ?\n", err)
		os.Exit(1)
	}

	err = json.NewEncoder(os.Stdout).Encode(t)
	if err != nil {
		fmt.Printf("Can't write to stdout\n")
		os.Exit(1)
	}
}
