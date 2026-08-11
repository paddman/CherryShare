# CherryShare Architecture

## Product boundary

CherryShare is split into a control plane, data plane, sync plane, security plane, and AI plane so storage traffic does not become coupled to administration or model workloads.

## High-level architecture

```text
Web / Desktop / Mobile
        |
        v
API Gateway + Auth
        |
  +-----+------------------------------+
  |            Control Plane           |
  | Tenant | User | RBAC | Quota       |
  | Sharing | Audit | Policy | Billing |
  +-----+------------------------------+
        |
        +------------------+
        |                  |
        v                  v
   Sync Plane          Data Plane
Watcher/Journal      Upload/Download
Chunk + Hash         Version/Metadata
Conflict Resolver    S3 Object Storage
        |                  |
        +--------+---------+
                 |
                 v
            Event Bus (NATS)
              /       \
             v         v
      Security Plane   AI Plane
      AV/DLP/Ransom    OCR/RAG/Search
      Audit/Policies   Summary/Tagging
```

## Core design rules

1. Store file bytes in S3-compatible object storage, not PostgreSQL.
2. PostgreSQL stores metadata, ownership, versions, permissions and audit references.
3. Uploads are resumable and chunk-addressed.
4. File versions are immutable references; logical files point to the active version.
5. Every object belongs to a tenant from day one.
6. Authorization is enforced server-side on every object operation.
7. AI is asynchronous and event-driven. Normal file access must not depend on an LLM being available.
8. Security scans happen before risky objects become broadly shareable.
9. Desktop sync uses a local journal/database instead of repeatedly scanning entire trees.
10. Provider adapters keep S3, identity, search and AI engines replaceable.

## Initial service modules

- Identity and tenant service
- File metadata service
- Upload/session service
- Object storage adapter
- Sharing service
- Version service
- Audit service
- Sync protocol
- Event publisher/consumer

## AI roadmap

- OCR and text extraction
- Semantic indexing
- Natural-language file search
- Document summaries
- Auto tagging and classification
- Duplicate/near-duplicate detection
- Sensitive document classification
- Permission-risk suggestions

## Security roadmap

- Malware scanning
- DLP policies
- Ransomware change-rate detection
- Immutable recovery snapshots
- Device sessions and trust policy
- Country/IP restrictions
- Signed and expiring shares
- Full audit trail
