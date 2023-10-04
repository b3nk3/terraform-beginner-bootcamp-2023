variable "user_uuid" {
  description = "User UUID"
  type        = string

  validation {
    condition     = can(regex("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", var.user_uuid))
    error_message = "User UUID must be in the format of a UUID (e.g., 123e4567-e89b-12d3-a456-426655440000)"
  }
}

variable "assets_filepath" {
  description = "Filepath to the index.html file"
  type        = string
}
variable "index_html_filepath" {
  description = "Filepath to the index.html file"
  type        = string

  validation {
    condition     = fileexists(var.index_html_filepath)
    error_message = "The provided path for index.html does not exist."
  }
}

variable "error_html_filepath" {
  description = "Filepath to the error.html file"
  type        = string

  validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The provided path for error.html does not exist."
  }
}

variable "content_version" {
  description = "The content version"
  type        = number

  validation {
    condition     = var.content_version >= 1
    error_message = "Content version must be a positive number starting with 1"
  }
}
