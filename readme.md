## HTTP Adjust task

This tool is written in golang, it achieves the task goals which are

- ✔️ Making http requests and prints the address of the request along with the MD5 hash of the response
- ✔️ Ability to perform the requests in parallel so that the tool can complete sooner. T
- ✔️ Ability to perform the requests in parallel so that the tool can complete sooner. T
    
#### Compile:

```bash
make compile
```

___
#### Run Examples:

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
This keeps watching go files and recompile after any update

___
#### TODO:

- Write complete Reame file
- Write tests