### **Recopilación de Códigos PlantUML para PeluDog**

A continuación se presentan los códigos finales para cada uno de los diagramas realizados:

> **Nota:** Los diagramas específicos del módulo de autenticación se encuentran en el archivo `DiagramasAutenticacion.md`

---

### **1. Modelo de Unidades Organizacionales (Organigrama)**

Este diagrama muestra la estructura jerárquica formal y las líneas de reporte dentro del consultorio PeluDog.

```plantuml
@startuml
' --- Estilos Opcionales para mejorar la apariencia ---
skinparam titleFontSize 20
skinparam rectangle {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  BorderThickness 2
  FontColor #1A237E
  Shadowing true
  roundCorner 15
}
skinparam note {
  BackgroundColor #FFF9C4
  borderColor #F57F17
  fontColor #3E2723
}
skinparam arrow {
  Color #1565C0
  Thickness 2
}

title Modelo de Unidades Organizacionales (Organigrama)

' --- 1. Definimos todos los roles (las cajas del organigrama) ---
rectangle "<b>Veterinaria Principal</b>\nDra. Génesis Conesa" as Principal
rectangle "<b>Veterinarios</b>" as Vets
rectangle "<b>Asistente Administrativo</b>" as Admin
rectangle "<b>Veterinarios auxiliares / Técnicos Veterinarios</b>" as Auxiliares

' --- 2. Definimos las relaciones jerárquicas con flechas ---
' La Veterinaria Principal está arriba, los demás se conectan a ella.
Principal --> Vets
Principal --> Admin

' Los Auxiliares se conectan a los Veterinarios.
Vets --> Auxiliares

' --- 3. Añadimos la nota aclaratoria al rol correspondiente ---
note right of Auxiliares
  Equipo de apoyo convocado
  solo para cirugía o
  procedimientos que
  lo requieran.
end note

@enduml
```

---

### **2. Modelo Jerárquico de Objetivos**

Este diagrama de mente (Mind Map) visualiza cómo los objetivos específicos (SMART) contribuyen a las áreas estratégicas y al objetivo general del negocio.

```plantuml
@startmindmap
' --- Estilos para un look profesional y claro ---
<style>
mindmapDiagram {
  node {
    BackgroundColor #E3F2FD
    LineColor #1565C0
    FontName Arial
    RoundCorner 15
    Shadowing true
  }
  .level1 {
    BackgroundColor #1565C0
    FontColor #FFFFFF
    FontSize 14
    LineThickness 2
  }
  .level2 {
    BackgroundColor #64B5F6
    FontColor #000000
    FontSize 12
    LineThickness 1.5
  }
  .level3 {
    BackgroundColor #BBDEFB
    FontSize 11
    HorizontalAlignment left
  }
  arrow {
    LineColor #4682B4
    LineThickness 1.5
  }
}
</style>

title Modelo Jerárquico de Objetivos - PeluDog

' --- Nodo Raíz: El Objetivo General del Negocio ---
* **Proveer atención médica veterinaria de alta calidad**\npara mantener y mejorar la salud y\nel bienestar de las mascotas. <<level1>>

' --- Ramas Principales: Las Categorías Estratégicas ---
** Satisfacción del Cliente <<level2>>
*** Tasa de Retención: <b>85%</b>\n(Cierre del año)
*** Encuestas de Satisfacción: <b>>75/100</b>\n(Implementar en 1 mes, lograr en 1er trimestre)
*** Aumento de Referidos: <b>+50%</b>\n(Próximos 3 meses)

** Eficiencia Operativa <<level2>>
*** Reducción Errores de Caja: <b>-90%</b>\n(De 20 a 2 errores/semana a fin de año)
*** Reducción Tiempo por Consulta: <b>-15%</b>\n(En 4 meses)
*** Reducción Tiempo Recordatorios: <b>-70%</b>\n(En 2 meses)

** Rentabilidad y Sostenibilidad <<level2>>
*** Clientes Nuevos: <b>+20%</b>\n(Fin de año)
*** Citas Facturadas: <b>+15%</b>\n(Fin de año)
*** Cumplimiento Vacunación: <b>+50%</b>\n(Próximos 6 meses)

** Cumplimiento Normativo <<level2>>
*** Mantener <b>100%</b> de cumplimiento\ncon normativas y regulaciones.

@endmindmap
```

---

### **3. Modelo de Proceso de Negocio (Diagrama de Actividad)**

Este diagrama de actividad con carriles (swimlanes) modela el flujo de trabajo principal, mostrando las responsabilidades de cada actor en el ciclo de atención al cliente.

```plantuml
@startuml
' --- Estilos para mejorar la legibilidad y apariencia ---
<style>
title {
  FontSize 20
  FontColor #333333
}
activityDiagram {
  backgroundColor #FFFFFF
  borderColor #A0A0A0

  partition {
    backgroundColor #E3F2FD
    borderColor #1565C0
    fontColor #1A237E
  }

  activity {
    backgroundColor #FFFFFF
    borderColor #1565C0
    fontColor #000000
    roundCorner 15
  }

  arrow {
    color #1565C0
    thickness 2
  }

  note {
    backgroundColor #FFF9C4
    borderColor #F57F17
    fontColor #3E2723
  }
}
</style>

title Modelo de Proceso de Negocio: Flujo de Atención en PeluDog

' --- Definición de Carriles (Swimlanes) para cada Actor ---
|Dueño de Mascota|
start
:Necesita atención\npara su mascota;
:Contacta al consultorio\n(Teléfono, Presencial, Web);

|Asistente|
:Recibe solicitud de cita;

if (¿Es cliente nuevo?) then (Sí)
  :<b>[Gestión de Clientes]</b>\nRegistra datos del dueño\ny la mascota en CRM;
  note right
    Proceso apoyado
    por el CRM.
  end note
else (No)
  :Busca cliente/mascota\nen el sistema CRM;
endif

:<b>[Gestión de Agenda]</b>\nAgenda la cita en el calendario del CRM;
:Envía confirmación de cita;
note right
  Los recordatorios automáticos
  son una función clave del CRM
  para reducir el tiempo invertido
  (Objetivo de Eficiencia).
end note

|Dueño de Mascota|
:Asiste a la cita con\nsu mascota;

|Asistente|
:Realiza la recepción del paciente;

|Veterinario|
:<b>[Atención Clínica]</b>\nConsulta historial clínico en CRM;
:Realiza anamnesis y examen físico;
:Registra hallazgos, diagnóstico\ny plan de tratamiento en la\nHistoria Clínica Electrónica;
note left
  Proceso central
  apoyado por el CRM
  para cumplir el
  Objetivo General.
end note

if (¿Requiere procedimiento especial?) then (Sí)
  |Asistente|
  :<b>[Gestión de Procedimientos]</b>\nCoordina agenda para cirugía/\nhospitalización según disponibilidad;
else (No)
endif

|Asistente|
:<b>[Registro de Ingresos]</b>\nGenera cuenta de servicios\ny productos utilizados;
:Registra el pago en el sistema;
note right
  Registro básico de ingresos,
  la facturación fiscal está
  fuera del alcance del CRM.
end note

|Dueño de Mascota|
:Realiza el pago;

|Veterinario|
if (¿Requiere seguimiento o\npróxima vacunación?) then (Sí)
  |Asistente|
  :<b>[Seguimiento Post-Atención]</b>\nPrograma recordatorios automáticos\n(vacunas, desparasitación, control);
  note right
    Proceso clave para
    fomentar la retención
    (Objetivo de Satisfacción).
  end note
else (No)
endif

|Dueño de Mascota|
:Recibe recordatorios futuros;
stop

@enduml
```

---

### **4. Modelo de Objetos de Negocio (Diagrama de Clases Completo)**

Este es el diagrama de clases final y completo, que representa todos los objetos de negocio, incluyendo usuarios, servicios, turnos y los registros detallados del historial clínico, así como sus relaciones.

```plantuml
@startuml
' --- Estilos para un diagrama limpio y profesional ---
<style>
title {
  FontSize 20
  FontColor #333333
}
classDiagram {
  backgroundColor #FFFFFF
  borderColor #A0A0A0

  class {
    backgroundColor #E3F2FD
    borderColor #1565C0
    fontColor #1A237E
    roundCorner 15
    shadowing true
  }

  arrow {
    color #1565C0
    thickness 2
  }

  note {
    backgroundColor #FFF9C4
    borderColor #F57F17
    fontColor #3E2723
  }
}
</style>

title Modelo de Objetos de Negocio (Completo) - Dominio de PeluDog

' --- Definición de Clases Principales y de Usuario ---

abstract class Usuario {
  +NombreUsuario
  +Rol
}
class Veterinario extends Usuario {
  +ID_Colegio
}
class Asistente extends Usuario {
  +Cargo
}
class Turno {
  +FechaInicio
  +FechaFin
}
class Cliente {
  +Nombre
  +Apellido
  +ID_Identificacion
}
class Mascota {
  +Nombre
  +Especie
  +Raza
}
class Cita {
  +FechaHora
  +Motivo
  +Estado
}
class Servicio {
  +NombreServicio
  +PrecioBase
}
class ComprobanteVenta {
  +Numero
  +Fecha
  +Total
}

' --- Clases del Historial Clínico ---

class HistoriaClinica {
  .. Contenedor Lógico ..
}
class EntradaConsulta {
  +Fecha
  +Anamnesis
  +Diagnostico
  +PlanTratamiento
}
class Prescripcion {
  +Medicamento
  +Dosis
  +Frecuencia
}
class SolicitudExamen {
  +TipoExamen
  +MuestraRequerida
}
class ArchivoAdjunto {
  +NombreArchivo
  +Tipo
  +URL
}
class RegistroVacunacion {
  +ProductoUtilizado
  +Lote
  +FechaProximaDosis
}


' --- Definición de Relaciones ---

' --- Relaciones de Usuarios y Turnos ---
Usuario "1..*" -- "0..*" Turno : tiene asignado >
Veterinario --|> Usuario
Asistente --|> Usuario

' --- Relaciones de Clientes y Mascotas ---
Cliente "1" -- "1..*" Mascota : tiene >
Mascota "1" -- "1" HistoriaClinica : tiene >

' --- Relaciones de Citas y Servicios ---
Cita "1" -- "1" Cliente : es de >
Cita "1" -- "1" Mascota : es para >
Cita "1" -- "1" Veterinario : es con >
Cita "1" -- "1" Servicio : es para un >
Cita "1" ..> "1" EntradaConsulta : genera (si es consulta)

' --- Relaciones del Historial Clínico ---
HistoriaClinica "1" -- "1..*" EntradaConsulta : contiene >
EntradaConsulta "1" -- "1" Veterinario : realizada por >
' Una consulta puede generar CERO o MUCHOS de los siguientes registros
EntradaConsulta "1" -- "0..*" Prescripcion : genera >
EntradaConsulta "1" -- "0..*" SolicitudExamen : genera >
EntradaConsulta "1" -- "0..*" ArchivoAdjunto : puede tener >
EntradaConsulta "1" -- "0..*" RegistroVacunacion : puede incluir >

' --- Relaciones de Facturación ---
ComprobanteVenta "1" -- "1" Cliente : se emite a >
ComprobanteVenta "1..*" -- "1..*" Servicio : detalla >
EntradaConsulta "1" ..> "0..1" ComprobanteVenta : puede generar >


note "Facturación fiscal y gestión de inventario\ndetallada están fuera del alcance del CRM." as N1
note "La clase 'Usuario' permite gestionar\npermisos basados en el Rol\n(Veterinario, Asistente)." as N2

@enduml
```
