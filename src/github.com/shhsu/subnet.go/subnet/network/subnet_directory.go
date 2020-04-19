package network

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/shhsu/subnet.go/subnet/structure"
)

var createTrie = structure.NewBasicBinaryPrefixTree

const segmentMaskHigh = uint32(255 << 24)

func ParseIPv4(str string) (structure.PrefixKey, error) {
	segmentMask := segmentMaskHigh
	strCursor := 0
	key := structure.PrefixKey(0)
	for segmentShift := uint32(24); segmentMask > 0; segmentShift -= 8 {
		dotIndex := strings.IndexRune(str[strCursor:], '.')
		var segmentString string
		if dotIndex == -1 {
			if segmentShift == 0 {
				segmentString = str[strCursor:]
			} else {
				return 0, errors.Errorf("missing Segment for IP string %s", str)
			}
		} else {

			segmentString = str[strCursor : strCursor+dotIndex]
		}
		segmentValue, err := strconv.Atoi(segmentString)
		if err != nil {
			return 0, errors.Wrapf(err, "error parsing IP string %s, segment %s", str, segmentString)
		}
		key |= structure.PrefixKey(uint32(segmentValue)<<segmentShift | segmentMask)
		segmentMask >>= 8
	}
	return key, nil
}

func ParseCIDR(cidr string) (ip structure.PrefixKey, depth int, err error) {
	slashIndex := strings.IndexRune(cidr, '/')
	if slashIndex <= 0 {
		err = errors.Errorf("error parsing cidr string %s, missing or misplaced slash", cidr)
		return
	}
	if ip, err = ParseIPv4(cidr[0:slashIndex]); err != nil {
		return
	}
	if depth, err = strconv.Atoi(cidr[slashIndex+1:]); err != nil {
		err = errors.Wrapf(err, "error parsing cidr string %s", cidr)
		return
	}
	return
}

type SubnetDirectory struct {
	trie structure.BinaryPrefixTree
}

func NewSubnetDirectory() *SubnetDirectory {
	return &SubnetDirectory{
		trie: createTrie(),
	}
}

func (dir *SubnetDirectory) AddSubnet(cidrString string, label interface{}) (interface{}, error) {
	if ip, depth, err := ParseCIDR(cidrString); err != nil {
		return nil, err
	} else {
		return dir.trie.AddOrReplace(ip, depth, label), nil
	}
}

func (dir *SubnetDirectory) GetSubnet(ipV4 string) (interface{}, error) {
	if ip, err := ParseIPv4(ipV4); err != nil {
		return nil, err
	} else {
		return dir.trie.Get(ip), nil
	}
}
