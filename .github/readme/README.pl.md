<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Ikona TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/demo.gif?raw=true" alt="Prezentacja TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <a href="README.es.md">Español</a> •
    <a href="README.fr.md">Français</a> •
    <a href="README.it.md">Italiano</a> •
    <a href="README.ja.md">日本語</a> •
    <strong>Polski</strong> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI to lekka aplikacja desktopowa z przyjaznym interfejsem dla TrID – potężnego narzędzia do skanowania i analizy plików. Umożliwia łatwy wybór lub przeciąganie plików na ekran główny, aby rozpocząć lokalne skanowanie i wykrywać nierozpoznane typy plików.

Aplikacja korzysta z natywnej implementacji algorytmu identyfikacji plików TrID w Go, zapewniając szybkie i dokładne wykrywanie bez zewnętrznych zależności.

> [!TIP]
> Pobierz TridUI ze [strony wydań](https://github.com/JMcrafter26/TridUI/releases)

## Funkcje

- 🚀 Szybkie skanowanie oparte na natywnym Go
- 🎯 Dokładna identyfikacja typów plików z użyciem definicji TrID
- 💻 Aplikacja wieloplatformowa (Windows, macOS, Linux)
- 🔒 100% lokalne przetwarzanie – żadne dane nie opuszczają komputera
- 🎨 Nowoczesny, intuicyjny interfejs
- 📊 Szczegółowe wyniki z poziomami ufności
- 🔄 Obsługa przeciągnij‑i‑upuść
- 🔁 Automatyczne aktualizacje definicji jednym kliknięciem
- 📅 Informacje o dacie ostatniej aktualizacji i liczbie definicji

## Spis treści

<details>
<summary>Kliknij, aby rozwinąć</summary>

- [Funkcje](#funkcje)
- [Spis treści](#spis-treści)
- [Prezentacja i zrzuty ekranu](#prezentacja-i-zrzuty-ekranu)
  - [Filmy demonstracyjne](#filmy-demonstracyjne)
  - [Zrzuty ekranu](#zrzuty-ekranu)
- [Instalacja](#instalacja)
  - [Wymagania wstępne](#wymagania-wstępne)
    - [Opcja 1: Automatyczne pobieranie (Zalecane)](#opcja-1-automatyczne-pobieranie-zalecane)
    - [Opcja 2: Ręczna instalacja](#opcja-2-ręczna-instalacja)
  - [Budowanie ze źródeł](#budowanie-ze-źródeł)
- [Użycie](#użycie)
- [Szczegóły techniczne](#szczegóły-techniczne)
  - [Architektura](#architektura)
  - [Implementacja skanera TrID](#implementacja-skanera-trid)
- [Licencja i atrybucja](#licencja-i-atrybucja)
- [Współpraca](#współpraca)
  - [Tłumaczenia](#tłumaczenia)

</details>

## Prezentacja i zrzuty ekranu

### Filmy demonstracyjne

<details>
<summary>Kliknij, aby rozwinąć</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### Zrzuty ekranu

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="Zrzut TrID UI 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="Zrzut TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="Zrzut TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Pokaż więcej</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="Zrzut TrID UI 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="Zrzut TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="Zrzut TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="Zrzut TrID UI 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="Zrzut TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Instalacja

> [!TIP]
> Gotowe binaria znajdziesz na [stronie wydań](https://github.com/JMcrafter26/TridUI/releases).

### Wymagania wstępne

Aplikacja może automatycznie pobierać i aktualizować plik definicji TrID.

#### Opcja 1: Automatyczne pobieranie (Zalecane)

1. Uruchom TrID UI
2. Otwórz Ustawienia
3. Kliknij „Download Definitions” lub „Check for Updates”
4. Aplikacja automatycznie pobierze i zainstaluje najnowsze definicje

#### Opcja 2: Ręczna instalacja

1. Pobierz plik definicji TrID (`triddefs.trd`) ze strony [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Umieść `triddefs.trd` w katalogu danych aplikacji:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

W Ustawieniach możesz użyć przycisku „Open App Dir”, aby otworzyć właściwy katalog.

### Budowanie ze źródeł

> **📖 Pełna dokumentacja budowania:** zobacz [`build/README.md`](../../build/README.md) po szczegóły i pomoc.

**Szybkie budowanie:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**Co robią skrypty build:**
- ✅ Weryfikują wymagania (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Sprawdzają zależności systemowe
- ✅ Wykrywają opcjonalne narzędzia (UPX, NSIS, create-dmg)
- ✅ Automatycznie wykrywają architekturę
- ✅ Tworzą pakiety dystrybucyjne

**Minimalne wymagania:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Instalacja Wails CLI:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Lokalizacje wyjściowe:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux/{arch}`

## Użycie

1. Uruchom TrID UI
2. Kliknij lub przeciągnij plik na interfejs
3. Zobacz wyniki skanowania ze współczynnikami ufności
4. Najlepsze dopasowanie jest wyróżnione u góry
5. Dodatkowe możliwe dopasowania są poniżej

## Szczegóły techniczne

### Architektura

- **Backend**: Go (framework Wails)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Silnik TrID**: Czysta implementacja w Go (pakiet `/trid`)

### Implementacja skanera TrID

Skaner TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) to czysta implementacja w Go, która:

- Parsuje pliki TRD według specyfikacji binarnej
- Wykonuje dopasowania wzorców na określonych offsetach
- Wspiera dopasowania łańcuchów dla lepszej dokładności
- Oblicza poziomy ufności na podstawie wag wzorców
- Zwraca uporządkowane wyniki ze szczegółami typu pliku

> Specyfikację formatu TRD znajdziesz na [Mark0.net](https://mark0.net/soft-trid-format.html).

## Licencja i atrybucja

TrID UI to oprogramowanie open‑source na licencji GNU AGPLv3. UI zostało stworzone przez Cufiy (JMcrafter26) i bazuje na TrID autorstwa [Marco Pontello](https://mark0.net/).
Szczegóły w pliku LICENSE.

Skaner `trid.go` to czysta implementacja w Go autorstwa JMcrafter26 na licencji GNU AGPLv3.

Ikona aplikacji bazuje na ikonie „eye” z icons8.com.

## Współpraca

Wszelkie kontrybucje są mile widziane! Jeśli chcesz pomóc, zrób fork repozytorium i wyślij pull request. Przy większych zmianach prosimy najpierw o otwarcie issue w celu omówienia.

### Tłumaczenia

TrID UI potrzebuje Twojej pomocy, aby dotrzeć do szerszego grona! Obecne tłumaczenia są generowane maszynowo i mogą zawierać nieścisłości.

Jak dodać tłumaczenie:

1. Zrób fork repozytorium
2. Utwórz nowy branch dla tłumaczenia
3. Dodaj pliki w katalogu `translations`
4. Wyślij pull request

Dziękujemy za pomoc w zwiększaniu dostępności TrID UI!
