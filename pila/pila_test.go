package pila_test

import (
	TDAPila "tdas/pila"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Pruebas con Pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.EqualValues(t, true, pila.EstaVacia())

}

func TestPilaEsLifo(t *testing.T) {
	t.Log("Pruebas con una pila de enteros mantiene la invariante de pila (LIFO)")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	// Si el elemento se apiló correctamente, debe ser "3"
	require.EqualValues(t, 3, pila.VerTope())
	pila.Apilar(2)
	//Apilar en una pila que no esta vacia modifica el tope
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(1)
	// Desapilar devuelve el tope
	require.EqualValues(t, 1, pila.Desapilar())
	//Desapilar modifica el tope
	require.EqualValues(t, 2, pila.VerTope())
}

func TestDeVolumen(t *testing.T) {
	t.Log("Prueba con una pila de enteros de varios elementos que implica redimension")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 100000; i++ {
		pila.Apilar(i)
	}
	j := 100000
	for !pila.EstaVacia() {
		require.EqualValues(t, j, pila.VerTope())
		pila.Desapilar()
		j--
	}
	// Una pila a la cual se le desapilan todos sus elementos se comporta como recien creada
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.NotNil(t, pila)
}

func TestPilaDeStrings(t *testing.T) {
	t.Log("Pruebas con una pila de strings")
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Esto")
	pila.Apilar("Es")
	pila.Apilar("Una")
	pila.Apilar("Pila")
	pila.Apilar("En")
	pila.Apilar("Go")
	pila.Apilar("De")
	pila.Apilar("Algo")
	pila.Apilar("2")
	arr := []string{"Esto", "Es", "Una", "Pila", "En", "Go", "De", "Algo", "2"}
	i := len(arr) - 1
	for !pila.EstaVacia() {
		require.EqualValues(t, arr[i], pila.VerTope())
		i--
		pila.Desapilar()
	}
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.NotNil(t, pila)

	s := "La pila conserva un estado valido despues de vaciarla"
	pila.Apilar(s)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, s, pila.VerTope())
	require.NotNil(t, pila)
}

func TestPilaDeFloats(t *testing.T) {
	t.Log("Pruebas con una pila de floats")
	pila := TDAPila.CrearPilaDinamica[float64]()
	for i := 0.0; i < 10000.0; i += 0.5 {
		pila.Apilar(i)
	}
	j := 9999.5
	for !pila.EstaVacia() {
		require.EqualValues(t, j, pila.VerTope())
		j -= 0.5
		pila.Desapilar()
	}
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.NotNil(t, pila)

	n := 0.5
	pila.Apilar(n)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, n, pila.VerTope())
	require.NotNil(t, pila)

}
