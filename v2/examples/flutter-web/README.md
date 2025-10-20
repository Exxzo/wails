# Example: Flutter Web + Wails

This example demonstrates running a Flutter Web app inside Wails' window with the standard dev workflow.

Dev:

- `wails dev` runs Flutter web-server and proxies it into the Wails window.
- Go backend restarts on change; Flutter stays connected.

Build:

- `wails build` runs `flutter build web` and embeds the built assets.
