# Conceptos de Go 🚀

![Go](https://img.shields.io/badge/Go-1.x-00ADD8?style=for-the-badge&logo=go) ![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white) ![Dev Containers](https://img.shields.io/badge/Dev%20Containers-04AA5D?style=for-the-badge&logo=visual-studio-code&logoColor=white)

Este repositorio es una introducción a **GoLang**, diseñado para mostrar su funcionamiento a través de ejemplos prácticos y claros. 💡 Se encuentra en desarrollo activo y hace uso de **Dev Containers** y **Docker** para proporcionar un entorno de desarrollo consistente y fácil de configurar.

## ✨ Características Principales

*   **Ejemplos claros y detallados:** Demostraciones de los conceptos fundamentales de Go, desde variables hasta estructuras de control, organizados modularmente.
*   **Entorno de desarrollo listo:** Configurado con Docker y Dev Containers para una puesta en marcha rápida y sin complicaciones, asegurando portabilidad.
*   **Modular y extensible:** Cada concepto se presenta en su propio archivo o directorio, facilitando la exploración, el aprendizaje y la adición de nuevos ejemplos.
*   **Enfoque práctico:** Centrado en cómo se utilizan las características de Go en escenarios reales, ideal para principiantes y como referencia rápida.

## 📋 Requisitos Previos

Para ejecutar y explorar este proyecto, necesitarás tener instalado lo siguiente:

*   **Go:** La versión 1.x o superior de Go. Puedes descargarla desde [golang.org](https://golang.org/dl/).
*   **Docker Desktop:** Para ejecutar contenedores Docker y Docker Compose. Disponible en [docker.com](https://www.docker.com/products/docker-desktop).
*   **Docker Compose:** Generalmente se incluye con Docker Desktop.
*   **Visual Studio Code (Opcional pero recomendado):** Con la extensión "Dev Containers" para una experiencia de desarrollo integrada y fluida.

## 🚀 Instrucciones de Instalación

Sigue estos pasos para poner en marcha el proyecto en tu máquina local:

### 1. Clonar el repositorio
```bash
git clone https://github.com/tu_usuario/Primeros-Pasos-Go.git # Reemplaza con la URL real de tu repo
cd Primeros-Pasos-Go
```

### 2. Configurar el entorno con Dev Containers (Recomendado)

Si usas VS Code con la extensión Dev Containers:
*   Abre el proyecto en VS Code.
*   VS Code detectará la configuración del Dev Container y te preguntará si deseas reabrir en un contenedor. Haz clic en "Reopen in Container".
*   Esto construirá la imagen de Docker y configurará el entorno de desarrollo automáticamente. ¡Listo para empezar!

### 3. Configuración manual (Alternativa)

Si prefieres no usar Dev Containers o solo quieres ejecutar Docker Compose:
*   Navega a la raíz del proyecto donde se encuentra `docker-compose.yml`.
*   Levanta los servicios de Docker:
    ```bash
    docker-compose up -d
    ```
*   Para ingresar al contenedor y ejecutar comandos Go:
    ```bash
    docker exec -it <nombre_del_servicio_go> /bin/bash # Ajusta el nombre del servicio si es diferente
    ```

## 📚 Guía de Uso

Una vez que tengas el entorno configurado, puedes explorar y ejecutar los ejemplos de Go fácilmente:

### Ejecutar ejemplos individuales de Go

1.  **Navega al directorio `src`:**
    ```bash
    cd src
    ```

2.  **Ejecuta un archivo Go:**
    Por ejemplo, para ejecutar el ejemplo de variables:
    ```bash
    go run variables/main.go
    ```
    O para el ejemplo de arrays:
    ```bash
    go run arrays/main.go
    ```
    Repite esto para cualquier archivo `main.go` dentro de los subdirectorios de `src`.

### Usando el Dev Container

Si estás utilizando el Dev Container, ya estarás dentro de un entorno donde Go está instalado y configurado. Puedes ejecutar los comandos `go run` directamente desde la terminal integrada de VS Code sin preocuparte por la configuración local.

## 🌳 Estructura del Proyecto

La estructura del proyecto está organizada de manera lógica para facilitar el aprendizaje:
```
Primeros-Pasos-Go/
├── src/
│   ├── arrays/
│   │   └── main.go
│   ├── bytes/
│   │   └── main.go
│   ├── condicionales/
│   │   └── main.go
│   ├── constantes/
│   │   └── main.go
│   ├── conversiones/
│   │   └── main.go
│   ├── operadores/
│   │   └── main.go
│   ├── tipos-datos/
│   │   └── main.go
│   ├── variables/
│   │   └── main.go
│   └── go.mod
└── docker-compose.yml
```

*   **`src/`**: Contiene todo el código fuente de los ejemplos de Go. Cada subdirectorio representa un concepto diferente (ej. `arrays`, `variables`, `condicionales`), y cada uno contiene un archivo `main.go` con su respectiva implementación.
*   **`go.mod`**: Define el módulo Go del proyecto y gestiona las dependencias (principalmente la librería estándar de Go en este caso).
*   **`docker-compose.yml`**: Archivo de configuración para Docker Compose, utilizado para definir y ejecutar el entorno de desarrollo basado en Docker.

## 🛠️ Tecnologías Utilizadas

Este proyecto hace uso de las siguientes tecnologías:

*   **Go (Golang):** Lenguaje de programación principal del proyecto, enfocado en simplicidad, eficiencia y concurrencia.
*   **Docker:** Plataforma para desarrollar, enviar y ejecutar aplicaciones en contenedores, garantizando entornos consistentes.
*   **Docker Compose:** Herramienta para definir y ejecutar aplicaciones Docker de múltiples contenedores, facilitando la orquestación.
*   **Dev Containers:** Entornos de desarrollo basados en contenedores, integrados con VS Code, para una experiencia de desarrollo sin fricciones.