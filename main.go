package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting program....")

	municipios := []string{
		"Vila Flores",
		"Cabixi",
		"Alto Paraiso",
		"Nova Uniao",
		"Bujari",
		"Sao Bernardo do campo",
	}

	notas := []int{
		130, 356, 195,
		240, 891, 340,
	}

	for x := 1; x < 20; x++ {

		filename := fmt.Sprintf("nota_%d.csv", x)

		f, err := os.Create("./Datasets/new_notas/" + filename)

		writer := csv.NewWriter(f)

		msg := []string{"NO_MUNICIPIO_RESIDENCIA;NOTA2"}
		defer writer.Flush()
		writer.Write(msg)

		for y := 1; y < 1000; y++ {
			nMuni := rand.Int() % len(municipios)
			nNotas := rand.Int() % len(notas)
			values := fmt.Sprintf("%s;%d", municipios[nMuni], notas[nNotas])
			msg := []string{values}
			writer.Write(msg)
		}

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		fmt.Println("done file: " + filename)

		time.Sleep(time.Second * 5)
	}

}
