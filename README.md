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

## Commentaires :

Question de l'un d'entre vous: "par rapport à l'étape v2 : "Modifiez la relation entre les entités en m:n". Une relation m:n correspond-elle à une relation de type :
Table 1 (id, ...)
Table 2 (id, ...)
Table 3 (id, table1_id, table2_id, ...)
?
Mais la relation 1:n créée à l'étape v1 doit toujours exister ? C'est à dire qu'il doit toujours y avoir quelque chose du type :
Table1_maitre (id, ...)
Table1_detail (id, talbe1_maitre_id, ...)"

Réponse: non, la relation m:n remplace la relation 1:n. Le but de cette partie du devoir est de traiter un changement de schéma avec conservation de toutes les données possibles, soit manuellement, soit avec un outil de migration tiers, soit avec un outil faisant partie de l'ORM que vous utilisez. Dans le cas d'une migration 1:n vers m:n, il n'y a pas de perte d'information, les éléments de la table détail se retrouvant simplement reliée par la table de relation, avec une seule valeur dans la relation m:n

Information en plus

Question de l'un(e) d'entre vous, qui ne parvient pas à réaliser l'intégralité des étapes du devoir dans l'ordre prévu: est-il possible de ne rendre que des parties du devoir non reliées les unes aux autres ?

Réponse: oui. Pour avoir 20/20, il faut tout faire, ou presque tout et avoir des bonus (présentation, commentaires, tests...); et c'est globalement plus facile en faisant toutes les étapes dans l'ordre. Mais toutes les questions sont notées séparément. Donc si vous séchez sur certaines questions, mais parvenez à faire certaines autres plus loin, n'hésitez pas. En revanche, prenez soin de mettre un README.md à la racine de votre projet pour expliquer ce que vous avez fait et comment le vérifier pour la notation.
