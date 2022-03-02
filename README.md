# CertLS
Gather information about remote SSL certificates. 

## Output
Information about the remote SSL certificates prints to terminal.

Outputs:

* text
* csv
* json
* yaml

## Input file
All server/hosts are stored in a JSON or YAML file (using flag --host-file=<path>)

## Build
To build certls

```
go build -o certls cmd/certls/main.go
```

### Build releases
You can use the bash script:
```
build.sh
```
This will build and include git versions in binarys.

It will build for:

* linux amd64/arm64
* windows amd64/arm64
* darwin (macos) amd64/arm64

WARNING
This will remove all untracked files from your git repository.

# License
```
The MIT License (MIT)

Copyright (c) 2021 Micke Sommar <me@mickesommar.com>

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