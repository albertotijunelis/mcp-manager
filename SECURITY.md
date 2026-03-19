# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| < latest | :x:               |

Only the latest release of mcp-manager is supported with security updates.

## Reporting a Vulnerability

If you discover a security vulnerability in mcp-manager, please report it responsibly.

**Do NOT open a public GitHub issue for security vulnerabilities.**

Instead, email:

📧 **albertotijunelis@gmail.com**

**Subject line:** `mcp-manager security`

### What to include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response time

- **Acknowledgment:** Within 48 hours
- **Initial assessment:** Within 5 business days
- **Fix timeline:** Depends on severity, typically within 14 days for critical issues

### What to expect

1. You'll receive an acknowledgment within 48 hours
2. We'll investigate and assess the severity
3. We'll work on a fix and coordinate the disclosure
4. You'll be credited in the release notes (unless you prefer anonymity)

## Scope

This policy covers:

- The `mcp-manager` CLI binary
- The install script (`install.sh`)
- The registry format and sync mechanism
- The GitHub Actions CI/CD workflows

## Best Practices

mcp-manager follows these security practices:

- No credentials are stored in plaintext (env vars only)
- Registry sync uses HTTPS
- Binary downloads verify SHA256 checksums
- Minimal filesystem permissions (0750 for directories, 0600 for configs)
- No execution of untrusted code from the registry

Thank you for helping keep mcp-manager safe!
