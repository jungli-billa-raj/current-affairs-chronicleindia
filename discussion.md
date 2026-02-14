### Why Go fits *exactly*

* **Static HTML** → `net/http` + `goquery`
* **Single request** → simple, predictable
* **In-memory** → structs + slices
* **List → detail view** → state machine (clean in Go)
* **Single binary** → `go build` (drop into `/usr/bin`)
* **Linux first, Windows easy** → `GOOS=windows GOARCH=amd64 go build`

---

### Minimal stack (don’t overthink)

* **TUI**: `bubbletea`
* **Styling**: `lipgloss`
* **Scraping**: `net/http` + `goquery`
* **CLI args**: `flag` or `cobra` (flag is enough)

---

### Mental model (keep it simple)

* `Model` = current view + article list + selected article
* `Update` = key presses (`j/k/enter/q`)
* `View` = list view OR article view

Think **vim-like navigation**, not menus.

---

### Skeleton layout

```
cmd/
  app/
    main.go
internal/
  tui/
    model.go
    update.go
    view.go
  scrape/
    fetch.go
    parse.go
```

---

### One non-obvious tip

**Normalize scraped text early**
(strip whitespace, fix line breaks) → TUIs look ugly otherwise.

---

If you want next:

* a **20-line Bubble Tea starter**
* keybindings design
* or a **clean article rendering strategy (paging, wrapping)**
