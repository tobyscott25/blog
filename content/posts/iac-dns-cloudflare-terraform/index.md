---
date: 2024-03-25T18:04:33+11:00
title: "Infrastructure as Code DNS records with Cloudflare and Terraform"
description: "..."
tags:
  [
    "DNS",
    "Infrastructure",
    "Infrastructure as Code",
    "Terraform",
    "Hashicorp",
    "Cloudflare",
    "GitHub Actions",
    "Pipelines",
    "CICD",
    "GitOps",
    "AWS",
    "S3",
    "DynamoDB",
    "Deployment",
    "Production",
    "Security",
    "Networking",
  ]
# author: ["Toby Scott", "Other example contributor"]
hidden: true
draft: true
---

I've been meaning to start using Terraform to define infrastructure as code in my personal projects for a while now. So, I couldn't help myself when I learnt that Cloudflare has a Terraform provider! What better way to start than with my domains' DNS records, and applying the changes through GitHub actions?

### What is a Terraform backend

Terraform has the concept of a "backend" which stores the state of your infrastructure and associates what you have defined in code with the actual resources that have been provisioned. **By default, your state is local.** So if you run the `terraform init`, `terraform plan` and `terraform apply` commands, you will see a `terraform.tfstate` file created in the directory you ran the commands in. **This becomes a problem if you have a team of people** working with the same infrastructure, or perhaps you want to manage your infrastructure (via Terraform) from multiple devices. **The solution to this is using a remote "backend".**

There a number of [Terraform backends](https://developer.hashicorp.com/terraform/language/settings/backends/configuration) to choose from, and I decided to go with [S3](https://developer.hashicorp.com/terraform/language/settings/backends/s3). So instead of storing the state file locally on my machine, it is stored in an S3 bucket in the cloud. In addition, the S3 backend supports using a DynamoDB table for "state-locking" to ensure that **only one** actor is editing the state at any given time.

This means we can now alter the defined infrastructure from any device, and only one person can make a change at any given time.

> Note: The device must have the required AWS credentials to access the S3 bucket and DynamoDB table, but more on that later.

### Defining the resources for the backend

So we have a Cloudflare Terraform provider for defining Cloudflare resources, but we also have an AWS Terraform provider for defining AWS resources, so not use it now since we need to create some AWS resources? This feels oddly recursive because we are defining with Terraform, the resources that Terraform itself uses, but it is great practice!

As mentioned before, we need an S3 bucket to hold the state file and a DynamoDB table for state-locking. Let's take a look at defining these resources as code with Terraform.

```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS provider
provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "tobyscott-tf-state"
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_sse" {
  bucket = aws_s3_bucket.terraform_state.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_versioning" "terraform_state_versioning" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_dynamodb_table" "terraform_state_locks" {
  name         = "tobyscott-tf-state-locks" # Table name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}
```

Here we have told terraform which providers it needs, told the AWS provider which region we want the resources deployed to, and declaratively defined the resources we need and some options such as encryption and bucket versioning.

Because this will be attempting to manipulate AWS resources, we need to set the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables in the shell we're running the Terraform commands from.

```bash
export AWS_ACCESS_KEY_ID='***'
export AWS_SECRET_ACCESS_KEY='***'
```

Finally, we need to `terraform init`, `terraform plan` and `terraform apply`. If you look in AWS you should see your shiny new resources that you created by defining in code.

### CI/CD

1. Create a GitHub repo. I named mine `infra`. Original, I know.
2.
