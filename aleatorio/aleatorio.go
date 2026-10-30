package aleatorio

import (
    "crypto/rand"
    "errors"
)

func BytesAleatorios(longitud int) ([]byte, error) {
    if longitud <= 0 {
        return nil, errors.New("la longitud debe ser mayor que cero")
    }

    bytesAleatorios := make([]byte, longitud)
    _, err := rand.Read(bytesAleatorios)
    if err != nil {
        return nil, err
    }

    return bytesAleatorios, nil
}