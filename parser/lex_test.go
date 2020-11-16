package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := RequestMock(bts)

	fmt.Printf("%+v", resp.Data.Payment)
	assert.NoError(t, err)
	assert.Equal(t, "R", resp.Data.Type)
	assert.Equal(t, "1", resp.Data.ReceiptID)
	assert.Equal(t, float64(200), resp.Data.Amount)
	assert.Equal(t, "1", resp.Data.AmountPbp)
	assert.Equal(t, "880", resp.Data.PointsPbp)
	assert.Equal(t, "NQ", resp.Data.PosVersion)
	assert.Equal(t, "10.152.152.79", resp.Data.IpCashDesk)
}

func BenchmarkLexer(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := RequestMock(bts)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXml(b *testing.B) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := DiscountRequestXMLEnvelope{}
			err := xml.Unmarshal(bts, &r)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
