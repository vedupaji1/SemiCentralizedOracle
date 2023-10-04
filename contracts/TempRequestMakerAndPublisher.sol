// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.8;

contract RequestReceiver {
    event ReceivedRequest(address indexed recepientContract);

    function createRequest(address recepientContract) external {
        emit ReceivedRequest(recepientContract);
    }
}
// 0x26226C0B001471910bDFcDC3B4A07e659dC7eD66

contract DataReceiver {
    event ReceivedData(uint256 data);

    function handleRequestFulfillment(uint256 data) external {
        emit ReceivedData(data);
    }

    receive() external payable {}
}
// 0x003dBE12277abcfe5B1e41f7454dc75EF3FC4dA8