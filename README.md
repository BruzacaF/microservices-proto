# gRPC Protocol Buffer Definitions

Este repositório contém as definições de Protocol Buffer para os microsserviços gRPC.

## Estrutura

```
microservices-proto/
├── order/
│   └── order.proto       # Definição do serviço Order
└── README.md
```

## Usando os Proto Files

Para compilar os arquivos `.proto` para Go:

```bash
protoc --go_out=golang --go-grpc_out=golang order/order.proto
```

### Geração automatizada (recomendada)

Incluí dois scripts para facilitar a geração dos stubs Go:

- `run_proto.sh` — para ambientes Unix/WSL
- `run_proto.ps1` — para PowerShell no Windows

Exemplo (PowerShell):

```powershell
cd microservices-proto
.\run_proto.ps1
```

Se `protoc` ou os plugins (`protoc-gen-go`, `protoc-gen-go-grpc`) não estiverem no `PATH`, instale-os antes de rodar os scripts.

## Serviços Disponíveis

### Order Service

**Método:** `Create`
- Cria um novo pedido com itens

**Request:**
```protobuf
message CreateOrderRequest {
  int32 costumer_id = 1;
  repeated OrderItem order_items = 2;
  float total_price = 3;
}
```

**Response:**
```protobuf
message CreateOrderResponse {
  int32 order_id = 1;
}
```

## Mensagens

### OrderItem
```protobuf
message OrderItem {
  string product_code = 1;
  float unit_price = 2;
  int32 quantity = 3;
}
```

## Referências

- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers)
- [gRPC Go Documentation](https://grpc.io/docs/languages/go/)
