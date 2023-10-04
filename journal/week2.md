# Terraform Beginner Bootcamp 2023 - week 2 journal

## Working with Ruby

### Bundler

A package manager for Ruby, used to install packages called gems.

https://bundler.io/

#### Installing gems

Create a gemfile and define gems then run `bundle install` to install the gems globally. (Unlike Node which installs locally in `node_modules`)

#### Executing ruby scripts

Use `bundle exec` to tell the ruby scripts to use the gems.

## Mock server

### Sinatra

Micro framework for ruby to build web apps.
https://sinatrarb.com/

### Running our server

```sh
cd ./terratowns_mock_server
bundle install
bundle exec ruby server.rb
```
