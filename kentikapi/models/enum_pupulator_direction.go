// Code generated by "enumer -output enum_pupulator_direction.go -type PopulatorDirection"; DO NOT EDIT.

//
package models

import (
	"fmt"
)

const _PopulatorDirectionName = "SRCDSTEITHER"

var _PopulatorDirectionIndex = [...]uint8{0, 3, 6, 12}

func (i PopulatorDirection) String() string {
	if i < 0 || i >= PopulatorDirection(len(_PopulatorDirectionIndex)-1) {
		return fmt.Sprintf("PopulatorDirection(%d)", i)
	}
	return _PopulatorDirectionName[_PopulatorDirectionIndex[i]:_PopulatorDirectionIndex[i+1]]
}

var _PopulatorDirectionValues = []PopulatorDirection{0, 1, 2}

var _PopulatorDirectionNameToValueMap = map[string]PopulatorDirection{
	_PopulatorDirectionName[0:3]:  0,
	_PopulatorDirectionName[3:6]:  1,
	_PopulatorDirectionName[6:12]: 2,
}

// PopulatorDirectionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PopulatorDirectionString(s string) (PopulatorDirection, error) {
	if val, ok := _PopulatorDirectionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PopulatorDirection values", s)
}

// PopulatorDirectionValues returns all values of the enum
func PopulatorDirectionValues() []PopulatorDirection {
	return _PopulatorDirectionValues
}

// IsAPopulatorDirection returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PopulatorDirection) IsAPopulatorDirection() bool {
	for _, v := range _PopulatorDirectionValues {
		if i == v {
			return true
		}
	}
	return false
}
