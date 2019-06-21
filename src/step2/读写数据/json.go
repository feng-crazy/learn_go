package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
	Size int64
	//TMap map[string]int
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	tmpMap := make(map[string]int)
	tmpMap["0.25"] = 2
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none", 1024 }
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	//js, _ := json.MarshalforHTML()
	fmt.Printf("JSON format: %s\n", js)
	fmt.Printf("%T\n", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("","	")
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	file.Sync()
	offsetv, offerr := file.Seek(0,0)
	if offerr != nil {
		log.Println("Error in encoding json")
	}
	fmt.Println(offsetv, offerr)

	dec := json.NewDecoder(file)
	//t, err := dec.Token()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Token %T: %v\n", t, t)
	var dcv VCard
	err = dec.Decode(&dcv)
	if err != nil {
		log.Println("Error in encoding json")
	}
	fmt.Printf("dcv %T: %v\n", dcv, dcv)
	fmt.Println(dcv.FirstName, dcv.LastName)

	var unm VCard
	err = json.Unmarshal(js, &unm)
		if err != nil {
		log.Println("Error in encoding json")
	}
	fmt.Printf("unm %T: %v\n", unm, unm)

}