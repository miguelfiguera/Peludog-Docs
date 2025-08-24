# Diseño de la Base de Datos - PeluDog CRM

## 1. Introducción

Este documento describe la estructura y el diseño de la base de datos para el sistema PeluDog CRM. El diseño se basa en la información recopilada de los documentos de arquitectura del sistema.

La base de datos seleccionada es **MySQL 8.0+**, y la interacción con ella se realizará a través del ORM de Ruby on Rails, **ActiveRecord**. Esto significa que los nombres de las tablas seguirán la convención de pluralización de Rails y todas las tablas incluirán campos `id` como clave primaria y `created_at`/`updated_at` para el control de versiones.

## 2. Diagrama de Entidad-Relación

Un diagrama visual de la relación entre las tablas es fundamental para entender el modelo. El diseño aquí descrito se basa en el diagrama de clases que se encuentra en `Imagenes/DiagramaDeClases.png`. Se recomienda consultar dicho diagrama para una representación visual.

## 3. Definición de Entidades y Tablas

A continuación se detallan las tablas principales del sistema, sus campos y las relaciones entre ellas.

### Tabla: `users`

Almacena la información de los empleados de la clínica que utilizan el sistema (veterinarios, asistentes, administradores).

| Campo             | Tipo          | Descripción                                       |
| ----------------- | ------------- | ------------------------------------------------- |
| `id`              | `bigint` (PK) | Identificador único para cada usuario.            |
| `name`            | `string`      | Nombre completo del usuario.                      |
| `email`           | `string`      | Correo electrónico para el inicio de sesión (único). |
| `password_digest` | `string`      | Hash de la contraseña para autenticación segura.  |
| `role`            | `string`      | Rol del usuario (`admin`, `veterinario`, `asistente`). |
| `created_at`      | `datetime`    | Fecha y hora de creación del registro.            |
| `updated_at`      | `datetime`    | Fecha y hora de la última actualización.          |

### Tabla: `clients`

Contiene la información de los dueños de las mascotas.

| Campo          | Tipo          | Descripción                                     |
| -------------- | ------------- | ----------------------------------------------- |
| `id`           | `bigint` (PK) | Identificador único para cada cliente.          |
| `name`         | `string`      | Nombre del cliente.                             |
| `last_name`    | `string`      | Apellido del cliente.                           |
| `phone_number` | `string`      | Número de teléfono de contacto.                 |
| `email`        | `string`      | Correo electrónico del cliente.                 |
| `address`      | `text`        | Dirección física del cliente.                   |
| `created_at`   | `datetime`    | Fecha y hora de creación del registro.          |
| `updated_at`   | `datetime`    | Fecha y hora de la última actualización.        |

### Tabla: `pets`

Almacena la información de cada mascota (paciente).

| Campo       | Tipo          | Descripción                                     |
| ----------- | ------------- | ----------------------------------------------- |
| `id`        | `bigint` (PK) | Identificador único para cada mascota.          |
| `client_id` | `bigint` (FK) | Referencia al dueño de la mascota (`clients.id`). |
| `name`      | `string`      | Nombre de la mascota.                           |
| `species`   | `string`      | Especie del animal (ej. "Canino", "Felino").    |
| `breed`     | `string`      | Raza de la mascota.                             |
| `birth_date`| `date`        | Fecha de nacimiento de la mascota.              |
| `gender`    | `string`      | Sexo de la mascota (ej. "Macho", "Hembra").     |
| `created_at`| `datetime`    | Fecha y hora de creación del registro.          |
| `updated_at`| `datetime`    | Fecha y hora de la última actualización.        |

### Tabla: `appointments`

Gestiona las citas programadas en la clínica.

| Campo             | Tipo          | Descripción                                                       |
| ----------------- | ------------- | ----------------------------------------------------------------- |
| `id`              | `bigint` (PK) | Identificador único para cada cita.                               |
| `pet_id`          | `bigint` (FK) | Referencia a la mascota que asiste a la cita (`pets.id`).         |
| `user_id`         | `bigint` (FK) | Referencia al veterinario asignado a la cita (`users.id`).        |
| `appointment_time`| `datetime`    | Fecha y hora programada para la cita.                             |
| `status`          | `string`      | Estado de la cita (`programada`, `confirmada`, `cancelada`, `completada`). |
| `reason`          | `text`        | Motivo o descripción de la consulta.                              |
| `created_at`      | `datetime`    | Fecha y hora de creación del registro.                            |
| `updated_at`      | `datetime`    | Fecha y hora de la última actualización.                          |

### Tabla: `medical_records`

Contiene el historial médico detallado de cada mascota.

| Campo           | Tipo          | Descripción                                                     |
| --------------- | ------------- | --------------------------------------------------------------- |
| `id`            | `bigint` (PK) | Identificador único para cada registro clínico.                 |
| `pet_id`        | `bigint` (FK) | Referencia a la mascota (`pets.id`).                            |
| `appointment_id`| `bigint` (FK) | Referencia a la cita donde se generó este registro (`appointments.id`). |
| `diagnosis`     | `text`        | Diagnóstico realizado por el veterinario.                       |
| `treatment`     | `text`        | Tratamiento prescrito.                                          |
| `notes`         | `text`        | Notas adicionales del veterinario.                              |
| `created_at`    | `datetime`    | Fecha y hora de creación del registro.                          |
| `updated_at`    | `datetime`    | Fecha y hora de la última actualización.                        |

### Tabla: `payments`

Registra las transacciones financieras asociadas a las citas.

| Campo            | Tipo          | Descripción                                                       |
| ---------------- | ------------- | ----------------------------------------------------------------- |
| `id`             | `bigint` (PK) | Identificador único para cada pago.                               |
| `appointment_id` | `bigint` (FK) | Referencia a la cita asociada al pago (`appointments.id`).        |
| `amount`         | `decimal`     | Monto total del pago.                                             |
| `payment_method` | `string`      | Método de pago (`efectivo`, `tarjeta`, `transferencia`).          |
| `status`         | `string`      | Estado del pago (`pagado`, `pendiente`).                          |
| `transaction_date`| `datetime`    | Fecha y hora en que se realizó la transacción.                    |
| `created_at`     | `datetime`    | Fecha y hora de creación del registro.                            |
| `updated_at`     | `datetime`    | Fecha y hora de la última actualización.                          |

## 4. Tablas Auxiliares (ActiveStorage)

Para la gestión de archivos (fotos de mascotas, documentos, etc.), Ruby on Rails con ActiveStorage creará y gestionará automáticamente las siguientes tablas. No es necesario definirlas manualmente.

-   **`active_storage_blobs`**: Almacena los metadatos de cada archivo (nombre, tipo, tamaño, etc.).
-   **`active_storage_attachments`**: Tabla polimórfica que une los blobs (archivos) con los modelos de Rails que los poseen (ej. un `Pet` tiene una foto adjunta).
-   **`active_storage_variant_records`**: Almacena la información sobre las diferentes variantes de un mismo archivo (ej. thumbnails de una imagen).