# 📰 Blog Platform

A **modern blog platform** with a **powerful backend, an intuitive admin panel, and a dynamic client interface**.  
This project consists of three main applications:  

- **Backend** – A Golang-powered API for handling blog posts, users, and authentication.  
- **CRM** – A React + Vite admin panel for managing content and users.  
- **Client** – A Next.js + Vite frontend for reading and interacting with blog content.  

---

## **Project Description**

This project is designed to provide a **fast, scalable, and user-friendly** blogging experience.  
It follows a **microservices-like approach**, separating **backend, admin panel, and client** into distinct services.

### **Technologies Used**
- **Backend:** Golang, Echo, PostgreSQL  
- **CRM:** React, Vite, TypeScript  
- **Client:** Next.js, Vite, TypeScript  
- **Deployment:** Docker, Kubernetes (planned)  

### **Challenges & Future Improvements**
- Implement **GraphQL API** for optimized data fetching.  
- Improve **SEO & performance** for the blog.  
- Add **full authentication & authorization** using JWT.  

---

## **Table of Contents**
- [Installation](#installation)
- [Usage](#how-to-use-the-project)
- [Project Structure](#project-structure)
- [Contributing](#contributing)

---

## **Installation**

To run this project locally, follow these steps:

### **1. Clone the repository**
```sh
git clone https://github.com/nnniyaz/blog.git
cd blog
```

### **2. Backend Setup (Golang)**
```sh
cd backend
make run
```
Runs on: http://localhost:8080

### **3. CRM Setup (React + Vite + Typescript)**
```sh
cd crm
npm install
npm run dev
```
Runs on: http://localhost:3000

### **4. Client Setup (Next.js + Vite + Typescript)**
```sh
cd client
npm install
npm run dev
```
Runs on: http://localhost:3001

---

## **How to Use the Project**
Once all services are running:  
- **Visit `http://localhost:3001`** to browse and read blog posts.  
- **Log into `http://localhost:3000`** to manage blog content.  
- **Use the API at `http://localhost:8080`** for backend operations.  

#### **Authentication**
- Default admin credentials:  
  - **Username:** `admin`  
  - **Password:** `password123` (change it in `.env`!)  

---

## **Project Structure**
```
blog/
│── backend/      # Golang API
│── crm/          # Admin panel (React + Vite)
│── client/       # Blog frontend (Next.js + Vite)
│── README.md     # Documentation
```

---

## **Contributing**
Want to contribute? Feel free to fork the repo and submit a PR!

**Anything you want to add or modify?** Let me know!
