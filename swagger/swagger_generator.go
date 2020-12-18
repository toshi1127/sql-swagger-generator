package swagger

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ikaiguang/go-sqlparser"
	"github.com/pkg/errors"
)

type (
	ServiceInfo struct {
		Name   string
		Host   string
		Prefix string
	}

	Generator struct {
		ServiceInfo ServiceInfo
		Resources   map[string]Resource
		Queries     []string
	}

	TemplateData struct {
		ServiceInfo ServiceInfo
		Resources   string
		Definitions string
	}

	Resource struct {
		Title      string
		Path       string
		Definition Definition
		Get        Request
	}

	Request struct {
		Params []RequestParam
	}

	RequestParam struct {
		In          string
		Name        string
		Required    bool
		Type        FieldType
		Description string
	}

	Definition struct {
		Name   string
		Fields []DefinitionField
		Ref    string
	}

	DefinitionField struct {
		Name       string
		Type       FieldType
		IsNullable sqlparser.BoolVal
	}

	FieldType struct {
		Name            string
		ExtraProperties map[string]string
	}
)

func getTypeFromSqlColumnType(col string, enumValues []string) FieldType {
	switch col {
	case "varchar":
		return FieldType{"string", nil}
	case "text":
		return FieldType{"string", nil}
	case "char":
		return FieldType{"string", nil}
	case "datetime":
		return FieldType{"string", map[string]string{
			"format": "date-time",
		}}
	case "date":
		return FieldType{"string", map[string]string{
			"format": "date",
		}}
	case "timestamp":
		return FieldType{"string", map[string]string{
			"format": "date-time",
		}}
	case "tinyint":
		return FieldType{"integer", map[string]string{
			"format": "int64",
		}}
	case "smallint":
		return FieldType{"integer", map[string]string{
			"format": "int64",
		}}
	case "int":
		return FieldType{"integer", map[string]string{
			"format": "int64",
		}}
	case "bigint":
		return FieldType{"integer", map[string]string{
			"format": "int64",
		}}
	case "enum":
		return FieldType{"string", map[string]string{
			"enum": strings.Replace("\n      - "+strings.Join(enumValues, " \n      - "), "'", "", -1),
		}}
	}

	fmt.Printf("%v\n", col)
	return FieldType{col + " [unsupported by swagen]", nil}
}

func (g *Generator) Generate(outputDirPath string) (string, error) {
	resources, err := g.generateResources(outputDirPath)
	if err != nil {
		return "", err
	}

	fullTempl, err := template.New("resources").Parse(fullTemplate)
	if err != nil {
		return "", errors.WithStack(err)
	}

	resourceTempl, err := template.New("resources").Parse(resourceTemplate)
	if err != nil {
		return "", errors.WithStack(err)
	}

	definitionTempl, err := template.New("definitions").Parse(definitionTemplate)
	if err != nil {
		return "", errors.WithStack(err)
	}

	resourcesBuf := bytes.Buffer{}
	for _, resource := range resources {
		err = resourceTempl.Execute(&resourcesBuf, resource)
		if err != nil {
			return "", errors.WithStack(err)
		}
	}

	definitionsBuf := bytes.Buffer{}
	for _, resource := range resources {
		_, err := os.Stat("./" + outputDirPath + "/models")
		if err != nil {
			err = os.Mkdir("./"+outputDirPath+"/models", 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		f, err := os.Create("./" + outputDirPath + "/models/" + resource.Title + ".yml")
		if err != nil {
			return "", errors.WithStack(err)
		}
		err = definitionTempl.Execute(f, resource.Definition)
		if err != nil {
			return "", errors.WithStack(err)
		}
	}

	data := TemplateData{
		ServiceInfo: g.ServiceInfo,
		Resources:   resourcesBuf.String(),
		Definitions: definitionsBuf.String(),
	}
	fullBuf := bytes.Buffer{}
	err = fullTempl.Execute(&fullBuf, data)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return fullBuf.String(), err
}

func (g *Generator) generateResources(outputDirPath string) ([]Resource, error) {
	var resources []Resource

	for tableName, resource := range g.Resources {
		for _, query := range g.Queries {

			tokenizer := sqlparser.NewStringTokenizer(query)
			stat, err := sqlparser.ParseNext(tokenizer)
			if err != nil {
				return nil, err
			}

			ddl, ok := stat.(*sqlparser.DDL)
			if !ok {
				return nil, errors.New("SQL parsing failed.")
			}

			if ddl.NewName.Name.CompliantName() == tableName {
				fields := make([]DefinitionField, 0, len(ddl.TableSpec.Columns))

				if ddl.Action != "create" {
					return nil, errors.New("Contains SQL other than the Create statement")
				}

				for _, col := range ddl.TableSpec.Columns {
					fields = append(fields, DefinitionField{
						Name:       col.Name.String(),
						Type:       getTypeFromSqlColumnType(col.Type.Type, col.Type.EnumValues),
						IsNullable: !col.Type.NotNull,
					})
				}

				resource.Path = strings.ReplaceAll(tableName, "_", "-")
				resource.Definition.Fields = fields
				resource.Definition.Ref = outputDirPath
				resources = append(resources, resource)
			}
		}
	}

	return resources, nil
}
