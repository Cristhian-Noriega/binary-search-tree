package diccionario_test

import (
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
	"time"
)

const _VOLUMEN = 30000

func TestAbbVacio(t *testing.T) {
	t.Log("Comprueba que un ABB vacio se comporta como tal")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que el ABB con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())

	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestDatoReemplazado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestABBBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se los borra, revisando que en todo momento " +
		"el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestBorrar2Hijos(t *testing.T) {
	t.Log("Se borra el nodo aunque tenga 2 hijos")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	abb.Guardar(50, 50)
	abb.Guardar(35, 35)
	abb.Guardar(70, 70)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(75, 75)
	abb.Guardar(76, 76)
	abb.Borrar(70)
	require.EqualValues(t, false, abb.Pertenece(70))
	require.EqualValues(t, true, abb.Pertenece(75))
	require.EqualValues(t, true, abb.Pertenece(76))

}

func randomGen(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func TestABBVolumen(t *testing.T) {
	t.Log("Creamos un ABB con muchas elementos y chequeamos que se comporte correctamente")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})

	aux := []int{}
	for i := 0; i < _VOLUMEN; i++ {
		num := randomGen(0, _VOLUMEN)
		if !abb.Pertenece(num) {
			aux = append(aux, num)
		}
		abb.Guardar(num, num)
	}

	for _, num := range aux {
		require.EqualValues(t, true, abb.Pertenece(num))
		require.EqualValues(t, num, abb.Obtener(num))
		require.EqualValues(t, num, abb.Borrar(num))
	}

	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(80))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(80) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(80) })
}

func TestIteradorInOrder(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	clave1 := 1
	clave2 := 2
	clave3 := 3
	clave4 := 4
	clave5 := 5
	clave6 := 6
	clave7 := 7
	abb.Guardar(clave5, clave5)
	abb.Guardar(clave3, clave3)
	abb.Guardar(clave7, clave7)
	abb.Guardar(clave2, clave2)
	abb.Guardar(clave4, clave4)
	abb.Guardar(clave6, clave6)
	abb.Guardar(clave1, clave1)
	inicio, fin := 2, 5
	abb.IterarRango(&inicio, &fin, func(clave, dato int) bool {
		return clave < 4
	})
}

func TestIteradorInternoSinCorte(t *testing.T) {
	t.Log("Chequeamos que si la funcion de visita del iterador interno siempre devuelve true se recorre todo el arbol")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1 := "Hola"
	clave2 := "Algoritmos"
	clave3 := "Arbol"
	valor1 := "Chau"
	valor2 := "Programacion"
	valor3 := "Binario"
	claves := []string{clave2, clave3, clave1}
	valores := []string{valor2, valor3, valor1}
	abb.Guardar(clave1, valor1)
	abb.Guardar(clave2, valor2)
	abb.Guardar(clave3, valor3)
	i := 0
	ptr := &i
	abb.Iterar(func(clave string, valor string) bool {
		require.EqualValues(t, clave, claves[i])
		require.EqualValues(t, valor, valores[i])
		*ptr = *ptr + 1
		return true
	})
}

func TestIteradorExterno(t *testing.T) {
	t.Log("Chequeamos que el Iterador Externo se comporta como un IteradorRango sin un rango definido")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1 := "Hola"
	clave2 := "Algoritmos"
	clave3 := "Arbol"
	valor1 := "Chau"
	valor2 := "Programacion"
	valor3 := "Binario"
	claves := []string{clave2, clave3, clave1}
	valores := []string{valor2, valor3, valor1}
	abb.Guardar(clave1, valor1)
	abb.Guardar(clave2, valor2)
	abb.Guardar(clave3, valor3)
	iter := abb.Iterador()
	i := 0
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, claves[i])
		require.EqualValues(t, valor, valores[i])
		i++
		iter.Siguiente()
	}

}
func TestIteradorExternoSinRango(t *testing.T) {
	t.Log("Validamos que IteradorRango sin un rango definido se comporta como Iterador de ABB")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1 := "Hola"
	clave2 := "Algoritmos"
	clave3 := "Arbol"
	valor1 := "Chau"
	valor2 := "Programacion"
	valor3 := "Binario"
	claves := []string{clave2, clave3, clave1}
	valores := []string{valor2, valor3, valor1}
	abb.Guardar(clave1, valor1)
	abb.Guardar(clave2, valor2)
	abb.Guardar(clave3, valor3)
	iter := abb.IteradorRango(nil, nil)
	i := 0
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, claves[i])
		require.EqualValues(t, valor, valores[i])
		i++
		iter.Siguiente()
	}
}

func TestIteradorExternoConRango(t *testing.T) {
	t.Log("Chequeamos que IteradorRango itere solo sobre un rango acotado")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	clave1 := 50
	clave2 := 75
	clave3 := 80
	clave4 := 35
	clave5 := 40
	clave6 := 20
	clave7 := 70
	claves := []int{clave6, clave4, clave5, clave1, clave7, clave2, clave3}
	abb.Guardar(clave1, clave1)
	abb.Guardar(clave2, clave2)
	abb.Guardar(clave3, clave3)
	abb.Guardar(clave4, clave4)
	abb.Guardar(clave5, clave5)
	abb.Guardar(clave6, clave6)
	abb.Guardar(clave7, clave7)
	inicio, fin := 35, 70
	iter := abb.IteradorRango(&inicio, &fin)
	i := 1
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, clave, claves[i])
		i++
		iter.Siguiente()
	}
}
func TestIteradorInternoSinRango(t *testing.T) {
	t.Log("Validamos que IterarRango sin un rango definido se comporta como Iterar de ABB")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1 := "Hola"
	clave2 := "Algoritmos"
	clave3 := "Arbol"
	valor1 := "Chau"
	valor2 := "Programacion"
	valor3 := "Binario"
	claves := []string{clave2, clave3, clave1}
	valores := []string{valor2, valor3, valor1}
	abb.Guardar(clave1, valor1)
	abb.Guardar(clave2, valor2)
	abb.Guardar(clave3, valor3)
	i := 0
	ptr := &i
	abb.IterarRango(nil, nil, func(clave string, valor string) bool {
		require.EqualValues(t, clave, claves[i])
		require.EqualValues(t, valor, valores[i])
		*ptr = *ptr + 1
		return true
	})
}

func TestIteradorInternoConRango(t *testing.T) {
	t.Log("Chequeamos que IterarRango itere solo sobre un rango acotado")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	clave1 := 50
	clave2 := 75
	clave3 := 80
	clave4 := 35
	clave5 := 40
	clave6 := 20
	clave7 := 70
	claves := []int{clave6, clave4, clave5, clave1, clave7, clave2, clave3}
	abb.Guardar(clave1, clave1)
	abb.Guardar(clave2, clave2)
	abb.Guardar(clave3, clave3)
	abb.Guardar(clave4, clave4)
	abb.Guardar(clave5, clave5)
	abb.Guardar(clave6, clave6)
	abb.Guardar(clave7, clave7)
	inicio, fin := 35, 70
	i := 1
	ptr := &i
	abb.IterarRango(&inicio, &fin, func(clave int, dato int) bool {
		require.EqualValues(t, clave, claves[i])
		*ptr = *ptr + 1
		return true
	})
}

func TestIteradorInternosSinDesde(t *testing.T) {
	t.Log("Chequeamos que IterarRango itere solo sobre un rango acotado")
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	clave1 := 50
	clave2 := 75
	clave3 := 80
	clave4 := 35
	clave5 := 40
	clave6 := 20
	clave7 := 70
	claves := []int{clave6, clave4, clave5, clave1, clave7, clave2, clave3}
	abb.Guardar(clave1, clave1)
	abb.Guardar(clave2, clave2)
	abb.Guardar(clave3, clave3)
	abb.Guardar(clave4, clave4)
	abb.Guardar(clave5, clave5)
	abb.Guardar(clave6, clave6)
	abb.Guardar(clave7, clave7)
	fin := 70
	i := 0
	ptr := &i
	abb.IterarRango(nil, &fin, func(clave int, dato int) bool {
		require.EqualValues(t, clave, claves[i])
		*ptr = *ptr + 1
		return true
	})
}

func TestIteradorInternoCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	clave1 := 50
	clave2 := 75
	clave3 := 80
	clave4 := 35
	clave5 := 40
	clave6 := 20
	clave7 := 70
	claves := []int{clave6, clave4, clave5, clave1, clave7, clave2, clave3}
	abb.Guardar(clave1, clave1)
	abb.Guardar(clave2, clave2)
	abb.Guardar(clave3, clave3)
	abb.Guardar(clave4, clave4)
	abb.Guardar(clave5, clave5)
	abb.Guardar(clave6, clave6)
	abb.Guardar(clave7, clave7)
	i := 0
	ptr := &i
	abb.Iterar(func(clave, valor int) bool {
		require.EqualValues(t, clave, claves[i])
		*ptr = *ptr + 1
		return clave < 60
	})
}
