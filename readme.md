# Currency converter

Go application to convert currencies by hourly updated prizes on console or by API

## Run

.env:
```
#Set up your Fixer Api key
API_KEY=""

#console / api
RUN_TYPE=""
```
#### dev
```bash
# help:
go run . -h

# To run in console:
go run . -run=console
#or use env exept -run flag

# To run API
go run . -run=api
#or use env exept -run flag
```

#### prod

```bash
# build:
go build -o [target] . 

#set up env like above

#run binary
```


## Usage

### API

```url
http://localhost:8080/?amount=1000&input=eur&output=usd
```

### Console

Answer the questions
