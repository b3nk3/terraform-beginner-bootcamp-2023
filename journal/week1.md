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
