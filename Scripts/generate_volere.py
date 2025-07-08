import json
import os
import re
from pathlib import Path

BASE_DIR = Path(__file__).resolve().parent.parent
JSON_PATH = BASE_DIR / "JsonRequisitos" / "requisitosTotales.json"
OUTPUT_DIR = BASE_DIR / "plantillasVolere"
OUTPUT_DIR.mkdir(exist_ok=True)

with JSON_PATH.open(encoding="utf-8") as f:
    data = json.load(f)

for rf in data.get("funcionales", []):
    safe_name = re.sub(r"[^a-zA-Z0-9_-]+", "_", rf["nombre"].lower())
    filename = OUTPUT_DIR / f"{rf['id']}_{safe_name}.md"
    criterios = "\n".join(f"- {c}" for c in rf.get("criteriosAceptacion", []))

    content = f"""# Volere Plantilla: {rf['nombre']}

**ID:** {rf['id']}

**Módulo/Área:** {rf['modulo']}

**Descripción:** {rf['descripcion']}

**Prioridad:** {rf['prioridad']}

## Criterios de Aceptación
{criterios}

## Fuente
Dra. Génesis / Equipo PeluDog

## Razón/Valor de negocio
Describe cómo este requisito aporta valor al consultorio.
"""

    filename.write_text(content, encoding="utf-8")

print(f"Se generaron {len(data.get('funcionales', []))} plantillas en {OUTPUT_DIR}")
