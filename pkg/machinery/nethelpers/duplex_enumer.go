// Code generated by "enumer -type=Duplex -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
)

const (
	_DuplexName_0 = "HalfFull"
	_DuplexName_1 = "Unknown"
)

var (
	_DuplexIndex_0 = [...]uint8{0, 4, 8}
	_DuplexIndex_1 = [...]uint8{0, 7}
)

func (i Duplex) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _DuplexName_0[_DuplexIndex_0[i]:_DuplexIndex_0[i+1]]
	case i == 255:
		return _DuplexName_1
	default:
		return fmt.Sprintf("Duplex(%d)", i)
	}
}

var _DuplexValues = []Duplex{0, 1, 255}

var _DuplexNameToValueMap = map[string]Duplex{
	_DuplexName_0[0:4]: 0,
	_DuplexName_0[4:8]: 1,
	_DuplexName_1[0:7]: 255,
}

// DuplexString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DuplexString(s string) (Duplex, error) {
	if val, ok := _DuplexNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Duplex values", s)
}

// DuplexValues returns all values of the enum
func DuplexValues() []Duplex {
	return _DuplexValues
}

// IsADuplex returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Duplex) IsADuplex() bool {
	for _, v := range _DuplexValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Duplex
func (i Duplex) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Duplex
func (i *Duplex) UnmarshalText(text []byte) error {
	var err error
	*i, err = DuplexString(string(text))
	return err
}
