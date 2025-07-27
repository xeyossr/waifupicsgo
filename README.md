# waifupicsgo 🖼️

A simple Go client for waifu.pics — fetch random SFW/NSFW anime images effortlessly!

## Install

```bash
go get github.com/xeyossr/waifupicsgo
```

## Usage

```go
import "github.com/xeyossr/waifupicsgo"

// SFW examples
img, _ := waifupicsgo.SFW.GetOne(waifupicsgo.SFWHug.String())
imgs, _ := waifupicsgo.SFW.GetMany(waifupicsgo.SFWHug.String())

// NSFW examples
img, _ := waifupicsgo.NSFW.GetOne(waifupicsgo.NSFWNeko.String())
imgs, _ := waifupicsgo.NSFW.GetMany(waifupicsgo.NSFWNeko.String())
```

### 🔹 `GetOne(category string)`

Fetches a single random image from the given category.

### 🔹 `GetMany(category string)`

Fetches multiple unique images from the given category.

## Categories
- `SFW`: waifu, neko, shinobu, megumin, bully, cuddle, cry, hug, awooo, kiss, lick, pat, smug, bonk, yeet, blush, smile, wave, highfive, handhold, nom, bite, glomp, slap, kill, kick, happy, wink, poke, dance, cringe

- `NSFW`: waifu, neko, trap, blowjob

Each category is defined as a typed constant and can be passed using `.String()`.

## License

This project is licensed under the [GNU General Public License v3.0 (GPL-3.0)](https://opensource.org/licenses/GPL-3.0).