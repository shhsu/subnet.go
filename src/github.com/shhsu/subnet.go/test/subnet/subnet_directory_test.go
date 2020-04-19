package subnet

import (
	"testing"

	"github.com/shhsu/subnet.go/subnet/network"
	"github.com/shhsu/subnet.go/subnet/util"
	"github.com/stretchr/testify/assert"
)

type tcParseIPv4 struct {
	Input string
	Error error
}

func TestParseIPv4(t *testing.T) {
	tcs := []tcParseIPv4{
		{Input: "0.0.0.0"},
		{Input: "128.0.0.0"},
		{Input: "0.64.0.0"},
		{Input: "0.0.32.0"},
		{Input: "0.0.0.8"},
		{Input: "1.2.3.4"},
		{Input: "255.255.255.255"},
		{Input: "0.123.231.64"},
		{Input: "32.0.53.4"},
		{Input: "22.243.0.114"},
		{Input: "90.99.135.0"},
	}
	for _, tc := range tcs {
		prefix, err := network.ParseIPv4(tc.Input)
		if tc.Error == nil {
			assert.Equal(t, tc.Error, err)
		} else {
			ip := util.ToIP(prefix)
			assert.Equal(t, tc.Input, ip)
		}
	}
}
