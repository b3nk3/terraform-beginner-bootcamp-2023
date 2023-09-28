# Terraform Beginner Bootcamp 2023

## Table of contents

- [Terraform Beginner Bootcamp 2023](#terraform-beginner-bootcamp-2023)
  - [Table of contents](#table-of-contents)
  - [Commit messages](#commit-messages)
  - [Versioning](#versioning)
  - [Install Terraform CLI](#install-terraform-cli)
    - [Considerations for Linux Distros - Changes to the CLI installation](#considerations-for-linux-distros---changes-to-the-cli-installation)
    - [Refactoring into bash script](#refactoring-into-bash-script)
      - [Linux permissions](#linux-permissions)
      - [Shebang](#shebang)
    - [Gitpod lifecycle](#gitpod-lifecycle)
    - [Working with env vars](#working-with-env-vars)
      - [env command](#env-command)
      - [set/unset env vars](#setunset-env-vars)
      - [Scoping env vars](#scoping-env-vars)
      - [Persisting env vars in Gitpod](#persisting-env-vars-in-gitpod)
    - [AWS CLI installation](#aws-cli-installation)
      - [Check if logged in to AWS](#check-if-logged-in-to-aws)
  - [Terraform Basics](#terraform-basics)
    - [Terraform Registry](#terraform-registry)
    - [Terraform CLI](#terraform-cli)
    - [Example provider](#example-provider)
    - [Output details from executed terraform](#output-details-from-executed-terraform)
    - [Terraform Cloud issues](#terraform-cloud-issues)

## Commit messages

For clean commit messages we're using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## Versioning

This project will make use of [Semantic Versioning](semver.org) for tagging.

Given a version number **MAJOR.MINOR.PATCH**, increment the:

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.

## Install Terraform CLI

[https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli](Official installation guide)

### Considerations for Linux Distros - Changes to the CLI installation

_Note_: This project uses Ubuntu. Do check relevant docs for your OS and make changes accordingly

The original steps used in `gitpod.yml` have been deprecated therefore an update was necessary.

### Refactoring into bash script

Script location: [./bin/install_terraform_cli.sh](./bin/install_terraform_cli.sh)

New guide is considerably longer hence branching it out into a standalone bash script.

This allows for:

- tidier, more readable code in Gitpod Task File ([`.gitpod.yaml`](./.gitpod.yaml))
- make debugging easier
- better portability of install script into other projects that use Terraform CLI

#### Linux permissions

To make our script executable, we need to run: `chmod +x ./bin/install_terraform_cli.sh`

#### Shebang

Tell the OS how to run our program.
[More info](<https://en.wikipedia.org/wiki/Shebang_(Unix)>)

### Gitpod lifecycle

To allow for our tools (TF/AWS CLIs) to install on workspace restart, we call the install from a `before` section

[Execution order](https://www.gitpod.io/docs/configure/workspaces/tasks#execution-order)

### Working with env vars

#### env command

We can list out all env vars using `env` command
We can filter env vars by piping to grep eg. `env | grep AWS_`

#### set/unset env vars

In terminal we can set with

```sh
export HELLO=World
```

Unset with

```sh
unset HELLO
```

Temporarily for the script we are running on command line

```sh
HELLO='world' ./bin/my_script.sh
```

Within `.sh` file without export keyword

```sh
#!/usr/bin/env bash

HELLO=World

echo $HELLO
```

#### Scoping env vars

Env vars are scoped to individual terminal windows. If you want to use them across windows, we need to set them globally.

#### Persisting env vars in Gitpod

Env vars can be persisted in Gitpod Secrets Storage.

```sh
gp env HELLO='world'
```

All future workspaces will set the env vars for all terminals opened.

non sensitive env vars can also be set in `.gitpod.yml`

### AWS CLI installation

AWS CLI is installed in this project vie bash script [./bin/install_aws_cli.sh](./bin/install_aws_cli.sh.)

[Official Docs](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
[AWS CLI Env vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

#### Check if logged in to AWS

```bash
aws sts get-caller-identity
```

If successful it returns a JSON object

```json
{
  "UserId": "AKIAIOSFODNN7EXAMPLE",
  "Account": "123456789012",
  "Arn": "arn:aws:iam::123456789012:user/benszabo"
}
```

AWS Access Key need to be generated for the appropriate IAM user on the AWS Console.

## Terraform Basics

Infrastructure as Code -

### [Terraform Registry](https://registry.terraform.io/)

Terraform stores providers and modules (plus extra bits) in its registry.

- _Providers_ are interfaces for the likes of AWS, Azure, GCP etc. and give use a way to create resources
- _Modules_ are a way to make large TF files modular and reusable

### Terraform CLI

_NOTE_: We've aliased the `terraform` command to `tf` in our project.

To acquire the our provider binaries we should always run `terraform init`
To plan we run `terraform plan`
To apply we run `terraform apply` - then confirm/deny as appropriate or use the `--auto-approve` flag
To destroy our infra we run `terraform destroy` - it will tear everything down - `--auto-approve` works here too

- `.terraform` contains binaries
- `.terraform.lock.hcl` contains provider versions - **should be commited**
- `terraform.tfstate` contains current state of our infra - _**should not be commited**_ as it can contains sensitive information
- `terraform.backup.tfstate` previous version of state - _**don't commit**_

[more details on CLI](https://developer.hashicorp.com/terraform/cli/commands)

### Example provider

[Terraform Random Provider](https://registry.terraform.io/providers/hashicorp/random/latest/docs)

### Output details from executed terraform

We can write out information created by terraform
[Terraform Outputs Documentation](https://developer.hashicorp.com/terraform/language/values/outputs)

### Terraform Cloud issues

- TF Cloud defaults to remote run, which fails due to lacking AWS auth something or other - short term solution is to switch the workspace to `local` execution
- Automated TF Cloud authentication file generation
