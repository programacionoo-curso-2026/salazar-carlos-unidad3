# Streaming Go

## Proyecto Final - Programación con Go

Sistema backend para una plataforma de streaming desarrollado utilizando el lenguaje de programación Go.

El proyecto integra los principales contenidos estudiados durante las cuatro unidades de la asignatura, incluyendo programación con Go, estructuras de datos, programación orientada a objetos, concurrencia, testing y desarrollo de servicios web.

---

# 1. Datos del proyecto

**Nombre del proyecto:** Streaming Go

**Tipo de proyecto:** Servicios Web 

**Lenguaje principal:** Go

**Asignatura:** Programación con Go

**Docente:** DAVID HERNAN GUEVARA GUEVARA

**Grupo:**
- Elkin Barrera
- Carlos Salazar
- Willy Mendoza

---

# 2. Descripción del proyecto

Streaming Go es una aplicación backend que simula el funcionamiento de una plataforma de streaming.

El sistema permite administrar diferentes elementos relacionados con una plataforma audiovisual, como cuentas de usuarios, perfiles, contenidos, géneros, categorías, proveedores de video, favoritos, historial de reproducción, comentarios, valoraciones, planes y pagos.

La aplicación expone diferentes servicios web mediante HTTP y utiliza JSON como formato de intercambio de información entre el cliente y el servidor.

El proyecto fue desarrollado aplicando los conceptos aprendidos durante las cuatro unidades de la asignatura.

---

# 3. Objetivo general

Desarrollar una aplicación backend para una plataforma de streaming utilizando Go, integrando estructuras de datos, programación orientada a objetos, manejo de errores, concurrencia, testing y servicios web REST con serialización JSON.

---

# 4. Objetivos específicos

- Aplicar la sintaxis y características principales del lenguaje Go.
- Utilizar funciones, métodos y estructuras.
- Implementar Arrays, Slices y Maps.
- Aplicar conceptos de encapsulación.
- Utilizar interfaces y polimorfismo.
- Implementar manejo de errores.
- Utilizar Goroutines para ejecutar procesos concurrentes.
- Utilizar Channels para la comunicación entre procesos concurrentes.
- Desarrollar servicios web utilizando HTTP.
- Implementar operaciones CRUD.
- Utilizar JSON para la comunicación entre cliente y servidor.
- Implementar pruebas utilizando las herramientas de testing de Go.
- Organizar el proyecto mediante paquetes.
- Integrar todos los conocimientos en una aplicación funcional.

---

# 5. Problemática

Las plataformas de streaming necesitan administrar grandes cantidades de información relacionada con usuarios, perfiles, contenido audiovisual, historial de reproducción, favoritos, comentarios, valoraciones y suscripciones.

El proyecto plantea una solución backend simplificada que permite gestionar estos elementos mediante servicios web.

La arquitectura permite separar la lógica de negocio, los modelos de datos, el acceso a la información y la comunicación HTTP, facilitando la organización y mantenimiento del código.

---

# 6. Funcionalidades principales

El sistema permite realizar las siguientes operaciones:

### Cuentas

- Crear cuentas.
- Listar cuentas.
- Consultar una cuenta por ID.

### Perfiles

- Crear perfiles.
- Listar perfiles.

### Contenidos

- Crear contenidos.
- Listar contenidos.
- Consultar contenidos por ID.
- Actualizar contenidos.
- Eliminar contenidos.

### Favoritos

- Registrar favoritos.
- Listar favoritos.

### Comentarios

- Registrar comentarios.
- Listar comentarios.

### Valoraciones

- Registrar valoraciones.
- Listar valoraciones.

### Historial

- Registrar reproducciones.
- Consultar historial.

### Planes

- Crear planes.
- Listar planes.

### Proveedores

- Registrar proveedores de video.
- Listar proveedores.

### Pagos

- Registrar pagos.
- Listar pagos.

### Estadísticas

- Obtener información estadística del sistema mediante procesos concurrentes.
- Utilizar Goroutines.
- Utilizar Channels.

---

# 7. Tecnologías utilizadas

El proyecto utiliza las siguientes tecnologías:

- **Go:** lenguaje principal de programación.
- **net/http:** implementación del servidor HTTP y servicios web.
- **encoding/json:** serialización y deserialización JSON.
- **Goroutines:** ejecución concurrente de procesos.
- **Channels:** comunicación entre Goroutines.
- **Go Testing:** pruebas automatizadas.
- **Git:** control de versiones.
- **GitHub:** almacenamiento y gestión del repositorio.

---