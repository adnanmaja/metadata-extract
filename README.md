# Simple Metadata Extraction App

## Project Overview
- **Purpose:** Upload an image and receive extracted EXIF and basic image information as JSON.
- **Stack:** Go backend using Gin, `imagemeta` and `go-exif`; Vue + Vite frontend.

## Repository Structure
- `backend/` – Go server and API handlers
- `backend/main.go` – server setup, CORS and rate-limiting middleware
- `backend/api/post.go` – upload handler (`/api/upload`)
- `backend/services/metadata.go` – metadata extraction helpers
- `frontend/` – Vue app (Vite) providing the upload UI

## How it works
- The frontend uploads a file to `POST http://localhost:8080/api/upload`.
- The backend saves the upload to a temporary file, then:
  - decodes image metadata with `imagemeta` → `exif2.Exif`
  - extracts flattened EXIF tags with `github.com/dsoprea/go-exif/v3`
  - reads image width/height via `image.DecodeConfig`
- The server responds with JSON grouped into `basic`, `camera`, and `location` sections.

## Example JSON Response (abbreviated)
```
{
  "basic": { "FileName": "photo.jpg", "FileSize": 1.23, "FileType": "jpeg", "DateTimeOriginal": "...", "ImageWidth": 4000, "ImageHeight": 3000 },
  "camera": { "Make": "Canon", "Model": "EOS 80D", "FNumber": 2.8, "ISO": 100 },
  "location": { "GPSLatitude": 51.5074, "GPSLongitude": null, "GPSAltitude": 15 }
}
```

## Running Locally

Backend (Go):
```bash
cd backend
go mod download
go run .
# server listens on :8080
```

Frontend (Vue + Vite):
```bash
cd frontend
npm install
npm run dev
# open the dev server URL (usually http://localhost:5173)
```

## API Usage
- Upload with `curl`:
```bash
curl -X POST "http://localhost:8080/api/upload" -F "file=@/path/to/photo.jpg"
```

## Notes & Caveats
- Rate limiting is per-client IP using `rate.NewLimiter(3, 5)` in `backend/main.go`.
- CORS is open (`AllowOrigins: ["*"]`) — tighten in production.
- Some fields (e.g., longitude, city) may be `null` if not available.
- Error handling in helper functions is conservative; consider propagating errors for clearer API error responses.

## Suggestions / Next Steps
- Surface extraction errors as HTTP error responses from the upload handler.
- Return full GPS coordinates when available.
- Add basic integration tests for the `POST /api/upload` flow.

---
Created from the code in this repository; see `backend/` and `frontend/` for implementation.

## Example API Response Structures

Below are more explicit example responses the backend may return.

- Full successful response:

```
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

- Partial response when GPS or some tags are missing (note `null` values):

```
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

- Error response example (bad request / no file):

```
HTTP/1.1 400 Bad Request
{
  "error": "something"
}
```

These examples reflect how `backend/api/post.go` currently shapes responses; field availability depends on the uploaded image's EXIF data.
