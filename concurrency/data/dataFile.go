package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
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
    file, err := os.Open("data.json")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)

    var nodes Nodes
    json.Unmarshal(byteValue, &nodes)

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

