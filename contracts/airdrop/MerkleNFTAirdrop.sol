// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../../packages/openzeppelin/contracts/access/Ownable.sol";
import "../../packages/openzeppelin/contracts/utils/math/SafeMath.sol";
import "../../packages/openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "../../packages/openzeppelin/contracts/token/ERC721/IERC721.sol";

contract MerkleNFTAirdrop is Ownable {
    using SafeMath for uint256;

    event Claimed(address claimant, uint256 week, uint256 balance);
    event TrancheAdded(uint256 tranche, bytes32 merkleRoot);
    event TrancheExpired(uint256 tranche);

    IERC721 public token;

    mapping(uint256 => bytes32) public merkleRoots;
    mapping(uint256 => mapping(address => bool)) public claimed;
    uint256 public tranches;

    constructor(IERC721 _token) {
        token = _token;
    }

    function seedNewTranche(bytes32 _merkleRoot)
    public onlyOwner returns (uint256 trancheId)
    {
        trancheId = tranches;
        merkleRoots[trancheId] = _merkleRoot;

        tranches = tranches.add(1);

        emit TrancheAdded(trancheId, _merkleRoot);
    }

    function expireTranche(uint256 _trancheId)
    public onlyOwner
    {
        merkleRoots[_trancheId] = bytes32(0);
        emit TrancheExpired(_trancheId);
    }

    function redeem(
        address _account,
        uint256 _tranche,
        uint256 _tokenId,
        bytes32[] memory _merkleProof
    )
    public
    {
        _redeem(_account, _tranche, _tokenId, _merkleProof);
        _disburse(_account, _tokenId);
    }

    function redeemMany(
        address _account,
        uint256[] memory _tranches,
        uint256[] memory _tokenIds,
        bytes32[][] memory _merkleProofs
    )
    public
    {
        uint256 len = _tranches.length;
        require(len == _tokenIds.length && len == _merkleProofs.length, "Mismatching inputs");

        for (uint256 i = 0; i < len; i++) {
            _redeem(_account, _tranches[i], _tokenIds[i], _merkleProofs[i]);
            _disburse(_account, _tokenIds[i]);
        }
    }

    function verify(
        address _account,
        uint256 _tranche,
        uint256 _tokenId,
        bytes32[] memory _merkleProof
    )
    public view returns (bool valid)
    {
        return _verify(_account, _tranche, _tokenId, _merkleProof);
    }

    function _redeem(
        address _account,
        uint256 _tranche,
        uint256 _tokenId,
        bytes32[] memory _merkleProof
    )
    internal
    {
        require(_tranche < tranches, "Tranche cannot be in the future");

        require(!claimed[_tranche][_account], "Account has already claimed");
        require(_verify(_account, _tranche, _tokenId, _merkleProof), "Incorrect merkle proof");

        claimed[_tranche][_account] = true;

        emit Claimed(_account, _tranche, _tokenId);
    }

    function _verify(
        address _account,
        uint256 _tranche,
        uint256 _tokenId,
        bytes32[] memory _merkleProof
    )
    internal view returns (bool valid)
    {
        bytes32 leaf = _leaf(_account, _tokenId);
        return MerkleProof.verify(_merkleProof, merkleRoots[_tranche], leaf);
    }

    function _disburse(
        address _account,
        uint256 _tokenId
    )
    private
    {
        require(token.ownerOf(_tokenId) != address(0), "ERC721: approved query for nonexistent token");
        token.safeTransferFrom(_msgSender(), _account, _tokenId);
    }

    function _leaf(address account, uint256 tokenId)
    internal pure returns (bytes32)
    {
        return keccak256(abi.encodePacked(tokenId, account));
    }
}
