package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDASistema "tp2/diseno_alumnos/sistema"
)

const (
	CANCELADO         string = "0"
	AGREGARARCHIVO    string = "agregar_archivo"
	VERTABLERO        string = "ver_tablero"
	INFOVUELO         string = "info_vuelo"
	PRIORIDADVUELOS   string = "fin-votar"
	SEGUIMIENTOVUELOS string = "siguiente_vuelo"
	BORRAR            string = "borrar"
)

func main() {
	sistema := TDASistema.CrearSistemaImplementacion()
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		// Este codigo lee el comando desde la entrada estándar
		comando, parametros := parsearComando(s.Text())

		// lista de comandos de la aplicacion de voto
		switch comando {
		case AGREGARARCHIVO:
			fmt.Println(sistema.AgregarArchivo(parametros[0]))
		case VERTABLERO:
			fmt.Println(sistema.Tablero(parametros))
		case INFOVUELO:
			fmt.Println(sistema.InfoVuelo(parametros))
		case PRIORIDADVUELOS:
			fmt.Println(sistema.PrioridadVuelos(parametros[0]))
		case SEGUIMIENTOVUELOS:
			fmt.Println(sistema.SiguienteVuelo(parametros))
		case BORRAR:
			fmt.Println(sistema.Borrar(parametros))
		default:
			fmt.Println("Error en comando")
		}
	}
}

func parsearComando(linea string) (string, []string) {
	inputs := strings.Split(linea, " ")
	return inputs[0], inputs[1:]
}
