# GH Runner Regger

A simple program to turn GitHub App IDs and private keys into a runner registration token. This is useful for setting up self-hosted runners in a CI/CD pipeline.

To set it up, create a new github app and generate a private key for it. Make sure to install the app in the organization you want to use it in.
The app needs read/write access to github runners. Then, use this tool to generate a registration token for the app.
You can then use the token to register a runner.

## Usage

````bash
```bash
$ ./gh-runner-regger ref+gcpsecrets://secrets-project/private-key app-id myOrg
````

The first argument can either point to a local `pem` file or any valid `vals` URL, see below for more info on `vals`.
It should represent a private key for a GitHub App. app-id is the numeric ID of the GitHub App. `myOrg` is the name of the organization where
the app is installed.

Note that this approach can only be used to register org level runners, not repo level runners.

## See also

- [vals](https://GitHub.com/helmfile/vals) - A tool for injecting values into YAML files from a variety of sources

## MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
