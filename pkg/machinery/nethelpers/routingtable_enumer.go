// Code generated by "enumer -type=RoutingTable -linecomment -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
)

const (
	_RoutingTableName_0 = "unspec"
	_RoutingTableName_1 = "defaultmainlocal"
)

var (
	_RoutingTableIndex_0 = [...]uint8{0, 6}
	_RoutingTableIndex_1 = [...]uint8{0, 7, 11, 16}
)

func (i RoutingTable) String() string {
	switch {
	case i == 0:
		return _RoutingTableName_0
	case 253 <= i && i <= 255:
		i -= 253
		return _RoutingTableName_1[_RoutingTableIndex_1[i]:_RoutingTableIndex_1[i+1]]
	default:
		return fmt.Sprintf("RoutingTable(%d)", i)
	}
}

var _RoutingTableValues = []RoutingTable{0, 253, 254, 255}

var _RoutingTableNameToValueMap = map[string]RoutingTable{
	_RoutingTableName_0[0:6]:   0,
	_RoutingTableName_1[0:7]:   253,
	_RoutingTableName_1[7:11]:  254,
	_RoutingTableName_1[11:16]: 255,
}

// RoutingTableString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoutingTableString(s string) (RoutingTable, error) {
	if val, ok := _RoutingTableNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RoutingTable values", s)
}

// RoutingTableValues returns all values of the enum
func RoutingTableValues() []RoutingTable {
	return _RoutingTableValues
}

// IsARoutingTable returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RoutingTable) IsARoutingTable() bool {
	for _, v := range _RoutingTableValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for RoutingTable
func (i RoutingTable) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for RoutingTable
func (i *RoutingTable) UnmarshalText(text []byte) error {
	var err error
	*i, err = RoutingTableString(string(text))
	return err
}
