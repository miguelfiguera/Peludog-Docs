<style>
@page {
    size: A4;
    margin: 2cm;
}

body {
    font-family: Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    margin: 0;
    padding: 20px;
    background: white;
}

.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 40px;
    min-height: 80vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.header {
    text-align: center;
    border-bottom: 3px solid #2c3e50;
    padding-bottom: 20px;
    margin-bottom: 40px;
}

.volere-title {
    font-size: 28px;
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 10px;
}

.requisito-id {
    font-size: 20px;
    font-weight: bold;
    color: #e74c3c;
    background: #f8f9fa;
    padding: 8px 16px;
    border-radius: 4px;
    display: inline-block;
}

.section {
    margin: 25px 0;
    text-align: left;
}

.section-title {
    font-size: 18px;
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 10px;
    border-left: 4px solid #3498db;
    padding-left: 15px;
}

.section-content {
    font-size: 14px;
    padding: 10px 0;
    text-align: justify;
}

.metadata {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
    margin: 20px 0;
}

.metadata-item {
    background: #f8f9fa;
    padding: 15px;
    border-radius: 6px;
    border-left: 4px solid #3498db;
}

.metadata-label {
    font-weight: bold;
    color: #2c3e50;
    margin-bottom: 5px;
}

.criterios {
    background: #f8f9fa;
    padding: 20px;
    border-radius: 6px;
    border-left: 4px solid #27ae60;
}

.criterios ul {
    margin: 10px 0;
    padding-left: 20px;
}

.criterios li {
    margin-bottom: 8px;
    text-align: justify;
}

.footer {
    text-align: center;
    margin-top: 40px;
    padding-top: 20px;
    border-top: 1px solid #bdc3c7;
    font-size: 12px;
    color: #7f8c8d;
}

@media print {
    body {
        print-color-adjust: exact;
    }
    .container {
        page-break-inside: avoid;
    }
}
</style>

<div class="container">

<div class="header">

# Plantilla Volere

## RF-RH004

</div>

## Requisito

**Desactivar/Activar Empleado**

### Metadatos

| Campo | Valor |
|-------|-------|
| **Módulo/Área** | Gestión de Recursos Humanos y Servicios Internos |
| **Prioridad** | Media - Permite gestionar altas/bajas de empleados asegurando control de acceso; CU-RH01. |
| **Caso de Uso** | CU-RH01 |


## Descripción

El Administrador/Gerente podrá desactivar temporal o permanentemente la cuenta de un empleado, impidiendo su acceso al sistema, o reactivarla.

## Criterios de Aceptación

- El sistema debe permitir cambiar el estado de un empleado a 'activo' o 'inactivo'.
- Los empleados inactivos no deben poder iniciar sesión.
- El historial de acciones del empleado debe conservarse.


## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="footer">

*Plantilla Volere - Proyecto Socio Tecnologico: Plataforma CRM para PeluDog*

</div>

</div>