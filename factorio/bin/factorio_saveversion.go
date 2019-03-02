package main

import (
	"fmt"
	"os"

	"github.com/oko/go-factorio/factorio/savefile"
	"github.com/oko/logif"
	flag "github.com/spf13/pflag"
)

var ()

func main() {
	flag.Parse()
	fns := flag.Args()

	for _, fn := range fns {
		sf, err := savefile.ReadSavefile(fn)
		if err != nil {
			logif.Errorf("failed to read save file %s: %s", fn, err)
			os.Exit(1)
		}
		fmt.Printf("%s %s\n", fn, &sf.Version)
	}
}
