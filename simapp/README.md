# Configure environment

Install [Go 1.14+](https://golang.org/dl/)

`export PATH="$HOME/go/bin:$PATH"`

# Build

Execute in the Cosmos root dir
```
make build-simd
./simapp/init.sh
```

# Start the Cosmos node

nsd start --with-tendermint=false --transport=grpc

# Cosmos CLI commands

## First check the accounts to ensure they have funds
nscli query account $(nscli keys show jack -a)
nscli query account $(nscli keys show alice -a)

## Buy your first name using your coins from the genesis file
nscli tx nameservice buy-name jack.id 5nametoken --from jack --chain-id namechain

## Set the value for the name you just bought
nscli tx nameservice set-name jack.id 8.8.8.8 --from jack --chain-id namechain

## Try out a resolve query against the name you registered
nscli query nameservice resolve jack.id

## Try out a whois query against the name you just registered
nscli query nameservice whois jack.id

## Alice buys name from jack
nscli tx nameservice buy-name jack.id 10nametoken --from alice --chain-id namechain

## Alice decides to delete the name she just bought from jack
nscli tx nameservice delete-name jack.id --from alice --chain-id namechain
