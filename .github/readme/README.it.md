<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Icona TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/demo.gif?raw=true" alt="Dimostrazione TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <a href="README.es.md">Español</a> •
    <a href="README.fr.md">Français</a> •
    <strong>Italiano</strong> •
    <a href="README.ja.md">日本語</a> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI è un’app desktop leggera che offre un’interfaccia intuitiva per TrID, un potente strumento per la scansione e l’analisi dei file. Con TrID UI puoi selezionare o trascinare i file sulla schermata Home per avviare scansioni locali e rilevare tipi di file non riconosciuti.

L’app utilizza un’implementazione nativa in Go dell’algoritmo di identificazione file TrID, offrendo rilevazioni rapide e accurate senza dipendenze esterne.

> [!TIP]
> Scarica TridUI dalla [pagina delle release](https://github.com/JMcrafter26/TridUI/releases)

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JMcrafter26/TridUI?style=flat&logo=go)](https://github.com/JMcrafter26/TridUI)
[![Release GitHub (ultima)](https://img.shields.io/github/v/release/JMcrafter26/TridUI?style=flat&label=ultima+release&logo=github)](https://github.com/JMcrafter26/TridUI/releases/latest)
[![Issue GitHub](https://img.shields.io/github/issues/JMcrafter26/TridUI?style=flat&logo=github)](https://github.com/JMcrafter26/TridUI/issues)
[![Actions Status](https://img.shields.io/github/actions/workflow/status/JMcrafter26/TridUI/release.yml?branch=main&label=build&logo=github&style=flat)](https://github.com/JMcrafter26/TridUI/actions/workflows/release.yml)

<div>
<a href="https://fmhy.net/file-tools#file-info-metadata:~:text=TridUI" target="_blank" rel="noopener noreferrer">
<img alt="badge name" src="https://cdn.jsdelivr.net/gh/JMcrafter26/badges@main/src/assets/available/fmhy/cozy.svg">
</a>
<a href="https://wails.io" target="_blank" rel="noopener noreferrer">
  <img alt="badge name" src="https://cdn.jsdelivr.net/gh/JMcrafter26/badges@main/src/assets/built-with/wails/cozy.svg">
</a>
<a href="https://svelte.dev" target="_blank" rel="noopener noreferrer">
  <img alt="badge name" src="https://cdn.jsdelivr.net/gh/JMcrafter26/badges@main/src/assets/built-with/svelte/cozy-minimal.svg">
</a>
<a href="https://golang.org" target="_blank" rel="noopener noreferrer">
  <img alt="badge name" src="https://cdn.jsdelivr.net/gh/JMcrafter26/badges@main/src/assets/built-with/go/cozy-minimal.svg">
</a>
<a href="https://daisyui.com" target="_blank" rel="noopener noreferrer">
  <img alt="badge name" src="https://cdn.jsdelivr.net/gh/JMcrafter26/badges@main/src/assets/built-with/daisyui/cozy-minimal.svg">
</a>
</div>

## Funzionalità

- 🚀 Scansione rapida basata su Go nativo
- 🎯 Identificazione precisa dei tipi di file tramite definizioni TrID
- 💻 Applicazione desktop multipiattaforma (Windows, macOS, Linux)
- 🔒 Elaborazione 100% locale – nessun dato lascia il tuo computer
- 🎨 Interfaccia moderna e intuitiva
- 📊 Risultati dettagliati con punteggi di confidenza
- 🔄 Supporto drag‑and‑drop dei file
- 🔁 Aggiornamenti automatici delle definizioni con un clic
- 🖥️ Interfaccia a riga di comando (CLI) per utenti avanzati

## Indice

<details>
<summary>Fai clic per espandere</summary>

- [Funzionalità](#funzionalità)
- [Indice](#indice)
- [Dimostrazioni e schermate](#dimostrazioni-e-schermate)
  - [Video dimostrativi](#video-dimostrativi)
  - [Schermate](#schermate)
- [Configurazione](#configurazione)
  - [Prerequisiti](#prerequisiti)
    - [Opzione 1: Download automatico (Consigliata)](#opzione-1-download-automatico-consigliata)
    - [Opzione 2: Installazione manuale](#opzione-2-installazione-manuale)
  - [Compilazione dai sorgenti](#compilazione-dai-sorgenti)
- [Utilizzo](#utilizzo)
- [Dettagli tecnici](#dettagli-tecnici)
  - [Architettura](#architettura)
  - [Implementazione dello scanner TrID](#implementazione-dello-scanner-trid)
- [Licenza e attribuzioni](#licenza-e-attribuzioni)
- [Contribuire](#contribuire)
  - [Traduzioni](#traduzioni)

</details>

## Dimostrazioni e schermate

### Video dimostrativi

<details>
<summary>Fai clic per espandere</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Schermate

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/home.png?raw=true" alt="Screenshot TrID UI 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/scan.png?raw=true" alt="Screenshot TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/about.png?raw=true" alt="Screenshot TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Mostra di più</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings.png?raw=true" alt="Screenshot TrID UI 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings2.png?raw=true" alt="Screenshot TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings3.png?raw=true" alt="Screenshot TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/unknown.png?raw=true" alt="Screenshot TrID UI 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/about.png?raw=true" alt="Screenshot TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Configurazione

> [!TIP]
> Trovi i binari precompilati nella [pagina delle release](https://github.com/JMcrafter26/TridUI/releases).

### Prerequisiti

L’app può scaricare e aggiornare automaticamente il file delle definizioni TrID!

#### Opzione 1: Download automatico (Consigliata)

1. Avvia TrID UI
2. Vai su Impostazioni
3. Clicca "Download Definitions" o "Check for Updates"
4. L’app scaricherà e installerà automaticamente le ultime definizioni

#### Opzione 2: Installazione manuale

1. Scarica il file delle definizioni TrID (`triddefs.trd`) da [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Posiziona `triddefs.trd` nella directory dati dell’applicazione:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

Puoi usare il pulsante "Open App Dir" nelle Impostazioni per aprire la posizione corretta.

### Compilazione dai sorgenti

> **📖 Documentazione completa build:** vedi [`build/README.md`](../../build/README.md) per istruzioni dettagliate e troubleshooting.

**Build rapida:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Cosa fanno gli script:**
- ✅ Verificano i prerequisiti (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Controllano le dipendenze di sistema
- ✅ Rilevano strumenti opzionali (UPX, NSIS, create-dmg)
- ✅ Rilevano l’architettura automaticamente
- ✅ Creano pacchetti distribuibili

**Requisiti minimi:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Installa Wails CLI:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Posizioni output:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## Utilizzo

1. Avvia TrID UI
2. Clicca o trascina un file sull’interfaccia
3. Visualizza i risultati con punteggi di confidenza
4. La migliore corrispondenza è evidenziata in alto
5. Le possibili corrispondenze aggiuntive sono elencate sotto

## Dettagli tecnici

### Architettura

- **Backend**: Go (framework Wails)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Motore TrID**: Implementazione in Go puro (pacchetto `/trid`)

### Implementazione dello scanner TrID

Lo scanner TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) è un’implementazione "clean‑room" in Go che:

- Analizza i file TRD (definizioni TrID) secondo la specifica binaria
- Esegue pattern matching a determinati offset
- Supporta il matching di stringhe per maggiore accuratezza
- Calcola i punteggi di confidenza basati sui pesi dei pattern
- Restituisce risultati ordinati con dettagli sul tipo di file

> Trovi la specifica del formato TRD su [Mark0.net](https://mark0.net/soft-trid-format.html).

## Licenza e attribuzioni

TrID UI è software open‑source con licenza GNU AGPLv3. La UI è sviluppata da Cufiy (alias JMcrafter26) e si basa su TrID di [Marco Pontello](https://mark0.net/).
Consulta il file LICENSE per maggiori dettagli.

Lo scanner `trid.go` è una implementazione clean‑room in Go di JMcrafter26 con licenza GNU AGPLv3.

L’icona dell’app è basata sull’icona "eye" di icons8.com.

<a href="https://github.com/JMcrafter26/TridUI/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=JMcrafter26/TridUI" />
</a>

## Contribuire

Contributi ben accetti! Se vuoi contribuire a TrID UI, fai un fork del repository e invia una pull request. Per modifiche importanti, apri prima una issue per discuterne.

### Traduzioni

TrID UI ha bisogno del tuo aiuto per raggiungere più persone! Le traduzioni attuali sono generate automaticamente e potrebbero contenere inesattezze.

Per contribuire alle traduzioni:

1. Fai il fork del repository
2. Crea un nuovo branch per la traduzione
3. Aggiungi i file nella cartella `translations`
4. Invia una pull request

Grazie per aiutare a rendere TrID UI accessibile a più utenti!
