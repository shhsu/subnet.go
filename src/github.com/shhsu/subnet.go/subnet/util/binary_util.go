package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shhsu/subnet.go/subnet/network"
	"github.com/shhsu/subnet.go/subnet/structure"
)

// ToBinary prints the PrefixKey in binary format
func ToBinary(key structure.PrefixKey) string {
	charArray := make([]rune, structure.PrefixKeyBits)
	for i := uint64(0); i < structure.PrefixKeyBits; i++ {
		charArray[i] = rune('0' + structure.ChildIndex(key, int(i)))
	}
	return string(charArray)
}

// ToIP prints the prefix key as IP Addresses, but if the input is invalid
// there will be extra chunks (purposely left in for debugging convenience)
func ToIP(key network.PrefixKey) string {
	return ToChunks(structure.PrefixKey(key), 4)
}

// ToChunks prints the prefix key in 8 bits chunks
func ToChunks(key structure.PrefixKey, minChunks int) string {
	if minChunks <= 0 {
		panic(fmt.Sprintf("Invalid numChunks %d", minChunks))
	}
	intStack := make([]int, 0)
	for ; key > 0; key /= 256 {
		segment := key % 256
		intStack = append(intStack, int(segment))
	}
	for len(intStack) < minChunks {
		intStack = append(intStack, 0)
	}
	sb := strings.Builder{}
	for i := len(intStack) - 1; i > 0; i-- {
		sb.WriteString(strconv.Itoa(intStack[i]))
		sb.WriteRune('.')
	}
	sb.WriteString(strconv.Itoa(intStack[0]))
	return sb.String()
}
