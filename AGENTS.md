# Rules for AI agents and contributors

This repository is **donutvpn-core**: the Go layer between
[donutvpn-sing-box](donutvpn-sing-box/) (a fork of sing-box, a git submodule here) and the DonutVpn
apps. It builds the sing-box configuration from the user's profile and settings, runs it, and
exposes it to the app over gRPC and libbox. These rules hold for every change, by a person or an
AI agent. The sing-box fork has its own rules in
[donutvpn-sing-box/AGENTS.md](donutvpn-sing-box/AGENTS.md) and its `donutvpn-feature` skill; follow
them whenever a change touches that submodule.

## 1. Know the layers, and put the change in the right one

| Layer | Where | Job |
|---|---|---|
| upstream sing-box | `github.com/sagernet/sing-box` | the engine. **Never edited.** |
| donutvpn-sing-box | `donutvpn-sing-box/` (replaces upstream in `go.mod`) | the engine plus DonutVpn features. General, driven only by its config. |
| donutvpn-core | this repo | decides what config to build and with which values, runs it, serves the app. |
| donutvpn-app | `github.com/donutvpn/donutvpn-app` | the UI. Talks to donutvpn-core only through the gRPC API in `v2/hcore/*.proto`. |

A change goes in the lowest layer that can hold it without knowing anything about the layers
above it. Use this table before writing code:

| You need | Put it in |
|---|---|
| a protocol, transport, DNS, group, routing or connection behaviour | donutvpn-sing-box, as a config option with a default |
| the value of such an option, or whether DonutVpn uses it | donutvpn-core, `v2/config/` (`builder.go`, `chain.go`) |
| a setting the user can change | a field of `DonutVpnOptions` (`v2/donutvpnoptions/`), then the builder, then the app |
| reading a share link or subscription | `ray2sing/` for links, `v2/config/` for sing-box JSON |
| something the app asks the core, or the core tells the app | the gRPC API, `v2/hcore/hcore.proto` and `hcore_service.proto` |
| start/stop, crash and diagnostics handling, the admin background service | `v2/hcore/` |
| a platform difference (Android, iOS, macOS tunnel, Windows, Linux, router) | `Makefile` tags, `v2/hutils/`, `platform/`, conditions in the builder |
| build tags, cronet/naive, CI | `Makefile` first; CI reads it |

## 2. Keep donutvpn-sing-box changes minimal

- Do not change code that comes from `github.com/sagernet/sing-box`. It is replaced by
  `./donutvpn-sing-box` in `go.mod`; every edited upstream line is a conflict at the next rebase of
  the fork.
- In the fork, new code goes in `h_*.go` files; a changed upstream line ends with `//H`. Tests are
  `h_*_test.go`, docs are in English and Persian under `docs/donutvpn/`. All of this is in the
  fork's `AGENTS.md`; read it before touching the submodule.
- Before adding anything to the fork, read `donutvpn-sing-box/docs/donutvpn/` and search the code.
  Most "missing" features already exist under another name. Extend what exists; never add a
  parallel option, type or package.
- A fork change is committed on the fork's `v5` branch and pushed first. Then this repo bumps the
  submodule pointer in its own commit, `build(deps): bump donutvpn-sing-box to <id>`. The same
  holds for `ray2sing/` and the forks under `donutvpn-sing-box/replace/`.

## 3. donutvpn-sing-box stays general

- Everything the fork does is controlled by its JSON config: every behaviour is an option in
  `option/` with a documented default that is the same for every user of the fork.
- The fork knows nothing about DonutVpn: no profiles, app settings, platforms, or outbound tags
  such as `select`, `lowest`, `balance`.
- A fixed number in the fork (a tolerance, timeout, count, URL) is a bug. Make it an option and
  give it a sensible default when unset or zero.
- Old configs keep loading with the same meaning. New options are additions only.
- A config the fork cannot run fails when it loads, with an error that says what to change.

## 4. donutvpn-core sets the parameters

- donutvpn-core is the policy layer. It chooses which fork options to send and their values, in
  `v2/config/`, from `DonutVpnOptions` and the parsed profile.
- When the fork gains an option that matters to DonutVpn, set it explicitly in the builder with
  the value DonutVpn wants, and a one-line comment on why that value. Do not rely on the fork's
  default for anything the app depends on.
- User-facing settings are fields of `DonutVpnOptions`, with the default in
  `DefaultDonutVpnOptions()`. There is no second settings path.
- **Never hide a donutvpn-sing-box bug here.** No retries, delays, extra conditions, swallowed
  errors or state guessing that make a failure in the fork invisible. Fix it in the fork, bump the
  submodule, and say so in the commit body with the fork commit id. donutvpn-core may only add
  handling that is correct whether or not the bug exists, such as surfacing the error to the app.
- The gRPC API is a contract with the app. A changed `.proto` is a breaking change unless it only
  adds fields or RPCs; say in the commit what the app must regenerate or change.

## 5. Build

- Submodules are required: `git submodule update --init --recursive`.
- Build tags live only in the `Makefile`: `CORE_TAGS`, `PROTOCOL_TAGS`, `CLI_ADD_TAGS` and the
  per-platform `*_ADD_TAGS`. `TAGS` (libraries), `CLI_TAGS` (the CLI) and `TUNNEL_TAGS` (the macOS
  system extension) are composed with `$(call tags,...)`. CI and `cmd.sh` read them with
  `make -s print-<VAR>`. Never copy a tag list into a workflow, script or doc.
- Naive (cronet) is part of every build unless disabled with `make NAIVE=0 <target>`, for targets
  that have no cronet library.
- Targets: `make android`, `make ios`, `make macos`, `make linux-amd64`, `make windows-amd64`,
  `make cli`. After a `.proto` change run `make protos` (Go stubs and `docs/donutvpnrpc.md`); the
  app regenerates its Dart stubs from the same files.
- Never commit built libraries, `bin/`, `data/`, logs, `*.json` data files, keys, or real profile
  links, IPs or tokens.

## 6. Test

```bash
TAGS=$(make -s print-TAGS)
go vet -tags "$TAGS" ./v2/...
go test -tags "$TAGS" ./v2/...
```

- Tests are `*_test.go` next to the code. `.gitignore` excludes `**/*_test.go`, so they are not
  committed; write them anyway, keep them on disk, and run them before every commit.
- A bug fix comes with a test that fails without the fix and passes with it. Check both.
- A config change is checked on a built config: `make cli`, then build a config from a profile
  (`./bin/DonutVpnCli config build <profile>` or `cmd.sh`) and read the option in the output.
- A change the app will see is verified in donutvpn-app on Android and on one desktop platform.

## 7. Document

- Anything donutvpn-app must change or know goes where the app's developers look:
  `donutvpn-sing-box/docs/donutvpn/app_considerations.md` for fork changes, and the commit body
  plus the related app pull request for core changes (with the proto diff when the API changed).
- A user-visible change gets a line in `HISTORY.md` under "(unreleased)".

## 8. Commit

- One feature or fix per commit, Conventional Commits: `feat(config): ...`, `fix(hcore): ...`,
  `build(deps): bump donutvpn-sing-box to <id>`, `docs: ...`, `build(Makefile): ...`. The body
  says why, and how it was tested.

- `git fetch git && git rebase git/main` before pushing (the remote is named `git`). Push the
  fork and `ray2sing` before pushing a commit here that bumps their pointers.
- If a change spans repos, link the fork commit or pull request in the commit body.
