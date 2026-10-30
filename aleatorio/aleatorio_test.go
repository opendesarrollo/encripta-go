package aleatorio

import "testing"

func TestAleatorio(t *testing.T) {
    t.Logf("Hola mundo")
}


func TestBytesAleatorios(t *testing.T) {
    datos, err := BytesAleatorios(2)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("Bytes generados: %d", len(datos))
    t.Logf("Datos en HEX:   %X", datos)
    t.Logf("Datos en bytes: %v", datos)
}