
const address = "0xDb7168F98C5323714a0BE7f473a70043EFC25D7A";
let contractInstance, web3, currentAccount;

window.onload = () => {
  bootstrapWeb3()
}

bootstrapWeb3 = () => {
  if (typeof window.ethereum !== "undefined") {
    window.ethereum.enable().then(() => {
      web3 = new Web3(window.ethereum);
      
      contractInstance = new web3.eth.Contract(abi, address);
    });
  } else {
    console.log("Установите метамаску");

    Web3.providers.HttpProvider.prototype.sendAsync =
      Web3.providers.HttpProvider.prototype.send;
    web3 = new Web3(
      new Web3.providers.HttpProvider("http://localhost:7545")
    );
    contractInstance = new web3.eth.Contract(abi, address);
  }
  setInterval(() => {
    refreshAccount();
  }, 500);
}

async function refreshAccount() {
  const accs = await web3.eth.getAccounts();
  if (accs.length === 0) {
    console.log("Аккаунтов нет");
    return;
  }
  currentAccount = accs[0];
  console.log('AAAAAAAAAAAAA', currentAccount)
};

async function getFromStorage() {
  const result = await contractInstance.methods.get_userlist().call();
  console.log('AAAAAAAAAAAAA', result)
  let resultStr = ``
  result.forEach((elem) => {
    resultStr+=`<li>${elem}</li>`
  })
  document.querySelector('.resultList').innerHTML = resultStr;
}


