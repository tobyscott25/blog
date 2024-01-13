---
date: 2024-01-08T10:35:52+11:00
title: "Authenticating Git and GitHub CLI over SSH"
description: "Configuring Git and GitHub CLI to authenticate over SSH is easy, more secure, and a great step for those wanting a deeper understanding of Git configuration and authentication."
tags:
  [
    "Linux",
    "Kubuntu",
    "Ubuntu",
    "GitHub",
    "Git",
    "SSH",
    "Keys",
    "Passphrase",
    "Authentication",
    "Configuration",
    "CLI",
    "Shell",
    "ed25519",
    "Identity",
    "Security"
  ]
# author: ["Toby Scott", "Other example contributor"]
hidden: false
draft: false
---

### Prerequisites

Install GitHub CLI if you haven't already

```bash
sudo apt install gh
```

Check that you are logged out.

```bash
gh auth status

You are not logged into any GitHub hosts. Run gh auth login to authenticate.
```

If you're still logged in, logout.

```bash
gh auth status

github.com
  ✓ Logged in to github.com as tobyscott25 (keyring)
  ✓ Git operations for github.com configured to use https protocol.
  ✓ Token: gho_************************************
  ✓ Token scopes: gist, read:org, repo, workflow

gh auth logout

✓ Logged out of github.com account 'tobyscott25'
```

### Generating an SSH key pair

Generate an SSH key pair with `ssh-keygen`. Choose a long and secure passphrase that you will remember.

```bash
ssh-keygen -t ed25519 -C "your.email@example.com"
```

By default, it will create a `id_ed25519` private key file and `id_ed25519.pub` public key file in your `~/.ssh` directory.

### Authenticating GitHub CLI via SSH

GitHub CLI will automatically generate and use a new key pair if there's no existing SSH config for the `github.com` host. So let's configure it now.

```bash
vim ~/.ssh/config
```

Add the following config, just remember to update the `User` to your own GitHub username and update the `IdentityFile` to the path to your generated private key.

```conf
Host github.com
  HostName github.com
  User tobyscott25
  IdentityFile ~/.ssh/id_ed25519
```

Now we've configured SSH for the `github.com` host, let's authenticate with the specified key. You can specify the protocol with the `-p` flag, like so:

```bash
gh auth login -p ssh
```

GitHub CLI will detect that you have configured an existing key for `github.com` so it won't generate a new one, but it will still ask you to sign in to GitHub through a web browser, enter a login code and upload the public key to GitHub.

Just follow the prompts, but make sure to upload the public key to GitHub, that part is vital for authenticating over SSH. Now confirm you are authenticated by running the following commands:

```bash
# Check the GitHub CLI authentication status
gh auth status

# Use GitHub CLI to view your repos
gh repo list
```

### Configuring Git itself

We've authenticated the GitHub CLI over SSH, but we're still yet to actually configure Git itself so we can make commits. You will need to specify your name and email address like so:

```bash
git config --global user.name "Toby Scott"
git config --global user.email "your.email@example.com"
```

Only the commit author name and email are required, but there is a lot more you can configure if you wish to. Personally, I like to do the following:

```bash
# Change the default branch name from 'master' to 'main'
git config --global init.defaultBranch main

# Set Vim as the default editor for writing commit messages
git config --global core.editor "vim"
```

### Clone a repository

The command syntax for cloning a repository over SSH is slightly different from HTTPS. Here's how you do it:

```bash
# Clone over HTTPS
git clone https://github.com/tobyscott25/blog

# Clone over SSH
git clone git@github.com:tobyscott25/blog
```

### Using SSH agent to store your key and manage your passphrase

If the repository you cloned is private you will have been asked to enter your SSH key's passphrase. You will be prompted for this every time it requires authentication when talking to the remote (GitHub), which can get annoying fast. Let's use `ssh-agent` to handle that for us.

First, run the following command to start the SSH agent in the background:

```bash
eval $(ssh-agent -s)
```

That command will start it if it's not already running. Now let's add your private key to the agent.

```bash
ssh-add $HOME/.ssh/id_ed25519
```

You may need to update the path to match the correct path to your SSH private key. Finally, you'll be prompted for your passphrase, and it will be added to the SSH agent. You can confirm that it has been added by running the following command to list all keys that have been added:

```bash
ssh-add -l
```

Lastly, we want to start the SSH agent automatically if it's not already running when we open a shell, add these lines to your `~/.bash_profile` if you use bash, or `~/.zprofile` if you use ZSH:

```sh
if [ -z "$SSH_AUTH_SOCK" ]; then
  # Check for a currently running instance of the agent
  RUNNING_AGENT="`ps -ax | grep 'ssh-agent -s' | grep -v grep | wc -l | tr -d '[:space:]'`"
  if [ "$RUNNING_AGENT" = "0" ]; then
    # Launch a new instance of the agent
    ssh-agent -s &> $HOME/.ssh/ssh-agent
  fi
  eval `cat $HOME/.ssh/ssh-agent`
fi
```

(Script sourced from: https://code.visualstudio.com/remote/advancedcontainers/sharing-git-credentials)

### Conclusion

Congratulations, you've successfully configured Git to authenticate over SSH. 🎉

Note: I'm running Kubuntu, the commands and paths may differ slightly if you're on Mac or Windows, but essentially the steps wil be the same.
