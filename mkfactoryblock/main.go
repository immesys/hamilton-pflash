package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

const Version = "4"

func main() {
	//mkfactoryblock 1 uniqueid designator publickey16 privatekey16
	if len(os.Args) != 7 {
		fmt.Printf("usage: mkfactoryblock %s uniqueid:dec designator:hex symmkey:hex publickey:hex privatekey:hex\n", Version)
		os.Exit(1)
	}
	if os.Args[1] != Version {
		fmt.Printf("wrong mkfactoryblock version (we are %s)\n", Version)
		os.Exit(1)
	}
	out := make([]byte, 1024)
	binary.LittleEndian.PutUint64(out[0:], 0x27c83f60f6b6e7c8)
	binary.LittleEndian.PutUint64(out[8:], uint64(time.Now().UnixNano()/1000))
	uniqueid, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	designator, err := strconv.ParseUint(os.Args[3], 16, 64)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	//00:12:6d:07:00:00:serial
	out[16] = 0x00
	out[17] = 0x12
	out[18] = 0x6d
	if designator == 0x8000 {
		out[19] = 0x08
	} else {
		out[19] = 0x07
	}
	out[20] = 0
	out[21] = 0
	out[22] = byte(uniqueid >> 8)
	out[23] = byte(uniqueid & 0xFF)

	out[24] = byte(uniqueid & 0xFF)
	out[25] = byte(uniqueid >> 8)
	//26 pad
	//27 pad
	binary.LittleEndian.PutUint64(out[28:], designator)
	//36 .. 47 pad
	//48 to 64 symm key
	symmkey, err := hex.DecodeString(os.Args[4])
	if err != nil || len(symmkey) != 16 {
		fmt.Println("error with symmkey")
		os.Exit(1)
	}
	copy(out[48:64], symmkey)
	pubkey, err := hex.DecodeString(os.Args[5])
	if err != nil || len(pubkey) != 32 {
		fmt.Println("error with pubkey")
		os.Exit(1)
	}
	copy(out[64:96], pubkey)
	privkey, err := hex.DecodeString(os.Args[6])
	if err != nil || len(privkey) != 32 {
		fmt.Println("error with privkey")
		os.Exit(1)
	}
	copy(out[96:128], privkey)
	fblock, err := os.Create("fblock.bin")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	fblock.Write(out)
	fblock.Close()
	os.Exit(0)
}
