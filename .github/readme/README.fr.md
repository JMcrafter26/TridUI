<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Icône TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.2/demo.gif?raw=true" alt="Démonstration TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <a href="README.es.md">Español</a> •
    <strong>Français</strong> •
    <a href="README.it.md">Italiano</a> •
    <a href="README.ja.md">日本語</a> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI est une application de bureau légère qui fournit une interface conviviale pour TrID, un outil puissant pour analyser les fichiers. Avec TrID UI, il suffit de sélectionner ou glisser-déposer des fichiers sur l’écran d’accueil pour lancer des analyses locales et détecter des types de fichiers inconnus.

L’application utilise une implémentation Go native de l’algorithme d’identification de fichiers TrID, offrant une détection rapide et précise sans dépendances externes.

> [!TIP]
> Téléchargez TridUI depuis la [page des releases](https://github.com/JMcrafter26/TridUI/releases)

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JMcrafter26/TridUI?style=flat&logo=go)](https://github.com/JMcrafter26/TridUI)
[![GitHub release (dernière)](https://img.shields.io/github/v/release/JMcrafter26/TridUI?style=flat&label=release+récente&logo=github)](https://github.com/JMcrafter26/TridUI/releases/latest)
[![Issues GitHub](https://img.shields.io/github/issues/JMcrafter26/TridUI?style=flat&logo=github)](https://github.com/JMcrafter26/TridUI/issues)
[![Actions Status](https://img.shields.io/github/actions/workflow/status/JMcrafter26/TridUI/release.yml?branch=main&label=build&logo=github&style=flat)](https://github.com/JMcrafter26/TridUI/actions/workflows/release.yml)

## Fonctionnalités

- 🚀 Analyse rapide basée sur Go natif
- 🎯 Identification précise des types de fichiers via les définitions TrID
- 💻 Application multi‑plateforme (Windows, macOS, Linux)
- 🔒 Traitement 100 % local – aucune donnée ne quitte votre machine
- 🎨 Interface moderne et intuitive
- 📊 Résultats détaillés avec scores de confiance
- 🔄 Glisser‑déposer des fichiers
- 🔁 Mises à jour automatiques des définitions en un clic
- 📅 Suivi de la date de mise à jour et du nombre de définitions

## Sommaire

<details>
<summary>Cliquer pour développer</summary>

- [Fonctionnalités](#fonctionnalités)
- [Sommaire](#sommaire)
- [Démonstration et captures](#démonstration-et-captures)
  - [Vidéos de démonstration](#vidéos-de-démonstration)
  - [Captures](#captures)
- [Installation](#installation)
  - [Prérequis](#prérequis)
    - [Option 1 : Téléchargement automatique (recommandé)](#option-1-téléchargement-automatique-recommandé)
    - [Option 2 : Installation manuelle](#option-2-installation-manuelle)
  - [Compilation depuis les sources](#compilation-depuis-les-sources)
- [Utilisation](#utilisation)
- [Détails techniques](#détails-techniques)
  - [Architecture](#architecture)
  - [Implémentation du scanner TrID](#implémentation-du-scanner-trid)
- [Licence et attributions](#licence-et-attributions)
- [Contribuer](#contribuer)
  - [Traductions](#traductions)

</details>

## Démonstration et captures

### Vidéos de démonstration

<details>
<summary>Cliquer pour développer</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Captures

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="Capture TrID UI 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="Capture TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="Capture TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Afficher plus</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="Capture TrID UI 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="Capture TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="Capture TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="Capture TrID UI 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="Capture TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Installation

> [!TIP]
> Des binaires pré‑compilés sont disponibles sur la [page des releases](https://github.com/JMcrafter26/TridUI/releases).

### Prérequis

L’application peut télécharger et mettre à jour automatiquement le fichier de définitions TrID.

#### Option 1 : Téléchargement automatique (recommandé)

1. Lancez TrID UI
2. Ouvrez les paramètres
3. Cliquez sur « Download Definitions » ou « Check for Updates »
4. L’application téléchargera et installera automatiquement les dernières définitions

#### Option 2 : Installation manuelle

1. Téléchargez le fichier de définitions TrID (`triddefs.trd`) depuis [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Placez `triddefs.trd` dans le répertoire de données de l’application :
   - **Windows** : `%APPDATA%\TridUI\triddefs.trd`
   - **macOS** : `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux** : `~/.local/share/TridUI/triddefs.trd`

Vous pouvez utiliser le bouton « Open App Dir » pour ouvrir l’emplacement adéquat.

### Compilation depuis les sources

> **📖 Documentation complète de build :** voir [`build/README.md`](../../build/README.md) pour les instructions détaillées et le dépannage.

**Build rapide :**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Ce que font les scripts :**
- ✅ Valident les prérequis (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Vérifient les dépendances système
- ✅ Détectent les outils optionnels (UPX, NSIS, create-dmg)
- ✅ Détectent automatiquement l’architecture
- ✅ Produisent des paquets distribuables

**Prérequis minimum :**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Installer Wails CLI :**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Emplacements de sortie :**
- Windows : `build/bin/windows/TridUI-win-{arch}.exe`
- macOS : `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux : `build/bin/linux/TridUI-linux-{arch}`

## Utilisation

1. Lancez TrID UI
2. Cliquez ou glissez‑déposez un fichier dans l’interface
3. Consultez les résultats avec les scores de confiance
4. La meilleure correspondance est mise en évidence en haut
5. Les correspondances supplémentaires apparaissent en dessous

## Détails techniques

### Architecture

- **Backend** : Go (framework Wails)
- **Frontend** : SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Moteur TrID** : implémentation Go pure (package `/trid`)

### Implémentation du scanner TrID

Le scanner TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) est une implémentation Go « clean‑room » qui :

- Analyse les fichiers TRD (définitions TrID) selon la spécification binaire
- Effectue des correspondances de motifs à des offsets donnés
- Gère la correspondance de chaînes pour une meilleure précision
- Calcule des scores de confiance basés sur les poids des motifs
- Retourne des résultats classés avec des informations détaillées

> La spécification du format TRD est disponible sur [Mark0.net](https://mark0.net/soft-trid-format.html).

## Licence et attributions

TrID UI est un logiciel open‑source sous licence GNU AGPLv3. L’interface est développée par Cufiy (alias JMcrafter26) et est basée sur TrID de [Marco Pontello](https://mark0.net/).
Reportez‑vous au fichier LICENSE pour plus de détails.

Le scanner `trid.go` est une implémentation « clean‑room » en Go par JMcrafter26, sous licence GNU AGPLv3.

L’icône de l’application est basée sur une icône œil de icons8.com.

<a href="https://github.com/JMcrafter26/TridUI/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=JMcrafter26/TridUI" />
</a>

## Contribuer

Les contributions sont les bienvenues ! Pour contribuer à TrID UI, forkez le dépôt et ouvrez une pull request. Pour des changements majeurs, ouvrez d’abord une issue pour en discuter.

### Traductions

TrID UI a besoin de votre aide pour toucher un plus large public ! Les traductions actuelles sont générées automatiquement et peuvent comporter des imprécisions.

Pour contribuer des traductions :

1. Forkez le dépôt
2. Créez une branche dédiée
3. Ajoutez vos fichiers dans le répertoire `translations`
4. Ouvrez une pull request

Merci d’aider à rendre TrID UI accessible au plus grand nombre !
