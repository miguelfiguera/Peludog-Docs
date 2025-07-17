# Informe de Pantallas y Funcionalidades Estimadas

A continuación, se presenta un desglose de las posibles pantallas y funcionalidades de la aplicación, basado en los requisitos funcionales y las vistas solicitadas.

---

## 1. Landing Page y Acceso Público

Esta sección corresponde a la cara visible de la clínica para usuarios no autenticados.

**Pantallas Estimadas:** 3-4

*   **1.1. Página Principal (Home)**
    *   **Sección "Quiénes Somos":** Presentación de la clínica, su misión, visión y el equipo.
    *   **Sección "Anuncios y Promociones":** Espacio dinámico para publicar jornadas de vacunación, descuentos, etc. (contenido gestionado desde el Admin Dashboard).
    *   **Sección "Servicios":** Descripción de los servicios ofrecidos (clínica, peluquería, etc.).
    *   **Llamada a la Acción (CTA):** Botones claros para "Agendar Cita" y "Acceso Clientes". ( Se puede agendar cita siendo un cliente no registrado, el veterinario o el asistente puede hacer el registro directo en el consultorio para personas menos tecnicas).

*   **1.2. Agendamiento de Citas (Público)**
    *   **Funcionalidades:**
        *   Consultar disponibilidad horaria de la clínica o veterinarios.
        *   Permitir a usuarios nuevos o existentes agendar una nueva cita.
        *   Formulario para registrar datos básicos del cliente y la mascota si no existen.

*   **1.3. Login / Registro de Clientes**
    *   **Funcionalidades:**
        *   Formulario de inicio de sesión para clientes existentes.
        *   Formulario para registrar un nuevo cliente.

---

## 2. Admin Dashboard (Panel de Administración)

Panel para el rol de Administrador/Gerente, enfocado en la gestión del negocio, personal y configuración del sistema.

**Pantallas Estimadas:** 6-8

*   **2.1. Vista General (Dashboard Principal)**
    *   **Funcionalidades:**
        *   Visualización de métricas clave: ingresos del día/mes, citas programadas, nuevos clientes.
        *   Accesos directos a las funciones más utilizadas.
        *   Generación de reportes de ingresos.

*   **2.2. Gestión de Personal (RRHH)**
    *   **Funcionalidades:**
        *   Registrar, visualizar y actualizar datos de empleados.
        *   Asignar y modificar roles (permisos).
        *   Activar/desactivar cuentas de empleados.

*   **2.3. Gestión de Turnos y Horarios**
    *   **Funcionalidades:**
        *   Crear y modificar tipos de turnos (mañana, tarde, guardia).
        *   Asignar turnos a los empleados en un calendario.
        *   Consultar la disponibilidad general del personal.

*   **2.4. Gestión Financiera**
    *   **Funcionalidades:**
        *   Consultar historial de pagos por múltiples filtros (fecha, servicio, cliente).
        *   Generar reportes de ingresos detallados y exportables.

*   **2.5. Gestión de Servicios de la Clínica**
    *   **Funcionalidades:**
        *   Definir y modificar los servicios internos que ofrece la clínica.
        *   Asignar personal responsable a cada servicio.
        *   Consultar listado de servicios activos.

*   **2.6. Gestión de Contenido (Landing Page)**
    *   **Funcionalidades:**
        *   Crear, editar y eliminar anuncios y promociones que se mostrarán en la página principal.

---

## 3. Dashboard Clínico (Panel para Veterinario y Asistente)

El corazón operativo de la clínica, donde se gestiona el día a día de pacientes y citas.

**Pantallas Estimadas:** 6-7

*   **3.1. Agenda y Citas**
    *   **Funcionalidades:**
        *   Visualizar la agenda de citas (diaria, semanal, mensual).
        *   Filtrar agenda por veterinario.
        *   Agendar, reprogramar y cancelar citas.
        *   Confirmar asistencia de pacientes (solo deben aparecer los pacientes que hayan confirmado).

*   **3.2. Gestión de Pacientes y Clientes**
    *   **Funcionalidades:**
        *   Buscador avanzado de clientes y mascotas.
        *   Formularios para registrar nuevos clientes y mascotas.
        *   Visualizar y actualizar la información de contacto de los clientes y los datos generales de las mascotas.

*   **3.3. Historial Clínico**
    *   **Funcionalidades:**
        *   Visualización completa del historial de una mascota.
        *   Resumen destacado de alergias, enfermedades cronicas y datos vitales.
        *   Visor de archivos adjuntos (radiografías, análisis).
        *   Filtro para ver historial de vacunación y desparasitación.
	*   Widget/componente que indique si esta o no al dia con sus vacunas.

*   **3.4. Registro de Consulta (Atención Clínica)**
    *   **Funcionalidades:**
        *   Formulario para una nueva entrada en el historial: anamnesis, diagnóstico, tratamiento.
        *   Generador de récipes (prescripciones) imprimibles.
        *   Generador de órdenes de laboratorio imprimibles.
        *   Registro de aplicación de vacunas y desparasitantes.
        *   Subida de archivos adjuntos a la consulta.
        *   Opción para editar la entrada recién creada.

*   **3.5. Registro y Consulta de Pagos**
    *   **Funcionalidades:**
        *   Registrar un pago asociándolo a servicios/productos.
        *   Generar comprobante de pago digital imprimible.
        *   Consultar el historial de pagos de un cliente/mascota.
        
*   **3.6. Emision de reportes o historia en caso de referencias a otros profesionales**
    *   **Funcionalidades:**
    	*   Emision de reportes en formato .json o csv si maneja el mismo software.
    	*   Emision de reportes pdf para impresion o para enviar por correo si no maneja el mismo software.
    	*   Si se escribe un reporte, debe haber doble pantalla, una que muestre lo que se esta escribiendo, quizas con un formato predefinido y otra que muestre el historial clinico mas reciente al lateral.

---

## 4. Espacio de Usuario (Portal del Cliente)

Área privada para que los dueños de mascotas gestionen su información y la de sus animales.

**Pantallas Estimadas:** 4-5

*   **4.1. Mi Perfil**
    *   **Funcionalidades:**
        *   Visualizar y actualizar sus datos personales (teléfono, dirección, etc.).

*   **4.2. Mis Mascotas**
    *   **Funcionalidades:**
        *   Ver listado de sus mascotas registradas.
        *   Visualizar y actualizar datos de sus mascotas (foto, señas particulares).
        *   Opción para registrar una nueva mascota a su nombre.

*   **4.3. Mis Citas**
    *   **Funcionalidades:**
        *   Ver historial y próximas citas.
        *   Agendar, reprogramar y cancelar sus propias citas.

*   **4.4. Historial de Pagos y Facturas**
    *   **Funcionalidades:**
        *   Consultar su historial de pagos.
        *   Descargar comprobantes o facturas.

*   **4.5. Historial Clínico (Vista de Consulta)**
    *   **Funcionalidades:**
        *   Consultar un resumen del historial clínico de sus mascotas (diagnósticos, vacunas, tratamientos).

---

### Resumen General

- **Landing Page:** ~4 pantallas principales.
- **Admin Dashboard:** ~6-8 áreas funcionales/pantallas.
- **Dashboard Clínico:** ~6-7 áreas funcionales/pantallas.
- **Espacio de Usuario:** ~5 áreas funcionales/pantallas.

**Total Estimado:** La aplicación podría constar de aproximadamente **20 a 24 pantallas o secciones principales**, cada una con múltiples funcionalidades y componentes.
