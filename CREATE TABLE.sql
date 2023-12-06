-- CREACIÓN DE LA TABLA
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type_event_id INTEGER NOT NULL,
    description TEXT NOT NULL,
    date TIMESTAMPTZ DEFAULT current_timestamp,
    status VARCHAR(50) NOT NULL,
    management_required BOOLEAN,
    management_status VARCHAR(50)
    --CONSTRAINT fk_type_event FOREIGN KEY (type_event_id) REFERENCES type_events(id)
);


-- INSERTAR EVENTO DE EJEMPLO
INSERT INTO events (name, type_event_id, description, status, management_required, management_status)
VALUES ('Evento', 2, 'Descripción dl evento', 'REVISADO', true, 'Requiere gestión');


SELECT * FROM public.events;

