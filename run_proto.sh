#!/usr/bin/env bash
set -euo pipefail

# Example: generate Go gRPC stubs for all proto files into a golang directory
# Requires: protoc, protoc-gen-go, protoc-gen-go-grpc in PATH

PROTO_ROOT="$(cd "$(dirname "$0")" && pwd)"
OUT_DIR="$PROTO_ROOT/golang"

mkdir -p "$OUT_DIR"

protoc \
  --proto_path="$PROTO_ROOT" \
  --go_out="$OUT_DIR" --go_opt=paths=source_relative \
  --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative \
  "$PROTO_ROOT"/order/order.proto

echo "Generated Go stubs in $OUT_DIR"
#!/usr/bin/env bash
set -euo pipefail

# Generates Go code from protobufs in this repo.
PROTO_DIR="$(cd "$(dirname "$0")" && pwd)"
OUT_DIR="$PROTO_DIR/golang"

mkdir -p "$OUT_DIR"

protoc \
  --proto_path="$PROTO_DIR" \
  --go_out="$OUT_DIR" \
  --go-grpc_out="$OUT_DIR" \
  payment/payment.proto

echo "Generated Go protobuf files into $OUT_DIR"
