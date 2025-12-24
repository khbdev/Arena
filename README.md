

# 🧩 ARENA — Telegram-Based Online Test Platform (Microservices)

## 🎯 Project Overview

ARENA is a **Telegram-based online testing platform** built using a **microservices architecture**.
The system is designed to generate tests using AI, deliver them to users via Telegram, process answers asynchronously, and provide automated results and statistics.

All services are **independently deployable**, containerized with Docker, and communicate using **gRPC and RabbitMQ**, ensuring scalability, fault isolation, and high performance.

---

## ⚙️ Architecture Overview

* **Gateway Service** acts as the entry point for all Telegram bot interactions.
* **User Service** manages users and roles via gRPC.
* **AI Service** generates tests asynchronously using message queues.
* **Results Service** processes user answers and calculates scores.
* **Notification Service** delivers results and statistics to users and teachers.
* **Test Section Service** controls test visibility and correctness logic.

The system heavily relies on **RabbitMQ for async jobs**, **Redis for caching and temporary storage**, and **gRPC for internal service communication**.

---

## 🧱 Microservices

### 1️⃣ Gateway Service (Telegram Entry Point)

* Handles all Telegram bot requests
* Routes messages to internal services
* Implements Clean Architecture
* Acts as the system entry point

**🧰 Tech Stack:**
Go · Telegram Bot API · Clean Architecture · Docker

---

### 2️⃣ User Service

* Manages user data using Telegram ID
* Independent service
* Provides gRPC APIs for other services
* Built using Clean Architecture

**🧰 Tech Stack:**
Go · gRPC Server · MySQL · Clean Architecture

---

### 3️⃣ AI Service

* Fully asynchronous test generation service
* Listens to RabbitMQ prompt queues
* Uses buffered channels (capacity: 50)
* Runs 5 goroutines for concurrent processing
* Stores AI-generated tests in Redis with TTL = 2 hours
* Pushes Redis keys back to queues for downstream services

**Key Characteristics:**

* High-load capable
* Non-blocking
* Independent from other services

**🧰 Tech Stack:**
Go · RabbitMQ · Redis · Goroutines & Channels · Layered Architecture

---

### 4️⃣ Results Service

* Consumes user answers from RabbitMQ
* Fetches AI-generated tests from Redis using keys
* Compares answers and calculates results
* Pushes user statistics to message queues

**🧰 Tech Stack:**
Go · RabbitMQ · Redis · Async Processing

---

### 5️⃣ Notification Service

* Listens to two RabbitMQ queues:

  * AI notification queue
  * User statistics queue
* Sends results to:

  * Students
  * Teachers
* Uses goroutines and channels for concurrency

**🧰 Tech Stack:**
Go · RabbitMQ · Concurrent Workers

---

### 6️⃣ Test Section Service

* Handles test section validation
* Uses gRPC communication
* Logic based on Telegram ID:

  * Students receive `correct = null`
  * Teachers receive full correct answers
* Reads test data from Redis

**🧰 Tech Stack:**
Go · gRPC Server · Redis

---

## 🔄 Communication Flow

* **Synchronous:** gRPC (internal services)
* **Asynchronous:** RabbitMQ (AI generation, results, notifications)
* **Cache & Temp Storage:** Redis

---

## 🐳 Deployment & Infrastructure

* Each service has its own `Dockerfile`
* All services are orchestrated using `docker-compose`
* Deployed on a VPS server
* Configuration handled via `.env` files

```bash
docker-compose up -d --build
```

---

## 🧰 Technology Stack

* **Language:** Go
* **Architecture:** Microservices, Clean Architecture, Layered Architecture
* **Communication:** gRPC, RabbitMQ
* **Database:** MySQL
* **Cache:** Redis
* **Infrastructure:** Docker, Docker Compose
* **Bot Platform:** Telegram Bot API


---

## 👤 Author

Azizbek Xasanov
Telegram Bot: **@arena_rep_bot**

---

