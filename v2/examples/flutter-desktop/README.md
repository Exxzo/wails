# Example: Flutter Desktop + Wails (gRPC IPC)

This example demonstrates a Flutter desktop app communicating with a Wails Go backend via gRPC.

Dev:
- `wails dev` starts the Go gRPC server and launches `flutter run -d desktop`.
- Flutter reconnects to `127.0.0.1:50051` on server restart.

Build:
- `wails build` builds the Go server and the Flutter desktop binary.


