package monopoly

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetCurrentState(stub shim.ChaincodeStubInterface) (Board, error) {
	var chainTx ChainTX
	var board Board
	var positions []Position
	boardAsBytes, err := stub.GetState("board")
	if err != nil {
		return Board{}, errors.New("Could not find board")
	}
	err = json.Unmarshal(boardAsBytes, &board)
	if err != nil {
		return Board{}, errors.New("Error Unmarshalling json")
	}
	positionsAsBytes, err := stub.GetState("positions")
	if err != nil {
		return Board{}, errors.New("Could not find position")
	}
	err = json.Unmarshal(positionsAsBytes, &positions)
	if err != nil {
		return Board{}, errors.New("Error Unmarshalling json")
	}
	url := "https://207418d37fd4468698d9e7dd36373ba8-vp0.us.blockchain.ibm.com:5002/chain"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	chainAsString := string(responseData)

	err = json.Unmarshal([]byte(chainAsString), &chainTx)
	if err != nil {
		return Board{}, errors.New(string(responseData) + err.Error())
	}
	board.ChainHeight = chainTx.Height
	board.TransactionHash = chainTx.CurrentBlockHash
	board.Positions = positions
	return board, nil
}

func InitializeGame(stub shim.ChaincodeStubInterface) error {
	var board Board
	var player Player
	player.Balance = 15000
	board.Player1 = player
	board.Player1.CurrentTurn = true
	board.Player2 = player
	ResetPositions(stub)
	boardAsBytes, _ := json.Marshal(board)
	stub.PutState("board", boardAsBytes)
	return nil
}
