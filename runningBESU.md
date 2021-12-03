## Deploying smart contracts on Besu

#### 1) Download Besu packaged binaries
https://github.com/hyperledger/besu/releases

#### 2) Build Besu

Build Besu

After cloning, go to the besu directory.

Build Besu with the Gradle wrapper gradlew, omitting tests as follows:


```sh
./gradlew build -x test
```

Go to the distribution directory: 

```sh
cd build/distributions/
```
Expand the distribution archive: 
```sh
tar -xzf besu-<version>.tar.gz
```
Move to the expanded folder and display the Besu help to confirm installation. 
```sh
cd besu-<version>/
bin/besu --help
```
add besu to path
```sh
sudo nano ~/.bashrc
```
reload ~/.bashrc
```sh
source ~/.bashrc
```

#### 3) Deploy smart contracts to Besu 
```sh
npm install @truffle/hdwallet-provider
```

```sh
npm run besu
```

```sh
truffle migrate --reset --network besu
```

