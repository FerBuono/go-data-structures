package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarElementosInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	lista.InsertarPrimero(10)
	require.Equal(t, 10, lista.VerPrimero())
	require.NotEqual(t, 10, lista.VerUltimo())
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(30)
	require.Equal(t, 30, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestInsertarElementosFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.NotEqual(t, 10, lista.VerPrimero())
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	require.Equal(t, 30, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())
}
func TestInsertarBorrarElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("primero")
	require.Equal(t, "primero", lista.VerPrimero())
	require.Equal(t, "primero", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "primero", lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	lista.InsertarUltimo("ultimo")
	require.Equal(t, "ultimo", lista.VerPrimero())
	require.Equal(t, "ultimo", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "ultimo", lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	lista.InsertarUltimo("primero")
	lista.InsertarUltimo("ultimo")
	require.Equal(t, "primero", lista.VerPrimero())
	require.Equal(t, "ultimo", lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, "primero", lista.BorrarPrimero())
	require.Equal(t, "ultimo", lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	lista.InsertarPrimero("cuatro")
	lista.InsertarPrimero("tres")
	lista.InsertarPrimero("dos")
	lista.InsertarPrimero("uno")
	require.Equal(t, "uno", lista.VerPrimero())
	require.Equal(t, "cuatro", lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, "uno", lista.BorrarPrimero())
	require.Equal(t, "dos", lista.BorrarPrimero())
	require.Equal(t, "tres", lista.BorrarPrimero())
	require.Equal(t, "cuatro", lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 1000; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
	}
	require.Greater(t, lista.Largo(), 1000)
	require.False(t, lista.EstaVacia())

	for j := 0; j <= 1000; j++ {
		require.Equal(t, j, lista.VerPrimero())
		require.Equal(t, 1000, lista.VerUltimo())
		require.Equal(t, j, lista.BorrarPrimero())
	}
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestInsertarElementosConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()

	iter.Insertar(0)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, iter.VerActual(), lista.VerPrimero())

	for i := 0; i < lista.Largo()/2; i++ {
		iter.Siguiente()
	}

	actual := iter.VerActual()
	iter.Insertar(26)
	require.Equal(t, 26, iter.Siguiente())
	require.Equal(t, actual, iter.VerActual())

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, iter.VerActual(), lista.VerUltimo())
}

func TestBorrarElementosConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	for i := 1; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	iter = lista.Iterador()

	primero := lista.VerPrimero()
	require.Equal(t, primero, iter.Borrar())
	require.NotEqual(t, primero, lista.VerPrimero())

	for i := 0; i < lista.Largo()/2; i++ {
		iter.Siguiente()
	}

	actual := iter.VerActual()
	require.Equal(t, actual, iter.Borrar())
	require.NotEqual(t, actual, iter.VerActual())

	iter = lista.Iterador()

	for i := 0; i < lista.Largo()-1; i++ {
		iter.Siguiente()
	}

	ultimo := lista.VerUltimo()
	require.Equal(t, ultimo, iter.Borrar())
	require.NotEqual(t, ultimo, lista.VerUltimo())
}
func TestFuncionalidadIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	for i := 0; i <= 10; i++ {
		iterador.Insertar(i)
	}
	for i := 10; i >= 0; i-- {
		require.True(t, iterador.HaySiguiente())
		require.Equal(t, i, iterador.Siguiente())
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	iterador2 := lista.Iterador()
	for i := 10; i >= 0; i-- {
		require.True(t, iterador2.HaySiguiente())
		require.Equal(t, i, iterador2.Borrar())
	}
	require.Equal(t, 0, lista.Largo())
	for i := 0; i <= 10; i++ {
		iterador2.Insertar(i)
		if i == 5 {
			iterador2.Insertar(30)
		}
	}
	iterador3 := lista.Iterador()
	for i := 10; i >= 0; i-- {
		if i == 5 {
			require.True(t, iterador3.HaySiguiente())
			require.Equal(t, 30, iterador3.Borrar())
		}
		require.True(t, iterador3.HaySiguiente())
		require.Equal(t, i, iterador3.Borrar())
	}
}

func TestIteradorExtVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	for i := 0; i <= 2000; i++ {
		iterador.Insertar(i)
	}
	for i := 2000; i >= 0; i-- {
		require.True(t, iterador.HaySiguiente())
		require.Equal(t, i, iterador.Siguiente())
	}
	iterador2 := lista.Iterador()
	for i := 2000; i >= 0; i-- {
		require.True(t, iterador2.HaySiguiente())
		require.Equal(t, i, iterador2.Borrar())
	}
	require.Equal(t, 0, lista.Largo())
}
func TestFuncionalidadIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarPrimero(i)
	}
	datos_iterados := 0
	lista.Iterar(func(dato int) bool {
		datos_iterados++
		return true
	})
	require.Equal(t, 10, datos_iterados)
}

func TestIteradorFinalizaAlFinalizarLaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	suma_iterada := 0
	lista.Iterar(func(dato int) bool {
		require.Equal(t, dato, suma_iterada)
		suma_iterada++
		return true
	})
}

func TestIteradorFinalizaAlIndicarUsuario(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for i := 0; i < 20; i++ {
		lista.InsertarUltimo(i)
		suma += i
	}
	suma_iterada := 0
	lista.Iterar(func(dato int) bool {
		suma_iterada += dato
		return suma_iterada < 20
	})
	require.Less(t, suma_iterada, suma)
	suma_iterada = suma
	lista.Iterar(func(dato int) bool {
		suma_iterada = suma_iterada - (dato * 2)
		return suma_iterada > 0
	})
	require.Less(t, suma_iterada, suma)
}

func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numero := 0
	lista.Iterar(func(dato int) bool {
		numero += dato
		return true
	})
	require.Zero(t, numero)
}

func TestVolumenIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for i := 0; i < 2000; i++ {
		suma += i
		lista.InsertarUltimo(i)
	}
	suma_iterada := 0
	lista.Iterar(func(dato int) bool {
		suma_iterada += dato
		return true
	})
}
