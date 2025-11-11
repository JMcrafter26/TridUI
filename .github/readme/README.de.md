<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="TrID UI Icon" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>

</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.2/demo.gif?raw=true" alt="TrID UI Demonstration" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <strong>Deutsch</strong> •
    <a href="README.es.md">Español</a> •
    <a href="README.fr.md">Français</a> •
    <a href="README.it.md">Italiano</a> •
    <a href="README.ja.md">日本語</a> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI ist eine leichtgewichtige Desktop-Anwendung mit einer benutzerfreundlichen Oberfläche für TrID – ein leistungsstarkes Tool zum Scannen und Analysieren von Dateien. Mit TrID UI können Nutzer Dateien bequem auf dem Startbildschirm auswählen oder per Drag & Drop ablegen, um lokale Scans zu starten und unbekannte Dateitypen zu erkennen.

Die Anwendung nutzt eine native Go-Implementierung des TrID-Dateierkennungsalgorithmus und ermöglicht schnelle, präzise Erkennung ohne externe Abhängigkeiten.

> [!TIP]
> Lade TridUI von der [Releases-Seite](https://github.com/JMcrafter26/TridUI/releases) herunter

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JMcrafter26/TridUI?style=flat&logo=go)](https://github.com/JMcrafter26/TridUI)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/JMcrafter26/TridUI?style=flat&label=letztes+Release&logo=github)](https://github.com/JMcrafter26/TridUI/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/JMcrafter26/TridUI?style=flat&logo=github)](https://github.com/JMcrafter26/TridUI/issues)
[![Actions Status](https://img.shields.io/github/actions/workflow/status/JMcrafter26/TridUI/release.yml?branch=main&label=Build&logo=github&style=flat)](https://github.com/JMcrafter26/TridUI/actions/workflows/release.yml)

## Funktionen

- 🚀 Schnelles Scannen mit nativer Go-Implementierung
- 🎯 Präzise Dateityperkennung anhand von TrID-Definitionen
- 💻 Plattformübergreifende Desktop-App (Windows, macOS, Linux)
- 🔒 100% lokale Verarbeitung – keine Daten verlassen deinen Rechner
- 🎨 Moderne, intuitive Benutzeroberfläche
- 📊 Detailierte Treffer mit Vertrauenswerten
- 🔄 Drag-and-drop für Dateien
- 🔁 Automatische Definitions-Updates mit einem Klick
- 📅 Letztes Update-Datum und Definitionsanzahl im Blick

## Inhaltsverzeichnis

<details>
<summary>Zum Aufklappen klicken</summary>

- [Funktionen](#funktionen)
- [Inhaltsverzeichnis](#inhaltsverzeichnis)
- [Demo und Screenshots](#demo-und-screenshots)
  - [Demovideos](#demovideos)
  - [Screenshots](#screenshots)
- [Einrichtung](#einrichtung)
  - [Voraussetzungen](#voraussetzungen)
    - [Option 1: Automatischer Download (Empfohlen)](#option-1-automatischer-download-empfohlen)
    - [Option 2: Manuelle Installation](#option-2-manuelle-installation)
  - [Aus dem Quellcode bauen](#aus-dem-quellcode-bauen)
- [Benutzung](#benutzung)
- [Technische Details](#technische-details)
  - [Architektur](#architektur)
  - [TrID-Scanner-Implementierung](#trid-scanner-implementierung)
- [Lizenz und Danksagung](#lizenz-und-danksagung)
- [Mitwirken](#mitwirken)
  - [Übersetzungen](#übersetzungen)

</details>

## Demo und Screenshots

### Demovideos

<details>
<summary>Zum Aufklappen klicken</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Screenshots

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="TrID UI Screenshot 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="TrID UI Screenshot 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="TrID UI Screenshot 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Mehr anzeigen</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="TrID UI Screenshot 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="TrID UI Screenshot 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="TrID UI Screenshot 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="TrID UI Screenshot 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="TrID UI Screenshot 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Einrichtung

> [!TIP]
> Vorgebaute Binärdateien findest du auf der [Releases-Seite](https://github.com/JMcrafter26/TridUI/releases).

### Voraussetzungen

Die Anwendung kann die TrID-Definitionsdatei automatisch für dich herunterladen und aktualisieren!

#### Option 1: Automatischer Download (Empfohlen)

1. TrID UI starten
2. Zu Einstellungen wechseln
3. Auf „Definitionen herunterladen“ oder „Nach Updates suchen“ klicken
4. Die App lädt und installiert die neuesten Definitionen automatisch

#### Option 2: Manuelle Installation

1. Lade die TrID-Definitionsdatei (`triddefs.trd`) von [Mark0.net](https://mark0.net/soft-trid-deflist.html) herunter
2. Lege die Datei `triddefs.trd` im Anwendungsdatenordner ab:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

Über die Schaltfläche „App-Ordner öffnen“ in den Einstellungen gelangst du direkt zum richtigen Speicherort.

### Aus dem Quellcode bauen

> **📖 Ausführliche Build-Doku:** Siehe [`build/README.md`](../../build/README.md) für Details und Troubleshooting.

**Schneller Build:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Was die Build-Skripte machen:**
- ✅ Voraussetzungen prüfen (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Systemabhängigkeiten prüfen
- ✅ Optionale Tools erkennen (UPX, NSIS, create-dmg)
- ✅ Architektur automatisch erkennen
- ✅ Distributable-Pakete erzeugen

**Mindestanforderungen:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Wails CLI installieren:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Ausgabeorte:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## Benutzung

1. TrID UI starten
2. Datei anklicken oder per Drag & Drop in die Oberfläche ziehen
3. Scanergebnisse samt Vertrauenswerten ansehen
4. Der beste Treffer ist oben hervorgehoben
5. Weitere mögliche Treffer sind darunter gelistet

## Technische Details

### Architektur

- **Backend**: Go (Wails Framework)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **TrID-Engine**: Reine Go-Implementierung (Paket `/trid`)

### TrID-Scanner-Implementierung

Der TrID-Scanner ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) ist eine Clean-Room-Go-Implementierung, die:

- TRD-(TrID-Definition)-Dateien gemäß Binärformatspezifikation parst
- Pattern-Matching an spezifizierten Datei-Offsets durchführt
- String-Matching für höhere Genauigkeit unterstützt
- Vertrauenswertberechnung basierend auf Pattern-Gewichtungen vornimmt
- Gerankte Resultate mit detaillierten Dateityp-Informationen zurückgibt

> Die TRD-Formatspezifikation findest du auf [Mark0.net](https://mark0.net/soft-trid-format.html).

## Lizenz und Danksagung

TrID UI ist Open-Source-Software unter der GNU AGPLv3. Die UI wurde von Cufiy (aka JMcrafter26) entwickelt und basiert auf TrID von [Marco Pontello](https://mark0.net/).
Details stehen in der LICENSE-Datei.

Der Scanner `trid.go` ist eine Clean-Room-Go-Implementierung von JMcrafter26 und steht unter der GNU AGPLv3.

Das App-Icon basiert auf dem „eye“-Icon von icons8.com.

<a href="https://github.com/JMcrafter26/TridUI/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=JMcrafter26/TridUI" />
</a>

## Mitwirken

Beiträge sind willkommen! Wenn du zu TrID UI beitragen möchtest, forke das Repository und stelle einen Pull-Request mit deinen Änderungen. Für größere Änderungen eröffne bitte vorher ein Issue zur Abstimmung.

### Übersetzungen

TrID UI braucht deine Hilfe, um mehr Nutzer zu erreichen! Die aktuellen Übersetzungen sind maschinell erzeugt und können Ungenauigkeiten enthalten.

Wenn du Übersetzungen beitragen möchtest, gehe so vor:

1. Repository forken
2. Einen neuen Branch für deine Übersetzung anlegen
3. Deine Übersetzungsdateien im Verzeichnis `translations` hinzufügen
4. Einen Pull-Request mit deinen Änderungen erstellen

Vielen Dank, dass du hilfst, TrID UI für mehr Nutzer zugänglich zu machen!
