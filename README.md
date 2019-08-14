# Skyblog

- SQL: PostgresQL
- ORM: go-pg

# Getting started

```bash
go build -o skyblog main.go
./skyblog init
```

# Running tests

```bash
go test
```

# How to use this CLI

To see all the flags:

```bash
./skyblog help
```

Global structure of a command :

```bash
./skyblog <create|read|update|delete|list> <article|user> -flags
```

## Users

**Create** (Jean, Jacques, jj@gmail.com)

```bash
./skyblog create user --firstName Jean --lastName Jacques --email jj@gmail.com --password Azerty1234 --phone +33123456789
```

**Update** (#ID, Jean, Jacques, jj@gmail.com)

```bash
./skyblog update user --id eb8e7457-fb2f-4d68-9424-8240b8a19b04 --email mathias.lachot2@gmail.com
```

**Read** (#ID)

```bash
./skyblog read user --id eb8e7457-fb2f-4d68-9424-8240b8a19b04
```

**Delete** (#ID)

```bash
./skyblog delete user --id eb8e7457-fb2f-4d68-9424-8240b8a19b04
```

**List**

```bash
./skyblog list user
```

## Articles

**Create** (Dummy Name, Dummy Content, Created by Jean Jacques)

```bash
./skyblog create article --name "Dummy Name" --content "Dummy Content" --userEmail jj@gmail.com
```

**Update** (#ID, Dummy Name, Dummy Content, Created by Jean Jacques)

```bash
./skyblog update article --id u-u-i-d --name "Dummy Name" --content "Dummy Content" --userEmail jj@gmail.com
```

**Read**

```bash
./skyblog read article --id u-u-i-d
```

**Delete** (#ID)

```bash
./skyblog delete article --id u-u-i-d
```

**List**

```bash
./skyblog list article
```

# Consigne de base

Choisissez un stockage (SQL, NoSQL) et une API (sql, ORM, ODM…)
Ecrivez un programme Go qui prenne des options de command pour les étapes ci-dessous.

## v0:

Un programme unique capable de réaliser des commandes

## v1:

Modélisez deux entités en user/article (1:n)
Insérez des documents reliés, listez la base en hiérarchique
Modifiez un document maître, démontrez la cascade
Supprimez un document maître, démontrez la cascade

## v2:

Modifiez la relation entre les entités en m:n
Appliquez une migration des données existantes
Mêmes démonstrations: lister, modifier, supprimer
