# Encripta-Go

**Biblioteca criptográfica de código abierto para Go (Golang)**

Encripta-Go es una biblioteca de código abierto desarrollada en Go, diseñada para facilitar la integración de soluciones criptográficas en aplicaciones y servicios de software mediante una arquitectura modular, extensible y fácil de utilizar.

El proyecto nace con el propósito de acercar la criptografía al desarrollo de software, proporcionando herramientas reutilizables que permitan a los desarrolladores incorporar mecanismos de seguridad sin tener que implementar cada componente desde cero.

Su filosofía se fundamenta en la **simplicidad, la transparencia y la responsabilidad criptográfica**, priorizando implementaciones comprensibles, documentación técnica clara y buenas prácticas de desarrollo.

Encripta-Go busca evolucionar como una biblioteca que permita construir aplicaciones con funcionalidades criptográficas de forma organizada, mantenible y consciente de los requisitos de seguridad.

El proyecto forma parte de [OpenDesarrollo](https://opendesarrollo.com/), una iniciativa dedicada al desarrollo y publicación de herramientas tecnológicas de código abierto.

## 1. Características principales

- **Modularidad:** arquitectura organizada en paquetes independientes con responsabilidades claramente definidas.
- **Simplicidad:** interfaces comprensibles que facilitan la integración en proyectos Go.
- **Reutilización:** componentes diseñados para incorporarse en diferentes aplicaciones y servicios.
- **Transparencia:** documentación de los mecanismos implementados, sus propiedades y limitaciones.
- **Mantenibilidad:** estructura que facilita la revisión, evolución y ampliación de funcionalidades.
- **Seguridad:** desarrollo orientado a la aplicación responsable de principios y buenas prácticas criptográficas.
- **Código abierto:** distribución bajo la licencia Apache 2.0.

## 2. Paquetes disponibles

Encripta-Go organiza sus funcionalidades en paquetes independientes. Cada uno cuenta con documentación técnica que describe su funcionamiento, parámetros, resultados, ejemplos prácticos y pruebas unitarias.

| Paquete | Descripción | Documentación |
|---|---|---|
| [`aleatorio`](aleatorio/) | Generación de bytes aleatorios criptográficamente seguros mediante `crypto/rand`. | [Ver documentación](aleatorio/README.md) |
| [`otp`](otp/) | Cifrado y descifrado de información mediante One-Time Pad y operaciones XOR. | [Ver documentación](otp/README.md) |

## 3. Instalación

Para incorporar Encripta-Go como dependencia en un proyecto Go:

```bash
go get github.com/opendesarrollo/encripta-go
```

**Módulo:** `github.com/opendesarrollo/encripta-go`

Consulta la documentación de cada paquete para conocer sus funcionalidades y condiciones de utilización.

## 4. Seguridad y consideraciones

Encripta-Go promueve el uso responsable de herramientas criptográficas mediante implementaciones documentadas y con responsabilidades definidas.

La seguridad de una aplicación no depende únicamente de la biblioteca utilizada, sino también de su integración, configuración y los mecanismos de protección implementados.

Cada paquete documenta sus propiedades criptográficas, requisitos y limitaciones.

> **Importante:** el proyecto se encuentra en desarrollo y no cuenta con una auditoría de seguridad que avale su utilización para proteger información sensible en producción.

## 5. Ecosistema OpenDesarrollo

Encripta-Go forma parte de [OpenDesarrollo](https://opendesarrollo.com/), una iniciativa de desarrollo y publicación de herramientas tecnológicas de código abierto.

Dentro de sus proyectos también se encuentra [**Encripta**](https://encripta.opendesarrollo.com/), una aplicación orientada al cifrado y descifrado de información para usuarios finales.

Aunque comparten una identidad, ambos proyectos tienen propósitos diferentes: Encripta-Go proporciona componentes reutilizables para desarrolladores, mientras que Encripta ofrece una aplicación para usuarios finales.

**Enlaces relacionados:**

- [OpenDesarrollo](https://opendesarrollo.com/)
- [Encripta — Sitio web](https://encripta.opendesarrollo.com/)
- [Encripta — GitHub](https://github.com/opendesarrollo/Encripta)

## 6. Estado del proyecto

Encripta-Go es un proyecto de código abierto en evolución, con funcionalidades públicas disponibles y una arquitectura preparada para la incorporación progresiva de nuevos paquetes.

Actualmente no se ha declarado una versión estable.

## 7. Licencia

Este proyecto se distribuye bajo la **Licencia Apache 2.0**.

Consulta el archivo [LICENSE](LICENSE) para conocer los términos y condiciones de uso, modificación y distribución.

---

**Encripta-Go** · Un proyecto de [OpenDesarrollo](https://opendesarrollo.com/).

[Repositorio oficial en GitHub](https://github.com/opendesarrollo/encripta-go)