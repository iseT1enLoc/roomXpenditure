# 🏠 RoomXpenditure - Student Dormitory Management Backend

RoomXpenditure is a backend system built with **Go** and the **Gin web framework**, designed to efficiently manage student dormitory rooms and their expenditures. It supports core features like room allocation, student profiles, expenditure tracking, and role-based access.

---

## 📌 Features

- 🧑‍🎓 Student registration & management  
- 🏡 Room listing, assignment, and availability tracking  
- 💰 Monthly expenditure recording  
- 🔐 Role-based access control (Admin, Staff, Student)  
- 📦 Clean, modular architecture  
- 🐳 Docker support for easy deployment  

---

## 🧱 Tech Stack

| Component       | Technology      |
|----------------|-----------------|
| Language        | Go (Golang)     |
| Web Framework   | Gin             |
| ORM             | GORM            |
| Database        | PostgreSQL      |
| Containerization| Docker          |
| Config Format   | YAML            |

---

## 📁 Project Structure



roomXpenditure/ │ ├── api/ # API endpoints and route handlers ├── appcontext/ # Application-wide context (DI, configs) ├── config/ # App configurations ├── models/ # Struct definitions for DB tables ├── repository/ # Repositories for DB operations ├── services/ # Business logic layer ├── utils/ # Utilities and helpers │ ├── tables.sql # SQL schema to initialize DB ├── docker-compose.yml ├── command.txt # Commands for app usage ├── go.mod ├── go.sum └── main.go # Application entry point


---

## 🛠️ Setup & Installation

### 1. Clone the repository
```bash
git clone https://github.com/iseT1enLoc/roomXpenditure.git
cd roomXpenditure
### 2. Docker setup
```bash
sudo docker-compose up

### 2. Run application
go run main.go

---

## 🙏 Acknowledgments

This project was developed as part of a backend development initiative for student room and expenditure management. Special thanks to:

- **My mentors and instructors** for their guidance and support.
- **Open-source Go and Gin communities** for excellent libraries and documentation.
- **All contributors** who helped test, debug, and improve the system.

---

## 👨‍💻 Author

**iseT1enLoc**  
📧 [Email me](locnvt.it.com)  
🔗 [GitHub Profile](https://github.com/iseT1enLoc)  
🎓 Student | Backend Developer | Tech Enthusiast

---
