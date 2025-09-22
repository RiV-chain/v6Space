# v6Space™ - Decentralized IPv6 Overlay Network Infrastructure

## Introduction

v6Space™ is a fully end-to-end encrypted decentralized IPv6 overlay network infrastructure that enables secure connectivity between a wide spectrum of endpoint devices like IoT devices, desktop computers, or even routers.

It is lightweight, self-arranging, supported on multiple platforms and allows pretty much any IPv6-capable application to communicate securely with other v6Space™ nodes. v6Space™ does not require you to have IPv6 Internet connectivity - it also works over IPv4.

## Core Functionality

v6Space™ provides the decentralized IPv6 overlay network infrastructure for applications like CupLink™, which requires a VPN service to establish the decentralized mesh network environment for peer-to-peer communication.

### Key Differences from RiV-mesh

v6Space™ represents a significant evolution from the original RiV-mesh project:

- **Enhanced User Experience**: Modern, responsive web interface with professional design
- **MCP Server Integration**: Built-in Model Context Protocol server running on the decentralized overlay network
- **Self-Generated IPv6 Addressing**: Advanced cryptographic addressing with EUI-64 format
- **AI-First Design**: Native integration with AI platforms and tools
- **Improved Performance**: Optimized decentralized mesh networking with better peer discovery
- **Professional Attribution**: Licensed UI work by RiV Chain™ Limited under CC BY-NC 4.0

## Supported Platforms

v6Space™ works on a number of platforms, including Linux, macOS, Ubiquiti EdgeRouter, VyOS, Windows, FreeBSD, OpenBSD, OpenWrt, and Android.
## Building

If you want to build from source, as opposed to installing one of the pre-built packages:

1. Install [Go](https://golang.org) (requires Go 1.19 or later)
2. Clone this repository
3. Build for your platform:

### Linux/macOS
```bash
./build
```

### Windows
```powershell
go mod tidy
go build -o mesh.exe ./cmd/mesh
go build -o meshctl.exe ./cmd/meshctl
go build -o genkeys.exe ./cmd/genkeys
```

### Cross-compilation
You can cross-compile for other platforms and architectures by specifying the `GOOS` and `GOARCH` environment variables, e.g. `GOOS=windows ./build` or `GOOS=linux GOARCH=mipsle ./build`

... or generate an iOS framework with:
```
./contrib/mobile/build -i
```

... or generate an Android AAR bundle with:
```
./contrib/mobile/build -a
```

## Running

### Generate configuration

To generate static configuration, either generate a HJSON file (human-friendly, complete with comments):

```
./mesh -genconf > /path/to/mesh.conf
```

... or generate a plain JSON file (which is easy to manipulate programmatically):

```
./mesh -genconf -json > /path/to/mesh.conf
```

You will need to edit the `mesh.conf` file to add or remove peers, modify other configuration such as listen addresses or multicast addresses, etc.

### Run v6Space™

To run with the generated static configuration:

```
./mesh -useconffile /path/to/mesh.conf
```

To run in auto-configuration mode (which will use sane defaults and random keys at each startup, instead of using a static configuration file):

```
./mesh -autoconf
```

You will likely need to run v6Space™ as a privileged user or under `sudo`, unless you have permission to create TUN/TAP adapters. On Linux this can be done by giving the v6Space™ binary the `CAP_NET_ADMIN` capability.

### Access the Web Interface

Once v6Space™ is running, you can access the modern web interface at:

```
http://localhost:19019
```

The web interface provides:
- **Real-time mesh network status** and peer connections
- **Configuration editor** for easy setup management  
- **MCP server control** for AI integration features
- **Network analytics** and performance monitoring
- **Peer management** with decentralized mesh topology visualization

### API Control

You can also control v6Space™ programmatically using the REST API endpoints:

```bash
# Get node information
curl http://localhost:19019/api/self

# MCP Server Control
curl http://localhost:19019/api/mcp/status       # Get MCP server status
curl -X POST http://localhost:19019/api/mcp/start   # Start MCP server
curl -X POST http://localhost:19019/api/mcp/stop    # Stop MCP server
curl -X POST http://localhost:19019/api/mcp/sync    # Sync UI tools to server
curl http://localhost:19019/api/mcp/info        # Get MCP server info
curl http://localhost:19019/api/mcp/secret      # Get access secret (localhost only)
curl -X POST http://localhost:19019/api/mcp/secret/regenerate # Generate new secret

# View all API endpoints
curl http://localhost:19019/doc
```

## Modern Web UI & Advanced Features

### Enhanced User Interface
v6Space™ features a modern, responsive web interface (Licensed under CC BY-NC 4.0 - RiV Chain™ Limited) that provides:
- **Professional UI Design**: Clean, modern interface with responsive design
- **Real-time Network Monitoring**: Live status updates and decentralized overlay network visualization  
- **Configuration Management**: Easy-to-use configuration editor
- **Peer Management**: Visual peer discovery and decentralized mesh connection management
- **Network Analytics**: Comprehensive network statistics and performance metrics
- **Zero-Configuration Assets**: JavaScript and CSS files load automatically (no manual setup required)

## Tools and Features

### MCP (Model Context Protocol) Integration

v6Space™ includes a built-in **MCP server** that runs on top of the self-generated IPv6 mesh network address, providing:
- **AI Tool Integration**: Direct integration with AI platforms like Claude and OpenWebUI
- **Network Management Tools**: AI-assisted network configuration and troubleshooting
- **Real-time Server Control**: Start/stop MCP server directly from the web interface
- **Platform Agnostic**: Works across Windows, macOS, and Linux platforms
- **JSON-RPC 2.0 Compliance**: Full Model Context Protocol support
- **🔒 Secure Authentication**: Bearer token-based security for exposed tools and APIs
- **🔑 Dynamic Secret Management**: Cryptographically secure access tokens with regeneration capability

#### MCP Security Model
The MCP server implements robust security to protect your mesh network tools:

- **Bearer Token Authentication**: All MCP requests require `Authorization: Bearer <access_secret>` header
- **Localhost Secret Access**: Access secrets can only be retrieved from localhost (127.0.0.1/::1)
- **Cryptographic Secrets**: 256-bit randomly generated access tokens
- **Secret Rotation**: Regenerate access tokens anytime via API
- **Network Isolation**: MCP server runs on configurable port (default: 3001)

```bash
# Get current access secret (localhost only)
curl http://localhost:19019/api/mcp/secret

# Use secret to access MCP tools
curl -H "Authorization: Bearer v6s-ABC123..." http://localhost:3001/mcp/tools/list
```

### Self-Generated IPv6 Addressing

v6Space™ uses **EUI-64 format** for IPv6 address creation based on device identifiers, providing:
- **Static Cryptographic IP Addressing**: Secure, predictable IPv6 addresses
- **No DHCP Required**: Self-configuring network addressing
- **Peer Authentication**: Cryptographically verified peer identities

### Multicast Communication

v6Space™ implements multicast communication for:
- **Peer Discovery**: Automatic discovery of nearby mesh nodes
- **Auto-routing**: Dynamic decentralized mesh topology management
- **Zero-configuration Networking**: Plug-and-play decentralized mesh connectivity

## Documentation

- [Installing v6Space™](https://v6space.net/docs/installation-guide.html) - Complete installation guide with system requirements and platform-specific instructions
- [Configuring v6Space™](https://v6space.net/docs/configuration-guide.html) - Detailed configuration options and best practices
- [Frequently asked questions](https://v6space.net/docs/faq.html) - Common questions and troubleshooting guide
- [Version changelog](v6Space-develop/CHANGELOG.md) - Release notes and updates

## Downloads

### Pre-built Packages
- **Windows (x64)**: [Download v6Space™ for Windows](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-x64.msi)
- **Windows (x86)**: [Download v6Space™ for Windows 32-bit](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-x86.msi)
- **macOS (Intel)**: [Download v6Space™ for macOS Intel](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-macos-amd64.pkg)
- **macOS (Apple Silicon)**: [Download v6Space™ for macOS Apple Silicon](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-macos-arm64.pkg)
- **Linux (Ubuntu/Debian)**: [Download .deb package](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-amd64.deb)
- **Linux (RHEL/CentOS/Fedora)**: [Download .rpm package](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-1.x86_64.rpm)
- **FreeBSD**: [Download FreeBSD package](https://github.com/RiV-chain/v6Space-builds/releases/latest/download/v6space-0.5.0.1-freebsd-amd64.txz)

### All Downloads
- [View all releases and platforms](https://github.com/RiV-chain/v6Space-builds/releases) - Complete list of available downloads
- [Source code](https://github.com/RiV-chain/v6Space) - Build from source

### Security & API Documentation
- [Security Policy](SECURITY.md) - Comprehensive security guide and vulnerability reporting
- [MCP Security Guide](docs/MCP-SECURITY.md) - Detailed MCP server security documentation  
- [API Documentation](docs/API.md) - Complete REST API and MCP server API reference
- [GitHub Security Policy](.github/ISSUE_TEMPLATE/security.md) - How to report security issues

## Community

Connect with us on social media:

- [Twitter](https://x.com/rivchain)
- [Reddit](https://reddit.com/r/rivchain)
- [LinkedIn](https://linkedin.com/company/rivchain)

## Public peers
If you are operating v6Space™ peer and may create your pull request with your new peer or use existing ones: https://github.com/RiV-chain/public-peers

## Troubleshooting

### Common Issues

#### IPv6 Support Issues
If you encounter errors like:
```
An error occurred starting multicast: listen udp6 [::]:9001: socket: address family not supported by protocol
An error occurred starting TUN/TAP: operation not supported
```

**Solution**: Ensure your system has IPv6 support enabled. v6Space™ requires IPv6 functionality to operate correctly.

#### Permission Issues
If you get permission errors when starting v6Space™:

**Linux**: Give the binary the `CAP_NET_ADMIN` capability:
```bash
sudo setcap cap_net_admin+ep ./mesh
```

**Windows**: Run as Administrator or ensure you have TUN/TAP adapter permissions.

#### MCP Server Issues
If the MCP server fails to start:

1. Check if port 3001 is available
2. Verify the access secret is properly generated
3. Check the web interface for error messages
4. Restart v6Space™ and try again

## License

v6Space™ uses a dual-license model with clear legal and technical separation:

### Core Mesh Networking (LGPLv3)
- **License**: GNU Lesser General Public License v3.0 (LGPLv3)
- **Scope**: All mesh networking functionality, REST API, protocol implementations, command-line tools
- **Commercial Use**: Allowed with full LGPLv3 compliance requirements
- **Source Code**: Available at [GitHub](https://github.com/RiV-chain/v6Space)

### v6Space™ Interface (CC BY-NC 4.0)
- **License**: Creative Commons Attribution-NonCommercial 4.0 International
- **Scope**: Web interface, HTML/CSS/JavaScript assets, MCP server UI controls
- **Commercial Use**: Prohibited (personal and educational use only)
- **Attribution**: Required when using UI components

For complete licensing details, see:
- [Core License](LICENSE) - LGPLv3
- [Interface License](contrib/ui/mesh-ui/ui/LICENSE-UI) - CC BY-NC 4.0
- [Component Separation Guide](https://rivchain.org/legal/component-separation.html)
