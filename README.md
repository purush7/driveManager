
### Description:

  Simple application which helps in viewing and downloading files of a google drive folder

### Prerequisites:

- Go >= v1.16.7

### Building the binary:

- Copy the folder id,name and enter the preferred browser in `client/env.go`
- Run `go build`

### Commands:

```bash
./driveManager list 
./driveManager download

```

### How it works:

First time while you run the binary, it opens the consent page on browser and asks the consent and permissions. After getting permissions, it creates
`token.json` and it uses the refresh token for accessing the google drive folder.
> If the scope has been change, delete the token.json
