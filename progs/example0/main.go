package main

import (
	"fmt"
	"log"

	"os"

	"github.com/lleo/datafile"
	"github.com/pkg/errors"
)

type mystruct struct {
	String  string
	Integer int
	Boolean bool
}

func main() {
	fmt.Println("example0 filename")

	dataFileName := "example0.df"
	df, err := datafile.Create(dataFileName)
	if err != nil {
		log.Fatal(
			errors.Wrapf(
				err,
				"failed to datafile.Create filename: %s",
				filename,
			),
		)
	}

	loc, err = df.StoreData(&mystruct{"hello", 42, true})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.TempDir())

}
