---
theme: apple-basic
layout: intro-image
image: https://images.unsplash.com/photo-1667751867092-a19e8c66e3ea?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2274&q=80
---

<div class="absolute top-10">
    <span class="font-700">
        Nikita Barskov
    </span>
</div>

<div class="absolute bottom-10">
    <h1>Technology Demo Cloud SQL</h1>
    <p>30th November 2022</p>
</div>

---

# Agenda

- What is Cloud SQL?
- Are there any obscure issues?
- What is the expected developer experience with Cloud SQL?
- Coop Norge Cloud SQL Terraform module design proposal
- QA

---
layout: section
---

# What is Cloud SQL?
## Let's talk about this cloud resource

- What is that about?
- How does it work?
- How many different ways of deployment you might have?

---

# What is that about?

Cloud SQL provides a cloud-based alternative 
to local and on-premise MySQL, PostgreSQL, 
and SQL Server databases. 

## What problem is Cloud SQL solving?

You **should use** Cloud SQL if you want to spend 
less time managing your database and more time using it.

## How is it solving this problem?

Cloud SQL manages your database instance. You manage your data.

---
layout: image-right
image: 'static/cloud-sql-architecture.png'
---

# How does it work?

- Each Cloud SQL instance is powered by a virtual machine (VM)
  running on a host Google Cloud server.

- The database is stored on a scalable, durable network storage device called a
  persistent disk that attaches to the VM. 

- A static IP address sits in front of each VM to ensure that the IP address an
  application connects to persists throughout the lifetime of the Cloud SQL
  instance.

---
layout: statement
---

# How many ways do you potentially have to deploy Cloud SQL instance?

---

# The options

- You can deploy Cloud SQL instance with public or private IP
- You can require TLS enforcement
- You can use IAM based authentication

> I do not cover an actual implementation of the database
> want to choose, I do not cover specific parameters
> you have for a specific implementation.

Let's bring some boring math (combinatorics) and say 
we want to calculate **a sequence without repetition**:

$$
A^k_n = \dfrac{n!}{(n-k)!}, A^1_2 = 2
$$

And then calculate the number of combinations:

$$
(A^k_n)^m, (A^1_2)^3 = 8
$$

---
layout: fact
---

<iframe src="https://giphy.com/embed/8Odq0zzKM596g" width="480" height="270" frameBorder="0" class="giphy-embed" allowFullScreen></iframe>

---
layout: statement
---

# How many implementations do we have at Coop Norge?

---
layout: fact
---

## 3.

---
layout: section
---

# Are there any obscure issues?
## Let's talk about the challenges we had so far

---

# Generic questions

## Security

TLS, Google Cloud IAM, internal database roles.

### Example:

What is the difference between role `roles/cloudsql.instanceUser` and
role `roles/cloudsql.client`?

## Scalability

What type of the virtual machine I should choose for this
specific amount of data?

## Connection

How to connect to Cloud SQL from Google Managed resources,
API Platform or even from my MacBook?

---
layout: section
---

# What is the expected developer experience with Cloud SQL?
## How can we approach the challenges with Cloud SQL?

---

# As a developer, I want a product...

- allows me to deploy Cloud SQL instance,
- the instance is available from the source I configure,
- I can configure identities with permissions,
- I can configure the hardware,
- also has a documented way to access this Cloud SQL instance from
  GKE, Google Cloud resource or my MacBook.

---
layout: statement
---

# Laget med ~~🥲~~ 💙 hos CoopX
