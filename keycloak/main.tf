terraform {
  required_providers {
    keycloak = {
      source  = "mrparkers/keycloak"
      version = "~> 4.4.0"
    }
  }
}

provider "keycloak" {}

resource "keycloak_realm" "realm" {
  realm        = var.realm
  enabled      = var.enabled
  display_name = var.display_name

  registration_allowed           = true
  registration_email_as_username = true
  reset_password_allowed         = true
  login_with_email_allowed       = true
  duplicate_emails_allowed       = false
}

resource "keycloak_openid_client" "client" {
  realm_id = keycloak_realm.realm.id

  client_id                       = "app"
  name                            = "App"
  access_type                     = "CONFIDENTIAL"
  standard_flow_enabled           = true
  implicit_flow_enabled           = false
  direct_access_grants_enabled    = true
  service_accounts_enabled        = false
  valid_redirect_uris             = var.valid_redirect_uris
  valid_post_logout_redirect_uris = var.valid_post_logout_redirect_uris
}
