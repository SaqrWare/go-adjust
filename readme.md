## My http Adjust task

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