FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    vim \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -ms /bin/bash devuser
USER devuser

WORKDIR /workspace
