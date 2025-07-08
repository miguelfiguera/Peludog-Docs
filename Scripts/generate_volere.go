package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Requisitos struct {
	Funcionales   []Funcional   `json:"funcionales"`
	NoFuncionales []NoFuncional `json:"noFuncionales"`
}

type Funcional struct {
	ID                  string   `json:"id"`
	Modulo              string   `json:"modulo"`
	Nombre              string   `json:"nombre"`
	Prioridad           string   `json:"prioridad"`
	Descripcion         string   `json:"descripcion"`
	CriteriosAceptacion []string `json:"criteriosAceptacion"`
}

type NoFuncional struct {
	ID          string `json:"id"`
	Categoria   string `json:"categoria"`
	Descripcion string `json:"descripcion"`
}

func main() {
	baseDir, _ := filepath.Abs("../../")
	jsonPath := filepath.Join(baseDir, "JsonRequisitos", "requisitos.json")
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	var reqs Requisitos
	if err := json.Unmarshal(data, &reqs); err != nil {
		log.Fatal(err)
	}

	outDir := filepath.Join(baseDir, "plantillasVolere")
	os.MkdirAll(outDir, os.FileMode(0755))

	// Generate Volere template for each functional requirement
	for _, rf := range reqs.Funcionales {
		filename := filepath.Join(outDir, fmt.Sprintf("%s_%s.md", rf.ID, sanitize(rf.Nombre)))
		content := buildVolereTemplate(rf)
		if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Generadas %d plantillas Volere en %s\n", len(reqs.Funcionales), outDir)
}

func sanitize(s string) string {
	return strings.ReplaceAll(strings.ToLower(strings.ReplaceAll(s, " ", "_")), "/", "-")
}

func buildVolereTemplate(rf Funcional) string {
	return fmt.Sprintf(`# Volere Plantilla: %s

**ID:** %s

**Módulo/Área:** %s

**Descripción:** %s

**Prioridad:** %s

## Criterios de Aceptación
%s

## Fuente
Dra. Génesis / Equipo PeluDog

## Razón/Valor de negocio
Describe cómo este requisito aporta valor al consultorio.

## 
`, rf.Nombre, rf.ID, rf.Modulo, rf.Descripcion, rf.Prioridad, bulletList(rf.CriteriosAceptacion))
}

func bulletList(items []string) string {
	var b strings.Builder
	for _, it := range items {
		b.WriteString("- ")
		b.WriteString(it)
		b.WriteString("\n")
	}
	return b.String()
}
