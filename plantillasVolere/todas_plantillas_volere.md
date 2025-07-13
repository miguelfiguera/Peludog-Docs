<style>
@page {
    size: A4;
    margin: 1.5cm;
}

body {
    font-family: Arial, sans-serif;
    line-height: 1.4;
    color: #333;
    margin: 0;
    padding: 10px;
    background: white;
}

.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

.cover-page {
    text-align: center;
    padding: 50px 0;
    margin-bottom: 30px;
    border-bottom: 2px solid #2c3e50;
}

.cover-title {
    font-size: 32px;
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 15px;
}

.cover-subtitle {
    font-size: 22px;
    color: #7f8c8d;
    margin-bottom: 30px;
}

.cover-info {
    font-size: 15px;
    color: #2c3e50;
    line-height: 1.6;
}

/* This is the key change: hide the page break element for screen view */
.page-break {
    display: none;
}

.header {
    text-align: center;
    border-bottom: 2px solid #2c3e50;
    padding-bottom: 10px;
    margin-top: 30px; /* Add margin to separate templates */
    margin-bottom: 20px;
}

.volere-title {
    font-size: 24px;
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 8px;
}

.requisito-id {
    font-size: 18px;
    font-weight: bold;
    color: #e74c3c;
    background: #f8f9fa;
    padding: 6px 12px;
    border-radius: 4px;
    display: inline-block;
}

.section {
    margin: 15px 0;
    text-align: left;
}

.section-title {
    font-size: 16px;
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 8px;
    border-left: 3px solid #3498db;
    padding-left: 10px;
}

.section-content {
    font-size: 13px;
    padding: 5px 0;
    text-align: justify;
}

.metadata {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 15px;
    margin: 15px 0;
}

.metadata-item {
    background: #f8f9fa;
    padding: 10px;
    border-radius: 4px;
    border-left: 3px solid #3498db;
}

.metadata-label {
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 4px;
    font-size: 13px;
}

.criterios {
    background: #f8f9fa;
    padding: 15px;
    border-radius: 4px;
    border-left: 3px solid #27ae60;
}

.criterios ul {
    margin: 8px 0;
    padding-left: 18px;
}

.criterios li {
    margin-bottom: 6px;
    text-align: justify;
    font-size: 13px;
}

.footer {
    text-align: center;
    margin-top: 20px;
    padding-top: 10px;
    border-top: 1px solid #bdc3c7;
    font-size: 11px;
    color: #7f8c8d;
}

/* Estilos para requisitos no funcionales */
.nf-header {
    border-bottom: 2px solid #8e44ad;
}

.nf-title {
    color: #8e44ad;
}

.nf-section-title {
    color: #8e44ad;
    border-left: 3px solid #9b59b6;
}

.nf-metadata-item {
    border-left: 3px solid #9b59b6;
}

@media print {
    body {
        print-color-adjust: exact;
    }
    .page-break {
        display: block; /* Re-enable for printing */
        page-break-before: always;
    }
}
</style>

<div class="container">

<div class="cover-page">

<div class="cover-title">Plantillas Volere</div>

<div class="cover-subtitle">Proyecto Socio Tecnológico</div>

<div class="cover-info">
<strong>Plataforma CRM para PeluDog</strong><br/>
Requisitos Funcionales y No Funcionales<br/>
<br/>
<strong>Equipo:</strong> Dra. Génesis / Equipo PeluDog<br/>
<strong>Total de Requisitos:</strong> 41 Funcionales, 22 No Funcionales</div>

</div>

<div class="page-break"></div>

# Requisitos Funcionales

<div class="header">

## Plantilla Volere

### RF-GP001

</div>

## Requisito

**Registrar Nuevo Cliente**

### Metadatos

| Campo           | Valor                                                                                                                           |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                                      |
| **Prioridad**   | Alta - Necesario para iniciar el flujo de alta de clientes descrito en CU-GP01 y para disponer de datos de contacto confiables. |
| **Caso de Uso** | CU-GP01                                                                                                                         |

## Descripción

El Asistente o el Cliente/Propietario de Mascota podrán registrar la información de un nuevo dueño de mascota en el sistema.

## Criterios de Aceptación

- El sistema debe mostrar un formulario para el ingreso de datos del cliente (ej. nombre, apellido, DNI, teléfono, email, dirección).
- El sistema debe validar los campos obligatorios (ej. nombre, teléfono).
- Al guardar, el nuevo cliente debe crearse en la base de datos.
- El sistema debe permitir asociar una o más mascotas a este cliente durante el registro o posteriormente.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP002

</div>

## Requisito

**Registrar Nueva Mascota**

### Metadatos

| Campo           | Valor                                                                                                       |
| --------------- | ----------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                  |
| **Prioridad**   | Alta - Esencial para asociar mascotas a sus propietarios según CU-GP01 y mantener la base clínica completa. |
| **Caso de Uso** | CU-GP01                                                                                                     |

## Descripción

El Asistente, Cliente/Propietario de Mascota o Veterinario podrá registrar la información de una nueva mascota, asociándola a un cliente existente (o al propio cliente si es él quien realiza la acción).

## Criterios de Aceptación

- El sistema debe permitir seleccionar un cliente existente al cual asociar la mascota (si registra el Asistente/Veterinario) o asociarla automáticamente al perfil del Cliente/Propietario (si es él quien registra).
- El sistema debe mostrar un formulario para el ingreso de datos de la mascota (ej. nombre, especie, raza, sexo, fecha de nacimiento, número de chip, color, señas particulares, foto).
- El sistema debe validar los campos obligatorios (ej. nombre de mascota, especie).
- Al guardar, la nueva mascota debe crearse y vincularse correctamente al cliente.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP003

</div>

## Requisito

**Buscar Cliente y/o Mascota**

### Metadatos

| Campo           | Valor                                                                                                  |
| --------------- | ------------------------------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                             |
| **Prioridad**   | Alta - Facilita la localización rápida de clientes y mascotas durante la atención, soportando CU-GP01. |
| **Caso de Uso** | CU-GP01                                                                                                |

## Descripción

El Asistente o Veterinario podrá buscar clientes y/o mascotas existentes en el sistema utilizando diversos criterios.

## Criterios de Aceptación

- El sistema debe proveer campos de búsqueda para clientes (ej. por nombre, DNI, teléfono).
- El sistema debe proveer campos de búsqueda para mascotas (ej. por nombre, número de chip, nombre del dueño).
- Los resultados de la búsqueda deben mostrar información clave para identificar al cliente/mascota.
- El usuario debe poder seleccionar un cliente/mascota de los resultados para ver su información detallada.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP004

</div>

## Requisito

**Visualizar y Actualizar Datos de Cliente (Staff)**

### Metadatos

| Campo           | Valor                                                                                                                |
| --------------- | -------------------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                           |
| **Prioridad**   | Alta - Permite mantener actualizada la información de los clientes, mejorando la calidad de los datos según CU-GP01. |
| **Caso de Uso** | CU-GP01                                                                                                              |

## Descripción

El Asistente o Veterinario podrá ver y modificar la información de un cliente existente.

## Criterios de Aceptación

- Tras seleccionar un cliente, el sistema debe mostrar toda su información registrada.
- El sistema debe permitir la edición de los campos del cliente.
- El sistema debe guardar los cambios realizados en la base de datos.
- El sistema debe mostrar la lista de mascotas asociadas al cliente, con acceso a sus perfiles.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP005

</div>

## Requisito

**Visualizar y Actualizar Datos de Mascota (Staff)**

### Metadatos

| Campo           | Valor                                                                                                          |
| --------------- | -------------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                     |
| **Prioridad**   | Alta - Asegura que la información de las mascotas esté actualizada para decisiones clínicas, acorde a CU-GP01. |
| **Caso de Uso** | CU-GP01                                                                                                        |

## Descripción

El Asistente o Veterinario podrá ver y modificar la información general de una mascota existente.

## Criterios de Aceptación

- Tras seleccionar una mascota, el sistema debe mostrar toda su información general registrada (no la historia clínica detallada, que es parte de otro CU).
- El sistema debe permitir la edición de los campos de la mascota.
- El sistema debe guardar los cambios realizados en la base de datos.
- El sistema debe mostrar la información del cliente dueño de la mascota.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP006

</div>

## Requisito

**Visualizar y Actualizar Datos de Mascota (Cliente/Propietario)**

### Metadatos

| Campo           | Valor                                                                                                                   |
| --------------- | ----------------------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                              |
| **Prioridad**   | Alta - Empodera al cliente para gestionar los datos de sus mascotas, reduciendo carga administrativa; parte de CU-GP01. |
| **Caso de Uso** | CU-GP01                                                                                                                 |

## Descripción

El Cliente/Propietario de Mascota podrá ver y modificar la información de sus propias mascotas registradas.

## Criterios de Aceptación

- Tras seleccionar una de sus mascotas, el sistema debe mostrar toda su información general registrada (no la historia clínica detallada).
- El sistema debe permitir la edición de los campos de la mascota por parte de su dueño.
- El sistema debe guardar los cambios realizados en la base de datos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GP007

</div>

## Requisito

**Visualizar y Actualizar Datos Personales (Cliente/Propietario)**

### Metadatos

| Campo           | Valor                                                                                                                       |
| --------------- | --------------------------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas)                                                                                  |
| **Prioridad**   | Alta - Permite al cliente mantener sus datos actualizados, garantizando la exactitud de contacto y cumplimiento de CU-GP01. |
| **Caso de Uso** | CU-GP01                                                                                                                     |

## Descripción

El Cliente/Propietario de Mascota podrá ver y modificar la información de su propio perfil.

## Criterios de Aceptación

- Tras acceder a su perfil, el sistema debe mostrar al Cliente/Propietario toda su información personal registrada.
- El sistema debe permitir la edición de sus propios campos.
- El sistema debe guardar los cambios realizados en la base de datos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA001

</div>

## Requisito

**Consultar Disponibilidad Horaria**

### Metadatos

| Campo           | Valor                                                                                                    |
| --------------- | -------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Agenda y Citas                                                                                |
| **Prioridad**   | Alta - Proporciona visibilidad de horarios disponibles indispensable para programar citas según CU-GA01. |
| **Caso de Uso** | CU-GA01                                                                                                  |

## Descripción

El Asistente o Cliente/Propietario de Mascota podrá consultar los horarios disponibles en la agenda del consultorio o de un veterinario específico.

## Criterios de Aceptación

- El sistema debe mostrar una vista de calendario (diaria, semanal, mensual).
- El sistema debe permitir filtrar la disponibilidad por veterinario (si hay más de uno).
- Los bloques horarios ocupados deben estar claramente diferenciados de los disponibles.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA002

</div>

## Requisito

**Agendar Nueva Cita**

### Metadatos

| Campo           | Valor                                                                                       |
| --------------- | ------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Agenda y Citas                                                                   |
| **Prioridad**   | Alta - Habilita el flujo de reserva de citas descrito en CU-GA01, pieza clave del servicio. |
| **Caso de Uso** | CU-GA01                                                                                     |

## Descripción

El Asistente o Cliente/Propietario de Mascota podrá reservar una nueva cita para una mascota, seleccionando un horario disponible.

## Criterios de Aceptación

- El sistema debe permitir seleccionar una mascota existente (o enlazar con CU-GP01/RF-GP002 si es nueva o el cliente gestiona sus mascotas).
- El sistema debe permitir seleccionar un veterinario (si aplica), fecha y hora de la cita de entre los disponibles.
- El sistema debe permitir ingresar el motivo de la cita.
- Al confirmar, la cita debe registrarse y el horario debe marcarse como ocupado.
- (Opcional) El sistema debe poder enviar una confirmación de cita al cliente (email/SMS).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA003

</div>

## Requisito

**Reprogramar Cita**

### Metadatos

| Campo           | Valor                                                                                                   |
| --------------- | ------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Agenda y Citas                                                                               |
| **Prioridad**   | Media - Ofrece flexibilidad a clientes y staff para cambios de agenda, evitando cancelaciones; CU-GA01. |
| **Caso de Uso** | CU-GA01                                                                                                 |

## Descripción

El Asistente o Cliente/Propietario de Mascota podrá cambiar la fecha y/o hora de una cita existente (el cliente solo las suyas).

## Criterios de Aceptación

- El sistema debe permitir buscar y seleccionar una cita existente.
- El sistema debe permitir elegir una nueva fecha/hora disponible.
- Al confirmar, la cita original debe actualizarse y el bloqueo horario debe moverse.
- (Opcional) El sistema debe poder enviar una notificación de reprogramación al cliente.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA004

</div>

## Requisito

**Cancelar Cita**

### Metadatos

| Campo           | Valor                                                                                   |
| --------------- | --------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Agenda y Citas                                                               |
| **Prioridad**   | Media - Gestiona anulaciones de citas liberando horarios y evitando ausencias; CU-GA01. |
| **Caso de Uso** | CU-GA01                                                                                 |

## Descripción

El Asistente o Cliente/Propietario de Mascota podrá cancelar una cita agendada (el cliente solo las suyas).

## Criterios de Aceptación

- El sistema debe permitir buscar y seleccionar una cita existente.
- Al confirmar la cancelación, la cita debe marcarse como cancelada y el horario liberarse.
- (Opcional) El sistema debe poder enviar una notificación de cancelación al cliente.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA005

</div>

## Requisito

**Consultar Agenda (Vista Staff)**

### Metadatos

| Campo           | Valor                                                                                      |
| --------------- | ------------------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Agenda y Citas                                                                  |
| **Prioridad**   | Alta - Permite al personal visualizar la carga de trabajo y confirmar asistencia; CU-GA01. |
| **Caso de Uso** | CU-GA01                                                                                    |

## Descripción

El Asistente o Veterinario podrá visualizar la agenda de citas programadas.

## Criterios de Aceptación

- El sistema debe mostrar una vista de calendario con todas las citas (diaria, semanal, mensual).
- Debe ser posible filtrar por veterinario.
- Cada cita en la agenda debe mostrar información clave (hora, mascota, cliente, motivo).
- El sistema debe permitir al Asistente confirmar la asistencia del cliente a una cita.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-GA006

</div>

## Requisito

**Enviar Recordatorio Automático de Cita**

### Metadatos

| Campo           | Valor                                                                                                      |
| --------------- | ---------------------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Agenda y Citas                                                                                  |
| **Prioridad**   | Alta - Reduce inasistencias enviando recordatorios automáticos; vital para eficiencia operativa (CU-GA02). |
| **Caso de Uso** | CU-GA02                                                                                                    |

## Descripción

El sistema enviará recordatorios automáticos a los clientes sobre sus próximas citas.

## Criterios de Aceptación

- El sistema debe permitir configurar el canal de envío (ej. email, SMS).
- El sistema debe permitir configurar la antelación del recordatorio (ej. 24 horas antes).
- El recordatorio debe incluir los detalles de la cita (fecha, hora, mascota).
- (Opcional) El recordatorio puede incluir un enlace para confirmar o cancelar la cita (enlazando con RF-GA003, RF-GA004).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC001

</div>

## Requisito

**Visualizar Historial Clínico Completo**

### Metadatos

| Campo           | Valor                                                                                      |
| --------------- | ------------------------------------------------------------------------------------------ |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                            |
| **Prioridad**   | Alta - Permite revisión completa de antecedentes clínicos para mejor diagnóstico; CU-AC01. |
| **Caso de Uso** | CU-AC01                                                                                    |

## Descripción

El Veterinario podrá ver todas las entradas previas del historial clínico de la mascota seleccionada.

## Criterios de Aceptación

- El sistema debe mostrar las entradas del historial en orden cronológico inverso (o permitir ordenamiento).
- Cada entrada debe mostrar como mínimo: fecha, motivo de consulta, diagnóstico(s) principal(es), y veterinario que atendió.
- El usuario debe poder expandir una entrada para ver todos sus detalles (anamnesis, examen, tratamientos, notas, etc.).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC002

</div>

## Requisito

**Visualizar Resumen de Alergias y Datos Relevantes**

### Metadatos

| Campo           | Valor                                                                                  |
| --------------- | -------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                        |
| **Prioridad**   | Alta - Destaca información crítica (alergias) mejorando la seguridad clínica; CU-AC01. |
| **Caso de Uso** | CU-AC01                                                                                |

## Descripción

El sistema mostrará información crucial de la mascota de forma destacada al consultar su historial.

## Criterios de Aceptación

- El sistema debe mostrar claramente las alergias conocidas de la mascota, si existen.
- El sistema debe mostrar datos básicos de la mascota (especie, raza, sexo, edad/fecha de nacimiento, último peso registrado).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC003

</div>

## Requisito

**Visualizar Archivos Adjuntos del Historial**

### Metadatos

| Campo           | Valor                                                                           |
| --------------- | ------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                 |
| **Prioridad**   | Alta - Facilita acceso a pruebas complementarias adjuntas, esencial en CU-AC01. |
| **Caso de Uso** | CU-AC01                                                                         |

## Descripción

El Veterinario podrá acceder a los archivos digitales asociados a las entradas del historial clínico.

## Criterios de Aceptación

- Para cada entrada del historial, el sistema debe listar los archivos adjuntos.
- El sistema debe permitir la visualización o descarga de estos archivos.
- Se deben mostrar las descripciones o etiquetas de los archivos adjuntos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC004

</div>

## Requisito

**Consultar Registros Específicos de Vacunación y Desparasitación**

### Metadatos

| Campo           | Valor                                                                                       |
| --------------- | ------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                             |
| **Prioridad**   | Media - Agiliza consulta de vacunas/desparasitaciones para cumplimiento sanitario; CU-AC01. |
| **Caso de Uso** | CU-AC01                                                                                     |

## Descripción

El Veterinario podrá filtrar o acceder a una vista específica del historial para ver únicamente los registros de vacunación y desparasitación.

## Criterios de Aceptación

- El sistema debe permitir filtrar el historial para mostrar solo eventos de vacunación.
- Para cada vacuna: mostrar fecha, nombre de vacuna, lote, vencimiento, próxima revacunación.
- El sistema debe permitir filtrar el historial para mostrar solo eventos de desparasitación.
- Para cada desparasitación: mostrar fecha, producto, dosis, próxima desparasitación.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC005

</div>

## Requisito

**Iniciar Nueva Entrada de Historial Clínico**

### Metadatos

| Campo           | Valor                                                                                 |
| --------------- | ------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                       |
| **Prioridad**   | Alta - Inicia la documentación de una consulta, base de la historia clínica; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                               |

## Descripción

El Veterinario podrá crear una nueva entrada en el historial clínico para la mascota en consulta.

## Criterios de Aceptación

- El sistema debe permitir seleccionar la mascota (si no viene de una cita preseleccionada).
- La fecha y hora de la consulta deben autocompletarse con la actual, pero ser editables.
- La entrada debe asociarse al veterinario que está logueado.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC006

</div>

## Requisito

**Registrar Datos de la Consulta**

### Metadatos

| Campo           | Valor                                                                                           |
| --------------- | ----------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                                 |
| **Prioridad**   | Alta - Captura la información detallada de la consulta asegurando continuidad clínica; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                                         |

## Descripción

El Veterinario podrá registrar toda la información clínica relevante de la consulta.

## Criterios de Aceptación

- El sistema debe proveer campos para: Peso, Temperatura.
- El sistema debe proveer campo para Motivo de la consulta (texto libre y/o selección predefinida).
- El sistema debe proveer campo para Anamnesis (texto libre).
- El sistema debe proveer campo para Hallazgos del examen físico (texto libre, opcionalmente estructurado por sistema).
- El sistema debe proveer campo para Diagnóstico(s) (texto libre y/o selección de catálogo - ej. CIE-Vet como opcional RF023/RF027).
- El sistema debe proveer campo para Plan de tratamiento (texto libre general).
- El sistema debe proveer campo para Notas adicionales.
- Al guardar, la nueva entrada se añade al historial de la mascota.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC007

</div>

## Requisito

**Prescribir Tratamiento (Generar Récipe)**

### Metadatos

| Campo           | Valor                                                                               |
| --------------- | ----------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                     |
| **Prioridad**   | Alta - Estandariza y registra prescripciones, generando récipe imprimible; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                             |

## Descripción

Dentro de una consulta, el Veterinario podrá añadir prescripciones de medicamentos y generar un récipe.

## Criterios de Aceptación

- El sistema debe permitir añadir uno o más medicamentos a la prescripción.
- Para cada medicamento: registrar nombre, presentación/concentración, dosis, vía, frecuencia, duración.
- (Opcional Medio) El sistema puede ofrecer un catálogo de medicamentos para autocompletar.
- Las prescripciones deben quedar guardadas como parte de la entrada del historial clínico.
- El sistema debe permitir generar un documento de récipe (PDF imprimible) con datos del consultorio, veterinario, dueño, mascota, fecha y medicamentos prescritos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC008

</div>

## Requisito

**Solicitar Exámenes de Laboratorio**

### Metadatos

| Campo           | Valor                                                                                         |
| --------------- | --------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                               |
| **Prioridad**   | Alta - Permite solicitudes de laboratorio integradas en la historia, evitando papel; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                                       |

## Descripción

Dentro de una consulta, el Veterinario podrá solicitar exámenes de laboratorio y generar una orden.

## Criterios de Aceptación

- El sistema debe permitir seleccionar o ingresar los exámenes solicitados.
- (Opcional Medio) El sistema puede ofrecer un catálogo de exámenes disponibles.
- Las solicitudes de exámenes deben quedar guardadas como parte de la entrada del historial clínico.
- El sistema debe permitir generar una orden de exámenes (PDF imprimible) con datos del consultorio, veterinario, dueño, mascota, fecha y exámenes solicitados.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC009

</div>

## Requisito

**Adjuntar Archivos a la Entrada Clínica**

### Metadatos

| Campo           | Valor                                                                             |
| --------------- | --------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                   |
| **Prioridad**   | Alta - Conserva evidencias (imágenes, documentos) ligadas a la consulta; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                           |

## Descripción

Durante la creación o edición de una entrada clínica, el Veterinario podrá adjuntar archivos digitales.

## Criterios de Aceptación

- El sistema debe permitir cargar uno o más archivos.
- El sistema debe soportar formatos comunes (PDF, JPG, PNG, DOCX, DICOM).
- Los archivos deben quedar vinculados a la entrada específica del historial.
- El sistema debe permitir añadir una descripción/etiqueta a cada archivo.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC010

</div>

## Requisito

**Registrar Aplicación de Vacunas y Desparasitantes**

### Metadatos

| Campo           | Valor                                                                                        |
| --------------- | -------------------------------------------------------------------------------------------- |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                                              |
| **Prioridad**   | Alta - Registra vacunaciones y desparasitaciones con recordatorios futuros; CU-AC02 & AC004. |
| **Caso de Uso** | CU-AC02                                                                                      |

## Descripción

El Veterinario podrá registrar la aplicación de vacunas o desparasitantes como parte de la consulta.

## Criterios de Aceptación

- El sistema debe permitir seleccionar/ingresar el nombre de la vacuna, lote, fecha de vencimiento del producto.
- El sistema debe permitir calcular y/o registrar la próxima fecha de revacunación.
- El sistema debe permitir seleccionar/ingresar el nombre del desparasitante, dosis.
- El sistema debe permitir calcular y/o registrar la próxima fecha de desparasitación.
- Esta información debe quedar registrada en la entrada del historial clínico y ser visible en la consulta de registros específicos (RF-AC004).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-AC011

</div>

## Requisito

**Editar Entrada de Historial Clínico Reciente**

### Metadatos

| Campo           | Valor                                                                    |
| --------------- | ------------------------------------------------------------------------ |
| **Módulo/Área** | Atención Clínica (Historia Clínica Electrónica)                          |
| **Prioridad**   | Media - Ofrece ventana de corrección con trazabilidad auditada; CU-AC02. |
| **Caso de Uso** | CU-AC02                                                                  |

## Descripción

El Veterinario podrá modificar una entrada de historial clínico recientemente creada, con auditoría.

## Criterios de Aceptación

- El sistema debe permitir la edición de una entrada dentro de un período configurable (ej. 24h, antes del cierre diario).
- Cualquier modificación debe registrarse en una bitácora de auditoría (quién, qué, cuándo).
- Los campos críticos (ej. diagnóstico, prescripción) podrían tener restricciones adicionales o requerir justificación.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH001

</div>

## Requisito

**Registrar Nuevo Empleado**

### Metadatos

| Campo           | Valor                                                                        |
| --------------- | ---------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                             |
| **Prioridad**   | Alta - Impulsa la gestión de personal creando cuentas de empleados; CU-RH01. |
| **Caso de Uso** | CU-RH01                                                                      |

## Descripción

El Administrador/Gerente podrá registrar la información de un nuevo empleado en el sistema.

## Criterios de Aceptación

- El sistema debe mostrar un formulario para el ingreso de datos del empleado (ej. nombre, apellido, DNI, teléfono, email, dirección, fecha de contratación).
- El sistema debe validar los campos obligatorios (ej. nombre, DNI, rol).
- Al guardar, el nuevo empleado debe crearse en la base de datos.
- El sistema debe permitir asignar uno o más roles al empleado (ej. Veterinario, Asistente, Peluquero Canino, Limpieza, Administrativo).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH002

</div>

## Requisito

**Visualizar y Actualizar Datos de Empleado**

### Metadatos

| Campo           | Valor                                                                                |
| --------------- | ------------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                                     |
| **Prioridad**   | Alta - Mantiene actualizados los datos del personal para procesos internos; CU-RH01. |
| **Caso de Uso** | CU-RH01                                                                              |

## Descripción

El Administrador/Gerente podrá ver y modificar la información de un empleado existente.

## Criterios de Aceptación

- Tras seleccionar un empleado, el sistema debe mostrar toda su información registrada.
- El sistema debe permitir la edición de los campos del empleado.
- El sistema debe guardar los cambios realizados en la base de datos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH003

</div>

## Requisito

**Asignar/Modificar Rol de Empleado**

### Metadatos

| Campo           | Valor                                                                          |
| --------------- | ------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                               |
| **Prioridad**   | Alta - Controla roles y permisos del personal garantizando seguridad; CU-RH01. |
| **Caso de Uso** | CU-RH01                                                                        |

## Descripción

El Administrador/Gerente podrá asignar o modificar los roles de un empleado.

## Criterios de Aceptación

- El sistema debe permitir seleccionar uno o varios roles predefinidos para un empleado.
- El sistema debe actualizar los permisos de acceso del empleado según su rol.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH004

</div>

## Requisito

**Desactivar/Activar Empleado**

### Metadatos

| Campo           | Valor                                                                                     |
| --------------- | ----------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                                          |
| **Prioridad**   | Media - Permite gestionar altas/bajas de empleados asegurando control de acceso; CU-RH01. |
| **Caso de Uso** | CU-RH01                                                                                   |

## Descripción

El Administrador/Gerente podrá desactivar temporal o permanentemente la cuenta de un empleado, impidiendo su acceso al sistema, o reactivarla.

## Criterios de Aceptación

- El sistema debe permitir cambiar el estado de un empleado a 'activo' o 'inactivo'.
- Los empleados inactivos no deben poder iniciar sesión.
- El historial de acciones del empleado debe conservarse.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH005

</div>

## Requisito

**Crear/Modificar Turno**

### Metadatos

| Campo           | Valor                                                            |
| --------------- | ---------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                 |
| **Prioridad**   | Alta - Define turnos tipo necesarios para programación; CU-RH02. |
| **Caso de Uso** | CU-RH02                                                          |

## Descripción

El Administrador/Gerente podrá definir y modificar los tipos de turnos y guardias (ej. 'Turno Mañana', 'Guardia Nocturna', 'Descanso').

## Criterios de Aceptación

- El sistema debe permitir definir un nombre, hora de inicio y fin para cada tipo de turno.
- El sistema debe permitir asociar un turno a un día específico o a un rango de fechas.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH006

</div>

## Requisito

**Asignar Turno a Empleado**

### Metadatos

| Campo           | Valor                                                                        |
| --------------- | ---------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                             |
| **Prioridad**   | Alta - Asigna turnos a empleados, clave en planificación operativa; CU-RH02. |
| **Caso de Uso** | CU-RH02                                                                      |

## Descripción

El Administrador/Gerente podrá asignar turnos y guardias a empleados específicos para un período determinado.

## Criterios de Aceptación

- El sistema debe permitir seleccionar un empleado y asignarle uno o más turnos para días específicos.
- El sistema debe mostrar posibles conflictos de horario al asignar turnos.
- El sistema debe guardar las asignaciones de turnos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH007

</div>

## Requisito

**Consultar Horario de Empleado**

### Metadatos

| Campo           | Valor                                                                           |
| --------------- | ------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                                |
| **Prioridad**   | Alta - Da visibilidad al personal de su horario, reduciendo confusión; CU-RH02. |
| **Caso de Uso** | CU-RH02                                                                         |

## Descripción

Cualquier empleado podrá consultar su propio horario de turnos y guardias asignados. El Administrador/Gerente podrá consultar el horario de cualquier empleado.

## Criterios de Aceptación

- El sistema debe mostrar una vista de calendario con los turnos asignados al empleado.
- Debe ser posible filtrar por semana o mes.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH008

</div>

## Requisito

**Consultar Disponibilidad de Personal**

### Metadatos

| Campo           | Valor                                                                      |
| --------------- | -------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                           |
| **Prioridad**   | Media - Brinda vista global de disponibilidad para planificación; CU-RH02. |
| **Caso de Uso** | CU-RH02                                                                    |

## Descripción

El Administrador/Gerente podrá consultar la disponibilidad general del personal para planificar turnos o asignar tareas.

## Criterios de Aceptación

- El sistema debe mostrar una vista consolidada de los turnos de todo el personal.
- Debe ser posible filtrar por rol o por fecha.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH009

</div>

## Requisito

**Crear/Modificar Tipo de Servicio Interno**

### Metadatos

| Campo           | Valor                                                                          |
| --------------- | ------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                               |
| **Prioridad**   | Alta - Permite definir servicios internos que mejoran oferta clínica; CU-RH03. |
| **Caso de Uso** | CU-RH03                                                                        |

## Descripción

El Administrador/Gerente podrá definir nuevos tipos de servicios que la clínica ofrece (ej. 'Servicio de Emergencia', 'Farmacia', 'Delivery de Medicamentos', 'Atención a Domicilio', 'Hospedaje').

## Criterios de Aceptación

- El sistema debe permitir crear un nombre y una descripción para cada servicio.
- El sistema debe permitir activar o desactivar un servicio.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH010

</div>

## Requisito

**Asignar Personal a Servicio Interno**

### Metadatos

| Campo           | Valor                                                                             |
| --------------- | --------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                                  |
| **Prioridad**   | Media - Asigna responsables a servicios internos asegurando seguimiento; CU-RH03. |
| **Caso de Uso** | CU-RH03                                                                           |

## Descripción

El Administrador/Gerente podrá asignar personal específico como responsable o participante de un tipo de servicio interno.

## Criterios de Aceptación

- El sistema debe permitir seleccionar uno o más empleados para asociarlos a un tipo de servicio.
- Esta asignación puede influir en la visibilidad de tareas o reportes relacionados con ese servicio.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-RH011

</div>

## Requisito

**Consultar Servicios Internos Activos**

### Metadatos

| Campo           | Valor                                                                          |
| --------------- | ------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos                               |
| **Prioridad**   | Baja - Ofrece inventario de servicios activos para control y reporte; CU-RH03. |
| **Caso de Uso** | CU-RH03                                                                        |

## Descripción

El Administrador/Gerente podrá consultar la lista de servicios internos que la clínica tiene definidos y activos.

## Criterios de Aceptación

- El sistema debe mostrar una lista de todos los servicios internos con su estado (activo/inactivo) y descripción.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG001

</div>

## Requisito

**Registrar Pago de Servicio/Producto**

### Metadatos

| Campo           | Valor                                                                                   |
| --------------- | --------------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pagos y Facturación                                                          |
| **Prioridad**   | Alta - Registra ingresos de manera formal y asocia servicios, base financiera; CU-PG01. |
| **Caso de Uso** | CU-PG01                                                                                 |

## Descripción

El Asistente podrá registrar un pago recibido de un cliente, asociándolo a uno o varios servicios o productos.

## Criterios de Aceptación

- El sistema debe permitir seleccionar el cliente y/o la mascota asociada al pago.
- El sistema debe permitir seleccionar los servicios o productos por los cuales se realiza el pago (idealmente vinculados a una factura o consulta previa).
- El sistema debe permitir ingresar el monto pagado.
- El sistema debe permitir seleccionar el método de pago (ej. efectivo, tarjeta de crédito/débito, transferencia, cheque).
- El sistema debe validar que el monto pagado sea consistente con el total a pagar (o registrar un saldo pendiente/a favor).
- Al guardar, el pago debe registrarse en la base de datos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG002

</div>

## Requisito

**Generar Comprobante de Pago**

### Metadatos

| Campo           | Valor                                                                 |
| --------------- | --------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pagos y Facturación                                        |
| **Prioridad**   | Alta - Genera notas de pago para transparencia con clientes; CU-PG01. |
| **Caso de Uso** | CU-PG01                                                               |

## Descripción

El sistema debe poder generar un comprobante o recibo de pago para el cliente.

## Criterios de Aceptación

- El comprobante debe incluir detalles del consultorio, fecha, monto, método de pago, y los servicios/productos pagados.
- El comprobante debe ser imprimible (ej. PDF).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG003

</div>

## Requisito

**Consultar Pagos por Cliente/Mascota (Staff)**

### Metadatos

| Campo           | Valor                                                                            |
| --------------- | -------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pagos y Facturación                                                   |
| **Prioridad**   | Alta - Permite rastrear pagos por cliente, apoyo a atención y cobranza; CU-PG02. |
| **Caso de Uso** | CU-PG02                                                                          |

## Descripción

El Asistente o Administrador/Gerente podrá consultar todos los pagos realizados por un cliente o para una mascota específica.

## Criterios de Aceptación

- El sistema debe permitir buscar pagos asociados a un cliente o mascota.
- Los resultados deben mostrar la fecha, monto, servicios/productos pagados y método de pago.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG004

</div>

## Requisito

**Consultar Pagos por Fecha/Servicio (Staff)**

### Metadatos

| Campo           | Valor                                                                          |
| --------------- | ------------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Pagos y Facturación                                                 |
| **Prioridad**   | Media - Facilita análisis histórico de ingresos por fechas/servicios; CU-PG02. |
| **Caso de Uso** | CU-PG02                                                                        |

## Descripción

El Administrador/Gerente podrá consultar pagos filtrados por rango de fechas o por tipo de servicio.

## Criterios de Aceptación

- El sistema debe permitir filtrar pagos por un rango de fechas.
- El sistema debe permitir filtrar pagos por uno o varios tipos de servicio/producto.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG006

</div>

## Requisito

**Consultar Propio Historial de Pagos y Facturas (Cliente/Propietario)**

### Metadatos

| Campo           | Valor                                                                           |
| --------------- | ------------------------------------------------------------------------------- |
| **Módulo/Área** | Gestión de Pagos y Facturación                                                  |
| **Prioridad**   | Media - Empodera al cliente a revisar sus pagos, fomentando confianza; CU-PG02. |
| **Caso de Uso** | CU-PG02                                                                         |

## Descripción

El Cliente/Propietario de Mascota podrá consultar su propio historial de pagos y descargar las facturas asociadas.

## Criterios de Aceptación

- El sistema debe mostrar al Cliente/Propietario una lista de sus pagos realizados, con fecha, monto y concepto.
- El sistema debe permitir al Cliente/Propietario visualizar y/o descargar las facturas o comprobantes asociados a sus pagos.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

<div class="header">

## Plantilla Volere

### RF-PG005

</div>

## Requisito

**Generar Reporte de Ingresos**

### Metadatos

| Campo           | Valor                                                                    |
| --------------- | ------------------------------------------------------------------------ |
| **Módulo/Área** | Gestión de Pagos y Facturación                                           |
| **Prioridad**   | Alta - Provee reportes de ingresos para decisiones financieras; CU-PG03. |
| **Caso de Uso** | CU-PG03                                                                  |

## Descripción

El sistema debe poder generar reportes de ingresos consolidados por períodos (diario, semanal, mensual, anual).

## Criterios de Aceptación

- El reporte debe mostrar el total de ingresos para el período seleccionado.
- El reporte puede desglosar los ingresos por tipo de servicio/producto o método de pago.
- El reporte debe ser exportable (ej. CSV, PDF).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="page-break"></div>

# Requisitos No Funcionales

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-USA-001

</div>

## Requisito

**Facilidad de Aprendizaje**

### Metadatos

| Campo         | Valor      |
| ------------- | ---------- |
| **Categoría** | Usabilidad |

## Descripción

Un veterinario o asistente nuevo, con conocimientos informáticos básicos, debería poder realizar tareas principales (p. ej., crear una nueva entrada clínica, consultar el historial, prescribir medicación) con menos de 2 horas de formación.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-USA-002

</div>

## Requisito

**Eficiencia de Uso**

### Metadatos

| Campo         | Valor      |
| ------------- | ---------- |
| **Categoría** | Usabilidad |

## Descripción

Una vez que dominen el sistema, los usuarios deben poder completar tareas comunes y frecuentes (p. ej., crear una receta estándar después de ingresar un diagnóstico) en un promedio de 3-5 interacciones del usuario (clics o selecciones).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-USA-003

</div>

## Requisito

**Satisfacción del Usuario**

### Metadatos

| Campo         | Valor      |
| ------------- | ---------- |
| **Categoría** | Usabilidad |

## Descripción

El sistema debería alcanzar una puntuación de satisfacción del usuario de al menos 7/10 en encuestas posteriores a la implementación centradas en la facilidad de uso y el soporte al flujo de trabajo.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-USA-004

</div>

## Requisito

**Prevención y Recuperación de Errores**

### Metadatos

| Campo         | Valor      |
| ------------- | ---------- |
| **Categoría** | Usabilidad |

## Descripción

El sistema debería proporcionar mensajes de confirmación claros para acciones críticas (p. ej., guardar un registro clínico, eliminar un archivo adjunto) y permitir la corrección fácil de errores de ingreso de datos antes del envío final, cuando sea apropiado. La eliminación de datos críticos debería requerir una confirmación explícita.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-PER-001

</div>

## Requisito

**Tiempo de Respuesta**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Rendimiento y Escalabilidad |

## Descripción

La carga del historial clínico completo de una mascota (hasta 50 entradas) no debería tardar más de 3 segundos. El guardado de una nueva entrada clínica (incluyendo datos básicos de consulta, una receta y una orden de laboratorio) debería completarse en 2 segundos. Los resultados de búsqueda de pacientes o clientes deberían aparecer en 2 segundos para consultas típicas.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-PER-002

</div>

## Requisito

**Usuarios Concurrentes**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Rendimiento y Escalabilidad |

## Descripción

El sistema debe soportar hasta 10 usuarios concurrentes activos (veterinarios y asistentes) realizando tareas clínicas y administrativas típicas sin una degradación notable del rendimiento (es decir, los tiempos de respuesta no deberían aumentar en más del 20% en comparación con un escenario de un solo usuario).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-PER-003

</div>

## Requisito

**Volumen de Datos**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Rendimiento y Escalabilidad |

## Descripción

El sistema debe gestionar y recuperar datos de manera eficiente para al menos 10.000 mascotas y 500.000 entradas de historial clínico durante 5 años sin una degradación significativa del rendimiento.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SCL-001

</div>

## Requisito

**Escalabilidad**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Rendimiento y Escalabilidad |

## Descripción

La arquitectura del sistema debería permitir un aumento del 50% en el número de usuarios y del 100% en el volumen de datos durante los próximos 3 años solo con escalado de hardware o cambios mínimos de configuración de software, sin requerir una re-arquitectura importante.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-AVA-001

</div>

## Requisito

**Tiempo de Actividad del Sistema**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Disponibilidad y Fiabilidad |

## Descripción

El módulo clínico debe estar disponible el 99.5% del tiempo durante las horas operativas definidas de la clínica (p. ej., de lunes a sábado, de 8 AM a 8 PM).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-REL-001

</div>

## Requisito

**Integridad de los Datos**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Disponibilidad y Fiabilidad |

## Descripción

Todos los datos clínicos guardados en el sistema deben almacenarse y recuperarse con precisión sin corrupción ni pérdida. Las transacciones (p. ej., guardar una entrada clínica) deben ser atómicas.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-REL-002

</div>

## Requisito

**Copia de Seguridad y Recuperación**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Disponibilidad y Fiabilidad |

## Descripción

Objetivo de Punto de Recuperación (RPO): En caso de una falla mayor, la pérdida de datos no debe exceder las 24 horas (copias de seguridad diarias). Objetivo de Tiempo de Recuperación (RTO): El sistema debería ser restaurable a un estado operativo en un plazo de 4 horas tras una falla mayor.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-REL-003

</div>

## Requisito

**Tolerancia a Fallos**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Disponibilidad y Fiabilidad |

## Descripción

El sistema debería manejar los errores comunes con elegancia (p. ej., interrupciones de red durante el envío de datos) e informar al usuario adecuadamente, intentando preservar los datos ingresados siempre que sea posible.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-REL-004

</div>

## Requisito

**Offline First**

### Metadatos

| Campo         | Valor                       |
| ------------- | --------------------------- |
| **Categoría** | Disponibilidad y Fiabilidad |

## Descripción

El sistema debera manejar alternativas para los casos donde se pierda la conexión con el servidor a través de localStorage tanto en movil como en web.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SEC-001

</div>

## Requisito

**Control de Acceso**

### Metadatos

| Campo         | Valor     |
| ------------- | --------- |
| **Categoría** | Seguridad |

## Descripción

El acceso a los datos clínicos del paciente debe estar restringido según los roles de usuario (p. ej., el Veterinario puede crear/editar/ver todos los datos clínicos; el Asistente puede tener acceso de solo lectura al historial clínico o acceso de escritura restringido a ciertas secciones).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SEC-002

</div>

## Requisito

**Confidencialidad de los Datos**

### Metadatos

| Campo         | Valor     |
| ------------- | --------- |
| **Categoría** | Seguridad |

## Descripción

Los datos sensibles de pacientes y clientes (información de identificación personal, registros clínicos) deben protegerse contra el acceso no autorizado. Los datos deberían estar encriptados en tránsito (p. ej., HTTPS) y en reposo (para la base de datos).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SEC-003

</div>

## Requisito

**Pistas de Auditoría**

### Metadatos

| Campo         | Valor     |
| ------------- | --------- |
| **Categoría** | Seguridad |

## Descripción

Todas las creaciones, modificaciones y eliminaciones de registros clínicos, recetas y órdenes de laboratorio deben registrarse con el ID de usuario, la marca de tiempo y los detalles del cambio.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SEC-004

</div>

## Requisito

**Autenticación**

### Metadatos

| Campo         | Valor     |
| ------------- | --------- |
| **Categoría** | Seguridad |

## Descripción

Los usuarios deben ser autenticados mediante un mecanismo seguro de nombre de usuario y contraseña. Se deben aplicar políticas de contraseña (complejidad, caducidad, bloqueo).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-SEC-005

</div>

## Requisito

**Validación de Entradas**

### Metadatos

| Campo         | Valor     |
| ------------- | --------- |
| **Categoría** | Seguridad |

## Descripción

El sistema debe validar todas las entradas del usuario para prevenir vulnerabilidades de seguridad comunes como la inyección SQL o Cross-Site Scripting (XSS).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-MAI-001

</div>

## Requisito

**Modularidad**

### Metadatos

| Campo         | Valor                           |
| ------------- | ------------------------------- |
| **Categoría** | Mantenibilidad y Compatibilidad |

## Descripción

El módulo clínico debería desarrollarse con un diseño modular para facilitar actualizaciones, correcciones de errores y futuras mejoras con un impacto mínimo en otras partes del CRM.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-MAI-002

</div>

## Requisito

**Testeabilidad**

### Metadatos

| Campo         | Valor                           |
| ------------- | ------------------------------- |
| **Categoría** | Mantenibilidad y Compatibilidad |

## Descripción

El sistema debería diseñarse para permitir pruebas automatizadas de funcionalidades clave.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-COM-001

</div>

## Requisito

**Compatibilidad con Navegadores**

### Metadatos

| Campo         | Valor                           |
| ------------- | ------------------------------- |
| **Categoría** | Mantenibilidad y Compatibilidad |

## Descripción

La interfaz de usuario debe ser completamente funcional y mostrarse correctamente en las dos últimas versiones de los principales navegadores web (p. ej., Chrome, Firefox, Edge, Safari).

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

<div class="page-break"></div>

<div class="header nf-header">

## Plantilla Volere - Requisito No Funcional

### RNF-COM-002

</div>

## Requisito

**Exportación de Datos**

### Metadatos

| Campo         | Valor                           |
| ------------- | ------------------------------- |
| **Categoría** | Mantenibilidad y Compatibilidad |

## Descripción

El sistema debería proporcionar un mecanismo para que los usuarios autorizados exporten datos clínicos de pacientes en un formato común y estructurado (p. ej., CSV, PDF) para mascotas individuales si se requiere para transferencia o por razones legales.

## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito no funcional asegura la calidad, rendimiento y confiabilidad del sistema CRM, contribuyendo a una experiencia de usuario óptima y operaciones estables del consultorio.

---

</div>
