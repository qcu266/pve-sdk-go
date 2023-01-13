package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/qcu266/pve-sdk-go/tools/gen-api/gen"
)

var schemaFile, outputDir string

func main() {

	flag.StringVar(&schemaFile, "schema-file", "schema/apidata.js",
		"specific pve api schema file. https://github.com/proxmox/pve-docs/raw/master/api-viewer/apidata.js")
	flag.StringVar(&outputDir, "output-dir", "pve",
		"specific pve api client code output dir")

	flag.Parse()

	if schemaFile == "" {
		fmt.Println("schema-file not specific.")
		os.Exit(1)
	}

	apiSchemas, err := gen.ParseSchemaFile(schemaFile)
	if err != nil {
		fmt.Println("parse api schema file failed. err: ", err)
		os.Exit(1)
	}

	err = gen.GenApis(outputDir, apiSchemas)
	if err != nil {
		fmt.Println("gen api failed. err: ", err)
		os.Exit(1)
	}

}
