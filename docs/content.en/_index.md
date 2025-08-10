---
title: Introduction
bookHeadingAnchor: false
---

<div class="book-hero">

<div class="book-hero">

# Gitctx Documentation

{{<button relref="/docs/intro">}}Get Started with gitctx{{</button>}}

</div>

{{% columns %}}

- ## What is gitctx?

  **`gitctx`** is a lightweight CLI tool written in Go for managing multiple Git contexts (identities) effortlessly.  
  If you work with multiple Git profiles—such as **personal** and **work**—you can switch between them instantly without editing your `~/.gitconfig` manually.

  - Store unlimited named contexts in `~/.gitctx/`
  - Apply a context globally or locally per repository
  - Manage SSH/GPG keys per context
  - Quickly toggle between your **current** and **previous** identity

- ## Why use gitctx?

  Switching Git identities is a common pain point for developers working across different organizations, OSS projects, or personal work.  
  `gitctx` centralizes your configurations, reduces mistakes (like committing with the wrong email), and lets you manage them through an intuitive CLI or simple interactive prompts.

{{% /columns %}}

{{% columns %}}

- {{< card title="Create Contexts" image="" >}}

  # Add and Configure

  Create a new context with `gitctx add <name>` or import your existing `~/.gitconfig`.  
  Prefill values with flags, or use an interactive prompt to set up `user.name`, `user.email`, SSH keys, and more.
  {{< /card >}}

- {{< card title="Switch in Seconds" image="" >}}

  # Quick Toggle

  Swap between your **current** and **previous** contexts with a single command:  
  `gitctx switch` — perfect for bouncing between work and personal projects.
  {{< /card >}}

- {{< card title="Apply Anywhere" image="" >}}

  # Global or Local

  Apply a context globally (`~/.gitconfig`) or only to the nearest repo (`.git/config`) without touching other projects.
  {{< /card >}}

{{% /columns %}}

{{% columns %}}

- {{< card title="Organized Storage" >}}

  ### Everything in One Place

  All contexts are stored in `~/.gitctx/` as plain `.gitconfig` files, alongside metadata tracking your current and last-used context.
  {{< /card >}}

- {{< card title="Safe and Secure" >}}

  ### Optional Key Encryption

  Future features include encrypted SSH key storage using tools like **Age** or **SOPS** to keep your identities safe across devices.
  {{< /card >}}

- {{< card title="Future Integrations" >}}

  ### Smarter Git

  Planned integrations include **LazyGit** plugins, smart context detection based on remotes, and syncing across devices.
  {{< /card >}}

{{% /columns %}}
