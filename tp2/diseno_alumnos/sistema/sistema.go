package sistema

type Sistema interface {

	//agregararchivo procesa de forma completa un archivo csv
	AgregarArchivo(parametro string) string
	//Tablero muestra los K vuelos ordenados por la fecha enviada
	Tablero(parametros []string) string
	//InfoVuelo devuelve toda la informacion sobre un vuelo
	InfoVuelo(parametros []string) string
	//muerstra los codigos de los k vuelos que tienen mayor prioridad
	PrioridadVuelos(string) string
	//siguientevuelo busca el siguiente vuelo entre los dos destinos
	SiguienteVuelo(parametros []string) string
	//borrar elimina todos los vuelos entre las fechas de los parametros
	Borrar(parametros []string) string
}
