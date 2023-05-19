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
	if err != nil && !os.IsExist(err) {
		// Crear datos de ejemplo
		people := []Person{
			//{Name: "Cesar", Age: 29},
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

		fmt.Println("El archivo config.json ha sido creado.")
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
	age := flag.Int("age", 0, "Edad de la persona")
	show := flag.Bool("show", false, "Mostrar el contenido de config.json")
	flag.Parse()

	if *show {
		fmt.Println("el archivo config.jsdon tiene")
		for _, p := range people {
			fmt.Printf("Nombre: %s, Edad: %d\n", p.Name, "\n", p.Age)
		}
		return
	}

	if *name != "" && *age != 0 {
		// Agregar nueva persona a la lista
		newPerson := Person{Name: *name, Age: *age}
		people = append(people, newPerson)

		// Actualizar el archivo config.json con la nueva lista de personas
		file, err := os.Create("config.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(people)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Se ha agregado la persona con nombre '%s' y edad '%d' al archivo config.json.\n", *name, *age)
	} else if *name != "" {
		for _, p := range people {
			if p.Name == *name {
				fmt.Println("Los nombres de las personas en el archivo config.json son:")
				for _, p := range people {
					fmt.Println(p.Name)
				}
				return
			}
		}

		fmt.Printf("El nombre '%s' no fue encontrado en el archivo config.json.\n", *name)
	} else if *age != 0 {
		for _, p := range people {
			if p.Age == *age {
				fmt.Println("Las edades de las personas en el archivo config.json son:")
				for _, p := range people {
					fmt.Println(p.Age)
				}
				return
			}
		}

		fmt.Printf("La edad '%d' no fue encontrada en el archivo config.json.\n", *age)
	} else {
		fmt.Println("Debe de porporcionar un valor para la bandera Name y Age")
	}
}
