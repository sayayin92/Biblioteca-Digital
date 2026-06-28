# 📚 Biblioteca Digital en Go

Sistema de gestión de biblioteca desarrollado en **Go (Golang)** utilizando el framework **Gin**, enfocado en la aplicación de conceptos de **Programación Orientada a Objetos (POO)** y arquitectura en capas.

---

# 📌 Datos del grupo

- **Integrante:** YUSI ARISTEGA PEÑAFIEL  
- **Curso:** 3A  
- **Materia:** Programación Orientada a Objetos 1  
- **Fecha de entrega:** 27 de junio de 2026  

---

# 🚀 Descripción del proyecto

Este proyecto es una **API REST** que permite gestionar una biblioteca digital con las siguientes funcionalidades:

- 👤 Gestión de usuarios  
- 📚 Gestión de libros  
- 📖 Gestión de préstamos  

---

# 🧱 Arquitectura del proyecto

Biblioteca-Digital/
│
├── main.go
├── handlers/
│   ├── libros.go
│   ├── usuarios.go
│   └── prestamos.go
├── models/
│   ├── libro.go
│   ├── usuario.go
│   └── prestamo.go
├── database/
│   └── conexion.go
├── routes/
│   └── routes.go
└── go.mod

---

# ⚙️ Tecnologías utilizadas

- Go (Golang)  
- Gin Web Framework  
- API REST  
- Arquitectura en capas  
- Structs (POO en Go)  
- Slices (almacenamiento en memoria)  

---

# ▶️ Instalación y ejecución

## 1. Clonar el repositorio

git clone https://github.com/sayayin92/Biblioteca-Digital.git

## 2. Entrar al proyecto

cd Biblioteca-Digital

## 3. Instalar dependencias

go mod tidy

## 4. Ejecutar el servidor

go run main.go

---

# 🌐 Servidor

http://localhost:8080

---

# 📌 Endpoints de la API

## 👤 Usuarios

POST /usuarios

{
  "nombre": "Juan",
  "email": "juan@gmail.com"
}

GET /usuarios

---

## 📚 Libros

POST /libros

{
  "titulo": "El Quijote",
  "autor": "Cervantes"
}

GET /libros  
GET /libros/:id  

---

## 📖 Préstamos

POST /prestamos

{
  "usuario_id": 1,
  "libro_id": 1
}

GET /prestamos

---

# 🧠 Funcionamiento interno

- Structs para modelar entidades  
- Handlers para lógica de negocio  
- Routes para endpoints  
- Base de datos en memoria (slices)  

---

# ⚠️ Limitaciones

- No utiliza base de datos persistente  
- Los datos se pierden al reiniciar el servidor  
- No incluye autenticación de usuarios  

---

# 🚀 Mejoras futuras

- MySQL o PostgreSQL  
- JWT (autenticación)  
- Validaciones avanzadas  
- Swagger (documentación API)  
- Deploy en la nube  

---

# 👨‍💻 Autor
YUSI ARISTEGA PEÑAFIEL  


---

# 💡 Conclusión

Este proyecto permite comprender los fundamentos del desarrollo backend en Go, aplicando conceptos de Programación Orientada a Objetos, arquitectura en capas y creación de APIs REST, sirviendo como base para proyectos más avanzados.
