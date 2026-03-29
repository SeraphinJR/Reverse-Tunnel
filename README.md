# 🌌 Rift

> A lightweight, multiplexed reverse tunnel written in pure Go. Expose local servers to the public internet behind strict NATs and firewalls.

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green.svg)

**Rift** is a custom-built, zero-dependency reverse proxy network that acts as a private alternative to tools like `ngrok`. It establishes a persistent, bidirectional TCP tunnel from your local machine to a public cloud relay, allowing you to securely expose local environments to the world.

Whether you need to rapidly expose a Spring Boot backend to judges during a live hackathon demo, test webhooks, or share a local DevOps dashboard with your team, Rift handles the routing seamlessly.

## ✨ Core Features

* **True Multiplexing:** Utilizes HashiCorp's `yamux` to multiplex hundreds of simultaneous HTTP requests over a single persistent TCP connection. 
* **Bulletproof Resiliency:** The Local Agent features an automatic exponential backoff algorithm. If you drop off Wi-Fi or change networks, the tunnel silently re-establishes itself in the background.
* **Bypass Strict Firewalls:** Initiates outbound connections from the local machine, natively bypassing aggressive home ISP and corporate NAT firewalls.
* **Micro-Footprint:** Compiles down to a single, statically linked binary. The cloud relay can be deployed in a bare `scratch` Docker container weighing less than 15MB.

## 🏗️ Architecture

Rift is composed of two distinct binaries that work in tandem:

1. **The Public Relay (`/relay`)**: A lightweight server deployed on a public cloud VPS. It listens for incoming web traffic (Port 80/8000) and routes it down the active multiplexed tunnel (Port 7000).
2. **The Local Agent (`/agent`)**: A background process running on your local machine. It maintains the persistent TCP connection to the Relay and forwards incoming tunnel traffic to your local development server (e.g., `localhost:8080`).

## 🚀 Getting Started

### 1. Deploy the Cloud Relay
The relay is designed to be deployed on any public Linux VPS (GCP, AWS, DigitalOcean). 

```bash
# Clone the repository
git clone [https://github.com/SeraphnJR/rift.git](https://github.com/SeraphinJR/rift.git)
cd rift/relay

# Build and run the Docker container
docker build -t rift-relay .
docker run -d --name rift-relay --restart unless-stopped -p 7000:7000 -p 80:8000 rift-relay
