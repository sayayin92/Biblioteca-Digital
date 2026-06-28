📚 Biblioteca Digital en Go (2026)

Sistema de gestión de biblioteca desarrollado en Go (Golang) utilizando el framework Gin, enfocado en la aplicación de conceptos de Programación Orientada a Objetos (POO) y arquitectura en capas.

📌 Datos del grupo
Integrante: YUSI ARISTEGA PEÑAFIEL
Curso: 3A
Materia: Programación Orientada a Objetos 1
Fecha de entrega: 28 de junio de 2026
🚀 Descripción del proyecto

Este proyecto consiste en una API REST para la gestión de una biblioteca digital, que permite administrar:

Usuarios 👤
Libros 📚
Préstamos 📖

El sistema está desarrollado utilizando structs en Go, separación por capas y lógica modular, simulando principios de Programación Orientada a Objetos.

🧱 Arquitectura del proyecto
Biblioteca-Digital/
│
├── main.go
│
├── handlers/
│   ├── libros.go
│   ├── usuarios.go
│   └── prestamos.go
│
├── models/
│   ├── libro.go
│   ├── usuario.go
│   └── prestamo.go
│
├── database/
│   └── conexion.go
│
├── routes/
│   └── routes.go
│
└── go.mod
⚙️ Tecnologías utilizadas
Go (Golang)
Gin Web Framework
Arquitectura en capas
API REST con JSON
Structs (POO en Go)
Slices (almacenamiento en memoria)
▶️ Instalación y ejecución
1. Clonar el repositorio
git clone https://github.com/sayayin92/Biblioteca-Digital.git
2. Entrar al proyecto
cd Biblioteca-Digital
3. Instalar dependencias
go mod tidy
4. Ejecutar el servidor
go run main.go
🌐 Servidor

El proyecto se ejecuta en:

http://localhost:8080
📌 Endpoints de la API
👤 Usuarios
➕ Crear usuario
POST /usuarios
{
  "nombre": "Juan",
  "email": "juan@gmail.com"
}
📄 Listar usuarios
GET /usuarios
📚 Libros
➕ Crear libro
POST /libros
{
  "titulo": "El Quijote",
  "autor": "Cervantes"
}
📄 Listar libros
GET /libros
🔎 Buscar libro por ID
GET /libros/:id
📖 Préstamos
➕ Crear préstamo
POST /prestamos
{
  "usuario_id": 1,
  "libro_id": 1
}
📄 Listar préstamos
GET /prestamos
🧠 Funcionamiento interno

El sistema utiliza:

structs para modelar entidades
handlers para la lógica de negocio
routes para definir endpoints
database en memoria con slices
separación por módulos para escalabilidad
⚠️ Limitaciones
No utiliza base de datos persistente
Los datos se pierden al reiniciar el servidor
No incluye autenticación
Validaciones básicas
🚀 Mejoras futuras
Integración con PostgreSQL o MySQL
Autenticación JWT 🔐
Validaciones avanzadas
Documentación Swagger
Deploy en la nube (Render / Railway)
👨‍💻 Autor
Integrante: YUSI ARISTEGA PEÑAFIEL
Curso: 3A
Materia: Programación Orientada a Objetos 1
Fecha: 27 de junio de 2026
💡 Conclusión

Este proyecto representa una base sólida para el aprendizaje de backend en Go, aplicando arquitectura en capas, POO y desarrollo de APIs REST.
