token:
	solc --abi --bin contracts/token/MyNFTToken.sol -o contracts/token/artifacts
	abigen --bin=contracts/token/artifacts/MyNFTToken.bin --abi=contracts/token/artifacts/MyNFTToken.abi --pkg=contracts --out=contracts/token/MyNFTToken.go

airdrop:
	solc --abi --bin contracts/airdrop/MerkleNFTAirdrop.sol -o contracts/airdrop/artifacts
	abigen --bin=contracts/airdrop/artifacts/MerkleNFTAirdrop.bin --abi=contracts/airdrop/artifacts/MerkleNFTAirdrop.abi --pkg=contracts --out=contracts/airdrop/MerkleNFTAirdrop.go
