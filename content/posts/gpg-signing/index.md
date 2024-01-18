---
date: 2022-04-18T18:48:05+11:00
title: "Signing your commits with GPG (GNU Privacy Guard)"
description: "Signing your commits with a GPG private key verifies it was really you who committed the changes."
tags: ["gpg", "git", "signed commits", "cryptography", "security", "identity", "verification"]
cover:
  image: "cover.jpg"
  alt: "Encryption"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

## 1. Pre-requisites

### Install GPG and PINEntry if not already installed

For Mac, install via [Homebrew](https://brew.sh/):

```bash
brew install gpg pinentry-mac
```

For Linux, install via your distro's package manager.

For Windows, you can either use WSL and install via the distro's package manager, or download [GPG4Win](https://www.gpg4win.org/get-gpg4win.html) and install the following components:
  - GnuPG
  - Kleopatra
  - GpgOL
  - GpgEX

### Check what GPG key pairs you have already

```bash
# List all public keys
gpg --list-keys

# List all private keys
gpg --list-secret-keys
```

## 2. Generate or import your GPG key pair

### Option 1: Generate a new GPG key pair

```bash
gpg --full-generate-key
```

- Select `RSA and RSA`
- Choose key size of `4096` bits
- Choose when you want your key to expire
- Enter your details as prompted
  - You must use the email that will be tied to your commits (more on this below)
  - Use the comment to describe the purpose for the GPG key ("Git commit signing key")
- Confirm your details and you will be prompted to enter a "passphrase". You will use this passphrase when signing your commits to prove it was you who made the commit. Make sure you remember your passphrase, and do not share it with anyone.

### Option 2: Import an existing GPG key pair from another computer

On the computer that has the existing GPG key pair, export the key pair and write the exported keys to files. This example uses a GPG key pair tied to `hi@tobyscott.dev`.

```bash
# Export the public key to a file
gpg --export -a "hi@tobyscott.dev" > public.key

# Export the private key to a file. This step will ask you for the key's passphrase.
gpg --export-secret-key -a "hi@tobyscott.dev" > private.key
```

Securely transfer these files to the new computer, and now perform all following steps on the new computer.

- Import the GPG key pair from the files

```bash
# Import the public key from the file
gpg --import public.key

# Import the private key from the file. This step will ask you for the key's passphrase.
gpg --import private.key
```

- Trust your imported key pair in the new environment. If you run `$ gpg --list-keys` you will see the `uid` line now reads `[unknown]` rather than `[ultimate]`. This is because your newly imported key pair doesn't exist in your new environment's GPG trust database (stored at `~/.gnupg/trustdb.gpg`). Let's trust it now.

```bash
gpg --edit-key hi@tobyscott.dev
```

This will open an interactive prompt. Run `trust`, choose "I trust ultimately" and confirm your choice. Finally run `save` to save and exit the interactive prompt.

Now list your GPG keys and notice that `[unknown]` has changed to `[ultimate]`.

## 3. Add your public GPG key to your GitHub Account (see reference [here](https://docs.github.com/en/authentication/managing-commit-signature-verification/adding-a-new-gpg-key-to-your-github-account))

You need to give GitHub your GPG public key so they are able to verify your signed commits against it. Remember to never give anyone your private key or passphrase, not even GitHub!

List the long form of the GPG keys for which you have both a public and private key:

```bash
gpg --list-secret-keys --keyid-format=long
```

It should output something like:

```
/home/toby/.gnupg/pubring.kbx
-----------------------------
sec   rsa4096/9359241A3D7708E5 2024-01-18 [SC]
      1C66CEB6A85DDFE83D3695899359241A3D7708E5
uid                 [ultimate] Toby Scott (Commit signing key) <hi@tobyscott.dev>
ssb   rsa4096/51B12CCB2D65F7B3 2024-01-18 [E]
```

The part *after* `rsa4096/` on the *sec* line is the GPG key ID. In the example above, the GPG key ID is: `9359241A3D7708E5`.

Now export the public key to a file with the following command:

```bash
gpg --armor --export 9359241A3D7708E5 > public.key
```

Copy your entire GPG public key, all the way from `-----BEGIN PGP PUBLIC KEY BLOCK-----` to `-----END PGP PUBLIC KEY BLOCK-----` to your clipboard.

Go to GitHub, navigate to **Settings > Access > SSH and GPG keys**
- Click **New GPG key**
- In the “Key” field, paste the GPG key you just copied
- Click **Add GPG key** and enter your password to confirm the action.

## 4. Configuring Git on your machine to use GPG

Tell Git where to find the GPG executable:

```bash
git config --global gpg.program $(which gpg)
```

Tell Git to sign all commits:

```bash
git config --global commit.gpgsign true
```

Tell Git which GPG key to use to sign the commits by referencing the GPG key ID:

```bash
git config --global user.signingkey "9359241A3D7708E5"
```

## Final notes

The global Git config will be used for all repos on your machine. If you want to configure a specific repo differently, make sure your working directory is the repo you want to configure, and replace the `--global` flag with `--local` instead.

Also note, the email used in the GPG key must match the email tied to the commits (set in your git config). Additionally, GitHub requires that email to match your GitHub email address. If you have email privacy turned on the email will look something like `12345678+username@users.noreply.github.com`, so that is the email you will have to use when creating your GPG key pair. However, you can manually set the email address tied to commits in a specific repo with `git config --local user.email "hi@tobyscott.dev`", but you will have to ensure that email address is added to your GitHub account. You can read more about this [here](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-user-account/managing-email-preferences/setting-your-commit-email-address) in the GitHub docs._
