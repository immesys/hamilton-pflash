package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/immesys/bw2/crypto"
	"github.com/immesys/bw2/objects"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: inspectent sk16/vk16/sk64/vk64 file\n")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Printf("could not open file: %v\n", err)
		os.Exit(1)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
		os.Exit(1)
	}
	if contents[0] != objects.ROEntityWKey {
		fmt.Printf("file is not a signing entity")
		os.Exit(1)
	}
	ro, err := objects.LoadRoutingObject(int(contents[0]), contents[1:])
	if err != nil {
		fmt.Printf("could not load entity: %v\n", err)
		os.Exit(1)
	}
	ent := ro.(*objects.Entity)
	switch os.Args[1] {
	case "sk16":
		fmt.Printf("%064x\n", ent.GetSK())
	case "sk64":
		fmt.Printf("%s\n", crypto.FmtKey(ent.GetSK()))
	case "vk16":
		fmt.Printf("%064x\n", ent.GetVK())
	case "vk64":
		fmt.Printf("%s\n", crypto.FmtKey(ent.GetVK()))
	default:
		fmt.Printf("unknown field %v\n", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
