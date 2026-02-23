# PLATAFORMA DE GESTIÓN CRM PARA EL CONSULTORIO VETERINARIO PELUDOG

**Autores:**
- Miguel Figuera (CI: 23.558.789)
- Iromy Leon (CI: V-30.243.131)
- Alejandra Herde (CI: V-23.711.974)

**Tutora:** Yuly Delgado

**Venezuela, Julio 2025**

---

# DEDICATORIA

---

**Alejandra Herde**

A mi familia, por su apoyo incondicional, sus consejos y por creer en mí en cada momento. Su esfuerzo y confianza han sido mi mayor motivación.

---

**Miguel Figuera**

A Chepino, el gato pepino, todo este esfuerzo es un homenaje a tu amistad y cariño. Espérame en el Bifrost junto a Lupita; allá nos vemos muchacho.

A Pepperina, Calico, Randori y Ruperto; la vida no es vida si no hay un gato presente.

---

**Iromy Leon**

*(Pendiente — falta su dedicatoria)*

---

<div style="page-break-before: always;"></div>

# AGRADECIMIENTO

---

**Miguel Figuera**

A mi madre Marisol Quintero, porque si no es por ella nunca habría pensado en regresar a la universidad para una segunda carrera.

A Rafael Hernández, por la fe, las clases gratis y el esfuerzo invertido en darle forma a mi conocimiento.

A Diego Solórzano, el mejor product owner que existe en el mundo.

A Simple-C, por corregir todas mis malas prácticas y ser fuente infinita de conocimiento.

A Génesis Conesa, por las ideas y el apoyo en esta loca aventura de hacer un producto que ya existe pero ponerlo al alcance de todos.

A Matthew Le Tiec, el mejor profesor de Ruby y Rails que existe.

Al Slack de Ruby on Rails, su ayuda colectiva convirtió este proyecto en algo posible.

---

**Alejandra Herde**

A mis compañeros, quienes formaron parte de este proceso de aprendizaje y crecimiento.

---

**Iromy Leon**

*(Pendiente — faltan su dedicatoria y agradecimientos)*

---

---

# RESUMEN

La presente investigación tiene como objetivo el desarrollo e implementación de una plataforma de Gestión de Relaciones con los Clientes (CRM) para el consultorio veterinario PeluDog, ubicado en la urbanización La Mora I, La Victoria, estado Aragua. Mediante un enfoque de Investigación Acción Participativa (IAP), se realizó un diagnóstico participativo con la Dra. Génesis Conesa y su equipo de trabajo, el cual evidenció deficiencias críticas en los procesos manuales de gestión de citas, registro de historias clínicas, control de vacunación, comunicación con los clientes y registro de pagos. Estas carencias, lejos de ser exclusivas de PeluDog, representan un problema común entre consultorios veterinarios en Venezuela, donde la mayoría carece de herramientas digitales accesibles y adaptadas a su realidad. Como respuesta, se diseñó un sistema CRM integral utilizando Ruby on Rails como framework de backend, React como librería de interfaz de usuario e Inertia.js como puente entre ambos, permitiendo una experiencia de aplicación moderna sin la complejidad de una API separada. El proceso contempló el levantamiento de 41 requisitos funcionales y 22 no funcionales, la elaboración de casos de uso detallados por módulo, el modelado del dominio de la aplicación y el diseño de la arquitectura técnica. El sistema abarca la gestión de clientes y mascotas, agenda de citas, historias clínicas electrónicas, recordatorios automáticos, gestión de personal y registro básico de ingresos. Los resultados del proceso de investigación y diseño demuestran la viabilidad técnica, operativa y económica del proyecto. La plataforma se publicará como código abierto con el propósito de que otros consultorios veterinarios puedan adoptarla y adaptarla libremente, contribuyendo así al fortalecimiento de la atención veterinaria comunitaria y a la democratización del acceso tecnológico en el sector.

**Palabras clave:** CRM, consultorio veterinario, gestión de clientes, historias clínicas electrónicas, código abierto, Investigación Acción Participativa, Ruby on Rails, React, Inertia.js.

---

# INTRODUCCIÓN

En Venezuela, los consultorios veterinarios enfrentan un desafío operativo que rara vez se aborda desde el ámbito tecnológico: la gestión integral de sus procesos administrativos y clínicos se realiza, en la gran mayoría de los casos, de forma completamente manual. Agendas escritas en cuadernos, historias clínicas en carpetas de papel, esquemas de vacunación que dependen de la memoria del profesional y registros de pagos dispersos en hojas sueltas constituyen la realidad cotidiana de estos establecimientos. Esta situación no solo compromete la eficiencia del servicio, sino que pone en riesgo la continuidad del cuidado animal y limita la capacidad de crecimiento de negocios que cumplen un rol fundamental en el bienestar de las comunidades.

El consultorio veterinario PeluDog, ubicado en la urbanización La Mora I, La Victoria, estado Aragua, no es ajeno a esta realidad. Bajo la dirección de la Dra. Génesis Conesa, PeluDog ofrece servicios de consulta veterinaria, hospedaje y peluquería canina a precios accesibles, atendiendo a una población diversa de propietarios de mascotas provenientes de distintos grupos socioeconómicos de la ciudad. Más allá de su labor como negocio, el consultorio participa activamente en campañas de vacunación y jornadas de atención comunitaria, lo que lo posiciona como un pilar en la salud animal de su entorno. Sin embargo, el crecimiento sostenido de su cartera de pacientes ha hecho que los procesos manuales se conviertan en un obstáculo cada vez más difícil de sortear.

Ante esta problemática, la presente investigación se enmarca en la modalidad de Proyecto Socio Tecnológico bajo el enfoque de Investigación Acción Participativa (IAP), con el propósito de desarrollar una plataforma de Gestión de Relaciones con los Clientes (CRM) que digitalice y optimice los procesos medulares del consultorio. A través de un diagnóstico participativo realizado junto a la Dra. Génesis y su equipo, se identificaron deficiencias críticas en cinco áreas fundamentales: la gestión de citas, el registro de historias clínicas, el control de esquemas de vacunación, la comunicación con los clientes y el registro de pagos. Cada una de estas carencias genera consecuencias tangibles que van desde la pérdida de información clínica hasta la imposibilidad de realizar un seguimiento proactivo de la salud de los pacientes.

La solución propuesta consiste en un sistema CRM integral desarrollado con Ruby on Rails como framework de backend, React como librería de interfaz de usuario e Inertia.js como capa de integración entre ambos, conformando un stack tecnológico moderno que permite una experiencia de usuario fluida sin la complejidad de mantener una API separada. El sistema abarca la gestión de clientes y mascotas, la agenda de citas, las historias clínicas electrónicas, los recordatorios automáticos, la gestión de personal y el registro básico de ingresos. El proceso de desarrollo contempló el levantamiento de 41 requisitos funcionales y 22 no funcionales, la elaboración de casos de uso detallados, el modelado del dominio y el diseño completo de la arquitectura técnica.

Un aspecto distintivo de este proyecto es su vocación de servicio comunitario. Durante el proceso de investigación se evidenció que las dificultades de PeluDog no son un caso aislado, sino un problema estructural que afecta a la mayoría de los consultorios veterinarios del país. Por esta razón, la plataforma se concibe como un proyecto de código abierto, diseñado para que otros consultorios puedan adoptarlo y adaptarlo a sus necesidades particulares sin costo alguno. Esta decisión refleja el compromiso del equipo con la democratización del acceso tecnológico y el fortalecimiento de la atención veterinaria en Venezuela.

El presente documento se estructura en tres fases según el esquema de Proyecto Socio Tecnológico de la UNETI. La Fase I desarrolla el diagnóstico situacional, incluyendo la descripción del contexto, la identificación y análisis de problemas, los objetivos, las bases legales y la vinculación con las líneas estratégicas de desarrollo nacional. La Fase II presenta la planificación del proyecto, abarcando el estudio de factibilidad, el cronograma de actividades, el plan de acción y la metodología de desarrollo. La Fase III expone los resultados y logros obtenidos, las conclusiones y las recomendaciones derivadas de la investigación.

---

# FASE I — DIAGNÓSTICO SITUACIONAL

---

## DIAGNÓSTICO PARTICIPATIVO — SITUACIONAL

### Descripción del contexto

El consultorio veterinario PeluDog, dirigido por la Dra. Génesis Conesa, se encuentra ubicado en la urbanización La Mora I, La Victoria, municipio José Félix Ribas, estado Aragua. Ofrece servicios integrales de salud animal que incluyen consultas veterinarias, hospedaje y peluquería canina a precios accesibles, atendiendo a una población diversa de propietarios de mascotas provenientes de distintos estratos socioeconómicos de la ciudad. Más allá de su operación como negocio, el consultorio participa activamente en campañas de vacunación y jornadas de atención comunitaria, lo que lo posiciona como un referente en la salud animal de su entorno.

La Dra. Génesis Conesa y su asistente constituyen el equipo base del consultorio, apoyadas por veterinarios auxiliares en procedimientos especiales. El espacio físico del consultorio, integrado en una casa familiar, está adaptado para la prestación de servicios veterinarios y cuenta con las condiciones necesarias para consultas, hospedaje y peluquería canina.

El diagnóstico participativo se realizó mediante entrevistas a profundidad con la Dra. Génesis y su asistente, observación directa de los procesos operativos del consultorio, y talleres de trabajo colaborativo, siguiendo el enfoque de Investigación Acción Participativa (IAP).

### Naturaleza de la Comunidad

La urbanización La Mora I es un sector residencial consolidado de la ciudad de La Victoria, compuesto por edificios multifamiliares y viviendas unifamiliares. Alberga familias de diversos estratos socioeconómicos, muchas de las cuales son propietarias de mascotas. El sector cuenta con servicios públicos completos, buena conectividad vial y una dinámica comunitaria activa.

---

# LOCALIZACIÓN GEOGRÁFICA

## Mapa y Ubicación

Las coordenadas del consultorio veterinario PeluDog según Google Maps son: latitud 10.2294169, longitud -67.3026786. Se encuentra ubicado en el urbanismo La Mora I, tercera entrada, avenida 11 con avenida 36, La Victoria, estado Aragua.

## Ubicación Geográfica

El consultorio veterinario PeluDog se encuentra en el municipio José Félix Ribas, específicamente en la urbanización La Mora I, una zona residencial consolidada de la ciudad de La Victoria, estado Aragua, Venezuela. El establecimiento opera en la intersección de la avenida 11 con la avenida 36, frente a los edificios residenciales del sector.

## Comunidad o Institución

Consultorio Veterinario PeluDog, dirigido por la Dra. Génesis Conesa. El consultorio funciona integrado en una casa familiar adaptada para la prestación de servicios veterinarios, incluyendo áreas de consulta, hospedaje y peluquería canina.

## Puntos y Referencias

El consultorio se ubica frente a los edificios residenciales de La Mora I, en la intersección de la avenida 11 con la avenida 36, accediendo por la tercera entrada del urbanismo.

## Vías de Acceso

El acceso principal al consultorio se realiza a través de la tercera entrada del urbanismo La Mora I. Esta entrada conecta directamente con las avenidas principales del sector, lo que permite una llegada cómoda tanto en vehículo particular como en transporte público. La urbanización se encuentra bien comunicada con el centro de La Victoria y con las vías que conectan hacia El Consejo y el terminal de pasajeros.

## Servicios públicos existentes

La zona cuenta con todos los servicios públicos básicos en funcionamiento adecuado: suministro de agua potable por tubería, servicio eléctrico, conectividad a internet, y servicio de aseo urbano con recolección regular. Esta disponibilidad de servicios resulta favorable para la operación del consultorio y para la implementación de herramientas tecnológicas como el CRM propuesto, ya que garantiza la conectividad necesaria para el uso continuo de la plataforma.

## Medios de transporte de la comunidad o institución

La urbanización La Mora I cuenta con amplia cobertura de transporte público. Las rutas de autobús que cubren el recorrido "La Mora – Centro" pasan directamente por el sector, facilitando el acceso desde el casco central de La Victoria. Adicionalmente, las rutas que conectan La Victoria con El Consejo y con el terminal de pasajeros transitan por las cercanías del urbanismo. El sector también es atendido por servicios de mototaxi y carritos por puesto, lo que ofrece múltiples alternativas de movilidad para los clientes del consultorio. Esta accesibilidad en el transporte público contribuye a que PeluDog pueda atender a una clientela diversa proveniente de distintas zonas de la ciudad.

---

## HISTORIA DE VIDA DE LA COMUNIDAD

La Victoria, capital del municipio José Félix Ribas del estado Aragua, es una de las ciudades más antiguas de Venezuela. Sus orígenes se remontan a 1593, cuando existía como aldea indígena, y fue erigida formalmente como pueblo de doctrina el 18 de noviembre de 1620 bajo el nombre de Nuestra Señora de Guadalupe de La Victoria. En 1795, el Rey Carlos IV le otorgó el título de Villa, y en 1814 fue elevada a la categoría de ciudad por decreto del Libertador Simón Bolívar, tras la célebre Batalla de La Victoria del 12 de febrero de ese año, en la que el General José Félix Ribas, al mando de jóvenes seminaristas y universitarios, derrotó a las fuerzas realistas. Este hecho histórico convirtió a La Victoria en símbolo de la juventud venezolana, y cada 12 de febrero se celebra allí el Día de la Juventud.

Históricamente, la región fue reconocida por sus haciendas productoras de caña de azúcar, maíz y frutas en los fértiles valles de Aragua. A partir de la década de 1960, La Victoria experimentó un acelerado proceso de industrialización que transformó su paisaje económico y generó un significativo crecimiento demográfico. Según datos del Instituto Nacional de Estadísticas (INE), para 2020 la ciudad contaba con aproximadamente 172.981 habitantes, cifra que alcanza los 281.285 en su área metropolitana, conformando una conurbación con las localidades de Las Tejerías y El Consejo.

La urbanización La Mora I, donde se ubica el consultorio veterinario PeluDog, forma parte del crecimiento residencial que acompañó la expansión industrial de la ciudad. Se trata de un sector residencial consolidado, compuesto por edificios multifamiliares y viviendas unifamiliares, con acceso completo a servicios públicos básicos y buena conectividad vial hacia el centro de La Victoria y las localidades circundantes. La comunidad de La Mora I alberga familias de diversos estratos socioeconómicos, muchas de las cuales son propietarias de mascotas y constituyen la clientela natural del consultorio PeluDog.

---

## ORGANIZACIONES VINCULADAS AL PROYECTO

Las principales organizaciones e instituciones vinculadas al desarrollo de este proyecto son:

- **Consultorio Veterinario PeluDog:** Institución beneficiaria directa del proyecto, liderada por la Dra. Génesis Conesa. Ofrece servicios de consulta veterinaria, hospedaje y peluquería canina, y participa activamente en campañas de vacunación y jornadas de atención comunitaria.
- **Universidad Nacional Experimental de las Telecomunicaciones e Informática (UNETI):** Institución académica que enmarca el presente Proyecto Socio Tecnológico, aportando la formación técnica y metodológica para el desarrollo del sistema CRM.
- **Colegio de Médicos Veterinarios del Estado Aragua:** Organismo gremial que regula el ejercicio profesional veterinario en la región, cuyas normativas deben ser contempladas en el diseño del sistema.
- **Comunidad de código abierto de Ruby on Rails:** Comunidad de desarrolladores que proporciona el ecosistema tecnológico sobre el cual se construye la plataforma, y que constituye un recurso fundamental para la sostenibilidad del proyecto a largo plazo.

---

## ANÁLISIS DE PROBLEMA

El problema central identificado en el consultorio veterinario PeluDog es la **ausencia de herramientas digitales para la gestión integral de los procesos administrativos y clínicos**, lo cual genera ineficiencias operativas, pérdida de información y limitaciones en la calidad del servicio.

### Causas del problema:

Las causas que originan esta situación pueden clasificarse en tres niveles:

**Causas estructurales:**
- Inexistencia de soluciones CRM accesibles y adaptadas a la realidad de los consultorios veterinarios en Venezuela.
- Alto costo de las plataformas comerciales existentes, que las hace inaccesibles para pequeños consultorios.
- Limitada cultura de digitalización en el sector veterinario del país.

**Causas directas:**
- Registro manual de citas en cuadernos, lo que genera errores, duplicaciones y dificultades para optimizar la agenda.
- Historias clínicas en papel que dificultan el acceso rápido a la información y son susceptibles a extravío o deterioro.
- Control de vacunación manual que aumenta el riesgo de omitir fechas de refuerzo.
- Comunicación reactiva con los clientes por falta de herramientas automatizadas de recordatorios.
- Registro de pagos disperso que impide comprender el flujo de caja del negocio.
- Gestión indiferenciada de servicios (consulta, peluquería, hospedaje) que tienen procesos y costos distintos.

**Causas contribuyentes:**
- Coordinación manual con laboratorios y otros profesionales, generando retrasos en la comunicación de resultados.
- Crecimiento sostenido de la cartera de pacientes que hace insostenible la gestión manual.

### Efectos del problema:

**Efectos directos:**
- Pérdida y deterioro de información clínica de los pacientes.
- Errores en la programación de citas que afectan la satisfacción de los clientes.
- Incumplimiento de esquemas de vacunación por falta de seguimiento oportuno.
- Imposibilidad de generar reportes para la toma de decisiones.
- Discrepancias frecuentes en las auditorías de caja semanales (promedio de 20 errores por semana).

**Efectos a largo plazo:**
- Limitación de la capacidad de crecimiento del negocio.
- Riesgo de pérdida de clientes por deficiencias en el servicio.
- Incapacidad de ofrecer atención veterinaria preventiva de forma sistemática.
- Desventaja competitiva frente a consultorios que adopten herramientas digitales.

---

## MATRIZ DE PRIORIZACIÓN DE PROBLEMAS

La siguiente matriz evalúa los problemas identificados según tres criterios: **Frecuencia** (con qué regularidad se presenta), **Impacto** (qué tan grave es su efecto en la operación y los clientes) y **Viabilidad de solución** (qué tan factible es resolverlo con el CRM propuesto). La escala utilizada es de 1 (bajo) a 5 (alto).

| N° | Problema identificado | Frecuencia | Impacto | Viabilidad | Total | Prioridad |
|----|----------------------|:----------:|:-------:|:----------:|:-----:|:---------:|
| 1 | Gestión manual de citas con errores y duplicaciones | 5 | 5 | 5 | 15 | **Alta** |
| 2 | Historias clínicas en papel susceptibles a extravío | 5 | 5 | 5 | 15 | **Alta** |
| 3 | Control de vacunación manual sin recordatorios | 4 | 5 | 5 | 14 | **Alta** |
| 4 | Registro de pagos manual con discrepancias en caja | 5 | 4 | 5 | 14 | **Alta** |
| 5 | Comunicación reactiva con clientes (sin recordatorios automáticos) | 4 | 4 | 5 | 13 | **Alta** |
| 6 | Gestión indiferenciada de servicios (consulta, peluquería, hospedaje) | 4 | 3 | 4 | 11 | **Media** |
| 7 | Coordinación manual con laboratorios externos | 3 | 3 | 3 | 9 | **Media** |

**Criterio de selección:** Se priorizan los problemas con puntuación total igual o superior a 13 puntos como problemas de alta prioridad, que serán abordados en la primera fase de implementación del CRM.

---

## PROPÓSITO U OBJETIVOS

### Objetivo General

Desarrollar e implementar una plataforma de Gestión de Relaciones con los Clientes (CRM) de código abierto para el consultorio veterinario PeluDog, que digitalice y optimice los procesos de gestión de citas, historias clínicas, control de vacunación, comunicación con los clientes y registro de pagos, mediante un enfoque de Investigación Acción Participativa.

### Objetivos Específicos

1. Diagnosticar la situación actual de los procesos administrativos y clínicos del consultorio veterinario PeluDog mediante técnicas participativas de recolección de información (entrevistas, observación directa y talleres de trabajo).

2. Documentar los requisitos funcionales y no funcionales del sistema CRM a partir de las necesidades identificadas en el diagnóstico participativo, utilizando plantillas Volere y casos de uso detallados.

3. Diseñar la arquitectura técnica del sistema CRM utilizando Ruby on Rails, React e Inertia.js, incluyendo el modelo de dominio, el modelo de base de datos y los mockups de interfaz de usuario.

4. Desarrollar los módulos esenciales del CRM (Citas, Propietarios/Pacientes, Mascotas, Historial Médico, Nueva Consulta, Pagos, Reportes Médicos, Personal, Turnos y configuración SMTP) siguiendo la metodología Cascada.

5. Publicar el código fuente del sistema como proyecto de código abierto para que otros consultorios veterinarios puedan adoptarlo y adaptarlo a sus necesidades.

---

## VINCULACIÓN CON LAS LÍNEAS ESTRATÉGICAS DE DESARROLLO DE LA NACIÓN E IMPACTO SOCIAL

### Teórico – Conocimiento

El presente proyecto se fundamenta en el enfoque de Investigación Acción Participativa (IAP), el cual promueve la construcción colectiva de soluciones a partir del diálogo entre los investigadores y la comunidad beneficiaria. Desde el punto de vista técnico, se apoya en los principios de la ingeniería de software, específicamente en las metodologías de levantamiento de requisitos, diseño de arquitectura de sistemas y desarrollo de aplicaciones web. El marco conceptual del CRM (Customer Relationship Management) proporciona la base teórica para el diseño de un sistema que priorice la gestión de las relaciones con los clientes como eje central de la operación del negocio.

### Técnico – Ámbito de acción

El ámbito técnico del proyecto comprende el desarrollo de una plataforma web utilizando Ruby on Rails como framework de backend, React como librería de interfaz de usuario e Inertia.js como capa de integración. Esta decisión tecnológica responde a la necesidad de construir una aplicación moderna, mantenible y de código abierto, que pueda ser replicada por otros consultorios veterinarios. El sistema se diseñó con 41 requisitos funcionales y 22 no funcionales, organizados en módulos de gestión de pacientes, agenda de citas, atención clínica, gestión de personal y registro de ingresos.

### Plan de Desarrollo Económico y Social de la Nación

El proyecto se vincula con el Plan de la Patria de las 7 Grandes Transformaciones 2025–2031 (Ley Orgánica publicada en Gaceta Oficial N° 6.907 Extraordinario del 24/05/2025), específicamente con las siguientes transformaciones:

**Sexta Transformación: Ecosocialismo, Ciencia y Tecnología.** Esta transformación promueve el desarrollo de tecnología pertinente y revolucionaria para la descolonización y el desarrollo nacional. El proyecto PeluDog contribuye a este objetivo al desarrollar tecnología propia, de código abierto, que resuelve un problema concreto del sector veterinario venezolano sin depender de soluciones importadas o de alto costo. La publicación del código fuente permite la apropiación social del conocimiento tecnológico, principio central de esta transformación.

**Primera Transformación: Económica.** El plan plantea la liberación de las fuerzas productivas y el impulso de nuevas formas de organización que pongan los medios de producción al servicio de la sociedad. El CRM PeluDog, al ser de código abierto, democratiza el acceso a herramientas de gestión empresarial para pequeños emprendimientos veterinarios, contribuyendo a la formalización y eficiencia de estas microempresas.

Asimismo, el proyecto se alinea con el Plan de la Patria 2019–2025 (Tercer Plan Socialista de Desarrollo Económico y Social, Gaceta Oficial Extraordinaria N° 6.442 del 03/04/2019), en su objetivo de promover el desarrollo de las tecnologías de información al servicio de la sociedad y el fortalecimiento de las capacidades productivas nacionales.

### Bases Legales

El marco legal que sustenta el presente proyecto se compone de las siguientes normativas:

**Constitución de la República Bolivariana de Venezuela (1999).** Gaceta Oficial Extraordinaria N° 5.453 del 24 de marzo de 2000.

- **Artículo 110:** Establece que el Estado reconoce el interés público de la ciencia, la tecnología, el conocimiento, la innovación y sus aplicaciones, así como los servicios de información, por ser instrumentos fundamentales para el desarrollo económico, social y político del país. Este artículo fundamenta la pertinencia de desarrollar soluciones tecnológicas como el CRM propuesto para mejorar la gestión de servicios veterinarios.

- **Artículo 108:** Dispone que los centros educativos deben incorporar el conocimiento y aplicación de las nuevas tecnologías y sus innovaciones, lo cual sustenta la naturaleza académica de este Proyecto Socio Tecnológico como espacio de formación y aplicación de conocimientos.

- **Artículo 98:** Reconoce la libertad de creación, incluyendo el derecho a la producción y divulgación de la obra científica y tecnológica, respaldando la decisión de publicar el sistema como software de código abierto.

**Ley Orgánica de Ciencia, Tecnología e Innovación (LOCTI).** Publicada en Gaceta Oficial N° 38.242 del 3 de agosto de 2005, con reforma en 2010 y 2014.

- **Artículo 1:** Establece como objeto de la ley dirigir la generación de ciencia, tecnología e innovación con base en el ejercicio pleno de la soberanía nacional y la aplicación de conocimientos populares y académicos para la solución de problemas concretos de la sociedad. El proyecto PeluDog constituye una aplicación directa de este principio al utilizar conocimientos de ingeniería de software para resolver un problema real del sector veterinario.

- **Artículo 2:** Declara las actividades científicas, tecnológicas y de innovación como de interés público para el ejercicio de la soberanía nacional, reforzando la importancia de desarrollar tecnología propia.

**Ley de Infogobierno.** Publicada en Gaceta Oficial N° 40.274 del 17 de octubre de 2013.

- **Artículo 1:** Establece los principios y lineamientos que rigen el uso de las tecnologías de información para mejorar la gestión pública y los servicios a las personas, promoviendo el desarrollo de tecnologías de información libres, garantizando la independencia tecnológica y la apropiación social del conocimiento. Aunque esta ley se dirige al Poder Público, sus principios de promoción del software libre y el conocimiento abierto inspiran y respaldan la filosofía de código abierto del proyecto PeluDog.

- **Artículo 34:** Establece que el desarrollo, adquisición, implementación y uso de las tecnologías de información por el Poder Público tiene como base el conocimiento libre, principio que el proyecto adopta de forma voluntaria al publicar su código fuente bajo una licencia abierta.

**Ley de Ejercicio de la Medicina Veterinaria.** Publicada en Gaceta Oficial el 19 de septiembre de 1968.

- **Artículo 1:** Establece que el ejercicio de la medicina veterinaria se rige por esta ley, su reglamento, los reglamentos internos y las normas de ética profesional del Colegio de Médicos Veterinarios. El CRM incorpora en su diseño los procesos necesarios para que los profesionales veterinarios cumplan con las normativas de registro clínico y seguimiento de pacientes.

- **Artículo 2:** Señala que el ejercicio de la profesión impone dedicación al estudio de las disciplinas que impliquen su desarrollo científico, fundamentando la necesidad de herramientas que faciliten el registro y análisis de datos clínicos.

**Código Deontológico de la Medicina Veterinaria.** Emitido por la Federación de Colegios de Médicos Veterinarios de Venezuela (FCMVV).

- Establece las normas éticas para el ejercicio profesional, incluyendo las obligaciones en cuanto al manejo de historiales clínicos, la confidencialidad de la información de los pacientes y los estándares de atención. El módulo de atención clínica del CRM fue diseñado para facilitar el cumplimiento de estas obligaciones deontológicas.

---

## POBLACIÓN BENEFICIADA

### Población beneficiada directa

La población beneficiada directa del proyecto está conformada por:

- **Equipo del consultorio PeluDog:** La Dra. Génesis Conesa, su asistente y los veterinarios auxiliares que participan en procedimientos especiales, quienes utilizarán el sistema diariamente para la gestión de citas, historias clínicas, recordatorios y registros de pagos.
- **Propietarios de mascotas atendidos por PeluDog:** Los clientes actuales y futuros del consultorio, quienes se beneficiarán de una atención más organizada, recordatorios automáticos de vacunación y citas, y un seguimiento más efectivo de la salud de sus mascotas. Se estima que el consultorio atiende a varios cientos de familias propietarias de mascotas provenientes de distintas zonas de La Victoria y sus alrededores.
- **Las mascotas:** Como pacientes del consultorio, se benefician directamente de un mejor seguimiento clínico, cumplimiento de esquemas de vacunación y detección temprana de enfermedades gracias a la digitalización de sus historias clínicas.

### Población beneficiada indirecta

La población beneficiada indirecta comprende:

- **Otros consultorios veterinarios de La Victoria y Venezuela:** Al ser publicado como software de código abierto, cualquier consultorio veterinario pequeño o mediano podrá adoptar la plataforma para digitalizar sus propios procesos, abandonando la gestión en papel y mejorando la fiabilidad de sus historias clínicas a lo largo del tiempo. Según datos del INE, el municipio José Félix Ribas cuenta con más de 214.000 habitantes, lo que implica una demanda significativa de servicios veterinarios atendida por múltiples consultorios que enfrentan problemáticas similares.
- **La comunidad de La Victoria:** Al fortalecer la capacidad operativa de los consultorios veterinarios locales, se contribuye a mejorar la salud animal de la comunidad, fomentando la medicina veterinaria preventiva y reduciendo la incidencia de enfermedades por falta de seguimiento.
- **La comunidad de desarrolladores de software libre:** El proyecto aporta al ecosistema de código abierto una solución especializada para un nicho poco atendido, generando conocimiento técnico documentado y reutilizable.
- **La comunidad académica de la UNETI:** El proyecto sirve como referencia metodológica y técnica para futuros Proyectos Socio Tecnológicos que busquen aplicar tecnologías de la información a la resolución de problemas comunitarios concretos.

---

## RAZONES DEL PROYECTO

La razón fundamental del proyecto radica en la necesidad de resolver un problema concreto y verificado: el consultorio veterinario PeluDog opera la totalidad de sus procesos administrativos y clínicos de forma manual, lo que genera pérdida de información, errores recurrentes, incumplimiento de esquemas de vacunación e imposibilidad de crecimiento organizado. Esta situación fue documentada a través del diagnóstico participativo, que evidenció un promedio de 20 discrepancias semanales en las auditorías de caja como indicador objetivo de la magnitud del problema.

Sin embargo, la investigación reveló que esta problemática no es exclusiva de PeluDog. La gran mayoría de los consultorios veterinarios en Venezuela carecen de herramientas digitales accesibles y adaptadas a su realidad, debido al alto costo de las plataformas comerciales existentes y a la inexistencia de soluciones locales de código abierto. Esta constatación amplió el alcance del proyecto más allá de una solución puntual: la plataforma se concibe como un bien público tecnológico que cualquier consultorio veterinario del país puede adoptar sin costo.

Adicionalmente, el proyecto responde a la necesidad académica de vincular la formación universitaria con la solución de problemas reales de la comunidad, principio central de los Proyectos Socio Tecnológicos de la UNETI y del enfoque de Investigación Acción Participativa. El desarrollo de un sistema de esta complejidad permite a los integrantes del equipo aplicar y consolidar competencias en ingeniería de software, gestión de proyectos y trabajo con comunidades beneficiarias.

---

# FASE II — PLANIFICACIÓN DEL PROYECTO

---

## FACTIBILIDAD DEL PROYECTO

### Factibilidad Técnica

El proyecto es técnicamente viable dado que se apoya en tecnologías maduras, estables y de código abierto ampliamente respaldadas por sus comunidades de desarrollo. El stack seleccionado está compuesto por Ruby on Rails 7+ como framework de backend, React como librería de interfaz de usuario e Inertia.js como capa de integración entre ambos, con MySQL 8.0+ como sistema de gestión de base de datos. Estas tecnologías cuentan con documentación extensa, ecosistemas de librerías robustos y comunidades activas que garantizan soporte a largo plazo.

La arquitectura diseñada sigue el patrón de tres capas (presentación, lógica de negocio y persistencia de datos), lo que facilita el mantenimiento, la escalabilidad y la incorporación de nuevas funcionalidades. El sistema se diseñó para operar inicialmente en un servidor básico (2 cores, 2 GB RAM, 50 GB SSD), con capacidad de escalamiento vertical hasta soportar entre 250 y 300 usuarios concurrentes sin necesidad de modificar la arquitectura base.

El equipo de desarrollo cuenta con tres integrantes con formación en ingeniería informática y conocimientos en las tecnologías seleccionadas, lo que asegura la capacidad técnica para llevar a cabo el proyecto.

### Factibilidad Operativa

El consultorio PeluDog cuenta con las condiciones operativas necesarias para adoptar el sistema CRM. La Dra. Génesis Conesa y su asistente han manifestado su disposición a participar activamente en todas las fases del proyecto, desde el levantamiento de requerimientos hasta la capacitación y puesta en marcha.

El equipo del consultorio posee conocimientos suficientes para operar herramientas digitales básicas y cuenta con experiencia en el uso de dispositivos como laptops, tablets y teléfonos inteligentes, lo que facilita la transición del registro manual al sistema digital. Adicionalmente, la infraestructura física del consultorio, aunque integrada en una casa familiar, es adecuada para la prestación de servicios veterinarios y para la operación del sistema.

### Factibilidad Económica

El proyecto presenta una viabilidad económica favorable. El consultorio está dispuesto a invertir aproximadamente $25 dólares mensuales en el alojamiento de la aplicación en un servidor, lo que constituye el único costo recurrente del sistema. Este monto es accesible para la operación del negocio y representa una fracción mínima de los ingresos mensuales del consultorio.

Al tratarse de un proyecto desarrollado con tecnologías de código abierto y en el marco de un Proyecto Socio Tecnológico universitario, no se generan costos de licenciamiento de software ni honorarios de desarrollo. Los beneficios económicos esperados incluyen la reducción de errores en el registro de pagos (que actualmente genera un promedio de 20 discrepancias semanales en las auditorías de caja), la optimización del tiempo de consulta, y el incremento en la retención de clientes gracias al seguimiento automatizado.

### Recursos Humanos

| Rol | Persona | Función en el proyecto |
|-----|---------|----------------------|
| Desarrollador / Analista de Requisitos | Miguel Figuera (CI: 23.558.789) | Levantamiento de requerimientos, diseño de arquitectura, desarrollo backend y frontend |
| Desarrolladora | Iromy Leon (CI: V-30.243.131) | Desarrollo de módulos, documentación técnica, pruebas |
| Desarrolladora | Alejandra Herde (CI: V-23.711.974) | Desarrollo de módulos, diseño de interfaz, pruebas |
| Tutora Académica | Yuly Delgado | Orientación metodológica y supervisión del Proyecto Socio Tecnológico |
| Veterinaria Principal / Beneficiaria | Dra. Génesis Conesa | Participación en diagnóstico, validación de requisitos, pruebas de usuario |
| Asistente del consultorio | — | Participación en entrevistas, validación de flujos operativos, pruebas de usuario |

### Recursos Tecnológicos

**Infraestructura del consultorio:**
- 1 laptop con acceso a internet
- 1 tablet
- Teléfonos inteligentes
- 1 mini UPS (garantiza acceso al sistema durante cortes eléctricos)
- Conexión a internet estable

**Stack tecnológico del sistema:**
- Ruby on Rails 7+ (backend)
- React (frontend)
- Inertia.js (integración frontend-backend)
- MySQL 8.0+ (base de datos)
- Tailwind CSS (estilos)
- JWT (autenticación)
- ActiveStorage (gestión de archivos multimedia)

**Herramientas de desarrollo:**
- GitHub (control de versiones y gestión de proyecto)
- Visual Studio Code (editor de código)
- Tablero Kanban en Trello (seguimiento de tareas; se seleccionó esta plataforma en lugar de GitHub Issues porque la Dra. Génesis, como parte activa del proceso participativo, necesitaba visualizar y validar el avance del proyecto en una herramienta amigable y accesible sin conocimientos técnicos)

---

## CRONOGRAMA DE ACTIVIDADES (DIAGRAMA DE GANTT)

El siguiente cronograma refleja la ejecución real del proyecto en el período comprendido entre febrero de 2025 y enero de 2026, organizado por meses.

| Actividad | Feb | Mar | Abr | May | Jun | Jul | Ago | Sep | Oct | Nov | Dic | Ene 2026 |
|-----------|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:--------:|
| **Fase 1: Preparación y diagnóstico participativo** | ██ | | | | | | | | | | | |
| **Fase 2: Recopilación de información (entrevistas, observación)** | ██ | ██ | | | | | | | | | | |
| **Fase 3: Análisis y documentación de requisitos (41 RF, 22 RNF)** | | ██ | ██ | | | | | | | | | |
| **Fase 4: Validación y aprobación de requisitos** | | | ██ | | | | | | | | | |
| **Fase 5a: Diseño de arquitectura y modelado de datos** | | | ██ | ██ | | | | | | | | |
| **Fase 5b: Diseño de mockups e interfaz de usuario** | | | | ██ | ██ | ██ | | | | | | |
| **Fase 5c: Desarrollo backend (Rails + MySQL + JWT)** | | | | | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ |
| **Fase 5d: Desarrollo frontend (React + Inertia.js + Tailwind)** | | | | | | ██ | ██ | ██ | ██ | ██ | ██ | ██ |
| **Fase 5e: Pruebas continuas e integración** | | | | | | | ██ | ██ | ██ | ██ | ██ | ██ |
| **Fase 6: Capacitación, despliegue y seguimiento** | | | | | | | | | | | ██ | ██ |
| **Redacción del documento PST** | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ |

---

## MÓDULOS DESARROLLADOS

El sistema CRM PeluDog se organizó en dos portales diferenciados por rol de usuario, con un total de **27 páginas funcionales** distribuidas de la siguiente manera:

### Portal del Propietario (Owner) — 7 páginas

| N° | Módulo | Ruta | Sección |
|----|--------|------|---------|
| 1 | Inicio (Dashboard) | /dashboard | General |
| 2 | Mi Perfil | /profile | General |
| 3 | Mascotas | /my-pets | Mis Mascotas |
| 4 | Citas | /my-appointments | Mis Mascotas |
| 5 | Historial Médico | /my-medical-history | Mis Mascotas |
| 6 | Mis Apadrinamientos | /my-sponsorships | Otros |
| 7 | Historial de Pagos | /my-payments | Otros |

### Portal Administrativo / Veterinario — 20 páginas

**Sección Clínica (7 páginas):**

| N° | Módulo | Ruta |
|----|--------|------|
| 1 | Citas | /admin/appointments |
| 2 | Propietarios / Pacientes | /admin/owners |
| 3 | Mascotas | /admin/pets |
| 4 | Historial Médico | /admin/medical-history |
| 5 | Nueva Consulta | /admin/new-consultation |
| 6 | Pagos | /admin/payments |
| 7 | Reportes Médicos | /admin/medical-reports |

**Sección Gestión (8 páginas):**

| N° | Módulo | Ruta |
|----|--------|------|
| 8 | Métricas | /admin/metrics |
| 9 | Personal | /admin/staff |
| 10 | Turnos | /admin/shifts |
| 11 | Finanzas | /admin/financial |
| 12 | Productos y Servicios | /admin/products-services |
| 13 | Gestión de Contenido | /admin/content |
| 14 | Adopciones | /admin/adoptions |
| 15 | Apadrinamientos | /admin/sponsorships |

**Sección Configuración (5 páginas):**

| N° | Módulo | Ruta |
|----|--------|------|
| 16 | Configuración de Email | /admin/settings/email |
| 17 | Servidor SMTP | /admin/settings/smtp |
| 18 | Moneda y Símbolos | /admin/settings/currency |
| 19 | Configuración General | /admin/settings/general |
| 20 | Manual de Usuario | /admin/user-manual |

---

## PLAN DE ACCIÓN

| Metas | Actividades | Responsables | Recursos | Medios de verificación | Fecha | Hora | Observaciones |
|-------|-------------|--------------|----------|----------------------|-------|------|---------------|
| Establecer las bases del proceso participativo | Reunión inicial con la Dra. Génesis y su asistente; definición de alcance del CRM; preparación de guía de entrevista; creación de plantillas de documentación | Miguel Figuera | Laptop, acceso a internet, plantillas de documentación | Acta de reunión firmada, documento de alcance | Febrero 2025 | 9:00 AM | Fase 1: Preparación |
| Comprender los procesos actuales y necesidades del consultorio | Entrevistas individuales con Dra. Génesis y asistente; observación directa del flujo de trabajo; análisis de documentos existentes (historias clínicas, hojas de citas) | Miguel Figuera, Iromy Leon, Alejandra Herde, Dra. Génesis, Asistente | Laptop, grabadora, formatos de entrevista | Transcripciones de entrevistas, notas de observación, fotografías | Febrero – Marzo 2025 | 8:00 AM – 12:00 PM | Fase 2: Recopilación |
| Traducir necesidades en requisitos técnicos | Taller participativo para definir RF y RNF; creación de casos de uso; priorización de requerimientos con plantillas Volere | Miguel Figuera, Iromy Leon, Alejandra Herde, Dra. Génesis, Asistente | Laptop, proyector, pizarra, plantillas Volere | Documento de requisitos (41 RF, 22 RNF), diagramas de casos de uso | Marzo – Abril 2025 | 9:00 AM – 1:00 PM | Fase 3: Análisis |
| Asegurar que los requisitos reflejan las necesidades reales | Presentación de requisitos documentados; sesión de retroalimentación; ajustes según comentarios | Miguel Figuera, Dra. Génesis, Asistente | Laptop, proyector, documento de requisitos impreso | Documento de requisitos aprobado y firmado por la Dra. Génesis | Abril 2025 | 10:00 AM | Fase 4: Validación |
| Diseñar la arquitectura técnica del sistema | Diseño de arquitectura de 3 capas; modelado de base de datos; diagramas UML y de secuencia; definición de módulos | Miguel Figuera, Iromy Leon, Alejandra Herde | Laptops, PlantUML, GitHub | Documento de arquitectura, diagramas de secuencia, modelo de datos | Abril – Mayo 2025 | — | Fase 5a: Diseño de arquitectura |
| Diseñar las interfaces de usuario del sistema | Diseño de mockups para los 27 módulos (7 portal owner + 20 portal admin); validación visual con Dra. Génesis | Miguel Figuera, Iromy Leon, Alejandra Herde | Laptops, herramientas de diseño | Mockups aprobados de los 27 módulos, organizados por portal y sección | Mayo – Julio 2025 | — | Fase 5b: Diseño de mockups |
| Desarrollar el backend del sistema | Implementación de modelos, controladores, API, autenticación JWT, ActiveStorage, base de datos MySQL; desarrollo iterativo por módulos | Miguel Figuera, Iromy Leon, Alejandra Herde | Laptops, GitHub, servidor de desarrollo, Rails + MySQL | Código fuente en GitHub, tablero Kanban actualizado, pruebas unitarias | Junio 2025 – Enero 2026 | — | Fase 5c: Desarrollo backend |
| Desarrollar el frontend del sistema | Implementación de las 27 páginas en React + Inertia.js + Tailwind CSS; integración con backend; diseño responsive | Miguel Figuera, Iromy Leon, Alejandra Herde | Laptops, GitHub, React + Inertia.js + Tailwind | Código fuente en GitHub, demos funcionales, 27 páginas operativas | Julio 2025 – Enero 2026 | — | Fase 5d: Desarrollo frontend |
| Garantizar el correcto funcionamiento del sistema | Pruebas funcionales por módulo; pruebas de integración; pruebas de usabilidad con Dra. Génesis y asistente | Miguel Figuera, Iromy Leon, Alejandra Herde, Dra. Génesis, Asistente | Laptops, servidor de staging, formularios de prueba | Reportes de pruebas, bitácora de errores corregidos | Agosto 2025 – Enero 2026 | — | Fase 5e: Pruebas continuas |
| Preparar al equipo y poner en marcha el sistema | Capacitación práctica; despliegue en producción; migración de datos iniciales; acompañamiento en primeros días de uso | Miguel Figuera, Iromy Leon, Alejandra Herde, Dra. Génesis | Servidor de producción (~$25/mes), dominio, certificado SSL, manual de usuario | Sistema en línea, lista de asistencia, encuesta de satisfacción | Enero 2026 | — | Fase 6: Capacitación y despliegue |

---

## METODOLOGÍA CASCADA

### Descripción de la metodología

La metodología Cascada (Waterfall) es un modelo de desarrollo de software que organiza el proceso en fases secuenciales y lineales, donde cada fase debe completarse antes de iniciar la siguiente. Fue descrito formalmente por Winston W. Royce en 1970 y se caracteriza por su estructura ordenada, la documentación rigurosa de cada etapa y la claridad en los entregables esperados al finalizar cada fase.

Las fases clásicas de la metodología Cascada son:

1. **Análisis de requisitos:** Recopilación y documentación exhaustiva de las necesidades del sistema.
2. **Diseño del sistema:** Traducción de los requisitos en una arquitectura técnica y un diseño detallado.
3. **Implementación (codificación):** Desarrollo del código fuente según el diseño establecido.
4. **Pruebas (verificación):** Validación de que el sistema cumple con los requisitos especificados.
5. **Despliegue:** Puesta en producción del sistema en el entorno del usuario final.
6. **Mantenimiento:** Corrección de errores y mejoras posteriores a la implementación.

### ¿Cómo se relaciona la Metodología Cascada con el proyecto?

El desarrollo del CRM PeluDog se organizó siguiendo la estructura secuencial propia de la metodología Cascada. Las fases macro del proyecto se ejecutaron en orden estricto: primero se completó el diagnóstico y levantamiento de requisitos, luego el diseño, después la implementación, y finalmente las pruebas y el despliegue. Cada fase produjo entregables documentados que sirvieron como entrada para la fase siguiente.

A continuación se presenta la correspondencia entre las fases Cascada y las fases del proyecto:

**1. Análisis de Requisitos (Febrero – Abril 2025)**

Corresponde a las fases 1 a 4 del plan de acción. Se realizó el diagnóstico participativo con la Dra. Génesis y su equipo mediante entrevistas, observación directa y talleres participativos. Se documentaron 41 requisitos funcionales y 22 no funcionales organizados en casos de uso por módulo, utilizando plantillas Volere. Los requisitos fueron validados y aprobados por la Dra. Génesis antes de avanzar al diseño.

**Entregables:** Documento de diagnóstico, documento de requisitos, diagramas de casos de uso, plantillas Volere.

**2. Diseño del Sistema (Abril – Julio 2025)**

Corresponde a las fases 5a y 5b del plan de acción. Se diseñó la arquitectura de tres capas (Rails + React/Inertia.js + MySQL), se modeló la base de datos, se elaboraron diagramas UML y de secuencia, y se diseñaron los mockups de las 27 páginas del sistema (7 del portal del propietario y 20 del portal administrativo/veterinario). Los mockups fueron validados con la Dra. Génesis.

**Entregables:** Documento de arquitectura, modelo de datos, diagramas de secuencia, mockups de 27 páginas.

**3. Implementación (Junio 2025 – Enero 2026)**

Corresponde a las fases 5c y 5d del plan de acción. Se ejecutó en dos líneas de trabajo paralelas:

- **Backend (Junio 2025 – Enero 2026):** Implementación de modelos, controladores, autenticación JWT, ActiveStorage y base de datos MySQL en Ruby on Rails.
- **Frontend (Julio 2025 – Enero 2026):** Implementación de las 27 páginas funcionales en React con Inertia.js y Tailwind CSS, integradas con el backend.

Dentro de esta fase, el equipo organizó el trabajo en **iteraciones** utilizando un tablero Kanban en Trello. Se eligió Trello en lugar de GitHub Issues porque la Dra. Génesis, como participante activa del proceso de validación, necesitaba una plataforma visual e intuitiva donde pudiera seguir el avance del desarrollo, priorizar funcionalidades y dar retroalimentación sin requerir conocimientos técnicos. Trello cumplió el mismo objetivo de gestión de tareas que GitHub Issues, pero con una interfaz amigable que facilitó la participación de todos los involucrados. Los módulos se desarrollaron en el siguiente orden:

*Iteración 1 — Módulos base:* Autenticación, registro de Propietarios/Pacientes, registro de Mascotas.
*Iteración 2 — Módulo clínico:* Citas, Nueva Consulta, Historial Médico, Reportes Médicos.
*Iteración 3 — Módulo de gestión:* Personal, Turnos, Pagos.
*Iteración 4 — Configuración de comunicaciones:* Configuración de servidor SMTP para el envío de correos electrónicos (recordatorios de citas, notificaciones de vacunación).

**Nota sobre el alcance del PST:** El diseño completo del sistema contempla 27 páginas organizadas en dos portales (7 del propietario y 20 del administrador/veterinario). El alcance del presente Proyecto Socio Tecnológico se limita a la implementación de los módulos esenciales para la operación clínica y administrativa: Citas, Propietarios/Pacientes, Mascotas, Historial Médico, Nueva Consulta, Pagos, Reportes Médicos, Personal, Turnos, y la configuración SMTP. Los módulos adicionales (Métricas, Finanzas, Productos y Servicios, Gestión de Contenido, Adopciones, Apadrinamientos, portal del propietario y configuraciones adicionales) se encuentran actualmente en construcción pero quedan fuera del alcance académico de este PST. Esta decisión responde a que el proyecto, dado el nivel de profundidad alcanzado en su planificación y diseño, trascendió el ámbito académico y se encuentra en proceso de convertirse en un producto real de código abierto destinado a la comunidad veterinaria venezolana.

**Entregables:** Código fuente en GitHub, módulos funcionales del scope del PST (Citas, Propietarios/Pacientes, Mascotas, Historial Médico, Nueva Consulta, Pagos, Reportes Médicos, Personal, Turnos, configuración SMTP), tablero Kanban en Trello actualizado.

**4. Pruebas (Agosto 2025 – Enero 2026)**

Corresponde a la fase 5e del plan de acción. Se realizaron pruebas funcionales por módulo, pruebas de integración entre portales, y pruebas de usabilidad con la Dra. Génesis y su asistente de forma continua a medida que se completaban las iteraciones de desarrollo.

**Entregables:** Reportes de pruebas, bitácora de errores corregidos.

**5. Despliegue (Enero 2026)**

Corresponde a la fase 6 del plan de acción. Se desplegó el sistema en el servidor de producción, se realizó la migración de datos iniciales, se capacitó al equipo del consultorio y se brindó acompañamiento durante los primeros días de operación.

**Entregables:** Sistema en producción, manual de usuario, acta de capacitación.

**6. Mantenimiento (Posterior al despliegue)**

Fase de seguimiento continuo para la corrección de errores, incorporación de mejoras basadas en la retroalimentación de uso real, y publicación del código fuente como proyecto de código abierto.

### Justificación de la elección metodológica

La metodología Cascada resulta apropiada para este proyecto por las siguientes razones:

- **Requisitos bien definidos desde el inicio:** El diagnóstico participativo permitió identificar con claridad las necesidades del consultorio antes de iniciar el diseño y el desarrollo.
- **Equipo pequeño con roles definidos:** La estructura del equipo (3 desarrolladores, 1 tutora, 1 beneficiaria) se adapta naturalmente a un flujo secuencial de trabajo.
- **Documentación académica:** El contexto de Proyecto Socio Tecnológico requiere documentación rigurosa de cada fase, lo cual es una fortaleza intrínseca del modelo Cascada.
- **Alcance acotado:** El alcance del Proyecto Socio Tecnológico se delimitó a los módulos esenciales para la operación clínica y administrativa del consultorio: Citas, Propietarios/Pacientes, Mascotas, Historial Médico, Nueva Consulta, Pagos, Reportes Médicos, Personal, Turnos, y la configuración de envío de correos a través de un servidor SMTP. Los módulos restantes (Métricas, Finanzas, Productos y Servicios, Gestión de Contenido, Adopciones, Apadrinamientos, portal del propietario y configuraciones adicionales) forman parte del diseño completo del sistema pero su implementación queda fuera del alcance de este PST. Esta delimitación clara, junto con las exclusiones explícitas (facturación fiscal, inventario avanzado), reduce la incertidumbre y favorece la planificación secuencial.

Dentro de la fase de implementación, el equipo empleó iteraciones y un tablero Kanban en Trello para organizar el desarrollo de los módulos de forma progresiva. Se seleccionó Trello como herramienta de gestión porque permitía la participación directa de la Dra. Génesis en el seguimiento del avance, algo que no habría sido posible con herramientas técnicas como GitHub Issues. Esta práctica es compatible con la Cascada y representa una adaptación pragmática para gestionar la complejidad del desarrollo (27 páginas en 2 portales) sin alterar la secuencia general del proyecto. Las iteraciones no modificaron requisitos ni rediseñaron la arquitectura; simplemente dividieron el trabajo de codificación en bloques manejables dentro de la fase de implementación.

---

# FASE III — RESULTADOS Y LOGROS

---

## RESULTADOS Y LOGROS

El presente Proyecto Socio Tecnológico se propuso diagnosticar, diseñar y desarrollar una plataforma de Gestión de Relaciones con los Clientes (CRM) para el consultorio veterinario PeluDog, ubicado en la urbanización La Mora I, La Victoria, estado Aragua. A lo largo de las fases de diagnóstico, planificación e implementación, se obtuvieron resultados concretos y verificables que se describen a continuación, organizados según las fases de la metodología Cascada adoptada.

---

## RESULTADO

### Fase de Análisis de Requisitos

El diagnóstico participativo, realizado mediante entrevistas a profundidad con la Dra. Génesis Conesa y su asistente, observación directa de los procesos del consultorio y talleres de trabajo colaborativo, permitió identificar cinco áreas críticas de gestión manual: citas, historias clínicas, control de vacunación, comunicación con clientes y registro de pagos. La matriz de priorización determinó que los cinco problemas principales superaban el umbral de 13 puntos sobre 15, lo que confirmó la urgencia de una intervención tecnológica integral.

Como resultado de esta fase se produjo un documento de requisitos exhaustivo que contiene 41 requisitos funcionales y 22 requisitos no funcionales, organizados en módulos y documentados con plantillas Volere. Cada requisito funcional fue acompañado de su respectivo caso de uso detallado, incluyendo actores, precondiciones, flujo principal, flujos alternativos, poscondiciones y criterios de aceptación. Este nivel de detalle excede significativamente lo habitual en proyectos académicos de este nivel y constituye en sí mismo un aporte metodológico al proceso de desarrollo de software en contextos comunitarios.

### Fase de Diseño del Sistema

A partir de los requisitos aprobados, se diseñó una arquitectura de tres capas utilizando Ruby on Rails como framework de backend, React como librería de interfaz de usuario e Inertia.js como capa de integración. Esta combinación tecnológica permite una experiencia de aplicación moderna de página única (SPA) sin la complejidad de mantener una API REST separada.

Los entregables de esta fase incluyen:

El modelo de dominio de la aplicación, que define las entidades centrales del sistema (propietarios, mascotas, citas, consultas, historias clínicas, pagos, personal, turnos) y sus relaciones. El modelo de base de datos relacional en MySQL, normalizado y diseñado para soportar la operación clínica y administrativa del consultorio. Los diagramas de secuencia en PlantUML que documentan los flujos de interacción entre los actores y el sistema para cada caso de uso. Los mockups de las 27 páginas del sistema, organizadas en dos portales diferenciados (7 páginas para el portal del propietario y 20 para el portal administrativo/veterinario), validados con la Dra. Génesis entre mayo y julio de 2025.

### Fase de Implementación

El desarrollo del sistema se ejecutó entre junio de 2025 y enero de 2026, organizado en cuatro iteraciones dentro de la fase de implementación Cascada:

**Iteración 1 — Módulos base:** Se implementó el sistema de autenticación con JWT, el registro y gestión de propietarios/pacientes, y el registro de mascotas con sus datos clínicos básicos. Estos módulos constituyen la base de datos central sobre la cual operan todos los demás módulos del sistema.

**Iteración 2 — Módulo clínico:** Se desarrollaron los módulos de Citas (creación, modificación, cancelación y visualización en calendario), Nueva Consulta (registro de examen físico, diagnóstico, tratamiento y prescripción), Historial Médico (línea de tiempo de todas las atenciones de cada mascota) y Reportes Médicos (generación de documentos clínicos).

**Iteración 3 — Módulo de gestión:** Se implementaron los módulos de Personal (registro de veterinarios y asistentes con roles diferenciados), Turnos (programación de horarios del equipo) y Pagos (registro de transacciones vinculadas a consultas y servicios).

**Iteración 4 — Configuración de comunicaciones:** Se desarrolló la configuración del servidor SMTP para el envío de correos electrónicos, habilitando los recordatorios automáticos de citas y notificaciones de vacunación, una de las necesidades más urgentes identificadas en el diagnóstico.

El seguimiento de las iteraciones se realizó mediante un tablero Kanban en Trello, plataforma seleccionada para permitir la participación directa de la Dra. Génesis en la supervisión del avance sin requerir conocimientos técnicos.

### Fase de Pruebas

Las pruebas se ejecutaron de forma continua a medida que se completaban las iteraciones, abarcando pruebas funcionales por módulo (verificación de que cada funcionalidad cumple con los requisitos especificados), pruebas de integración entre módulos (por ejemplo, que una consulta registrada aparezca correctamente en el historial médico y genere el cargo correspondiente en pagos) y pruebas de usabilidad con la Dra. Génesis y su asistente como usuarios finales.

### Resultados cuantitativos del proyecto

A modo de síntesis, el proyecto produjo los siguientes entregables verificables:

Un documento de diagnóstico participativo con la identificación de 7 problemas priorizados mediante matriz cuantitativa. Un documento de requisitos con 41 requisitos funcionales y 22 no funcionales documentados con plantillas Volere. Casos de uso detallados para cada requisito funcional, con flujos principales y alternativos. Un documento de arquitectura técnica con modelo de dominio, modelo de base de datos y diagramas de secuencia en PlantUML. Mockups de 27 páginas organizadas en 2 portales (propietario y administrador/veterinario). Un sistema funcional con 10 módulos implementados dentro del alcance del PST: Citas, Propietarios/Pacientes, Mascotas, Historial Médico, Nueva Consulta, Pagos, Reportes Médicos, Personal, Turnos, y configuración SMTP. Despliegue exitoso en servidor de producción en enero de 2026, con capacitación del equipo del consultorio y migración de datos iniciales. Código fuente publicado en GitHub como repositorio de código abierto.

### Alcance extendido del proyecto

Un resultado no previsto inicialmente fue que el nivel de profundidad alcanzado en la planificación, el diseño y la implementación trascendió el ámbito estrictamente académico. El sistema completo contempla 27 páginas funcionales en dos portales, de las cuales los módulos fuera del alcance del PST (Métricas, Finanzas, Productos y Servicios, Gestión de Contenido, Adopciones, Apadrinamientos, portal del propietario y configuraciones adicionales) se encuentran actualmente en construcción. Esta decisión responde a la convicción del equipo de que el proyecto tiene potencial real como producto de código abierto para la comunidad veterinaria venezolana, y que limitarlo exclusivamente al alcance académico habría significado desaprovechar la oportunidad de generar un impacto social concreto y sostenible.

---

## CONCLUSIÓN

El desarrollo del CRM PeluDog demuestra que es posible construir herramientas tecnológicas de calidad profesional desde el ámbito universitario, aplicando metodologías rigurosas de ingeniería de software en un contexto de servicio comunitario. El enfoque de Investigación Acción Participativa permitió que la solución tecnológica respondiera a necesidades reales y no a supuestos teóricos, lo que se refleja en el nivel de detalle de los requisitos y en la validación continua de cada entregable por parte de la beneficiaria directa.

El diagnóstico participativo reveló que las dificultades de PeluDog no constituyen un caso aislado, sino un problema estructural del sector veterinario venezolano, donde la mayoría de los consultorios opera con procesos completamente manuales por carecer de herramientas digitales accesibles y adaptadas a su realidad. Esta constatación motivó la decisión de publicar el proyecto como código abierto, transformando lo que podría haber sido una solución puntual para un consultorio en una plataforma replicable para toda la comunidad veterinaria del país.

Desde el punto de vista técnico, la elección de Ruby on Rails, React e Inertia.js demostró ser acertada. El stack permite un desarrollo eficiente con un equipo pequeño, produce una experiencia de usuario moderna y fluida, y mantiene una base de código mantenible y extensible. La arquitectura de tres capas, diseñada desde el inicio para soportar escalamiento vertical y eventual escalamiento horizontal, garantiza que el sistema pueda crecer junto con las necesidades del consultorio y de otros consultorios que lo adopten en el futuro.

La metodología Cascada, adaptada con iteraciones internas en la fase de implementación y con el uso de Trello como herramienta de seguimiento colaborativo, resultó adecuada para el contexto del proyecto. La secuencia clara de fases facilitó la documentación académica requerida por el esquema UNETI, mientras que las iteraciones internas permitieron gestionar la complejidad del desarrollo de forma pragmática y con retroalimentación continua.

El proyecto cumplió con los objetivos planteados: se diagnosticaron los problemas del consultorio mediante un proceso participativo genuino, se diseñó una solución técnica completa y documentada, se implementaron los módulos esenciales para la operación clínica y administrativa, y se desplegó el sistema en producción en enero de 2026, completando así el ciclo completo de la metodología Cascada. Más allá de los objetivos académicos, el proyecto se encuentra en proceso de convertirse en un producto real, lo que constituye un testimonio del impacto que puede tener la vinculación universidad-comunidad cuando se ejecuta con compromiso y rigor.

---

## RECOMENDACIONES

A la comunidad de PeluDog y la Dra. Génesis Conesa:

Se recomienda continuar fortaleciendo el uso del sistema CRM en la operación diaria del consultorio, asegurando que tanto la Dra. Génesis como su asistente consoliden el dominio de todos los flujos del sistema (registro de pacientes, gestión de citas, registro de consultas y generación de reportes). Es importante mantener un período de operación en paralelo con los registros manuales durante las primeras semanas para garantizar la continuidad de la información durante la transición completa. Se sugiere que el consultorio designe a una persona responsable de la administración básica del sistema (respaldos, gestión de usuarios) y que se mantenga un canal de comunicación directo con el equipo de desarrollo para la notificación de incidencias.

A futuros equipos de desarrollo y mantenimiento:

Se recomienda completar la implementación de los módulos que se encuentran fuera del alcance del presente PST pero que ya están diseñados y documentados, priorizando el portal del propietario (que permitirá a los clientes consultar el historial de sus mascotas y gestionar citas de forma autónoma) y el módulo de Métricas (que proporcionará indicadores de gestión para la toma de decisiones). Es fundamental mantener la documentación técnica actualizada a medida que se incorporen nuevas funcionalidades, siguiendo los estándares de documentación establecidos en el repositorio. Se recomienda implementar pruebas automatizadas (unitarias y de integración) para garantizar la estabilidad del sistema a medida que crece en complejidad.

A la comunidad veterinaria venezolana:

Se invita a los consultorios veterinarios del país a evaluar la adopción de la plataforma, que estará disponible como código abierto en GitHub. El sistema fue diseñado con la replicabilidad como criterio de arquitectura, lo que permite su instalación y configuración en cualquier servidor con requisitos mínimos (2 cores, 2 GB RAM, 50 GB SSD, aproximadamente $25 mensuales de hosting). Se recomienda que los consultorios interesados formen alianzas con universidades o estudiantes de informática de su localidad para facilitar el proceso de instalación y personalización.

A la Universidad Nacional Experimental de las Telecomunicaciones e Informática (UNETI):

Se recomienda considerar este proyecto como referencia metodológica para futuros Proyectos Socio Tecnológicos que involucren desarrollo de software, particularmente en lo referente a la aplicación de la Investigación Acción Participativa en el levantamiento de requisitos y al uso de plantillas Volere para la documentación formal de requisitos funcionales y no funcionales. Se sugiere promover la creación de repositorios institucionales de código abierto donde los proyectos de grado puedan publicarse y continuarse por cohortes sucesivas, maximizando el impacto social de la producción académica.

---

## BIBLIOGRAFÍA

Constitución de la República Bolivariana de Venezuela. (2000). Gaceta Oficial Extraordinaria N° 5.453 del 24 de marzo de 2000.

Ley Orgánica de Ciencia, Tecnología e Innovación (LOCTI). (2005). Gaceta Oficial N° 38.242 del 3 de agosto de 2005. Con reformas de 2010 y 2014.

Ley de Infogobierno. (2013). Gaceta Oficial N° 40.274 del 17 de octubre de 2013.

Ley de Ejercicio de la Medicina Veterinaria. (1968). Gaceta Oficial del 19 de septiembre de 1968.

Plan de la Patria de las 7 Grandes Transformaciones 2025–2031. (2025). Ley Orgánica publicada en Gaceta Oficial N° 6.907 Extraordinario del 24 de mayo de 2025.

Plan de la Patria 2019–2025. Tercer Plan Socialista de Desarrollo Económico y Social de la Nación. (2019). Gaceta Oficial Extraordinaria N° 6.442 del 3 de abril de 2019.

Federación de Colegios de Médicos Veterinarios de Venezuela (FCMVV). (s.f.). Código Deontológico de la Medicina Veterinaria.

Fals Borda, O. (1999). Orígenes universales y retos actuales de la IAP. *Análisis Político*, (38), 73–90.

Fals Borda, O. y Rahman, M. A. (1991). *Acción y conocimiento: cómo romper el monopolio con investigación-acción participativa*. CINEP.

Pressman, R. S. (2014). *Ingeniería del software: un enfoque práctico* (8.ª ed.). McGraw-Hill Education.

Sommerville, I. (2016). *Ingeniería de software* (10.ª ed.). Pearson Educación.

Royce, W. W. (1970). Managing the development of large software systems. *Proceedings of IEEE WESCON*, 26, 1–9.

Robertson, S. y Robertson, J. (2012). *Mastering the Requirements Process: Getting Requirements Right* (3.ª ed.). Addison-Wesley Professional.

Laudon, K. C. y Laudon, J. P. (2020). *Sistemas de información gerencial: administración de la empresa digital* (16.ª ed.). Pearson Educación.

Buttle, F. y Maklan, S. (2019). *Customer Relationship Management: Concepts and Technologies* (4.ª ed.). Routledge.

Ruby on Rails Community. (2025). *Ruby on Rails Guides*. Recuperado de https://guides.rubyonrails.org/

React Team. (2025). *React Documentation*. Recuperado de https://react.dev/

Inertia.js. (2025). *Inertia.js Documentation*. Recuperado de https://inertiajs.com/

Instituto Nacional de Estadística (INE). (2020). *XIV Censo Nacional de Población y Vivienda*. República Bolivariana de Venezuela.

Wikipedia. (s.f.). La Victoria (Aragua). Recuperado de https://es.wikipedia.org/wiki/La_Victoria_(Aragua)

Alcaldía del Municipio José Félix Ribas. (s.f.). Historia del municipio. La Victoria, estado Aragua.
