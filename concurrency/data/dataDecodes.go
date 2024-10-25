package main

import (
    "encoding/json"
    "fmt"
    "os"
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
    file, err := os.Open("data.json")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    dec := json.NewDecoder(file)

    var nodes Nodes
    if err := dec.Decode(&nodes); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    typeANodes := make([]Node, 0)

    for _, node := range nodes.Nodes {
        if node.Type == "A" {
            typeANodes = append(typeANodes, node)
        }
    }

    fmt.Println("Nodes of type A:")
    for _, node := range typeANodes {
        fmt.Printf("IP: %s, Network: %s\n", node.IP, node.Network)
    }
}
