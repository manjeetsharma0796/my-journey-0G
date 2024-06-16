package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/joho/godotenv"

	"github.com/0glabs/0g-storage-client/common/blockchain"
	"github.com/0glabs/0g-storage-client/contract"
	"github.com/0glabs/0g-storage-client/core"
	"github.com/0glabs/0g-storage-client/node"
	"github.com/0glabs/0g-storage-client/transfer"
	// "github.com/joho/godotenv"
)

type uploadArgs struct {
	file      string
	tags      string
	url       string
	contract  string
	key       string
	force     bool
	taskSize  uint
}

func uploadFile(client *node.Client) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	args := uploadArgs{
		file:     "test.txt",
		tags:     "0x",
		url:      "https://rpc-testnet.0g.ai",
		contract: "0xb8F03061969da6Ad38f0a4a9f8a86bE71dA3c8E7",
		key:      os.Getenv("PRIVATE_KEY"),
		force:    false,
		taskSize: 10,
	}
	w3client := blockchain.MustNewWeb3(args.url, args.key)
	defer w3client.Close()

	contractAddr := common.HexToAddress(args.contract)

	flowContract, err := contract.NewFlowContract(contractAddr, w3client)
	if err != nil {
		fmt.Println("Error creating contract:", err)
		return
	}

	uploader, err := transfer.NewUploader(flowContract, []*node.Client{client})
	if err != nil {
		fmt.Println("Error creating uploader:", err)
		return
	}

	opt := transfer.UploadOption{
		Tags:    hexutil.MustDecode(args.tags),
		Force:   args.force,
		TaskSize: args.taskSize,
	}

	file, err := core.Open(args.file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if err := uploader.Upload(file, opt); err != nil {
		fmt.Println("Error uploading file:", err)
	}
}

func main() {
	ip := "https://rpc-storage-testnet.0g.ai"

	client, err := node.NewClient(ip)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	status, err := client.ZeroGStorage().GetStatus()
	if err != nil {
		fmt.Println("Error getting status:", err)
		return
	}

	fmt.Println(status)
	appendToFile("status.txt", status)

	uploadFile(client)
}



// Function to append data to a file
func appendToFile(filename string, data string) {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append the data
	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Data successfully appended to file")
	}
}