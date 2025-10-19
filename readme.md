# 🧠 GoHabits
**Gamified Habit Tracker API (Go + Fiber + GORM + PostgreSQL + JWT + RabbitMQ Ready)**

GoHabits is a backend API built with **Golang (Fiber)** using **Domain-Driven Design (DDD)** principles.
It’s designed as a **habit tracker with a rewarding system**, where users can earn points for completing daily habits and compete on leaderboards.

---

## 🌍 Deskripsi (Bahasa Indonesia)
**GoHabits** adalah proyek backend API berbasis **Golang (Fiber)** dengan konsep **Domain-Driven Design (DDD)**.
Tujuannya adalah membuat **aplikasi pelacak kebiasaan (habit tracker)** yang memiliki sistem **reward dan leaderboard** agar pengguna termotivasi untuk konsisten.

Aplikasi ini dirancang agar mudah dikembangkan dengan fitur tambahan seperti notifikasi asinkron (RabbitMQ), dashboard admin, dan integrasi frontend (Vue/Nuxt).

---

## 🚀 Features
- 🧩 **Clean Architecture (DDD)** — pemisahan domain, application, dan infrastructure yang jelas.
- 🔐 **JWT Authentication** — endpoint register, login, dan `/auth/me`.
- 🗃️ **PostgreSQL + GORM** — ORM andal dan fleksibel.
- ⚙️ **Config Loader** — environment berbasis `.env` dengan dependency injection.
- 📨 **RabbitMQ Ready** — disiapkan untuk event async seperti habit completion dan reward.
- 🧰 **Modular Codebase** — mudah dikembangkan untuk fitur tambahan.

---

## 🧱 Tech Stack
| Layer | Library |
|-------|----------|
| **Language** | Go 1.22+ |
| **Framework** | Fiber |
| **ORM** | GORM |
| **Database** | PostgreSQL |
| **Auth** | JWT |
| **Architecture** | Domain-Driven Design (DDD) |

---

## 📂 Folder Structure
```bash
internal/
 ├─ domain/        # Entities & repositories
 ├─ application/   # Services, DTOs, business logic
 ├─ interfaces/    # HTTP handlers & routes
 └─ infra/         # Config, database, middleware
