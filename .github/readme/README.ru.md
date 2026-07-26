<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Иконка TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/demo.gif?raw=true" alt="Демонстрация TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <a href="README.es.md">Español</a> •
    <a href="README.fr.md">Français</a> •
    <a href="README.it.md">Italiano</a> •
    <a href="README.ja.md">日本語</a> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <strong>Русский</strong> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI — легкое настольное приложение с удобным интерфейсом для TrID, мощного инструмента для сканирования и анализа файлов. Вы можете выбрать файл или перетащить его на главный экран, чтобы запустить локальное сканирование и определить неизвестные типы файлов.

Приложение использует нативную реализацию на Go алгоритма идентификации файлов TrID, обеспечивая быструю и точную идентификацию без внешних зависимостей.

> [!TIP]
> Скачайте TridUI на [странице релизов](https://github.com/JMcrafter26/TridUI/releases)

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JMcrafter26/TridUI?style=flat&logo=go)](https://github.com/JMcrafter26/TridUI)
[![GitHub release (последний)](https://img.shields.io/github/v/release/JMcrafter26/TridUI?style=flat&label=последний+релиз&logo=github)](https://github.com/JMcrafter26/TridUI/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/JMcrafter26/TridUI?style=flat&logo=github)](https://github.com/JMcrafter26/TridUI/issues)
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

## Возможности

- 🚀 Быстрое сканирование на нативном Go
- 🎯 Точная идентификация типов файлов с помощью определений TrID
- 💻 Кроссплатформенность (Windows, macOS, Linux)
- 🔒 100% локальная обработка — данные не покидают ваш компьютер
- 🎨 Современный и интуитивный интерфейс
- 📊 Подробные результаты с оценкой доверия
- 🔄 Поддержка перетаскивания файлов (drag & drop)
- 🔁 Автообновление определений в один клик
- 🖥️ Интерфейс командной строки (CLI) для продвинутых пользователей

## Оглавление

<details>
<summary>Развернуть</summary>

- [Возможности](#возможности)
- [Оглавление](#оглавление)
- [Демонстрация и скриншоты](#демонстрация-и-скриншоты)
  - [Демонстрационные видео](#демонстрационные-видео)
  - [Скриншоты](#скриншоты)
- [Установка](#установка)
  - [Предварительные требования](#предварительные-требования)
    - [Вариант 1: Автозагрузка (рекомендуется)](#вариант-1-автозагрузка-рекомендуется)
    - [Вариант 2: Ручная установка](#вариант-2-ручная-установка)
  - [Сборка из исходников](#сборка-из-исходников)
- [Использование](#использование)
- [Технические детали](#технические-детали)
  - [Архитектура](#архитектура)
  - [Реализация сканера TrID](#реализация-сканера-trid)
- [Лицензия и благодарности](#лицензия-и-благодарности)
- [Вклад](#вклад)
  - [Переводы](#переводы)

</details>

## Демонстрация и скриншоты

### Демонстрационные видео

<details>
<summary>Развернуть</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Скриншоты

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/home.png?raw=true" alt="TrID UI Скриншот 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/scan.png?raw=true" alt="TrID UI Скриншот 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/about.png?raw=true" alt="TrID UI Скриншот 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Показать ещё</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings.png?raw=true" alt="TrID UI Скриншот 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings2.png?raw=true" alt="TrID UI Скриншот 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/settings3.png?raw=true" alt="TrID UI Скриншот 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/unknown.png?raw=true" alt="TrID UI Скриншот 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.2.0/about.png?raw=true" alt="TrID UI Скриншот 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Установка

> [!TIP]
> Готовые бинарные файлы доступны на [странице релизов](https://github.com/JMcrafter26/TridUI/releases).

### Предварительные требования

Приложение может автоматически загружать и обновлять файл определений TrID.

#### Вариант 1: Автозагрузка (рекомендуется)

1. Запустите TrID UI
2. Откройте Настройки
3. Нажмите «Download Definitions» или «Check for Updates»
4. Приложение автоматически загрузит и установит последние определения

#### Вариант 2: Ручная установка

1. Скачайте файл определений TrID (`triddefs.trd`) с [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Поместите `triddefs.trd` в каталог данных приложения:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

В Настройках можно нажать «Open App Dir», чтобы открыть нужную папку.

### Сборка из исходников

> **📖 Полная документация по сборке:** см. [`build/README.md`](../../build/README.md) для подробностей и устранения неполадок.

**Быстрая сборка:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Что делают скрипты:**
- ✅ Проверяют предпосылки (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Проверяют системные зависимости
- ✅ Обнаруживают дополнительные инструменты (UPX, NSIS, create-dmg)
- ✅ Автоматически определяют архитектуру
- ✅ Создают дистрибутивные пакеты

**Минимальные требования:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Установка Wails CLI:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Каталоги вывода:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## Использование

1. Запустите TrID UI
2. Нажмите или перетащите файл в интерфейс
3. Просмотрите результаты со значениями доверия
4. Лучшее совпадение выделено сверху
5. Дополнительные совпадения перечислены ниже

## Технические детали

### Архитектура

- **Backend**: Go (Wails)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Движок TrID**: Чистая реализация на Go (пакет `/trid`)

### Реализация сканера TrID

Сканер TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) — clean‑room реализация на Go, которая:

- Разбирает файлы TRD по бинарной спецификации
- Выполняет сопоставление шаблонов на заданных смещениях
- Поддерживает сопоставление строк для повышения точности
- Вычисляет доверие на основе весов шаблонов
- Возвращает ранжированные результаты с подробной информацией

> Спецификацию формата TRD см. на [Mark0.net](https://mark0.net/soft-trid-format.html).

## Лицензия и благодарности

TrID UI — ПО с открытым исходным кодом под лицензией GNU AGPLv3. Интерфейс разработан Cufiy (aka JMcrafter26) и основан на TrID от [Marco Pontello](https://mark0.net/).
См. файл LICENSE для подробностей.

Сканер `trid.go` — clean‑room реализация на Go от JMcrafter26 под GNU AGPLv3.

Иконка приложения основана на иконке «eye» с icons8.com.

<a href="https://github.com/JMcrafter26/TridUI/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=JMcrafter26/TridUI" />
</a>

## Вклад

Будем рады любому вкладу! Форкните репозиторий и создайте pull request. Для существенных изменений сначала создайте issue для обсуждения.

### Переводы

TrID UI нуждается в вашей помощи, чтобы охватить больше пользователей! Текущие переводы созданы машиной и могут содержать неточности.

Как помочь с переводами:

1. Сделайте fork репозитория
2. Создайте ветку для перевода
3. Добавьте файлы в каталог `translations`
4. Откройте pull request

Спасибо за помощь в развитии доступности TrID UI!
