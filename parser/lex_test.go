package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func GetFieldByNameFromXML(b []byte, name string) (trTyp string) {
	decoder := xml.NewDecoder(bytes.NewBuffer(b))
	inType := false
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			inType = se.Name.Local == name
		case xml.CharData:
			if inType {
				return string(se)
			}
		}
	}

	return trTyp
}

func TestLex(t *testing.T) {
	b, err :=ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	_, items := Lex("mylexer", b)
	elements := make(map[string][]byte)
	for res := range items {
		if res.Kind == TokenPayment {
			fmt.Println("=========")
		}
	   fmt.Println(string(res.Value),res.Kind)
	}
	fmt.Println(string(elements["type"]))
}

func TestOld(t *testing.T) {
	bts, err :=ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	elements := make(map[string]interface{})
	str:=GetFieldByNameFromXML(bts,"type")
	elements["type"]=str
	t.Log(str)
}


func BenchmarkOld(b *testing.B) {
	bts, err :=ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			elements := make(map[string]interface{})
			str:=GetFieldByNameFromXML(bts,"type")
			elements["type"]=str
		}
	})

}

func BenchmarkLex(b *testing.B) {
	bts, err :=ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, items := Lex("mylexer", bts)
			elements := make(map[string][]byte)
			for res := range items {
				if res.Kind == TokenTyp {
					elements["type"] = res.Value
				}
			}

		}
	})




}
