package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//func TestLex(t *testing.T) {
//	b, err := ioutil.ReadFile("../fixtures/request_B.xml")
//	if err != nil {
//		t.Fatal(err)
//	}
//	_, items := Lex("mylexer", b)
//	for res := range items {
//		fmt.Println(string(res.Value), res.Kind)
//	}
//}

func TestLex2(t *testing.T) {
	b, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	res := GetNewXml(b, []byte("type"))
	fmt.Println(string(res))
}

func TestOld(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	elements := make(map[string]interface{})
	str := GetFieldByNameFromXML(bts, "type")
	elements["type"] = str
	t.Log(str)
}

func BenchmarkOld(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			elements := make(map[string]interface{})
			str := GetFieldByNameFromXML(bts, "type")
			elements["type"] = str
		}
	})

}

func BenchmarkLex(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			elements := make(map[string]interface{})
			res:=GetNewXml(bts, []byte("type"))
			elements["type"] = res
		}
	})
}
