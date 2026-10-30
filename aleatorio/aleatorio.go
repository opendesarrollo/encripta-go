package aleatorio

import (
    "crypto/rand"
    "errors"
)

// BytesAleatorios genera una secuencia de bytes aleatorios
// criptográficamente seguros mediante crypto/rand.
//
// Utiliza un generador de números aleatorios criptográficamente
// seguro (CSPRNG) para producir valores impredecibles adecuados
// para operaciones criptográficas.
//
// Parámetros:
//
//   - longitud: cantidad de bytes aleatorios que se desean
//     generar. Debe ser un número entero mayor que cero.
//
// Valores de retorno:
//
//   - datosCriptograficamenteSeguros: secuencia de bytes aleatorios
//     criptográficamente seguros, cuya longitud corresponde
//     exactamente al valor solicitado.
//
//   - error: error si la longitud indicada es menor o igual
//     a cero, o si falla la generación de los bytes aleatorios.
//
// Seguridad:
//
// Los bytes se generan mediante crypto/rand, que utiliza una
// fuente de aleatoriedad criptográficamente segura.
//
// Los valores generados son impredecibles desde el punto de vista
// criptográfico, siempre que la fuente de aleatoriedad sea segura.
//
// La generación aleatoria no garantiza que dos secuencias
// producidas en operaciones diferentes sean siempre distintas.
func BytesAleatorios(longitud int) ([]byte, error) {
    if longitud <= 0 {
        return nil, errors.New("la longitud debe ser mayor que cero")
    }

    datos := make([]byte, longitud)
    _, err := rand.Read(datos)
    if err != nil {
        return nil, err
    }

    return datos, nil
}