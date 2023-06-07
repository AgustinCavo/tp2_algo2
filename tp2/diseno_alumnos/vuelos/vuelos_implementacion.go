package vuelos

import "strconv"

// Vuelo tiene guardada la información de un vuelo.
type vueloImplementacion struct {
	numero     string
	aerolinea  string
	origen     string
	destino    string
	numeroCola string
	prioridad  string
	fecha      string
	delay      string
	tiempoAire string
	cancelado  bool
}

func CrearVuelo(datos []string, cancelado bool) Vuelo {
	vuelo := new(vueloImplementacion)
	vuelo.numero = datos[0]
	vuelo.aerolinea = datos[1]
	vuelo.origen = datos[2]
	vuelo.destino = datos[3]
	vuelo.numeroCola = datos[4]
	vuelo.prioridad = datos[5]
	vuelo.fecha = datos[6]
	vuelo.delay = datos[7]
	vuelo.tiempoAire = datos[8]
	vuelo.cancelado = cancelado
	return vuelo
}

func (v *vueloImplementacion) FechaAerolineaNumero() []string {
	datos := []string{v.fecha, v.aerolinea, v.numero}
	return datos
}

func (v *vueloImplementacion) Recorrido() string {
	return v.origen + "-" + v.destino
}

func (v *vueloImplementacion) InfoCompleta() []string {
	estado := "no"
	if v.cancelado {
		estado = "si"
	}
	datos := []string{v.numero, v.aerolinea, v.origen, v.destino, v.numeroCola, v.fecha, v.delay, v.tiempoAire, estado}

	return datos
}
func (v *vueloImplementacion) PrioridadNumero() (int, int) {
	numeroP, _ := strconv.Atoi(v.prioridad)
	numeroV, _ := strconv.Atoi(v.numero)
	return numeroP, numeroV
}
