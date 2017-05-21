package monopoly

import (
	"encoding/json"

	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func InsertPositions() {

}

func ResetPositions(stub shim.ChaincodeStubInterface) error {

	var positions []Position

	positions = append(positions, Position{
		ID:   1,
		Type: "go",
	})
	positions = append(positions, Position{
		Cost: 60,
		ID:   2,
		Type: "chest",
	})
	positions = append(positions, Position{
		Cost: 200,
		ID:   3,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 100,
		ID:   4,
		Type: "incometax",
	})
	positions = append(positions, Position{
		Cost: 200,
		ID:   5,
		Type: "railroad",
	})
	positions = append(positions, Position{
		Cost: 100,
		ID:   6,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 100,
		ID:   7,
		Type: "chance",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   8,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 150,
		ID:   9,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 0,
		ID:   10,
		Type: "visiting",
	})
	positions = append(positions, Position{
		Cost: 150,
		ID:   11,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 160,
		ID:   12,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   13,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 150,
		ID:   14,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 200,
		ID:   15,
		Type: "railroad",
	})
	positions = append(positions, Position{
		Cost: 170,
		ID:   16,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 180,
		ID:   17,
		Type: "chest",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   18,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   19,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 100,
		ID:   20,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   21,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 0,
		ID:   22,
		Type: "chance",
	})
	positions = append(positions, Position{
		Cost: 90,
		ID:   23,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   24,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 150,
		ID:   25,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   26,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 100,
		ID:   27,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 150,
		ID:   28,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 110,
		ID:   29,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 180,
		ID:   30,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   31,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 110,
		ID:   32,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 140,
		ID:   33,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 90,
		ID:   34,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 110,
		ID:   35,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 0,
		ID:   36,
		Type: "chance",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   37,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 0,
		ID:   38,
		Type: "luxurytax",
	})
	positions = append(positions, Position{
		Cost: 90,
		ID:   39,
		Type: "property",
	})
	positions = append(positions, Position{
		Cost: 120,
		ID:   40,
		Type: "property",
	})

	positionsAsBytes, err := json.Marshal(positions)
	if err != nil {
		return errors.New("bleh")
	}
	return stub.PutState("positions", positionsAsBytes)

}
