# Imagen base de Go oficial
FROM golang:1.25

# Instala utilidades básicas
RUN apt-get update && apt-get install -y \
    git \
    vim \
    curl \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Crea directorio de trabajo
WORKDIR /workspace

# Configura Go path (opcional)
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# Muestra versión de Go al iniciar
CMD ["go", "version"]
