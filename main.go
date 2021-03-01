package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tkircsi/proto/person"
	"google.golang.org/protobuf/proto"
)

func main() {
	fname := "../boldi_v3.bin"
	p := person.Person{}
	err := readFromFile(fname, &p)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	fmt.Println(p.String())
}

func readFromFile(fn string, pb proto.Message) error {
	buf, err := os.ReadFile(fn)
	log.Println("protobuf message length: ", len(buf), "bytes")
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buf, pb)
	if err != nil {
		return err
	}
	return nil
}
