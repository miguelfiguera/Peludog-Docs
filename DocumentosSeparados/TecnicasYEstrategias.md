# **Técnicas y Estrategias para la Validación, Seguimiento y Control de Cambios de Requisitos**

---

## 1. Validación de Requisitos

| Técnica                                                      | Propósito                                                                             | Justificación                                                                                      | Cuándo se Aplica                                                       |
| ------------------------------------------------------------ | ------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| **Revisión guiada (Walk-through)**                           | Repasar cada requisito con el _Product Owner_ (Dra. Génesis) y el equipo.             | Garantiza entendimiento común, detecta ambigüedades. Adecuado para un equipo pequeño y disponible. | Al finalizar cada iteración de levantamiento (CU-GP01, CU-GA01, etc.). |
| **Casos de prueba de aceptación (ATDD)**                     | Convertir criterios de aceptación en pruebas concretas (Given-When-Then).             | Facilita la comprobación objetiva; permite a la clínica validar si se cubren sus procesos reales.  | Se redactan junto al requisito; se ejecutan en demo de sprint.         |
| **Prototipos de UI de baja fidelidad**                       | Mostrar pantallas clave (agenda, historia clínica).                                   | La Dra. Génesis y asistentes comprenden mejor requisitos visuales; reduce retrabajo.               | Antes de desarrollar los módulos GP, GA y AC.                          |
| **Matriz de trazabilidad Requisito ↔ Caso de Uso ↔ CU Test** | Asegurar que todos los RF/RNF están cubiertos por algún artefacto de diseño y prueba. | Evita requisitos huérfanos u oro añadido; agiliza auditorías.                                      | Mantener viva en cada sprint.                                          |
| **Checklist de RNF**                                         | Lista de verificación para Usabilidad, Rendimiento, Seguridad, etc.                   | RNF suelen olvidarse; checklist rápida garantiza atención sistemática.                             | Revisión de cierre de sprint y revisión de diseño.                     |

## 2. Seguimiento de Requisitos

1. **Herramienta de gestión**: Se usará un tablero Kanban en GitHub Issues o Jira con los siguientes campos mínimos:
   - ID (ej. RF-GP001), Título, Descripción, Prioridad, Estado, Responsable, Vínculos.
2. **Estados propuestos**
   - _Propuesto_ ➜ _En análisis_ ➜ _Aprobado_ ➜ _Implementado_ ➜ _Validado_ ➜ _Cerrado_.
3. **Indicadores de seguimiento**
   - % RF implementados vs. aprobados.
   - Velocidad de cierre de requisitos por sprint.
   - Nº de defectos encontrados por requisito (calidad).
4. **Reunión quincenal de seguimiento**
   - Revisar métricas, riesgo, bloqueos.
   - Ajustar prioridades junto con la Dra. Génesis (rol _Product Owner_).

## 3. Control de Cambios

| Estrategia                             | Detalle                                                                                                    | Razón de elección                                                                 |
| -------------------------------------- | ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| **Baseline de requisitos**             | Crear una línea base al final de la fase de levantamiento. Cambios posteriores requieren solicitud formal. | Mantiene estabilidad para el desarrollo incremental.                              |
| **Solicitud de Cambio (CR)**           | Plantilla ligera: Descripción, motivación, impacto en alcance, tiempo, costo, prioridad.                   | Documenta decisiones y evita comunicación informal que derive en malentendidos.   |
| **Comité de Control de Cambios (CCB)** | Miembros: Tutor académico, Dra. Génesis, Líder de Desarrollo, Analista de Requisitos.                      | Asegura balance entre valor clínico, esfuerzo técnico y restricciones académicas. |
| **Análisis de Impacto**                | Toda CR se analiza respecto a: casos de uso afectados, UI, base de datos, pruebas y cronograma.            | Minimiza riesgos de regresión y sobrecosto.                                       |
| **Versionado Semántico de Requisitos** | Major.Minor.Patch (1.0.0 inicial). Incremento _Minor_ para nuevos RF; _Patch_ para correcciones menores.   | Facilita trazabilidad con versiones de software entregadas.                       |
| **Histórico en Git**                   | Los documentos `Requisitos.md`, matrices y CR se almacenan en control de versiones.                        | Permite auditoría completa y fácil reversión.                                     |

## 4. Flujo Simplificado de Cambio

```
sequenceDiagram
    participant Solicitante
    participant Analista as Analista Req.
    participant CCB
    participant Dev as Equipo Dev.

    Solicitante->>Analista: Envía CR
    Analista->>CCB: Pre-análisis + recomendación
    CCB-->>Analista: Aprueba / Rechaza
    Analista->>Dev: Actualiza backlog, documentos
    Dev-->>Solicitante: Entrega nueva versión
```

## 5. Buenas Prácticas Complementarias

- **Definir Definition of Done**: Incluye que el requisito esté implementado, probado, revisado por la Dra. Génesis y documentado.
- **Etiquetas de riesgo**: _Alto_, _Medio_, _Bajo_ para priorizar CR que afectan despliegue clínico.
- **Retrospectivas**: Evaluar eficacia de validación y control de cambios cada sprint y ajustar proceso.

## 6. Conclusión

La combinación de revisiones colaborativas, pruebas de aceptación, prototipado y trazabilidad garantiza que los requisitos representen fielmente las necesidades de PeluDog y puedan ser verificados objetivamente. El tablero Kanban y la línea base ofrecen visibilidad continua, mientras que el proceso liviano de CR protege al proyecto de la volatilidad sin ralentizar la entrega. Estas estrategias equilibran rigurosidad y agilidad, adecuándose al tamaño del equipo y al contexto académico-comunitario del CRM.
