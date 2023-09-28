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
