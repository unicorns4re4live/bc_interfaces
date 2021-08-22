
const address = "0xDb7168F98C5323714a0BE7f473a70043EFC25D7A";

const App = {
  data() {
    return {
      web3: undefined,
      currentAccount: undefined,
      contractInstance: undefined,
      results: {
        num: undefined,
      },
      sendData: {
        num: 0,
        address: '0x',
        pwhash: '0x'
      },
    };
  },
  methods: {
    bootstrapWeb3() {
      if (typeof window.ethereum !== "undefined") {
        window.ethereum.enable().then(() => {
          this.web3 = new Web3(window.ethereum);
          
          this.contractInstance = new this.web3.eth.Contract(abi, address);
        });
      } else {
        console.log("Установите метамаску");

        Web3.providers.HttpProvider.prototype.sendAsync =
          Web3.providers.HttpProvider.prototype.send;
        this.web3 = new Web3(
          new Web3.providers.HttpProvider("http://localhost:7545")
        );
        this.contractInstance = new this.web3.eth.Contract(abi, address);
      }
    },
    async refreshAccount() {
      const accs = await this.web3.eth.getAccounts();
      if (accs.length === 0) {
        console.log("Аккаунтов нет");
        return;
      }
      this.currentAccount = accs[0];
    },
    async getFromStorage() {
      const result = await this.contractInstance.methods.get_userlist().call();
      this.results.num = result;
    },
    async createUser() {
      const result = await this.contractInstance.methods.create_user(this.sendData.pwhash).send({ from: this.currentAccount });
      this.results.num = result;
      console.log('AAAAAAAAAAAAA', result)
    },

  },
  created() {
    this.bootstrapWeb3();
    setInterval(() => {
      this.refreshAccount();
    }, 500);
  },
};

Vue.createApp(App).mount("#app");
