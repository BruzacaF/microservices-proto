Param()

# PowerShell script to generate Go protobuf/gRPC stubs
# Requires: protoc, protoc-gen-go, protoc-gen-go-grpc in PATH

$protoRoot = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
$outDir = Join-Path $protoRoot 'golang'

if (!(Test-Path $outDir)) { New-Item -ItemType Directory -Path $outDir | Out-Null }

& protoc --proto_path=$protoRoot `
	--go_out=$outDir --go_opt=paths=source_relative `
	--go-grpc_out=$outDir --go-grpc_opt=paths=source_relative `
	(Join-Path $protoRoot 'order\order.proto')

Write-Output "Generated Go stubs in $outDir"

# PowerShell script to generate Go protobuf stubs for Windows users
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$OutDir = Join-Path $ScriptDir 'golang'
if (-not (Test-Path $OutDir)) { New-Item -ItemType Directory -Path $OutDir | Out-Null }

protoc --proto_path=$ScriptDir --go_out=$OutDir --go-grpc_out=$OutDir payment\payment.proto
Write-Host "Generated Go protobuf files into $OutDir"
