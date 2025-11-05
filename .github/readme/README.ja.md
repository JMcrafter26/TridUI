<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="TrID UI アイコン" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/demo.gif?raw=true" alt="TrID UI デモ" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

<p align="center">
  <span style="font-size: 0.95em; opacity: .8">
    <a href="../../README.md">English</a> •
    <a href="README.de.md">Deutsch</a> •
    <a href="README.es.md">Español</a> •
    <a href="README.fr.md">Français</a> •
    <a href="README.it.md">Italiano</a> •
    <strong>日本語</strong> •
    <a href="README.pl.md">Polski</a> •
    <a href="README.pt.md">Português</a> •
    <a href="README.ru.md">Русский</a> •
    <a href="README.zh.md">简体中文</a>
  </span>
</p>

TrID UI は、ファイルのスキャンと解析に強力な TrID を、扱いやすい UI で利用できる軽量デスクトップアプリです。ホーム画面でファイルを選択またはドラッグ＆ドロップするだけでローカルスキャンを開始し、未認識のファイルタイプを検出できます。

このアプリは、TrID のファイル識別アルゴリズムを Go でネイティブ実装しており、外部依存なしで高速かつ正確な検出を実現します。

> [!TIP]
> [Releases ページ](https://github.com/JMcrafter26/TridUI/releases)から TridUI をダウンロードできます

## 特長

- 🚀 ネイティブ Go による高速スキャン
- 🎯 TrID 定義による高精度のファイルタイプ識別
- 💻 クロスプラットフォーム対応（Windows・macOS・Linux）
- 🔒 100% ローカル処理 – データは外部に送信されません
- 🎨 モダンで直感的な UI
- 📊 信頼度スコア付きの詳細な結果
- 🔄 ファイルのドラッグ＆ドロップ対応
- 🔁 ワンクリックで定義ファイルを自動更新
- 📅 最終更新日と定義数を表示

## 目次

<details>
<summary>クリックして展開</summary>

- [特長](#特長)
- [目次](#目次)
- [デモとスクリーンショット](#デモとスクリーンショット)
  - [デモ動画](#デモ動画)
  - [スクリーンショット](#スクリーンショット)
- [セットアップ](#セットアップ)
  - [前提条件](#前提条件)
    - [オプション1: 自動ダウンロード（推奨）](#オプション1-自動ダウンロード推奨)
    - [オプション2: 手動インストール](#オプション2-手動インストール)
  - [ソースからビルド](#ソースからビルド)
- [使い方](#使い方)
- [技術詳細](#技術詳細)
  - [アーキテクチャ](#アーキテクチャ)
  - [TrID スキャナ実装](#trid-スキャナ実装)
- [ライセンスとクレジット](#ライセンスとクレジット)
- [コントリビュート](#コントリビュート)
  - [翻訳](#翻訳)

</details>

## デモとスクリーンショット

### デモ動画

<details>
<summary>クリックして展開</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### スクリーンショット

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="TrID UI スクリーンショット 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="TrID UI スクリーンショット 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="TrID UI スクリーンショット 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>さらに表示</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="TrID UI スクリーンショット 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="TrID UI スクリーンショット 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="TrID UI スクリーンショット 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="TrID UI スクリーンショット 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="TrID UI スクリーンショット 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## セットアップ

> [!TIP]
> ビルド済みバイナリは [Releases ページ](https://github.com/JMcrafter26/TridUI/releases) にあります。

### 前提条件

アプリは TrID 定義ファイルを自動でダウンロード・更新できます。

#### オプション1: 自動ダウンロード（推奨）

1. TrID UI を起動
2. 設定を開く
3. 「Download Definitions」または「Check for Updates」をクリック
4. アプリが最新の定義を自動でダウンロードして適用します

#### オプション2: 手動インストール

1. [Mark0.net](https://mark0.net/soft-trid-deflist.html) から TrID 定義ファイル (`triddefs.trd`) をダウンロード
2. `triddefs.trd` を以下のアプリデータディレクトリに配置:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

設定の「Open App Dir」ボタンで保存先を開けます。

### ソースからビルド

> **📖 詳細なビルド手順:** [`build/README.md`](../../build/README.md) を参照してください。

**クイックビルド:**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**ビルドスクリプトの内容:**
- ✅ 前提条件の検証 (Go 1.22+, Wails CLI, Node.js, pnpm)
- ✅ システム依存関係の確認
- ✅ 追加ツールの検出 (UPX, NSIS, create-dmg)
- ✅ アーキテクチャの自動検出
- ✅ 配布用パッケージの作成

**最小要件:**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**Wails CLI のインストール:**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**出力先:**
- Windows: `build/bin/windows/TridUI-win-{arch}.exe`
- macOS: `build/bin/darwin/TridUI-macOS-{arch}.dmg` (+ .app, .zip)
- Linux: `build/bin/linux/TridUI-linux-{arch}`

## 使い方

1. TrID UI を起動
2. クリックまたはドラッグ＆ドロップでファイルを指定
3. 信頼度スコア付きの結果を確認
4. 最上位の一致が上部にハイライト表示
5. その他の候補は下部に表示

## 技術詳細

### アーキテクチャ

- **バックエンド**: Go (Wails フレームワーク)
- **フロントエンド**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **TrID エンジン**: 純粋な Go 実装（`/trid` パッケージ）

### TrID スキャナ実装

TrID スキャナ（[`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)）は、クリーンルームな Go 実装で、以下を行います:

- TRD（TrID 定義）ファイルをバイナリ仕様に基づいてパース
- 指定オフセットでのパターンマッチング
- 精度向上のための文字列マッチング
- パターン重みづけに基づく信頼度スコアの計算
- 詳細情報付きの順位付け結果を返却

> TRD 形式の仕様は [Mark0.net](https://mark0.net/soft-trid-format.html) を参照してください。

## ライセンスとクレジット

TrID UI は GNU AGPLv3 ライセンスのオープンソースです。UI は Cufiy（JMcrafter26）によって開発され、[Marco Pontello](https://mark0.net/) の TrID に基づいています。
詳細は LICENSE を参照してください。

`trid.go` スキャナは JMcrafter26 によるクリーンルームな Go 実装で、GNU AGPLv3 ライセンスです。

アプリアイコンは icons8.com の eye アイコンを基にしています。

## コントリビュート

コントリビューション歓迎です！変更を提案する場合はリポジトリをフォークし、Pull Request を送ってください。大きな変更の場合は、まず Issue を開いて議論してください。

### 翻訳

TrID UI をより多くのユーザーに届けるため、翻訳のご協力をお願いします。現在の翻訳は機械生成であり、正確でない場合があります。

翻訳をコントリビュートする手順:

1. リポジトリをフォーク
2. 翻訳用のブランチを作成
3. `translations` ディレクトリに翻訳ファイルを追加
4. Pull Request を作成

ご協力ありがとうございます！
