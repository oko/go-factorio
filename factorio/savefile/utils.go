package savefile

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Version struct {
	Major uint16
	Minor uint16
	Patch uint16
	Devel uint16
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Devel)
}

func ReadVersionHeader(r io.Reader) (*Version, error) {
	v := &Version{}
	var err error
	err = binary.Read(r, binary.LittleEndian, &v.Major)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &v.Minor)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &v.Patch)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &v.Devel)
	if err != nil {
		return nil, err
	}
	return v, nil
}
