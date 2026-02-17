
# 📚 Current Affairs TUI

A minimal, beautiful Terminal User Interface (TUI) application built with **Go + Bubble Tea** to browse and read Current Affairs articles directly from your terminal.

Built for speed, simplicity, and a clean reading experience.

NOTE: "enter" key functionality has not been added yet to view the article. Working on it
---

## ✨ Features

* 🔎 Scrapes articles by **month** and **year**
* 📜 Scrollable article list
* 📖 Full-screen article reader
* ⌨️ Vim-style navigation (`j/k`)
* 🖥️ Runs in alternate screen (like `nvim`)
* 📦 Single static binary

---

## 🛠 Built With

* [Bubble Tea](https://github.com/charmbracelet/bubbletea)
* [Bubbles](https://github.com/charmbracelet/bubbles)
* Go standard library

---

## 🚀 Build

Make sure you have Go installed (1.20+ recommended).

```bash
git clone <your-repo-url>
cd current-affairs
go mod tidy
go build -o current-affairs
```

Optional: install globally

```bash
sudo mv current-affairs /usr/bin/
```

Now you can run it from anywhere.

---

## 🖥 Usage

```bash
current-affairs -m <month> [-y <year>]
```

### Examples

```bash
current-affairs -m jan
current-affairs -m feb -y 2024
```

### Flags

| Flag            | Description                         | Required |
| --------------- | ----------------------------------- | -------- |
| `-m`, `--month` | Month name (`jan`, `january`, etc.) | ✅ Yes    |
| `-y`, `--year`  | Year (default: 2026)                | ❌ No     |

---

## 🎮 Controls

### List View

* `j` / `↓` → Move down
* `k` / `↑` → Move up
* `Enter` → Open article
* `q` / `Ctrl+C` → Quit

### Article View

* `j` / `k` / `↑` / `↓` → Scroll
* `Esc` → Back to list
* `q` / `Ctrl+C` → Quit

---

## 🧠 Architecture

```
Scraper → []article → bubbles/list → viewport → screen
```

* `list` handles browsing
* `viewport` handles reading
* State controls which view is active

---

## 🐧 Platform

* Designed for Linux
* Cross-compile for Windows:

```bash
GOOS=windows GOARCH=amd64 go build
```

---

## 🎯 Philosophy

Minimal.
Fast.
Terminal-native.

---

