# Final Project Go Batch 56

## Struktur Folder

```shell
final-project-golang
├───config
├───controllers
├───database
│   └───migrations
├───repositories
├───routes
├───structs
└───utils
```

## Endpoint untuk Manajemen Kost

Berikut adalah endpoint untuk menggunakan aplikasi ini

BASE URL (local): `http://localhost:8080`

| Path                   | Method | Keterangan                               |
| ---------------------- | ------ | ---------------------------------------- |
| `/lodgers`             | GET    | Mengambil semua data penghuni kost       |
| `/lodgers`             | POST   | Menambah data penghuni kost              |
| `/lodgers/:id`         | PUT    | Memperbarui data penghuni kost           |
| `/lodgers/:id`         | DELETE | Menghapus data penghuni kost             |
| `/rooms`               | GET    | Mengambil semua data kamar kost          |
| `/rooms`               | POST   | Menambah data kamar kost                 |
| `/rooms/:id`           | PUT    | Memperbarui data kamar kost              |
| `/rooms/:id`           | DELETE | Menghapus data kamar kost                |
| `/rents`               | GET    | Mengambil semua data info sewa           |
| `/rents`               | POST   | Menambah data info sewa                  |
| `/rents/:id`           | DELETE | Menghapus data info sewa                 |
| `/transaction/duedate` | GET    | Mengambil data transaksi kost (due date) |
| `/transaction`         | POST   | Menambah data transaksi                  |
| `/transaction/:id`     | DELETE | Menghapus data transaksi                 |
