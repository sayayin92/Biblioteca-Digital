# 📚 Biblioteca Digital en Go

## 📌 Datos del grupo

- **Integrante:** YUSI ARISTEGA PEÑAFIEL
- **Curso:** 3A
- **Materia:** Programación Orientada a Objetos 1

---

## 🎯 Objetivo del programa

Desarrollar una API REST para gestionar una biblioteca digital, aplicando los conceptos aprendidos durante el curso de Programación Orientada a Objetos I, como el uso de estructuras (`structs`), modularidad, manejo de datos en memoria, arquitectura en capas y desarrollo de servicios web utilizando Go y el framework Gin.

---

## ⚙️ Principales funcionalidades

La API permite realizar las siguientes operaciones:

### 👤 Gestión de usuarios
- Registrar usuarios.
- Listar usuarios registrados.
- Actualizar información de un usuario.
- Eliminar usuarios.

### 📚 Gestión de libros
- Registrar libros.
- Listar libros disponibles.
- Buscar un libro por su identificador.
- Actualizar información de un libro.
- Eliminar libros.

### 📖 Gestión de préstamos
- Registrar préstamos de libros.
- Consultar los préstamos registrados.

---

## 🛠 Tecnologías utilizadas

- Go (Golang)
- Framework Gin
- API REST
- JSON
- Arquitectura en capas
- Base de datos en memoria mediante slices

---

## 📂 Estructura del proyecto

```text
Biblioteca-Digital/
│
├── main.go
├── handlers/
├── models/
├── database/
├── routes/
└── go.mod
```

---

## ▶️ Ejecución

1. Clonar el repositorio:

```bash
git clone https://github.com/sayayin92/Biblioteca-Digital.git
```

2. Ingresar al proyecto:

```bash
cd Biblioteca-Digital
```

3. Instalar dependencias:

```bash
go mod tidy
```

4. Ejecutar el servidor:

```bash
go run main.go
```

La API estará disponible en:

```
http://localhost:8080
```

---

## 📅 Fecha

**27 de junio de 2026**

---

## 👨‍💻 Autor

**YUSI ARISTEGA PEÑAFIEL**

---

## 💡 Conclusión

Este proyecto permitió aplicar los conocimientos adquiridos durante la asignatura de Programación Orientada a Objetos I mediante el desarrollo de una API REST para una biblioteca digital. Se implementaron estructuras de datos, organización modular del código, servicios web y operaciones CRUD, fortaleciendo las habilidades en el desarrollo de aplicaciones backend con Go.
