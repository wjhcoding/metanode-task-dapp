// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Counter {
    uint256 private count;

    event CountIncremented(uint256 newCount);
    event CountDecremented(uint256 newCount);

    function increment() external {
        count += 1;
        emit CountIncremented(count);
    }

    function decrement() external {
        require(count > 0, "Counter: count cannot be negative");
        count -= 1;
        emit CountDecremented(count);
    }

    function getCount() external view returns (uint256) {
        return count;
    }
}