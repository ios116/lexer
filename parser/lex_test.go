package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
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
		{
			name:   "type",
			result: "R",
		},
		{
			name:   "ip_cash_desk",
			result: "10.152.152.79",
		},
		{
			name:   "pos_version",
			result: "NQ",
		},
	}
    s:=time.Now()
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			item:=item
			_, items := Lex(bts)
			inType := false
			for res := range items {
				if res.Kind == TokenEOF {
					break
				}
				switch {
				case res.Kind == StartElement:
					if 0 == bytes.Compare(res.Value, []byte(item.name)) {
						inType = true
					}
				case res.Kind == CharData:
					if inType {
						if item.result != string(res.Value) {
							t.Fatal("Error")
						}
						return
					}
				}
			}

		})
	}
	t.Log(time.Since(s))

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
			elements := make(map[string]interface{})
			str := GetFieldByNameFromXML(bts, "ip_cash_desk")
			if "10.152.152.79" != str {
				b.Fatal("errrrrrrr")
			}
			elements["type"] = str
		}
	})

}

func BenchmarkNew(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
    v:=	 []byte("ip_cash_desk")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			elements := make(map[string]interface{})
			res := GetNewXml(bts, v)
			if "10.152.152.79" != string(res) {
				b.Fatal("errrrrrrr")
			}
			elements["type"] = res
		}
	})
}

func BenchmarkLex2(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			elements := make(map[string]interface{})
			_, items := Lex( bts)
			inType := false
			for res := range items {
				if res.Kind == TokenEOF {
					break
				}
				switch {
				case res.Kind == StartElement:
					if 0 == bytes.Compare(res.Value, []byte("ip_cash_desk")) {
						inType = true
					}
				case res.Kind == CharData:
					if inType {
						elements["type"] = res.Value
						if "10.152.152.79" != string(res.Value) {
							b.Fatal("eee")
						}
						break
					}
				}
			}
		}

	})

}
