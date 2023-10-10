# Terraform Beginner Bootcamp 2023

This project aims to cover the basics of infrastructure as code (IaC) using Terraform.

## What are we building?

A cloud infrastructure for website hosting and deployment into the 3rd party [TerraTowns website](https://terratowns.cloud/)'s topics.

### Architecture diagram

![Architecture diagram](image.png)

### Usage

_Assumption: using Gitpod_

1. Declare the website(s) details in the `terraform.tfvars`
   ```hcl
     BENS_RETRO_GAMES_TERRAHOME = {
        name = "Retro games from my childhood"
        description = "Played on a state of the art Intel Pentium I"
        town = "gamers-grotto"
        content_version = 1
     }
   ```
2. Create a folder in `public` using the name of the VAR from above (e.g. `BENS_RETRO_GAMES_TERRAHOME`) and place at least an index.html and error.html file it for deployment. Optionally add an assets folder with images(`jpg,jpeg,png,gif`)
3. run `tf init`
4. run `tf apply`
5. review then apply/deny the plan to deploy the infra

PS: run `tf destroy` to initiate the teardown of the infrastructure

## Weekly journals

- [Week 0 journal](journal/week0.md)
- [Week 1 journal](journal/week1.md)
- [Week 2 journal](journal/week2.md)
