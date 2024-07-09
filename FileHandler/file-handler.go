package main

import (
    "fmt"
    "os"
)

func main() {
    writeFile()
    readFile()
}

func writeFile() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString("Hello, world!\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File successfully created and written to.")
}

func readFile() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Read %d bytes:\n%s\n", count, data[:count])
}
