# Plataforma de gestión de CRM - PeluDog

**Integrantes:**

- Miguel Figuera C.I: 23.558.789
- Iromy Leon C.I: V-30.243.131
- Alejandra Herde C.I: V-23.711.974

**Tutora:**

- Yuly Delgado

**Fecha del último commit:** sábado, 12 de julio de 2025

## Introducción

Este proyecto documenta el desarrollo de un Sistema de Gestión de Relaciones con los Clientes (CRM) para el consultorio veterinario PeluDog. El objetivo es digitalizar y optimizar la gestión de citas, historias clínicas y otros procesos clave para mejorar la calidad del servicio y la eficiencia operativa.

La propuesta plantea desarrollar un sistema CRM integral para el consultorio veterinario PeluDog, ubicado en La Mora I (La Victoria, Aragua), con el fin de digitalizar la gestión de citas, historias clínicas, vacunaciones, comunicaciones y pagos. El diagnóstico participativo detecta ineficiencias en los procesos manuales actuales y define como prioridad automatizar la agenda y los registros médicos.

## Documentos Principales

- **[Documento de Requisitos](./DocumentoDeRequisitos/DocumentoDeRequisitos.md):** Contiene la especificación completa de los requisitos funcionales y no funcionales del sistema.
- **[Modelado de la Aplicación](./ModeladoDelDominioDeLaAplicacion/ModeladoDeDominioDeLaAplicacion.md):** Describe el modelo del dominio, los procesos de negocio, los actores y los objetos de negocio del sistema.

## Estructura del Repositorio

Este repositorio funciona como un índice para la documentación del proyecto. A continuación se describe el contenido de cada directorio principal:

### [DocumentosSeparados/](./DocumentosSeparados/)

Alberga los documentos que describen cada una de las partes del proyecto de forma individual.

### [ModeladoDelDominioDeLaAplicacion/](./ModeladoDelDominioDeLaAplicacion/)

Contiene el documento dedicado al modelado completo de la aplicación.

### [DocumentoDeRequisitos/](./DocumentoDeRequisitos/)

Contiene la entrega final y consolidada del documento de requisitos del sistema.

### [CodigoDiagramasPlantUml/](./CodigoDiagramasPlantUml/)

Guarda el código fuente en formato PlantUML utilizado para generar varios de los diagramas del proyecto.

### [Imagenes/](./Imagenes/)

Aloja todos los diagramas, gráficos e imágenes utilizados en la documentación.

### [plantillasVolere/](./plantillasVolere/)

Contiene todas las plantillas de requisitos, inspiradas en el enfoque Volere, redactadas en formato Markdown.

### [Impresion/](./Impresion/)

Contiene las versiones en PDF de los documentos principales para su fácil impresión y visualización.

- **[DocumentoDeRequisitos.pdf](./Impresion/DocumentoDeRequisitos.pdf)**
- **[ModeladoDeLaAplicacion.pdf](./Impresion/ModeladoDeLaAplicacion.pdf)**

## Índice de Diagramas

A continuación se presenta un índice de los diagramas incluidos en el proyecto, organizados por categoría.

### Diagramas de Casos de Uso

- [Diagrama de Casos de Uso General](./Imagenes/DiagramaDeCasosDeUsoGeneral.png)
- [Actualizar Historia Clínica](./Imagenes/CasosDeUso/ActualizarHistoriaClinica.png)
- [Autenticación de Usuarios](./Imagenes/CasosDeUso/AutenticacionDeUsuarios.png)
- [Confirmación de Citas](./Imagenes/CasosDeUso/confirmacionDeCitas.png)
- [Consultar Historia Clínica](./Imagenes/CasosDeUso/consultarHistoriaClinico.png)
- [Consultar Historial de Pagos](./Imagenes/CasosDeUso/consultarHistorialPagos.png)
- [Generar Reporte Financiero](./Imagenes/CasosDeUso/generarReporteFinanciero.png)
- [Gestionar Servicios de la Clínica](./Imagenes/CasosDeUso/gestionarServiciosClinica.png)
- [Gestionar Turnos](./Imagenes/CasosDeUso/GestionarTurnos.png)
- [Gestión de Citas](./Imagenes/CasosDeUso/gestionDeCitas.png)
- [Gestión de Personal](./Imagenes/CasosDeUso/GestionDePersonal.png)
- [Gestión de Pacientes y Clientes](./Imagenes/CasosDeUso/gestionPacientesClientes.png)
- [Registrar Pago](./Imagenes/CasosDeUso/registraPago.png)

### Diagramas de Secuencia

- [CU-AC01: Consultar Historial Clínico](./Imagenes/DiagramasDeSecuencia/CU-AC01.png)
- [CU-AC02: Registrar Nueva Consulta](./Imagenes/DiagramasDeSecuencia/CU-AC02.png)
- [CU-AU01: Autenticar Usuario](./Imagenes/DiagramasDeSecuencia/CU-AU01.png)
- [CU-GA01: Gestionar Citas](./Imagenes/DiagramasDeSecuencia/CU-GA01.png)
- [CU-GA02: Gestionar Recordatorios](./Imagenes/DiagramasDeSecuencia/CU-GA02.png)
- [CU-GP01: Gestionar Pacientes y Clientes](./Imagenes/DiagramasDeSecuencia/CU-GP01.png)
- [CU-PG01: Registrar Pagos](./Imagenes/DiagramasDeSecuencia/CU-PG01.png)
- [CU-PG02: Consultar Historial de Pagos](./Imagenes/DiagramasDeSecuencia/CU-PG02.png)
- [CU-PG03: Generar Reportes Financieros](./Imagenes/DiagramasDeSecuencia/CU-PG03.png)
- [CU-RH01: Gestionar Personal](./Imagenes/DiagramasDeSecuencia/CU-RH01.png)
- [CU-RH02: Gestionar Turnos](./Imagenes/DiagramasDeSecuencia/CU-RH02.png)
- [CU-RH03: Gestionar Servicios](./Imagenes/DiagramasDeSecuencia/CU-RH03.png)

### Flujos Alternativos de Casos de Uso

- [FA-AC01-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-AC01-01.png)
- [FA-AC02-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-AC02-01.png)
- [FA-AC02-02](./Imagenes/FlujosAlternativosCasosDeUso/FA-AC02-02.png)
- [FA-GA01-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-GA01-01.png)
- [FA-GA01-02](./Imagenes/FlujosAlternativosCasosDeUso/FA-GA01-02.png)
- [FA-GA02-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-GA02-01.png)
- [FA-GP01-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-GP01-01.png)
- [FA-GP01-02](./Imagenes/FlujosAlternativosCasosDeUso/FA-GP01-02.png)
- [FA-PG01-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-PG01-01.png)
- [FA-PG01-02](./Imagenes/FlujosAlternativosCasosDeUso/FA-PG01-02.png)
- [FA-PG03-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-PG03-01.png)
- [FA-RH01-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-RH01-01.png)
- [FA-RH02-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-RH02-01.png)
- [FA-RH03-01](./Imagenes/FlujosAlternativosCasosDeUso/FA-RH03-01.png)

### Otros Diagramas

- [Diagrama de Clases](./Imagenes/DiagramaDeClases.png)
- [Diagrama de Actividad](./Imagenes/DiagramaDeActividad.png)
- [Flujo de Trabajo en Github](./Imagenes/FlujoDeTrabajoEnGithub.png)
- [Modelo Jerárquico de Objetivos](./Imagenes/modeloJerarquicoDeObjetivos.png)
- [Organigrama de PeluDog](./Imagenes/organigramaPelugod.png)
- [Procesos de Negocio](./Imagenes/ProcesosDeNegocio.png)