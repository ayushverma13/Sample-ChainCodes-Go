// When yosu build reusable pieces of code, you will develop a package as a shared library. But when you develop executable programs, you will use the package “main” for making the package as an executable program. The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. 
package main

import (
  "fmt"
  "log"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
   contractapi.Contract
   }

// Asset describes basic details of what makes up a simple asset
type Asset struct {
    ID             string `json:"ID"`
    Color          string `json:"color"`
    Size           int    `json:"size"`
    Owner          string `json:"owner"`
    AppraisedValue int    `json:"appraisedValue"`
   }

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
	  {ID: "asset1", Color: "blue", Size: 5, Owner: "Tomoko", AppraisedValue: 300},
	  {ID: "asset2", Color: "red", Size: 5, Owner: "Brad", AppraisedValue: 400},
	  {ID: "asset3", Color: "green", Size: 10, Owner: "Jin Soo", AppraisedValue: 500},
	  {ID: "asset4", Color: "yellow", Size: 10, Owner: "Max", AppraisedValue: 600},
	  {ID: "asset5", Color: "black", Size: 15, Owner: "Adriana", AppraisedValue: 700},
	  {ID: "asset6", Color: "white", Size: 15, Owner: "Michel", AppraisedValue: 800},
	}

 for _, asset := range assets {
	assetJSON, err := json.Marshal(asset)
	if err != nil {
	  return err
	}

	err = ctx.GetStub().PutState(asset.ID, assetJSON)
	if err != nil {
	  return fmt.Errorf("failed to put to world state. %v", err)
	}
  }

  return nil
}