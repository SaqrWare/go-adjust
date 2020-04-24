## HTTP Adjust task

This tool is written in golang, it achieves the task goals which are

- ✔️ Making http requests and prints the address of the request along with the MD5 hash of the response.
- ✔️ Ability to perform the requests in parallel so that the tool can complete sooner.
- ✔️ Ability to limit the number of parallel requests.
- ✔️ The tool must accept a flag to indicate this limit, and it should default to 10 if the flag is not provided.
- ✔️ Unit tests
    
#### Compile:
Make sure go is installed and run the following command

```bash
make compile
```
This compiles the code and saves the binary file as ./bin/myhttp
___
#### Running tests:
Run the following command

```bash
make test
```
___

#### Usage Examples:
After compiling the code you can use the tool like the following examples  
`-parallel` flag represents the maximum number of the parallel requests 
```bash
./bin/myhttp  adjust.com saqrware.com   
./bin/myhttp -parallel 3  adjust.com saqrware.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com   
```

___ 

#### Development

First you need to install entr
```bash
sudo apt install entr -y
```
Then run 
```bash
make watch
```
This keeps watching go files and recompile after any changes
