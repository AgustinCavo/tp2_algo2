package sistema

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAErrores "tp2/diseno_alumnos/errores"
	TDAVuelos "tp2/diseno_alumnos/vuelos"
	TDAColaPrioridad "tp2/tdas/cola_prioridad"
	TDADicionario "tp2/tdas/diccionario"
)

type sistemaImplementacion struct {
	tablero          TDADicionario.DiccionarioOrdenado[string, TDAVuelos.Vuelo] //un abb con clave fecha-aero-num y dato vuelo
	vuelosInfo       TDADicionario.Diccionario[string, TDAVuelos.Vuelo]         // un hash con clave numero vuelo
	vecPrioritarios  TDAColaPrioridad.ColaPrioridad[TDAVuelos.Vuelo]            // un heap de prioridades de los vuelos
	siguientesVuelos TDADicionario.Diccionario[string, TDADicionario.DiccionarioOrdenado[string, TDAVuelos.Vuelo]]
} //un hash con clave partida-destino que dato sea un abb ordenado por fecha-aero-num que tenga los vuelos
func comparacionPrioNum(vueloA, vueloB TDAVuelos.Vuelo) int {
	PrioriA := vueloA.PrioridadNumero()
	PrioriB := vueloB.PrioridadNumero()

	if strings.Compare(PrioriA[0], PrioriB[0]) > 0 {
		return 2
	} else if strings.Compare(PrioriA[0], PrioriB[0]) == 0 {
		if strings.Compare(PrioriA[1], PrioriB[1]) > 0 {
			return 2
		} else if strings.Compare(PrioriA[1], PrioriB[1]) == 0 {
			return 0
		}
		return -1
	}
	return -1
}

func CrearSistemaImplementacion() Sistema {
	sistema := new(sistemaImplementacion)

	sistema.tablero = TDADicionario.CrearABB[string, TDAVuelos.Vuelo](strings.Compare)
	sistema.vuelosInfo = TDADicionario.CrearHash[string, TDAVuelos.Vuelo]()
	sistema.vecPrioritarios = TDAColaPrioridad.CrearHeap(comparacionPrioNum)
	sistema.siguientesVuelos = TDADicionario.CrearHash[string, TDADicionario.DiccionarioOrdenado[string, TDAVuelos.Vuelo]]()

	return sistema
}
func abrirArchivo(path string) *os.File {

	file, err := os.Open(path)
	if err != nil {
		e := new(TDAErrores.ErrorLeerArchivo)
		fmt.Println(e.Error())
		os.Exit(1)
		defer file.Close()
	}
	return file
}
func (s *sistemaImplementacion) AgregarArchivo(parametros string) string {

	file := abrirArchivo(parametros)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		datos := strings.Split(linea, ",")
		vuelo := TDAVuelos.CrearVuelo(datos)

		fan := vuelo.FechaAerolineaNumero()
		clave := fan[0] + fan[1] + fan[2]
		s.tablero.Guardar(clave, vuelo)
		s.vuelosInfo.Guardar(fan[2], vuelo)
		s.vecPrioritarios.Encolar(vuelo)
		if s.siguientesVuelos.Pertenece(vuelo.Recorrido()) {
			arbol := s.siguientesVuelos.Obtener(vuelo.Recorrido())
			arbol.Guardar(clave, vuelo)
			s.siguientesVuelos.Guardar(vuelo.Recorrido(), arbol)
		} else {
			arbol := TDADicionario.CrearABB[string, TDAVuelos.Vuelo](strings.Compare)
			arbol.Guardar(clave, vuelo)
			s.siguientesVuelos.Guardar(vuelo.Recorrido(), arbol)
		}

	}

	defer file.Close()
	return "OK"
}
func (s *sistemaImplementacion) Tablero(parametros []string) string {
	vuelo := s.vuelosInfo.Obtener("10")
	return vuelo.Recorrido()
}
func (s *sistemaImplementacion) InfoVuelo(parametros []string) string {
	vuelo := s.vuelosInfo.Obtener("10")
	return vuelo.Recorrido()
}
func (s *sistemaImplementacion) PrioridadVuelos(cantidad string) string {
	return "OK"
}
func (s *sistemaImplementacion) SiguienteVuelo(parametros []string) string {
	return "OK"
}
func (s *sistemaImplementacion) Borrar(parametros []string) string {
	return "OK"
}
