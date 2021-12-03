# Sberbank Hackathon

goal: create a hyperledger chaincode contract where users can rent items they own to other users


completed: created a smart contract that allows users to rent their property to other users - see /contracts/rental.sol

completed: created Solidity to Go bindings for rental.sol


#### Rental Smart Contract live on Rinkeby:
https://rinkeby.etherscan.io/address/0xa22eecf77ba8d3d47b24612d2d1f87b4c61e3050#writeContract



## using geth abigen tool:

Create Solidity Go Bindings

```sh
abigen --sol rental.sol --pkg main --out ../goBindings/rental.go
```


## Why Sberbank should look into using a permissioned blockchain that allows developers to use Solidity 

#### Potential Risk Detection System of Hyperledger Fabric Smart Contract based on Static Analysis

1) https://easychair.org/publications/preprint_open/7c6R


*"Due  to  the  lack  of 
mature development specifications for smart contracts 
using  general-purpose  programming  language,  there  are 
often  potential  risks  in  the  smart  contracts  related  to  the 
characteristics  of  Hyperledger  Fabric.  It  will  bring  many 
inconveniences  and  potential  safety  hazards  to  users  after 
the smart contracts are deployed."*

#### Hyperledger Fabric: A Distributed Operating System for Permissioned Blockchains

2) https://arxiv.org/pdf/1801.10228.pdf

*"Unfortunately, generic languages pose many problems for ensuring deterministic execution. Even if the application developer
does not introduce obviously non-deterministic operations, hidden
implementation details can have the same devastating effect (e.g., a
map iterator is not deterministic in Go)."*



## Hyperledger Besu - EVM compatible permissioned blockchain

https://besu.hyperledger.org/en/stable/



--------
## @devs - useful links:

#### Deploying solidity to hyperledger fabric - couldn't figure it out :(

https://medium.com/coinmonks/solidity-smart-contract-on-hyperledger-fabric-3d50f25e577b


https://hyperledger-fabric.readthedocs.io/en/latest/deploy_chaincode.html#install-the-chaincode-package



Go binding
https://geth.ethereum.org/docs/dapp/native-bindings

https://hackernoon.com/a-step-by-step-guide-to-testing-and-deploying-ethereum-smart-contracts-in-go-9fc34b178d78git 


Install Hyperledger
https://medium.com/hackernoon/hyperledger-fabric-installation-guide-74065855eca9

EVM Chaincode
https://github.com/hyperledger/fabric-chaincode-evm

https://www.ledgerinsights.com/hyperledger-fabric-integrates-ethereum-smart-contracts-evm-blockchain/

https://www.youtube.com/watch?v=uAPZ2QMSJaI

https://github.com/ethereum/go-ethereum
