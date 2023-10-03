# Terraform Beginner Bootcamp 2023 - week 1 journal

## Root module structure

```
PROJECT_ROOT
│
├── main.tf                 # everything else.
├── variables.tf            # stores the structure of input variables
├── terraform.tfvars        # the data of variables we want to load into our terraform project
├── providers.tf            # defined required providers and their configuration
├── outputs.tf              # stores our outputs
└── README.md               # required for root modules
```

[Standard module structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform Cloud variables

Two types can be set `environment variables` and `terraform variables`. They can be set to sensitive, which makes them read only once set.

_Note_:If the run is set to `local` for the project, the `Variables` menu is not available in the settings panel

## Dealing with configuration drift

If resources are altered manually, ie. via ClickOps, TF will try to restore it based on the state file when ran it.

### Fixing missing resources with `Terraform Import`

If we want to pull ClickOps resources into TF, we can try [Terrafrom Import](https://developer.hashicorp.com/terraform/cli/import)

## Terraform modules

Modules are the main way to package and reuse resource configurations with Terraform.
[More info](https://developer.hashicorp.com/terraform/language/modules)

## Working with files in Terraform

There are a lot of functions available on [https://developer.hashicorp.com/terraform/language/functions/](https://developer.hashicorp.com/terraform/language/functions/)

### path

There is a special variable in Terraform called `path`, that allows us to reference local paths.

[Terraform `path` reference](https://developer.hashicorp.com/terraform/language/expressions/references#filesystem-and-workspace-info)

### File changes

Terraform does not monitor file content changes. The way to have it check for changes in files is using an [`ETag`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) and its value a file's `md5` hash using the Terraform function [`filemd5`](https://developer.hashicorp.com/terraform/language/functions/filemd5)

### Consideration

Managing files with Terraform is not necessarily considered a best practice - use common sense.

## Terraform locals sources

Terraform local values (or "locals") assign a name to an expression or value. Using locals simplifies your Terraform configuration – since you can reference the local multiple times, you reduce duplication in your code. Locals can also help you write more readable configuration by using meaningful names rather than hard-coding values.

[Reference](https://developer.hashicorp.com/terraform/tutorials/configuration-language/locals)

## Terraform data sources

Data sources allow Terraform to use information defined outside of Terraform, defined by another separate Terraform configuration, or modified by functions.
[Reference](https://developer.hashicorp.com/terraform/language/data-sources)

## Terraform Lifecycle

[Reference](https://developer.hashicorp.com/terraform/tutorials/state/resource-lifecycle)

## terrafrom_data resource

[Reference](https://developer.hashicorp.com/terraform/language/resources/terraform-data)

## Provisioners

Allow you to execute commands on compute instances e.g. AWS CLI command.
NOT RECOMMENDED BY HASHICORP as config management tools like Ansible are better suited.

Tho functionality exists.

- `local-exec`: executes on the machine running the `terraform` command
- `remote-exec`: executes on the specified machine. (Requires auth)

[Reference](https://developer.hashicorp.com/terraform/language/resources/provisioners/syntax)
