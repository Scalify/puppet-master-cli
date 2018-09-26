# puppet-master-cli

> A simple and smart CLI for the puppet-master api.

## Usage

You can use the CLI either by installing the go package through running `go install github.com/Scalify/puppet-master-cli`
or by running the docker image:

```bash
docker run --rm -it scalify/puppet-master-cli --help
```

which prints something like the following:
```
$ docker run --rm -it scalify/puppet-master-cli --help
Unable to find image 'scalify/puppet-master-cli:latest' locally
latest: Pulling from scalify/puppet-master-cli
4fe2ade4980c: Pull complete
b2340255bddd: Pull complete
4fb5a75d3b92: Pull complete
Digest: sha256:537f603d389f4a8b9a870c7fa08b06b4fa6c69837a8cac574e08982d6ada404b
Status: Downloaded newer image for scalify/puppet-master-cli:latest
A simple and smart CLI for the puppet-master api.

Usage:
  puppet-master-cli [command]

Available Commands:
  exec        Execute a single job
  help        Help about any command

Flags:
  -h, --help      help for puppet-master-cli
      --verbose   Verbose mode enables detailed logs

Use "puppet-master-cli [command] --help" for more information about a command.
```

## Available commands

### exec
Right now only the `exec` command is available, which takes a directory of code, modules and vars and executes it
against the puppet-master, specified through environment variables (endpoint, apiToken). Example:

**Assuming you are running from the root of this directory, using the `test` directory, and are running the self hosted example from [here](https://github.com/Scalify/puppet-master/tree/master/examples/self_hosted)!** 
```bash
docker run \
    -v "$(pwd)/test:/test" \
    --network puppet_master \
    -e "PUPPET_MASTER_ENDPOINT=http://gateway" \
    -e "PUPPET_MASTER_API_TOKEN=puppet" \
    --rm -it scalify/puppet-master-cli exec /test
```

should print something like this:

```
$ docker run \
>     -v "$(pwd)/test:/test" \
>     --network puppet_master \
>     -e "PUPPET_MASTER_ENDPOINT=http://gateway" \
>     -e "PUPPET_MASTER_API_TOKEN=puppet" \
>     --rm -it scalify/puppet-master-cli exec /test
INFO[0000] Executing jobs from directory /test           cmd=exec
INFO[0000] Loaded code file code.mjs, vars file vars.json, module files [modules/shared.mjs]  cmd=exec
INFO[0000] Created job with id 4830a025-4bbb-4072-b74f-eb9d61d867ea  cmd=exec
INFO[0000] Job has status created                        cmd=exec
INFO[0001] Job has status done                           cmd=exec


Job:
UUID                |4830a025-4bbb-4072-b74f-eb9d61d867ea |
Status              |done                                 |
Duration            |391                                  |
Error               |                                     |
Created at          |2018-09-26 09:39:46 +0000 UTC        |
Started at          |2018-09-26 09:39:46 +0000 UTC        |
Finished at         |2018-09-26 09:39:47 +0000 UTC        |


Logs:
2018-09-26 09:39:46 +0000 UTC |DEBUG               |Setting page viewport to width 1920 / height 1080                       |
2018-09-26 09:39:46 +0000 UTC |DEBUG               |Setting default language to en                                          |
2018-09-26 09:39:46 +0000 UTC |DEBUG               |Linking module shared to file:///puppet-master/code (Available: shared) |
2018-09-26 09:39:47 +0000 UTC |INFO                | 185.232.23.186                                                         |
2018-09-26 09:39:47 +0000 UTC |INFO                |Code took 391ms to execute.                                             |


Results:
ip                  | 185.232.23.186     |


INFO[0001] Done.                                         cmd=exec
```

## License

Copyright 2018 Scalify GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
