# Flutter FFI + Wails

This template builds a Go library using `-buildmode=c-shared` and loads it in a Flutter app via `dart:ffi`. Works for desktop, mobile, and web platforms.

## Dev

- `wails dev` will:
  - automatically run `dart run ffigen` to generate Dart FFI bindings
  - rebuild the Go shared library on changes
  - touch `frontend/lib/reload.trigger` to hint Flutter to reload
  - run `flutter run` to launch the Flutter app

## Build

- `wails build` will create the shared library in `build/bin/`.
- Build your Flutter app separately (`flutter build <platform>`), bundling the `.so/.dylib/.dll` as needed per platform.

## Prerequisites

Install required tools:
```bash
# Install LLVM/clang (required by ffigen)
brew install llvm  # macOS
# or apt-get install clang on Ubuntu

# Install ffigen in your Flutter project
dart pub add --dev ffigen
```

## Configuration

- FFI bindings are configured in `frontend/ffigen.yaml`
- Generated bindings go to `frontend/lib/src/ffi/wails_ffi_bindings.dart`
- Adjust the header path in `ffigen.yaml` to match your generated C header location

## Notes

- Ensure the Flutter app watches `lib/reload.trigger` to trigger hot-reload (simple file watcher).
- Use the generated Dart bindings to load symbols from the library and marshal data.
- The template includes `ffi` package dependency for FFI operations.
