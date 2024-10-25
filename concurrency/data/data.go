package main

import (
    "encoding/json"
    "fmt"
    "sync"
)

type Node struct {
    IP      string `json:"ip"`
    Network string `json:"network"`
    Type    string `json:"type"`
}

type Nodes struct {
    Nodes map[string]Node `json:"nodes"`
}

func main() {
    data := []byte(`{"nodes": {"node1": {"ip":"123", "network":"123", "type":"A"}, "node2": {"ip":"456", "network":"56", "type":"B"}}}`)

    var nodes Nodes
    json.Unmarshal(data, &nodes)

    var wg sync.WaitGroup
    typeACh := make(chan Node)

    for _, node := range nodes.Nodes {
        wg.Add(1)
        go func(node Node) {
            defer wg.Done()
            if node.Type == "A" {
                typeACh <- node
            }
        }(node)
    }

    go func() {
        wg.Wait()
        close(typeACh)
    }()

    var typeANodes []Node
    for node := range typeACh {
        typeANodes = append(typeANodes, node)
    }

    fmt.Println("Nodes of type A:")
    for _, node := range typeANodes {
        fmt.Printf("IP: %s, Network: %s\n", node.IP, node.Network)
    }
}
