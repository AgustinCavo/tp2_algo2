package sistema

import (
	"strings"
	TDAVuelos "tp2/diseno_alumnos/vuelos"
	TDAColaPrioridad "tp2/tdas/cola_prioridad"
	TDADicionario "tp2/tdas/diccionario"
)

type sistemaImplementacion struct {
	tablero          TDADicionario.DiccionarioOrdenado[TDAVuelos.Vuelo, TDAVuelos.Vuelo] //un abb con clave fecha y dato vuelo
	vuelosInfo       TDADicionario.Diccionario[string, TDAVuelos.Vuelo]                  // un hash con clave numero vuelo
	vecPrioritarios  TDAColaPrioridad.ColaPrioridad[TDAVuelos.Vuelo]                     // un heap de prioridades de los vuelos
	siguientesVuelos TDADicionario.Diccionario[string, TDADicionario.DiccionarioOrdenado[string, TDAVuelos.Vuelo]]
} //un hash con clave partida-destino que dato sea un abb ordenado por fecha que tenga los vuelos
func comparacionFechaAeroNum(vueloA, vueloB TDAVuelos.Vuelo) int {
	datosA := vueloA.FechaAerolineaNumero()
	datosB := vueloB.FechaAerolineaNumero()

	for i := 0; i < len(datosA); i++ {
		if strings.Compare(datosA[i], datosB[i]) > 0 {
			return 2
		} else if strings.Compare(datosA[i], datosB[i]) < 0 {
			return -1
		}
	}
	return 0
}
func comparacionPrioNum(vueloA, vueloB TDAVuelos.Vuelo) int {
	PrioriA, NumA := vueloA.PrioridadNumero()
	PrioriB, NumB := vueloB.PrioridadNumero()

	if PrioriA > PrioriB {
		return 2
	} else if PrioriA == PrioriB {
		if NumA > NumB {
			return 2
		} else if NumA == NumB {
			return 0
		}
		return -1
	}
	return -1
}
func CrearSistemaImplementacion() Sistema {
	sistema := new(sistemaImplementacion)

	sistema.tablero = TDADicionario.CrearABB[TDAVuelos.Vuelo, TDAVuelos.Vuelo](comparacionFechaAeroNum)
	sistema.vuelosInfo = TDADicionario.CrearHash[string, TDAVuelos.Vuelo]()
	sistema.vecPrioritarios = TDAColaPrioridad.CrearHeap(comparacionPrioNum)
	sistema.siguientesVuelos = TDADicionario.CrearHash[string, TDADicionario.DiccionarioOrdenado[string, TDAVuelos.Vuelo]]()

	return sistema
}

func (s *sistemaImplementacion) AgregarArchivo(parametros []string) string {
	vuelo := s.vuelosInfo.Obtener("10")
	return vuelo.Recorrido()
}
func (s *sistemaImplementacion) Tablero(parametros []string) string {
	vuelo := s.vuelosInfo.Obtener("10")
	return vuelo.Recorrido()
}
func (s *sistemaImplementacion) InfoVueloDeshacer(parametros []string) string {
	vuelo := s.vuelosInfo.Obtener("10")
	return vuelo.Recorrido()
}
func (s *sistemaImplementacion) PrioridadVuelos(cantidad int) {

}
func (s *sistemaImplementacion) SiguienteVuelo(parametros []string) {

}
func (s *sistemaImplementacion) Borrar(parametros []string) {

}
