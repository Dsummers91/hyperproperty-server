package monopoly

import (
	"build-chaincode/util"
	"errors"
	"math/rand"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func RollDice(stub shim.ChaincodeStubInterface) error {
	user, err := util.GetCurrentBlockchainUser(stub)
	if err != nil {
		return errors.New(err.Error())
	}
	board, _ := GetCurrentState(stub)

	numberRolled := rand.Intn(10) + 2

	if user.UserID == "player1" {
		board.Player1.CurrentTurn = true
		if !board.Player1.CurrentTurn {
			return errors.New("Player1 is not active")
		}
		newPosition := board.Player1.CurrentPosition.ID + numberRolled
		if newPosition > 40 {
			newPosition -= 40
			board.Player1.Balance += 200
		}
		board.Player1.CurrentPosition = board.Positions[newPosition]
		board.Player1.HasRolled = true
	} else if user.UserID == "player2" {
		newPosition := board.Player1.CurrentPosition.ID + numberRolled
		if newPosition > 40 {
			newPosition -= 40
			board.Player1.Balance += 200
		}
		if !board.Player2.CurrentTurn || board.Player2.HasRolled {
			return errors.New("Player2 is not active")
		}
		board.Player2.CurrentPosition = board.Positions[newPosition]
		board.Player2.HasRolled = true
	}

	boardAsBytes, _ := json.Marshal(board)
	stub.PutState("board", boardAsBytes)
	return nil
}

func PlayerAction(stub shim.ChaincodeStubInterface, action string) error {
	user, err := util.GetCurrentBlockchainUser(stub)
	if err != nil {
		return errors.New(err.Error())
	}
	board, _ := GetCurrentState(stub)
	var player *Player
	var opponent string
	if user.UserID == "player1" {
		player = &board.Player1
		opponent = "player2"
	} else {
		player = &board.Player2
		opponent = "player1"
	}
	if !player.CurrentTurn {
		return errors.New("Is not players current turn to buy")
	}

	if action == "buy" && board.Positions[player.CurrentPosition.ID-1].BelongsTo != opponent {
		player.Balance -= player.CurrentPosition.Cost
		if player.Balance >= 0 {
			player.PositionsOwned = append(player.PositionsOwned, player.CurrentPosition)
			board.Positions[player.CurrentPosition.ID-1].BelongsTo = user.UserID
		}
	}
	board.Player1.CurrentTurn = !board.Player1.CurrentTurn
	board.Player2.CurrentTurn = !board.Player2.CurrentTurn
	board.Player1.HasRolled = false
	board.Player2.HasRolled = false
	boardAsBytes, _ := json.Marshal(board)
	positionsAsBytes, _ := json.Marshal(board.Positions)
	stub.PutState("board", boardAsBytes)
	stub.PutState("positions", positionsAsBytes)
	return nil
}
