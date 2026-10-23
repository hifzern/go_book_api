# Book App

API buku dalam Go + klien desktop WPF, digabung jadi satu repo.

| Path | Isi | Stack |
|---|---|---|
| [`api/`](./api), [`cmd/`](./cmd), [`tests/`](./tests) | REST API buku | Go |
| [`client/`](./client) | Konsumen REST API (desktop) | C# / WPF |

## Menjalankan API

```bash
go run ./cmd
go test ./...
```

## Klien

Buka `client/bookAppsClient.sln` di Visual Studio, lalu jalankan `bookAppsClient`.
