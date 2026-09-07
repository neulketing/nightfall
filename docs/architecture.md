# Nightfall Documentation

## Architecture

Nightfall is a modular attack emulation suite that simulates advanced persistent threat (APT) campaigns. It follows a phased approach:

1. **Reconnaissance**: OSINT collection, vulnerability scanning, target profiling
2. **Initial Access**: Phishing, browser exploits, supply chain attacks
3. **Persistence**: Backdoors, rootkits, service persistence
4. **Privilege Escalation**: Local and remote privilege escalation
5. **Lateral Movement**: Network traversal, credential harvesting
6. **Credential Access**: Password dumping, keyloggers, browser theft
7. **Exfiltration**: Data staging, C2 channels, covert transfer

## Modules

### Core (`core/`)
- Workspace management
- Logging and auditing
- Configuration management

### Reconnaissance (`recon/`)
- OSINT data collection
- Vulnerability scanning
- Target profiling

### Initial Access (`initial-access/`)
- Spear-phishing email generation
- Browser exploit delivery
- Supply chain attack simulation

### Persistence (`persistence/`)
- Backdoor deployment
- Rootkit installation
- Service persistence mechanisms

### Privilege Escalation (`privilege-escalation/`)
- Local privilege escalation exploits
- Remote privilege escalation
- Token manipulation

### Lateral Movement (`lateral-movement/`)
- Network traversal
- Credential harvesting
- Service discovery

### Credential Access (`credential-access/`)
- Password dumping
- Keylogger deployment
- Browser credential theft

### Exfiltration (`exfiltration/`)
- Data staging
- C2 channel communication
- Covert data transfer

### Evasion (`evasion/`)
- EDR bypass techniques
- Log manipulation
- Traffic obfuscation

### C2 (`c2/`)
- Command and control server
- Agent management
- Protocol implementation

### AI (`ai/`)
- Attack vector selection
- Payload optimization
- Target profiling

## Usage

### Initialization
```bash
nightfall init /path/to/workspace
```

### Reconnaissance
```bash
nightfall recon scan --target example.com --output /path/to/workspace
```

### Phishing Campaign
```bash
nightfall phishing generate --template spearphishing --targets admin@example.com user@example.com
```

### C2 Server
```bash
nightfall c2 server --addr :8080 --output /path/to/workspace
```

### Attack Campaign
```bash
nightfall campaign execute --target example.com --strategy spearphishing --output /path/to/workspace
```

## Testing

```bash
make test
```

## Dependencies

- Go 1.21+
- Python 3.11+
- Rust 1.75+
- Docker & Docker Compose
