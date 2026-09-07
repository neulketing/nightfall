# Nightfall — Advanced Threat Emulation Suite

A comprehensive self-hosted toolkit for simulating and researching advanced persistent threat (APT) techniques, zero-day exploits, and modern attack vectors.

## Architecture

```
nightfall/
├── core/           # Shared infrastructure, config, logging
├── reconnaissance/ # OSINT, target discovery, vulnerability scanning
├── initial-access/ # Phishing, spear-phishing, browser exploits
├── persistence/    # Backdoors, rootkits, service persistence
├── privilege-escalation/ # Local/remote privilege escalation
├── lateral-movement/ # Network traversal, credential harvesting
├── credential-access/ # Password dumping, keyloggers, browser theft
├── exfiltration/   # Data staging, C2 channels, covert transfer
├── evasion/        # EDR bypass, log manipulation, obfuscation
├── c2/             # Command & Control frameworks
├── ai/             # AI-powered attack automation
├── supply-chain/   # Developer tool trojanization
├── edge/           # Edge device, IoT, OT attacks
├── honeypots/      # Decoy systems for detection research
├── tests/          # Comprehensive test suite
├── docs/           # Documentation, usage guides, research notes
└── tools/          # Utility scripts, helpers, converters
```

## Features

- **Full MITRE ATT&CK coverage** across 14+ tactics
- **AI-powered** attack selection and payload generation
- **Modular design** — swap components without breaking others
- **Automation-first** — scripted end-to-end attack scenarios
- **Research-grade** — detailed logging and analysis capabilities

## Prerequisites

- Go 1.21+
- Python 3.11+
- Rust 1.75+
- Docker & Docker Compose
- Virtual machine platform (VMware, VirtualBox, or Proxmox)

## Quick Start

```bash
# Clone and initialize
git clone https://github.com/neulketing/nightfall
cd nightfall

# Install dependencies
./install.sh

# Run self-test suite
make test

# Start C2 server
./nightfall c2 server

# Run reconnaissance scan
./nightfall recon scan --target example.com
```

## License

Proprietary — Internal use only.
