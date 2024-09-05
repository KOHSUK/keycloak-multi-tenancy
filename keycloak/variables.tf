variable "realm" {
  type        = string
  description = "value for realm"
}

variable "display_name" {
  type        = string
  description = "value for display_name"
}

variable "enabled" {
  type        = bool
  description = "value for enabled"
}

variable "valid_redirect_uris" {
  type        = list(string)
  description = "value for valid_redirect_uris"
}

variable "valid_post_logout_redirect_uris" {
  type        = list(string)
  description = "value for valid_post_logout_redirect_uris"
}