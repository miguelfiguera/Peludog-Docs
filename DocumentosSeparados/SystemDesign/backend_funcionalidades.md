# Backend Funcionalidades - PeluDog CRM

Este documento define la estructura del backend para el sistema CRM de PeluDog, incluyendo endpoints, modelos, controladores y namespaces organizados por funcionalidades.

## Arquitectura General

**Framework:** Ruby on Rails 7+ (modo `api_only`)  
**Ruby Version:** 3.3+  
**Base de datos:** MySQL 8.0+  
**Autenticación:** JWT (stateless, sin sesiones del lado del servidor)  
**API:** RESTful JSON API  
**Orquestación:** Docker Compose para todos los servicios  
**Proxy:** NGINX como proxy inverso y servidor de archivos estáticos

---

## 1. Namespace: Auth

### Controladores
- `AuthController`

### Endpoints

#### POST /api/auth/login
**Propósito:** Autenticación de usuarios  
**Parámetros:**
```json
{
  "email": "string",
  "password": "string"
}
```
**Respuesta:**
```json
{
  "token": "jwt_token",
  "user": {
    "id": "integer",
    "email": "string",
    "role": "string",
    "name": "string"
  }
}
```

#### POST /api/auth/logout
**Propósito:** Cerrar sesión (invalidar token)

#### POST /api/auth/validate
**Propósito:** Validar token JWT activo

---

## 2. Namespace: Admin

### Controladores
- `Admin::DashboardController`
- `Admin::StaffController`
- `Admin::ShiftsController`
- `Admin::FinanceController`
- `Admin::ServicesController`
- `Admin::ContentController`

### Endpoints

#### Dashboard Principal
- **GET /api/admin/dashboard/metrics**  
  Métricas generales: ingresos, citas, nuevos clientes

#### Gestión de Personal
- **GET /api/admin/staff**  
  Listar todos los empleados
- **POST /api/admin/staff**  
  Registrar nuevo empleado
- **GET /api/admin/staff/:id**  
  Obtener datos específicos de empleado
- **PUT /api/admin/staff/:id**  
  Actualizar datos de empleado
- **DELETE /api/admin/staff/:id**  
  Desactivar empleado
- **PUT /api/admin/staff/:id/role**  
  Cambiar rol/permisos de empleado

#### Gestión de Turnos
- **GET /api/admin/shifts**  
  Obtener calendario de turnos
- **POST /api/admin/shifts**  
  Crear nuevo turno
- **PUT /api/admin/shifts/:id**  
  Modificar turno existente
- **DELETE /api/admin/shifts/:id**  
  Eliminar turno
- **GET /api/admin/shifts/availability**  
  Consultar disponibilidad general

#### Gestión Financiera
- **GET /api/admin/finance/payments**  
  Historial de pagos con filtros
- **GET /api/admin/finance/reports**  
  Generar reportes de ingresos
- **GET /api/admin/finance/export**  
  Exportar datos financieros

#### Gestión de Servicios
- **GET /api/admin/services**  
  Listar servicios de la clínica
- **POST /api/admin/services**  
  Crear nuevo servicio
- **PUT /api/admin/services/:id**  
  Actualizar servicio
- **DELETE /api/admin/services/:id**  
  Eliminar servicio
- **PUT /api/admin/services/:id/staff**  
  Asignar personal a servicio

#### Gestión de Contenido
- **GET /api/admin/content/announcements**  
  Listar anuncios y promociones
- **POST /api/admin/content/announcements**  
  Crear nuevo anuncio
- **PUT /api/admin/content/announcements/:id**  
  Editar anuncio
- **DELETE /api/admin/content/announcements/:id**  
  Eliminar anuncio

---

## 3. Namespace: Clinical

### Controladores
- `Clinical::AppointmentsController`
- `Clinical::PatientsController`
- `Clinical::ClientsController`
- `Clinical::MedicalRecordsController`
- `Clinical::ConsultationsController`
- `Clinical::PaymentsController`
- `Clinical::ReportsController`

### Endpoints

#### Gestión de Citas
- **GET /api/clinical/appointments**  
  Agenda de citas (filtros: fecha, veterinario)
- **POST /api/clinical/appointments**  
  Agendar nueva cita
- **GET /api/clinical/appointments/:id**  
  Detalles de cita específica
- **PUT /api/clinical/appointments/:id**  
  Reprogramar cita
- **DELETE /api/clinical/appointments/:id**  
  Cancelar cita
- **PUT /api/clinical/appointments/:id/confirm**  
  Confirmar asistencia
- **GET /api/clinical/appointments/availability**  
  Consultar disponibilidad horaria

#### Gestión de Pacientes y Clientes
- **GET /api/clinical/clients**  
  Buscador avanzado de clientes
- **POST /api/clinical/clients**  
  Registrar nuevo cliente
- **GET /api/clinical/clients/:id**  
  Datos específicos del cliente
- **PUT /api/clinical/clients/:id**  
  Actualizar datos del cliente
- **GET /api/clinical/clients/:id/pets**  
  Listar mascotas del cliente
- **POST /api/clinical/clients/:id/pets**  
  Registrar nueva mascota al cliente
- **GET /api/clinical/pets/:id**  
  Datos específicos de la mascota
- **PUT /api/clinical/pets/:id**  
  Actualizar datos de la mascota

#### Historial Clínico
- **GET /api/clinical/pets/:id/medical-history**  
  Historial completo de la mascota
- **GET /api/clinical/pets/:id/medical-history/summary**  
  Resumen de alergias, enfermedades crónicas
- **GET /api/clinical/pets/:id/medical-history/export**  
  Exportar historial completo (JSON, Markdown, PDF)
- **GET /api/clinical/pets/:id/vaccinations**  
  Historial de vacunación
- **GET /api/clinical/pets/:id/vaccinations/status**  
  Estado actual de vacunas (al día o no)
- **GET /api/clinical/pets/:id/attachments**  
  Archivos adjuntos (radiografías, análisis)

#### Registro de Consultas
- **POST /api/clinical/consultations**  
  Crear nueva entrada en historial
- **GET /api/clinical/consultations/:id**  
  Detalles de consulta específica
- **PUT /api/clinical/consultations/:id**  
  Editar consulta reciente
- **POST /api/clinical/consultations/:id/attachments**  
  Subir archivos adjuntos
- **GET /api/clinical/consultations/:id/prescription**  
  Generar receta imprimible
- **GET /api/clinical/consultations/:id/lab-order**  
  Generar orden de laboratorio
- **POST /api/clinical/pets/:id/vaccinations**  
  Registrar aplicación de vacuna
- **POST /api/clinical/pets/:id/dewormings**  
  Registrar desparasitación

#### Gestión de Pagos
- **POST /api/clinical/payments**  
  Registrar nuevo pago
- **GET /api/clinical/payments/:id/receipt**  
  Generar comprobante imprimible
- **GET /api/clinical/clients/:id/payments**  
  Historial de pagos del cliente
- **GET /api/clinical/pets/:id/payments**  
  Historial de pagos de la mascota

#### Reportes Profesionales
- **POST /api/clinical/reports/referral**  
  Generar reporte de referencia
- **GET /api/clinical/reports/:id/export**  
  Exportar reporte (JSON, CSV, PDF)
- **GET /api/clinical/pets/:id/summary**  
  Resumen clínico para reportes

---

## 4. Namespace: Client

### Controladores
- `Client::ProfileController`
- `Client::PetsController`
- `Client::AppointmentsController`
- `Client::PaymentsController`
- `Client::MedicalRecordsController`

### Endpoints

#### Gestión de Perfil
- **GET /api/client/profile**  
  Datos personales del cliente
- **PUT /api/client/profile**  
  Actualizar datos personales

#### Mis Mascotas
- **GET /api/client/pets**  
  Listar mascotas del cliente
- **POST /api/client/pets**  
  Registrar nueva mascota
- **GET /api/client/pets/:id**  
  Detalles de mascota específica
- **PUT /api/client/pets/:id**  
  Actualizar datos de la mascota

#### Mis Citas
- **GET /api/client/appointments**  
  Historial y próximas citas
- **POST /api/client/appointments**  
  Agendar nueva cita
- **PUT /api/client/appointments/:id**  
  Reprogramar cita propia
- **DELETE /api/client/appointments/:id**  
  Cancelar cita propia

#### Historial de Pagos
- **GET /api/client/payments**  
  Historial de pagos del cliente
- **GET /api/client/payments/:id/receipt**  
  Descargar comprobante

#### Historial Clínico (Vista Cliente)
- **GET /api/client/pets/:id/medical-summary**  
  Resumen del historial clínico
- **GET /api/client/pets/:id/medical-history/export**  
  Exportar historial completo de su mascota (JSON, Markdown, PDF)
- **GET /api/client/pets/:id/vaccinations**  
  Historial de vacunación

---

## 5. Namespace: Public

### Controladores
- `Public::AppointmentsController`
- `Public::ContentController`

### Endpoints

#### Agendamiento Público
- **GET /api/public/appointments/availability**  
  Consultar disponibilidad para citas públicas
- **POST /api/public/appointments**  
  Agendar cita sin registro previo
- **POST /api/public/clients**  
  Registro rápido de cliente nuevo

#### Contenido Público
- **GET /api/public/announcements**  
  Anuncios y promociones activos
- **GET /api/public/services**  
  Servicios disponibles de la clínica

---

## Modelos de Datos Principales

### User
```ruby
# Usuarios del sistema con roles integrados
# Roles: admin, vet_admin, veterinario, asistente, lab, cliente
enum role: { admin: 0, vet_admin: 1, veterinario: 2, asistente: 3, lab: 4, cliente: 5 }
has_one :client_profile # (si rol es 'cliente')
has_many :pets # (como dueño, si rol es 'cliente')
has_many :appointment_staff # (como personal clínico)
has_many :appointments, through: :appointment_staff # (como personal asignado)
has_many :user_schedules
has_many :shift_templates, through: :user_schedules
has_many :payment_notes # (como cliente)
```

### ClientProfile
```ruby
# Perfil específico para usuarios con rol 'cliente'
belongs_to :user
```

### Pet
```ruby
# Mascotas/pacientes
belongs_to :user # (dueño con rol 'cliente')
has_many :appointments
has_many :medical_records
has_many :vaccination_records
```

### Appointment
```ruby
# Citas médicas
belongs_to :pet
belongs_to :service
belongs_to :payment_note, optional: true
has_one :medical_record
has_many :appointment_staff
has_many :staff, through: :appointment_staff, source: :user
```

### MedicalRecord
```ruby
# Registros del historial clínico
belongs_to :appointment
belongs_to :pet
has_many :prescriptions
has_many :lab_orders
has_many_attached :files
```

### PaymentNote
```ruby
# Notas de pago (facturas)
belongs_to :user # (cliente)
has_many :appointments
has_many :payment_note_items
has_many :payments
```

### Payment
```ruby
# Pagos realizados
belongs_to :payment_note
```

### Service
```ruby
# Servicios de la clínica
has_many :appointments
has_many :payment_note_items, as: :item
```

### Product
```ruby
# Productos de la clínica
has_many :payment_note_items, as: :item
```

### VaccinationRecord
```ruby
# Registro de vacunas
belongs_to :pet
belongs_to :medical_record, optional: true
```

### ShiftTemplate
```ruby
# Plantillas de turno
has_many :user_schedules
```

### UserSchedule
```ruby
# Horarios asignados
belongs_to :user
belongs_to :shift_template
```

### AppointmentStaff
```ruby
# Personal asignado a citas
belongs_to :appointment
belongs_to :user
```

### Prescription
```ruby
# Prescripciones médicas
belongs_to :medical_record
```

### LabOrder
```ruby
# Órdenes de laboratorio
belongs_to :medical_record
```

### PaymentNoteItem
```ruby
# Items en notas de pago
belongs_to :payment_note
belongs_to :item, polymorphic: true # Service o Product
```

---

## Middlewares y Políticas

### Authentication Middleware
- Validación de JWT en rutas protegidas
- Manejo de expiración de tokens

### Authorization Policies
- **AdminPolicy**: Acceso completo a todas las funcionalidades
- **VetAdminPolicy**: Acceso completo a namespace admin + funciones clínicas
- **VeterinarianPolicy**: Acceso completo a funciones clínicas
- **AssistantPolicy**: Acceso limitado a funciones clínicas y gestión de citas
- **LabPolicy**: Acceso específico a resultados de laboratorio
- **ClientPolicy**: Acceso solo a sus propios datos

### Rate Limiting
- Límites por endpoint según criticidad
- Protección especial en endpoints de autenticación

---

## Validaciones y Reglas de Negocio

### Appointments
- No permitir citas en horarios ocupados
- Validar disponibilidad del veterinario
- Confirmar cliente y mascota válidos

### Medical Records
- Solo veterinarios pueden crear/editar
- Campos obligatorios según tipo de consulta
- Validación de archivos adjuntos

### Payments
- Asociación obligatoria con cliente
- Validación de montos positivos
- Generación automática de comprobantes

### Users/Staff
- Email único en el sistema
- Roles válidos: admin, vet_admin, veterinario, asistente, lab, cliente
- Activación/desactivación soft delete
- Control de acceso basado en roles (RBAC)

---

## Configuraciones Adicionales

### Background Jobs (SolidQueue + MySQL)

#### Sistema de Recordatorios de Citas
- **AppointmentReminderJob**: Envía recordatorios 24 horas antes de la cita
- **AppointmentConfirmationJob**: Solicita confirmación 48 horas antes
- **NoShowFollowUpJob**: Seguimiento a pacientes que no asistieron
- **Programación**: Cron job diario que programa recordatorios

#### Sistema de Seguimiento Veterinario
- **VaccinationReminderJob**: Notifica esquemas de vacunación próximos a vencer
- **DewormingReminderJob**: Recordatorios de desparasitación según calendario
- **VaccinationStatusUpdateJob**: Actualiza estado de vacunas (al día/vencidas) diariamente
- **FollowUpAppointmentJob**: Programa citas de seguimiento para casos críticos
- **Programación**: Cron job semanal para revisar esquemas

#### Sistema de Pagos y Cobranza
- **PaymentReminderJob**: Recordatorios de pagos pendientes
- **OverduePaymentJob**: Notificaciones de pagos vencidos con escalación
- **PaymentStatusUpdateJob**: Actualiza estado de notas de pago automáticamente
- **Programación**: Cron job diario para revisar vencimientos

#### Sistema de Reportes Automáticos
- **DailyReportJob**: Genera reportes diarios de actividad
- **WeeklyFinanceReportJob**: Reportes financieros semanales para administración
- **MonthlyClientReportJob**: Reportes mensuales de retención de clientes
- **VaccinationComplianceReportJob**: Reportes de cumplimiento de vacunación
- **MedicalHistoryExportJob**: Genera historiales clínicos en segundo plano (JSON, Markdown, PDF)
- **Programación**: Cron jobs programados según frecuencia

#### Configuración SolidQueue
```ruby
# config/solid_queue.yml
production:
  dispatchers:
    - polling_interval: 1
      batch_size: 500
  workers:
    - queues: critical
      threads: 3
    - queues: default,reminders
      threads: 5
    - queues: reports,cleanup
      threads: 2
```

#### Colas de Trabajo
- **critical**: Pagos, confirmaciones críticas
- **reminders**: Recordatorios de citas y vacunas
- **reports**: Generación de reportes y exportación de historiales clínicos
- **cleanup**: Tareas de mantenimiento y limpieza
- **default**: Trabajos generales

### File Storage (ActiveStorage)
- **Inicial**: Almacenamiento local en disco con ActiveStorage
- **Escalamiento**: Migración a DigitalOcean Spaces (S3 compatible) para arquitectura distribuida
- **CDN**: Integración nativa con CDN de DigitalOcean para mejor performance
- **Tipos soportados**: JPEG, PNG, WebP, PDF
- **Límites**: 50MB por archivo inicialmente
- **Organización**: Por fecha y tipo con limpieza automática de archivos temporales

### Sistema de Backups Automatizado
- **Herramienta**: Whenever gem para programar tareas cron
- **Frecuencia**: Diaria a las 2:00 AM
- **Retención**: 6 backups más recientes
- **Compresión**: Automática con gzip
- **Restauración**: Comando específico desde Rails

### Logging y Monitoreo
- **Audit trail**: Cambios críticos registrados
- **Logs de acceso**: Por roles de usuario
- **Monitoreo de errores**: Para detección temprana
- **Escalamiento**: Prometheus para observabilidad en arquitectura distribuida