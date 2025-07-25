# **Final Project SpecialAcademy \- Level Up Your Golang Skills**

## Simple Merchant API

Project ini adalah implementasi sederhana untuk pengelolaan merchant menggunakan Golang boilerplate (GoFiber). API ini mendukung operasi CRUD merchant, account, dan transaction, serta integrasi dengan database(MySQL, Redis, MongoDB), message queue(RabbitMQ).

## Latar Belakang

Sistem ini dibuat untuk mendemonstrasikan arsitektur sederhana merchant api (pengelolaan data merchant, generate credential account, generate qris, dan create transaction), serta integrasi dengan beberapa komponen seperti database, message queue, dan autentikasi dengan signature.

## System Design Architecture

Flow Merchant API
![Flow Merchant API](./docs/architecture/Flow%20Merchant%20API%20.png)

Merchant API Design
![API Design](./docs/architecture/Merchant%20API%20Design%20.png)

Detail arsitektur folder pada `docs/architecture/project-structure.md`.

## Database Schema Design

![Database Diagram](./docs/architecture/DB%20Diagram.png)
Skema database dapat ditemukan pada folder `database/migration/`.

## Tech Stacks

- GoFiber (Web Framework)
- GORM (ORM)
- MongoDB
- MySQL
- Redis
- RabbitMQ
- Docker & Docker Compose
- Swagger (API Documentation)
- Go Migrate (Database Migration)

## Cara Menjalankan

1. Clone repository ini
2. Jalankan perintah berikut untuk build dan menjalankan service menggunakan Docker Compose:
   ```bash
   docker-compose up --build
   ```
3. Dokumentasi API tersedia di `/docs/swagger.yaml` dan dapat diakses melalui Swagger UI.

## Konfigurasi

Konfigurasi database, message queue, dan environment dapat diatur pada folder `config/` dan file `docker-compose.yaml`.

## Kontribusi

Silakan buat pull request atau issue untuk perbaikan dan pengembangan lebih lanjut.

## Lisensi

MIT License
