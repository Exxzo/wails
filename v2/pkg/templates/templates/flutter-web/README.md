# Flutter Web + Wails

This template wires a Flutter Web app into Wails' native window and dev workflow.

## Dev

- `wails dev` will:
  - run `flutter run -d web-server --web-port=34115 --web-hostname=127.0.0.1`
  - proxy the dev server into the Wails window (URL autodetected from stdout)
  - rebuild and restart Go backend on changes
  - inject `frontend/wails_bridge.js` so Flutter Web can call Wails runtime

## Build

- `wails build` runs `flutter build web` and embeds `frontend/build/web` into the app binary.

## Notes

     - Ensure Flutter SDK is installed and `flutter pub get` works in `frontend/`
