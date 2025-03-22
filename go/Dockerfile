FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y \
    wget \
    tar \
    build-essential \
    git \
    curl \
    make \
    gcc \
    vim \
    tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.24.0
ARG GO_ARCH=arm64
RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && \
    rm go${GO_VERSION}.linux-${GO_ARCH}.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /app

CMD ["bash"]
