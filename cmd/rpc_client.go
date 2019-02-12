package cmd

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"context"

	"google.golang.org/grpc"
	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-zenwords/rpc/proto"
)

var rpcClient = cli.Command{
	Name: "client",
	Usage: "Start RPC client",
	Description: "Run zenWords gRPC client for check service",
	Action: runRPCClient,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Default host to connect the server"),
		stringFlag("port, p", "50051", "Default port for connect to the server"),
	},
}

func runRPCClient(c *cli.Context) error {
	address         := c.String("address")
	port            := c.String("port")
	addressWithPort := fmt.Sprintf("%s:%s", address, port)

	//
	connection, err := grpc.Dial(addressWithPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to dail to %s: %v", addressWithPort, err)
	}
	defer connection.Close()

	client := proto.NewContentServiceClient(connection)

	//
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR! Unknown error when read the input value")
			continue
		}

		// Remove all \n and trim the space
		text = strings.Trim(strings.Replace(text, "\n", "", -1), " ")

		if len(text) <= 0 {
			continue
		}

		if strings.Compare(text, "exit") == 0 {
			fmt.Println("Bye!")
			break
		}

		response, err := client.Replace(
			context.Background(),
			&proto.ContentReplaceRequest{
				Content: text,
			},
		)
		if err != nil {
			log.Fatalf("Content replace response error: %v", err)
		}

		fmt.Println(response.GetResult())
	}

	return nil
}
