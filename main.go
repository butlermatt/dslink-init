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
		fmt.Printf("Usage:\n\n\t%s <package name> <node name>\n\n", name)
		fmt.Printf("Arguements:\n\n")
		fmt.Println("\tpackage name\tThis will be appended to the dslink name such as: dslink-myPackage")
		fmt.Println("\tnode name\tDefault node name for the link within DGLux.")
		fmt.Println("")
		os.Exit(1)
	}

	pName := strings.TrimSpace(os.Args[1])
	nName := strings.TrimSpace(os.Args[2])

	createFile("README.md", templates.README, pName, nName)
	createFile("dslink.json", templates.DSLinkJSON, pName, nName)
	createFile("pubspec.yaml", templates.PubSpec, pName, nName)

	createUseDirectory("bin")
	createFile("run.dart", templates.DartRun, pName, nName)

	createUseDirectory("..")

	createUseDirectory("lib")
	createFile(strings.ToLower(pName+".dart"), "", pName, nName)
	createFile("models.dart", "", pName, nName)
	createUseDirectory("src")
	createUseDirectory("models")
	createUseDirectory("..")
	createUseDirectory("nodes")
}

func createUseDirectory(directory string) {
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

func createFile(filename, template, pName, nName string) {
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