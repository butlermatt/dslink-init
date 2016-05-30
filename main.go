package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/butlermatt/dslink-init/templates"
)

func main() {
	if len(os.Args) != 3 {
		name := path.Base(os.Args[0])
		fmt.Printf("%s is a tool for creating DSA DSLink project skeletons\n\n", name)
		fmt.Printf("Usage:\n\n\t%s <packageName> <nodeName>\n\n", name)
		fmt.Printf("Arguments:\n\n")
		fmt.Println("\tpackageName\tThis will be appended to the dslink name such as: dslink_packageName")
		fmt.Println("\tnodeName\tDefault node name for the link within DGLux.")
		fmt.Println("")
		os.Exit(1)
	}

	pName := strings.TrimSpace(os.Args[1])
	nName := strings.TrimSpace(os.Args[2])

	mkFile("README.md", templates.README, pName, nName)
	mkFile("dslink.json", templates.DSLinkJSON, pName, nName)
	mkFile("pubspec.yaml", templates.PubSpec, pName, nName)

	mkChDir("bin")
	mkFile("run.dart", templates.DartRun, pName, nName)

	mkChDir("..")

	mkChDir("lib")
	mkFile(strings.ToLower(pName+".dart"), "", pName, nName)
	mkFile("models.dart", "", pName, nName)
	mkChDir("src")
	mkChDir("models")
	mkChDir("..")
	mkChDir("nodes")
}

func mkChDir(directory string) {
	if directory != ".." {
		log.Println("Creating directory:", directory)
		err := os.Mkdir(directory, 0755)
		if err != nil {
			if os.IsExist(err) {
				log.Println("Directory exists. Entering")
			} else {
				log.Fatal(err)
			}
		}
	}

	log.Println("Changing to directory:", directory)
	err := os.Chdir(directory)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}

func mkFile(filename, template, pName, nName string) {
	log.Println("Creating file: ", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if template != "" {
		_, err = fmt.Fprintf(file, template, pName, nName)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = file.Chmod(0644)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}
