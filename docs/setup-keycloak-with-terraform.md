# Setup Keycloak with Terraform

## Prerequisites

- Terraform CLI

## Open Keycloak Admin Console

![alt text](image.png)

Open http://localhost:8080 and login with the following credentials:

- Username: admin
- Password: admin

## Create a client for the Terraform Provider

- Click `Create Client`

![alt text](image-1.png)

- Fill in the form with the following information:
  - Client type: openid-connect
  - Client ID: terraform
  - Name: Terraform

![alt text](image-2.png)

- Client authentication: On
- Service accounts roles: On
- Other settings: off

![alt text](image-3.png)

- Leave both `Root URL` and `Home URL` blank and click `Save`

![alt text](image-4.png)

- Open `Service Account` tab and click `Assign Role`

![alt text](image-5.png)

- Select all roles and click `Assign`

![alt text](image-6.png)

- Select `Filter by realm roles` and select `admin`