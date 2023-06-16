# gupTrade

This is a highly extensible and flexible investment trading bot inspired by my brother, Guppy. It is intended for use with algorithms that execute trades programatically, most likely
on a particular interval of time, with the hope to support several different brokerage platforms and trading strategies.

This project is ultimately a single executable that is centered around a `configuration.yaml` file, an example of which can be seen
in `infra/local/configuration.yaml`. 

This currently is very much a work in progress and the design is very likely to change drastically. I just want to save my progress in a repo so here it is.

## Names & Definitions:
Broker: This refers to the stock broker the trades are being executed on, e.g. Schwab, Fidelity, Webull, Robin Hood, etc.
        The brokers supported in this project will be those which have good API support, like Tradier, or Alpaca and some others

Strategy: A Strategy is a concept defined in this project, which contains the actual logic for executing a desired algorithm. This is the 
          portion that will do any market evaluations, choose what trades to execute, or what to buy/sell, etc.

Trader: A Trader is a concept defined in this project, it contains and runs several strategies to execute, and has the client for a broker attached to it.

So there are several Traders, each of which is for one brokerage platform, and each trader can contain several strategies to execute on that particular platform
This hierarchy can be seen in the example `configuration.yaml` file.
Note: you can create several Traders with the same broker. So you can run multiple users with the one binary(though idk why you would, probably a better idea to keep users in separate deployments)