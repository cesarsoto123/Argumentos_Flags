package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Verificar si el archivo existe
	_, err := os.Stat("config.json")
	if err == nil || os.IsExist(err) {
		fmt.Println("El archivo people.json ya existe. No se creará un nuevo archivo.")
	} else {
		// Crear datos de ejemplo
		people := []Person{

			{Name: "Cesar", Age: 29},
		}

		// Abrir archivo para escritura
		file, err := os.Create("config.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Escribir datos en formato JSON en el archivo
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(people)
		if err != nil {
			panic(err)
		}

		fmt.Println("El archivo people.json ha sido creado.")
	}

	// Leer archivo JSON
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var people []Person
	err = json.NewDecoder(file).Decode(&people)
	if err != nil {
		panic(err)
	}

	name := flag.String("name", "", "Nombre de la persona")

	if *name != "" {

		for _, p := range people {
			if p.Name == *name {
				fmt.Printf("El nombre '%s' fue encontrado en el archivo config.json.\n", *name)
				return
			}
		}

		fmt.Printf("El nombre '%s' no fue encontrado en el archivo config.json.\n", *name)
	} else {

		fmt.Println("Los nombres de las personas en el archivo people.json son:")
		for _, p := range people {
			fmt.Println(p.Name)
		}
	}
}
