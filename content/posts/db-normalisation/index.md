---
date: 2024-03-25T18:04:33+11:00
title: "Database normalisation"
description: "..."
tags: []
# author: ["Toby Scott", "Other example contributor"]
hidden: true
draft: true
---

### Steps for Transformation:

1. **Transform Strong Entities**: Create tables for all strong entities, which include all their attributes. The primary key in the relational table corresponds to the entity's key attribute.
2. **Transform Weak Entities**: Create tables for weak entities and include their attributes along with a foreign key to the owning strong entity.
3. **Transform 1-to-1 Relationships**: For each 1-to-1 relationship, choose one of the tables and add a foreign key to it that references the other table.
4. **Transform 1-to-Many Relationships**: In the table corresponding to the "many" side of the relationship, add a foreign key that references the table on the "one" side.
5. **Transform Many-to-Many Relationships**: Create a new table to represent the relationship and include foreign keys that reference the tables on both sides of the relationship.
6. **Transform Multi-Valued Attributes**: Create a new table for each multi-valued attribute, and use a foreign key to reference the owning entity.
7. **Transform n-ary Relationships**: Create a new table for the n-ary relationship, including foreign keys to reference all participating entities.
8. **Transform Specialisation/Generalisation**: Depending on the type (disjoint/overlapping, total/partial), choose one of the following options:
   - **Option 8A**: Separate tables for superclass and each subclass.
   - **Option 8B**: Tables only for subclasses, including superclass attributes.
   - **Option 8C**: Single table with a 'type' attribute.
   - **Option 8D**: Single table with multiple 'type' attributes.
9. **Transform Union Types (Categories)**: Create a new table for the category with a surrogate key. Use this surrogate key as a foreign key in the tables for the defining superclasses.

### Final Step:

- **Repeat Steps 2 to 7** for tables corresponding to subclasses created in Steps 8 and 9.

These steps provide a systematic way to convert an EER diagram into a set of relational tables, which is a crucial part of database design.

# Comparing the above steps with a

### Similarities:

1. **Strong Entities to Tables**: Both approaches recommend creating a table for each strong entity, including its attributes and primary key.
2. **Weak Entities to Tables**: Both mention that a table should be created for weak entities, incorporating a composite primary key that includes the primary key of the owning entity.
3. **1:1 Relationships**: Both suggest that for 1:1 relationships, one table should contain a foreign key referencing the primary key of the other table.
4. **1:N Relationships**: Both suggest that the table on the "many" side should contain a foreign key referencing the primary key of the table on the "one" side.
5. **M:N Relationships**: Both recommend creating a new table for many-to-many relationships, incorporating foreign keys from both participating tables.
6. **Multi-valued Attributes**: Both recommend creating a separate table for multi-valued attributes.
7. **n-ary Relationships**: Both suggest that an n-ary relationship should result in a new table with foreign keys referencing all participating entities.

### Differences:

1. **Specialisation/Generalisation**: Your content doesn't explicitly mention transforming specialisation/generalisation hierarchies from the EER model, while my earlier explanation did. This is a step that applies when you're moving from an EER model to a relational model, and it's often crucial for database design.
2. **Union Types (Categories)**: My original list also included the transformation of union types or categories, which is more specific to EER diagrams and is not mentioned in your content.
3. **Foreign Key Addition**: Your content specifically states that foreign keys are not drawn in the ER model and are only added during the transformation to tables, which is a good practice to avoid confusion during the ER modelling phase.
4. **Listing of Final Tables**: Your content emphasises the importance of listing the final tables, which is a good practice for clarity and documentation but wasn't explicitly mentioned in my initial list.

Overall, the guidelines in your content are quite comprehensive for ER model to relational table transformations, but they might not cover some specific cases that arise in Enhanced Entity-Relationship (EER) models.
