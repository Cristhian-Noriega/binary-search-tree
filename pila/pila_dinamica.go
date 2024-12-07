package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	CAPACIDAD_INICIAL    = 8
	INCREMENTO_CAPACIDAD = 2
	CUATRO               = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionarArreglo(cap(pila.datos) * INCREMENTO_CAPACIDAD)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++

}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	if pila.cantidad*CUATRO <= cap(pila.datos) {

		pila.redimensionarArreglo(cap(pila.datos) / INCREMENTO_CAPACIDAD)
	}
	viejo_tope := pila.datos[pila.cantidad-1]
	pila.cantidad--
	return viejo_tope
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionarArreglo(capacidad int) {
	nuevos_datos := make([]T, capacidad)
	copy(nuevos_datos, pila.datos)
	pila.datos = nuevos_datos
}
