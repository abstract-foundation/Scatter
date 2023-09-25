# Ozone (A database built for blockchain)
Ozone would be the database that powers the Rivative Network.
Ozone would have it's own objects (`ledger`, `block`, `blockchain`, `etc..`).
These objects would be how data is stored from the network. Ozone would also allow
for the creation of custom data types on the fly.

```createType.oz

sc = object `smartcontract` {
    creator `account`
    domain `domain`
}

oz.LoadEnvironment("locale")

oz.Publish(sc)


```