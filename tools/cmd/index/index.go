package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/odo-devfiles/registry/tools/types"
	"gopkg.in/yaml.v2"
)

// genIndex generate new index from meta.yaml files in dir.
// meta.yaml file is expected to be in dir/<devfiledir>/meta.yaml
func genIndex(dir string) ([]types.MetaIndex, error) {

	var index []types.MetaIndex

	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range dirs {
		if file.IsDir() {
			var meta types.Meta
			var devfile types.Devfile200

			metaFile, err := ioutil.ReadFile(filepath.Join(dir, file.Name(), "meta.yaml"))
			if err != nil {
				return nil, err
			}
			err = yaml.Unmarshal(metaFile, &meta)
			if err != nil {
				return nil, err
			}

			self := fmt.Sprintf("/%s/%s/%s", filepath.Base(dir), file.Name(), "devfile.yaml")

			// Parse the devfile and retrieve its list of projects
			devfileFile, err := ioutil.ReadFile(filepath.Join(dir, file.Name(), "devfile.yaml"))
			if err != nil {
				return nil, err
			}
			err = yaml.Unmarshal(devfileFile, &devfile)
			if err != nil {
				return nil, err
			}

			metaIndex := types.MetaIndex{
				Meta: meta,
				Links: types.Links{
					Self: self,
				},
				Projects: getDevfileProjects(devfile),
			}
			index = append(index, metaIndex)
		}
	}
	return index, nil
}

func main() {
	devfiles := flag.String("devfiles-dir", "", "Directory containing devfiles.")
	output := flag.String("index", "", "Index filaname. This is where the index in JSON format will be saved.")

	flag.Parse()

	if *devfiles == "" {
		log.Fatal("Provide devfile directory.")
	}

	if *output == "" {
		log.Fatal("Provide index file.")
	}

	index, err := genIndex(*devfiles)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		log.Fatal(err)

	}
	err = ioutil.WriteFile(*output, b, 0644)
	if err != nil {
		log.Fatal(err)

	}
}

// getDevfileProjects iterates through a devfiles list of projects and returns a list of their names
func getDevfileProjects(devfile types.Devfile200) []string {
	var projects []string
	for _, project := range devfile.Projects {
		projects = append(projects, project.Name)
	}
	return projects
}
