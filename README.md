# 📊 RESTful API with Go and Gin to Any Benchmarking

<table>
    <tr>
        <td>
            <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
        </td>
        <td>
            <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" />
        </td>
        <td>
            <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" />
        </td>
        <td>
            <img src="https://img.shields.io/badge/Shell_Script-121011?style=for-the-badge&logo=gnu-bash&logoColor=white" />
        </td>
        </td>
    </tr>
</table>

## 🚀 Introduction
This project aims to provide a RESTful API implemented in [Go](https://go.dev/) using the Gin framework for benchmarking various sorting algorithms. The API follows the principles of Clean Architecture, emphasizing modularity, testability, and maintainability.

<img width="1000" src="https://technology.riotgames.com/sites/default/files/articles/116/golangheader.png" alt="golang logo">


## 📋 Requirements
> [!NOTE]
> To start this project, you will need Docker and Docker Compose installed on your machine.

## ⏳ Next Steps
The project is still under development and the next updates will focus on the following tasks:

- [x] Implementation of various sorting algorithms (Bubble Sort, Insertion Sort, Merge Sort, Quick Sort, Selection Sort).
- [ ] Serialization support for multiple formats (JSON, Protocol Buffers, FlatBuffers).

## 📌 Features
- Dockerized for easy deployment and scalability.
- Unit tests for all components, including mocks for interfaces.
- Dependency injection for managing dependencies and facilitating testing.

## 🛠️ Project Structure
The project follows a modular structure inspired by Clean Architecture:
- **cmd/server/main.go**: Entry point for the application.
- **di/**: Dependency injection configurations for various components.
- **internal/**: Core application logic.
  - **application/**: Use cases and DTOs.
  - **domain/entity/**: entities.
  - **infrastructure/api/**: API controllers and routes.
  - **infrastructure/repository/**: Repository implementations.
- **pkg/**: Reusable packages and utilities.
- **tests/**: Test helpers and mocks.

## ✅ Installation
To run the project, ensure you have Go and Docker installed on your machine. Then, follow these steps:
1. Clone the repository.
2. Navigate to the project directory.
3. Run `docker-compose up --build` to start the application.
```bash
docker-compose up --build
```

## 🌟 Usage
Once the application is running, you can interact with the API using tools like Postman or cURL. Here are some sample endpoints:
```
http://localhost:3000
```
- `/health`: Check the health of the API.
- `/benchmark/sort`: Benchmark different sorting algorithms.

## ✌️ Acknowledgements

Hi, I'm Brazilian, my name is Diego Marques Vieira and my English isn't the best haha. This project is just to implement clean code and clean architecture studios in Golang, thanks for reviewing and if you can leave a star 🌟!

<img width="1000" src="https://i.pinimg.com/originals/fb/af/44/fbaf443c014bf40b95cfa35121572b25.gif" alt="study gif">

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.
