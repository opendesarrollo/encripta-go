
# Paquete `otp`

El paquete `otp` forma parte de **Encripta-Go** e implementa el mecanismo de cifrado **One-Time Pad (OTP)**, también conocido como cifrado de libreta de un solo uso.

Este mecanismo utiliza una clave aleatoria de la misma longitud que el mensaje y aplica la operación lógica **XOR (OR exclusivo)** para cifrar y descifrar información.

La implementación utiliza el paquete `aleatorio` de Encripta-Go, que genera claves mediante `crypto/rand`, un generador de números aleatorios criptográficamente seguro (CSPRNG).

> **Alcance:** el paquete permite cifrar y descifrar secuencias de bytes. La seguridad del mecanismo depende de la generación, protección y utilización correcta de la clave. No proporciona autenticación ni verificación de integridad.

## 1. Funcionalidades disponibles

| Función | Propósito |
|---|---|
| `Cifrar` | Cifra un mensaje mediante XOR y genera una clave aleatoria de la misma longitud. |
| `Descifrar` | Recupera el mensaje original mediante XOR utilizando el mensaje cifrado y su clave correspondiente. |

## 2. Referencia técnica

### 2.1. `Cifrar`

**Firma:** `func Cifrar(mensaje []byte) ([]byte, []byte, error)`

**Descripción**

Recibe un mensaje representado como una secuencia de bytes y genera una clave aleatoria criptográficamente segura de la misma longitud.

Posteriormente, aplica XOR entre cada byte del mensaje y el byte correspondiente de la clave.

La función devuelve el mensaje cifrado, la clave utilizada y un posible error.

**Parámetros**

| Nombre | Tipo | Descripción | Restricción |
|---|---|---|---|
| `mensaje` | `[]byte` | Secuencia de bytes que se desea cifrar. | No puede estar vacía. |

**Valores de retorno**

| Retorno | Tipo | Descripción |
|---|---|---|
| `cifrado` | `[]byte` | Secuencia de bytes resultante de aplicar XOR al mensaje y la clave. |
| `clave` | `[]byte` | Clave aleatoria generada para cifrar el mensaje. |
| `error` | `error` | `nil` si la operación finaliza correctamente; un error si el mensaje está vacío o falla la generación de la clave. |

**Comportamiento esperado**

- Valida que el mensaje contenga al menos un byte.
- Determina la longitud del mensaje.
- Genera una clave aleatoria de esa misma longitud utilizando `aleatorio.BytesAleatorios`.
- Aplica XOR a cada par de bytes del mensaje y la clave.
- Devuelve el mensaje cifrado y la clave generada.

### 2.2. `Descifrar`

**Firma:** `func Descifrar(mensajeCifrado []byte, clave []byte) ([]byte, error)`

**Descripción**

Recupera el mensaje original aplicando XOR entre los bytes del mensaje cifrado y los bytes de la clave utilizada durante el cifrado.

Debido a las propiedades de XOR, aplicar nuevamente la misma clave sobre el mensaje cifrado permite recuperar los bytes originales.

**Parámetros**

| Nombre | Tipo | Descripción | Restricción |
|---|---|---|---|
| `mensajeCifrado` | `[]byte` | Secuencia de bytes cifrados. | No puede estar vacía. |
| `clave` | `[]byte` | Clave utilizada durante el cifrado. | No puede estar vacía y debe tener la misma longitud que el mensaje cifrado. |

**Valores de retorno**

| Retorno | Tipo | Descripción |
|---|---|---|
| `mensaje` | `[]byte` | Secuencia de bytes recuperados mediante XOR. |
| `error` | `error` | `nil` si la operación finaliza correctamente; un error si los parámetros no cumplen las validaciones. |

**Comportamiento esperado**

- Valida que el mensaje cifrado no esté vacío.
- Verifica que la clave no esté vacía.
- Comprueba que la clave y el mensaje cifrado tengan la misma longitud.
- Aplica XOR entre cada byte del mensaje cifrado y su correspondiente byte de la clave.
- Devuelve la secuencia de bytes recuperada.

## 3. Funcionamiento paso a paso

### 3.1. Ejemplo: cifrar un mensaje de 4 bytes

Supongamos que una aplicación necesita cifrar un mensaje compuesto por **4 bytes**.

**Operación:** `Cifrar([72 79 76 65])`

El mensaje contiene la siguiente secuencia:

`[72 79 76 65]`

Cada valor representa un byte del mensaje original.

**Paso 1. Validación del mensaje.**

La función determina la longitud de la secuencia recibida.

| Validación | Resultado |
|---|---|
| Mensaje recibido | `[72 79 76 65]` |
| Longitud | 4 bytes |
| Mensaje vacío | No |
| Resultado | Válido |

Como el mensaje contiene datos, continúa con la operación.

**Paso 2. Generación de la clave aleatoria.**

La función solicita al paquete `aleatorio` generar una clave de 4 bytes, porque esa es la longitud del mensaje.

**Operación interna:** `BytesAleatorios(4)`

**Ejemplo de clave generada:**

`[163 241 12 123]`

| Posición | Mensaje original | Clave generada |
|---|---:|---:|
| Byte 1 | 72 | 163 |
| Byte 2 | 79 | 241 |
| Byte 3 | 76 | 12 |
| Byte 4 | 65 | 123 |

La clave debe conservarse de forma segura porque será necesaria para recuperar el mensaje original.

**Paso 3. Aplicación de la operación XOR.**

La función aplica XOR entre cada byte del mensaje original y el byte correspondiente de la clave.

XOR es una operación lógica que compara dos valores bit a bit.

El resultado es `1` cuando los bits son diferentes y `0` cuando son iguales.

**Ejemplo de operación:**

| Posición | Mensaje | Clave | Operación XOR | Resultado |
|---|---:|---:|---|---:|
| Byte 1 | 72 | 163 | `72 ^ 163` | 235 |
| Byte 2 | 79 | 241 | `79 ^ 241` | 190 |
| Byte 3 | 76 | 12 | `76 ^ 12` | 64 |
| Byte 4 | 65 | 123 | `65 ^ 123` | 58 |

Los resultados de cada operación forman el mensaje cifrado.

**Paso 4. Resultado del cifrado.**

La función devuelve el mensaje cifrado, la clave generada y `nil` como error.

**Mensaje cifrado:**

```text
[235 190 64 58]
```

**Clave generada:**

```text
[163 241 12 123]
```

**Resumen del resultado**

| Elemento | Resultado |
|---|---|
| Función | `Cifrar` |
| Mensaje original | `[72 79 76 65]` |
| Longitud | 4 bytes |
| Mensaje cifrado | `[235 190 64 58]` |
| Clave | `[163 241 12 123]` |
| Error | `nil` |

> Los valores de la clave y del mensaje cifrado son ilustrativos. Como la clave se genera aleatoriamente, dos ejecuciones con el mismo mensaje pueden producir resultados diferentes.

### 3.2. Ejemplo: descifrar el mensaje

Para recuperar el mensaje original, se utiliza el mensaje cifrado y la misma clave generada durante el cifrado.

**Operación:**

`Descifrar([235 190 64 58], [163 241 12 123])`

**Paso 1. Validación de los parámetros.**

La función verifica que ambos segmentos contengan datos y tengan la misma longitud.

| Validación | Resultado |
|---|---|
| Mensaje cifrado | `[235 190 64 58]` |
| Clave | `[163 241 12 123]` |
| Longitud del mensaje cifrado | 4 bytes |
| Longitud de la clave | 4 bytes |
| Resultado | Válido |

**Paso 2. Aplicación de XOR.**

La función aplica XOR entre el mensaje cifrado y la clave correspondiente.

| Posición | Mensaje cifrado | Clave | Operación XOR | Resultado |
|---|---:|---:|---|---:|
| Byte 1 | 235 | 163 | `235 ^ 163` | 72 |
| Byte 2 | 190 | 241 | `190 ^ 241` | 79 |
| Byte 3 | 64 | 12 | `64 ^ 12` | 76 |
| Byte 4 | 58 | 123 | `58 ^ 123` | 65 |

**Paso 3. Recuperación del mensaje original.**

Los bytes resultantes forman nuevamente el mensaje original.

**Resultado:**

```text
[72 79 76 65]
```

| Elemento | Resultado |
|---|---|
| Función | `Descifrar` |
| Mensaje cifrado | `[235 190 64 58]` |
| Clave | `[163 241 12 123]` |
| Mensaje recuperado | `[72 79 76 65]` |
| Error | `nil` |

**Conclusión:** al aplicar XOR con la misma clave utilizada durante el cifrado, se recuperan los bytes originales.

## 4. Otros ejemplos de funcionamiento

Los siguientes ejemplos muestran operaciones de cifrado y descifrado con diferentes longitudes.

Los valores son ilustrativos y representan directamente los bytes contenidos en cada `[]byte`.

### Ejemplo 1. Mensaje de 2 bytes

**Entrada:** `Cifrar([10 20])`

**Clave de ejemplo:** `[5 7]`

**Mensaje cifrado:** `[15 19]`

**Descifrado:** `Descifrar([15 19], [5 7])`

**Resultado:** `[10 20]`

### Ejemplo 2. Mensaje de 4 bytes

**Entrada:** `Cifrar([72 79 76 65])`

**Clave de ejemplo:** `[163 241 12 123]`

**Mensaje cifrado:** `[235 190 64 58]`

**Descifrado:** `Descifrar([235 190 64 58], [163 241 12 123])`

**Resultado:** `[72 79 76 65]`

### Ejemplo 3. Mensaje de 8 bytes

**Entrada:** `Cifrar([1 2 3 4 5 6 7 8])`

**Clave de ejemplo:** `[10 20 30 40 50 60 70 80]`

**Mensaje cifrado:** `[11 22 29 44 55 58 65 88]`

**Descifrado:** `Descifrar([11 22 29 44 55 58 65 88], [10 20 30 40 50 60 70 80])`

**Resultado:** `[1 2 3 4 5 6 7 8]`

### Relación entre las longitudes

| Longitud del mensaje | Longitud de la clave | Longitud del cifrado |
|---|---|---|
| 2 bytes | 2 bytes | 2 bytes |
| 4 bytes | 4 bytes | 4 bytes |
| 8 bytes | 8 bytes | 8 bytes |
| 16 bytes | 16 bytes | 16 bytes |
| 32 bytes | 32 bytes | 32 bytes |

La implementación mantiene la misma longitud para el mensaje original, la clave y el mensaje cifrado.

## 5. Validaciones y manejo de errores

### 5.1. Función `Cifrar`

| Situación | Entrada | Cifrado | Clave | Error |
|---|---|---|---|---|
| Mensaje válido | `Cifrar([10 20])` | 2 bytes | 2 bytes | `nil` |
| Mensaje vacío | `Cifrar([])` | `nil` | `nil` | `el mensaje no puede estar vacio` |
| Fallo al generar la clave | Mensaje no vacío | `nil` | `nil` | Error devuelto por `aleatorio.BytesAleatorios` |

### 5.2. Función `Descifrar`

| Situación | Resultado |
|---|---|
| Mensaje cifrado vacío | Error: `el mensaje cifrado no puede estar vacio` |
| Clave vacía | Error: `la clave no puede estar vacia` |
| Longitudes diferentes | Error: `el mensaje cifrado y la clave deben tener la misma longitud` |
| Parámetros válidos | Devuelve el mensaje recuperado y `nil` |

**Ejemplos de entradas incorrectas**

**Caso 1. Mensaje cifrado vacío**

`Descifrar([], [10 20])`

Resultado: `nil` y error `el mensaje cifrado no puede estar vacio`.

**Caso 2. Clave vacía**

`Descifrar([10 20], [])`

Resultado: `nil` y error `la clave no puede estar vacia`.

**Caso 3. Longitudes diferentes**

`Descifrar([10 20 30], [5 7])`

Resultado: `nil` y error `el mensaje cifrado y la clave deben tener la misma longitud`.

> La validación de longitud no comprueba que la clave sea la correcta. Si se proporciona una clave diferente de la original, pero con la longitud adecuada, la función devolverá bytes sin generar necesariamente un error.

## 6. Seguridad y consideraciones técnicas

### 6.1. Fundamento criptográfico

El mecanismo One-Time Pad se basa en la combinación de un mensaje y una clave aleatoria mediante XOR.

El OTP puede proporcionar **secreto perfecto**, siempre que se cumplan sus condiciones criptográficas fundamentales:

- La clave es verdaderamente aleatoria y uniforme.
- La clave tiene la misma longitud que el mensaje.
- La clave se utiliza una sola vez.
- La clave permanece secreta y solo está disponible para las partes autorizadas.

La implementación utiliza `crypto/rand` indirectamente a través del paquete `aleatorio`, para obtener claves criptográficamente seguras.

Esto proporciona seguridad computacional basada en el generador utilizado, pero no demuestra por sí solo secreto perfecto en sentido matemático.

### 6.2. Uso único de la clave

**Una clave OTP nunca debe reutilizarse para cifrar mensajes diferentes.**

Si dos mensajes se cifran con la misma clave, un atacante que disponga de ambos textos cifrados puede obtener la operación XOR entre los mensajes originales.

Esta situación compromete las propiedades de seguridad del mecanismo.

### 6.3. Protección de la clave

La función `Cifrar` devuelve tanto el mensaje cifrado como la clave utilizada.

La aplicación que consume la biblioteca es responsable de proteger la clave.

La clave debe mantenerse confidencial y distribuirse mediante mecanismos seguros cuando sea necesario compartirla con el destinatario.

La clave no debe exponerse públicamente ni almacenarse junto al mensaje cifrado sin medidas de protección apropiadas.

### 6.4. Integridad y autenticación

La operación XOR no proporciona mecanismos para verificar si los bytes cifrados fueron modificados.

Por lo tanto, este paquete no garantiza:

- Integridad del mensaje cifrado.
- Autenticación del remitente.
- Detección de modificaciones maliciosas.
- Protección frente a sustitución de mensajes cifrados.

Para sistemas que requieren confidencialidad e integridad, debe considerarse un esquema de cifrado autenticado, como AES-GCM, adecuado al caso de uso.

### 6.5. Limitaciones del descifrado

La función `Descifrar` verifica únicamente que los parámetros cumplan las condiciones de longitud y contenido.

No puede determinar si una clave corresponde realmente al mensaje cifrado.

Una clave incorrecta de la longitud esperada puede producir una secuencia distinta del mensaje original sin devolver un error.

### 6.6. Alcance del paquete

Las funciones implementan operaciones de cifrado y descifrado mediante XOR.

No implementan por sí mismas:

- Almacenamiento seguro de claves.
- Distribución segura de claves.
- Autenticación de mensajes.
- Verificación de integridad.
- Administración del ciclo de vida de las claves.
- Prevención de reutilización de claves.

Estas responsabilidades corresponden a la aplicación que utiliza la biblioteca.

## 7. Dependencias

El paquete utiliza una dependencia interna de Encripta-Go y una biblioteca estándar de Go.

| Dependencia | Propósito |
|---|---|
| `github.com/opendesarrollo/encripta-go/aleatorio` | Generación de claves aleatorias criptográficamente seguras. |
| `errors` | Construcción de errores de validación. |

La generación de claves depende del funcionamiento del paquete `aleatorio`.

## 8. Pruebas unitarias y comandos de ejecución

Las pruebas del paquete `otp` pueden ejecutarse mediante la herramienta estándar `go test`.

**Ejecuta los siguientes comandos desde la raíz del repositorio Encripta-Go.**

### 8.1. Ejecutar todas las pruebas del paquete `otp`

```bash
go test ./otp -v
```

Ejecuta todas las pruebas unitarias definidas en el paquete `otp`.

### 8.2. Ejecutar la prueba de cifrado y descifrado

```bash
go test ./otp -run '^TestCifrarYDescifrar$' -v
```

Ejecuta exclusivamente la prueba `TestCifrarYDescifrar`, si está definida en el paquete.

### 8.3. Ejecutar la prueba sin utilizar resultados en caché

```bash
go test ./otp -run '^TestCifrarYDescifrar$' -v -count=1
```

Ejecuta nuevamente la prueba sin reutilizar resultados almacenados en caché.

### 8.4. Consultar la cobertura de pruebas

```bash
go test ./otp -cover
```

Muestra el porcentaje de cobertura de sentencias alcanzado por las pruebas del paquete.

### 8.5. Ejecutar todas las pruebas del módulo

```bash
go test ./... -v
```

Ejecuta las pruebas de todos los paquetes del proyecto Encripta-Go.

**Significado de las opciones:**

- `-v`: muestra información detallada de las pruebas.
- `-run`: selecciona pruebas mediante una expresión regular.
- `-count=1`: evita reutilizar resultados de pruebas almacenados en caché.
- `-cover`: presenta la cobertura de sentencias.

---

**Proyecto:** [Encripta-Go](https://github.com/opendesarrollo/encripta-go) · **Paquete:** `otp`