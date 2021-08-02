package u256

import (
	"math/big"

	"github.com/holiman/uint256"
	"github.com/umbracle/go-web3/utils"
)

type Int struct {
	val *big.Int
}

func New(v interface{}) Int {
	switch val := v.(type) {
	case int:
		return Int{big.NewInt(int64(val))}
	case int64:
		return Int{big.NewInt(int64(val))}
	case uint8:
		return Int{big.NewInt(int64(val))}
	case uint16:
		return Int{big.NewInt(int64(val))}
	case uint32:
		return Int{big.NewInt(int64(val))}
	case uint:
		return Int{big.NewInt(0).SetUint64(uint64(val))}
	case uint64:
		return Int{big.NewInt(0).SetUint64(val)}
	case *big.Int:
		return Int{val}
	case big.Int:
		return Int{&val}
	case uint256.Int:
		return Int{val.ToBig()}
	case *uint256.Int:
		return Int{val.ToBig()}
	case Int:
		return val
	case *Int:
		return *val
	default:
		panic("")
	}
}

func Add(val ...interface{}) Int {
	result := New(0)
	for _, v := range val {
		result = result.Add(New(v))
	}

	return result
}

func Mul(val ...interface{}) Int {
	utils.EnsureTrue(len(val) >= 1)
	result := New(1)
	for _, v := range val {
		result = result.Mul(New(v))
	}

	return result
}

func (self Int) Mul(val Int) Int {
	return Int{big.NewInt(0).Mul(self.val, val.val)}
}

func (self Int) Add(val Int) Int {
	return Int{big.NewInt(0).Add(self.val, val.val)}
}

func Sub(val1, val2 interface{}) Int {
	return New(val1).Sub(New(val2))
}

func (self Int) Sub(val Int) Int {
	return Int{big.NewInt(0).Sub(self.val, val.val)}
}

func (self Int) Div(val Int) Int {
	return Int{big.NewInt(0).Div(self.val, val.val)}
}

func (self Int) Uint64() uint64 {
	utils.EnsureTrue(self.val.IsUint64())
	return self.val.Uint64()
}

func (self Int) ToBigInt() *big.Int {
	return self.val
}

func (a Int) MulUint64(rhs uint64) Int {
	return a.Mul(New(rhs))
}

func (self Int) Exp(val Int) Int {
	return Int{big.NewInt(0).Exp(self.val, val.val, nil)}
}

func (self Int) ExpUint8(val uint8) Int {
	return self.Exp(New(val))
}

func (a Int) DivUint64(rhs uint64) Int {
	return a.Div(New(rhs))
}

func (self Int) IsZero() bool {
	return self.ToBigInt().Sign() == 0
}

func (self Int) LessThan(rhs Int) bool {
	if self.ToBigInt().Cmp(rhs.ToBigInt()) == -1 {
		return true
	}

	return false
}

func (self Int) String() string {
	return self.val.String()
}

func (self Int) ToFixNum(precise uint64) string {
	return toStringByPrecise(self.ToBigInt(), precise)
}

func toStringByPrecise(bigNum *big.Int, precise uint64) string {
	result := ""
	destStr := bigNum.String()
	destLen := uint64(len(destStr))
	if precise >= destLen { // add "0.000..." at former of destStr
		var i uint64 = 0
		prefix := "0."
		for ; i < precise-destLen; i++ {
			prefix += "0"
		}
		result = prefix + destStr
	} else { // add "."
		pointIndex := destLen - precise
		result = destStr[0:pointIndex] + "." + destStr[pointIndex:]
	}
	result = removeZeroAtTail(result)
	return result
}

// delete no need "0" at last of result
func removeZeroAtTail(str string) string {
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	str = str[:i+1]
	// delete "." at last of result
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}
	return str
}
