// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../packages/openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../../packages/openzeppelin/contracts/access/Ownable.sol";

contract MyNFTToken is ERC721, Ownable {
    constructor() ERC721("MyNFTToken", "MTK") {}

    function safeMint(address to, uint256 tokenId) public onlyOwner {
        _safeMint(to, tokenId);
    }
}
