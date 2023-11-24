# Scout

## Purpose

Scout continuously probes pods on separate kubernetes nodes to provide insight into network latency, restrictions, and reliability.

## Components

Scout can be deployed via a daemonset or deployment. Each pod should listen for requests from remote containers, while continously probing the same remote containers. Scout must determine the IPs/pod names of its associated containers. Metrics are available at localhost:8080/metrics.

Scout containers need access to kube apiserver to retrieve details about other scout pods.

## Example output

```
2023/11/24 21:27:11 scout pods:
2023/11/24 21:27:11 	node: kind-worker
2023/11/24 21:27:11 		scout-c64f88c47-b8ll2.default.10.244.1.34
2023/11/24 21:27:11 	node: kind-worker
2023/11/24 21:27:11 		scout-c64f88c47-cndfr.default.10.244.1.35
2023/11/24 21:27:11 	node: kind-worker
2023/11/24 21:27:11 		scout-c64f88c47-z6xcx.default.10.244.1.33
2023/11/24 21:27:11 probing: scout-c64f88c47-b8ll2 -> scout-c64f88c47-b8ll2.default.10.244.1.34 @ node: kind-worker
2023/11/24 21:27:11 	connection start: 2023-11-24 21:27:11.884794208 +0000 UTC m=+15.035662049
2023/11/24 21:27:11 	latency connection: 235µs
2023/11/24 21:27:11 	latency write request: 82.083µs
2023/11/24 21:27:11 	latency server processing: 413.708µs
2023/11/24 21:27:11 	latency content transfer: 170.584µs
2023/11/24 21:27:11 	latency total: 1.388959ms
2023/11/24 21:27:11 	response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout"}}
2023/11/24 21:27:11 probing: scout-c64f88c47-b8ll2 -> scout-c64f88c47-cndfr.default.10.244.1.35 @ node: kind-worker
2023/11/24 21:27:11 	connection start: 2023-11-24 21:27:11.886367542 +0000 UTC m=+15.037235341
2023/11/24 21:27:11 	latency connection: 275.167µs
2023/11/24 21:27:11 	latency write request: 46.041µs
2023/11/24 21:27:11 	latency server processing: 315.25µs
2023/11/24 21:27:11 	latency content transfer: 68.541µs
2023/11/24 21:27:11 	latency total: 809.458µs
2023/11/24 21:27:11 	response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout"}}
2023/11/24 21:27:11 probing: scout-c64f88c47-b8ll2 -> scout-c64f88c47-z6xcx.default.10.244.1.33 @ node: kind-worker
2023/11/24 21:27:11 	connection start: 2023-11-24 21:27:11.887332708 +0000 UTC m=+15.038200466
2023/11/24 21:27:11 	latency connection: 284.709µs
2023/11/24 21:27:11 	latency write request: 36.042µs
2023/11/24 21:27:11 	latency server processing: 501.042µs
2023/11/24 21:27:11 	latency content transfer: 109.958µs
2023/11/24 21:27:11 	latency total: 1.059958ms
2023/11/24 21:27:11 	response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout"}}
```

## Promql metrics

```
# HELP scout_total_conn_duration gauge for total connection creation duration
# TYPE scout_total_conn_duration gauge
scout_total_conn_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-b8ll2",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 1288
scout_total_conn_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-cndfr",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 1540
scout_total_conn_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-z6xcx",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 1682
# HELP scout_total_dns_duration gauge for total dns duration
# TYPE scout_total_dns_duration gauge
scout_total_dns_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-b8ll2",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 0
scout_total_dns_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-cndfr",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 0
scout_total_dns_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-z6xcx",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 0
# HELP scout_total_latency gauge for total latency ms
# TYPE scout_total_latency gauge
scout_total_latency{dest_node="kind-worker",dest_pod="scout-c64f88c47-b8ll2",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 5
scout_total_latency{dest_node="kind-worker",dest_pod="scout-c64f88c47-cndfr",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 3
scout_total_latency{dest_node="kind-worker",dest_pod="scout-c64f88c47-z6xcx",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 3
# HELP scout_total_requests counts the number of requests
# TYPE scout_total_requests counter
scout_total_requests{dest_node="kind-worker",dest_pod="scout-c64f88c47-b8ll2",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 5
scout_total_requests{dest_node="kind-worker",dest_pod="scout-c64f88c47-cndfr",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 6
scout_total_requests{dest_node="kind-worker",dest_pod="scout-c64f88c47-z6xcx",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 5
# HELP scout_total_server_processing_duration gauge for total server processing duration
# TYPE scout_total_server_processing_duration gauge
scout_total_server_processing_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-b8ll2",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 2782
scout_total_server_processing_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-cndfr",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 2862
scout_total_server_processing_duration{dest_node="kind-worker",dest_pod="scout-c64f88c47-z6xcx",src_node="kind-worker",src_pod="scout-c64f88c47-b8ll2"} 2562
```

## Future additions

There is a basic client/server UDP setup under the /udp dir. Potentially include this alongside tcp probes
