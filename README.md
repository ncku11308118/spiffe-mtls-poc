# SPIFFE mTLS PoC

## Architecture

- CA
  - HashiCorp Vault (Optional)
  - SPIRE Server
- Server Node
  - SPIRE Agent
  - Traefik Proxy
  - Workload
- Client Node
  - SPIRE Agent
  - SPIFFE Helper
  - Workload

```mermaid
---
title: Trust Domain
---

flowchart BT
  subgraph ca[CA]
    ca.spire-server[SPIRE Server]
  end

  subgraph server[Server Node]
  	server.spire-agent[SPIRE Agent]
	  server.traefik-proxy[Traefik Proxy]
  	server.workload[Workload]
  end

  subgraph client[Client Node]
  	client.spire-agent[SPIRE Agent]
  	client.spiffe-helper[SPIFFE Helper]
  	client.workload[Workload]
  end

  server.traefik-proxy --> server.spire-agent --> ca.spire-server
  client.spiffe-helper --> client.spire-agent --> ca.spire-server
  client.workload == mTLS ==> server.traefik-proxy ==> server.workload
```

## Setup

To run CA:

```sh
docker compose up --detach
```

To run server:

```sh
docker compose --profile server-node up --detach
```

To run client:

```sh
docker compose --profile client-node up --detach
```
