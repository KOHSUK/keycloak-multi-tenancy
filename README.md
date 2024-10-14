# Keycloak Multi Tenancy

## Keycloak を設定する

以下を参考に設定を行います。

[Keycloak の設定方法](./keycloak/setup.md)

## Web

```
npm i
npm run dev
```

## API

```
# DB migration
task dbmigration

# Build
task build

# Start
./idapi

# health check
curl http://localhost:7070/health

```
