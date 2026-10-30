// Package otp implementa la técnica de cifrado
// One-Time Pad.
package otp

import (
	"github.com/opendesarrollo/encripta-go/aleatorio"
    "errors"
)

func Cifrar(mensaje []byte) ([]byte, []byte, error) {
	mensajeBytes := len(mensaje)
    if mensajeBytes == 0 {
        return nil, nil, errors.New("el mensaje no puede estar vacio")
    }

    clave, err := aleatorio.BytesAleatorios(mensajeBytes)
    if err != nil {
        return nil, nil, err
    }

    cifrado := make([]byte, mensajeBytes)

    for i := range mensaje {
        cifrado[i] = mensaje[i] ^ clave[i]
    }

    return cifrado, clave, nil
}


func Descifrar(mensajeCifrado []byte, clave []byte) ([]byte, error) {
	mensajeCifradoBytes := len(mensajeCifrado)
	claveBytes := len(clave)

    if mensajeCifradoBytes == 0 {
        return nil, errors.New("el mensaje cifrado no puede estar vacio")
    }

    if claveBytes == 0 {
        return nil, errors.New("la clave no puede estar vacia")
    }

    if mensajeCifradoBytes != claveBytes {
        return nil, errors.New("el mensaje cifrado y la clave deben tener la misma longitud")
    }

    mensaje := make([]byte, mensajeCifradoBytes)

    for i := range mensajeCifrado {
        mensaje[i] = mensajeCifrado[i] ^ clave[i]
    }

    return mensaje, nil
}