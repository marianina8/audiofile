# audiofile
In Chapter 3, we discuss a command line interface built from scratch which handles generating metadata from uploaded audio files, local flat file storage and retrieval of audio metadata.  This CLI is just an example and for reference to the chapter.  It was created on MacOS and other operating systems have not been tested.

Only the upload and get commands are currently implemented.  To expand upon this CLI for your own education, I suggest forking the repo and:
* implementing the list command.
* modifying the code to run on other operating systems
* adding new commands
* implementing new storage type other than flat local file storage

## Within the root of the audiofile folder, to start the API:
go run cmd/api/main.go
### NOTE
To change the default port, 80, pass in the new port value with the `-p` flag.

## To generate the audiofile command line interface:
go build -o audiofile-cli cmd/cli/main.go

## To call the audiofile command line interface:
./audiofile-cli

### NOTE
The API must be started and running before the CLI.  First, build the API:
go build -o audiofile-api cmd/api/main.go
Then run it in a separate terminal:
./audiofile-api
