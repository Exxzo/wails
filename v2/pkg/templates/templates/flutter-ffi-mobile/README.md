# Flutter FFI Mobile + Wails

This template creates a Flutter Mobile app (Android/iOS) with FFI bindings to a Go shared library.

## Architecture

- **Backend**: Go shared library built with `-buildmode=c-shared`
- **Frontend**: Flutter mobile app loads the shared library via `dart:ffi`

## Dev

- `wails dev` will:
  - automatically run `dart run ffigen` to generate Dart FFI bindings
  - rebuild the Go shared library on changes
  - touch `frontend/lib/reload.trigger` to hint Flutter to reload
  - run `flutter run` to launch the Flutter mobile app

## Build

- `wails build` will create the shared library in `build/bin/`.
- `wails build --platform android` builds Android APK with shared library
- `wails build --platform ios` builds iOS app with shared library
- Build your Flutter app separately (`flutter build apk` or `flutter build ios`), bundling the `.so/.dylib/.dll` as needed per platform.

## Prerequisites

Install required tools:
```bash
# Install LLVM/clang (required by ffigen)
brew install llvm  # macOS
# or apt-get install clang on Ubuntu

# Install ffigen in your Flutter project
dart pub add --dev ffigen

# For Android development
flutter doctor --android-licenses

# For iOS development (macOS only)
sudo xcode-select --switch /Applications/Xcode.app/Contents/Developer
```

## Configuration

- FFI bindings are configured in `frontend/ffigen.yaml`
- Generated bindings go to `frontend/lib/src/ffi/wails_ffi_bindings.dart`
- Adjust the header path in `ffigen.yaml` to match your generated C header location

## Notes

- Ensure the Flutter app watches `lib/reload.trigger` to trigger hot-reload (simple file watcher).
- Use the generated Dart bindings to load symbols from the library and marshal data.
- The template includes `ffi` package dependency for FFI operations.
- For mobile development, ensure Android Studio/Xcode is properly configured
