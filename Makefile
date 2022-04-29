token:
	solc --abi --bin contracts/token/MyNFTToken.sol -o contracts/token/artifacts
	abigen --bin=contracts/token/artifacts/MyNFTToken.bin --abi=contracts/token/artifacts/MyNFTToken.abi --pkg=contracts --out=contracts/token/MyNFTToken.go

airdrop:
	solc --abi --bin contracts/airdrop/MerkleAirdrop.sol -o contracts/airdrop/artifacts
	abigen --bin=contracts/airdrop/artifacts/MerkleAirdrop.bin --abi=contracts/airdrop/artifacts/MerkleAirdrop.abi --pkg=contracts --out=contracts/airdrop/MerkleAirdrop.go
