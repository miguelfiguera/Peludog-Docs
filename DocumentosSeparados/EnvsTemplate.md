# Variables de Entorno - PeluDog CRM MVP

Este documento contiene las variables de entorno necesarias para el MVP del sistema PeluDog CRM, optimizadas para una instalación inicial simple y funcional.

---

## Archivo .env.mvp (Configuración MVP)

```bash
# ========================
# CONFIGURACIÓN PRINCIPAL
# ========================
RAILS_ENV=production
SECRET_KEY_BASE=generate_with_rails_secret_minimum_64_chars_long

# ========================
# BASE DE DATOS
# ========================
DATABASE_URL=mysql2://peludog:peludog123@mysql:3306/peludog_production
DB_USERNAME=peludog
DB_PASSWORD=peludog123
DB_NAME=peludog_production
DB_HOST=mysql

# ========================
# RAILS CONFIGURACIÓN BÁSICA
# ========================
RAILS_MAX_THREADS=10
WEB_CONCURRENCY=2
PORT=3000

# ========================
# ACTIVESTORAGE (LOCAL)
# ========================
ACTIVE_STORAGE_SERVICE=local
MAX_FILE_SIZE=20971520  # 20MB en bytes

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=your_jwt_secret_minimum_32_chars_generate_random
JWT_EXPIRATION_TIME=24h

# ========================
# EMAIL BÁSICO (OPCIONAL PARA MVP)
# ========================
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_gmail_app_password
FROM_EMAIL=noreply@peludog.local

# ========================
# MYSQL (DOCKER)
# ========================
MYSQL_ROOT_PASSWORD=root123
MYSQL_DATABASE=peludog_production
MYSQL_USER=peludog
MYSQL_PASSWORD=peludog123

# ========================
# BACKUP BÁSICO
# ========================
BACKUP_RETENTION_DAYS=3
BACKUP_TIME=2:00

# ========================
# CONFIGURACIÓN REGIONAL
# ========================
TZ=America/Caracas
```

---

## Archivo .env.development (Para Desarrollo Local)

```bash
# ========================
# CONFIGURACIÓN DE DESARROLLO
# ========================
RAILS_ENV=development
SECRET_KEY_BASE=development_secret_not_for_production

# ========================
# BASE DE DATOS LOCAL
# ========================
DATABASE_URL=mysql2://root:password@localhost:3306/peludog_development
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=peludog_development
DB_HOST=localhost

# ========================
# RAILS CONFIGURACIÓN DESARROLLO
# ========================
RAILS_MAX_THREADS=5
WEB_CONCURRENCY=1
PORT=3000

# ========================
# ACTIVESTORAGE LOCAL
# ========================
ACTIVE_STORAGE_SERVICE=local
MAX_FILE_SIZE=10485760   # 10MB para desarrollo

# ========================
# JWT DESARROLLO
# ========================
JWT_SECRET_KEY=development_jwt_not_for_production
JWT_EXPIRATION_TIME=24h

# ========================
# EMAIL DESARROLLO (OPCIONAL)
# ========================
SMTP_HOST=localhost
SMTP_PORT=1025  # MailCatcher para testing
FROM_EMAIL=dev@peludog.local

# ========================
# CONFIGURACIÓN REGIONAL
# ========================
TZ=America/Caracas
```

---

## Generación de Claves Seguras

Para generar las claves necesarias de forma segura:

```bash
# Generar SECRET_KEY_BASE
rails secret

# O manualmente
openssl rand -hex 64

# Generar JWT_SECRET_KEY
openssl rand -hex 32
```

---

## Notas Importantes para MVP

1. **Simplicidad Primero:** Esta configuración prioriza facilidad de instalación sobre optimizaciones complejas
2. **Archivos Locales:** ActiveStorage configurado para almacenamiento local (sin S3 inicialmente)
3. **Base de Datos Simple:** MySQL básico sin réplicas ni optimizaciones avanzadas
4. **Email Opcional:** Las notificaciones por email son opcionales para el MVP
5. **Backups Básicos:** Sistema simple de backups, ampliable posteriormente
6. **Escalabilidad:** Estas variables son fácilmente expandibles cuando se necesite escalar

---

## Escalamiento Futuro

Cuando el sistema crezca, estas variables se pueden extender agregando:
- Redis para caché y sesiones
- Configuraciones de load balancer
- Múltiples instancias de base de datos
- Storage distribuido (S3/MinIO)
- Configuraciones avanzadas de email y notificaciones

Ver **AnexosArquitectura.md** para configuraciones completas de escalamiento.

---

## E2. Variables de Entorno - Escalamiento Vertical

### E2.1 Archivo .env.vertical_scaling

```bash
# ========================
# CONFIGURACIÓN PRINCIPAL
# ========================
RAILS_ENV=production
SECRET_KEY_BASE=your_secret_key_base_here_minimum_64_chars_long_random_string

# ========================
# BASE DE DATOS (OPTIMIZADA)
# ========================
DATABASE_URL=mysql2://peludog:secure_password@mysql:3306/peludog_production
DB_USERNAME=peludog
DB_PASSWORD=secure_password_here
DB_NAME=peludog_production
DB_HOST=mysql
DB_POOL_SIZE=30  # Incrementado para más conexiones

# ========================
# REDIS (CACHE Y SESIONES)
# ========================
REDIS_URL=redis://redis:6379/0
REDIS_POOL_SIZE=30

# ========================
# RAILS CONFIGURACIÓN (ESCALADO)
# ========================
RAILS_MAX_THREADS=25      # Incrementado de 20 a 25
WEB_CONCURRENCY=6         # Incrementado de 4 a 6
RAILS_SERVE_STATIC_FILES=false
PORT=3000

# ========================
# PUMA CONFIGURACIÓN AVANZADA
# ========================
PUMA_WORKER_TIMEOUT=60
PUMA_WORKER_BOOT_TIMEOUT=60
PUMA_PRELOAD_APP=true

# ========================
# ACTIVESTORAGE
# ========================
ACTIVE_STORAGE_SERVICE=local
MAX_FILE_SIZE=104857600   # 100MB en bytes (incrementado)
ACTIVE_STORAGE_CACHE_TTL=86400  # 24 horas

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=your_jwt_secret_key_here_minimum_32_chars
JWT_EXPIRATION_TIME=24h

# ========================
# EMAIL (NOTIFICACIONES)
# ========================
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password_here
FROM_EMAIL=noreply@peludog.com
EMAIL_POOL_SIZE=5

# ========================
# MYSQL (OPTIMIZADA PARA VERTICAL SCALING)
# ========================
MYSQL_ROOT_PASSWORD=rootpassword_secure
MYSQL_DATABASE=peludog_production
MYSQL_USER=peludog
MYSQL_PASSWORD=secure_password_here

# Configuraciones MySQL optimizadas
MYSQL_INNODB_BUFFER_POOL_SIZE=4G
MYSQL_MAX_CONNECTIONS=300
MYSQL_QUERY_CACHE_SIZE=256M

# ========================
# BACKUP CONFIGURACIÓN
# ========================
BACKUP_RETENTION_DAYS=6
BACKUP_TIME=2:00
CLEANUP_TIME=2:30
BACKUP_COMPRESSION=gzip

# ========================
# LOGGING Y MONITOREO
# ========================
LOG_LEVEL=info
LOG_ROTATION=daily
ENABLE_METRICS=true

# ========================
# TIMEZONE
# ========================
TZ=America/Caracas
```

---

## E3. Variables de Entorno - Escalamiento Horizontal

### E3.1 Archivo .env.scaling

```bash
# ========================
# CONFIGURACIÓN PRINCIPAL
# ========================
RAILS_ENV=production
SECRET_KEY_BASE=your_secret_key_base_here_minimum_64_chars_long_random_string

# ========================
# BASE DE DATOS (MASTER-REPLICA)
# ========================
DB_USERNAME=peludog
DB_PASSWORD=secure_password_here
DB_NAME=peludog_production

# Master Database
DB_PRIMARY_HOST=mysql-master
DB_PRIMARY_PORT=3306

# Replica Database
DB_REPLICA_HOST=mysql-replica
DB_REPLICA_PORT=3306

# Pool de conexiones para múltiples instancias
DB_POOL_SIZE=25

# ========================
# REDIS CLUSTER
# ========================
REDIS_URL=redis://redis-master:6379/0
REDIS_REPLICA_URL=redis://redis-replica:6379/0
REDIS_POOL_SIZE=25

# ========================
# RAILS CONFIGURACIÓN (MÚLTIPLES INSTANCIAS)
# ========================
RAILS_MAX_THREADS=25
WEB_CONCURRENCY=6
RAILS_SERVE_STATIC_FILES=false
PORT=3000

# Configuración específica para load balancer
RAILS_RELATIVE_URL_ROOT=/
FORCE_SSL=true

# ========================
# LOAD BALANCER CONFIGURACIÓN
# ========================
BACKEND_INSTANCES=3
NGINX_WORKER_PROCESSES=4
NGINX_WORKER_CONNECTIONS=2048
NGINX_KEEPALIVE_TIMEOUT=65
NGINX_CLIENT_MAX_BODY_SIZE=100M

# ========================
# ACTIVESTORAGE (DISTRIBUIDO)
# ========================
ACTIVE_STORAGE_SERVICE=amazon
MAX_FILE_SIZE=104857600   # 100MB

# AWS S3 / MinIO Configuration
AWS_ACCESS_KEY_ID=minioadmin
AWS_SECRET_ACCESS_KEY=minioadmin123456
AWS_REGION=us-east-1
AWS_BUCKET=peludog-storage
AWS_ENDPOINT=http://storage-server:9000  # Para MinIO

# MinIO específico
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin123456
MINIO_ENDPOINT=http://storage-server:9000
MINIO_BUCKET=peludog-storage

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=your_jwt_secret_key_here_minimum_32_chars
JWT_EXPIRATION_TIME=24h

# ========================
# EMAIL (NOTIFICACIONES)
# ========================
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password_here
FROM_EMAIL=noreply@peludog.com
EMAIL_POOL_SIZE=10

# ========================
# MYSQL MASTER
# ========================
MYSQL_MASTER_ROOT_PASSWORD=rootpassword_secure
MYSQL_MASTER_DATABASE=peludog_production
MYSQL_MASTER_USER=peludog
MYSQL_MASTER_PASSWORD=secure_password_here
MYSQL_MASTER_SERVER_ID=1

# ========================
# MYSQL REPLICA
# ========================
MYSQL_REPLICA_ROOT_PASSWORD=rootpassword_secure
MYSQL_REPLICA_DATABASE=peludog_production
MYSQL_REPLICA_USER=peludog
MYSQL_REPLICA_PASSWORD=secure_password_here
MYSQL_REPLICA_SERVER_ID=2

# ========================
# SESIONES COMPARTIDAS
# ========================
SESSION_STORE=redis
SESSION_KEY=_peludog_session
SESSION_SECRET=your_session_secret_key_here
SESSION_EXPIRE=604800  # 1 semana en segundos

# ========================
# BACKUP CONFIGURACIÓN
# ========================
BACKUP_RETENTION_DAYS=7
BACKUP_TIME=2:00
CLEANUP_TIME=2:30
BACKUP_COMPRESSION=gzip
BACKUP_S3_BUCKET=peludog-backups

# ========================
# LOGGING Y MONITOREO
# ========================
LOG_LEVEL=info
LOG_ROTATION=daily
ENABLE_METRICS=true
PROMETHEUS_PORT=9090
GRAFANA_PORT=3001

# ========================
# SSL/TLS
# ========================
SSL_CERT_PATH=/etc/ssl/certs/peludog.crt
SSL_KEY_PATH=/etc/ssl/private/peludog.key
FORCE_SSL=true

# ========================
# TIMEZONE
# ========================
TZ=America/Caracas
```

---

## E4. Variables de Entorno - Arquitectura Distribuida

### E4.1 Archivo .env.distributed

```bash
# ========================
# CONFIGURACIÓN PRINCIPAL
# ========================
RAILS_ENV=production
SECRET_KEY_BASE=your_secret_key_base_here_minimum_64_chars_long_random_string
APP_HOST=peludog.com
API_VERSION=v1

# ========================
# BASE DE DATOS (SERVIDORES DEDICADOS)
# ========================
DB_USERNAME=peludog
DB_PASSWORD=secure_password_here
DB_NAME=peludog_production

# Master Database (Servidor dedicado)
DB_PRIMARY_HOST=10.0.1.10
DB_PRIMARY_PORT=3306

# Read Replicas (Servidores dedicados)
DB_REPLICA_1_HOST=10.0.1.11
DB_REPLICA_1_PORT=3306
DB_REPLICA_2_HOST=10.0.1.12
DB_REPLICA_2_PORT=3306

# Pool de conexiones distribuidas
DB_POOL_SIZE=30
DB_TIMEOUT=5000
DB_RETRY_ATTEMPTS=3

# ========================
# REDIS CLUSTER (SERVIDORES DEDICADOS)
# ========================
REDIS_CLUSTER_URLS=redis://10.0.2.10:6379,redis://10.0.2.11:6379,redis://10.0.2.12:6379
REDIS_POOL_SIZE=30
REDIS_TIMEOUT=1000

# ========================
# RAILS CONFIGURACIÓN (SERVIDORES DE APLICACIÓN)
# ========================
RAILS_MAX_THREADS=30
WEB_CONCURRENCY=8
RAILS_SERVE_STATIC_FILES=false
PORT=3000

# Configuración para múltiples servidores de app
APP_SERVER_COUNT=6
INSTANCE_ID=${HOSTNAME}

# ========================
# LOAD BALANCER (HAProxy/NGINX)
# ========================
LB_PRIMARY_HOST=10.0.0.10
LB_SECONDARY_HOST=10.0.0.11
HEALTH_CHECK_PATH=/health
HEALTH_CHECK_INTERVAL=5s

# Servidores de aplicación
APP_SERVER_1=10.0.3.10:3000
APP_SERVER_2=10.0.3.11:3000
APP_SERVER_3=10.0.3.12:3000
APP_SERVER_4=10.0.3.13:3000
APP_SERVER_5=10.0.3.14:3000
APP_SERVER_6=10.0.3.15:3000

# ========================
# ACTIVESTORAGE (S3 COMPATIBLE)
# ========================
ACTIVE_STORAGE_SERVICE=amazon
MAX_FILE_SIZE=209715200   # 200MB

# AWS S3 / Servidor de almacenamiento dedicado
AWS_ACCESS_KEY_ID=your_access_key_here
AWS_SECRET_ACCESS_KEY=your_secret_key_here
AWS_REGION=us-east-1
AWS_BUCKET=peludog-production-files
AWS_ENDPOINT=https://s3.peludog.com  # Servidor S3 dedicado

# CDN para archivos estáticos
CDN_HOST=cdn.peludog.com
CDN_ENABLED=true

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=your_jwt_secret_key_here_minimum_32_chars
JWT_EXPIRATION_TIME=24h
JWT_REFRESH_TOKEN_EXPIRATION=7d

# ========================
# EMAIL (SERVIDOR DEDICADO)
# ========================
SMTP_HOST=mail.peludog.com
SMTP_PORT=587
SMTP_USERNAME=system@peludog.com
SMTP_PASSWORD=your_email_password_here
FROM_EMAIL=noreply@peludog.com
EMAIL_POOL_SIZE=20

# ========================
# SESIONES DISTRIBUIDAS
# ========================
SESSION_STORE=redis_cluster
SESSION_KEY=_peludog_session
SESSION_SECRET=your_session_secret_key_here
SESSION_EXPIRE=604800  # 1 semana
STICKY_SESSIONS=false

# ========================
# BACKUP CONFIGURACIÓN (DISTRIBUIDA)
# ========================
BACKUP_RETENTION_DAYS=14
BACKUP_TIME=2:00
CLEANUP_TIME=2:30
BACKUP_COMPRESSION=gzip

# Backup a múltiples ubicaciones
BACKUP_PRIMARY_S3_BUCKET=peludog-backups-primary
BACKUP_SECONDARY_S3_BUCKET=peludog-backups-secondary
BACKUP_OFFSITE_ENABLED=true

# ========================
# MONITORING Y OBSERVABILIDAD
# ========================
LOG_LEVEL=info
LOG_ROTATION=daily
ENABLE_METRICS=true

# Prometheus
PROMETHEUS_HOST=monitoring.peludog.com
PROMETHEUS_PORT=9090

# Grafana
GRAFANA_HOST=metrics.peludog.com
GRAFANA_PORT=3000

# Alertmanager
ALERTMANAGER_HOST=alerts.peludog.com
ALERTMANAGER_PORT=9093

# ========================
# SEGURIDAD
# ========================
SSL_CERT_PATH=/etc/ssl/certs/peludog.crt
SSL_KEY_PATH=/etc/ssl/private/peludog.key
FORCE_SSL=true
HSTS_MAX_AGE=31536000

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=1000
RATE_LIMIT_PERIOD=3600

# ========================
# PERFORMANCE
# ========================
ENABLE_CACHING=true
CACHE_EXPIRES=3600
ENABLE_COMPRESSION=true
ENABLE_HTTP2=true

# ========================
# TIMEZONE
# ========================
TZ=America/Caracas

# ========================
# FEATURES FLAGS
# ========================
ENABLE_API_VERSIONING=true
ENABLE_ANALYTICS=true
ENABLE_AUDIT_LOG=true
ENABLE_BACKGROUND_JOBS=true
```

---

## E5. Variables de Desarrollo

### E5.1 Archivo .env.development

```bash
# ========================
# CONFIGURACIÓN DE DESARROLLO
# ========================
RAILS_ENV=development
SECRET_KEY_BASE=development_secret_key_base_not_for_production

# ========================
# BASE DE DATOS
# ========================
DATABASE_URL=mysql2://root:password@localhost:3306/peludog_development
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=peludog_development
DB_HOST=localhost

# ========================
# REDIS
# ========================
REDIS_URL=redis://localhost:6379/0

# ========================
# RAILS CONFIGURACIÓN
# ========================
RAILS_MAX_THREADS=5
WEB_CONCURRENCY=1
RAILS_SERVE_STATIC_FILES=true
PORT=3000

# ========================
# ACTIVESTORAGE
# ========================
ACTIVE_STORAGE_SERVICE=local
MAX_FILE_SIZE=10485760   # 10MB para desarrollo

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=development_jwt_secret_not_for_production
JWT_EXPIRATION_TIME=24h

# ========================
# EMAIL (DESARROLLO)
# ========================
SMTP_HOST=localhost
SMTP_PORT=1025  # MailCatcher
SMTP_USERNAME=
SMTP_PASSWORD=
FROM_EMAIL=dev@peludog.local

# ========================
# LOGGING
# ========================
LOG_LEVEL=debug
ENABLE_METRICS=false

# ========================
# TIMEZONE
# ========================
TZ=America/Caracas
```

---

## E6. Variables de Testing

### E6.1 Archivo .env.test

```bash
# ========================
# CONFIGURACIÓN DE TESTING
# ========================
RAILS_ENV=test
SECRET_KEY_BASE=test_secret_key_base_not_for_production

# ========================
# BASE DE DATOS
# ========================
DATABASE_URL=mysql2://root:password@localhost:3306/peludog_test
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=peludog_test
DB_HOST=localhost

# ========================
# REDIS
# ========================
REDIS_URL=redis://localhost:6379/1

# ========================
# RAILS CONFIGURACIÓN
# ========================
RAILS_MAX_THREADS=5
WEB_CONCURRENCY=1
RAILS_SERVE_STATIC_FILES=true
PORT=3001

# ========================
# ACTIVESTORAGE
# ========================
ACTIVE_STORAGE_SERVICE=test
MAX_FILE_SIZE=1048576   # 1MB para tests

# ========================
# JWT AUTENTICACIÓN
# ========================
JWT_SECRET_KEY=test_jwt_secret_not_for_production
JWT_EXPIRATION_TIME=1h

# ========================
# EMAIL (TESTING)
# ========================
SMTP_HOST=localhost
SMTP_PORT=1025
FROM_EMAIL=test@peludog.test

# ========================
# TESTING
# ========================
PARALLEL_WORKERS=4
ENABLE_COVERAGE=true
HEADLESS_BROWSER=true

# ========================
# TIMEZONE
# ========================
TZ=America/Caracas
```