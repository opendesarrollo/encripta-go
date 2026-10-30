# Encripta Go

**Librería criptográfica One-Time Pad para Go**

Encripta Go es un proyecto de código abierto orientado al desarrollo de una librería criptográfica en el lenguaje de programación **Go (Golang)**, basada en la técnica de cifrado **One-Time Pad (OTP)**.

El proyecto tiene como objetivo proporcionar una implementación sencilla, segura y documentada, fundamentada en operaciones matemáticas de cifrado, generación de claves criptográficamente seguras y administración del material criptográfico.

Su desarrollo busca facilitar la integración de esta técnica en aplicaciones, servicios y proyectos de software, priorizando la correcta implementación de sus principios criptográficos y la prevención de errores que puedan comprometer la confidencialidad de la información.

> **Estado del proyecto:** En desarrollo. La implementación criptográfica y la API pública aún no están disponibles.

## 1. Fundamento criptográfico

### ¿Qué es One-Time Pad?

**One-Time Pad (OTP)**, conocido en español como *cifrado de libreta de un solo uso*, es una técnica de criptografía simétrica que permite cifrar información mediante una clave secreta aleatoria cuya longitud debe ser igual a la del mensaje.

Su característica fundamental es que **cada clave, o segmento de ella, debe utilizarse una sola vez y nunca reutilizarse**.

A diferencia de otros algoritmos de cifrado simétrico, One-Time Pad puede proporcionar una propiedad denominada **secreto perfecto**, siempre que se cumplan rigurosamente las condiciones matemáticas necesarias.

### Principios fundamentales

La seguridad de One-Time Pad depende de cuatro condiciones:

1. **Aleatoriedad:** la clave debe ser verdaderamente aleatoria, uniforme e independiente del mensaje.
2. **Longitud:** la clave debe contener suficiente material aleatorio para cubrir la totalidad del mensaje.
3. **Uso único:** cada porción de la clave debe utilizarse exactamente una vez para cifrar.
4. **Confidencialidad:** la clave debe permanecer secreta y ser conocida únicamente por las partes autorizadas.

El incumplimiento de cualquiera de estas condiciones puede comprometer las propiedades de seguridad del sistema.

## 2. Fundamento matemático

La implementación binaria de One-Time Pad puede realizarse utilizando la operación lógica **XOR (OR exclusivo)**.

### Cifrado

`C = M ⊕ K`

### Descifrado

`M = C ⊕ K`

Donde:

- **M:** mensaje original.
- **K:** clave secreta.
- **C:** mensaje cifrado.
- **⊕:** operación XOR.

La operación XOR permite recuperar el mensaje original aplicando nuevamente la misma clave sobre el texto cifrado.

Esta propiedad constituye la base matemática de la variante binaria del sistema.

## 3. Secreto perfecto

El secreto perfecto es una propiedad de seguridad definida en la teoría de la información y formalizada por el matemático **Claude Shannon**.

Un sistema One-Time Pad ideal proporciona secreto perfecto cuando las claves utilizadas son uniformemente aleatorias, independientes, secretas y no reutilizadas.

Bajo estas condiciones, el texto cifrado no revela información sobre el mensaje original.

Esta garantía matemática no depende del lenguaje de programación empleado, sino del cumplimiento de las condiciones criptográficas del sistema.

**Importante:** la utilización de XOR o de un generador criptográficamente seguro no garantiza, por sí sola, el secreto perfecto.

## 4. Generación de claves criptográficas

Encripta Go contempla utilizar el paquete estándar `crypto/rand` de Go para generar material criptográfico.

Este paquete proporciona acceso a un generador de números aleatorios criptográficamente seguro, adecuado para la generación de claves y otros valores secretos.

Se debe distinguir entre la seguridad computacional de un generador criptográfico y la aleatoriedad ideal requerida para demostrar matemáticamente el secreto perfecto.

Por lo tanto, el uso de `crypto/rand` no constituye, por sí solo, una demostración de secreto perfecto.

## 5. Administración de claves

La correcta administración del material criptográfico es uno de los principales desafíos de One-Time Pad.

El diseño de Encripta Go buscará contemplar los siguientes aspectos:

- Generación segura del material criptográfico.
- Prevención de reutilización accidental de claves.
- Control del consumo de segmentos de claves.
- Validación de longitudes de los mensajes y claves.
- Protección de las claves contra accesos no autorizados.
- Manejo adecuado de errores.
- Separación de las responsabilidades de cifrado y administración de claves.

La distribución segura de claves entre las partes involucradas deberá abordarse como un componente independiente del proceso de cifrado.

## 6. Consideraciones de seguridad

One-Time Pad ofrece confidencialidad bajo sus condiciones matemáticas ideales, pero **no proporciona por sí mismo autenticación ni integridad de los mensajes**.

Esto significa que la información cifrada podría ser modificada por un tercero sin que el sistema detecte necesariamente la alteración.

La implementación deberá considerar estas limitaciones y documentar los controles necesarios para utilizar la técnica de forma responsable.

La librería no deberá considerarse apta para proteger información sensible en producción sin las pruebas, revisiones y validaciones de seguridad correspondientes.

## 7. Principios de desarrollo

Encripta Go se desarrollará siguiendo los siguientes principios:

**Simplicidad:** mantener una API comprensible y evitar dependencias innecesarias.

**Corrección:** implementar operaciones matemáticas correctas, validaciones explícitas y pruebas automatizadas.

**Seguridad:** priorizar el manejo responsable de claves y la prevención de errores criptográficos.

**Transparencia:** documentar las características, los supuestos de seguridad y las limitaciones de la implementación.

**Mantenibilidad:** organizar el código para facilitar su revisión, mantenimiento y evolución.

## 8. Estado del proyecto

Actualmente, Encripta Go se encuentra en su etapa inicial de diseño.

La implementación, las pruebas automatizadas y la documentación de la API se incorporarán conforme avance el desarrollo.

No se ha publicado todavía una versión estable.

## 9. Licencia

Este proyecto se distribuye bajo la **Licencia Apache 2.0**.

Consulte el archivo [LICENSE](LICENSE) para conocer los términos y condiciones de uso, modificación y distribución.

---

**Encripta Go** — Proyecto de código abierto para el desarrollo de herramientas criptográficas en Go.