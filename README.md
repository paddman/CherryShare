# CherryShare

**AI-native Secure Private Cloud**

CherryShare is a self-hosted file sync, sharing, collaboration, security, and AI knowledge platform. It is designed as a modern alternative to traditional private-cloud file platforms, with S3-native storage, multi-tenant architecture, zero-trust controls, and AI features built in from the start.

## Vision

CherryShare = File Sync + Collaboration + AI Knowledge + Zero Trust Security

## Planned capabilities

- Web file manager
- Desktop sync client
- S3-compatible object storage
- Chunked/resumable upload
- File versioning and conflict handling
- Public/private sharing with expiry and passwords
- User, group, tenant, RBAC, quota and audit logs
- WebDAV compatibility
- LDAP / Active Directory / OIDC / MFA
- Semantic search, OCR, document summaries and natural-language file search
- Malware scanning, DLP, ransomware behavior detection and immutable versions
- Multi-tenant / white-label / subscription-ready architecture

## Stack

- Backend: Go
- Web: React / Next.js
- Metadata: PostgreSQL
- Object storage: MinIO / S3-compatible storage
- Cache: Redis
- Event bus: NATS
- Search: OpenSearch (planned)
- AI: pluggable OpenAI-compatible/local LLM providers (planned)

## Repository structure

```text
cmd/server/          CherryShare API server
internal/config/     configuration
internal/httpapi/    HTTP routes and handlers
internal/storage/    storage abstractions
web/                 web client (next phase)
docs/                architecture and roadmap
deploy/               local/production deployment assets
```

## Local development

```bash
docker compose up -d postgres redis minio nats
go run ./cmd/server
```

Then open `http://localhost:8080/healthz`.

## Status

Early development scaffold. The first implementation focuses on a clean storage/control-plane foundation before adding desktop sync and AI workflows.
