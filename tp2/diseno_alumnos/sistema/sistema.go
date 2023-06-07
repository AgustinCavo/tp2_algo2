package sistema

type Sistema interface {

	//Ingresar ejecuta el comando que permite indicar que llego un votador a la mesa y comprueba que este en el padron
	AgregarArchivo(parametros []string) string
	//Votar permite realizar la votacion al Presidente,Gobernador o Intendente siempre que el dni este
	Tablero(parametros []string) string
	//Deshacer retira la ultima accion realizada por el votante que halla realizado con votar
	InfoVueloDeshacer(parametros []string) string
	//FinVotar termina el proceso de votacion para el votante actual
	PrioridadVuelos(cantidad int)
	//FinalizarVotacion termina el proceso global de votacion y realiza el escrutinio de los votos
	SiguienteVuelo(parametros []string)
	//MostrarResultados muestra los resultados globales de la votacion (Presidente,Gobernador,Intendente)
	Borrar(parametros []string)
}
