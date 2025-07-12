### **Recopilación de Códigos PlantUML para Flujos Alternativos**

A continuación se presentan los códigos para modelar los flujos alternativos y de error de los principales casos de uso.

---

### **1. CU-GP01: Gestionar Información de Clientes y Mascotas**

#### **Flujo Alternativo: Cliente ya existe al intentar registrar**

```plantuml
@startuml
title FA-GP01-01: Cliente ya existe al intentar registrar

actor "Asistente/Cliente" as Actor
participant "Sistema" as Sistema

activate Actor
Actor -> Sistema: Inicia registro de Nuevo Cliente (RF-GP001)
activate Sistema

Sistema -> Actor: Presenta formulario de registro
Actor -> Sistema: Ingresa datos (incluyendo DNI ya existente) y envía
Sistema -> Sistema: Valida DNI y detecta duplicado
Sistema --> Actor: Muestra mensaje de error "El DNI ya está registrado"
deactivate Sistema

Actor -> Sistema: Solicita buscar cliente por DNI
activate Sistema
Sistema -> Sistema: Busca cliente
Sistema --> Actor: Muestra perfil del cliente existente
deactivate Sistema
deactivate Actor
@enduml
```

#### **Flujo Alternativo: Datos inválidos en formulario**

```plantuml
@startuml
title FA-GP01-02: Datos inválidos en formulario

actor "Actor" as Actor
participant "Sistema" as Sistema

activate Actor
Actor -> Sistema: Rellena formulario con datos inválidos (ej. email sin @)
activate Sistema
Sistema -> Sistema: Valida campos y detecta error de formato
Sistema --> Actor: Muestra mensaje de error específico junto al campo inválido
deactivate Sistema
deactivate Actor
@enduml
```

---

### **2. CU-GA01: Gestionar Citas**

#### **Flujo Alternativo: Horario seleccionado ya no está disponible**

```plantuml
@startuml
title FA-GA01-01: Horario no disponible al agendar

actor "Actor" as Actor
participant "Sistema" as Sistema

activate Actor
Actor -> Sistema: Consulta disponibilidad horaria (RF-GA001)
activate Sistema
Sistema --> Actor: Muestra calendario con horarios disponibles
deactivate Sistema

Actor -> Sistema: Selecciona un horario y envía para agendar (RF-GA002)
activate Sistema
Sistema -> Sistema: Verifica disponibilidad del horario en tiempo real
note right: Otro usuario agendó en el mismo instante.
Sistema --> Actor: Muestra mensaje "El horario seleccionado ya no está disponible. Por favor, elija otro."
Sistema -> Sistema: Refresca la vista del calendario
Sistema --> Actor: Muestra calendario actualizado
deactivate Sistema
deactivate Actor
@enduml
```

#### **Flujo Alternativo: Cancelación de cita fuera del plazo permitido**

```plantuml
@startuml
title FA-GA01-02: Cancelación fuera de plazo

actor "Cliente/Asistente" as Actor
participant "Sistema" as Sistema

activate Actor
Actor -> Sistema: Busca y selecciona una cita para cancelar (RF-GA004)
activate Sistema
Sistema -> Sistema: Verifica la fecha de la cita y la política de cancelación (ej. > 24h)
Sistema --> Actor: Muestra mensaje de error "No se puede cancelar la cita. Fuera del plazo permitido."
deactivate Sistema
deactivate Actor
@enduml
```

---

### **3. CU-AC02: Registrar Nueva Consulta**

#### **Flujo Alternativo: Intento de edición de historial clínico fuera de plazo**

```plantuml
@startuml
title FA-AC02-01: Edición de historial fuera de plazo

actor "Veterinario" as Veterinario
participant "Sistema" as Sistema

activate Veterinario
Veterinario -> Sistema: Selecciona una entrada de historial clínico para editar (RF-AC011)
activate Sistema
Sistema -> Sistema: Comprueba la fecha de la entrada y el período de edición permitido (ej. 24h)
Sistema --> Veterinario: Bloquea la edición y muestra el mensaje "No es posible editar esta entrada, ha superado el tiempo límite."
deactivate Sistema
deactivate Veterinario
@enduml
```

#### **Flujo Alternativo: Carga de archivo no soportado**

```plantuml
@startuml
title FA-AC02-02: Carga de archivo no soportado

actor "Veterinario" as Veterinario
participant "Sistema" as Sistema

activate Veterinario
Veterinario -> Sistema: Intenta adjuntar un archivo (RF-AC009)
activate Sistema
Sistema -> Actor: Muestra diálogo para seleccionar archivo
Actor -> Sistema: Selecciona un archivo con formato no permitido (ej. .exe) y lo carga
Sistema -> Sistema: Valida la extensión del archivo
Sistema --> Actor: Rechaza el archivo y muestra un mensaje de error "Formato de archivo no soportado. Por favor, use PDF, JPG, PNG."
deactivate Sistema
deactivate Veterinario
@enduml
```

---

### **4. CU-PG01: Registrar Pagos**

#### **Flujo Alternativo: Se registra un pago parcial**

```plantuml
@startuml
title FA-PG01-01: Registro de pago parcial

actor "Asistente" as Asistente
participant "Sistema" as Sistema

activate Asistente
Asistente -> Sistema: Inicia registro de pago para un servicio de 100$ (RF-PG001)
activate Sistema
Sistema --> Asistente: Muestra formulario de pago
Asistente -> Sistema: Ingresa un monto de 50$ y método de pago
Sistema -> Sistema: Valida que el monto es menor al total
Sistema -> Sistema: Registra el pago de 50$
Sistema -> Sistema: Calcula y actualiza el saldo pendiente a 50$
Sistema --> Asistente: Muestra confirmación de pago y el nuevo saldo pendiente
deactivate Sistema
deactivate Asistente
@enduml
```

---

### **5. CU-AC01: Consultar Historial Clínico**

#### **Flujo Alternativo: Búsqueda de mascota sin resultados**

Este es un flujo muy común que ocurre cuando el Asistente o Veterinario busca una mascota que no está registrada o comete un error al escribir el criterio de búsqueda.

```plantuml
@startuml
title FA-AC01-01: Búsqueda de mascota sin resultados

actor "Asistente/Veterinario" as Actor
participant "Sistema CRM" as Sistema

activate Actor
Actor -> Sistema: Busca mascota por nombre o ID (RF-AC001)
activate Sistema
Sistema -> Sistema: Realiza consulta en la base de datos
Sistema --> Actor: Muestra mensaje: "No se encontraron mascotas con el criterio de búsqueda."
deactivate Sistema
deactivate Actor
@enduml
```

---

### **6. CU-GA02: Gestionar Recordatorios de Citas**

#### **Flujo Alternativo: Falla en el servicio de envío de notificaciones**

Este flujo es crucial para la robustez del sistema. Ocurre cuando el servicio externo (para enviar emails o SMS) no está disponible o responde con un error.

```plantuml
@startuml
title FA-GA02-01: Falla en el servicio de envío de notificaciones

participant "Scheduler" as Scheduler
participant "Sistema CRM" as Sistema
participant "Servicio de Notificaciones" as Notificador

activate Scheduler
Scheduler -> Sistema: Solicita procesar recordatorios (RF-GA005)
activate Sistema
Sistema -> Sistema: Obtiene lista de citas para recordar
loop para cada cita
    Sistema -> Notificador: Enviar recordatorio (email/SMS)
    activate Notificador
    Notificador --> Sistema: Error de envío (ej. API caída, credenciales inválidas)
    deactivate Notificador
    Sistema -> Sistema: Registra el fallo del envío para esta cita
    note right: El sistema puede reintentar más tarde o\nmarcar la cita para seguimiento manual.
end
Sistema --> Scheduler: Proceso finalizado con errores registrados
deactivate Sistema
deactivate Scheduler
@enduml
```

---

### **7. CU-PG01: Registrar Pagos**

#### **Flujo Alternativo: Transacción de pago con tarjeta rechazada**

Este flujo se activa cuando se utiliza una pasarela de pago y la transacción es denegada por el banco emisor o la propia pasarela.

```plantuml
@startuml
title FA-PG01-02: Transacción de pago rechazada

actor "Asistente" as Asistente
participant "Sistema CRM" as Sistema
participant "Pasarela de Pago" as Pasarela

activate Asistente
Asistente -> Sistema: Ingresa datos de pago (monto, tarjeta de crédito) (RF-PG001)
activate Sistema
Sistema -> Pasarela: Solicitar procesamiento de pago
activate Pasarela
Pasarela --> Sistema: Transacción Rechazada (ej. fondos insuficientes, tarjeta inválida)
deactivate Pasarela
Sistema --> Asistente: Muestra mensaje de error: "Pago rechazado. Razón: [Motivo del rechazo]."
deactivate Sistema
deactivate Asistente
@enduml
```

---

### **8. CU-PG03: Generar Reportes Financieros**

#### **Flujo Alternativo: Generación de reporte sin datos**

Este escenario ocurre si el Administrador selecciona un rango de fechas o filtros para los cuales no existe ninguna transacción o dato.

```plantuml
@startuml
title FA-PG03-01: Generación de reporte sin datos

actor "Administrador" as Admin
participant "Sistema CRM" as Sistema

activate Admin
Admin -> Sistema: Solicita generar reporte con filtros específicos (RF-PG005)
activate Sistema
Sistema -> Sistema: Consulta la base de datos con los filtros
Sistema -> Sistema: No se encuentran registros
Sistema --> Admin: Muestra mensaje: "No se encontraron datos para los filtros seleccionados."
deactivate Sistema
deactivate Admin
@enduml
```

---

### **9. CU-RH01: Gestionar Personal de la Clínica**

#### **Flujo Alternativo: Intento de registrar empleado con DNI/email duplicado**

Similar al registro de clientes, es fundamental evitar la duplicidad de empleados para mantener la integridad de los datos.

```plantuml
@startuml
title FA-RH01-01: Empleado con DNI/email duplicado

actor "Administrador" as Admin
participant "Sistema CRM" as Sistema

activate Admin
Admin -> Sistema: Ingresa datos para registrar nuevo empleado (RF-RH001)
activate Sistema
Sistema -> Sistema: Valida datos y detecta DNI/email duplicado
Sistema --> Admin: Muestra mensaje de error: "Ya existe un empleado con ese DNI o email."
deactivate Sistema
deactivate Admin
@enduml
```

---

### **10. CU-RH02: Gestionar Turnos y Guardias**

#### **Flujo Alternativo: Conflicto de horario al asignar turno**

Este flujo previene que un administrador asigne a un empleado un turno que se solape con otro ya existente en su calendario.

```plantuml
@startuml
title FA-RH02-01: Conflicto de horario al asignar turno

actor "Administrador" as Admin
participant "Sistema CRM" as Sistema

activate Admin
Admin -> Sistema: Selecciona empleado y asigna nuevo turno (fecha/hora) (RF-RH005)
activate Sistema
Sistema -> Sistema: Valida la disponibilidad del empleado para el nuevo turno
note right: El sistema detecta que el nuevo turno se\nsolapa con un turno ya asignado.
Sistema --> Admin: Muestra error: "Conflicto de horario. El empleado ya tiene un turno asignado en ese rango."
deactivate Sistema
deactivate Admin
@enduml
```

---

### **11. CU-RH03: Gestionar Servicios de la Clínica**

#### **Flujo Alternativo: Intento de crear servicio con nombre duplicado**

Para evitar confusión, el nombre de cada servicio ofrecido por la clínica debería ser único.

```plantuml
@startuml
title FA-RH03-01: Intento de crear servicio con nombre duplicado

actor "Administrador" as Admin
participant "Sistema CRM" as Sistema

activate Admin
Admin -> Sistema: Ingresa datos para crear nuevo servicio (RF-RH008)
activate Sistema
Sistema -> Sistema: Valida si ya existe un servicio con el mismo nombre
Sistema --> Admin: Muestra error: "Ya existe un servicio con ese nombre. Por favor, elija otro."
deactivate Sistema
deactivate Admin
@enduml
```
