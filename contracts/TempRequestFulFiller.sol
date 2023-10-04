// SPDX-License-Identifier: LGPL 3.0
pragma solidity 0.8.8;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./BLS.sol";

contract RequestFulFiller is Ownable {
    struct SignatureMessage {
        uint256 chainId;
        address verifier;
        address recipient;
        uint256 data;
    }

    event AggregatedPubKeyRegistered(bytes32 indexed aggPubKeyHash);
    event PacketVerified(address indexed recipient, uint256 packetId);
    event RequestFulFilled(address indexed receiver, uint256 indexed data);

    mapping(bytes32 => bool) public isMembers;
    mapping(bytes32 => bool) public isRegisteredAggPubKey;
    mapping(uint256 => bool) public isUsedData;

    constructor(uint256[4][] memory members) {
        uint16 totalMembers = uint16(members.length);
        for (uint16 i = 0; i < totalMembers; ) {
            isMembers[keccak256(abi.encode(members[i]))] = true;
            unchecked {
                i++;
            }
        }
    }

    function fullFillRequestUsingUnRegisteredAggPubKey(
        uint256[4][] calldata signersPubKey,
        uint256[2] calldata aggregatedSignature,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2,
        address receiver,
        uint256 data
    ) external {
        require(
            isUsedData[data] == false,
            "Passed Data Is Already Used, Passed Data Must Be All Time Unique"
        );
        bytes32 aggPubKeyG2Hash = keccak256(abi.encode(aggregatedPubKeyG2));
        // require(
        //     isRegisteredAggPubKey[aggPubKeyG2Hash] == false,
        //     "Aggregated PubKey Is Already Registered, Use `verifyProposalFromRegisteredAggPubKey` Method"
        // );
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        uint16 totalSigners = uint16(signersPubKey.length);
        for (uint16 i = 0; i < totalSigners; ) {
            require(
                isMembers[keccak256(abi.encode(signersPubKey[i]))] == true,
                "Passed Signer PubKey Is Not A Member"
            );
            unchecked {
                i++;
            }
        }
        require(
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggregatedPubKeyG1,
                signersPubKey
            ) == true,
            "Invalid AggregatedPubKeyG2"
        );
        require(
            BLS.verifySingle(
                aggregatedSignature,
                aggregatedPubKeyG2,
                BLS.hashToPoint(
                    abi.encode(
                        SignatureMessage(
                            block.chainid,
                            address(this),
                            receiver,
                            data
                        )
                    )
                )
            ) == true,
            "Signature Verification Failed"
        );
        isRegisteredAggPubKey[aggPubKeyG2Hash] = true;
        emit AggregatedPubKeyRegistered(aggPubKeyG2Hash);
        (bool success, ) = receiver.call(
            abi.encodeWithSignature("handleRequestFulfillment(uint256)", data)
        );
        require(
            success == true,
            "Failed To Successfully Deliver Data To Receiver, Something Wrong From Receiver Side"
        );
        isUsedData[data] = true;
        emit RequestFulFilled(receiver, data);
    }

    function fullFillRequestUsingRegisteredAggPubKey(
        uint256[2] calldata aggregatedSignature,
        uint256[4] calldata registeredAggPubKeyG2,
        address receiver,
        uint256 data
    ) external {
        require(
            isUsedData[data] == false,
            "Passed Data Is Already Used, Passed Data Must Be All Time Unique"
        );
        bytes32 aggPubKeyG2Hash = keccak256(abi.encode(registeredAggPubKeyG2));
        require(
            isRegisteredAggPubKey[aggPubKeyG2Hash] == true,
            "Aggregated PubKey Is Not Registered"
        );
        require(
            BLS.verifySingle(
                aggregatedSignature,
                registeredAggPubKeyG2,
                BLS.hashToPoint(
                    abi.encode(
                        SignatureMessage(
                            block.chainid,
                            address(this),
                            receiver,
                            data
                        )
                    )
                )
            ) == true,
            "Signature Verification Failed"
        );
        (bool success, ) = receiver.call(
            abi.encodeWithSignature("handleRequestFulfillment(uint256)", data)
        );
        require(
            success == true,
            "Failed To Successfully Deliver Data To Receiver, Something Wrong From Receiver Side"
        );
        isUsedData[data] = true;
        emit RequestFulFilled(receiver, data);
    }

    function registerAggPubKeyFromBasePubKeys(
        uint256[4][] calldata signersPubKey,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2
    ) external {
        bytes32 aggPubKeyG2Hash = keccak256(abi.encode(aggregatedPubKeyG2));
        require(
            isRegisteredAggPubKey[aggPubKeyG2Hash] == false,
            "Aggregated PubKey Is Already Registered"
        );
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        uint16 totalSigners = uint16(signersPubKey.length);
        for (uint16 i = 0; i < totalSigners; ) {
            require(
                isMembers[keccak256(abi.encode(signersPubKey[i]))] == true,
                "Passed Signer PubKey Is Not A Member"
            );
            unchecked {
                i++;
            }
        }
        require(
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggregatedPubKeyG1,
                signersPubKey
            ) == true,
            "Invalid AggregatedPubKeyG2"
        );
        isRegisteredAggPubKey[aggPubKeyG2Hash] = true;
        emit AggregatedPubKeyRegistered(aggPubKeyG2Hash);
    }

    function registerAggPubKeyFromRegisteredAggPubKeys(
        uint256[4][] calldata registeredAggPubKeys,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2
    ) external {
        bytes32 aggPubKeyG2Hash = keccak256(abi.encode(aggregatedPubKeyG2));
        require(
            isRegisteredAggPubKey[aggPubKeyG2Hash] == false,
            "Aggregated PubKey Is Already Registered"
        );
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        uint16 totalAggPubKeysPassed = uint16(registeredAggPubKeys.length);
        for (uint16 i = 0; i < totalAggPubKeysPassed; ) {
            require(
                isRegisteredAggPubKey[
                    keccak256(abi.encode(registeredAggPubKeys[i]))
                ] == true,
                "Invalid Registered Aggregated PubKeys Are Passed"
            );
            unchecked {
                i++;
            }
        }
        require(
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggregatedPubKeyG1,
                registeredAggPubKeys
            ) == true,
            "Invalid AggregatedPubKeyG2"
        );
        isRegisteredAggPubKey[aggPubKeyG2Hash] = true;
        emit AggregatedPubKeyRegistered(aggPubKeyG2Hash);
    }

    function registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(
        uint256[4][] calldata signersPubKey,
        uint256[4][] calldata registeredAggPubKeys,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2
    ) external {
        bytes32 aggPubKeyG2Hash = keccak256(abi.encode(aggregatedPubKeyG2));
        require(
            isRegisteredAggPubKey[aggPubKeyG2Hash] == false,
            "Aggregated PubKey Is Already Registered"
        );
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        uint16 totalAggPubKeysPassed = uint16(registeredAggPubKeys.length);
        uint16 totalBasePubKeys = uint16(signersPubKey.length);
        uint256[4][] memory basePubKeys;
        basePubKeys = new uint256[4][](
            (totalBasePubKeys + totalAggPubKeysPassed)
        );
        uint16 curBaseCount;
        while (curBaseCount < totalBasePubKeys) {
            require(
                isMembers[keccak256(abi.encode(signersPubKey[curBaseCount]))] ==
                    true,
                "Passed Signer PubKey Is Not A Member PubKey"
            );
            basePubKeys[curBaseCount] = signersPubKey[curBaseCount];
            unchecked {
                curBaseCount++;
            }
        }
        for (uint16 j = 0; j < totalAggPubKeysPassed; ) {
            require(
                isRegisteredAggPubKey[
                    keccak256(abi.encode(registeredAggPubKeys[j]))
                ] == true,
                "Invalid Registered Aggregated PubKeys Are Passed"
            );
            basePubKeys[curBaseCount] = registeredAggPubKeys[j];
            unchecked {
                curBaseCount++;
                j++;
            }
        }
        require(
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggregatedPubKeyG1,
                basePubKeys
            ) == true,
            "Invalid AggregatedPubKeyG2"
        );
        isRegisteredAggPubKey[aggPubKeyG2Hash] = true;
        emit AggregatedPubKeyRegistered(aggPubKeyG2Hash);
    }

    /********** Setter Functions **********/

    function setNewMembers(uint256[4][] memory members) external onlyOwner {
        uint16 totalMembers = uint16(members.length);
        for (uint16 i = 0; i < totalMembers; ) {
            isMembers[keccak256(abi.encode(members[i]))] = true;
            unchecked {
                i++;
            }
        }
    }

    function removeMembers(uint256[4][] memory members) external onlyOwner {
        uint16 totalMembers = uint16(members.length);
        for (uint16 i = 0; i < totalMembers; ) {
            isMembers[keccak256(abi.encode(members[i]))] = false;
            unchecked {
                i++;
            }
        }
    }

    /********** View Functions **********/

    function isValidPublicKey(
        uint256[4] memory publicKey
    ) external pure returns (bool) {
        return BLS.isValidPublicKey(publicKey);
    }

    function isValidSignature(
        uint256[2] memory signature
    ) external pure returns (bool) {
        return BLS.isValidSignature(signature);
    }

    function getMessageData(
        address recipient,
        uint256 data
    ) external view returns (bytes memory, bytes memory) {
        uint256[2] memory hashToPointData = BLS.hashToPoint(
            abi.encode(
                SignatureMessage(block.chainid, address(this), recipient, data)
            )
        );
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function hashToPoint(
        bytes memory data
    ) external view returns (bytes memory, bytes memory) {
        uint256[2] memory hashToPointData = BLS.hashToPoint(data);
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function verifyHelpedAggregationPublicKeys(
        uint256[2] memory pubkeyG1,
        uint256[4] memory pubkeyG2
    ) external view returns (bool) {
        return BLS.verifyHelpedAggregationPublicKeys(pubkeyG1, pubkeyG2);
    }

    function verifyHelpedAggregationPublicKeysMultiple(
        uint256[2] memory aggPubkeyG1,
        uint256[4][] memory pubkeysG2
    ) external view returns (bool) {
        return
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggPubkeyG1,
                pubkeysG2
            );
    }
}
// 0xC072FA167EAc058c30701976689B8817B216270B