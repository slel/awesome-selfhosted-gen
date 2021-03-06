package exporter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/n8225/awesome-selfhosted-gen/pkg/parse"
)

//ImportYaml imports yaml files into a slice of structs
func ImportYaml(ght string) (l parse.List) {
	files, err := ioutil.ReadDir("./list")
	if err != nil {
		log.Fatal(err)
	}
	var i = 0
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") == true {
			i++
			e := parse.Entry{}

			fFile, err := os.Open("./list/" + f.Name())
			if err != nil {
				fmt.Println(err)
			}
			fbytes, _ := ioutil.ReadAll(fFile)
			err = yaml.Unmarshal(fbytes, &e)
			if err != nil {
				fmt.Println(err)
			}
			e.ID = i
			//fmt.Printf("%d - %s Imported from %s\n", e.ID, e.Name, f.Name())
			g := parse.GetGitdata(e, ght)

			e.Stars = g.Stars
			e.Updated = g.Updated
			e.Gitdata = *g
			l.Entries = append(l.Entries, e)

		}
	}
	fmt.Printf("Added %d of %d yaml files found. There are now %d entries on the list.\n", i, len(files), len(l.Entries))
	l.TagList = parse.MakeTags(l.Entries)
	l.LangList = parse.MakeLangs(l.Entries)
	l.CatList = parse.MakeCats(l.Entries)

	return
}
