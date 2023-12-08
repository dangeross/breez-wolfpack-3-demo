# Breez SDK Demo for Wolfpack 3
```
% go run . -d data/test
__________                                                                 
\______   \_______   ____   ____ ________                                  
 |    |  _/\_  __ \_/ __ \_/ __ \\___   /                                  
 |    |   \ |  | \/\  ___/\  ___/ /    /                                   
 |______  / |__|    \___  >\___  >_____ \                                  
        \/              \/     \/      \/                                  
           __      __      .__   _____                     __     ________  
          /  \    /  \____ |  |_/ ____\__________    ____ |  | __ \_____  \ 
          \   \/\/   /  _ \|  |\   __\\____ \__  \ _/ ___\|  |/ /   _(__  < 
           \        (  <_> )  |_|  |  |  |_> > __ \\  \___|    <   /       \
            \__/\  / \____/|____/__|  |   __(____  /\___  >__|_ \ /______  /
                 \/                   |__|       \/     \/     \/        \/ 

sdk> 
```

## What's the plan

Add the Breez SDK to this barebones CLI to connect, receive and send payments. The `final` branch has the completed demo code.

*Add the SDK* - Add some code to access the SDK and listen to SDK events 

*Connect* - Register a node using a Greenlight Invite Code

*Query node state* - Get the node pubkey, on-chain / lightning liquidity

*Receive a payment* - Create an invoice, then receive a payment while opening a channel

*Send a payment* - Pay an invoice using our node