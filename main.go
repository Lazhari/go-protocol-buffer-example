package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/lazhari/protobuf-example-go/src/complex"
	enumpb "github.com/lazhari/protobuf-example-go/src/enum_example"
	simplepb "github.com/lazhari/protobuf-example-go/src/simple"
)

func main() {

	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   1,
				Name: "Second Message",
			},
		},
	}

	fmt.Println(cm.String())
}

func doEnum() {
	ep := enumpb.EnumMessage{
		Id:           100,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	fmt.Println(ep.String())
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println("Message as json string \n", smAsString)

	sm2 := &simplepb.SimpleMessage{}

	fromJSON(smAsString, sm2)

	fmt.Println("Successfully created the protocol buffer from as json string \n", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Can't convert to json", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)

	if err != nil {
		log.Fatalln("Can't convert string to json", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content from simple.bin", sm2)
}

func readFromFile(fn string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fn)

	if err != nil {
		log.Fatalln("Something went wrong when we try to open file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("couldn't put the bytes into the protovole buffers struct", err2)
		return err2
	}
	return nil
}

func writeToFile(fn string, pb proto.Message) error {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fn, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple message",
		SampleList: []int32{1, 4, 7, 8},
	}

	sm.Name = "Renamed name"

	return &sm
}
