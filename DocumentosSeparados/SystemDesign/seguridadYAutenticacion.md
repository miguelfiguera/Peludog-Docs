# Seguridad y Autenticación - PeluDog CRM

## Autenticación con JWT

### Implementación Principal
JWT (JSON Web Tokens) será el mecanismo principal de autenticación en PeluDog CRM, proporcionando:

- **Stateless Authentication**: No requiere almacenamiento de sesiones en el servidor
- **Escalabilidad**: Permite distribución horizontal sin compartir estado de sesión
- **Flexibilidad**: Compatible con aplicaciones web (SPA) y preparado para móvil nativo
- **Seguridad**: Tokens firmados digitalmente con algoritmos robustos (RS256/HS256)

### Estructura del Token
```json
{
  "header": {
    "alg": "RS256",
    "typ": "JWT"
  },
  "payload": {
    "user_id": 123,
    "role": "admin",
    "exp": 1641513600,
    "iat": 1640908800,
    "iss": "peludog-crm"
  }
}
```

### Flujo de Autenticación
1. Usuario envía credenciales (email/password)
2. Servidor valida credenciales contra base de datos
3. Servidor genera JWT firmado con clave privada
4. Cliente almacena token (localStorage/sessionStorage)
5. Cliente incluye token en header Authorization para requests subsecuentes
6. Servidor valida token en cada request protegido

### Gestión de Tokens
- **Access Token**: Duración de 7 días (estándar suficiente para inicio)
- **Almacenamiento**: localStorage/sessionStorage del navegador (aplicación web)
- **Expiración**: Usuario debe autenticarse nuevamente después de 7 días
- **Revocación**: Lista negra opcional para tokens comprometidos (implementación futura)
- **Escalabilidad**: Refresh tokens se agregarán en fase de crecimiento vertical
- **Compatibilidad futura**: Preparado para AsyncStorage cuando se extienda a móvil nativo

## Seguridad Integrada de Rails

### Protecciones Nativas

#### CSRF (Cross-Site Request Forgery)
- **protect_from_forgery**: Habilitado por defecto en ApplicationController
- **Tokens CSRF**: Generados automáticamente para formularios
- **Verificación**: Rails valida tokens en requests POST/PUT/DELETE
- **API Mode**: Uso de tokens JWT elimina necesidad de CSRF tokens

#### SQL Injection Prevention
- **Active Record**: Sanitización automática de parámetros
- **Prepared Statements**: Queries parametrizadas por defecto
- **Strong Parameters**: Filtrado explícito de parámetros permitidos

```ruby
def user_params
  params.require(:user).permit(:name, :email, :role)
end
```

#### XSS (Cross-Site Scripting) Protection
- **html_safe**: Escapado automático de contenido HTML
- **Content Security Policy**: Headers configurables para prevenir XSS
- **Sanitización**: Helper methods para limpiar input malicioso

#### Mass Assignment Protection
- **Strong Parameters**: Previene asignación masiva no autorizada
- **attr_accessible**: Control granular de atributos modificables

#### Session Security
- **Secure Cookies**: Configuración automática para HTTPS
- **HttpOnly**: Previene acceso JavaScript a cookies de sesión
- **SameSite**: Protección contra ataques CSRF modernos

### Configuraciones de Seguridad Adicionales
```ruby
# config/application.rb
config.force_ssl = true
config.ssl_options = {
  redirect: { status: 301, body: [] },
  secure_cookies: true
}

# Content Security Policy
config.content_security_policy do |policy|
  policy.default_src :self
  policy.script_src :self
  policy.style_src :self, :unsafe_inline
end
```

## Nginx como Proxy Inverso

### Ventajas de Seguridad

#### Rate Limiting
```nginx
# Limitación de requests por IP
limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;
limit_req zone=api burst=20 nodelay;

# Protección contra ataques de fuerza bruta
location /api/auth/login {
    limit_req zone=login burst=5 nodelay;
}
```

#### SSL/TLS Termination
- **Certificados SSL**: Manejo centralizado de certificados
- **Perfect Forward Secrecy**: Configuración de cipher suites seguros
- **HSTS**: Strict Transport Security headers automáticos

```nginx
ssl_protocols TLSv1.2 TLSv1.3;
ssl_ciphers ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256;
ssl_prefer_server_ciphers off;
add_header Strict-Transport-Security "max-age=63072000" always;
```

#### Headers de Seguridad
```nginx
# Prevención de clickjacking
add_header X-Frame-Options "SAMEORIGIN" always;

# Prevención de MIME type sniffing
add_header X-Content-Type-Options "nosniff" always;

# XSS Protection
add_header X-XSS-Protection "1; mode=block" always;

# Referrer Policy
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
```

#### Filtrado de Requests Maliciosos
```nginx
# Bloqueo de user agents sospechosos
if ($http_user_agent ~* (nmap|nikto|wikto|sf|sqlmap|bsqlbf|w3af|acunetix|havij|appscan)) {
    return 444;
}

# Bloqueo de métodos HTTP no permitidos
if ($request_method !~ ^(GET|POST|PUT|DELETE|OPTIONS)$ ) {
    return 405;
}

# Limitación de tamaño de body
client_max_body_size 10M;
```

#### DDoS Protection
- **Connection Limiting**: Límite de conexiones concurrentes por IP
- **Request Size Limiting**: Prevención de ataques de buffer overflow
- **Timeout Configuration**: Prevención de slow attacks

```nginx
limit_conn_zone $binary_remote_addr zone=conn_limit_per_ip:10m;
limit_conn conn_limit_per_ip 20;

client_body_timeout 12;
client_header_timeout 12;
send_timeout 10;
```

#### Ocultación de Información del Servidor
```nginx
# Ocultar versión de nginx
server_tokens off;

# Headers personalizados
add_header X-Powered-By "PeluDog-CRM" always;
```

### Arquitectura de Proxy Inverso
```
Internet → Nginx (Puerto 80/443) → Rails App (Puerto 3000)
```

**Beneficios**:
- Rails app no expuesta directamente a internet
- Nginx maneja SSL termination y compresión
- Filtrado de requests antes de llegar a Rails
- Servicio de archivos estáticos optimizado
- Load balancing para múltiples instancias de Rails

## Medidas de Seguridad Adicionales

### Auditoría y Logging
- **Rails Logger**: Registro de todas las acciones críticas
- **Nginx Access Logs**: Monitoreo de patrones de tráfico
- **Fail2Ban**: Bloqueo automático de IPs maliciosas

### Monitoreo de Seguridad
- **Intrusion Detection**: Alertas para patrones sospechosos
- **Log Analysis**: Análisis automatizado de logs de seguridad
- **Health Checks**: Monitoreo continuo de servicios críticos

### Backup y Recuperación
- **Encrypted Backups**: Respaldos cifrados de base de datos
- **Disaster Recovery**: Plan de recuperación ante incidentes
- **Data Retention**: Políticas de retención de datos sensibles