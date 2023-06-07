package vuelos

// Vuelo modela un vuelo en nuestro sistema de aeropuerto
type Vuelo interface {

	//Fecha retorna la fecha del vuelo, aerolinea y numero de vuelo
	FechaAerolineaNumero() []string

	//Recorrido devuelve la combinacion del origen-destino. Sino, nil.
	Recorrido() string

	//InfoCompleta devolvera todos los datos de vueloTambién puede devolver error en caso que el votante ya hubiera terminado antes su proceso de votación.
	InfoCompleta() []string

	//PrioridadNumero devolvera la prioridad y el numero de vuelo
	PrioridadNumero() (int, int)
}
