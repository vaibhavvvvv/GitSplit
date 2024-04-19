// module.exports = {
//     solidity: "0.8.24", 
//     networks: {
//       amoy: {
//         url: "https://80002.rpc.thirdweb.com",
//         chainId: 80002 
//       }
//     }
//   };    


require('dotenv').config();
require("@nomiclabs/hardhat-ethers");

const { API_URL, PRIVATE_KEY } = process.env;

module.exports = {
   solidity: "0.8.24",
   defaultNetwork: "amoy",
   networks: {
      hardhat: {},
      amoy: {
         networkId: 80002,
         url: API_URL,
         accounts: [`0x${PRIVATE_KEY}`]
      }
   },
}