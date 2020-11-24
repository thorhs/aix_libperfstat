package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type Field struct {
	Name  string
	Typ   string
	Descr string
}

type InputData struct {
	Name       string
	ItemName   string
	Id         string
	Labels     []string
	Fields     []Field
	InitStruct string
	InitStmt   string
}

func main() {
	var line string

	if len(os.Args) != 2 {
		fmt.Printf("%s input\n", os.Args[0])
		os.Exit(1)
	}

	base := path.Base(os.Args[1])
	ext := path.Ext(base)
	dsName := strings.TrimSuffix(base, ext)
	dsType := ext[1:]
	dsTemplateName := fmt.Sprintf("%s.template", dsType)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening input file: %s\n", err)
		os.Exit(1)
	}

	output, err := os.Create(fmt.Sprintf("%s.go", dsName))
	if err != nil {
		fmt.Printf("Error opening outptu file: %s\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)

	dsItemName := ""
	dsId := ""
	dsLabels := make([]string, 0)
	dsInitStruct := ""
	dsInitStmt := ""

	// 'total' data sources don't include header nor init line
	if dsType != "total" {
		if !scanner.Scan() {
			fmt.Printf("Unable to read header line of data source\n")
			os.Exit(1)
		}

		line = scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			fmt.Printf("Too few items of header line of data souce, header line was: '%s'\n", line)
			os.Exit(1)
		}

		dsItemName = tokens[0]
		dsId = tokens[1]
		if len(tokens) > 2 {
			dsLabels = tokens[2:]
		}

		if !scanner.Scan() {
			fmt.Printf("Unable to read initialization line of data source\n")
			os.Exit(1)
		}

		line = scanner.Text()
		tokens = strings.SplitN(line, " ", 2)
		if len(tokens) != 2 {
			fmt.Printf("Too few items of initialization line of data souce, header line was: '%s'\n", line)
			os.Exit(1)
		}

		dsInitStruct = tokens[0]
		dsInitStmt = tokens[1]
	}

	fields := make([]Field, 0)

	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ReplaceAll(line, "\t\t", "\t")
		tokens := strings.SplitN(line, "\t", 3)
		if len(tokens) != 3 {
			fmt.Printf("Too few items for data item, line was: '%s'\n", line)
			os.Exit(1)
		}

		fields = append(fields, Field{tokens[1], tokens[0], tokens[2]})
	}

	templ, err := template.New("").Funcs(template.FuncMap{
		"Title":    strings.Title,
		"Camel":    SnakeToCamel,
		"Reserved": Reserved,
	}).ParseFiles(path.Join("templates", dsTemplateName))
	if err != nil {
		fmt.Printf("Error opening template file: %s\n", err)
		os.Exit(1)
	}

	inputData := InputData{
		dsName,
		dsItemName,
		dsId,
		dsLabels,
		fields,
		dsInitStruct,
		dsInitStmt,
	}

	err = templ.ExecuteTemplate(output, dsTemplateName, inputData)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		os.Exit(1)
	}
}

func SnakeToCamel(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

func Reserved(input string) string {
	switch input {
	case "type":
		return "_type"
	default:
		return input
	}
}
