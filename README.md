# Metadata Extractor

## Gambaran Umum Sistem

Aplikasi ini menyediakan layanan ekstraksi metadata dari file gambar melalui antarmuka REST API. Pengguna dapat mengunggah file gambar dan menerima respons JSON yang berisi informasi EXIF serta metadata dasar dari gambar tersebut.

**Teknologi yang digunakan:**

- **Backend:** Go (Golang) dengan framework Gin
- **Library Metadata:** `imagemeta`, `go-exif/v3` (github.com/dsoprea/go-exif/v3)
- **Frontend:** Vue.js dengan Vite (dideploy terpisah)

---

## Arsitektur & Infrastruktur

```
┌─────────────────────────┐         ┌──────────────────────────────────┐
│  Frontend               │         │  Backend                         │
│  GitHub Pages           │────────▶│  Azure Container Apps            │
│  (Vue + Vite)           │  HTTP   │  (Go + Gin)                      │
└─────────────────────────┘  POST   └──────────────────────────────────┘
```

- **Frontend** dideploy di **GitHub Pages** dan berfungsi sebagai antarmuka pengunggahan file.
- **Backend** dideploy di **Azure Container Apps**, menerima permintaan dari frontend dan memproses ekstraksi metadata.

---

## Konfigurasi Deployment

### Azure Container Apps

Backend berjalan sebagai container di Azure Container Apps. Gunakan connection string / environment variable berikut untuk konfigurasi:

| Parameter | Nilai |
|---|---|
| Container App URL | `https://metadata-extractor.livelymoss-c733e49a.southeastasia.azurecontainerapps.io` |
| Port yang Diekspos | `8080` |


**Contoh base URL untuk API:**
```
```

### GitHub Pages (Frontend)

Static single page web app menggunka Vue.js, di-deploy menggunakan GitHub Actions

---

## Endpoint API

### `POST /api/upload`

Endpoint utama untuk mengunggah gambar dan menerima metadata yang diekstrak.

**URL:**
```
POST https://https://metadata-extractor.livelymoss-c733e49a.southeastasia.azurecontainerapps.io/api/upload
```

**Content-Type:** `multipart/form-data`

**Parameter Form:**

| Field | Tipe | Wajib | Deskripsi |
|---|---|---|---|
| `file` | File | Ya | File gambar yang akan diekstrak metadatanya (JPEG, PNG, dll.) |

**Contoh permintaan menggunakan `curl`:**
```bash
curl -X POST "https://https://metadata-extractor.livelymoss-c733e49a.southeastasia.azurecontainerapps.io/api/upload" \
     -F "file=@/path/to/photo.jpg"
```

---

## Format Respons

Respons dikembalikan dalam format JSON dan dikelompokkan ke dalam tiga bagian utama: `basic`, `camera`, dan `location`.

### Respons Berhasil — Data Lengkap (`200 OK`)

```json
{
  "basic": {
    "FileName": "IMG_1234.JPG",
    "FileSize": 2.45,
    "FileType": "jpeg",
    "DateTimeOriginal": "2022:09:15 14:22:10",
    "ImageWidth": 6000,
    "ImageHeight": 4000,
    "Software": "Adobe Photoshop 22.0"
  },
  "camera": {
    "Make": "Canon",
    "Model": "EOS 5D Mark IV",
    "LensInfo": "EF24-70mm f/2.8L II USM",
    "FNumber": 2.8,
    "ExposureTime": "1/250",
    "ISO": 100,
    "FocalLength": 35,
    "CameraSerial": "123456789",
    "LensSerial": "987654321"
  },
  "location": {
    "GPSLatitude": 51.507351,
    "GPSLongitude": -0.127758,
    "GPSAltitude": 15.0,
    "City": "London",
    "Country": "United Kingdom"
  }
}
```

### Respons Berhasil — Data Parsial (`200 OK`)

Apabila sebagian tag EXIF tidak tersedia dalam file gambar (misalnya, tidak ada data GPS), field terkait akan bernilai `null`.

```json
{
  "basic": {
    "FileName": "photo.jpg",
    "FileSize": 0.98,
    "FileType": "png",
    "DateTimeOriginal": null,
    "ImageWidth": 1920,
    "ImageHeight": 1080,
    "Software": null
  },
  "camera": {
    "Make": "Nikon",
    "Model": "D850",
    "LensInfo": null,
    "FNumber": null,
    "ExposureTime": null,
    "ISO": null,
    "FocalLength": null,
    "CameraSerial": null,
    "LensSerial": null
  },
  "location": {
    "GPSLatitude": null,
    "GPSLongitude": null,
    "GPSAltitude": null,
    "City": null,
    "Country": null
  }
}
```

### Referensi Field Respons

**Bagian `basic`**

| Field | Tipe | Deskripsi |
|---|---|---|
| `FileName` | string | Nama file yang diunggah |
| `FileSize` | float | Ukuran file dalam MB |
| `FileType` | string | Tipe/format file (mis. `jpeg`, `png`) |
| `DateTimeOriginal` | string \| null | Waktu pengambilan foto asli |
| `ImageWidth` | integer | Lebar gambar dalam piksel |
| `ImageHeight` | integer | Tinggi gambar dalam piksel |
| `Software` | string \| null | Perangkat lunak yang digunakan untuk memproses gambar |

**Bagian `camera`**

| Field | Tipe | Deskripsi |
|---|---|---|
| `Make` | string \| null | Merek kamera |
| `Model` | string \| null | Model kamera |
| `LensInfo` | string \| null | Informasi lensa yang digunakan |
| `FNumber` | float \| null | Nilai aperture (f-number) |
| `ExposureTime` | string \| null | Kecepatan rana (shutter speed) |
| `ISO` | integer \| null | Sensitivitas ISO |
| `FocalLength` | float \| null | Panjang fokus dalam mm |
| `CameraSerial` | string \| null | Nomor seri kamera |
| `LensSerial` | string \| null | Nomor seri lensa |

**Bagian `location`**

| Field | Tipe | Deskripsi |
|---|---|---|
| `GPSLatitude` | float \| null | Koordinat lintang (latitude) |
| `GPSLongitude` | float \| null | Koordinat bujur (longitude) |
| `GPSAltitude` | float \| null | Ketinggian dalam meter |
| `City` | string \| null | Nama kota (jika tersedia) |
| `Country` | string \| null | Nama negara (jika tersedia) |

### Respons Gagal

**`400 Bad Request`** — Tidak ada file yang diunggah atau permintaan tidak valid:
```json
{
  "error": "pesan error"
}
```

---

## Middleware & Keamanan

### Rate Limiting

Backend menerapkan pembatasan laju permintaan per alamat IP klien menggunakan `rate.NewLimiter`. Konfigurasi saat ini:

```go
rate.NewLimiter(3, 5)
// 3 permintaan per detik, dengan burst maksimum 5 permintaan
```

Konfigurasi ini dapat disesuaikan di `backend/main.go` sesuai kebutuhan beban produksi.

### CORS (Cross-Origin Resource Sharing)

Saat ini CORS dikonfigurasi untuk menerima permintaan dari semua origin:

```go
AllowOrigins: ["*"]
```
---

## Cara Menjalankan Secara Lokal

### Backend (Go)

```bash
cd backend
go mod download
go run .
# Server berjalan di http://localhost:8080
```

### Frontend (Vue + Vite)

```bash
cd frontend
npm install
npm run dev
# Buka URL dev server (biasanya http://localhost:5173)
```

Saat menjalankan secara lokal, pastikan konfigurasi frontend mengarah ke `http://localhost:8080` sebagai base URL backend, bukan ke URL Azure Container Apps.

---
