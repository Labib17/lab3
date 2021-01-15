  
const { Client } = require("./accounts/client");

const client = Client("http://localhost:15000");

(async () => {
   let account;
    console.log("Scenariy 1");
    try {
        account = await client.fetchAccounts();
        console.log("Balance of each account");
        console.table(account);
    } catch (err) {
        console.log(`Problem listing account: `, err);
    }
})();

(async () => {
    
    console.log("Scenariy 2");
    try {
        const transfer = await client.transferMoney(
            1000,
            2,
            3
        );
    } catch (err) {
        console.log(`Problem with transfer: `, err);
    }
})();
