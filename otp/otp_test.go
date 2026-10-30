package otp

import "testing"

func TestOTP(t *testing.T) {
    t.Logf("Hola mundo")
}


func TestCifrarYDescifrar(t *testing.T) {
    mensaje := []byte("Hola mundo")

    cifrado, clave, err := Cifrar(mensaje)
    if err != nil {
        t.Fatal(err)
    }

    descifrado, err := Descifrar(cifrado, clave)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("Mensaje original:	%s", mensaje)
    t.Logf("Mensaje cifrado: 	%v", cifrado)
    t.Logf("Clave: 				%v", clave)
    t.Logf("Mensaje cifrado:	%X", cifrado)
    t.Logf("Clave: 				%X", clave)
    t.Logf("Mensaje descifrado: %s", descifrado)

    if string(descifrado) != string(mensaje) {
        t.Fatal("El mensaje descifrado no coincide con el original")
    }
}