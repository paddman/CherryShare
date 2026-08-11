# CherryShare Roadmap

## Phase 1 — Foundation

- [x] Go API scaffold
- [x] Health/version endpoint
- [x] Local PostgreSQL / Redis / NATS / MinIO stack
- [ ] Configuration package and secret loading
- [ ] PostgreSQL migrations
- [ ] Tenant model
- [ ] User/session model
- [ ] RBAC middleware
- [ ] S3 storage adapter
- [ ] Audit event model

## Phase 2 — File Core

- [ ] Create folders
- [ ] List directory
- [ ] Upload sessions
- [ ] Multipart/resumable upload
- [ ] SHA-256 / chunk hashing
- [ ] Download streaming
- [ ] Move / rename / delete / restore
- [ ] Version history
- [ ] Quotas
- [ ] Trash retention

## Phase 3 — Sharing

- [ ] Internal user/group shares
- [ ] Public links
- [ ] Password-protected links
- [ ] Share expiry
- [ ] Download/upload-only links
- [ ] Link revocation
- [ ] Sharing audit log

## Phase 4 — Web Application

- [ ] CherryShare design system
- [ ] Login / first boot
- [ ] File explorer
- [ ] Grid/list views
- [ ] Upload manager
- [ ] Share drawer
- [ ] Version history UI
- [ ] Admin console
- [ ] Storage dashboard
- [ ] Activity timeline

## Phase 5 — Desktop Sync

- [ ] Desktop agent architecture
- [ ] Local SQLite journal
- [ ] Filesystem watcher
- [ ] Incremental scan fallback
- [ ] Chunk/delta synchronization
- [ ] Conflict resolution
- [ ] Selective sync
- [ ] Offline queue
- [ ] Bandwidth limits
- [ ] Windows shell integration
- [ ] macOS/Linux clients

## Phase 6 — Enterprise

- [ ] LDAP / Active Directory
- [ ] OIDC / SSO
- [ ] MFA
- [ ] Device/session management
- [ ] Multi-tenant administration
- [ ] White-label configuration
- [ ] Usage metering
- [ ] Subscription/billing hooks
- [ ] HA deployment
- [ ] Backup/restore
- [ ] S3 replication / object lock

## Phase 7 — Security

- [ ] Malware scanning pipeline
- [ ] DLP engine
- [ ] Sensitive-data classification
- [ ] Ransomware behavior detection
- [ ] Mass-delete/change protection
- [ ] Immutable recovery versions
- [ ] IP/country access policies
- [ ] Risk-based share controls

## Phase 8 — Cherry AI

- [ ] Text extraction
- [ ] OCR
- [ ] Embedding/index pipeline
- [ ] Semantic file search
- [ ] Ask-your-files RAG
- [ ] Document summaries
- [ ] Automatic tags
- [ ] Duplicate/near-duplicate detection
- [ ] Permission-risk recommendations
- [ ] Local/OpenAI-compatible LLM providers

## Phase 9 — Collaboration

- [ ] WebDAV
- [ ] Office integration (OnlyOffice/Collabora compatible)
- [ ] Comments and mentions
- [ ] File requests
- [ ] Team spaces
- [ ] Calendar/contact integration where useful

## Release targets

- `v0.1`: authenticated file API + S3 storage + tenant isolation
- `v0.2`: sharing + versions + web file manager
- `v0.3`: desktop sync MVP
- `v0.4`: enterprise auth + security controls
- `v0.5`: AI search and document intelligence
- `v1.0`: production-ready private cloud platform
