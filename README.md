# Event Manager API

Esta API permite gestionar eventos y organizarlos según el estado de revisión.
En la API se puedan agregar, editar, leer, y borrar la información de los eventos.

## Instalación

1. Clona este repositorio: `git clone https://github.com/Wainer-debug/event-manager.git`
2. Ingresa al directorio del proyecto: `cd event-manager`
3. Instala las dependencias: `go mod tidy`
4. Configura la base de datos: (Agrega instrucciones si es necesario)

## Configuración

1. Configura el archivo `.env` con la información de la base de datos y otras configuraciones si es necesario.

## Uso

1. Inicia la aplicación: `go run main.go`
2. Accede a la documentación de la API: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Endpoints

- **GET /events:** Obtener la lista de eventos.
- **POST /events:** Crear un nuevo evento.
- **PUT /events/:id:** Actualizar un evento existente.
- **DELETE /events/:id:** Eliminar un evento.
- **GET /events/management_required:** Obtener eventos revisados por Requiere Gestión / Sin Gestión.

## Ejemplos

Agrega ejemplos de cómo utilizar cada endpoint.

## Pruebas Unitarias

1. Ejecuta las pruebas unitarias: `go test ./...`

## Documentación de la API

Accede a la documentación de la API en [Swagger](http://localhost:8080/swagger/index.html) mientras la aplicación está en ejecución.

