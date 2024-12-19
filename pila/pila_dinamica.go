package pila


const (
	_CAPACIDAD_INICIAL = 10
	_FACTOR_REDIMENSION = 2
	_CARGA_MINIMA = _FACTOR_REDIMENSION * _FACTOR_REDIMENSION
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	return pila
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	p.validar_pila_vacia()
	return p.datos[p.cantidad-1]

}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * _FACTOR_REDIMENSION)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	p.validar_pila_vacia()
	elem := p.datos[p.cantidad-1]
	if cap(p.datos) > _CAPACIDAD_INICIAL && p.cantidad <= cap(p.datos) / _CARGA_MINIMA {
		p.redimensionar(cap(p.datos) / _FACTOR_REDIMENSION)
	}
	p.cantidad--
	return elem
}

func (p *pilaDinamica[T]) redimensionar(cap int) {
	nuevos_datos := make([]T, cap)
	copy(nuevos_datos, p.datos)
	p.datos = nuevos_datos
}

func (p *pilaDinamica[T]) validar_pila_vacia() {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
}
