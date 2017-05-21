package main

import (
	"build-chaincode/entities"
	"build-chaincode/monopoly"
	"build-chaincode/util"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("fabric-boilerplate")

//======================================================================================================================
//	 Structure Definitions
//======================================================================================================================
//	SimpleChaincode - A blank struct for use with Shim (An IBM Blockchain included go file used for get/put state
//					  and other IBM Blockchain functions)
//==============================================================================================================================
type Chaincode struct {
}

//======================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Passes the
//  		 initial arguments passed are passed on to the called function.
//======================================================================================================================

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, functionName string, args []string) ([]byte, error) {
	logger.Infof("Invoke is running " + functionName)

	if functionName == "init" {
		return t.Init(stub, "init", args)
	} else if functionName == "resetIndexes" {
		return nil, util.ResetIndexes(stub, logger)
	} else if functionName == "addUser" {
		return nil, t.addUser(stub, args[0], args[1])
	} else if functionName == "addTestdata" {
		return nil, t.addTestdata(stub, args[0])
	} else if functionName == "rollDice" {
		return nil, monopoly.RollDice(stub)
	} else if functionName == "playeraction" {
		return nil, monopoly.PlayerAction(stub, args[0])
	} else if functionName == "playerstart" {
		return nil, monopoly.InitializeGame(stub)
	}

	return nil, errors.New("Received unknown invoke function name")
}

//======================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, functionName string, args []string) ([]byte, error) {
	logger.Infof("Query is running " + functionName)

	result, err := t.GetQueryResult(stub, functionName, args)
	if err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

func (t *Chaincode) GetQueryResult(stub shim.ChaincodeStubInterface, functionName string, args []string) (interface{}, error) {
	if functionName == "getUser" {
		user, err := util.GetUser(stub, args[0])
		if err != nil {
			return nil, err
		}

		return user, nil
	} else if functionName == "authenticateAsUser" {
		user, err := util.GetUser(stub, args[0])
		if err != nil {
			logger.Infof("User with id %v not found.", args[0])
		}

		return t.authenticateAsUser(stub, user, args[1]), nil
	} else if functionName == "getThingsByUserID" {
		thingsByUserID, err := util.GetThingsByUserID(stub, args[0])
		if err != nil {
			return nil, errors.New("could not retrieve things by user id: " + args[0] + ", reason: " + err.Error())
		}

		return thingsByUserID, nil
	} else if functionName == "getCurrentState" {
		return monopoly.GetCurrentState(stub)
	}

	return nil, errors.New("Received unknown query function name")
}

//======================================================================================================================
//  Main - main - Starts up the chaincode
//======================================================================================================================

func main() {
	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogInfo)

	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode: %s", err)
	}
}

//======================================================================================================================
//  Init Function - Called when the user deploys the chaincode
//======================================================================================================================

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

//======================================================================================================================
//  Invoke Functions
//======================================================================================================================

func (t *Chaincode) addUser(stub shim.ChaincodeStubInterface, index string, userJSONObject string) error {
	id, err := util.WriteIDToBlockchainIndex(stub, util.UsersIndexName, index)
	if err != nil {
		return errors.New("Error creating new id for user " + index)
	}

	err = stub.PutState(string(id), []byte(userJSONObject))
	if err != nil {
		return errors.New("Error putting user data on ledger")
	}

	return nil
}

func (t *Chaincode) addTestdata(stub shim.ChaincodeStubInterface, testDataAsJson string) error {

	err := monopoly.InitializeGame(stub)
	if err != nil {
		return errors.New("Error initilizing Game")
	}
	var testData entities.TestData
	err = json.Unmarshal([]byte(testDataAsJson), &testData)
	if err != nil {
		return errors.New("Error while unmarshalling testdata")
	}

	for _, user := range testData.Users {
		userAsBytes, err := json.Marshal(user)
		if err != nil {
			return errors.New("Error marshalling testUser, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, user.UserID, util.UsersIndexName, userAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	for _, thing := range testData.Things {
		thingAsBytes, err := json.Marshal(thing)
		if err != nil {
			return errors.New("Error marshalling testThing, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, thing.ThingID, util.ThingsIndexName, thingAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	return nil
}

//======================================================================================================================
//		Query Functions
//======================================================================================================================

func (t *Chaincode) authenticateAsUser(stub shim.ChaincodeStubInterface, user entities.User, passwordHash string) entities.UserAuthenticationResult {
	if user == (entities.User{}) {
		fmt.Println("User not found")
		return entities.UserAuthenticationResult{
			User:          user,
			Authenticated: false,
		}
	}

	if user.Hash != passwordHash {
		fmt.Println("Hash does not match")
		return entities.UserAuthenticationResult{
			User:          user,
			Authenticated: false,
		}
	}

	return entities.UserAuthenticationResult{
		User:          user,
		Authenticated: true,
	}
}
