variable "terratowns_access_token" {
  type = string
}
variable "terratowns_endpoint" {
  type = string
}
variable "teacherseat_user_uuid" {
  type = string
}
variable "terrahomes" {
  type = map(object({
    name            = string
    description     = string
    town            = string
    content_version = number
  }))
}

variable "public_path" {
  type = string
}
