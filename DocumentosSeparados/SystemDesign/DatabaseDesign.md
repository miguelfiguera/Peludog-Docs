# Diseño de la Base de Datos - PeluDog CRM

## 1. Introducción

Este documento describe la estructura y el diseño de la base de datos para el sistema PeluDog CRM. El diseño se basa en la información recopilada de los documentos de arquitectura del sistema.

La base de datos seleccionada es **MySQL 8.0+**, y la interacción con ella se realizará a través del ORM de Ruby on Rails, **ActiveRecord**. Esto significa que los nombres de las tablas seguirán la convención de pluralización de Rails y todas las tablas incluirán campos `id` como clave primaria y `created_at`/`updated_at` para el control de versiones.

## 2. Diagrama de Entidad-Relación

Un diagrama visual de la relación entre las tablas es fundamental para entender el modelo. El diseño aquí descrito se basa en el diagrama de clases que se encuentra en `Imagenes/DiagramaDeClases.png`. Se recomienda consultar dicho diagrama para una representación visual.

## 3. Definición de Entidades y Tablas

Tras un análisis detallado de los requerimientos funcionales, el modelo de base de datos se expande para dar soporte a todas las entidades de negocio. Se adopta un **modelo de usuarios unificado con perfiles separados**, y se añaden tablas específicas para los distintos registros clínicos y de gestión, en lugar de almacenarlos como texto simple.

### Modelos Principales y de Usuario

#### Tabla: `users`
Almacena la información de autenticación y los datos básicos de todas las entidades (personal y clientes).

| Campo             | Tipo          | Descripción                                                                      |
| ----------------- | ------------- | -------------------------------------------------------------------------------- |
| `id`              | `bigint` (PK) | Identificador único para cada usuario.                                           |
| `first_name`      | `string`      | Nombre de pila del usuario.                                                      |
| `last_name`       | `string`      | Apellido del usuario.                                                            |
| `email`           | `string`      | Correo electrónico para el inicio de sesión (único).                             |
| `password_digest` | `string`      | Hash de la contraseña para autenticación segura.                                 |
| `role`            | `string`      | Rol del usuario (`admin`, `vet_admin`, `veterinario`, `asistente`, `lab`, `cliente`). |
| `is_active`       | `boolean`     | Para desactivar usuarios sin borrarlos (default: true).                          |
| `last_login_at`   | `datetime`    | Registra la fecha del último inicio de sesión para auditoría.                    |
| `created_at`      | `datetime`    | Fecha y hora de creación del registro.                                           |
| `updated_at`      | `datetime`    | Fecha y hora de la última actualización.                                         |

#### Tabla: `client_profiles`
Contiene la información de contacto específica de los usuarios con el rol `cliente`.

| Campo          | Tipo          | Descripción                                          |
| -------------- | ------------- | ---------------------------------------------------- |
| `id`           | `bigint` (PK) | Identificador único para cada perfil.                |
| `user_id`      | `bigint` (FK) | Relación uno a uno con el usuario (`users.id`).      |
| `phone_number` | `string`      | Número de teléfono de contacto.                      |
| `address`      | `text`        | Dirección física del cliente.                        |
| `notes`        | `text`        | Notas internas sobre el cliente (ej. preferencias). |
| `created_at`   | `datetime`    | Fecha y hora de creación del registro.               |
| `updated_at`   | `datetime`    | Fecha y hora de la última actualización.             |

#### Tabla: `pets`
Almacena la información de cada mascota (paciente).

| Campo       | Tipo          | Descripción                                                                 |
| ----------- | ------------- | --------------------------------------------------------------------------- |
| `id`        | `bigint` (PK) | Identificador único para cada mascota.                                      |
| `user_id`   | `bigint` (FK) | Referencia al dueño de la mascota (`users.id` con rol `cliente`).           |
| `name`      | `string`      | Nombre de la mascota.                                                       |
| `species`   | `string`      | Especie del animal (ej. "Canino", "Felino").                                |
| `breed`     | `string`      | Raza de la mascota.                                                         |
| `birth_date`| `date`        | Fecha de nacimiento de la mascota.                                          |
| `gender`    | `string`      | Sexo de la mascota (ej. "Macho", "Hembra").                                 |
| `color`     | `string`      | Color del pelaje de la mascota.                                                  |
| `chip_number` | `string`      | Número del microchip de identificación (si aplica).                              |
| `deleted_at`| `datetime`    | Para borrado lógico. Si tiene fecha, la mascota está "archivada".               |
| `created_at`| `datetime`    | Fecha y hora de creación del registro.                                      |
| `updated_at`| `datetime`    | Fecha y hora de la última actualización.                                      |

### Modelos de Gestión Clínica

#### Tabla: `appointments`
Gestiona las citas programadas en la clínica.

| Campo              | Tipo          | Descripción                                                                |
| ------------------ | ------------- | -------------------------------------------------------------------------- |
| `id`               | `bigint` (PK) | Identificador único para cada cita.                                        |
| `pet_id`           | `bigint` (FK) | Referencia a la mascota que asiste a la cita (`pets.id`).                  |
| `service_id`       | `bigint` (FK) | Referencia al servicio principal de la cita (`services.id`).               |
| `payment_note_id`  | `bigint` (FK) | Opcional. Referencia a la nota de pago que incluye esta cita (`payment_notes.id`). |
| `appointment_time` | `datetime`    | Fecha y hora programada para la cita.                                      |
| `status`           | `string`      | Estado de la cita (`programada`, `confirmada`, `cancelada`, `completada`). |
| `type`             | `string`      | Tipo de cita (`en_clinica`, `a_domicilio`, `telemedicina`).                |
| `reason`           | `text`        | Motivo o descripción de la consulta.                                       |
| `created_at`       | `datetime`    | Fecha y hora de creación del registro.                                     |
| `updated_at`       | `datetime`    | Fecha y hora de la última actualización.                                   |

#### Tabla: `appointment_staff` (Tabla de Unión)
Asocia uno o más veterinarios/asistentes a una cita.

| Campo            | Tipo          | Descripción                                           |
| ---------------- | ------------- | ----------------------------------------------------- |
| `id`             | `bigint` (PK) | Identificador único.                                  |
| `appointment_id` | `bigint` (FK) | Referencia a la cita (`appointments.id`).             |
| `user_id`        | `bigint` (FK) | Referencia al empleado asignado a la cita (`users.id`).|

#### Tabla: `medical_records`
Registra los detalles estructurados de una consulta o evento clínico.

| Campo                    | Tipo          | Descripción                                                                    |
| ------------------------ | ------------- | ------------------------------------------------------------------------------ |
| `id`                     | `bigint` (PK) | Identificador único para cada registro clínico.                                |
| `appointment_id`         | `bigint` (FK) | Referencia a la cita donde se generó este registro (`appointments.id`).        |
| `weight`                 | `decimal`     | Peso de la mascota en la consulta (ej. en kg).                                 |
| `temperature`            | `decimal`     | Temperatura de la mascota en la consulta (ej. en °C).                          |
| `anamnesis`              | `text`        | Historial y síntomas reportados por el dueño.                                  |
| `physical_exam_findings` | `text`        | Hallazgos del examen físico realizado por el veterinario.                      |
| `diagnosis`              | `text`        | Diagnóstico(s) principal(es) establecidos por el veterinario.                  |
| `treatment_plan`         | `text`        | Descripción del plan de tratamiento general.                                   |
| `notes`                  | `text`        | Notas adicionales o comentarios.                                               |
| `created_at`             | `datetime`    | Fecha y hora de creación del registro.                                         |
| `updated_at`             | `datetime`    | Fecha y hora de la última actualización.                                       |

**Nota sobre la edición:** Por reglas de negocio y para garantizar la integridad del historial, los registros médicos no podrán ser editados después de 7 días desde su creación. La única excepción será la tabla `lab_orders` para poder adjuntar los resultados de laboratorio posteriormente.

### Modelos de Detalle Clínico

Estas tablas añaden granularidad a los registros médicos.

#### Tabla: `prescriptions`

| Campo               | Tipo          | Descripción                                           |
| ------------------- | ------------- | ----------------------------------------------------- |
| `id`                | `bigint` (PK) | Identificador único.                                  |
| `medical_record_id` | `bigint` (FK) | Referencia al registro médico (`medical_records.id`). |
| `drug_name`         | `string`      | Nombre del medicamento.                               |
| `dosage`            | `string`      | Dosis indicada (ej. "1 comprimido", "5 ml").        |
| `frequency`         | `string`      | Frecuencia (ej. "Cada 12 horas").                   |
| `duration`          | `string`      | Duración del tratamiento (ej. "Por 7 días").        |

#### Tabla: `vaccination_records`

| Campo               | Tipo          | Descripción                                           |
| ------------------- | ------------- | ----------------------------------------------------- |
| `id`                | `bigint` (PK) | Identificador único.                                  |
| `pet_id`            | `bigint` (FK) | Referencia a la mascota (`pets.id`).                  |
| `medical_record_id` | `bigint` (FK) | Opcional, si se aplicó en consulta (`medical_records.id`). |
| `vaccine_name`      | `string`      | Nombre de la vacuna o desparasitante.                 |
| `lot_number`        | `string`      | Número de lote del producto.                          |
| `applied_at`        | `date`        | Fecha de aplicación.                                  |
| `next_due_date`     | `date`        | Fecha del próximo refuerzo.                           |

#### Tabla: `lab_orders`

| Campo               | Tipo          | Descripción                                           |
| ------------------- | ------------- | ----------------------------------------------------- |
| `id`                | `bigint` (PK) | Identificador único.                                  |
| `medical_record_id` | `bigint` (FK) | Referencia al registro médico (`medical_records.id`). |
| `test_name`         | `string`      | Nombre del exámen solicitado.                         |
| `status`            | `string`      | Estado (`solicitado`, `resultados_recibidos`).        |
| `results_summary`   | `text`        | Resumen de los resultados (opcional).                 |

### Modelos de Gestión y Administración

#### Tabla: `services`

| Campo         | Tipo          | Descripción                                           |
| ------------- | ------------- | ----------------------------------------------------- |
| `id`          | `bigint` (PK) | Identificador único.                                  |
| `name`        | `string`      | Nombre del servicio (ej. "Consulta", "Peluquería"). |
| `description` | `text`        | Descripción del servicio.                             |
| `price`       | `decimal`     | Precio base del servicio.                             |
| `is_active`   | `boolean`     | Indica si el servicio se ofrece actualmente.          |

#### Tabla: `products`

| Campo         | Tipo          | Descripción                                           |
| ------------- | ------------- | ----------------------------------------------------- |
| `id`          | `bigint` (PK) | Identificador único.                                  |
| `name`        | `string`      | Nombre del producto (ej. "Shampoo Medicado").       |
| `description` | `text`        | Descripción del producto.                             |
| `price`       | `decimal`     | Precio de venta del producto.                         |
| `is_active`   | `boolean`     | Indica si el producto está a la venta.                |

#### Tabla: `shift_templates` (Plantillas de Turno)
Define los bloques de horario reusables para el personal.

| Campo         | Tipo          | Descripción                                           |
| ------------- | ------------- | ----------------------------------------------------- |
| `id`          | `bigint` (PK) | Identificador único.                                  |
| `name`        | `string`      | Nombre descriptivo (ej. "Turno Mañana L-V").        |
| `day_of_week` | `integer`     | Día de la semana (0=Domingo, 1=Lunes, ...).           |
| `start_time`  | `time`        | Hora de inicio del turno (ej. `09:00`).               |
| `end_time`    | `time`        | Hora de fin del turno (ej. `17:00`).                  |

#### Tabla: `user_schedules` (Horarios de Usuario)
Asigna una plantilla de turno a un empleado para un período determinado.

| Campo               | Tipo          | Descripción                                                |
| ------------------- | ------------- | ---------------------------------------------------------- |
| `id`                | `bigint` (PK) | Identificador único.                                       |
| `user_id`           | `bigint` (FK) | Referencia al empleado (`users.id`).                       |
| `shift_template_id` | `bigint` (FK) | Referencia a la plantilla de turno (`shift_templates.id`). |
| `start_date`        | `date`        | Fecha en la que empieza a aplicar este horario.            |
| `end_date`          | `date`        | Opcional. Fecha en la que termina el horario.              |

### Modelos Financieros

Para manejar las transacciones de forma estructurada, se introduce el concepto de una "Nota de Pago" que funciona como una factura o recibo detallado.

#### Tabla: `payment_notes` (Notas de Pago)
Esta tabla es el corazón del módulo financiero. Cada nota agrupa todos los cargos de una o más citas.

| Campo            | Tipo          | Descripción                                                               |
| ---------------- | ------------- | ------------------------------------------------------------------------- |
| `id`             | `bigint` (PK) | Identificador único para la nota de pago.                                 |
| `user_id`        | `bigint` (FK) | Referencia al cliente al que se le emite la nota (`users.id`).            |
| `total_amount`   | `decimal`     | El monto total calculado de todos los ítems en la nota.                   |
| `status`         | `string`      | Estado de la nota (`pendiente`, `pagada_parcialmente`, `pagada`, `anulada`). |
| `due_date`       | `date`        | Fecha de vencimiento para notas de pago a crédito.                               |
| `created_at`     | `datetime`    | Fecha y hora de creación del registro.                                    |
| `updated_at`     | `datetime`    | Fecha y hora de la última actualización.                                  |

#### Tabla: `payment_note_items` (Ítems de la Nota de Pago)
Esta es una tabla polimórfica que detalla cada línea de servicio o producto dentro de una nota de pago.

| Campo             | Tipo          | Descripción                                                               |
| ----------------- | ------------- | ------------------------------------------------------------------------- |
| `id`              | `bigint` (PK) | Identificador único.                                                      |
| `payment_note_id` | `bigint` (FK) | Referencia a la nota de pago a la que pertenece (`payment_notes.id`).     |
| `item_id`         | `bigint`      | ID del ítem (puede ser un ID de `services` o `products`).                 |
| `item_type`       | `string`      | El tipo de modelo del ítem ("Service" o "Product").                     |
| `quantity`        | `integer`     | Cantidad del servicio o producto.                                         |
| `unit_price`      | `decimal`     | El precio del ítem al momento de la venta.                                |

#### Tabla: `payments` (Pagos)
Registra las transacciones monetarias que se aplican a una nota de pago.

| Campo             | Tipo          | Descripción                                                               |
| ----------------- | ------------- | ------------------------------------------------------------------------- |
| `id`              | `bigint` (PK) | Identificador único para cada pago.                                       |
| `payment_note_id` | `bigint` (FK) | Referencia a la nota de pago que se está abonando (`payment_notes.id`).   |
| `amount`          | `decimal`     | Monto pagado en esta transacción.                                         |
| `payment_method`  | `string`      | Método de pago (`efectivo`, `tarjeta`, `transferencia`).                  |
| `transaction_id`  | `string`      | ID de referencia externo (ej. de una transferencia, Zelle, Stripe). Nulable. |
| `transaction_date`| `datetime`    | Fecha y hora en que se realizó la transacción.                            |
| `created_at`      | `datetime`    | Fecha y hora de creación del registro.                                    |
| `updated_at`      | `datetime`    | Fecha y hora de la última actualización.                                  |

## 4. Relaciones Entre Modelos

A continuación se describen las principales relaciones entre los modelos, usando una sintaxis similar a la de Ruby on Rails para mayor claridad.

### User
- `has_one :client_profile` (si el rol es 'cliente')
- `has_many :pets` (como dueño de las mascotas)
- `has_many :appointment_staff`
- `has_many :appointments, through: :appointment_staff` (como personal de la clínica)
- `has_many :user_schedules`
- `has_many :shift_templates, through: :user_schedules`
- `has_many :payment_notes` (como cliente)

### ClientProfile
- `belongs_to :user`

### Pet
- `belongs_to :user` (su dueño)
- `has_many :appointments`
- `has_many :vaccination_records`

### Appointment
- `belongs_to :pet`
- `belongs_to :service`
- `belongs_to :payment_note` (opcional)
- `has_one :medical_record`
- `has_many :appointment_staff`
- `has_many :staff, through: :appointment_staff, source: :user` (el personal asignado)

### MedicalRecord
- `belongs_to :appointment`
- `has_many :prescriptions`
- `has_many :lab_orders`

### VaccinationRecord
- `belongs_to :pet`
- `belongs_to :medical_record` (opcional)

### Service & Product
- `has_many :payment_note_items, as: :item` (relación polimórfica)

### ShiftTemplate
- `has_many :user_schedules`

### UserSchedule
- `belongs_to :user`
- `belongs_to :shift_template`

### PaymentNote
- `belongs_to :user` (el cliente)
- `has_many :appointments`
- `has_many :payment_note_items`
- `has_many :payments`

### PaymentNoteItem
- `belongs_to :payment_note`
- `belongs_to :item, polymorphic: true` (apunta a un Service o a un Product)

### Payment
- `belongs_to :payment_note`

## 5. Tablas Auxiliares (Gestionadas por Rails)

Además de los modelos de negocio principales, el framework de Ruby on Rails gestionará automáticamente tablas para ciertas funcionalidades.

### ActiveStorage (Gestión de Archivos)
Para la gestión de archivos (fotos de mascotas, resultados de laboratorio, etc.), Rails usará:

-   **`active_storage_blobs`**: Almacena los metadatos de cada archivo.
-   **`active_storage_attachments`**: Une los archivos con los modelos que los poseen.
-   **`active_storage_variant_records`**: Almacena información sobre las variantes de imágenes (ej. thumbnails).

### SolidQueue (Background Jobs)
Para tareas que no deben bloquear al usuario (como enviar emails), se usa SolidQueue, un sistema de trabajos en segundo plano basado en la base de datos. Rails creará y gestionará automáticamente las siguientes tablas para su funcionamiento:

- **`solid_queue_jobs`**: Almacena la información de los trabajos en cola.
- **`solid_queue_scheduled_executions`**: Para trabajos programados a futuro.
- **`solid_queue_ready_executions`**: Trabajos listos para ser procesados.
- **`solid_queue_claimed_executions`**: Trabajos que están siendo procesados por un worker.
- **`solid_queue_blocked_executions`**: Trabajos bloqueados por alguna condición.
- **`solid_queue_failed_executions`**: Registro de los trabajos que han fallado.
- **`solid_queue_pauses`**: Para pausar y reanudar colas de trabajo.
- **`solid_queue_semaphores`**: Para controlar la concurrencia de ciertos trabajos.

## 6. Índices de la Base de Datos

Para garantizar un rendimiento óptimo de las consultas a medida que la aplicación escala, se recomienda la creación de los siguientes índices en la base de datos. Los índices aceleran drásticamente la recuperación de datos, especialmente en operaciones de búsqueda y filtrado.

-   **`users`**
    -   `email`: **(Único)** Esencial para búsquedas rápidas y garantizar que no haya duplicados.
    -   `role`: Para filtrar eficientemente por tipo de usuario.

-   **`pets`**
    -   `user_id`: Para encontrar todas las mascotas de un cliente de forma instantánea.
    -   `deleted_at`: Para excluir eficientemente las mascotas archivadas (soft delete).

-   **`appointments`**
    -   `pet_id`, `service_id`, `payment_note_id`: Para acelerar las uniones con otras tablas.
    -   `appointment_time`: Crítico para las consultas de la agenda por rangos de fecha.
    -   `status`: Para filtrar rápidamente por el estado de la cita.

-   **`appointment_staff`**
    -   `[appointment_id, user_id]`: **(Compuesto y Único)** Para búsquedas rápidas y prevenir duplicados.

-   **`vaccination_records`**
    -   `pet_id`: Para el historial de vacunación de una mascota.
    -   `next_due_date`: Esencial para el sistema de recordatorios.

-   **`shift_templates`**
    -   `day_of_week`: Para buscar rápidamente las plantillas de un día específico.

-   **`user_schedules`**
    -   `[user_id, start_date, end_date]`: **(Compuesto)** Para buscar eficientemente los horarios activos de un empleado.

-   **`payment_notes`**
    -   `user_id` y `status`: Para buscar notas de pago por cliente y por estado.

-   **`payment_note_items`**
    -   `[item_id, item_type]`: **(Polimórfico)** Para encontrar todas las notas que incluyen un servicio o producto específico.

-   **`payments`**
    -   `payment_note_id`: Para encontrar los pagos de una nota.
    -   `transaction_id`: Para buscar transacciones externas.