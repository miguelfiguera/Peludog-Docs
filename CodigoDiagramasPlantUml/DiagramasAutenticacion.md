### **Códigos PlantUML para Módulo de Autenticación - PeluDog**

A continuación se presentan los códigos PlantUML para generar todos los diagramas relacionados con el módulo de autenticación de usuarios.

---

### **1. Diagrama de Caso de Uso de Autenticación**

Este diagrama muestra el caso de uso de autenticación y sus relaciones con todos los usuarios del sistema.

```plantuml
@startuml
left to right direction

' --- Estilos para mejorar la apariencia ---
<style>
title {
  FontSize 20
  FontColor #333333
}
usecase {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
  roundCorner 15
  shadowing true
}
actor {
  BackgroundColor #FFFFFF
  BorderColor #1565C0
  FontColor #1A237E
}
</style>

title Caso de Uso: Autenticación y Gestión de Usuarios

actor "Asistente" as Asistente
actor "Veterinario" as Veterinario
actor "Cliente/Propietario" as Cliente
actor "Administrador/Gerente" as Admin

rectangle "Módulo de Autenticación" {
  usecase "Iniciar Sesión" as UC_Login
  usecase "Cerrar Sesión" as UC_Logout
  usecase "Recuperar Contraseña" as UC_Recovery
  usecase "Gestionar Políticas\nde Contraseña" as UC_Policies

  UC_Login .> UC_Policies : <<include>>
  UC_Recovery .> UC_Policies : <<include>>
}

' --- Relaciones con todos los usuarios ---
Asistente -- UC_Login
Asistente -- UC_Logout
Asistente -- UC_Recovery

Veterinario -- UC_Login
Veterinario -- UC_Logout
Veterinario -- UC_Recovery

Cliente -- UC_Login
Cliente -- UC_Logout
Cliente -- UC_Recovery

Admin -- UC_Login
Admin -- UC_Logout
Admin -- UC_Recovery

note right of UC_Policies
  Las políticas de contraseña se aplican
  automáticamente en todos los procesos
  de autenticación para garantizar
  la seguridad del sistema.
end note

@enduml
```

---

### **2. Diagrama de Secuencia Principal - CU-AU01**

Este diagrama muestra la secuencia completa del caso de uso de autenticación, incluyendo inicio de sesión, gestión de sesión y cierre de sesión.

```plantuml
@startuml
' --- Estilos para mejorar la legibilidad ---
<style>
title {
  FontSize 18
  FontColor #333333
}
participant {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
}
activate {
  BackgroundColor #BBDEFB
  BorderColor #1565C0
}
</style>

title Diagrama de Secuencia: CU-AU01 - Autenticar Usuario

actor "Usuario" as Usuario
participant "Sistema de Login" as Login
participant "Validador de Credenciales" as Validador
participant "Gestor de Sesiones" as Sesiones
participant "Gestor de Políticas" as Politicas
participant "Servicio de Email" as Email
database "Base de Datos" as BD

== Iniciar Sesión ==
Usuario -> Login: Accede a página de login
activate Login
Login --> Usuario: Muestra formulario de login
Usuario -> Login: Ingresa credenciales (usuario, contraseña)
Login -> Validador: Valida credenciales
activate Validador

Validador -> BD: Consulta usuario y contraseña
activate BD
BD --> Validador: Retorna datos del usuario
deactivate BD

Validador -> Politicas: Verifica políticas de contraseña
activate Politicas
Politicas --> Validador: Valida estado de contraseña
deactivate Politicas

alt Credenciales válidas
    Validador --> Login: Credenciales correctas
    Login -> Sesiones: Crear sesión para usuario
    activate Sesiones
    Sesiones -> BD: Registra sesión activa
    BD --> Sesiones: Sesión creada
    Sesiones --> Login: Token de sesión
    deactivate Sesiones
    Login --> Usuario: Redirige a dashboard principal
else Credenciales inválidas
    Validador -> BD: Registra intento fallido
    Validador --> Login: Credenciales incorrectas
    Login --> Usuario: Muestra error y aplica políticas de bloqueo
end

deactivate Validador
deactivate Login

== Gestión de Sesión Activa ==
note over Usuario, BD
  Usuario interactúa con el sistema
  La sesión se mantiene activa
end note

== Cerrar Sesión ==
Usuario -> Login: Selecciona "Cerrar Sesión"
activate Login
Login -> Sesiones: Invalidar sesión
activate Sesiones
Sesiones -> BD: Elimina sesión activa
BD --> Sesiones: Sesión eliminada
Sesiones --> Login: Sesión invalidada
deactivate Sesiones
Login --> Usuario: Redirige a página de login
deactivate Login

== Recuperación de Contraseña ==
Usuario -> Login: Selecciona "Recuperar Contraseña"
activate Login
Login --> Usuario: Solicita email registrado
Usuario -> Login: Ingresa email
Login -> BD: Verifica si email existe
activate BD
BD --> Login: Confirma usuario existente
deactivate BD

Login -> Politicas: Genera token temporal
activate Politicas
Politicas --> Login: Token seguro generado
deactivate Politicas

Login -> Email: Envía enlace de recuperación
activate Email
Email --> Login: Email enviado
deactivate Email

Login --> Usuario: Confirma envío de email
deactivate Login

Usuario -> Login: Accede a enlace de recuperación
activate Login
Login -> Politicas: Valida token temporal
activate Politicas
Politicas --> Login: Token válido
deactivate Politicas

Login --> Usuario: Muestra formulario de nueva contraseña
Usuario -> Login: Ingresa nueva contraseña
Login -> Politicas: Valida nueva contraseña
activate Politicas
Politicas --> Login: Contraseña cumple políticas
deactivate Politicas

Login -> BD: Actualiza contraseña
activate BD
BD --> Login: Contraseña actualizada
deactivate BD

Login --> Usuario: Confirma cambio exitoso
deactivate Login

@enduml
```

---

### **3. Flujo Alternativo FA-AU01-01: Credenciales Incorrectas**

Este diagrama muestra el flujo cuando el usuario ingresa credenciales incorrectas y las políticas de bloqueo.

```plantuml
@startuml
<style>
title {
  FontSize 16
  FontColor #333333
}
participant {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
}
activate {
  BackgroundColor #FFCDD2
  BorderColor #D32F2F
}
</style>

title Flujo Alternativo FA-AU01-01: Credenciales Incorrectas

actor "Usuario" as Usuario
participant "Sistema de Login" as Login
participant "Validador de Credenciales" as Validador
participant "Gestor de Políticas" as Politicas
database "Base de Datos" as BD

Usuario -> Login: Ingresa credenciales incorrectas
activate Login
Login -> Validador: Valida credenciales
activate Validador

Validador -> BD: Consulta usuario
activate BD
BD --> Validador: Usuario encontrado
deactivate BD

Validador -> Validador: Compara contraseñas
note right: Contraseña no coincide

Validador -> Politicas: Registra intento fallido
activate Politicas
Politicas -> BD: Incrementa contador de intentos fallidos
activate BD
BD --> Politicas: Contador actualizado
deactivate BD

Politicas -> Politicas: Verifica si supera límite (ej. 3 intentos)

alt Intentos dentro del límite
    Politicas --> Validador: Continúa permitiendo intentos
    Validador --> Login: Credenciales incorrectas
    Login --> Usuario: "Usuario o contraseña incorrectos. Intentos restantes: X"
else Límite de intentos superado
    Politicas -> BD: Bloquea cuenta temporalmente
    activate BD
    BD --> Politicas: Cuenta bloqueada
    deactivate BD
    Politicas --> Validador: Cuenta bloqueada
    Validador --> Login: Cuenta bloqueada
    Login --> Usuario: "Cuenta bloqueada temporalmente. Contacte al administrador."
end

deactivate Politicas
deactivate Validador
deactivate Login

@enduml
```

---

### **4. Flujo Alternativo FA-AU01-02: Cuenta Bloqueada**

Este diagrama muestra el comportamiento cuando un usuario intenta acceder con una cuenta ya bloqueada.

```plantuml
@startuml
<style>
title {
  FontSize 16
  FontColor #333333
}
participant {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
}
activate {
  BackgroundColor #FFCDD2
  BorderColor #D32F2F
}
</style>

title Flujo Alternativo FA-AU01-02: Cuenta Bloqueada

actor "Usuario" as Usuario
participant "Sistema de Login" as Login
participant "Validador de Credenciales" as Validador
participant "Gestor de Políticas" as Politicas
database "Base de Datos" as BD

Usuario -> Login: Intenta iniciar sesión
activate Login
Login -> Validador: Valida credenciales
activate Validador

Validador -> BD: Consulta estado de la cuenta
activate BD
BD --> Validador: Cuenta bloqueada (timestamp del bloqueo)
deactivate BD

Validador -> Politicas: Verifica tiempo de bloqueo
activate Politicas
Politicas -> Politicas: Calcula tiempo transcurrido desde bloqueo

alt Tiempo de bloqueo no ha expirado
    Politicas --> Validador: Cuenta aún bloqueada
    Validador --> Login: Acceso denegado - cuenta bloqueada
    Login --> Usuario: "Su cuenta está bloqueada. Tiempo restante: X minutos."
else Tiempo de bloqueo ha expirado
    Politicas -> BD: Desbloquea cuenta automáticamente
    activate BD
    BD --> Politicas: Cuenta desbloqueada
    deactivate BD
    Politicas -> BD: Resetea contador de intentos fallidos
    BD --> Politicas: Contador reseteado
    Politicas --> Validador: Cuenta desbloqueada - proceder con validación
    Validador --> Login: Continúa con proceso normal de login
    Login --> Usuario: Procede con autenticación normal
end

deactivate Politicas
deactivate Validador
deactivate Login

@enduml
```

---

### **5. Flujo Alternativo FA-AU01-03: Contraseña Expirada**

Este diagrama muestra el flujo cuando la contraseña del usuario ha expirado según las políticas establecidas.

```plantuml
@startuml
<style>
title {
  FontSize 16
  FontColor #333333
}
participant {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
}
activate {
  BackgroundColor #FFF9C4
  BorderColor #F57F17
}
</style>

title Flujo Alternativo FA-AU01-03: Contraseña Expirada

actor "Usuario" as Usuario
participant "Sistema de Login" as Login
participant "Validador de Credenciales" as Validador
participant "Gestor de Políticas" as Politicas
database "Base de Datos" as BD

Usuario -> Login: Ingresa credenciales válidas
activate Login
Login -> Validador: Valida credenciales
activate Validador

Validador -> BD: Consulta usuario y fecha de última actualización de contraseña
activate BD
BD --> Validador: Credenciales válidas + fecha de actualización
deactivate BD

Validador -> Politicas: Verifica expiración de contraseña
activate Politicas
Politicas -> Politicas: Calcula días desde última actualización
Politicas -> Politicas: Compara con política de expiración (ej. 90 días)

alt Contraseña expirada
    Politicas --> Validador: Contraseña expirada
    Validador --> Login: Requiere cambio de contraseña
    Login --> Usuario: "Su contraseña ha expirado. Debe cambiarla para continuar."

    Login --> Usuario: Muestra formulario de cambio de contraseña
    Usuario -> Login: Ingresa contraseña actual y nueva contraseña
    Login -> Politicas: Valida nueva contraseña
    Politicas -> Politicas: Verifica que cumple políticas de seguridad
    Politicas -> Politicas: Verifica que es diferente a la anterior

    alt Nueva contraseña válida
        Politicas --> Login: Nueva contraseña aceptada
        Login -> BD: Actualiza contraseña y fecha de actualización
        activate BD
        BD --> Login: Contraseña actualizada
        deactivate BD
        Login --> Usuario: "Contraseña actualizada exitosamente"
        Login --> Usuario: Continúa con proceso de login normal
    else Nueva contraseña inválida
        Politicas --> Login: Nueva contraseña no cumple políticas
        Login --> Usuario: "La nueva contraseña no cumple con los requisitos de seguridad"
        Login --> Usuario: Muestra formulario nuevamente
    end
else Contraseña vigente
    Politicas --> Validador: Contraseña vigente
    Validador --> Login: Procede con login normal
    Login --> Usuario: Redirige a dashboard principal
end

deactivate Politicas
deactivate Validador
deactivate Login

@enduml
```

---

### **6. Flujo Alternativo FA-AU01-04: Enlace de Recuperación Expirado**

Este diagrama muestra el comportamiento cuando el usuario intenta usar un enlace de recuperación que ya ha expirado.

```plantuml
@startuml
<style>
title {
  FontSize 16
  FontColor #333333
}
participant {
  BackgroundColor #E3F2FD
  BorderColor #1565C0
  FontColor #1A237E
}
activate {
  BackgroundColor #FFCDD2
  BorderColor #D32F2F
}
</style>

title Flujo Alternativo FA-AU01-04: Enlace de Recuperación Expirado

actor "Usuario" as Usuario
participant "Sistema de Login" as Login
participant "Gestor de Políticas" as Politicas
database "Base de Datos" as BD

Usuario -> Login: Accede a enlace de recuperación
activate Login
Login -> Politicas: Valida token de recuperación
activate Politicas

Politicas -> BD: Consulta token y fecha de expiración
activate BD
BD --> Politicas: Token encontrado con timestamp
deactivate BD

Politicas -> Politicas: Verifica si token ha expirado (ej. > 24 horas)

alt Token expirado
    Politicas -> BD: Invalida token expirado
    activate BD
    BD --> Politicas: Token eliminado
    deactivate BD
    Politicas --> Login: Token expirado
    Login --> Usuario: "El enlace de recuperación ha expirado."

    Login --> Usuario: Muestra opción "Solicitar nuevo enlace"
    Usuario -> Login: Solicita nuevo enlace de recuperación
    Login --> Usuario: Redirige a formulario de recuperación
    note right: Inicia proceso de recuperación nuevamente

else Token vigente
    Politicas --> Login: Token válido
    Login --> Usuario: Muestra formulario de nueva contraseña
    note right: Continúa con proceso normal de recuperación
end

deactivate Politicas
deactivate Login

@enduml
```
