<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Icono de TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.2/demo.gif?raw=true" alt="Demostración de TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <strong>Español</strong> •
    <a href="README.fr.md">Français</a> •
    <a href="README.it.md">Italiano</a> •
    <a href="README.ja.md">日本語</a> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI es una aplicación de escritorio ligera que ofrece una interfaz fácil de usar para TrID, una potente herramienta para escanear y analizar archivos. Con TrID UI, puedes seleccionar o arrastrar y soltar archivos en la pantalla de inicio para iniciar escaneos locales y detectar tipos de archivo no reconocidos.

La aplicación utiliza una implementación nativa en Go del algoritmo de identificación de archivos TrID, proporcionando una detección rápida y precisa sin dependencias externas.

> [!TIP]
> Descarga TridUI desde la [página de lanzamientos](https://github.com/JMcrafter26/TridUI/releases)

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JMcrafter26/TridUI?style=flat&logo=go)](https://github.com/JMcrafter26/TridUI)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/JMcrafter26/TridUI?style=flat&label=último+release&logo=github)](https://github.com/JMcrafter26/TridUI/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/JMcrafter26/TridUI?style=flat&logo=github)](https://github.com/JMcrafter26/TridUI/issues)
[![Actions Status](https://img.shields.io/github/actions/workflow/status/JMcrafter26/TridUI/release.yml?branch=main&label=build&logo=github&style=flat)](https://github.com/JMcrafter26/TridUI/actions/workflows/release.yml)

## Características

- 🚀 Escaneo rápido basado en Go nativo
- 🎯 Identificación precisa de tipos de archivo usando definiciones de TrID
- 💻 Aplicación de escritorio multiplataforma (Windows, macOS, Linux)
- 🔒 Procesamiento 100% local: ningún dato sale de tu equipo
- 🎨 Interfaz moderna e intuitiva
- 📊 Resultados detallados con puntuaciones de confianza
- 🔄 Compatibilidad con arrastrar y soltar archivos
- 🔁 Actualización automática de definiciones con un clic
- 📅 Seguimiento de la fecha de la última actualización y número de definiciones

## Tabla de contenidos

<details>
<summary>Haz clic para expandir</summary>

- [Características](#características)
- [Tabla de contenidos](#tabla-de-contenidos)
- [Demostración y capturas](#demostración-y-capturas)
  - [Vídeos de demostración](#vídeos-de-demostración)
  - [Capturas](#capturas)
- [Instalación](#instalación)
  - [Requisitos previos](#requisitos-previos)
    - [Opción 1: Descarga automática (Recomendada)](#opción-1-descarga-automática-recomendada)
    - [Opción 2: Instalación manual](#opción-2-instalación-manual)
  - [Compilar desde el código fuente](#compilar-desde-el-código-fuente)
- [Uso](#uso)
- [Detalles técnicos](#detalles-técnicos)
  - [Arquitectura](#arquitectura)
  - [Implementación del escáner TrID](#implementación-del-escáner-trid)
- [Licencia y atribución](#licencia-y-atribución)
- [Contribuir](#contribuir)
  - [Traducciones](#traducciones)

</details>

## Demostración y capturas

### Vídeos de demostración

<details>
<summary>Haz clic para expandir</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Capturas

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="Captura de TrID UI 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="Captura de TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="Captura de TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Mostrar más</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="Captura de TrID UI 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="Captura de TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="Captura de TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="Captura de TrID UI 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="Captura de TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Instalación

> [!TIP]
> Puedes encontrar binarios precompilados en la [página de lanzamientos](https://github.com/JMcrafter26/TridUI/releases).

### Requisitos previos

¡La aplicación puede descargar y actualizar automáticamente el archivo de definiciones de TrID por ti!

#### Opción 1: Descarga automática (Recomendada)

1. Inicia TrID UI
2. Abre Configuración
3. Haz clic en "Download Definitions" o "Check for Updates"
4. La aplicación descargará e instalará automáticamente las últimas definiciones

#### Opción 2: Instalación manual

1. Descarga el archivo de definiciones de TrID (`triddefs.trd`) desde [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Coloca `triddefs.trd` en el directorio de datos de la aplicación:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

Puedes usar el botón "Open App Dir" en Configuración para ir a la ubicación correcta.

### Compilar desde el código fuente

> **📖 Documentación completa de compilación:** Consulta [`build/README.md`](../../build/README.md) para instrucciones y resolución de problemas.

**Compilación rápida:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Qué hacen los scripts de compilación:**
- ✅ Validan requisitos (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Comprueban dependencias del sistema
- ✅ Detectan herramientas opcionales (UPX, NSIS, create-dmg)
- ✅ Detectan la arquitectura automáticamente
- ✅ Crean paquetes distribuibles

**Requisitos mínimos:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Instalar Wails CLI:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Ubicaciones de salida:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## Uso

1. Inicia TrID UI
2. Haz clic o arrastra y suelta un archivo en la interfaz
3. Consulta los resultados con puntuaciones de confianza
4. La mejor coincidencia aparece resaltada arriba
5. Otras coincidencias posibles se muestran debajo

## Detalles técnicos

### Arquitectura

- **Backend**: Go (framework Wails)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Motor TrID**: Implementación pura en Go (paquete `/trid`)

### Implementación del escáner TrID

El escáner TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) es una implementación en Go desde cero que:

- Analiza archivos TRD (definiciones de TrID) usando la especificación de formato binario
- Realiza coincidencia de patrones en offsets específicos
- Soporta coincidencia de cadenas para mayor precisión
- Calcula puntuaciones de confianza basadas en pesos de patrones
- Devuelve resultados ordenados con información detallada del tipo de archivo

> Puedes encontrar la especificación del formato TRD en [Mark0.net](https://mark0.net/soft-trid-format.html).

## Licencia y atribución

TrID UI es software de código abierto bajo la licencia GNU AGPLv3. La interfaz está desarrollada por Cufiy (aka JMcrafter26) y se basa en TrID de [Marco Pontello](https://mark0.net/).
Consulta el archivo LICENSE para más detalles.

El escáner `trid.go` es una implementación propia en Go por JMcrafter26 y está licenciado bajo GNU AGPLv3.

El ícono de la app se basa en el icono de ojo de icons8.com.

<a href="https://github.com/JMcrafter26/TridUI/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=JMcrafter26/TridUI" />
</a>

## Contribuir

¡Se agradecen las contribuciones! Si quieres contribuir a TrID UI, haz un fork del repositorio y envía un pull request con tus cambios. Para cambios de mayor envergadura, abre primero un issue para discutir la propuesta.

### Traducciones

¡TrID UI necesita tu ayuda para llegar a más personas! Las traducciones actuales están generadas por máquina y pueden contener imprecisiones.

Si quieres contribuir con traducciones, sigue estos pasos:

1. Haz fork del repositorio
2. Crea una rama nueva para tu traducción
3. Añade tus archivos de traducción en el directorio `translations`
4. Envía un pull request con tus cambios

¡Gracias por ayudar a hacer TrID UI accesible para más usuarios!
