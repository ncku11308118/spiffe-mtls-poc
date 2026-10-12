# SPIFFE mTLS PoC

## Architecture

- CA
  - HashiCorp Vault (Optional)
  - Open Policy Agent
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
config:
  layout: elk
---

graph BT
  subgraph ca[CA]
    ca.hashicorp-vault[HashiCorp Vault]
    ca.open-policy-agent[Open Policy Agent]
    ca.spire-server[SPIRE Server]
  end

  subgraph server[Server Node]
    server.spire-agent[SPIRE Agent]

    subgraph server.network[Network]
      server.traefik-proxy[Traefik Proxy]
      server.workload[Workload]
    end
  end

  subgraph client[Client Node]
    client.spire-agent[SPIRE Agent]

    subgraph client.volume[Volume]
      client.spiffe-helper[SPIFFE Helper]
      client.workload[Workload]
    end
  end

  server.spire-agent -- (1) --> ca.spire-server
  client.spire-agent -- (1) --> ca.spire-server
  ca.spire-server -- (2) --> ca.open-policy-agent
  ca.spire-server -- (3) --> ca.hashicorp-vault
  server.traefik-proxy -- (4) --> server.spire-agent
  client.spiffe-helper -- (4) --> client.spire-agent
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
