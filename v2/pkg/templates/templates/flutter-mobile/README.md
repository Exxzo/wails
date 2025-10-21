# Flutter Mobile + Wails

This template creates a Flutter Mobile app (Android/iOS) with a gRPC backend server. You can also use FFI instead of gRPC by changing the outputType.

## Architecture

- **Backend**: Go gRPC server (localhost:50051) with bound methods
- **Frontend**: Flutter mobile app connects via gRPC client to call backend methods

## Dev

- `wails dev` will:
  - automatically run `protoc` to generate Go and Dart gRPC code
  - build and start the Go gRPC server on `127.0.0.1:50051`
  - run `flutter run` for the Flutter mobile app
  - rebuild Go server on changes, restart automatically

## Build

- `wails build` creates the gRPC server binary
- `wails build --platform android` builds Android APK
- `wails build --platform ios` builds iOS app
- Flutter app should be built separately: `flutter build apk` or `flutter build ios`

## Prerequisites

Install required tools:
```bash
# Install protoc
brew install protobuf  # macOS
# or download from https://github.com/protocolbuffers/protobuf/releases

# Install Go protoc plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install Dart protoc plugin
dart pub global activate protoc_plugin
# Add ~/.pub-cache/bin to your PATH

# For Android development
flutter doctor --android-licenses

# For iOS development (macOS only)
sudo xcode-select --switch /Applications/Xcode.app/Contents/Developer
```

## Notes

- Ensure Flutter SDK is installed and `flutter pub get` works
- Flutter app needs `grpc` and `protobuf` packages (included in template)
- Server listens on `127.0.0.1:50051` by default
- Proto files should be placed in `proto/` directory
- Generated Go code goes to project root, Dart code to `frontend/lib/src/generated/`
- For mobile development, ensure Android Studio/Xcode is properly configured
