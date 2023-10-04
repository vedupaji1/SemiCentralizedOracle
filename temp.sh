#!/bin/bash

# Define the RPC endpoint URL
RPC_ENDPOINT="0.0.0.0:8000/nodeData"

# Construct the JSON-RPC request
JSON_REQUEST='{
    "method": "NodeRequests.NodeData", 
    "params": [{}],
    "id": 1
}'

# Send the JSON-RPC request using curl
curl -X POST -H 'Content-Type: application/json' -d "$JSON_REQUEST" "$RPC_ENDPOINT"


