# Flutter gRPC + Wails

This template creates a Flutter Desktop/Mobile app with a gRPC backend server.

## Architecture

- **Backend**: Go gRPC server (localhost:50051) with bound methods
- **Frontend**: Flutter app connects via gRPC client to call backend methods

## Dev

- `wails dev` will:
  - build and start the Go gRPC server on `127.0.0.1:50051`
  - run `flutter run -d linux --debug` for the Flutter app
  - rebuild Go server on changes, restart automatically

## Build

- `wails build` creates the gRPC server binary
- Flutter app should be built separately: `flutter build linux/windows/macos`

## Notes

- Ensure Flutter SDK is installed and `flutter pub get` works
- Flutter app needs `wails_dart_grpc` package to connect to the server
- Server listens on `127.0.0.1:50051` by default
