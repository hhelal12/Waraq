# Waraq — Database Schema Specification

This document details the architectural blueprint and relational database schema design for **Waraq**, a social mobile platform for book lovers. The schema has been custom engineered for **PostgreSQL**, incorporating strict relational integrity rules, indexing strategies, cascading constraints for administrative management, and infrastructure mappings for OCR text extraction pipelines.

---

## 1. System Entities & Relations

The Waraq application domain comprises six core domain entities and three relational junction mechanisms:

* **USERS**: The central domain identity. Manages local authentication credentials, Google OAuth sub-identifiers, and structural application routing privileges (Standard vs. Administrator).
* **BOOK**: The foundational catalog asset metadata. Tracks literary material uploaded to the index.
* **REVIEWS**: A transactional critical assessment asset bound structurally to a single `USER` and a single `BOOK`.
* **COMMENTS**: A multi-tiered engagement mechanism enabling nested dialogue threads attached to critical `REVIEWS`.
* **Quotes**: A specialized textual snippet asset extracted either through manual interface text input or via the automated OCR document scanner processing loop.
* **FOLLOWS**: A self-referential graph link logging unidirectional subscription relationships between unique instances of the `USERS` entity.
* **REVIEW_LIKES & QUOTE_LIKES**: Relational cross-reference junction structures tracking transactional social endorsement telemetry.

---

## 2. Complete Entity-Relationship Blueprint

### USERS
Maintains standard identity credentials, structural audit hooks, and platform security flags.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | `PRIMARY KEY`, `DEFAULT gen_random_uuid()` | Unique immutable internal surrogate system identifier. |
| `email` | `VARCHAR(255)` | `UNIQUE`, `NOT NULL` | Certified unique communication routing address. |
| `username` | `VARCHAR(50)` | `UNIQUE`, `NOT NULL` | Screen display moniker. Cleansed for alphanumeric characters. |
| `password_hash` | `VARCHAR(255)` | `NULLABLE` | Argon2id / bcrypt cryptographic passhash. Nullable for pure OAuth signups. |
| `google_id` | `VARCHAR(255)` | `UNIQUE`, `NULLABLE` | Multi-provider external key mapped from incoming Google ID tokens. |
| `role` | `VARCHAR(20)` | `NOT NULL`, `DEFAULT 'user'` | RBAC marker flag. Allowed options restrict to `user` or `admin`. |
| `created_at` | `TIMESTAMP WITH TIME ZONE` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | Temporal tracking record logging system registration. |

### BOOK
A centralized corpus register tracking metadata records of literary cataloging items.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | `PRIMARY KEY`, `DEFAULT gen_random_uuid()` | Core system immutable asset pointer. |
| `title` | `VARCHAR(255)` | `NOT NULL` | Full literal title catalog entry. Indexed for lexical lookups. |
| `author` | `VARCHAR(255)` | `NOT NULL` | Plain text naming value of the primary content producer. |
| `cover_image_url` | `VARCHAR(2048)` | `NULLABLE` | Fully qualified URL pointer routing to public storage object asset. |
| `description` | `TEXT` | `NULLABLE` | Uncapped synopsis detailing back-cover editorial summaries. |
| `created_at` | `TIMESTAMP WITH TIME ZONE` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | System ingestion timestamp log. |

### REVIEWS
The relational connection linking a user profile evaluation onto a specific ledger book.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | `PRIMARY KEY`, `DEFAULT gen_random_uuid()` | Unique evaluation transaction pointer. |
| `user_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | Identity reference link tracking the entry author. |
| `book_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `BOOK(id)`, `ON DELETE CASCADE` | System link back to targeted cataloged volume. |
| `review_text` | `TEXT` | `NOT NULL` | Main block of qualitative commentary. |
| `rating` | `INTEGER` | `NOT NULL`, `CHECK (rating >= 1 AND rating <= 5)` | Scaled metric ranking evaluating structural value. |
| `updated_at` | `TIMESTAMP WITH TIME ZONE` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | Mutating timestamp updated via system lifecycle triggers. |

### COMMENTS
Enables multi-user interaction loops underneath cataloged platform reviews.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | `PRIMARY KEY`, `DEFAULT gen_random_uuid()` | Transaction identifier. |
| `user_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | Author identity mapping link. |
| `review_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `REVIEWS(id)`, `ON DELETE CASCADE` | Parent anchor thread tracking contextual location. |
| `comment_text` | `TEXT` | `NOT NULL` | Explicit text string data block content. |
| `created_at` | `TIMESTAMP WITH TIME ZONE` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | Temporal ledger instantiation timestamp entry. |

### Quotes
Textual snippets emphasizing highlight segments extracted manually or with OCR assistance.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | `PRIMARY KEY`, `DEFAULT gen_random_uuid()` | Target database item tracking pointer. |
| `user_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | Ownership authorization linkage block. |
| `book_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `BOOK(id)`, `ON DELETE SET NULL` | Optional text context tracking anchor link. |
| `quote_text` | `TEXT` | `NOT NULL` | Verified electronic layout format characters. |
| `raw_image_url` | `VARCHAR(2048)` | `NULLABLE` | Remote path linking straight to un-extracted original image binary. |
| `created_at` | `TIMESTAMP WITH TIME ZONE` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | Immutable system submission sequence timing log. |

### FOLLOWS
A specialized intersection tracking table designed for asynchronous self-referencing relationship routing.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `follower_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | The active agent executing the subscription link sequence. |
| `following_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | The targeted model profile accepting the subscription link. |

* **Primary Key Strategy**: Composite Key composed of (`follower_id`, `following_id`).
* **Validation Rule**: A relational database constraint rule ensures `follower_id != following_id` to prevent profile loop anomalies.

### REVIEW_LIKES
High-throughput tracking junction mapping user profiles endorsing book review commentary.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `user_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | Link logging the endorsing user agent. |
| `review_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `REVIEWS(id)`, `ON DELETE CASCADE` | Target target critical assessment log pointer. |

* **Primary Key Strategy**: Composite Key composed of (`user_id`, `review_id`).

### QUOTE_LIKES
High-throughput tracking junction mapping user profiles endorsing platform quotes.

| Field Name | Data Type | Constraints / Modifiers | Description |
| :--- | :--- | :--- | :--- |
| `user_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `USERS(id)`, `ON DELETE CASCADE` | Link logging the endorsing user agent. |
| `quote_id` | `UUID` | `FOREIGN KEY` $
ightarrow$ `Quotes.id`, `ON DELETE CASCADE` | Target quote reference locator key. |

* **Primary Key Strategy**: Composite Key composed of (`user_id`, `quote_id`).

---

## 3. Relational Architecture Rules & Operations

### Data Normalization & Structural Integrity
The architecture is structured strictly inside **Third Normal Form (3NF)** to avoid transactional anomalies:
* User display properties (`username`, `email`) exist exclusively inside the `USERS` database segment. Downstream functional tables access identity contextual references strictly via the `user_id` immutable UUID foreign key.
* Total aggregated count metrics (such as total likes, total followers, or average rating scores) are not written into core structural columns to avoid stale cache sync states. These values are computed instantly during application query runs using high-performance subqueries, views, or real-time indexes.

### Admin Destructive Actions & Cascading Rules
Administrative maintenance procedures require automated downstream cleanups to avoid data corruption:
* `ON DELETE CASCADE` is set on every direct transactional child relation (`REVIEWS`, `COMMENTS`, `Quotes`, `FOLLOWS`, and all liking junctions). When an Admin flags a malicious account profile or a regular user deletes their account, removing that line from `USERS` triggers a hardware cascading thread that purges every downline element automatically.
* `ON DELETE SET NULL` is applied exclusively to the `book_id` field inside the `Quotes` table. This design preserves user quotes if an Admin deletes a `BOOK` catalog record. The text remains safe inside the platform, with its `book_id` reference simply shifting safely to `NULL`.

### Indexing Optimization Roadmap
To maintain microsecond server responses as social feeds expand, explicit secondary B-Tree database indexes must be declared manually on the following targets:
1.  `FOLLOWS (following_id)`: Accelerates computing reverse profile counting metrics (e.g., pulling a user's total count of active followers).
2.  `REVIEWS (book_id)` & `REVIEWS (user_id)`: Speeds up rendering chronological book landing walls and historical dashboard summaries.
3.  `COMMENTS (review_id)`: Speeds up deep parsing algorithms rebuilding complex conversational nested components underneath popular profile items.

### Image-to-Quote OCR Processing Sequence Pipeline
The implementation of the image capture and text conversion pipeline follows this strict sequence:
1.  **Binary Transport Sequence**: The mobile client takes an image of a physical page and uploads it to an isolated, secure object bucket storage location (e.g., Supabase Storage). This returns a secure public tracking locator string (`raw_image_url`).
2.  **OCR Character Isolation Sequence**: The backend pulls the file from storage and passes the stream to a computing OCR worker module to parse the physical assets into a raw text string.
3.  **Client Modification Interlock**: The processing server returns the raw string content back to the mobile interface screen, enabling the user to highlight, edit typos, or add contextual notes.
4.  **Database Commit Execution**: The user clicks "Post", prompting the application to write a clean row to the `Quotes` table, mapping the final string into `quote_text` and storing the original tracking link inside `raw_image_url` for administrative audit trails.