# 🌟 Event Manager API REST

Este es un proyecto que implementa una API la cual permite gestionar eventos y organizarlos según el estado de revisión. En la API se pueden agregar, editar, leer, y borrar la información de los eventos ...

## 📋 Requerimientos

Asegúrate de tener instalados los siguientes requisitos antes de ejecutar la API:

- PostgreSQL instalado en tu máquina.
- Go instalado en tu sistema.

## ⚙️ Instalación / Configuración

1. **Clona el repositorio:**
   ```bash
    git clone https://github.com/Wainer-debug/event-manager.git


2. **Ingresa al directorio del proyecto:**
   ```bash
   cd event-manager


3. **Instala las dependencias:**
    
    ```bash
    go mod tidy

4. **Configura la base de datos:**

    * Asegúrate de tener el servidor PostgreSQL en ejecución.
    * Crea una base de datos (Usa la sentencia .sql que se encuentra en el proyecto).
    * Configura las credenciales y ajusta configuración de la base de datos en el archivo utils/database.go.

5. **Genera la documentación Swagger:**

    ```bash
    go get -u github.com/swaggo/swag/cmd/swag
    swag init

6.  **Inicia el servidor:**

    ```bash
    go run main.go
    
La aplicación se ejecutará en http://localhost:8080.

## 🚀 Uso

- **Funcionalidades / Endpoints**

    🤖 Registro de eventos:
        
        POST /events: Crea un nuevo evento.
        
    🤖 Edición de eventos:
        
        PUT /events/:id: Actualiza un evento existente.

    🤖 Eliminación de eventos:
        
        DELETE /events/:id: Elimina un evento por su ID.

    🤖 Obtención de eventos:
        
        GET /events: Obtiene la lista completa de eventos.

    🤖 Organización de eventos revisados:
        
        GET /events/management_required: Obtiene eventos revisados, clasificados en “Requiere gestión” o “sin gestión” 

## 📄 Ejemplos de Eventos

- Puedes usar la función GenerateSampleEvents en el archivo main.go para generar algunos eventos de ejemplo.

- Utiliza herramientas como Postman o curl para interactuar con la API:

    - **Crear un evento**
        
        Puedes crear un evento usando un cliente HTTP como cURL o Postman. Aquí tienes un ejemplo con cURL:
    
        ```bash
        curl -X POST -H "Content-Type: application/json" -d 
        '{
            "name": "Evento 12",
            "type_event_id": 2,
            "description": "Descripción del Evento 12",
            "status": "REVISADO"
        }' 
        http://localhost:8080/events

    O con Postman, usando la interfaz gráfica.

    - **Obtener eventos**
        
        Puedes obtener todos los eventos con el siguiente comando cURL:
    
        ```bash
        curl http://localhost:8080/events

    O simplemente accediendo al endpoint /events desde tu navegador o Postman.

## 🧪 Pruebas Unitarias
    
* Puedes ejecutar las pruebas unitarias con el siguiente comando:

    ```bash
    go test ./...

🚧 Las pruebas unitarias quedan pendientes para una nueva versión.
        
## 📖 Documentación en Swagger
    
Accede a la documentación de la API en Swagger en la siguiente URL:

http://localhost:8080/swagger/index.html

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Si quieres contribuir a este proyecto, ¡no dudes en enviar un pull request!


## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT.