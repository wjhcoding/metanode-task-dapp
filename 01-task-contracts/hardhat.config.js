require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",

  networks:{
    localhost: {},
    sepolia:{
      url: "https://sepolia.infura.io/v3/" + process.env.INFURA_API_KEY,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY]
    },
    // mainnet:{
    //   url: "https://mainnet.infura.io/v3/" + process.env.INFURA_API_KEY,
    //   accounts: [process.env.DEPLOYER_PRIVATE_KEY]
    // }
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY
  }  
};
