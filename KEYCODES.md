# Key Code Reference

This file documents the Windows virtual-key codes used by this project.

## Project-Specific Keys

| Use | Key | Virtual-Key Code |
|-----|-----|------------------|
| Auto-shot reconfigure | `F1` | `0x70` |
| Auto-shot pause/resume | `F13` | `0x7C` |
| Default auto-shot trigger | `Mouse 1` | `0x01` |

## Mouse Buttons

| Input | Virtual-Key Code |
|-------|------------------|
| Mouse 1 | `0x01` |
| Mouse 2 | `0x02` |
| Mouse 3 | `0x04` |
| Mouse 4 | `0x05` |
| Mouse 5 | `0x06` |

## Keyboard Keys

### Control Keys

| Key | Virtual-Key Code |
|-----|------------------|
| Backspace | `0x08` |
| Tab | `0x09` |
| Enter | `0x0D` |
| Pause | `0x13` |
| Caps Lock | `0x14` |
| Escape | `0x1B` |
| Space | `0x20` |
| Page Up | `0x21` |
| Page Down | `0x22` |
| End | `0x23` |
| Home | `0x24` |
| Left Arrow | `0x25` |
| Up Arrow | `0x26` |
| Right Arrow | `0x27` |
| Down Arrow | `0x28` |
| Insert | `0x2D` |
| Delete | `0x2E` |

### Number Keys

| Key | Virtual-Key Code |
|-----|------------------|
| 0 | `0x30` |
| 1 | `0x31` |
| 2 | `0x32` |
| 3 | `0x33` |
| 4 | `0x34` |
| 5 | `0x35` |
| 6 | `0x36` |
| 7 | `0x37` |
| 8 | `0x38` |
| 9 | `0x39` |

### Letter Keys

| Key | Virtual-Key Code |
|-----|------------------|
| A | `0x41` |
| B | `0x42` |
| C | `0x43` |
| D | `0x44` |
| E | `0x45` |
| F | `0x46` |
| G | `0x47` |
| H | `0x48` |
| I | `0x49` |
| J | `0x4A` |
| K | `0x4B` |
| L | `0x4C` |
| M | `0x4D` |
| N | `0x4E` |
| O | `0x4F` |
| P | `0x50` |
| Q | `0x51` |
| R | `0x52` |
| S | `0x53` |
| T | `0x54` |
| U | `0x55` |
| V | `0x56` |
| W | `0x57` |
| X | `0x58` |
| Y | `0x59` |
| Z | `0x5A` |

### Numpad Keys

| Key | Virtual-Key Code |
|-----|------------------|
| Numpad 0 | `0x60` |
| Numpad 1 | `0x61` |
| Numpad 2 | `0x62` |
| Numpad 3 | `0x63` |
| Numpad 4 | `0x64` |
| Numpad 5 | `0x65` |
| Numpad 6 | `0x66` |
| Numpad 7 | `0x67` |
| Numpad 8 | `0x68` |
| Numpad 9 | `0x69` |
| Numpad * | `0x6A` |
| Numpad + | `0x6B` |
| Numpad - | `0x6D` |
| Numpad . | `0x6E` |
| Numpad / | `0x6F` |

### Function Keys

| Key | Virtual-Key Code |
|-----|------------------|
| F1 | `0x70` |
| F2 | `0x71` |
| F3 | `0x72` |
| F4 | `0x73` |
| F5 | `0x74` |
| F6 | `0x75` |
| F7 | `0x76` |
| F8 | `0x77` |
| F9 | `0x78` |
| F10 | `0x79` |
| F11 | `0x7A` |
| F12 | `0x7B` |
| F13 | `0x7C` |

### Modifier Keys

| Key | Virtual-Key Code |
|-----|------------------|
| Left Shift | `0xA0` |
| Right Shift | `0xA1` |
| Left Ctrl | `0xA2` |
| Right Ctrl | `0xA3` |
| Left Alt | `0xA4` |
| Right Alt | `0xA5` |

### Punctuation Keys

| Key | Virtual-Key Code |
|-----|------------------|
| ; | `0xBA` |
| = | `0xBB` |
| , | `0xBC` |
| - | `0xBD` |
| . | `0xBE` |
| / | `0xBF` |
| ` | `0xC0` |
| [ | `0xDB` |
| \ | `0xDC` |
| ] | `0xDD` |
| ' | `0xDE` |

## Notes

- `KeyNames` and `MouseButtonNames` in `internal/keyboard/keyboard.go` is used as the selection map for the project.