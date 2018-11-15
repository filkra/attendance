package csv

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"log"
	"os"
)

type StudentEntry struct { // Our example struct, you can use "-" to ignore a field
	Firstname      string `csv:"Vorname"`
	Lastname    string `csv:"Nachname"`
	Group     string `csv:"Übungsgruppe"`
}

type ExerciseGroups map[string][]string

func Load(srcDir string) ExerciseGroups {
	groups := ExerciseGroups{}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		students := []StudentEntry{}

		file, err := os.Open(fmt.Sprintf("%s/%s", srcDir, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		if err = gocsv.UnmarshalFile(file, &students); err != nil {
			log.Fatal(err)
		}

		for _, entry := range students {
			groups[entry.Group] = append(groups[entry.Group], fmt.Sprintf("%s %s", entry.Firstname, entry.Lastname))
		}
	}

	return groups
}
