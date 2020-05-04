package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CCHello struct{}

func main() {
	err := shim.Start(new(CCHello))
	if err != nil {
		fmt.Printf("start hello chaincode fail: %v", err)
	}
}

func (*CCHello) Name() string {
	return "hello"
}

func (cc *CCHello) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Printf("start to instantiate chaincode %q\n", cc.Name())

	_, args := stub.GetFunctionAndParameters()
	if len(args) != 2 {
		return shim.Error(fmt.Sprintf("wrong number of parameters, expected: 2, got: %d", len(args)))
	}

	fmt.Println("start to save data")

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("save data failed: %v", err))
	}

	fmt.Printf("instantiate chaincode %q succeed\n", cc.Name())
	return shim.Success(nil)
}

func (cc *CCHello) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()

	switch fun {
	case "query":
		return query(stub, args)
	default:
		fmt.Printf("method %q is not supported yet\n", fun)
		return shim.Error("unsupported method")
	}
}

func query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("wrong number of parameters for method %q, expected: 1, got: %d", "query", len(args)))
	}

	result, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("get data failed: %v", err))
	}
	if result == nil {
		return shim.Error(fmt.Sprintf("query get no result for key %q", args[0]))
	}
	return shim.Success(result)
}
