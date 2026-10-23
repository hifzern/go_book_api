# Book App

API buku dalam Go + klien desktop WPF, plus materi belajar Go web dari repo `go-web-services`.

| Path | Isi | Stack |
|---|---|---|
| [`api/`](./api), [`cmd/`](./cmd), [`tests/`](./tests) | REST API buku (versi tanpa auth) | Go, Gin, GORM |
| [`client/`](./client) | Konsumen REST API (desktop) | C# / WPF |
| [`learn/`](./learn) | 7 bab latihan Go web: hello world → Gin → Book API | Go |

## Menjalankan API

```bash
go run ./cmd
go test ./...
```

## Klien

Buka `client/bookAppsClient.sln` di Visual Studio, lalu jalankan `bookAppsClient`.

## Catatan `learn/`

`learn/07-Go-Book-Api` adalah **versi lain** dari API di root: punya JWT auth
(`POST /token`, `JWTAuthMiddleware`) yang belum ada di root, tapi belum punya
klien WPF. Modul Go-nya terpisah (`learn/07-Go-Book-Api/go.mod`), jadi
`go build ./...` di root tidak menyentuhnya — jalankan dari direktori itu sendiri.
