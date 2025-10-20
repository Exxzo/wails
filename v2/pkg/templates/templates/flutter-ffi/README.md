# Flutter FFI + Wails

This template builds a Go library using `-buildmode=c-shared` and loads it in a Flutter app via `dart:ffi`.

## Dev

- `wails dev` will:
  - rebuild the Go shared library on changes
  - touch `frontend/lib/reload.trigger` to hint Flutter to reload
  - run `flutter run` to launch the Flutter app

## Build

- `wails build` will create the shared library in `build/bin/`.
- Build your Flutter app separately (`flutter build <platform>`), bundling the `.so/.dylib/.dll` as needed per platform.

## Notes

- Ensure the Flutter app watches `lib/reload.trigger` to trigger hot-reload (simple file watcher).
- Use a Dart package (e.g. `wails_dart_ffi`) to load symbols from the library and marshal data.
