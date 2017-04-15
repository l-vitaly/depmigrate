package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	path = flag.String("path", "Path to directory", "-path=./")
)

func init() {
	flag.Parse()
}

func main() {
	if *path == "" {
		fmt.Println("Path must be required")
		os.Exit(1)
	}

	lockBytes, err := ioutil.ReadFile(*path + "/lock.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	manifestBytes, err := ioutil.ReadFile(*path + "/manifest.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var lock, manifest map[string]interface{}
	var lockToml, manifestToml bytes.Buffer

	err = json.Unmarshal(lockBytes, &lock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(manifestBytes, &manifest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = toml.NewEncoder(&lockToml).Encode(lock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = toml.NewEncoder(&manifestToml).Encode(manifest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*path+"/Gopkg.lock", lockToml.Bytes(), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*path+"/Gopkg.toml", manifestToml.Bytes(), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
