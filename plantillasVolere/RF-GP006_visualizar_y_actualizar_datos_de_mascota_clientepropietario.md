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

## RF-GP006

</div>

## Requisito

**Visualizar y Actualizar Datos de Mascota (Cliente/Propietario)**

### Metadatos

| Campo | Valor |
|-------|-------|
| **Módulo/Área** | Gestión de Pacientes (Clientes y Mascotas) |
| **Prioridad** | Alta - Empodera al cliente para gestionar los datos de sus mascotas, reduciendo carga administrativa; parte de CU-GP01. |
| **Caso de Uso** | CU-GP01 |


## Descripción

El Cliente/Propietario de Mascota podrá ver y modificar la información de sus propias mascotas registradas.

## Criterios de Aceptación

- Tras seleccionar una de sus mascotas, el sistema debe mostrar toda su información general registrada (no la historia clínica detallada).
- El sistema debe permitir la edición de los campos de la mascota por parte de su dueño.
- El sistema debe guardar los cambios realizados en la base de datos.


## Fuente

Dra. Génesis / Equipo PeluDog

## Razón/Valor de Negocio

Este requisito contribuye a la digitalización y eficiencia operativa del consultorio veterinario PeluDog, mejorando la experiencia del cliente y la gestión interna de procesos.

---

<div class="footer">

*Plantilla Volere - Proyecto Socio Tecnologico: Plataforma CRM para PeluDog*

</div>

</div>