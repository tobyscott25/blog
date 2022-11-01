---
date: 2022-04-18T18:48:05+11:00
title: "Signing your commits with GNU Privacy Guard (GPG)"
description: "Signing your commits with a GPG private key verifies it was really you who committed the changes."
tags: ["gpg", "git", "signed commits"]
---

## 1. Download and install GPG for your operating system
  - Windows: Download [GPG4Win](https://www.gpg4win.org/get-gpg4win.html) and install the following components: GnuPG, Kleopatra, GpgOL, GpgEX.
  - Mac (via [Homebrew](https://brew.sh/)):
```bash
$ brew install gpg
```
  - Linux (+ WSL): Install from your package manager
  
## 2. Generate or import your GPG key pair

### Import an existing GPG key pair from another computer

On the computer that already has the GPG key pair:

- Export the GPG key pair by email (this example uses a GPG key pair tied to `hi@tobyscott.dev`)
```bash
# Export the public key to a file
$ gpg --export -a "hi@tobyscott.dev" > public.key

# Export the private key to a file (this step will ask you for a passphrase)
$ gpg --export-secret-key -a "hi@tobyscott.dev" > private.key
```

Securely transfer these keys to the new computer. Now, on the new computer:

- Import the GPG key pair from the files
```bash
# Import the public key
$ gpg --import public.key

# Import the private key (this step will ask you for the passphrase you chose in the step above)
$ gpg --import private.key
```

- Trust your imported key pair in the new environment
> ***Note:** If you run `$ gpg --list-keys` you will see the `uid` line now reads `[ unknown ]` rather than `[ ultimate ]`. This is because your newly imported key pair doesn't exist in your new environment's GPG trust database (stored at `~/.gnupg/trustdb.gpg`).*
```
$ gpg --edit-key hi@tobyscott.dev

gpg> trust

Please decide how far you trust this user to correctly verify other users' keys
(by looking at passports, checking fingerprints from different sources, etc.)

  1 = I don't know or won't say
  2 = I do NOT trust
  3 = I trust marginally
  4 = I trust fully
  5 = I trust ultimately
  m = back to the main menu

Your decision? 5
Do you really want to set this key to ultimate trust? (y/N) y

gpg> save
```



### Generate a fresh GPG key pair

  - Open a terminal and run the following command:
```bash
$ gpg --full-generate-key
```
  - Select `RSA and RSA` (This will use RSA to firstly sign your commit, and then encrypt it).
  - Choose key size of `4096` bits
  - Choose when you want your key to expire
  - Enter your details as prompted
    - You must use the email that will be tied to your commits (more on this below)
    - Use the comment to describe the purpose for the GPG key
  - Confirm your details and you will be prompted to enter a "passphrase" (you will use this to "sign" your commits, do not share this with anyone)

## 3. Add your GPG key to your GitHub Account (see reference [here](https://docs.github.com/en/authentication/managing-commit-signature-verification/adding-a-new-gpg-key-to-your-github-account))
  - List the long form of the GPG keys for which you have both a public and private key:
```bash
$ gpg --list-secret-keys --keyid-format=long
```
It should output something like:
```
C:\Users\toby\AppData\Roaming\gnupg\pubring.kbx
------------------------------------
sec   4096R/3AA5C34371567BD2 2016-03-10 [expires: 2017-03-10]
uid                          [ultimate] Toby Scott (Key for tutorial) <hi@tobyscott.dev>
ssb   4096R/42B317FD4BA89E7A 2016-03-10
```
  - Print out your full GPG key, in ASCII armor format (using the GPG key ID, `3AA5C34371567BD2` in above example)
```bash
$ gpg --armor --export 3AA5C34371567BD2
```
  - Copy your GPG key, beginning with `-----BEGIN PGP PUBLIC KEY BLOCK-----` and ending with `-----END PGP PUBLIC KEY BLOCK-----`
  - In GitHub, navigate to `Settings > Access > SSH and GPG keys`
    - Click [New GPG key]
    - In the "Key" field, paste the GPG key you just copied
    - Click [Add GPG key] and enter your password to confirm the action.

## 4. Configuring Git to use GPG

  - Locate the installed GPG executable by running:
```bash
# Linux and Mac
$ which gpg

# Windows
$ where.exe gpg
```

  - Tell Git where to find the GPG executable:
```bash
$ git config --global gpg.program "path/to/exe/here"
```
On Mac and Linux, you can do this easily like so:
```bash
$ git config --global gpg.program $(which gpg)
```

- Tell Git to sign all commits:
> ***Note:** The global config will be used for all repos on your machine. If you want to configure a specific repo differently, you can `cd` into that repo and run the `$ git config` command with the `--local` flag instead.*
```bash
$ git config --global commit.gpgsign true
```

  - Tell Git which GPG key to use to sign the commits (use your own GPG key ID from before):
```bash
$ git config --global user.signingkey "3AA5C34371567BD2"
```

> *Final note: The email used in the GPG key must match the email tied to the commits (set in your git config). Additionally, GitHub requires that email to match your GitHub email address. If you have email privacy turned on the email will look something like `12345678+username@users.noreply.github.com`, so that is the email you will have to use when creating your GPG key pair. However, you can manually set the email address tied to commits in a specific repo with `git config --local user.email "hi@tobyscott.dev`", but you will have to ensure that email address is added to your GitHub account. You can read more about this [here](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-user-account/managing-email-preferences/setting-your-commit-email-address) in the GitHub docs.*
