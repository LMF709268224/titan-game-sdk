package test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	contracts "github.com/zscboy/titan-game-sdk/contracts/api"
	"github.com/zscboy/titan-game-sdk/contracts/client"
	"github.com/zscboy/titan-game-sdk/storage"
	"github.com/zscboy/titan-game-sdk/vrf/filrpc"
	"github.com/zscboy/titan-game-sdk/vrf/gamevrf"
)

var (
	chainHeight = int64(1009859)
	// filecoin bls private key
	filPrivateKey = []byte{152, 112, 35, 145, 21, 31, 99, 206, 204, 113, 33, 99, 241, 180, 157, 194, 91, 224, 34, 186, 137, 10, 136, 38, 133, 32, 109, 255, 59, 81, 45, 26}
	// filecoin bls public key
	filPublicKey = []byte{146, 209, 52, 147, 166, 127, 130, 148, 172, 13, 162, 254, 17, 85, 254, 151, 93, 182, 28, 218, 103, 106, 200, 115, 178, 101, 156, 74, 25, 214, 220, 136, 167, 32, 147, 231, 40, 250, 149, 109, 229, 58, 7, 135, 214, 93, 55, 169}
	replayData   = []byte{166, 151, 5, 61, 189, 201, 203, 69, 188, 20, 9, 50, 223, 153, 238, 59, 149, 71, 92, 205, 245, 57, 9, 168, 156, 163, 49, 215, 203, 159, 209, 245, 110, 78, 130, 62, 224, 136, 188, 64, 79, 245, 145, 21, 119, 13, 43, 8, 3, 231, 35, 65, 212, 42, 11, 44, 247, 146, 120, 206, 82, 252, 203, 131, 1, 13, 150, 229, 244, 12, 165, 170, 77, 27, 239, 148, 184, 106, 124, 46, 182, 222, 112, 241, 205, 168, 133, 58, 106, 104, 70, 68, 250, 70, 84, 27}

	ethPrivateKey   = "9382b3739fb3173143d35afb0f2b1ddc6cf1713db110563832c2b7b9108ffb28"
	contractAddress = "0x3f1AbDFC7f94b4F20412f142f0a42DA4E7fcf57a" // "0x5D7990C0487C57E3a0b57f2d3472600c37a5eE98"

	nodeURL = "http://172.25.9.91:1251/rpc/v1" //"https://api.calibration.node.glif.io/"

	sentCount     = 0
	receivedCount = 0
	gameInfoCount = 500
)

func TestOnGameServer(t *testing.T) {
	client := filrpc.New(
		filrpc.NodeURLOption(nodeURL),
	)

	tps, err := client.ChainGetTipSetByHeight(chainHeight)
	if err != nil {
		t.Fatal(err)
	}

	privateKey := filPrivateKey
	// publicKey := filPublicKey

	var entropy []byte
	gameRoundInfo := GameRoundInfo{
		GameID:    "abc-efg-hi",
		PlayerIDs: "a,b,c,d",
		RoundID:   "gogogogo1",
		ReplayID:  "bilibili",
	}

	buf := new(bytes.Buffer)
	err = gameRoundInfo.MarshalCBOR(buf)
	if err != nil {
		t.Fatal(err)
	}
	entropy = buf.Bytes()

	vrfout, err := gamevrf.FilGenerateVRFByTipSet(gamevrf.DomainSeparationTag_GameBasic, privateKey, tps, entropy)
	if err != nil {
		t.Fatal(err)
	}

	cid, err := storage.CalculateCID(bytes.NewReader(replayData))
	if err != nil {
		t.Fatal(err)
	}

	addr, err := address.NewBLSAddress(filPublicKey)
	if err != nil {
		t.Fatal(err)
	}

	// playing Game, and generate game result
	replay := &contracts.GameRoundReplay{
		DomainSeparationTag: int64(gamevrf.DomainSeparationTag_GameBasic),
		VRFHeight:           uint64(chainHeight),
		HashFunc:            "blake2b",
		VRFProof:            vrfout.Proof,
		Address:             addr.String(),
		ReplayCID:           cid.String(),
		GameInfo:            gameRoundInfoToContractGameRoundInfo(gameRoundInfo),
		GameResults:         generateGameResults(&gameRoundInfo),
	}

	if err = saveGameReplyWithContract(replay); err != nil {
		t.Fatal(err)
	}
}

func gameRoundInfoToContractGameRoundInfo(gameRoundInfo GameRoundInfo) contracts.GameRoundInfo {
	return contracts.GameRoundInfo{
		GameID:    gameRoundInfo.GameID,
		RoundID:   gameRoundInfo.RoundID,
		ReplayID:  gameRoundInfo.ReplayID,
		PlayerIDs: gameRoundInfo.PlayerIDs,
	}
}

func generateGameResults(gameRoundInfo *GameRoundInfo) []contracts.GameRoundResult {
	playerIDs := gameRoundInfo.PlayerIDs
	players := strings.Split(playerIDs, ",")

	results := make([]contracts.GameRoundResult, 0, len(players))
	for _, player := range players {
		result := contracts.GameRoundResult{
			PlayerID:     player,
			CurrentScore: 100,
			WinScore:     1,
		}
		results = append(results, result)
	}
	return results
}

// You have to deploy the contract before you can do that.
func saveGameReplyWithContract(replay *contracts.GameRoundReplay) error {
	c, err := client.New(
		client.PrivateKeyOption(ethPrivateKey),
		client.EndpointOption(nodeURL),
	)
	if err != nil {
		return err
	}

	// replayID := uuid.NewString()
	gameContractAddress := common.HexToAddress(contractAddress)
	instance, err := contracts.NewGameReplayContract(gameContractAddress, c.EthClient())
	if err != nil {
		return err
	}

	result, err := c.InvokeContract(0, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return instance.SaveGameReplay(opts, []contracts.GameRoundReplay{*replay})
	})
	if err != nil {
		return err
	}

	fmt.Println("replay id: ", replay.GameInfo.ReplayID)
	fmt.Println("save game replay OK: ", string(result))
	fmt.Println("querying game replay...")

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-time.After(2 * time.Minute):
			return fmt.Errorf("query order timeout!")
		case <-ticker.C:
			order, err := instance.GetGameReplay(&bind.CallOpts{
				Pending: true,
			}, replay.GameInfo.ReplayID)
			if err != nil {
				fmt.Println("get game replay err ", err.Error())
				continue
			}

			fmt.Println("Query game replay OK: ", order)
			return nil
		}
	}
}

func TestOnGameUser(t *testing.T) {
	c, err := client.New(
		client.PrivateKeyOption(ethPrivateKey),
		client.EndpointOption(nodeURL),
	)
	if err != nil {
		t.Fatal("new client err ", err.Error())
	}

	gameContractAddress := common.HexToAddress(contractAddress)
	instance, err := contracts.NewGameReplayContract(gameContractAddress, c.EthClient())
	if err != nil {
		t.Fatal("new contract instance err ", err.Error())
	}

	if instance == nil {
		t.Fatal("instance == nil")
	}

	replayID := "bilibili"
	replay, err := instance.GetGameReplay(nil, replayID)
	if err != nil {
		t.Fatal("new client err ", err.Error())
	}

	addr, err := address.NewFromString(replay.Address)
	if err != nil {
		t.Fatal(err)
	}

	filClient := filrpc.New(
		filrpc.NodeURLOption(nodeURL),
	)

	tps, err := filClient.ChainGetTipSetByHeight(chainHeight)
	if err != nil {
		t.Fatal(err)
	}

	gameRoundInfo := contractGameRoundInfoToGameRoundInfo(&replay.GameInfo)
	buf := new(bytes.Buffer)
	err = gameRoundInfo.MarshalCBOR(buf)
	if err != nil {
		t.Fatal(err)
	}
	entropy := buf.Bytes()

	vrfout := &gamevrf.VRFOut{Height: replay.VRFHeight, Proof: replay.VRFProof}
	err = gamevrf.FilVerifyVRFByTipSet(gamevrf.DomainSeparationTag(replay.DomainSeparationTag), addr, tps, entropy, vrfout)
	if err != nil {
		t.Fatal(err)
	}
}

func contractGameRoundInfoToGameRoundInfo(roundInfo *contracts.GameRoundInfo) GameRoundInfo {
	return GameRoundInfo{
		GameID:    roundInfo.GameID,
		RoundID:   roundInfo.RoundID,
		ReplayID:  roundInfo.ReplayID,
		PlayerIDs: roundInfo.PlayerIDs,
	}
}

func TestSaveGameReplays(t *testing.T) {
	fClient := filrpc.New(
		filrpc.NodeURLOption(nodeURL),
	)

	tps, err := fClient.ChainHead()
	if err != nil {
		fmt.Println(err)
		return
	}

	privateKey := filPrivateKey
	// publicKey := filPublicKey

	var entropy []byte
	gameRoundInfo := GameRoundInfo{
		GameID:    "abc-efg-hi",
		PlayerIDs: "a,b,c,d",
		RoundID:   "gogogogo1",
		ReplayID:  "bilibili",
	}

	buf := new(bytes.Buffer)
	err = gameRoundInfo.MarshalCBOR(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	entropy = buf.Bytes()

	vrfout, err := gamevrf.FilGenerateVRFByTipSet(gamevrf.DomainSeparationTag_GameBasic, privateKey, tps, entropy)
	if err != nil {
		fmt.Println(err)
		return
	}

	cid, err := storage.CalculateCID(bytes.NewReader(replayData))
	if err != nil {
		fmt.Println(err)
		return
	}

	addr, err := address.NewBLSAddress(filPublicKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// playing Game, and generate game resulPrintln(a)
	replay := &contracts.GameRoundReplay{
		DomainSeparationTag: int64(gamevrf.DomainSeparationTag_GameBasic),
		VRFHeight:           uint64(chainHeight),
		HashFunc:            "blake2b",
		VRFProof:            vrfout.Proof,
		Address:             addr.String(),
		ReplayCID:           cid.String(),
		GameInfo:            gameRoundInfoToContractGameRoundInfo(gameRoundInfo),
		GameResults:         generateGameResults(&gameRoundInfo),
	}

	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
		client.EndpointOption(nodeURL),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	nonce, err := c.Nonce()
	if err != nil {
		return
	}

	// if err = saveGameReplyWithContract2(nonce, *replay); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("nonce : ", nonce)
	var gameMap sync.Map

	for i := 0; i < gameInfoCount; i++ {
		n := nonce

		key := fmt.Sprintf("r_%d", n)
		gameMap.Store(key, nil)

		nonce++

		go saveGameReplyWithContract2(n, *replay)

	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	gameContractAddress := common.HexToAddress(contractAddress)
	instance, err := contracts.NewGameReplayContract(gameContractAddress, c.EthClient())
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case <-time.After(2 * time.Minute):
			return
		case <-ticker.C:
			gameMap.Range(func(key, value interface{}) bool {
				replayID := key.(string)
				_, err := instance.GetGameReplay(&bind.CallOpts{
					Pending: true,
				}, replayID)
				if err == nil {
					// 	fmt.Println(replay.GameInfo.ReplayID, " get game replay err ", err.Error())
					// } else {
					receivedCount++
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " Query game replay OK: ", replayID, " receivedCount: ", receivedCount)

					gameMap.Delete(replayID)
				}

				return true
			})

			if receivedCount >= sentCount {
				fmt.Println("end...")
				return
			}
		}
	}
}

// You have to deploy the contract before you can do that.
func saveGameReplyWithContract2(nonce uint64, replay contracts.GameRoundReplay) error {
	// fmt.Println("send nonce :", nonce)
	c, err := client.New(
		client.PrivateKeyOption(os.Getenv("PRIVATE_KEY")),
		client.EndpointOption(nodeURL),
	)
	if err != nil {
		fmt.Println("err :", err)
		return err
	}

	// replayID := uuid.NewString()
	gameContractAddress := common.HexToAddress(contractAddress)
	instance, err := contracts.NewGameReplayContract(gameContractAddress, c.EthClient())
	if err != nil {
		fmt.Println("NewGameReplayContract :", err)
		return err
	}

	replay.GameInfo.ReplayID = fmt.Sprintf("r_%d", nonce)
	list := []contracts.GameRoundReplay{replay}
	// for i := 0; i < 34; i++ {
	// 	list = append(list, replay)
	// }

	_, err = c.InvokeContract(nonce, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		result, err := instance.SaveGameReplay(opts, list)
		if err != nil {
			fmt.Println("SaveGameReplay :", err)
			return nil, err
		}
		fmt.Println("querying Hash : ", result.Hash())
		return result, nil
	})
	if err != nil {
		fmt.Println("InvokeContract :", err)
		return err
	}

	// fmt.Println("replay id: ", replay.GameInfo.ReplayID)
	// fmt.Println("save game replay OK: ", string(result))
	// fmt.Println("querying game replay...")

	sentCount++
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " send nonce :", nonce, " sentCount: ", sentCount)

	return nil
}
