package wallet

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ZeroCoinType uint32 = 0x80000000
)

type Path struct {
	Purpose      uint32
	CoinType     uint32
	Account      uint32
	Change       uint32
	AddressIndex uint32
}

var DefaultPath = Path{
	Purpose:      2147483692,
	CoinType:     2147483648,
	Account:      2147483648,
	Change:       0,
	AddressIndex: 0,
}

func String2Path(path string) Path {
	p := Path{}
	path = strings.TrimPrefix(path, "m/")
	paths := strings.Split(path, "/")
	if len(paths) != 5 {
		return p
	}
	p.Purpose = Str2Number(paths[0])
	p.CoinType = Str2Number(paths[1])
	p.Account = Str2Number(paths[2])
	p.Change = Str2Number(paths[3])
	p.AddressIndex = Str2Number(paths[4])

	return p
}

func Str2Number(str string) uint32 {
	num64, _ := strconv.ParseInt(strings.TrimSuffix(str, "'"), 10, 64)
	num := uint32(num64)
	if strings.HasSuffix(str, "'") {
		num += ZeroCoinType
	}
	return num
}

func (p Path) String() string {
	paths := []string{"m/"}

	if p.Purpose >= ZeroCoinType {
		num := p.Purpose - ZeroCoinType
		paths = append(paths, fmt.Sprintf("%d'/", num))
	} else {
		paths = append(paths, fmt.Sprintf("%d/", p.Purpose))
	}

	if p.CoinType >= ZeroCoinType {
		num := p.CoinType - ZeroCoinType
		paths = append(paths, fmt.Sprintf("%d'/", num))
	} else {
		paths = append(paths, fmt.Sprintf("%d/", p.CoinType))
	}

	if p.Account >= ZeroCoinType {
		num := p.Account - ZeroCoinType
		paths = append(paths, fmt.Sprintf("%d'/", num))
	} else {
		paths = append(paths, fmt.Sprintf("%d/", p.Account))
	}

	if p.Change >= ZeroCoinType {
		num := p.Change - ZeroCoinType
		paths = append(paths, fmt.Sprintf("%d'/", num))
	} else {
		paths = append(paths, fmt.Sprintf("%d/", p.Change))
	}

	if p.AddressIndex >= ZeroCoinType {
		num := p.AddressIndex - ZeroCoinType
		paths = append(paths, fmt.Sprintf("%d'", num))
	} else {
		paths = append(paths, fmt.Sprintf("%d", p.AddressIndex))
	}

	return strings.Join(paths, "")
}
