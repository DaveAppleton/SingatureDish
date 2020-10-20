pragma solidity ^0.7.0;
// SPDX-License-Identifier: Nun
contract recover {

    function ownerAddress(bytes32 hash, bytes32 r, bytes32 s, uint8 v ) public pure returns (address) {
        return ecrecover(hash, v, r, s);
    }
    
    /// @notice Extracts the r, s, and v components from the `sigData` field starting from the `offset`
    /// @dev Note: does not do any bounds checking on the arguments!
    /// @param sigData the signature data; could be 1 or more packed signatures.
    /// @param offset the offset in sigData from which to start unpacking the signature components.
    function extractSignature(bytes memory sigData, uint256 offset) internal pure returns  (bytes32 r, bytes32 s, uint8 v) {
        // Divide the signature in r, s and v variables
        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        // solium-disable-next-line security/no-inline-assembly
        assembly {
             let dataPointer := add(sigData, offset)
             r := mload(add(dataPointer, 0x20))
             s := mload(add(dataPointer, 0x40))
             v := byte(0, mload(add(dataPointer, 0x60)))
        }
    
        return (r, s, v);
    }
    
    function ownerAddressFromSig(bytes32 hash, bytes memory sigData) external pure returns (address) {
        bytes32 r;
        bytes32 s;
        uint8   v;
        (r,s,v) = extractSignature(sigData,0);
        return ownerAddress(hash,r,s,v);
    }
}