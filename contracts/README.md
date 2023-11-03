# titan-contracts

## Build contracts
    solcjs contracts/GameReplay.sol --output-dir ./build --bin --abi
    abigen --abi build/contracts_GameReplay_sol_GameReplayContract.abi --bin build/contracts_GameReplay_sol_GameReplayContract.bin --pkg contracts --type GameReplayContract --out ./api/game_replay.go
