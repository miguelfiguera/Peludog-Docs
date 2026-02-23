# Arquitectura Técnica Establecida - PeluDog CRM

## 1. Introducción y Filosofía

Esta arquitectura técnica definitiva establece el diseño técnico completo para el sistema PeluDog CRM. La filosofía central es crear un sistema robusto, mantenible y, fundamentalmente, fácil de instalar y desplegar para la comunidad de código abierto. Está diseñado para escalar desde una implementación inicial en un único servidor hasta una arquitectura distribuida conforme crezca la demanda.

La elección de tecnologías busca un equilibrio entre madurez, productividad de desarrollo, facilidad de escalamiento y un ecosistema sólido que facilite la implementación de las funcionalidades requeridas.

![Arquitectura Inicial](../../Imagenes/SystemDesignDiagramsImages/initialSetup.png)

---

## 2. Arquitectura de 3 Capas Containerizada

### **Capa de Backend (Servidor)**

- **Tecnología:** Ruby on Rails 7+ (modo `api_only`)
- **Ruby Version:** 3.3+
- **Responsabilidades:**

  - Exponer una **API RESTful** segura para que el frontend consuma los datos
  - Implementar toda la **lógica de negocio** y las reglas del dominio (cálculos de vacunación, gestión de citas, etc.)
  - Gestionar la **autenticación y autorización** de usuarios mediante JWT
  - Realizar **validaciones de datos** a nivel de modelo para garantizar la integridad
  - Gestionar **backups automáticos** de la base de datos
  - Manejar **archivos multimedia** con ActiveStorage

- **Modo `api_only`:** El backend no generará vistas HTML. Su única responsabilidad es recibir y responder a peticiones con formato JSON, manteniendo la arquitectura ligera y enfocada.

- **Autenticación (JWT):**
  - Sistema de autenticación stateless basado en **JSON Web Tokens**
  - **No se usarán sesiones del lado del servidor**
  - Controladores específicos para registro, inicio de sesión y recuperación de contraseña
  - Token JWT firmado que el cliente almacena y envía en la cabecera `Authorization: Bearer <token>`

### **Capa de Frontend (Aplicación Cliente)**

- **Tecnología:** **Expo (React Native) con TypeScript**
- **Versiones:** Expo SDK 52+, React Native 0.76+
- **Enfoque actual:** **Aplicación web** compilada desde Expo (preparada para extensión móvil futura)
- **Responsabilidades:**

  - Renderizar la **interfaz de usuario web** para veterinarios, asistentes y clientes
  - Gestionar el **estado de la aplicación** con React Hooks reutilizables
  - Realizar peticiones a la **API del backend**
  - Incluir **Landing Page integrada** como parte de la aplicación web
  - Ser **responsive** para funcionar en escritorio, tablet y móvil via navegador

- **Tecnologías Complementarias:**
  - **Estilos:** Tailwind CSS con NativeWind
  - **Notificaciones:** `react-native-toast-message`
  - **HTTP Client:** Axios con interceptores para JWT automático
  - **Tipado:** TypeScript para interfaces y tipos seguros

### **Capa de Persistencia (Base de Datos)**

- **Base de Datos:** MySQL 8.0+
- **Responsabilidades:**
  - Almacenar datos persistentes: usuarios, clientes, mascotas, historias clínicas, citas, pagos
  - Soporte para backups automáticos diarios
  - Gestión de volúmenes persistentes para archivos multimedia

---

## 3. Gestión de Archivos y ActiveStorage

### **Estrategia de Almacenamiento**

ActiveStorage se configurará inicialmente para persistencia local en disco, facilitando la instalación y configuración inicial. La estructura de almacenamiento será organizada para garantizar la integridad y facilitar futuras migraciones.

**Estructura de Almacenamiento Local:**

- Archivos originales organizados por fecha y tipo
- Variaciones de imagen generadas automáticamente (thumbnails)
- Archivos temporales con limpieza automática

**Tipos de Archivos Soportados:**

- **Imágenes:** JPEG, PNG, WebP (fotos de mascotas, documentos escaneados)
- **Documentos:** PDF (historias clínicas, reportes)
- **Límites:** 50MB por archivo inicialmente, escalable según necesidades

---

## 4. Sistema de Backups Automatizado

### **Estrategia de Backup con Whenever Gem**

El sistema implementará backups automáticos diarios utilizando la gema Whenever para programar tareas cron. Esta estrategia garantiza la continuidad del servicio y la protección de datos críticos.

**Configuración de Backups:**

- **Frecuencia:** Diaria a las 2:00 AM
- **Retención:** Se conservan los 6 backups más recientes.
- **Compresión:** Automática con gzip para optimizar espacio
- **Restauración:** Comando específico desde Rails para recuperación rápida

---

## 5. Orquestación y Despliegue con Docker

### **Estrategia de Containerización**

La arquitectura utiliza Docker Compose para orquestar todos los servicios necesarios, manteniendo la simplicidad de instalación mientras proporciona flexibilidad para el escalamiento futuro.

**Servicios Containerizados:**

- **NGINX:** Proxy inverso y servidor de archivos estáticos
- **Rails Backend:** API server con todas las dependencias
- **Frontend Web:** Contenedor que sirve la aplicación web compilada desde Expo (SPA/PWA).
- **MySQL:** Base de datos con configuraciones optimizadas

**Volúmenes Persistentes:**

- Base de datos MySQL
- Almacenamiento ActiveStorage
- Backups de base de datos
- Archivos estáticos del frontend

---

## 6. Estrategia de Escalamiento

### **Importancia del Escalamiento Vertical Primero**

El escalamiento vertical debe ser la primera opción antes de considerar el escalamiento horizontal. Esto es crítico por varias razones:

1. **Simplicidad Operacional:** Mantener un único servidor es significativamente más simple de administrar, monitorear y debuggear
2. **Menor Complejidad de Código:** No requiere modificaciones en la aplicación para manejar estado distribuido
3. **Costos Reducidos:** Un servidor más potente es generalmente más económico que múltiples servidores más pequeños
4. **Menor Latencia:** Sin overhead de comunicación entre servicios
5. **Facilidad de Backup y Restauración:** Un solo punto de datos facilita las operaciones de backup

### **6.1 Fase 1: Escalamiento Vertical Extensivo (0-300 usuarios concurrentes)**

![Escalamiento Vertical](../../Imagenes/SystemDesignDiagramsImages/verticalScaling.png)

Como primer paso en el escalamiento vertical, y antes de escalar horizontalmente, se puede introducir **Redis** para gestionar el caché de la aplicación y aliviar la carga sobre la base de datos.

**Hardware Progresivo:**

- **Básico (Inicial):** 2 cores, 2GB RAM, 50GB SSD
- **Medio:** 4 cores, 8GB RAM, 500GB SSD
- **Avanzado:** 8 cores, 16GB RAM, 1TB SSD
- **Potente:** 16 cores, 32GB RAM, 2TB NVMe SSD

**Optimización de Workers y Threads:**
La estrategia de escalamiento vertical consiste en empezar con una configuración conservadora (ej. 1 worker, 5 threads) y, a medida que se aumentan los recursos del servidor (CPU y RAM), incrementar gradualmente el número de workers para mejorar la capacidad de respuesta y el volumen de peticiones que el sistema puede manejar simultáneamente. El número de conexiones a la base de datos y la memoria caché de Redis también deben ajustarse en consonancia.

**Límites Reales del Escalamiento Vertical:**
Un servidor bien optimizado puede manejar entre 250-300 usuarios concurrentes reales (no solo conexiones) antes de necesitar escalamiento horizontal. Este límite considera:

- Operaciones complejas de base de datos (historias clínicas, reportes)
- Subida/descarga de archivos multimedia
- Procesamiento de imágenes con ActiveStorage
- Consultas pesadas de reportes y analíticas

### **6.2 Transición al Escalamiento Horizontal (300+ usuarios concurrentes)**

**Cuándo Escalar Horizontalmente:**

- CPU constantemente por encima del 80%
- Memoria RAM utilizada por encima del 85%
- Tiempo de respuesta de la base de datos > 200ms consistentemente
- Cola de requests en NGINX/Puma creciendo constantemente

### **6.3 Gestión de Variables de Entorno en Load Balancer**

**Problema de Consistencia de ENV:**
Cuando se escala horizontalmente, todas las instancias de Rails deben compartir las mismas variables de entorno para garantizar comportamiento consistente. Esto incluye:

- Claves de encriptación (SECRET_KEY_BASE, JWT_SECRET_KEY)
- Configuraciones de base de datos
- Configuraciones de Redis
- Configuraciones de ActiveStorage

**Solución con Archivos ENV Compartidos:**

1. **Archivo ENV Central:** Un único archivo `.env.scaling` compartido entre todas las instancias
2. **Montaje de Volumen:** Docker monta este archivo en todas las instancias de backend
3. **Sincronización:** Cualquier cambio en variables se refleja inmediatamente en todas las instancias
4. **Versionado:** Git tracks de archivos ENV para control de cambios

**Gestión de Sesiones Distribuidas:**

- **Redis Centralizado:** Todas las instancias comparten el mismo Redis para sesiones
- **JWT Stateless:** Preferir JWT sobre sesiones cuando sea posible
- **Sticky Sessions:** NGINX puede configurarse para mantener usuarios en la misma instancia

### **6.4 Separación de ActiveStorage y Race Conditions**

**Problemas con ActiveStorage Distribuido:**
Cuando múltiples instancias de Rails manejan subida de archivos simultáneamente, pueden ocurrir race conditions y corrupción de datos:

1. **Race Conditions en Escritura:** Dos instancias intentando escribir el mismo archivo
2. **Inconsistencia de Metadatos:** Base de datos actualizada en una instancia pero archivo no disponible en otras
3. **Limpieza de Archivos Temporales:** Conflictos cuando múltiples instancias limpian archivos

**Solución: Servidor de Storage Dedicado**

1. **MinIO como Servidor Dedicado:** Separar ActiveStorage a un contenedor MinIO dedicado
2. **API S3 Compatible:** Todas las instancias Rails hablan con MinIO via API S3
3. **Consistencia Garantizada:** MinIO maneja la consistencia y concurrencia automáticamente
4. **Eliminación de Race Conditions:** MinIO es responsible de la escritura atómica

**Proceso de Migración de Storage:**

1. **Instalación de MinIO:** Nuevo contenedor con volúmenes persistentes
2. **Configuración de Rails:** Cambiar ActiveStorage de `local` a `amazon` (MinIO)
3. **Migración de Archivos Existentes:** Rake task para transferir archivos locales a MinIO
4. **Verificación de Integridad:** Confirmar que todos los archivos fueron migrados correctamente
5. **Limpieza:** Eliminar archivos locales después de verificar migración

**Configuración de MinIO en Escalamiento:**

- **Replicación:** MinIO puede configurarse con múltiples nodos para redundancia
- **CDN:** DigitalOcean Spaces se integra nativamente con su propio servicio de CDN para mejorar la performance.
- **Backup:** MinIO se integra con estrategias de backup existentes

### 6.5 Base de Datos Distribuida (Master-Slave)

Para soportar la carga en un entorno horizontal, la base de datos también debe distribuirse. La estrategia a seguir es un modelo de replicación Master-Slave (o Primario-Réplica):

1.  **Servidor Maestro (Master/Primary):** Una única instancia de base de datos que maneja **todas las operaciones de escritura** (INSERT, UPDATE, DELETE). Esto garantiza la consistencia y la integridad de los datos al tener una única fuente de verdad.
2.  **Servidores Réplica (Slaves/Replicas):** Una o más instancias de base de datos que son copias de solo lectura del servidor maestro. Su único propósito es manejar **operaciones de lectura** (SELECT).
3.  **Flujo de Datos:** La aplicación Rails se configura para dirigir todas las consultas de escritura al servidor maestro y distribuir las consultas de lectura entre las diversas réplicas. Esto reduce drásticamente la carga sobre la base de datos principal, mejorando el rendimiento de las consultas, que suelen ser la operación más frecuente.
4.  **Alta Disponibilidad:** Este modelo también proporciona una capa de redundancia. Si el servidor maestro falla, una de las réplicas puede ser promovida para convertirse en el nuevo maestro, minimizando el tiempo de inactividad.

La configuración para que Rails pueda manejar esta topología se detalla en `AnexosArquitectura.md`, donde se define una conexión `primary` para escrituras y una `replica` para lecturas en `config/database.yml`.

### 6.6 Componentes de la Arquitectura Distribuida Final

![Escalamiento Horizontal](../../Imagenes/SystemDesignDiagramsImages/horizontalScaling.png)

La arquitectura para 1000+ usuarios concurrentes evoluciona a un sistema distribuido donde cada componente está desacoplado para maximizar el rendimiento y la disponibilidad. A continuación se detallan los componentes clave:

#### **CDN (Content Delivery Network)**

- **Función:** Actúa como la primera capa de contacto con el usuario. Es una red de servidores distribuidos geográficamente que almacena en caché los activos estáticos de la aplicación.
- **Detalle Técnico:** Cachea los archivos compilados del frontend (JavaScript, CSS) y los archivos subidos por los usuarios desde el servidor de Storage (MinIO). Esto reduce drásticamente la latencia para los usuarios de todo el mundo y disminuye la carga sobre la infraestructura principal.

#### **Load Balancer Principal (NGINX)**

- **Función:** Es el único punto de entrada a la infraestructura de la nube. Recibe todo el tráfico que no fue servido por la caché del CDN.
- **Detalle Técnico y Flujo de Peticiones:** Opera como un balanceador de carga de Capa 7, lo que le permite inspeccionar la URL de la petición para enrutar el tráfico de forma inteligente. El flujo es el siguiente:
  1. Una petición del usuario llega al Load Balancer.
  2. El LB analiza la ruta. Si la ruta comienza con `/api/`, la reconoce como una llamada a la API y la reenvía de forma segura a una de las instancias disponibles en la **flota de servidores Backend**.
  3. Si la ruta es cualquier otra (ej. `/`, `/citas`, `/perfil`), el LB la identifica como una petición para la aplicación web y la reenvía a una instancia de la **flota de servidores Frontend**.
- **Beneficio:** Este mecanismo de enrutamiento basado en la ruta es crucial. Permite que tanto el frontend como el backend compartan el mismo dominio y punto de entrada, simplificando la configuración de DNS y SSL/TLS, al tiempo que permite escalar cada flota de servidores de forma independiente.

#### **Flota de Servidores Frontend**

- **Función:** Un grupo de servidores idénticos cuya única responsabilidad es servir los archivos estáticos (HTML, CSS, JS) que componen la aplicación web compilada desde Expo (SPA/PWA).
- **Detalle Técnico:** Al ser servidores sin estado, se pueden añadir o quitar instancias fácilmente según la demanda de tráfico. El Load Balancer se encarga de distribuir las peticiones entre ellos, garantizando alta disponibilidad.

#### **Sistema de Monitoreo (Prometheus)**

- **Función:** Proporciona observabilidad completa sobre la salud y el rendimiento de toda la infraestructura.
- **Detalle Técnico:** Es un sistema centralizado que recolecta métricas de todos los componentes del sistema (Load Balancer, instancias de Frontend y Backend, Base de Datos, Redis, etc.). Esta recolección de datos (scraping) es fundamental para alimentar los dashboards de monitoreo y para el sistema de alertas críticas definido en la sección 7.

---

## 7. Monitoreo y Observabilidad

### **Métricas Clave de Escalamiento**

- **Performance:** Tiempo de respuesta API, throughput por instancia
- **Base de Datos:** Conexiones activas, queries lentas, replication lag
- **Archivos:** Espacio usado, velocidad de upload, race conditions detectadas
- **Backups:** Éxito/fallo de backups, tiempo de restauración
- **Workers:** Utilización de workers, queue depth, thread pool

### **Alertas Críticas**

- CPU por encima del 85% por más de 5 minutos
- Memoria RAM por encima del 90%
- Tiempo de respuesta de base de datos > 300ms
- Fallos de backup
- Errores en ActiveStorage (indicativo de race conditions)

---

## 8. Seguridad y Compliance

### **Medidas de Seguridad Escalables**

- **Datos Médicos:** Encriptación en reposo y tránsito, independiente de la arquitectura
- **Backups:** Encriptación de archivos de backup, múltiples ubicaciones en arquitectura distribuida
- **Acceso:** Logs de auditoría centralizados
- **SSL/TLS:** Certificados automáticos con Let's Encrypt, load balancer termination

### **Compliance en Arquitectura Distribuida**

- **LGPD/GDPR:** Derecho al olvido implementado en todas las instancias
- **Logs de Auditoría:** Centralización de logs para compliance
- **Retención de Datos:** Políticas consistentes across all services

---

## 9. Objetivo Estratégico de Replicabilidad

La arquitectura está diseñada para ser **fácilmente replicable** por otras clínicas veterinarias con mínima personalización. La estructura de datos común permite:

1. **Intercambio seguro de información** entre clínicas
2. **Referencia de pacientes** sin transferencias manuales
3. **Actualizaciones centralizadas** del sistema
4. **Comunidad de desarrollo** colaborativa

El proceso de instalación mantiene la simplicidad de `docker-compose up` en la fase inicial, con un path claro de escalamiento que no requiere reescritura de la aplicación base.

---

## 10. Documentación Técnica Complementaria

Para detalles técnicos específicos, configuraciones y código de implementación, consultar:

- **[AnexosArquitectura.md](./AnexosArquitectura.md)** - Configuraciones técnicas completas
- **[EnvsTemplate.md](./EnvsTemplate.md)** - Variables de entorno para todas las fases
