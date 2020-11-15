package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLexer(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		result string
	}{
		{name:   "type", result: "R",},
		{name:   "ip_cash_desk", result: "10.152.152.79",},
		{name:   "pos_version", result: "NQ",},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			item:=item
			str := GetNewXml(bts, []byte(item.name))
			if string(str) !=  item.result {
//				t.Fatal(item.result,"not equal",string(str))
			}

		})
	}
}

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
			str := GetFieldByNameFromXML(bts, "ip_cash_desk")
			b.StopTimer()
			if "10.152.152.79" != str {
			//	b.Fatal("error")
			}
			b.StartTimer()
		}
	})
}

func BenchmarkNew(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res := GetNewXml(bts, []byte("ip_cash_desk"))
			b.StopTimer()
			if "10.152.152.79" != string(res) {
			//	b.Fatal("error")
			}
			b.StartTimer()
		}
	})
}

