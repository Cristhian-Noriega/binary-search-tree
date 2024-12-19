package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN_PILA = 10000
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestLifo(t *testing.T) {
	t.Log("Hacemos pruebas para chequear invariantes")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < _VOLUMEN_PILA; i++ {
		pila.Apilar(i)
		require.EqualValues(t, pila.VerTope(), i)
	}
	for i := _VOLUMEN_PILA - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
}

func TestPilaConSrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("y")
	pila.Apilar("e")
	pila.Apilar("s")
	require.EqualValues(t, "s", pila.VerTope())
	require.EqualValues(t, "s", pila.Desapilar())
	require.EqualValues(t, "e", pila.VerTope())
	require.EqualValues(t, "e", pila.Desapilar())
	require.EqualValues(t, "y", pila.VerTope())
	require.EqualValues(t, "y", pila.Desapilar())
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
}

func TestPilaConFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(10.5)
	require.EqualValues(t, 10.5, pila.VerTope())
	pila.Apilar(100.8)
	require.EqualValues(t, 100.8, pila.VerTope())
	require.EqualValues(t, 100.8, pila.Desapilar())
	require.EqualValues(t, 10.5, pila.Desapilar())
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
}
