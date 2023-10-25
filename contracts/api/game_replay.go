// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GameRoundInfo is an auto generated low-level Go binding around an user-defined struct.
type GameRoundInfo struct {
	GameID    string
	RoundID   string
	ReplayID  string
	PlayerIDs string
}

// GameRoundReplay is an auto generated low-level Go binding around an user-defined struct.
type GameRoundReplay struct {
	DomainSeparationTag int64
	VRFHeight           uint64
	HashFunc            string
	VRFProof            []byte
	Address             string
	ReplayCID           string
	GameInfo            GameRoundInfo
	GameResults         []GameRoundResult
}

// GameRoundResult is an auto generated low-level Go binding around an user-defined struct.
type GameRoundResult struct {
	PlayerID     string
	CurrentScore uint64
	WinScore     uint64
}

// GameReplayContractMetaData contains all meta data concerning the GameReplayContract contract.
var GameReplayContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_replayID\",\"type\":\"string\"}],\"name\":\"getGameReplay\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"HashFunc\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"Address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayCID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PlayerIDs\",\"type\":\"string\"}],\"internalType\":\"structGameRound.Info\",\"name\":\"GameInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"PlayerID\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"CurrentScore\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"WinScore\",\"type\":\"uint64\"}],\"internalType\":\"structGameRound.Result[]\",\"name\":\"GameResults\",\"type\":\"tuple[]\"}],\"internalType\":\"structGameRound.Replay\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"HashFunc\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"Address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayCID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PlayerIDs\",\"type\":\"string\"}],\"internalType\":\"structGameRound.Info\",\"name\":\"GameInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"PlayerID\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"CurrentScore\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"WinScore\",\"type\":\"uint64\"}],\"internalType\":\"structGameRound.Result[]\",\"name\":\"GameResults\",\"type\":\"tuple[]\"}],\"internalType\":\"structGameRound.Replay[]\",\"name\":\"_replays\",\"type\":\"tuple[]\"}],\"name\":\"saveGameReplay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611fcb806101065f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b14610063578063b970994814610081578063e1d563d9146100b1578063f2fde38b146100cd575b5f80fd5b6100616100e9565b005b61006b6100fc565b6040516100789190610d30565b60405180910390f35b61009b60048036038101906100969190610e96565b610123565b6040516100a89190611233565b60405180910390f35b6100cb60048036038101906100c691906117de565b6107f8565b005b6100e760048036038101906100e2919061184f565b610888565b005b6100f161090a565b6100fa5f610988565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61012b610c73565b5f60018360405161013c91906118b4565b9081526020016040518091039020604051806101000160405290815f82015f9054906101000a900460070b60070b60070b81526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546101b2906118f7565b80601f01602080910402602001604051908101604052809291908181526020018280546101de906118f7565b80156102295780601f1061020057610100808354040283529160200191610229565b820191905f5260205f20905b81548152906001019060200180831161020c57829003601f168201915b50505050508152602001600282018054610242906118f7565b80601f016020809104026020016040519081016040528092919081815260200182805461026e906118f7565b80156102b95780601f10610290576101008083540402835291602001916102b9565b820191905f5260205f20905b81548152906001019060200180831161029c57829003601f168201915b505050505081526020016003820180546102d2906118f7565b80601f01602080910402602001604051908101604052809291908181526020018280546102fe906118f7565b80156103495780601f1061032057610100808354040283529160200191610349565b820191905f5260205f20905b81548152906001019060200180831161032c57829003601f168201915b50505050508152602001600482018054610362906118f7565b80601f016020809104026020016040519081016040528092919081815260200182805461038e906118f7565b80156103d95780601f106103b0576101008083540402835291602001916103d9565b820191905f5260205f20905b8154815290600101906020018083116103bc57829003601f168201915b50505050508152602001600582016040518060800160405290815f82018054610401906118f7565b80601f016020809104026020016040519081016040528092919081815260200182805461042d906118f7565b80156104785780601f1061044f57610100808354040283529160200191610478565b820191905f5260205f20905b81548152906001019060200180831161045b57829003601f168201915b50505050508152602001600182018054610491906118f7565b80601f01602080910402602001604051908101604052809291908181526020018280546104bd906118f7565b80156105085780601f106104df57610100808354040283529160200191610508565b820191905f5260205f20905b8154815290600101906020018083116104eb57829003601f168201915b50505050508152602001600282018054610521906118f7565b80601f016020809104026020016040519081016040528092919081815260200182805461054d906118f7565b80156105985780601f1061056f57610100808354040283529160200191610598565b820191905f5260205f20905b81548152906001019060200180831161057b57829003601f168201915b505050505081526020016003820180546105b1906118f7565b80601f01602080910402602001604051908101604052809291908181526020018280546105dd906118f7565b80156106285780601f106105ff57610100808354040283529160200191610628565b820191905f5260205f20905b81548152906001019060200180831161060b57829003601f168201915b505050505081525050815260200160098201805480602002602001604051908101604052809291908181526020015f905b8282101561077b578382905f5260205f2090600202016040518060600160405290815f82018054610689906118f7565b80601f01602080910402602001604051908101604052809291908181526020018280546106b5906118f7565b80156107005780601f106106d757610100808354040283529160200191610700565b820191905f5260205f20905b8154815290600101906020018083116106e357829003601f168201915b50505050508152602001600182015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190610659565b505050508152505090505f816060015151118360405160200161079e9190611971565b604051602081830303815290604052906107ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e591906119da565b60405180910390fd5b5080915050919050565b61080061090a565b5f815111610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083a90611a44565b60405180910390fd5b5f5b81518110156108845761087182828151811061086457610863611a62565b5b6020026020010151610a49565b808061087c90611ac5565b915050610845565b5050565b61089061090a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f590611b7c565b60405180910390fd5b61090781610988565b50565b610912610c6c565b73ffffffffffffffffffffffffffffffffffffffff166109306100fc565b73ffffffffffffffffffffffffffffffffffffffff1614610986576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097d90611be4565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f60018260c0015160400151604051610a6291906118b4565b90815260200160405180910390209050815f0151815f015f6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055508160200151815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160400151816001019081610ae49190611d9f565b508160600151816002019081610afa9190611ec6565b508160800151816003019081610b109190611d9f565b508160a00151816004019081610b269190611d9f565b508160c00151816005015f820151815f019081610b439190611d9f565b506020820151816001019081610b599190611d9f565b506040820151816002019081610b6f9190611d9f565b506060820151816003019081610b859190611d9f565b509050505f5b8260e0015151811015610c6757816009018360e001518281518110610bb357610bb2611a62565b5b6020026020010151908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f019081610bf49190611d9f565b506020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050508080610c5f90611ac5565b915050610b8b565b505050565b5f33905090565b6040518061010001604052805f60070b81526020015f67ffffffffffffffff16815260200160608152602001606081526020016060815260200160608152602001610cbc610cc9565b8152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d1a82610cf1565b9050919050565b610d2a81610d10565b82525050565b5f602082019050610d435f830184610d21565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610da882610d62565b810181811067ffffffffffffffff82111715610dc757610dc6610d72565b5b80604052505050565b5f610dd9610d49565b9050610de58282610d9f565b919050565b5f67ffffffffffffffff821115610e0457610e03610d72565b5b610e0d82610d62565b9050602081019050919050565b828183375f83830152505050565b5f610e3a610e3584610dea565b610dd0565b905082815260208101848484011115610e5657610e55610d5e565b5b610e61848285610e1a565b509392505050565b5f82601f830112610e7d57610e7c610d5a565b5b8135610e8d848260208601610e28565b91505092915050565b5f60208284031215610eab57610eaa610d52565b5b5f82013567ffffffffffffffff811115610ec857610ec7610d56565b5b610ed484828501610e69565b91505092915050565b5f8160070b9050919050565b610ef281610edd565b82525050565b5f67ffffffffffffffff82169050919050565b610f1481610ef8565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610f51578082015181840152602081019050610f36565b5f8484015250505050565b5f610f6682610f1a565b610f708185610f24565b9350610f80818560208601610f34565b610f8981610d62565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f610fb882610f94565b610fc28185610f9e565b9350610fd2818560208601610f34565b610fdb81610d62565b840191505092915050565b5f608083015f8301518482035f8601526110008282610f5c565b9150506020830151848203602086015261101a8282610f5c565b915050604083015184820360408601526110348282610f5c565b9150506060830151848203606086015261104e8282610f5c565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f606083015f8301518482035f86015261109e8282610f5c565b91505060208301516110b36020860182610f0b565b5060408301516110c66040860182610f0b565b508091505092915050565b5f6110dc8383611084565b905092915050565b5f602082019050919050565b5f6110fa8261105b565b6111048185611065565b93508360208202850161111685611075565b805f5b85811015611151578484038952815161113285826110d1565b945061113d836110e4565b925060208a01995050600181019050611119565b50829750879550505050505092915050565b5f61010083015f8301516111795f860182610ee9565b50602083015161118c6020860182610f0b565b50604083015184820360408601526111a48282610f5c565b915050606083015184820360608601526111be8282610fae565b915050608083015184820360808601526111d88282610f5c565b91505060a083015184820360a08601526111f28282610f5c565b91505060c083015184820360c086015261120c8282610fe6565b91505060e083015184820360e086015261122682826110f0565b9150508091505092915050565b5f6020820190508181035f83015261124b8184611163565b905092915050565b5f67ffffffffffffffff82111561126d5761126c610d72565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f80fd5b61129381610edd565b811461129d575f80fd5b50565b5f813590506112ae8161128a565b92915050565b6112bd81610ef8565b81146112c7575f80fd5b50565b5f813590506112d8816112b4565b92915050565b5f67ffffffffffffffff8211156112f8576112f7610d72565b5b61130182610d62565b9050602081019050919050565b5f61132061131b846112de565b610dd0565b90508281526020810184848401111561133c5761133b610d5e565b5b611347848285610e1a565b509392505050565b5f82601f83011261136357611362610d5a565b5b813561137384826020860161130e565b91505092915050565b5f6080828403121561139157611390611282565b5b61139b6080610dd0565b90505f82013567ffffffffffffffff8111156113ba576113b9611286565b5b6113c684828501610e69565b5f83015250602082013567ffffffffffffffff8111156113e9576113e8611286565b5b6113f584828501610e69565b602083015250604082013567ffffffffffffffff81111561141957611418611286565b5b61142584828501610e69565b604083015250606082013567ffffffffffffffff81111561144957611448611286565b5b61145584828501610e69565b60608301525092915050565b5f67ffffffffffffffff82111561147b5761147a610d72565b5b602082029050602081019050919050565b5f606082840312156114a1576114a0611282565b5b6114ab6060610dd0565b90505f82013567ffffffffffffffff8111156114ca576114c9611286565b5b6114d684828501610e69565b5f8301525060206114e9848285016112ca565b60208301525060406114fd848285016112ca565b60408301525092915050565b5f61151b61151684611461565b610dd0565b9050808382526020820190506020840283018581111561153e5761153d61127e565b5b835b8181101561158557803567ffffffffffffffff81111561156357611562610d5a565b5b808601611570898261148c565b85526020850194505050602081019050611540565b5050509392505050565b5f82601f8301126115a3576115a2610d5a565b5b81356115b3848260208601611509565b91505092915050565b5f61010082840312156115d2576115d1611282565b5b6115dd610100610dd0565b90505f6115ec848285016112a0565b5f8301525060206115ff848285016112ca565b602083015250604082013567ffffffffffffffff81111561162357611622611286565b5b61162f84828501610e69565b604083015250606082013567ffffffffffffffff81111561165357611652611286565b5b61165f8482850161134f565b606083015250608082013567ffffffffffffffff81111561168357611682611286565b5b61168f84828501610e69565b60808301525060a082013567ffffffffffffffff8111156116b3576116b2611286565b5b6116bf84828501610e69565b60a08301525060c082013567ffffffffffffffff8111156116e3576116e2611286565b5b6116ef8482850161137c565b60c08301525060e082013567ffffffffffffffff81111561171357611712611286565b5b61171f8482850161158f565b60e08301525092915050565b5f61173d61173884611253565b610dd0565b905080838252602082019050602084028301858111156117605761175f61127e565b5b835b818110156117a757803567ffffffffffffffff81111561178557611784610d5a565b5b80860161179289826115bc565b85526020850194505050602081019050611762565b5050509392505050565b5f82601f8301126117c5576117c4610d5a565b5b81356117d584826020860161172b565b91505092915050565b5f602082840312156117f3576117f2610d52565b5b5f82013567ffffffffffffffff8111156118105761180f610d56565b5b61181c848285016117b1565b91505092915050565b61182e81610d10565b8114611838575f80fd5b50565b5f8135905061184981611825565b92915050565b5f6020828403121561186457611863610d52565b5b5f6118718482850161183b565b91505092915050565b5f81905092915050565b5f61188e82610f1a565b611898818561187a565b93506118a8818560208601610f34565b80840191505092915050565b5f6118bf8284611884565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061190e57607f821691505b602082108103611921576119206118ca565b5b50919050565b7f47616d65207265706c6179206e6f7420666f756e643a200000000000000000005f82015250565b5f61195b60178361187a565b915061196682611927565b601782019050919050565b5f61197b8261194f565b91506119878284611884565b915081905092915050565b5f82825260208201905092915050565b5f6119ac82610f1a565b6119b68185611992565b93506119c6818560208601610f34565b6119cf81610d62565b840191505092915050565b5f6020820190508181035f8301526119f281846119a2565b905092915050565b7f5f7265706c6179732063616e206e6f7420656d707479000000000000000000005f82015250565b5f611a2e601683611992565b9150611a39826119fa565b602082019050919050565b5f6020820190508181035f830152611a5b81611a22565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f611acf82611abc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b0157611b00611a8f565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611b66602683611992565b9150611b7182611b0c565b604082019050919050565b5f6020820190508181035f830152611b9381611b5a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611bce602083611992565b9150611bd982611b9a565b602082019050919050565b5f6020820190508181035f830152611bfb81611bc2565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611c5e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c23565b611c688683611c23565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611ca3611c9e611c9984611abc565b611c80565b611abc565b9050919050565b5f819050919050565b611cbc83611c89565b611cd0611cc882611caa565b848454611c2f565b825550505050565b5f90565b611ce4611cd8565b611cef818484611cb3565b505050565b5b81811015611d1257611d075f82611cdc565b600181019050611cf5565b5050565b601f821115611d5757611d2881611c02565b611d3184611c14565b81016020851015611d40578190505b611d54611d4c85611c14565b830182611cf4565b50505b505050565b5f82821c905092915050565b5f611d775f1984600802611d5c565b1980831691505092915050565b5f611d8f8383611d68565b9150826002028217905092915050565b611da882610f1a565b67ffffffffffffffff811115611dc157611dc0610d72565b5b611dcb82546118f7565b611dd6828285611d16565b5f60209050601f831160018114611e07575f8415611df5578287015190505b611dff8582611d84565b865550611e66565b601f198416611e1586611c02565b5f5b82811015611e3c57848901518255600182019150602085019450602081019050611e17565b86831015611e595784890151611e55601f891682611d68565b8355505b6001600288020188555050505b505050505050565b5f819050815f5260205f209050919050565b601f821115611ec157611e9281611e6e565b611e9b84611c14565b81016020851015611eaa578190505b611ebe611eb685611c14565b830182611cf4565b50505b505050565b611ecf82610f94565b67ffffffffffffffff811115611ee857611ee7610d72565b5b611ef282546118f7565b611efd828285611e80565b5f60209050601f831160018114611f2e575f8415611f1c578287015190505b611f268582611d84565b865550611f8d565b601f198416611f3c86611e6e565b5f5b82811015611f6357848901518255600182019150602085019450602081019050611f3e565b86831015611f805784890151611f7c601f891682611d68565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220f2e98cd879fa0cb147a5189857e28f963b8ba0b594d4857f782e1ce4dbc14ff464736f6c63430008150033",
}

// GameReplayContractABI is the input ABI used to generate the binding from.
// Deprecated: Use GameReplayContractMetaData.ABI instead.
var GameReplayContractABI = GameReplayContractMetaData.ABI

// GameReplayContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameReplayContractMetaData.Bin instead.
var GameReplayContractBin = GameReplayContractMetaData.Bin

// DeployGameReplayContract deploys a new Ethereum contract, binding an instance of GameReplayContract to it.
func DeployGameReplayContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameReplayContract, error) {
	parsed, err := GameReplayContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameReplayContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameReplayContract{GameReplayContractCaller: GameReplayContractCaller{contract: contract}, GameReplayContractTransactor: GameReplayContractTransactor{contract: contract}, GameReplayContractFilterer: GameReplayContractFilterer{contract: contract}}, nil
}

// GameReplayContract is an auto generated Go binding around an Ethereum contract.
type GameReplayContract struct {
	GameReplayContractCaller     // Read-only binding to the contract
	GameReplayContractTransactor // Write-only binding to the contract
	GameReplayContractFilterer   // Log filterer for contract events
}

// GameReplayContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameReplayContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameReplayContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameReplayContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameReplayContractSession struct {
	Contract     *GameReplayContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GameReplayContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameReplayContractCallerSession struct {
	Contract *GameReplayContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GameReplayContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameReplayContractTransactorSession struct {
	Contract     *GameReplayContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GameReplayContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameReplayContractRaw struct {
	Contract *GameReplayContract // Generic contract binding to access the raw methods on
}

// GameReplayContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameReplayContractCallerRaw struct {
	Contract *GameReplayContractCaller // Generic read-only contract binding to access the raw methods on
}

// GameReplayContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameReplayContractTransactorRaw struct {
	Contract *GameReplayContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameReplayContract creates a new instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContract(address common.Address, backend bind.ContractBackend) (*GameReplayContract, error) {
	contract, err := bindGameReplayContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameReplayContract{GameReplayContractCaller: GameReplayContractCaller{contract: contract}, GameReplayContractTransactor: GameReplayContractTransactor{contract: contract}, GameReplayContractFilterer: GameReplayContractFilterer{contract: contract}}, nil
}

// NewGameReplayContractCaller creates a new read-only instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractCaller(address common.Address, caller bind.ContractCaller) (*GameReplayContractCaller, error) {
	contract, err := bindGameReplayContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractCaller{contract: contract}, nil
}

// NewGameReplayContractTransactor creates a new write-only instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GameReplayContractTransactor, error) {
	contract, err := bindGameReplayContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractTransactor{contract: contract}, nil
}

// NewGameReplayContractFilterer creates a new log filterer instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GameReplayContractFilterer, error) {
	contract, err := bindGameReplayContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractFilterer{contract: contract}, nil
}

// bindGameReplayContract binds a generic wrapper to an already deployed contract.
func bindGameReplayContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameReplayContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameReplayContract *GameReplayContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameReplayContract.Contract.GameReplayContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameReplayContract *GameReplayContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.Contract.GameReplayContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameReplayContract *GameReplayContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameReplayContract.Contract.GameReplayContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameReplayContract *GameReplayContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameReplayContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameReplayContract *GameReplayContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameReplayContract *GameReplayContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameReplayContract.Contract.contract.Transact(opts, method, params...)
}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _replayID) view returns((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[]))
func (_GameReplayContract *GameReplayContractCaller) GetGameReplay(opts *bind.CallOpts, _replayID string) (GameRoundReplay, error) {
	var out []interface{}
	err := _GameReplayContract.contract.Call(opts, &out, "getGameReplay", _replayID)

	if err != nil {
		return *new(GameRoundReplay), err
	}

	out0 := *abi.ConvertType(out[0], new(GameRoundReplay)).(*GameRoundReplay)

	return out0, err

}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _replayID) view returns((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[]))
func (_GameReplayContract *GameReplayContractSession) GetGameReplay(_replayID string) (GameRoundReplay, error) {
	return _GameReplayContract.Contract.GetGameReplay(&_GameReplayContract.CallOpts, _replayID)
}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _replayID) view returns((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[]))
func (_GameReplayContract *GameReplayContractCallerSession) GetGameReplay(_replayID string) (GameRoundReplay, error) {
	return _GameReplayContract.Contract.GetGameReplay(&_GameReplayContract.CallOpts, _replayID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameReplayContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractSession) Owner() (common.Address, error) {
	return _GameReplayContract.Contract.Owner(&_GameReplayContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractCallerSession) Owner() (common.Address, error) {
	return _GameReplayContract.Contract.Owner(&_GameReplayContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameReplayContract.Contract.RenounceOwnership(&_GameReplayContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameReplayContract.Contract.RenounceOwnership(&_GameReplayContract.TransactOpts)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xe1d563d9.
//
// Solidity: function saveGameReplay((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[])[] _replays) returns()
func (_GameReplayContract *GameReplayContractTransactor) SaveGameReplay(opts *bind.TransactOpts, _replays []GameRoundReplay) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "saveGameReplay", _replays)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xe1d563d9.
//
// Solidity: function saveGameReplay((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[])[] _replays) returns()
func (_GameReplayContract *GameReplayContractSession) SaveGameReplay(_replays []GameRoundReplay) (*types.Transaction, error) {
	return _GameReplayContract.Contract.SaveGameReplay(&_GameReplayContract.TransactOpts, _replays)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xe1d563d9.
//
// Solidity: function saveGameReplay((int64,uint64,string,bytes,string,string,(string,string,string,string),(string,uint64,uint64)[])[] _replays) returns()
func (_GameReplayContract *GameReplayContractTransactorSession) SaveGameReplay(_replays []GameRoundReplay) (*types.Transaction, error) {
	return _GameReplayContract.Contract.SaveGameReplay(&_GameReplayContract.TransactOpts, _replays)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.Contract.TransferOwnership(&_GameReplayContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.Contract.TransferOwnership(&_GameReplayContract.TransactOpts, newOwner)
}

// GameReplayContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameReplayContract contract.
type GameReplayContractOwnershipTransferredIterator struct {
	Event *GameReplayContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameReplayContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameReplayContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameReplayContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameReplayContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameReplayContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameReplayContractOwnershipTransferred represents a OwnershipTransferred event raised by the GameReplayContract contract.
type GameReplayContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameReplayContract *GameReplayContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameReplayContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameReplayContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractOwnershipTransferredIterator{contract: _GameReplayContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameReplayContract *GameReplayContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameReplayContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameReplayContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameReplayContractOwnershipTransferred)
				if err := _GameReplayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameReplayContract *GameReplayContractFilterer) ParseOwnershipTransferred(log types.Log) (*GameReplayContractOwnershipTransferred, error) {
	event := new(GameReplayContractOwnershipTransferred)
	if err := _GameReplayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
