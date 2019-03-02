package savefile

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/oko/logif"
)

type Savefile struct {
	Version Version
}

func removeSavenamePrefix(path string) string {
	split := strings.Split(path, "/")
	logif.Debugf("%s", split)
	if len(split) == 1 {
		return ""
	}
	return filepath.Join(split[1:]...)
}

func ReadSavefile(path string) (*Savefile, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var ver *Version
	for _, f := range zr.File {
		fn := removeSavenamePrefix(f.Name)
		switch fn {
		case "level.dat":
			logif.Debugf("found level.dat")
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			ver, err = ReadVersionHeader(rc)
		default:
			logif.Debugf("skipping %s", fn)
		}
	}
	if ver == nil {
		if err == nil {
			return nil, fmt.Errorf("%s does not contain a level.dat", path)
		} else {
			return nil, err
		}
	}
	return &Savefile{*ver}, nil
}
