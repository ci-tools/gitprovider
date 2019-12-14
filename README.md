# GitProvider

GitProvider is a CLI tool that allows to manage repositories/users and other resources via API. Currently supported following git servers:
* [GitLab](pkg/providers/gitlab/README.md)

## Usage

```sh
export SERVER=http://git.example.com
export TOKEN=<secret>
gitprovider repo ls
```