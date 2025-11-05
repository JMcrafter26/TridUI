<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="Ícone do TrID UI" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/demo.gif?raw=true" alt="Demonstração do TrID UI" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
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
    <strong>Português</strong> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

O TrID UI é um aplicativo desktop leve que fornece uma interface amigável para o TrID, uma poderosa ferramenta para escanear e analisar arquivos. Com o TrID UI, você pode selecionar ou arrastar arquivos para a tela inicial para iniciar varreduras locais e detectar tipos de arquivo não reconhecidos.

O aplicativo usa uma implementação nativa em Go do algoritmo de identificação de arquivos TrID, oferecendo detecção rápida e precisa sem dependências externas.

> [!TIP]
> Baixe o TridUI na [página de releases](https://github.com/JMcrafter26/TridUI/releases)

## Recursos

- 🚀 Varredura rápida baseada em Go nativo
- 🎯 Identificação precisa de tipos de arquivo usando definições TrID
- 💻 Aplicativo multiplataforma (Windows, macOS, Linux)
|- 🔒 Processamento 100% local – nenhum dado sai do seu computador
- 🎨 Interface moderna e intuitiva
- 📊 Resultados detalhados com pontuações de confiança
- 🔄 Suporte a arrastar e soltar arquivos
- 🔁 Atualizações automáticas de definições com um clique
- 📅 Acompanhe a data da última atualização e a contagem de definições

## Sumário

<details>
<summary>Clique para expandir</summary>

- [Recursos](#recursos)
- [Sumário](#sumário)
- [Demonstração e capturas](#demonstração-e-capturas)
  - [Vídeos de demonstração](#vídeos-de-demonstração)
  - [Capturas](#capturas)
- [Configuração](#configuração)
  - [Pré-requisitos](#pré-requisitos)
    - [Opção 1: Download automático (Recomendado)](#opção-1-download-automático-recomendado)
    - [Opção 2: Instalação manual](#opção-2-instalação-manual)
  - [Compilar a partir do código-fonte](#compilar-a-partir-do-código-fonte)
- [Uso](#uso)
- [Detalhes técnicos](#detalhes-técnicos)
  - [Arquitetura](#arquitetura)
  - [Implementação do scanner TrID](#implementação-do-scanner-trid)
- [Licença e atribuição](#licença-e-atribuição)
- [Contribuindo](#contribuindo)
  - [Traduções](#traduções)

</details>

## Demonstração e capturas

### Vídeos de demonstração

<details>
<summary>Clique para expandir</summary>

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
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="Captura TrID UI 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="Captura TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="Captura TrID UI 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>Mostrar mais</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="Captura TrID UI 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="Captura TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="Captura TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="Captura TrID UI 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="Captura TrID UI 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## Configuração

> [!TIP]
> Você encontra binários pré-compilados na [página de releases](https://github.com/JMcrafter26/TridUI/releases).

### Pré-requisitos

O aplicativo pode baixar e atualizar automaticamente o arquivo de definições do TrID.

#### Opção 1: Download automático (Recomendado)

1. Inicie o TrID UI
2. Abra as Configurações
3. Clique em “Download Definitions” ou “Check for Updates”
4. O app fará o download e instalará as últimas definições automaticamente

#### Opção 2: Instalação manual

1. Baixe o arquivo de definições do TrID (`triddefs.trd`) de [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Coloque o `triddefs.trd` no diretório de dados do aplicativo:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

Você pode usar o botão “Open App Dir” nas Configurações para abrir a pasta correta.

### Compilar a partir do código-fonte

> **📖 Documentação completa de build:** veja [`build/README.md`](../../build/README.md) para instruções detalhadas e solução de problemas.

**Build rápido:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**O que os scripts fazem:**
- ✅ Validam pré-requisitos (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ Checam dependências do sistema
- ✅ Detectam ferramentas opcionais (UPX, NSIS, create-dmg)
- ✅ Detectam arquitetura automaticamente
- ✅ Criam pacotes distribuíveis

**Requisitos mínimos:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Instalar Wails CLI:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**Locais de saída:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## Uso

1. Inicie o TrID UI
2. Clique ou arraste e solte um arquivo na interface
3. Veja os resultados com pontuações de confiança
4. A melhor correspondência aparece destacada no topo
5. Outras correspondências possíveis aparecem abaixo

## Detalhes técnicos

### Arquitetura

- **Backend**: Go (framework Wails)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **Motor TrID**: Implementação pura em Go (pacote `/trid`)

### Implementação do scanner TrID

O scanner TrID ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) é uma implementação em Go “clean-room” que:

- Analisa arquivos TRD (definições TrID) conforme a especificação binária
- Realiza correspondência de padrões em offsets especificados
- Suporta correspondência de strings para maior precisão
- Calcula pontuações de confiança baseadas nos pesos dos padrões
- Retorna resultados ranqueados com informações detalhadas

> A especificação do formato TRD está em [Mark0.net](https://mark0.net/soft-trid-format.html).

## Licença e atribuição

TrID UI é software open-source sob licença GNU AGPLv3. A UI foi desenvolvida por Cufiy (aka JMcrafter26) e é baseada no TrID de [Marco Pontello](https://mark0.net/).
Consulte o arquivo LICENSE para mais detalhes.

O scanner `trid.go` é uma implementação “clean-room” em Go por JMcrafter26, sob GNU AGPLv3.

O ícone do app é baseado no ícone “eye” do icons8.com.

## Contribuindo

Contribuições são bem-vindas! Para contribuir com o TrID UI, faça fork do repositório e envie um pull request. Para mudanças maiores, abra primeiro uma issue para discussão.

### Traduções

O TrID UI precisa da sua ajuda para alcançar mais pessoas! As traduções atuais são geradas por máquina e podem conter imprecisões.

Para contribuir com traduções:

1. Faça fork do repositório
2. Crie um branch para a tradução
3. Adicione seus arquivos no diretório `translations`
4. Envie um pull request

Obrigado por ajudar a tornar o TrID UI acessível para mais usuários!
