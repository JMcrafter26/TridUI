<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="../../icon.png" alt="TrID UI 图标" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/demo.gif?raw=true" alt="TrID UI 演示" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
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
    <a href="README.ru.md">Русский</a> •
    <strong>简体中文</strong>
  </span>
</p>

TrID UI 是一款轻量级桌面应用，为强大的 TrID 文件扫描与分析工具提供了友好的图形界面。用户可在首页选择或拖拽文件，即可发起本地扫描，便捷识别未知文件类型。

应用使用 TrID 文件识别算法的 Go 原生实现，无需外部依赖，提供快速而准确的类型检测。

> [!TIP]
> 在 [Releases 页面](https://github.com/JMcrafter26/TridUI/releases) 下载 TridUI

## 功能特性

- 🚀 基于 Go 原生实现的高速扫描
- 🎯 使用 TrID 定义进行精准文件类型识别
- 💻 跨平台桌面应用（Windows、macOS、Linux）
- 🔒 100% 本地处理——不会上传任何数据
- 🎨 现代、直观的用户界面
- 📊 搭配置信度评分的详细匹配结果
- 🔄 支持拖拽文件
- 🔁 一键自动更新定义文件
- 📅 显示最后一次更新日期与定义数量

## 目录

<details>
<summary>点击展开</summary>

- [功能特性](#功能特性)
- [目录](#目录)
- [演示与截图](#演示与截图)
  - [演示视频](#演示视频)
  - [截图](#截图)
- [安装与配置](#安装与配置)
  - [先决条件](#先决条件)
    - [方案一：自动下载（推荐）](#方案一自动下载推荐)
    - [方案二：手动安装](#方案二手动安装)
  - [从源码构建](#从源码构建)
- [使用方法](#使用方法)
- [技术细节](#技术细节)
  - [架构](#架构)
  - [TrID 扫描器实现](#trid-扫描器实现)
- [许可证与致谢](#许可证与致谢)
- [贡献](#贡献)
  - [翻译](#翻译)

</details>

## 演示与截图

### 演示视频

<details>
<summary>点击展开</summary>

https://github.com/user-attachments/assets/ecd4dbf3-77a3-4f07-8436-c1068e755d5f

https://github.com/user-attachments/assets/45d88137-3bf9-4c25-b516-6f344a1403a5

https://github.com/user-attachments/assets/766d55df-33e6-45d7-b2ae-cc4e02f55429

https://github.com/user-attachments/assets/c1adec87-dc68-4c0c-860f-f6f7d1cd1303

https://github.com/user-attachments/assets/6716fdbf-65c1-4c07-b8af-26a2912c84e6

https://github.com/user-attachments/assets/5c1e32e7-84ea-4815-9097-5134956f5e4d

https://github.com/user-attachments/assets/bde82ca9-fa8e-45a3-acd4-c31040aea11b

</details>

### 截图

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/home.png?raw=true" alt="TrID UI 截图 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scan.png?raw=true" alt="TrID UI 截图 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/scanning.png?raw=true" alt="TrID UI 截图 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />

<details>
 <summary>展开更多</summary>
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings.png?raw=true" alt="TrID UI 截图 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings2.png?raw=true" alt="TrID UI 截图 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
    <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/settings3.png?raw=true" alt="TrID UI 截图 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/unknown.png?raw=true" alt="TrID UI 截图 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
      <img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/assets/1.1.0/about.png?raw=true" alt="TrID UI 截图 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</details>
</div>

## 安装与配置

> [!TIP]
> 你可以在 [Releases 页面](https://github.com/JMcrafter26/TridUI/releases) 找到预构建的二进制包。

### 先决条件

应用可以为你自动下载并更新 TrID 定义文件！

#### 方案一：自动下载（推荐）

1. 启动 TrID UI
2. 打开设置
3. 点击“Download Definitions”或“Check for Updates”
4. 应用会自动下载并安装最新定义

#### 方案二：手动安装

1. 从 [Mark0.net](https://mark0.net/soft-trid-deflist.html) 下载 TrID 定义文件 (`triddefs.trd`)
2. 将 `triddefs.trd` 放到应用数据目录：
   - **Windows**：`%APPDATA%\TridUI\triddefs.trd`
   - **macOS**：`~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**：`~/.local/share/TridUI/triddefs.trd`

可在设置中点击“Open App Dir”按钮快速打开目标目录。

### 从源码构建

> **📖 完整构建文档：** 参见 [`build/README.md`](../../build/README.md) 获取详细说明与故障排查。

**快速构建：**

```bash
# Windows
.\build\build-windows.bat

# macOS
chmod +x build/build-darwin.sh && ./build/build-darwin.sh

# Linux
chmod +x build/build-linux.sh && ./build/build-linux.sh
```

**构建脚本将：**
- ✅ 验证前置条件（Go 1.22+、Wails CLI、Node.js、pnpm）
- ✅ 检查系统依赖
- ✅ 检测可选工具（UPX、NSIS、create-dmg）
- ✅ 自动检测架构
- ✅ 生成可发布包

**最低要求：**
- Go 1.22+ • Node.js 20+ • pnpm 10+ • Wails CLI

**安装 Wails CLI：**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

**输出位置：**
- Windows：`build/bin/windows/TridUI-win-{arch}.exe`
- macOS：`build/bin/darwin/TridUI-macOS-{arch}.dmg`（另含 .app、.zip）
- Linux：`build/bin/linux/TridUI-linux-{arch}`

## 使用方法

1. 启动 TrID UI
2. 单击或拖拽文件到界面
3. 查看带置信度评分的结果
4. 最佳匹配高亮显示在顶部
5. 其他可能匹配列于下方

## 技术细节

### 架构

- **后端**：Go（Wails 框架）
- **前端**：SvelteKit + TypeScript + DaisyUI（与 Tailwind CSS）
- **TrID 引擎**：纯 Go 实现（`/trid` 包）

### TrID 扫描器实现

TrID 扫描器（[`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)）为 clean‑room Go 实现：

- 按二进制格式规范解析 TRD（TrID 定义）文件
- 在指定偏移执行模式匹配
- 支持字符串匹配以提升准确度
- 根据模式权重计算置信度评分
- 返回带详细信息的排序结果

> TRD 格式规范见 [Mark0.net](https://mark0.net/soft-trid-format.html)。

## 许可证与致谢

TrID UI 采用 GNU AGPLv3 开源许可证。UI 由 Cufiy（即 JMcrafter26）开发，基于 [Marco Pontello](https://mark0.net/) 的 TrID。
详情请参见 LICENSE 文件。

`trid.go` 扫描器由 JMcrafter26 以 clean‑room 方式实现，采用 GNU AGPLv3 许可。

应用图标基于 icons8.com 的 eye 图标。

## 贡献

欢迎贡献！如欲参与 TrID UI 的开发，请 fork 本仓库并提交 Pull Request。对于重大更改，请先开一个 Issue 进行讨论。

### 翻译

TrID UI 需要你的帮助以覆盖更多用户！当前翻译由机器生成，可能存在不准确之处。

贡献翻译步骤：

1. fork 本仓库
2. 为你的翻译创建新分支
3. 将文件添加到 `translations` 目录
4. 提交 Pull Request

感谢你帮助让 TrID UI 更加易用！
