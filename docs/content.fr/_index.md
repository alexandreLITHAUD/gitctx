---
title: Introduction
bookHeadingAnchor: false
---

<div class="book-hero">

<div class="book-hero">

# Documentation Gitctx

{{<button relref="/docs/intro">}}Commencer avec gitctx{{</button>}}

</div>

{{% columns %}}

- ## Qu’est-ce que gitctx ?

  **`gitctx`** est un outil CLI léger écrit en Go pour gérer facilement plusieurs contextes (identités) Git.  
  Si vous utilisez plusieurs profils Git — comme **personnel** et **professionnel** — vous pouvez passer de l’un à l’autre instantanément, sans modifier manuellement votre `~/.gitconfig`.

  - Stockez un nombre illimité de contextes nommés dans `~/.gitctx/`
  - Appliquez un contexte globalement ou localement par dépôt
  - Gérez les clés SSH/GPG par contexte
  - Basculez rapidement entre votre **identité actuelle** et la **précédente**

- ## Pourquoi utiliser gitctx ?

  Changer d’identité Git est souvent une source de frustration pour les développeurs travaillant sur plusieurs organisations, projets open source ou travaux personnels.  
  `gitctx` centralise vos configurations, réduit les erreurs (comme un commit avec la mauvaise adresse e-mail) et vous permet de les gérer via une CLI intuitive ou des invites interactives simples.

{{% /columns %}}

{{% columns %}}

- {{< card title="Créer des Contextes" image="" >}}

  # Ajouter et Configurer

  Créez un nouveau contexte avec `gitctx add <nom>` ou importez votre `~/.gitconfig` existant.  
  Préremplissez les valeurs avec des options (`flags`) ou utilisez une invite interactive pour définir `user.name`, `user.email`, les clés SSH, et plus encore.
  {{< /card >}}

- {{< card title="Changer en Quelques Secondes" image="" >}}

  # Bascule Rapide

  Alternez entre votre **contexte actuel** et le **précédent** avec une seule commande :  
  `gitctx switch` — parfait pour passer rapidement d’un projet professionnel à un projet personnel.
  {{< /card >}}

- {{< card title="Appliquer Partout" image="" >}}

  # Global ou Local

  Appliquez un contexte globalement (`~/.gitconfig`) ou seulement au dépôt le plus proche (`.git/config`) sans impacter les autres projets.
  {{< /card >}}

{{% /columns %}}

{{% columns %}}

- {{< card title="Stockage Organisé" >}}

  ### Tout au Même Endroit

  Tous les contextes sont stockés dans `~/.gitctx/` sous forme de fichiers `.gitconfig` simples, accompagnés de métadonnées pour suivre votre contexte actuel et le dernier utilisé.
  {{< /card >}}

- {{< card title="Sécurisé et Fiable" >}}

  ### Chiffrement Optionnel des Clés

  Les fonctionnalités futures incluront le stockage chiffré des clés SSH grâce à des outils comme **Age** ou **SOPS** pour sécuriser vos identités sur plusieurs appareils.
  {{< /card >}}

- {{< card title="Intégrations Futures" >}}

  ### Git Plus Intelligent

  Les intégrations prévues incluent des plugins **LazyGit**, la détection intelligente du contexte en fonction des dépôts distants, et la synchronisation entre appareils.
  {{< /card >}}

{{% /columns %}}
