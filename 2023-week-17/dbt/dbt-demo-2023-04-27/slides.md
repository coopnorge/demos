---
theme: default
background: https://source.unsplash.com/collection/94734566/1920x1080
highlighter: shiki
lineNumbers: true
---

# Data Engineering and Integration
## Nikita Barskov, Senior ML Engineer at Coop Norge

2023-04-27

---

# What are we going to discuss today?

Data engineering and integration are two interdependent concepts
that play a crucial role in our work at Coop Norge.

**Data engineering** involves designing and building the infrastructure
and systems that enable organisations to collect, store, process,
and analyse large volumes of data.

In enterprise organisations, **integration** refers to the process of
connecting various systems and applications across the entire organisation
to ensure seamless communication and data exchange.

Today, I will demonstrate how data engineering principles and tools have
evolved up to the present day, and how integration can make use of these
principles to operate efficiently and pragmatically.

---

# The ETL and ELT

ETL (Extract, Transform, Load) and ELT (Extract, Load, Transform) are two data integration methods used to move data from various sources to a target data warehouse or data store.

- **ETL** involves extracting data from source systems, transforming it to meet
  the target system's requirements, and loading it into the destination data
  warehouse.

- **ELT** involves loading the data into the destination first and then
  transforming it within the destination system.

---

# Example

## ETL

Suppose you want to transfer data from your website's database to a data
warehouse. In this scenario, you would use ETL to extract data from the
e-commerce website's database, transform it into a suitable format for
the data warehouse, and then load it into the data warehouse.

## ELT

Suppose you receive real-time data from measuring devices and want to store this data in a data warehouse. In this scenario, you would use ELT to load the
real-time data directly into the data warehouse, and then transform it into a suitable format for downstream analysis and reporting.

---

# We focus on T

In both ETL and ELT, the "T" stands for transformation,
which involves a series of typical operations to transform
data from raw to high-quality data for consumption.
Some common data transformations include:

- Cleansing
- Deduplication
- Enrichment
- Normalisation
- Aggregation
- Deriving new data

These transformations are performed to ensure that the data is consistent,
accurate, and in the appropriate format for downstream analysis and reporting.
The specific transformations used will depend on the nature of the data being
processed and the goals of the data integration project.

---

# Let's do some time travelling

Before the advent of modern transformation tools, people interacted with data
transformations primarily through traditional Extract, Transform, Load (ETL)
processes using tools such as SQL and scripting languages like Python and Perl.

These ETL processes were often complex and difficult to manage, especially as
data volumes grew larger and the number of data sources increased. Maintaining
and updating ETL code required a significant amount of time and effort, and
errors could easily creep in, leading to incorrect or incomplete data.

Additionally, traditional ETL processes often lacked version control and
automated testing, making it challenging for teams to collaborate effectively
and ensure the quality of their transformations.

---

# dbt

**dbt (Data Build Tool)** is an open source tool for building and managing data
transformation pipelines. With dbt, you can define and execute complex
transformations on your data, including cleaning, aggregation, joining and
filtering.

By using dbt, you can:

- Modularise and centralise transformation code in a VCS
- Test and document your transformations
- Focus on the actual logic, while dbt takes care of materialisation.
