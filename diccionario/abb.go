package diccionario

import (
	TDAPila "tdas/pila"
)

type nodo_ab[K comparable, V any] struct {
	izq   *nodo_ab[K, V]
	der   *nodo_ab[K, V]
	clave K
	valor V
}

type abb[K comparable, V any] struct {
	raiz *nodo_ab[K, V]
	cmp  func(K, K) int
	cant int
}

type iteradorAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodo_ab[K, V]]
	desde *K
	hasta *K
	ab    *abb[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cmp: funcion_cmp, cant: 0}
}

func crearNodo[K comparable, V any](clave K, valor V) *nodo_ab[K, V] {
	return &nodo_ab[K, V]{izq: nil, der: nil, clave: clave, valor: valor}
}

func (ab *abb[K, V]) Guardar(clave K, valor V) {
	nodo := ab.buscarNodo(clave, &ab.raiz)
	if *nodo != nil {
		(*(*nodo)).valor = valor
	} else {
		*nodo = crearNodo(clave, valor)
		ab.cant++
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	return *(ab.buscarNodo(clave, &ab.raiz)) != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	nodo := ab.buscarNodo(clave, &ab.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*(*nodo)).valor
}

func (ab *abb[K, V]) Borrar(clave K) V {
	ptr := ab.buscarNodo(clave, &ab.raiz)
	if *ptr == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodo := *ptr
	valor := nodo.valor
	if nodo.izq == nil && nodo.der == nil {
		*ptr = nil
	} else if nodo.izq != nil && nodo.der == nil {
		*ptr = nodo.izq
	} else if nodo.izq == nil && nodo.der != nil {
		*ptr = nodo.der
	} else {
		reemplazantePtr := ab.buscaReemplazante(&nodo.der)
		reemplazante := *reemplazantePtr
		*reemplazantePtr = reemplazante.der
		nodo.valor, nodo.clave = reemplazante.valor, reemplazante.clave
		*reemplazantePtr = reemplazante.der
	}
	ab.cant--
	return valor
}

func (ab *abb[K, V]) Cantidad() int { return ab.cant }

func (ab *abb[K, V]) buscarNodo(clave K, act **nodo_ab[K, V]) **nodo_ab[K, V] {
	if *act == nil {
		return act
	}
	comp := ab.cmp((*(*act)).clave, clave)
	if comp < 0 {
		return ab.buscarNodo(clave, &(*(*act)).der)
	} else if comp > 0 {
		return ab.buscarNodo(clave, &(*(*act)).izq)
	}
	return act
}

func (ab *abb[K, V]) buscaReemplazante(actPtr **nodo_ab[K, V]) **nodo_ab[K, V] {
	act := *actPtr
	if act.izq == nil {
		return actPtr
	}
	return ab.buscaReemplazante(&act.izq)
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) { ab.IterarRango(nil, nil, visitar) }

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterar(ab.raiz, visitar, desde, hasta, ab)
}

func iterar[K comparable, V any](act *nodo_ab[K, V], visitar func(clave K, dato V) bool, desde *K, hasta *K, ab *abb[K, V]) bool {
	if act == nil {
		return true
	}
	if desde != nil && ab.cmp(act.clave, *desde) < 0 {
		return iterar(act.der, visitar, desde, hasta, ab)
	}
	if hasta != nil && ab.cmp(act.clave, *hasta) > 0 {
		return iterar(act.izq, visitar, desde, hasta, ab)
	}
	if !iterar(act.izq, visitar, desde, hasta, ab) {
		return false
	}
	if !visitar(act.clave, act.valor) {
		return false
	}
	return iterar(act.der, visitar, desde, hasta, ab)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

func (iter iteradorAbb[K, V]) HaySiguiente() bool { return !iter.pila.EstaVacia() }

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().valor
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := iter.pila.Desapilar()
	if desapilado.der != nil {
		recorrerRamaIzquierda(desapilado.der, iter)
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodo_ab[K, V]]()
	iteradorabb := iteradorAbb[K, V]{pila: pila, desde: desde, hasta: hasta, ab: ab}
	recorrerRamaIzquierda(ab.raiz, &iteradorabb)
	return &iteradorabb
}

func recorrerRamaIzquierda[K comparable, V any](act *nodo_ab[K, V], iter *iteradorAbb[K, V]) {
	if act == nil {
		return
	}
	if iter.desde != nil && iter.ab.cmp(act.clave, *iter.desde) < 0 {
		recorrerRamaIzquierda(act.der, iter)
		return
	}
	if iter.hasta != nil && iter.ab.cmp(act.clave, *iter.hasta) > 0 {
		recorrerRamaIzquierda(act.izq, iter)
		return
	}
	iter.pila.Apilar(act)
	recorrerRamaIzquierda(act.izq, iter)
}
