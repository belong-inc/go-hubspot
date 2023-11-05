package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type columnIndex int

const (
	Name columnIndex = iota
	InternalName
	FieldType
	Description
	GroupName
	FormField
	Options
	ReadOnlyValue
	ReadOnlyDefinition
	Calculated
	ExternalOptions
	Deleted
	HubspotDefined
	CreatedUser
	Usages
)

func main() {
	log.Printf("Running with arges: %s\n", os.Args[1:]) // Without command name.

	if len(os.Args) != 3 {
		log.Fatal("Missing required parameters: <objectName> <csvFilePath>")
	}

	objectName, csvFilePath := os.Args[1], os.Args[2]
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %s", err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %s", err)
	}
	sort.Slice(rows, func(i, j int) bool {
		// Order by InternalName ascending.
		return rows[i][InternalName] < rows[j][InternalName]
	})

	modelFields, internalNames := makeModelAndInternalNames(rows)
	out, err := filepath.Abs(fmt.Sprintf("../../%s_model.go", strings.ToLower(objectName)))
	if err != nil {
		log.Fatalf("Failed to get absolute path: %s", err)
	}

	if err := createFileFromTmpl(out, objectName, modelFields, internalNames); err != nil {
		log.Fatalf("Failed to create file from template: %s", err)
	}

	log.Printf("Generated code in: %s\n", out)
}

func createFileFromTmpl(outPath, objectName string, modelFields, internalNames []string) error {
	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("file open error: %w", err)
	}
	defer f.Close()

	t, err := template.ParseFiles("./model.tmpl")
	if err != nil {
		return fmt.Errorf("file parse error: %w", err)
	}
	data := map[string]interface{}{
		"ObjectName":    objectName,
		"ModelFields":   modelFields,
		"InternalNames": internalNames,
	}
	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}
	return nil
}

func makeModelAndInternalNames(rows [][]string) (model []string, names []string) {
	modelFields, internalNames := make([]string, 0, len(rows)), make([]string, 0, len(rows))
	for i, row := range rows {
		if i != 0 { // Skip header row.
			modelFields = append(modelFields, fmt.Sprintf("%s %s", snakeToCamel(row[InternalName]), switchHsType(row[FieldType])))
			internalNames = append(internalNames, row[InternalName])
		}
	}
	return modelFields, internalNames
}

func snakeToCamel(snakeStr string) string {
	var camelStr string
	for _, part := range strings.Split(snakeStr, "_") {
		camelStr += cases.Title(language.Und).String(part)
	}
	return camelStr
}

func switchHsType(fieldType string) string {
	switch fieldType {
	case "string", "enumeration":
		return "*HsStr"
	case "number":
		return "*HsInt"
	case "bool":
		return "*HsBool"
	case "datetime":
		return "*HsTime"
	default:
		return fieldType
	}
}
