# THE FINALS Auto-Ping

Automatically ping enemies while aiming in THE FINALS. Improve team communication and win more games — even in solo queue.

![Demo](https://github.com/user-attachments/assets/bff17d8c-ceb6-458f-b3ba-b447ce5a8e2a)

## ✨ Features

- **Auto-ping while aiming** — Hold right-click, enemies get pinged automatically
- **Repeating pings** — Keeps pinging every second while you aim
- **Lightweight** — Minimal CPU usage, runs quietly in the background
- **Safe** — Does not modify game files or interact with the game process

## 🎮 Download

1. Go to the [Releases](../../releases) page
2. Download the latest `.exe` file
3. Run it — no installation needed

> **Note:** Windows only.

## ⚙️ Default Controls

| Action  | Key                |
|---------|--------------------|
| Trigger | Right Mouse Button |
| Ping    | Left Control       |

> **Tip:** Make sure your in-game ping key is set to Left Control.

---

## 🛠️ Advanced: Build from Source

For developers who want to compile from source:

**Requirements:** [Go 1.25+](https://golang.org/dl/)

```bash
# Windows
go build -o autoping.exe main.go

# macOS / Linux (cross-compile)
GOOS=windows GOARCH=amd64 go build -o autoping.exe main.go
```

**Customization:** Edit `VK_RBUTTON` and `VK_LCONTROL` in `main.go` to change key bindings. See [Windows Virtual-Key Codes](https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes) for reference.

---

## 💬 Feedback & Suggestions

Have an idea or found a bug? Feel free to [open an issue](../../issues) or contact me directly.

---

## 📄 License

MIT
