# Xray-core — Custom Build (External Auth)

This is a custom fork of [XTLS/Xray-core](https://github.com/XTLS/Xray-core) with **pure external authentication** and **traffic statistics API** support.

## Changes

This fork adds the following modules on top of the upstream Xray-core:

- **`app/auth`** — External authentication module
  - HTTP-based auth forwarding: forwards user credentials to an external auth center via HTTP POST
  - `PermissiveValidator`: accepts all UUIDs at the protocol layer; actual authentication is delegated to the external auth center
  - Supports both connection-level and request-level authentication
  - User identity (authID) is written into `session.User.Email`, automatically propagated to stats, routing, and logging

- **`app/trafficstats`** — Traffic statistics and online user management
  - Real-time connection tracking via `ConnTracker`
  - HTTP API for traffic queries and user management:
    - `GET /traffic` — per-user uplink/downlink traffic
    - `GET /traffic?clear=1` — query and reset counters
    - `GET /online` — list of online users
    - `POST /kick` — disconnect a specific user
  - Optional management listener with shared-secret authentication

- **Configuration** — Two new top-level JSON fields
  - `"auth"` — auth center URL, node token, node ID, protocol, timeout
  - `"trafficStats"` — management API listen address and shared secret

- **VLESS inbound modifications**
  - Pure external auth mode: no `clients` array needed in VLESS inbound settings
  - `PermissiveValidator` registered as default validator when no users are configured
  - Every inbound connection triggers external authentication via `app/auth`

## Upstream

This fork tracks [XTLS/Xray-core](https://github.com/XTLS/Xray-core). All custom changes live on the `custom-main` branch.

```bash
git log main..custom-main --oneline   # view custom commits
```

To sync with upstream:

```bash
git fetch upstream
git checkout main && git reset --hard upstream/main && git push origin main --force
git checkout custom-main && git rebase main
```

## Compilation

```bash
CGO_ENABLED=0 go build -o xray -trimpath -ldflags="-s -w -buildid=" ./main
```

## License

[Mozilla Public License Version 2.0](LICENSE) — same as upstream.

## Credits

Forked from [XTLS/Xray-core](https://github.com/XTLS/Xray-core). All credit for the original work goes to the XTLS project and its contributors.
