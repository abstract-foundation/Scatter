package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Block represents a block in your network.
type Block struct {
	Number     int32
	Timestamp  int64
	Data       string
	PrevHash   string
	MerkleRoot string
}

func writeBlockToFile(block *Block, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the fields of the block to the file
	if err := binary.Write(file, binary.LittleEndian, block.Number); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, block.Timestamp); err != nil {
		return err
	}

	// Write the length of the string, followed by the string data
	dataBytes := []byte(block.Data)
	if err := binary.Write(file, binary.LittleEndian, int32(len(dataBytes))); err != nil {
		return err
	}
	if _, err := file.Write(dataBytes); err != nil {
		return err
	}

	// Repeat the process for PrevHash and MerkleRoot
	// ...

	return nil
}

func readBlockFromFile(filename string) (*Block, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var block Block

	// Read the fields of the block from the file
	if err := binary.Read(file, binary.LittleEndian, &block.Number); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &block.Timestamp); err != nil {
		return nil, err
	}

	// Read the length of the string, then read the string data
	var dataLength int32
	if err := binary.Read(file, binary.LittleEndian, &dataLength); err != nil {
		return nil, err
	}
	dataBytes := make([]byte, dataLength)
	if _, err := file.Read(dataBytes); err != nil {
		return nil, err
	}
	block.Data = string(dataBytes)

	// Repeat the process for PrevHash and MerkleRoot
	// ...

	return &block, nil
}

func main() {
	// Create a sample block
	block := &Block{
		Number:     1,
		Timestamp:  1632778800,
		Data:       "Hello, World!",
		PrevHash:   "00000000000000000000000000000000",
		MerkleRoot: "1234567890abcdef",
	}

	// Save the block to a binary file
	err := writeBlockToFile(block, "block1.dat")
	if err != nil {
		fmt.Println("Error saving block:", err)
		return
	}

	// Read the block from the binary file
	readBlock, err := readBlockFromFile("block1.dat")
	if err != nil {
		fmt.Println("Error reading block:", err)
		return
	}

	// Print the read block
	fmt.Printf("Block Number: %d\n", readBlock.Number)
	fmt.Printf("Timestamp: %d\n", readBlock.Timestamp)
	fmt.Printf("Data: %s\n", readBlock.Data)
	fmt.Printf("PrevHash: %s\n", readBlock.PrevHash)
	fmt.Printf("MerkleRoot: %s\n", readBlock.MerkleRoot)
}
