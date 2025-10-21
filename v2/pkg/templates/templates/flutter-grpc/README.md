# Flutter gRPC + Wails

This template creates a Flutter app with a gRPC backend server. Works for desktop, mobile, and web platforms.

## Architecture

- **Backend**: Go gRPC server (localhost:50051) with bound methods
- **Frontend**: Flutter app connects via gRPC client to call backend methods

## Dev

- `wails dev` will:
  - automatically run `protoc` to generate Go and Dart gRPC code
  - build and start the Go gRPC server on `127.0.0.1:50051`
  - run `flutter run -d linux --debug` for the Flutter app
  - rebuild Go server on changes, restart automatically

## Build

- `wails build` creates the gRPC server binary
- Flutter app should be built separately: `flutter build linux/windows/macos`

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
```

## Notes

- Ensure Flutter SDK is installed and `flutter pub get` works
- Flutter app needs `grpc` and `protobuf` packages (included in template)
- Server listens on `127.0.0.1:50051` by default
- Proto files should be placed in `proto/` directory
- Generated Go code goes to project root, Dart code to `frontend/lib/src/generated/`
