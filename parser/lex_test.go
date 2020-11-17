package parser

import (
	"encoding/xml"
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
    assert.NoError(t,err)
	assert.NoError(t, err)
	assert.Equal(t,20, len(resp.Products))
	assert.Equal(t,2, len(resp.Payment))
	assert.Equal(t, "R", resp.Type)
	assert.Equal(t, "2", resp.ReceiptID)
	assert.Equal(t, 1400.00, resp.Amount)
	assert.Equal(t, "1", resp.AmountPbp)
	assert.Equal(t, "880", resp.PointsPbp)

}

func TestXml(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/request_R.xml")
	if err != nil {
		t.Fatal(err)
	}
	r := DiscountRequestXMLEnvelope{}
	err = xml.Unmarshal(bts, &r)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t,20,len(r.Data.Products.Item))
	assert.Equal(t,2, len(r.Data.Payment.Item))
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
