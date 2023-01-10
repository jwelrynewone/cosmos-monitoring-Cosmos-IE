package cmd

import (
        "fmt"
	"log"

        "github.com/spf13/cobra"

//	cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos"
//        terra "github.com/node-a-team/Cosmos-IE/chains/terra"
//        iris "github.com/node-a-team/Cosmos-IE/chains/iris"
//        kava "github.com/node-a-team/Cosmos-IE/chains/kava"
//        iov "github.com/node-a-team/Cosmos-IE/chains/iov"
//        emoney "github.com/node-a-team/Cosmos-IE/chains/emoney"
//        band "github.com/node-a-team/Cosmos-IE/chains/bandprotocol"

	exporter "github.com/node-a-team/Cosmos-IE/exporter"

)

var (

	chainList []string

	// command로 안 받을 경우 defult 값 지정
	chain string = "okp4-nemeton-1"
        rpcAddr string = "localhost:36657"
        restAddr string = "localhost:1327"
	listenPort string = "36661"

	operAddr string= ""
	consHexAddr string = ""

	// for Terra
	feeDenom string ="ukrw"
)

// versionCmd represents the version command
var runCmd = &cobra.Command{
        Use:   "run",
        Short: "Validator Operator Address",
        Long: `Be sure to enter either Validator Operator Address or Consensus Hex Address
ex#1_Local REST Server) Cosmos-IE run --chain cosmos --oper-addr cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys
ex#2_Remote REST Server) Cosmos-IE run --chain cosmos --oper-addr cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys --rest-server 192.168.0.10:1317 
`,
        Run: func(cmd *cobra.Command, args []string) {
		check_chain()
		set_config()
		run()
        },
}

func init() {
        rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&chain, "chain", "c", "", "Chain name of the monitoring node")
        runCmd.MarkFlagRequired("chain")


	// 입력받은 값을 RpcAddr 변수에 저장, 입력 받지 않을 경우 default(tcp://localhost:26657)
        runCmd.Flags().StringVarP(&rpcAddr, "rpc-server", "", "localhost:26657", "<host>:<port> to Tendermint RPC interface for the selected chain")
	runCmd.Flags().StringVarP(&restAddr, "rest-server", "", "localhost:1317", "<host>:<port> to Rest-Server(LCD-Server) interface for the selected chain")

	runCmd.Flags().StringVarP(&listenPort, "port", "p", "26661", "Port to listen for Prometheus collector connections")

	runCmd.Flags().StringVarP(&operAddr, "oper-addr", "", "", "Operator address for Validator")
	runCmd.Flags().StringVarP(&consHexAddr, "cons-addr", "", "", "Consensus hex address for Validator")

	runCmd.Flags().StringVarP(&feeDenom, "oracle-fee-denom", "", "ukrw", "Denom type for Terra oracle fees")

}

func check_chain() {

	// chain check
	chainCheck := false
	for _, c := range chainList {
		if chain == c {
			chainCheck = true
		}
	}
	if !chainCheck {
		log.Fatal(fmt.Sprintf("[Error] %s is not supported", chain) +fmt.Sprint("\nList of supported chains: ", chainList))
	}


	if (operAddr == "" && consHexAddr == "") || (operAddr != "" && consHexAddr != "") {
		log.Fatal("[Error] Enter only one of --oper-addr or --cons-addr")
	} else if chain == "iov" && consHexAddr == "" {
		log.Fatal("[Error] In case of IOV chain, input only --cons-addr")
	} else if chain != "iov" && consHexAddr != "" {
		log.Fatal("[Error] If not an IOV chain, input only --oper-addr")
	}
}

func run() {
	switch chain {
//                case "cosmos":
//                        cosmos.Main(listenPort)
//                case "terra","emoney":
//                        exporter.Go(chain, listenPort)
//                case "iris":
//                        iris.Main(listenPort)
//                case "kava":
//                        kava.Main(listenPort)
//                case "iov":
//                        iov.Main(listenPort)
//		case "band":
//                        band.Main(listenPort)
		default:
			exporter.Go(chain, listenPort)

        }
}
