package database

import (
	"encoding/json"
	"io/ioutil"
)

var genesisJson = `
{
    "genesis_time": "2022-08-01T00:18:46.000000000Z",
    "chain_id": "the-blockchain-bar-ledger",
    "balances": {
        "mustafa": 1000000
    }
}`

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return genesis{}, err
	}

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return genesis{}, err
	}

	return loadedGenesis, nil
}

func writeGenesisToDisk(path string) error {
	return ioutil.WriteFile(path, []byte(genesisJson), 0644)
}
