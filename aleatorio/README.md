# Paquete `aleatorio`

El paquete `aleatorio` forma parte de **Encripta-Go** y permite generar secuencias de bytes aleatorios criptográficamente seguros mediante `crypto/rand`, de la biblioteca estándar de Go.

Utiliza un generador de números aleatorios criptográficamente seguro (**CSPRNG**, *Cryptographically Secure Pseudorandom Number Generator*), apropiado para producir datos impredecibles que pueden emplearse en operaciones criptográficas.

> **Alcance:** el paquete genera bytes aleatorios. No cifra ni descifra información y no garantiza que dos secuencias distintas nunca coincidan.

## 1. Funcionalidad disponible

| Función | Propósito |
|---|---|
| `BytesAleatorios` | Genera la cantidad solicitada de bytes aleatorios criptográficamente seguros. |

## 2. Referencia técnica

### `BytesAleatorios`

**Firma:** `func BytesAleatorios(longitud int) ([]byte, error)`

Genera una secuencia aleatoria de la longitud indicada. Si la operación se completa correctamente, devuelve exactamente esa cantidad de bytes.

**Parámetro**

| Nombre | Tipo | Descripción | Restricción |
|---|---|---|---|
| `longitud` | `int` | Número de bytes que se desea generar. | Mayor que cero. |

**Valores de retorno**

| Tipo | Resultado |
|---|---|
| `[]byte` | Secuencia de bytes de la longitud solicitada. |
| `error` | `nil` si se completa correctamente; un error si la longitud no es válida o falla la generación. |

Cada byte puede contener un valor entre **0 y 255**. La longitud se especifica en *bytes*, no en bits: **1 byte = 8 bits**.

## 3. Funcionamiento paso a paso

### Ejemplo: solicitar 4 bytes

**Operación:** `BytesAleatorios(4)`

**Paso 1. Validación.** La función recibe `longitud = 4` y comprueba que sea mayor que cero. Como la condición se cumple, continúa.

**Paso 2. Reserva de memoria.** Crea un segmento `[]byte` con cuatro posiciones. Antes de llenarlo, su contenido es `[0 0 0 0]`.

**Paso 3. Generación.** `crypto/rand` llena las cuatro posiciones con valores aleatorios. Por ejemplo:

| Posición | Valor de ejemplo |
|---|---:|
| Byte 1 | 163 |
| Byte 2 | 241 |
| Byte 3 | 12 |
| Byte 4 | 123 |

**Paso 4. Retorno.** La función devuelve los bytes generados y `nil` como error.

**Ejemplo de contenido del `[]byte` devuelto:**

```text
[163 241 12 123]
```

| Dato | Resultado |
|---|---|
| Longitud solicitada | 4 bytes |
| Longitud devuelta | 4 bytes (32 bits) |
| Contenido de ejemplo | `[163 241 12 123]` |
| Error | `nil` |

> Los números mostrados son un ejemplo, **no una salida fija**. La función devuelve datos de tipo `[]byte`; la notación `[163 241 12 123]` es su representación legible en valores decimales, no una transformación a hexadecimal o texto.

## 4. Otros ejemplos de resultado

| Operación | Ejemplo de bytes devueltos | Longitud |
|---|---|---|
| `BytesAleatorios(2)` | `[45 198]` | 2 bytes (16 bits) |
| `BytesAleatorios(4)` | `[163 241 12 123]` | 4 bytes (32 bits) |
| `BytesAleatorios(8)` | `[72 185 14 203 56 91 247 128]` | 8 bytes (64 bits) |
| `BytesAleatorios(16)` | `[34 178 91 207 56 12 243 87 166 42 198 73 15 221 104 159]` | 16 bytes (128 bits) |

Estos resultados son **ilustrativos**: los valores cambian entre ejecuciones y, en condiciones normales, solo se mantiene fija la cantidad de bytes solicitada.

## 5. Validaciones y manejo de errores

| Situación | Entrada | Datos devueltos | Error |
|---|---|---|---|
| Longitud válida | `BytesAleatorios(4)` | Un `[]byte` de 4 elementos | `nil` |
| Longitud igual a cero | `BytesAleatorios(0)` | `nil` | `la longitud debe ser mayor que cero` |
| Longitud negativa | `BytesAleatorios(-4)` | `nil` | `la longitud debe ser mayor que cero` |
| Fallo en la generación | Longitud positiva | `nil` | Error devuelto por `crypto/rand` |

La aplicación que consuma esta función debe comprobar el valor de `error` antes de utilizar los bytes obtenidos.

## 6. Seguridad y consideraciones técnicas

- **Aleatoriedad criptográfica:** la función utiliza `crypto/rand`, respaldado por mecanismos seguros del entorno de ejecución; no utiliza `math/rand`.
- **Impredecibilidad:** bajo los supuestos de seguridad del generador y de su fuente de aleatoriedad, no resulta viable predecir nuevos bytes a partir de salidas anteriores.
- **Unicidad:** las secuencias aleatorias **pueden repetirse**. No existe una garantía absoluta de que dos llamadas produzcan valores diferentes.
- **Longitud adecuada:** que una secuencia tenga 16 o 32 bytes no significa que sea automáticamente apropiada para cualquier algoritmo. La longitud necesaria depende del uso criptográfico.
- **Responsabilidad del consumidor:** la función solo genera bytes. Su almacenamiento, uso y protección dependen de la aplicación que los consuma.

## 7. Dependencias

El paquete utiliza únicamente componentes de la biblioteca estándar de Go; no requiere dependencias externas.

| Biblioteca | Uso |
|---|---|
| `crypto/rand` | Generación de bytes aleatorios criptográficamente seguros. |
| `errors` | Creación del error de validación de longitud. |

## 8. Pruebas unitarias y comandos de ejecución

Las pruebas del paquete están definidas en `aleatorio/aleatorio_test.go` y utilizan `testing`, la biblioteca estándar de pruebas de Go.

**Ejecuta los siguientes comandos desde la raíz del repositorio.**

### 8.1. Ejecutar todas las pruebas de `aleatorio`

```bash
go test ./aleatorio -v
```

### 8.2. Ejecutar únicamente `TestBytesAleatorios`

```bash
go test ./aleatorio -run '^TestBytesAleatorios$' -v
```

### 8.3. Repetir la prueba sin utilizar resultados en caché

```bash
go test ./aleatorio -run '^TestBytesAleatorios$' -v -count=1
```

### 8.4. Consultar la cobertura de pruebas del paquete

```bash
go test ./aleatorio -cover
```

### 8.5. Ejecutar todas las pruebas del módulo

```bash
go test ./... -v
```

**Significado de las opciones:** `-v` muestra información detallada, `-run` selecciona una prueba por su nombre, `-count=1` evita reutilizar resultados de pruebas almacenados en caché y `-cover` informa de la cobertura de sentencias.

---

**Proyecto:** [Encripta-Go](https://github.com/opendesarrollo/encripta-go) · **Paquete:** `aleatorio`