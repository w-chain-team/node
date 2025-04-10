package abis

const StakingJSONABI = `[
  {
    "inputs": [],
    "name": "validators",
    "outputs": [
      {
        "internalType": "address[]",
        "name": "",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "validatorBLSPublicKeys",
    "outputs": [
      {
        "internalType": "bytes[]",
        "name": "",
        "type": "bytes[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getRewardPool",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getRewardPerEpoch",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]`

const StressTestJSONABI = `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "number",
          "type": "uint256"
        }
      ],
      "name": "txnDone",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "getCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "sName",
          "type": "string"
        }
      ],
      "name": "setName",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
