const {ethers} = require('ethers')
const {MerkleTree} = require('merkletreejs')
const keccak256 = require('keccak256')
const tokens = require('./tokens.json')

function hashToken(tokenId, account) {
    return Buffer.from(ethers.utils.solidityKeccak256(['uint256', 'address'], [tokenId, account]).slice(2), 'hex')
}

const merkleTree = new MerkleTree(Object.entries(tokens).map(token => hashToken(...token)), keccak256, {sortPairs: true})

console.log(merkleTree.getHexRoot())

for (const [tokenId, account] of Object.entries(tokens)) {
    console.log(tokenId, account)
    const proof = merkleTree.getHexProof(hashToken(tokenId, account))
    console.log(proof)
}
