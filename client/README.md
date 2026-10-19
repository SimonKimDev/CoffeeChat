# Client

The CoffeeChat desktop app: **Wails v2** (Go backend) with a **React + TypeScript** frontend.

The app doesn't exist yet. You'll scaffold it in milestone **M8**:

```sh
go install github.com/wailsapp/wails/v2/cmd/wails@latest
wails doctor
wails init -n coffeechat -t react-ts -d .   # run from inside client/
wails dev
```

Planned layout once it exists:

```
client/
  main.go, app.go     Wails entry point and Go methods bound to the UI
  internal/
    crypto/           libsignal wrapper (M10)
    localstore/       passphrase-encrypted SQLite (M9)
    relay/            WebSocket client (M10)
  frontend/           React + TypeScript (Vite)
```
