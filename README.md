# **Final Project SpecialAcademy \- Level Up Your Golang Skills**

## Simple Microservice Merchant API

Project ini adalah implementasi microservice sederhana untuk pengelolaan merchant menggunakan Golang dan Fiber. API ini mendukung operasi CRUD merchant, akun, dan transaksi, serta integrasi dengan database, message queue, dan autentikasi JWT.

## Latar Belakang
Sistem ini dibuat untuk mendemonstrasikan arsitektur microservice, pengelolaan data merchant, serta integrasi dengan beberapa komponen seperti database, message queue, dan autentikasi.

## System Design Architecture
Lihat detail arsitektur pada `docs/architecture/project-structure.md`.

## Database Schema Design
Skema database dapat ditemukan pada folder `database/migration/`.

## Tech Stacks
- Golang & Fiber (Web Framework)
- GORM (ORM)
- MongoDB, MySQL, PostgreSQL
- RabbitMQ
- Docker & Docker Compose
- Swagger (API Documentation)

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