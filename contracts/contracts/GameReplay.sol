// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";
import "./GameRound.sol";

contract GameReplayContract is Ownable {
    mapping(string => GameRound.Replay) gameReplayMap;

    function saveGameReplay(GameRound.Replay[] memory _replays) public onlyOwner {
        require(_replays.length > 0, "_replays can not empty");

        for (uint256 i = 0; i < _replays.length; i++) {
            copyReplayToStorage(_replays[i]);
        }
    }

    function copyReplayToStorage(GameRound.Replay memory _replay) internal {
        GameRound.Replay storage storageReplay = gameReplayMap[_replay.GameInfo.ReplayID];
        storageReplay.DomainSeparationTag = _replay.DomainSeparationTag;
        storageReplay.VRFHeight = _replay.VRFHeight;
        storageReplay.HashFunc = _replay.HashFunc;
        storageReplay.VRFProof = _replay.VRFProof;
        storageReplay.Address = _replay.Address;
        storageReplay.ReplayCID = _replay.ReplayCID;
        storageReplay.GameInfo = _replay.GameInfo;

        for (uint256 i = 0; i < _replay.GameResults.length; i++) {
            storageReplay.GameResults.push(_replay.GameResults[i]);
        }
    }

    function getGameReplay(string memory _replayID) public view returns (GameRound.Replay memory) {
        GameRound.Replay memory replay = gameReplayMap[_replayID];
        require(bytes(replay.Address).length > 0,  string(abi.encodePacked("Game replay not found: ", _replayID)));
        
        return replay;
    }
}