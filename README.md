![Go](https://img.shields.io/badge/Go-1.21-blue)
![GUI](https://img.shields.io/badge/Fyne-GUI-green)
![Status](https://img.shields.io/badge/Project%20Status-Complete-brightgreen)

# 🧩 Quadchecker-GUI

A modern Go + Fyne GUI application that visually detects and validates ASCII-based quad patterns (QuadA → QuadE).
Built as part of my 01Talent x Nextera Go Piscine journey.

---

## 🧠 Overview

This tool extends the original Quad Checker command-line project with a graphical interface, background music, and animated ASCII logo.
It allows users to run quads interactively, visualize the generated shapes, and see which quad matches the input.

---

## Features

- 🎨 Built using the Fyne GUI Framework
- 🎵 Background music playback via beep
- ⚙️ Detects multiple matching quads alphabetically (QuadC || QuadE)
- 💫 Animated ASCII logo splash
- 🧩 Modular code structure

---

## 📁 Project Structure

```
Quadchecker-GUI/
├── main.go
├── go.mod
├── go.sum
├── /assets
│   └── pixify-230092.mp3
└── /Internal
    ├── quadA
    ├── quadB
    ├── quadC
    ├── quadD
    └── quadE
```

## 🚀 How to Run

```bash
git clone https://github.com/EssamGamal88/Quadchecker-GUI.git
cd Quadchecker-GUI
chmod +x Internal/quadA Internal/quadB Internal/quadC Internal/quadD Internal/quadE
go run .
```

### Build
```bash
go build -o quadchecker-gui .
./quadchecker-gui
```

---

## 🔧 Tech
- **Go** 1.21
- **Fyne v2** (GUI)
- **beep/mp3** (optional audio)

---

## 🧠 What I Learned

- Building a small **desktop GUI** with Fyne  
- Managing **string rendering** for ASCII shapes  
- Handling **process interaction & UX** (input → output → result)  
- Preparing a repo for growth (docs, structure, assets)

---

⭐ *If you liked this project or found it useful, feel free to star the repo and follow my journey on [GitHub](https://github.com/EssamGamal88)!*
