package miner

import "math/big"

type miner struct {
	address []byte
	transactions []message
}



type Staker interface{
	SetTreshHold(int)
	Stake(block)


}