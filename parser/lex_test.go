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
	b, err :=ioutil.ReadFile("../fixtures/request_B.xml")
	if err != nil {
		t.Fatal(err)
	}
	_, items := Lex("mylexer", b)
	for res := range items {
	   fmt.Println(string(res.Value),res.Kind)

	}
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


//func BenchmarkOld(b *testing.B) {
//	bts, err :=ioutil.ReadFile("../fixtures/request_R.xml")
//	if err != nil {
//		b.Fatal(err)
//	}
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			elements := make(map[string]interface{})
//			str:=GetFieldByNameFromXML(bts,"type")
//			elements["type"]=str
//		}
//	})
//
//}
//
//func BenchmarkLex(b *testing.B) {
//	bts, err :=ioutil.ReadFile("../fixtures/request_R.xml")
//	if err != nil {
//		b.Fatal(err)
//	}
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			_, items := Lex("mylexer", bts)
//			elements := make(map[string][]byte)
//			for res := range items {
//				if res.Kind == EndElement {
//					elements["type"] = res.Value
//				}
//			}
//
//		}
//	})
//}

func TestProcessor(t *testing.T) {

	str:=[]byte("hello")
	st:=bytes.HasPrefix(str[1:],[]byte("el"))
	t.Log(st)

	if str[1] == 'e' && str[2] == 'l' {
	  t.Log("YES")
	}
}

func BenchmarkLex2(b *testing.B) {
	str:=[]byte("hello")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
		   bytes.HasPrefix(str[1:],[]byte("el"))
		}
	})

}

func BenchmarkLex3(b *testing.B) {
	str:=[]byte("hello")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if str[1] == 'e' && str[2] == 'l' {

			}
		}
	})

}
