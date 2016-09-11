# leisp.go
The leisp programming language written in Go


## Installation

Run the following command to install leisp:

```bash
$ go install github.com/leizongmin/leisp
```

After finished install leisp, start the leisp REPL:

```
$ leisp


##       ######## ####  ######  ########
##       ##        ##  ##    ## ##     ##
##       ##        ##  ##       ##     ##
##       ######    ##   ######  ########
##       ##        ##        ## ##
##       ##        ##  ##    ## ##
######## ######## ####  ######  ##

Welcome to leisp 0.0.1

Copyright (c) 2016 Zongmin Lei <http://ucdok.com>

Type (help) and hit Enter for context help.
Press Ctrl+C to Exit.

leisp>
```


## Development

**leisp** use [gogo](https://github.com/leizongmin/gogo) to manage dependencies.
Please install **gogo** command firstly:

```bash
$ go get -u github.com/leizongmin/gogo
```

Then run the following commands to build:

```bash
# init
$ gogo clean && gogo init
# install dependencies
$ gogo install
# build leisp
$ gogo build
```


## License

```
The MIT License (MIT)

Copyright (c) 2016 Zongmin Lei <leizongmin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
