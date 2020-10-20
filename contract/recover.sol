pragma solidity ^0.7.0;

contract recover {

    function ownerAddress(bytes32 hash, bytes32 r, bytes32 s, uint8 v ) external pure returns address {
        return ecrecover(hash, v, r, s);
    }
}