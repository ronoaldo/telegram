package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	parsePayloadDesc(os.Stdin)
}

func parsePayloadDesc(in io.Reader) {
	fmt.Printf("package telegram\n\n")
	scaner := bufio.NewScanner(in)
	for scaner.Scan() {
		line := scaner.Text()
		
		var field, ftype, help string
		
		parts := strings.Split(line, "\t")
		switch {
		case len(parts) == 1:
			name := parts[0]
			if strings.TrimSpace(name) == "" {
				fmt.Printf("}\n\n")
			} else {
				// We are opening a type
				fmt.Printf("type %s struct {\n", parts[0])
			}
			continue
		case len(parts) == 3:
			field, ftype, help = parts[0], parts[1], parts[2]
			fmt.Printf("\t// %s\n", help)
			fmt.Printf("\t%s %s `json:\"%s,omitempty\"`\n", goFieldName(field), goFieldType(ftype), field)
		}
	}
}

func goFieldType(ftype string) string {
	switch ftype {
	case "Integer":
		return "int64"
	case "String":
		return "string"
	case "Float":
		return "float64"
	case "Boolean","True","False":
		return "bool"
	default:
		if strings.HasPrefix(ftype, "Array of Array of") {
			return strings.Replace(ftype, "Array of Array of ", "[][]*", -1)
		} else if strings.HasPrefix(ftype, "Array of") {
			return strings.Replace(ftype, "Array of ", "[]*", -1)
		} else {
			return "*" + ftype
		}
	}
}

func goFieldName(field string) string {
	names := strings.Split(field, "_")
	var buff bytes.Buffer
	for _, n := range names {
		buff.WriteString(strings.Title(n))
	}
	return buff.String()
}