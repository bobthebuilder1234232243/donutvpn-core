Changelog
=========


(unreleased)
------------

New
~~~
- Add the MasterDNS outbound (MasterDnsVPN over public resolvers), with
  masterdns:// links, and read SlipNet slipnet:// DNS tunnel profiles
  (DNSTT, NoizDNS, VayDNS, with SSH) as dnstt outbounds. [donutvpn]
- Read dnst:// links (the dnst URL spec, both forms) for the dnstt, vaydns
  and masterdns transports with the socks, ssh, shadowsocks and custom
  backends. [donutvpn]
- The smart resolver pool is one shared module inside the dnstt and masterdns
  outbounds (smart_pool as true or a settings object); the separate
  smart_dns_pool service is removed. [donutvpn]
- The smart resolver pool relays tunnel queries asynchronously and learns
  each resolver's rate (hmrd_multi_resolver_dns aaf22d9): MasterDNS with
  limited resolvers is about 8 times faster, DNSTT matches its own resolver
  handling, and the pool's sockets now use the outbound's socket protection.
  [donutvpn]

Fix
~~~
- Read share links again with libxray v1.261031.0, whose converter takes a
  key and returns JSON; the core did not build. [donutvpn]
- Ios build. [donutvpn]
- Version. [donutvpn]


v4.0.0 (2026-02-13)
-------------------

New
~~~
- Better proxy availblity when detour. [donutvpn]
- Add musl lib. [donutvpn]
- Add srun to run directly singbox without calling donutvpn core.
  [donutvpn]
- Fallback dns. [donutvpn]
- Add ios limits and ipv6 connectivity. [donutvpn]
- Add debouncer for sending proxy info. [donutvpn]
- Add pause and add urlencoded base64. [donutvpn]

Fix
~~~
- Update delete-artifact job condition to handle failures gracefully.
  [donutvpn]
- Workflow 3. [donutvpn]
- Workflow 2. [donutvpn]
- Workflow 1. [donutvpn]
- Issues. [donutvpn]
- 11. [donutvpn]
- 10. [donutvpn]
- 9. [donutvpn]
- 8. [donutvpn]
- 7. [donutvpn]
- 7. [donutvpn]
- 6. [donutvpn]
- 5. [donutvpn]
- 4. [donutvpn]
- 3. [donutvpn]
- No commit message. [donutvpn]
- Ideal timeout. [donutvpn]
- Share lan. [donutvpn]
- Bug. [donutvpn]
- Bug. [donutvpn]

Other
~~~~~
- Release: version 4.0.0. [donutvpn]
- Remove musl keep it for PR. [donutvpn]
- Add fake dns support, direct, redirect and tproxy. [donutvpn]
- Better replace management. [donutvpn]
- Add libs. [donutvpn]
- Add more os5. [donutvpn]
- Add more os4. [donutvpn]
- Add more os3. [donutvpn]
- Add more os2. [donutvpn]
- Add more os1. [donutvpn]
- Add more os1. [donutvpn]
- Add more os. [donutvpn]
- Fix workflow10. [donutvpn]
- Fix workflow9. [donutvpn]
- Fix workflow8. [donutvpn]
- Fix workflow8. [donutvpn]
- Fix workflow7. [donutvpn]
- Fix workflow6. [donutvpn]
- Fix workflow5. [donutvpn]
- Fix workflow4. [donutvpn]
- Fix workflow3. [donutvpn]
- Fix workflow2. [donutvpn]
- Fix workflow1. [donutvpn]
- Add naive workflow. [donutvpn]
- Add naive workflow. [donutvpn]
- Add naive workflow. [donutvpn]
- Remove naive. [donutvpn]
- Add naive. [donutvpn]
- Disable multi dns temporary. [donutvpn]
- Temp fix. [donutvpn]
- Add naive. [donutvpn]
- Add psiphon reconnect on network change, better url test only active
  on, better logging. [donutvpn]
- Update psiphon. [donutvpn]
- Better dns checker. [donutvpn]
- Fix bug and improvemnt. [donutvpn]
- Fix bug and improvemnt and add psiphon. [donutvpn]
- New balancer approach. [donutvpn]
- New balancer approach. [donutvpn]
- Fix awg. [donutvpn]
- Fix awg. [donutvpn]
- New add amnezia, wiregaurd raw config, wiregaurd noise. [donutvpn]
- Remove extra dependenceis. [donutvpn]
- Fixes and improvements. [donutvpn]
- Update and imporvement and fix bugs. [donutvpn]
- Add warp endpoints, disable warp tricks temporarly, fix bugs and
  improvements. [donutvpn]
- NEW: add pause and wake and fix: android stack. [donutvpn]
- Update workflow. [donutvpn]
- Update to singbox 1.13.0.beta6;add extended features. [donutvpn]
- Update macos. [donutvpn]
- Update logging. [donutvpn]


v3.2.0 (2025-07-11)
-------------------

New
~~~
- Add current profile and current outbound and config name into RPC.
  [donutvpn]
- Add xhttp support. [donutvpn]
- Add ntp and rule  and update to singbox 1.10. [donutvpn]
- Add routing rule. [donutvpn]
- Add NTP server. [donutvpn]
- Add custom url test and dns. [x]
- Update to singbox 1.10.3. [x]
- Update to singbox 10.3 and xray 24.12.18. [x]
- Big changes, stablity and more. [x]
- Add log level to cli. [x]
- Allow domain to be resolved from detour. [x]

Changes
~~~~~~~
- Remove unnecessory codes. [x]

Fix
~~~
- Bugs and improvements. [donutvpn]
- Initial state. [donutvpn]
- Direct and  bypass showing. [x]
- Ios lib. [x]
- Linux lib out. [x]
- Linux build. [x]
- Import. [x]
- Redrect Std error not supported in windows. [x]
- Bug. [donutvpn]

Other
~~~~~
- Update to v3.2.0. [donutvpn]
- Improve dns. [donutvpn]
- Add new regex dns resolver, better dns routing. [donutvpn]
- Add new regex dns resolver, better dns routing. [donutvpn]
- Update warp.go. [DonutVpn]
- Fix warp. [DonutVpn]
- Improve urltest fix: warp  bug. [donutvpn]
- Update. [donutvpn]
- Update linux build. [donutvpn]
- Fix allow insecure. [donutvpn]
- Merge pull request #101 from q4a/dev. [DonutVpn]

  add CI for arm64 Linux build
- Add CI for arm64 Linux build. [Aleksey Komarov]
- Improve support for xray, v2ray configs and fix dns bug. [donutvpn]
- Improve xray support and fix bug. [donutvpn]
- Add memory limit option. [donutvpn]
- Remove old xray config. [donutvpn]
- Better url test. [donutvpn]
- Imporve urltest and add debug symbol. [donutvpn]
- Add xhttp extra support. [donutvpn]
- Speedup staticip resolving and improve xray compatibility. [donutvpn]
- Suppoer full xray configs and update xray to 25.3.6. [donutvpn]
- Update to singbox xray. [donutvpn]
- Improve config generators. [x]
- Improvements and add saved config. [x]
- Update. [x]
- Update. [x]
- Update. [x]
- Update for better handling the background service. [x]
- Update macos. [x]
- Update url test. [x]
- Add more details. [x]
- Add live ipinfo. [x]
- Update. [x]
- Update. [x]
- Add time to log. [x]
- Update. [donutvpn]
- Refactor hcore. [donutvpn]
- Refactor config. [donutvpn]
- Remove old codes. [donutvpn]
- Merge redundant codes. [donutvpn]
- Remove old rpc. [donutvpn]
- Refactor platforms. [donutvpn]
- Refactor config. [donutvpn]
- Refactor cmd. [donutvpn]
- Remve flutter bridge. [donutvpn]
- Add multiplatform utils. [donutvpn]
- Fine grain service manager interface. [donutvpn]
- Refactor runnel service. [donutvpn]
- Better naming the extension functions, add admin-service-extension.
  [donutvpn]
- Remove tunnel service. [donutvpn]
- Update. [donutvpn]
- Update. [donutvpn]
- Update. [donutvpn]
- Fix bug. [donutvpn]
- Update. [donutvpn]
- Massive refactor of entire project with breaking changes. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.8 (2024-10-03)
-------------------

Fix
~~~
- Empty db issue. [donutvpn]

Other
~~~~~
- Release: version 3.1.8. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.7 (2024-10-02)
-------------------

Fix
~~~
- Bug. [donutvpn]

Other
~~~~~
- Release: version 3.1.7. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.6 (2024-10-02)
-------------------

New
~~~
- Add multiprocess safe database. [donutvpn]

Changes
~~~~~~~
- Add multi column fields and allow button inside form. [donutvpn]

Other
~~~~~
- Release: version 3.1.6. [donutvpn]
- Use more effienet data storage. [donutvpn]
- Update extensions. [donutvpn]
- Register extensions indivisually. [donutvpn]
- Update extensions. [donutvpn]
- Remove cancel and use better alternative close and Submit Data.
  [donutvpn]
- Change name to Close and remove cancel --< move it to submitdata.
  [donutvpn]
- Add Button Dialog Close and Ok. [donutvpn]
- Add support for multi col row to ui. [donutvpn]
- Update extensions. [donutvpn]
- Remove Button Field. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.5 (2024-09-29)
-------------------

Fix
~~~
- Exception in standalone mode. [donutvpn]

Other
~~~~~
- Release: version 3.1.5. [donutvpn]
- Better ci/cd. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.4 (2024-09-29)
-------------------
- Release: version 3.1.4. [donutvpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.3 (2024-09-29)
-------------------
- Release: version 3.1.3. [donutvpn]


v3.1.2 (2024-09-29)
-------------------
- Release: version 3.1.2. [donutvpn]
- Update README.md. [DonutVpn]
- Update README.md. [DonutVpn]
- Update docker-compose.yml. [DonutVpn]
- Update README.md. [DonutVpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.1 (2024-09-29)
-------------------

Fix
~~~
- Docker issues. [donutvpn]

Other
~~~~~
- Release: version 3.1.1. [DonutVpn]
- Update WRT package HASH. [donutvpn-com]


v3.1.0 (2024-09-29)
-------------------

Fix
~~~
- Bug. [donutvpn]

Other
~~~~~
- Release: version 3.1.0. [DonutVpn]
- Update WRT package HASH. [donutvpn-com]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Auto update wrt package on each release. [donutvpn]
- Add template extension and demo extension. [donutvpn]
- Update README.md. [DonutVpn]


v3.0.1 (2024-09-29)
-------------------

New
~~~
- Better docker management and add more supported platforms. [donutvpn]
- Store warp info if key starts with p{char} if key[2]=="_" then use the
  rest as license, remove forcing MTU when it has detour. [donutvpn]

Fix
~~~
- Docker. [DonutVpn]
- Windows build. [donutvpn]
- Docker image. [donutvpn]
- Docker image. [donutvpn]

Other
~~~~~
- Release: version 3.0.1. [DonutVpn]
- Release: version 3.1.0. [DonutVpn]
- Fix docker file. [donutvpn]
- Release: version 3.0.1. [DonutVpn]
- Merge pull request #69 from izhdahawrt/main. [DonutVpn]

  wrt: update core to 3.0.0
- Wrt: update core to 3.0.0. [Kaveh Dadgar]


v3.0.0 (2024-09-28)
-------------------

New
~~~
- Add ipscanner. [donutvpn]
- Add extension list. [donutvpn]
- Big Change, Add support for Extensions 😍 [donutvpn]
- Extensions v0. [x]

Changes
~~~~~~~
- Framgnet option. [donutvpn]
- DonutVpnOptions. [x]

Fix
~~~
- Typos. [donutvpn]
- Openwrt donutvpn.init typo. [Kaveh Dadgar]

  typo fix in openwrt donutvpn init file
- Openwrt makefile conffiles. [Kaveh Dadgar]

  fix openwrt makefile conffiles config according to new naming scheme on openwrt.
- Openwrt userid conflict with singbox. [Kaveh Dadgar]

  added a 1 in front of every user id seemed safe[1] and has been previously done for port numbers as well.
  https://en.wikipedia.org/wiki/User_identifier
- Allow connections from lan by default. [Kaveh Dadgar]

  this option allow clients connect to donutvpn's web-ui through ip addresses other than localhost loopback.
- Strict route option to false. [Kaveh Dadgar]

  we can prevent the device becoming unreachable by setting this
  option to false instead of previous true value.
- Typo. [donutvpn]

Other
~~~~~
- Release: version 3.0.0. [DonutVpn]
- Merge pull request #66 from izhdahawrt/izhdahawrt-wrtfocus. [DonutVpn]
- Add: openwrt import webui assets 2.3.0. [Kaveh Dadgar]

  import webui assets from latest 2.3.0 release packages for OpenWrt releases as well.
- Update: openwrt makefile to latest 2.3.0. [Kaveh Dadgar]

  updated to latest version.
  web-ui assets have been imported from release files.
  general polishing according to newly introduced changes.
  commands and config files have been simplified for easier interaction. ( service donutvpn restart. instead of typing down service donutvpn-core restart.
- Delete: openwrt old donutvpn-core.init. [Kaveh Dadgar]

  was not working
- Add: Openwrt new donutvpn.init. [Kaveh Dadgar]
- Delete: old donutvpn-core.conf. [Kaveh Dadgar]

  remove old donutvpn-core.conf for the sake of new donutvpn.conf
- Add: new openwrt configuration. [Kaveh Dadgar]

  shortened donutvpn-core to donutvpn for simplicity.
  removed options workingdir and user as it is not being used at the
  moment also we need root privilege for tun.
  added new option "config" referring to config file or url after -c switch.
  introduced option "appconfig" referring to app config file or url
   after -d switch.

  donutvpn-cli run -c <config file or sublink> -d <DonutVpnApp config file
  or URL>
  https://donutvpn.com/app/DonutVpnCli-guide/#
- Add: new geo format config. [Kaveh Dadgar]
- Remove:old geodb format from config. [Kaveh Dadgar]
- If an outbound has detour, use dns over proxy. [donutvpn]
- Ignore: extra files. [donutvpn]
- Update to sing 1.9.6, fix fragment issue, add: allow_insecure flag.
  [donutvpn]
- Remove: unused. [donutvpn]
- Add Start Extension method. [donutvpn]


v2.3.0 (2024-09-02)
-------------------

New
~~~
- Add option to use xray core when possible. [DonutVpn]

Fix
~~~
- Vmess xray links. [DonutVpn]
- Parse config issues. [DonutVpn]
- Xray core start multiinstance, add use xray core only. [DonutVpn]

Other
~~~~~
- Release: version 2.3.0. [DonutVpn]
- Update to singbox 1.9.4. [DonutVpn]
- Merge pull request #60 from alirezarm78/main. [DonutVpn]

  Fix docker network mode and volume mounts
- Change docker network mode to host. [Alireza Rahmani]
- Fix bugs for docker. [Alireza Rahmani]
- Update LICENSE.md. [DonutVpn]
- Avoid self loop route in android in tun. [DonutVpn]
- Test. [DonutVpn]
- Some changes for xray routing mark. [DonutVpn]
- Add intrrupt connection. [DonutVpn]


v2.1.5 (2024-08-05)
-------------------

Fix
~~~
- Warp random domains. [DonutVpn]
- Tunnel service. [DonutVpn]
- Bug? [DonutVpn]

Other
~~~~~
- Release: version 2.1.5. [DonutVpn]
- Revert changes. [DonutVpn]
- Test. [DonutVpn]
- Simple refactor config.go. [DonutVpn]
- Added default region. [DonutVpn]


v2.1.4 (2024-08-04)
-------------------

New
~~~
- Ignore utls if no fp is in the url, fix instance not stoped. [DonutVpn]

Other
~~~~~
- Release: version 2.1.4. [DonutVpn]


v2.1.3 (2024-08-04)
-------------------

Fix
~~~
- Reenable indeependent nat for vpn server. [DonutVpn]
- When no dns found for skydns. [DonutVpn]

Other
~~~~~
- Release: version 2.1.3. [DonutVpn]


v2.1.2 (2024-08-02)
-------------------
- Release: version 2.1.2. [DonutVpn]
- Add domains with .ru .ir .cn .tr as ignore list. [DonutVpn]


v2.1.1 (2024-08-02)
-------------------
- Release: version 2.1.1. [DonutVpn]


v2.1.0 (2024-08-02)
-------------------

Fix
~~~
- Bugs and improvement. [DonutVpn]
- Urltest bugs. [DonutVpn]

Other
~~~~~
- Release: version 2.1.0. [DonutVpn]
- Change outbound on the fly only on force test or when no outbound
  avilble. [DonutVpn]
- Better tunnel management and cache dns of url test. [DonutVpn]
- Add mark to xray. [DonutVpn]


v2.0.4 (2024-07-31)
-------------------
- Release: version 2.0.4. [DonutVpn]


v2.0.3 (2024-07-30)
-------------------

New
~~~
- Add fragment to xray if not available. [DonutVpn]

Fix
~~~
- Xray fragment bug, add loglevel to xray configs. [DonutVpn]

Other
~~~~~
- Release: version 2.0.3. [DonutVpn]


v2.0.2 (2024-07-30)
-------------------

Fix
~~~
- Config export issue. [DonutVpn]

Other
~~~~~
- Release: version 2.0.2. [DonutVpn]
- Always show fragment option. [DonutVpn]


v2.0.1 (2024-07-29)
-------------------
- Release: version 2.0.1. [DonutVpn]
- Update to xray core 1.8.23. [DonutVpn]


v2.0.0 (2024-07-28)
-------------------
- Release: version 2.0.0. [DonutVpn]
- Add xray fragment support. [DonutVpn]
- Fix warp over proxy bug. [DonutVpn]


v1.9.0 (2024-07-25)
-------------------

New
~~~
- Add xray support. [DonutVpn]

Other
~~~~~
- Release: version 1.9.0. [DonutVpn]


v1.7.0 (2024-07-17)
-------------------

New
~~~
- Remove for apple. [DonutVpn]

Other
~~~~~
- Release: version 1.7.0. [DonutVpn]
- Release: version 1.7.0. [DonutVpn]


v1.6.3 (2024-07-14)
-------------------
- Release: version 1.6.3. [DonutVpn]
- Add socks, phttp and phttps support. [DonutVpn]


v1.6.1 (2024-07-12)
-------------------

Changes
~~~~~~~
- Local sites now use direct dns server for routing. [DonutVpn]

Other
~~~~~
- Release: version 1.6.1. [DonutVpn]


v1.6.0 (2024-07-12)
-------------------

New
~~~
- Add knocker warp noise mode. [DonutVpn]

Other
~~~~~
- Release: version 1.6.0. [DonutVpn]
- Add socks, http, https support. [DonutVpn]


v1.5.1 (2024-07-09)
-------------------

Changes
~~~~~~~
- Move geoip into our github repository for more reliablity. [DonutVpn]

Other
~~~~~
- Release: version 1.5.1. [DonutVpn]


v1.5.0 (2024-07-06)
-------------------
- Release: version 1.5.0. [DonutVpn]
- Update to singbox 1.9.3. [DonutVpn]
- More ads website to block. [DonutVpn]
- Update  macos 11 to macos12. [DonutVpn]
- Add basic routing and block ads, remove old geo assets. [DonutVpn]


v1.3.6 (2024-06-01)
-------------------

Fix
~~~
- Bug. [DonutVpn]

Other
~~~~~
- Release: version 1.3.6. [DonutVpn]


v1.3.4 (2024-06-01)
-------------------

Fix
~~~
- Crash. [DonutVpn]

Other
~~~~~
- Release: version 1.3.4. [DonutVpn]
- Add default tag selector. [DonutVpn]


v1.3.0 (2024-05-31)
-------------------
- Release: version 1.3.0. [DonutVpn]
- Use fix ip for wow. [DonutVpn]
- Add warp in warp & reserved & more more. [DonutVpn]
- Add inbound option and custom flags. [DonutVpn]
- Update to singbox 1.9.0 and update to warp-plus. [DonutVpn]
- Allow full config. [DonutVpn]
- Merge pull request #32 from real-alchemist/main. [DonutVpn]

  Fix docker
- Update Docker/donutvpn.sh. [real-alchemist]
- Update Dockerfile. [real-alchemist]


v1.2.1 (2024-03-24)
-------------------
- Release: version 1.2.1. [DonutVpn]
- Change clash-api-secret to web-secret. [DonutVpn]
- Add default config for wrt. [DonutVpn]


v1.2.0 (2024-03-22)
-------------------

New
~~~
- Add makefile for open wrt. [donutvpn]

Fix
~~~
- Cli. [donutvpn]

Other
~~~~~
- Release: version 1.2.0. [donutvpn]
- Remove debug symbols for cli. [donutvpn]
- Change name of core to donutvpn-core. [donutvpn]
- Rename name to donutvpn core. [DonutVpn]


v1.1.0 (2024-03-19)
-------------------

New
~~~
- Add tproxy and fix crash on invalid fragment config. [DonutVpn]

Other
~~~~~
- Release: version 1.1.0. [DonutVpn]
- Release: version 1.1.0. [DonutVpn]
- Use dynamic idle timeout. [DonutVpn]


v1.0.1 (2024-03-18)
-------------------

New
~~~
- Add random secret if not provided. [DonutVpn]

Other
~~~~~
- Release: version 1.0.1. [DonutVpn]
- Fix; standalone issue and add secret. [DonutVpn]


v1.0.0 (2024-03-18)
-------------------

Fix
~~~
- Issue with singbox 1.8.9. [DonutVpn]

Other
~~~~~
- Release: version 1.0.0. [DonutVpn]


v0.18.0 (2024-03-17)
--------------------

New
~~~
- Add docker image to repository. [DonutVpn]

Fix
~~~
- Disable grpc core. [DonutVpn]
- Tunnel service issue. [DonutVpn]

Other
~~~~~
- Release: version 0.18.0. [DonutVpn]
- Ignore alpn in websocket. [DonutVpn]
- Update to singbox v1.8.9. [DonutVpn]
- Better vpn service handler. [DonutVpn]
- Lower response time. [DonutVpn]
- Refactor global and admin service to use grpc. [DonutVpn]
- Big refactor. make compatible v2 and v1 interface. [DonutVpn]
- Release: version 0.17.9.dev. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next-core.
  [DonutVpn]
- Add core grpc server. [problematicconsumer]
- Refactor docker. [DonutVpn]
- Refactor docker. [DonutVpn]
- Merge pull request #30 from real-alchemist/main. [DonutVpn]

  add dockerfile
- Add dockerfile. [real-alchemist]


v0.17.8 (2024-03-10)
--------------------

Fix
~~~
- Hysteria link. [DonutVpn]

Other
~~~~~
- Release: version 0.17.8. [DonutVpn]


v0.17.7 (2024-03-10)
--------------------

Fix
~~~
- Typo. [DonutVpn]

Other
~~~~~
- Release: version 0.17.7. [DonutVpn]
- Allow webui only if AllowConnectionFromLAN. [DonutVpn]
- Release: version 0.17.7. [DonutVpn]


v0.17.6 (2024-03-10)
--------------------

New
~~~
- Add useragent. [DonutVpn]

Fix
~~~
- Warp setting. [DonutVpn]

Other
~~~~~
- Release: version 0.17.6. [DonutVpn]


v0.17.4 (2024-03-10)
--------------------
- Release: version 0.17.4. [DonutVpn]
- Update ray2sing. [DonutVpn]


v0.17.1 (2024-03-10)
--------------------

Fix
~~~
- Libcore linux amd64. [DonutVpn]
- Linux lib path. [DonutVpn]

Other
~~~~~
- Release: version 0.17.1. [DonutVpn]


v0.17.0 (2024-03-09)
--------------------

New
~~~
- Add all linux support. [DonutVpn]
- Add cli interface with basic webui. [DonutVpn]
- Add donutvpnrpc. [DonutVpn]

Fix
~~~
- Path. [DonutVpn]
- Path issue in download and upload artificat. [DonutVpn]
- Conflict in upload artificat. [DonutVpn]
- Duplicate file name. [DonutVpn]
- Bugs and issues. [DonutVpn]
- Build issue. [DonutVpn]
- Build issue. [DonutVpn]

Other
~~~~~
- Release: version 0.17.0. [DonutVpn]
- Release: version 0.17.0. [DonutVpn]
- Release: version 0.17.0. [DonutVpn]
- Disable cgo for cli only build. [DonutVpn]
- Make release name uniform, use no so for cli only builds. [DonutVpn]
- Release: version 0.17.0. [DonutVpn]
- Update to singbox 1.8.8. [DonutVpn]
- Merge. [DonutVpn]
- Refactor options. [problematicconsumer]
- Fix json keys. [problematicconsumer]
- Add dynamic port. [DonutVpn]
- Update to singbox 1.8.7. [DonutVpn]


v0.16.8 (2024-02-27)
--------------------
- Release: version 0.16.8. [DonutVpn]
- Enable windows service again due to resolution of False Positive in
  windows defender. [DonutVpn]


v0.16.7 (2024-02-24)
--------------------
- Release: version 0.16.7. [DonutVpn]
- Disable windows service due to False Positive in Anti virus. [DonutVpn]


v0.16.4 (2024-02-22)
--------------------
- Release: version 0.16.4. [DonutVpn]
- Make up and down wireguard only with real event. [DonutVpn]


v0.16.1 (2024-02-22)
--------------------

New
~~~
- Add better support for network change! now warp will be reconnected
  after network change. [DonutVpn]
- Add early support for. [DonutVpn]

Changes
~~~~~~~
- Add more logs, add ip selector. [DonutVpn]

Fix
~~~
- Final resolution to network change issue. [DonutVpn]
- Change network issue. [DonutVpn]
- Warp configs, add warp delay, fix network switch. [DonutVpn]
- Bug in gomobile. [DonutVpn]

Other
~~~~~
- Release: version 0.16.1. [DonutVpn]
- Fix netswitch. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next-core.
  [DonutVpn]
- Change warp generation. [problematicconsumer]
- Add more logs. [DonutVpn]
- Add more logs. [DonutVpn]
- Add more logs. [DonutVpn]


v0.16.0 (2024-02-18)
--------------------
- Release: version 0.16.0. [problematicconsumer]
- Add syncthing ignore. [problematicconsumer]
- Add warp account methods. [problematicconsumer]


v0.15.15 (2024-02-16)
---------------------

Changes
~~~~~~~
- Add some logs for interface change listener. [DonutVpn]

Fix
~~~
- Bypass lan. [DonutVpn]

Other
~~~~~
- Release: version 0.15.15. [DonutVpn]
- Add more logs. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next-core.
  [DonutVpn]
- Change interval type. [problematicconsumer]


v0.15.12 (2024-02-15)
---------------------
- Release: version 0.15.12. [DonutVpn]
- Release: version 0.5.12. [DonutVpn]
- Fix config orders. [DonutVpn]


v0.15.11 (2024-02-15)
---------------------
- Release: version 0.15.11. [DonutVpn]


v0.15.10 (2024-02-15)
---------------------
- Release: version 0.15.10. [DonutVpn]
- Add support for regenrate config, make parallel call for warp config.
  [DonutVpn]


v0.15.0 (2024-02-14)
--------------------

New
~~~
- Filter tags includes §hide§ for example for detour. [DonutVpn]
- Add vpn service mode as seperated option, fix crash of wireguard in
  wireguard and more stable wireguard. [DonutVpn]
- Add fallback for commander if not working as a service. [DonutVpn]
- Configurable warp noise parameters, fix stopping issues. [DonutVpn]
- Retest proxies on 10 Failure (may cause too much battery usage)
  [DonutVpn]
- Add more checks in the input json. [DonutVpn]
- Add delay 65535 when node is unavilable. [DonutVpn]

Fix
~~~
- Bug in linux. [DonutVpn]
- MTU. [DonutVpn]
- Fragment. [DonutVpn]
- Multi receiver of command client issue. [DonutVpn]
- Bug and redundant chacks. [DonutVpn]
- Changedir, new: add support of partial singbox configs. [DonutVpn]

Other
~~~~~
- Release: version 0.15.0. [DonutVpn]
- Add auto recovery, send group with their active profiles, make warp
  more robust thanks to uoosef, support multiple receiver for command
  group. [DonutVpn]
- Use random warp ip if it can not be resolved. [DonutVpn]
- Do prestart dns checkings 10 times faster in parallel. [DonutVpn]
- Update go.sum. [DonutVpn]
- Create CONTRIBUTING.md. [DonutVpn]


v0.14.2 (2024-02-09)
--------------------

Fix
~~~
- Bug. [DonutVpn]

Other
~~~~~
- Release: version 0.14.2. [DonutVpn]


v0.14.1 (2024-02-09)
--------------------

Fix
~~~
- Appimage admin service bug. [DonutVpn-com]
- Linux make DonutVpnService bug. [DonutVpn]

Other
~~~~~
- Release: version 0.14.1. [DonutVpn]


v0.14.0 (2024-02-08)
--------------------

New
~~~
- Add CommandGroupInfoOnly. [DonutVpn]
- Move running admin service to core. [DonutVpn]
- Add cmd.sh to use a same tags with production. [DonutVpn]
- Make tunnel service working. [DonutVpn]
- Add chain support & fix warp bug. [DonutVpn]
- Add static ip resolver. [DonutVpn]
- Add default dns with fragment in doh. [DonutVpn]
- Add tunnel service. [DonutVpn]

Changes
~~~~~~~
- Change name of service to DonutVpnService. [DonutVpn]
- Activate tunnel service only in windows and linux. [DonutVpn]
- Ubuntu image to 20.04 for better compatibility. [DonutVpn]
- Release format to donutvpn-core-platfrom. [DonutVpn]

Fix
~~~
- Close box service in case of exception, use source is private and
  disallow make in windows. [DonutVpn]
- Bug of group only info which only zero and 1 group exists. [DonutVpn]
- Update wireguard imports. [DonutVpn]
- Unexpected warp exit! [DonutVpn]
- Compile issue in linux. [DonutVpn]
- Bug in release due to libcore.dll path. [DonutVpn]
- Release bug. [DonutVpn]
- More release bug. [DonutVpn]
- Release bug. [DonutVpn]
- Release bug. [DonutVpn]
- Panic in singbox. [DonutVpn]
- Bug in padding. [DonutVpn]

Other
~~~~~
- Release: version 0.14.0. [DonutVpn]
- Remove dependency to psiphon which use go1.20. [DonutVpn]
- Change to Stop State in case of an error occure. fix connection bug.
  [DonutVpn]
- Merge pull request #19 from omid-the-great/main. [DonutVpn]

  Chnage repository language to golang
- Improve: Undetectable dart api. [Omid The Great]
- Add internal build cli. [problematicconsumer]
- Change config options structure. [problematicconsumer]
- Change config builder using constant tag names. [problematicconsumer]
- Update sing. [DonutVpn]
- Merge. [DonutVpn]
- Merge. [DonutVpn]
- Merge pull request #17 from omid-the-great/main. [DonutVpn]

  add web_service for tunnel start/stop
- Add certification generator to cli. [Omid The Great]
- . [Omid The Great]
- Include tls fragment only to gRPC and WS. [DonutVpn]
- Remove extra settings that may create some issues. [DonutVpn]


v0.13.1 (2024-01-26)
--------------------
- Release: version 0.13.1. [DonutVpn]
- Fix bug in: fragment. [DonutVpn]


v0.13.0 (2024-01-26)
--------------------
- Release: version 0.13.0. [DonutVpn]
- Update to singbox1.8.4, fix vless http crash. [DonutVpn]


v0.12.8 (2024-01-26)
--------------------
- Release: version 0.12.8. [DonutVpn]


v0.12.7 (2024-01-26)
--------------------
- Release: version 0.12.7. [DonutVpn]


v0.12.6 (2024-01-26)
--------------------
- Release: version 0.12.6. [DonutVpn]


v0.12.5 (2024-01-26)
--------------------

New
~~~
- Add auto warp. [DonutVpn]
- Add auto warp option. [DonutVpn]

Other
~~~~~
- Release: version 0.12.5. [DonutVpn]


v0.12.2 (2024-01-26)
--------------------
- Release: version 0.12.2. [DonutVpn]
- Temporary fix: @consumer need resolution. [DonutVpn]


v0.12.1 (2024-01-25)
--------------------

New
~~~
- Add warp option. [DonutVpn]
- Add warp config. [DonutVpn]
- Add tls tricks only to tls configs. [DonutVpn]

Fix
~~~
- Bug. [DonutVpn]
- Bugs. [DonutVpn]
- Warp. [DonutVpn]

Other
~~~~~
- Release: version 0.12.1. [DonutVpn]
- Release: version 0.12.0. [DonutVpn]
- Speed up ping. [DonutVpn]
- Merge pull request #15 from 83hd4d/dev. [DonutVpn]

  Fix issues and Add features
- Update outbound.go. [DonutVpn]
- Fix service.go to work with singbox 1.8.2 Drop deprecated feautre
  (Clash API: cache_file and store_selected) Add feature (Cache File:
  path) Add new idle_timeout field for URLTest outbound Refactor
  outbound.go Add feature to omit TLSTricks and Fragment on VLESS
  Reality configs. [83hd4d]
- Update to singbox 1.8.2. [DonutVpn]
- Update to singbox 1.8.2. [DonutVpn]


v0.14.20 (2024-01-21)
---------------------
- Release: version 0.14.20. [DonutVpn]
- Update feature_request.yaml. [DonutVpn]
- Update bug_report.yaml. [DonutVpn]


v0.11.2 (2024-01-21)
--------------------

Fix
~~~
- Make. [DonutVpn]

Other
~~~~~
- Release: version 0.11.2. [DonutVpn]
- Fix json parser for files with comment. [problematicconsumer]
- Add mux. [problematicconsumer]
- Change builder. [problematicconsumer]
- Change outbound patch. [problematicconsumer]
- Merge branch 'main' [problematicconsumer]
- Revert changes for swift. [DonutVpn]


v0.14.11 (2024-01-20)
---------------------

New
~~~
- Add deb and rpm build. [DonutVpn]

Fix
~~~
- Versioning issue in ios singbox platform extension. [DonutVpn]
- Release bug. [DonutVpn]
- Release bug. [DonutVpn]
- Build linux. [DonutVpn]
- Stats bug in iOS. [DonutVpn]
- Some bugs in ios. [donutvpn]

Other
~~~~~
- Release: version 0.14.11. [DonutVpn]
- Merge pull request #351 from amirsaam/main. [DonutVpn]

  use universal bundleid from base.xconfig
- Use universal bundleid from base.xconfig. [Amir Mohammadi]
- Merge pull request #350 from amirsaam/main. [DonutVpn]

  ios log handler base functionality
- Log handler base functionality. [Amir Mohammadi]

  + fix makefile ios-temp-prepare upgrading pub
  + remove tvOS from libcore local spm
  +
- Add mux options. [problematicconsumer]
- Add mux. [problematicconsumer]
- Update release_message.md. [lymanjre]
- Update build.yml. [DonutVpn]
- Update changelog. [problematicconsumer]
- Update readme. [problematicconsumer]
- Fix ios stats. [problematicconsumer]
- Add reset tunnel option on ios. [problematicconsumer]
- Update build.yml. [DonutVpn]
- Fix release. [DonutVpn]
- Fix release. [DonutVpn]
- FIX: RELEASE BUG. [DonutVpn]
- Fix release bug in macos. [DonutVpn]
- Update build.yml. [DonutVpn]
- Update build.yml. [DonutVpn]
- Update build.yml. [DonutVpn]
- Update dependencies. [problematicconsumer]
- Fix not releasing linux packages. [DonutVpn]
- Merge pull request #347 from amirsaam/main. [DonutVpn]

  fix connection bug after bundle update
- Fix connection bug after bundle update. [Amir Mohammadi]
- Merge pull request #345 from amirsaam/main. [DonutVpn]

  update to latest itunes connect
- Update to latest itunes connect. [Amir Mohammadi]
- Revert changes for swift package. [DonutVpn]
- Merge pull request #344 from amirsaam/main. [DonutVpn]

  switch to local spm to load libcore
- Merge branch 'main' of https://github.com/amirsaam/donutvpn-next. [Amir
  Mohammadi]
- Merge branch 'donutvpn:main' into main. [Amir Mohammadi]
- Remove tvOS in singbox target. [Amir Mohammadi]
- Remove tvOS, previously added apple vision for ipad. [Amir Mohammadi]
- Handle xcode crash. [Amir Mohammadi]
- Switch to local spm. [Amir Mohammadi]
- Fix make ios. [DonutVpn]
- Release: version 0.14.1.dev. [DonutVpn]
- Bumpcore version to v0.11. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Merge pull request #343 from amirsaam/main. [DonutVpn]

  url scheme +  target dependency + general update
- Merge branch 'donutvpn:main' into main. [Amir Mohammadi]
- Add url scheme. [Amir Mohammadi]
- Update Package.resolved. [Amir Mohammadi]
- Target dependancy + pod update, package resolve. [Amir Mohammadi]
- Remove donutvpn://import/ from export option. [DonutVpn]
- Merge pull request #342 from amirsaam/main. [DonutVpn]

  getting libcore from spm, adding landscape mode
- Merge branch 'main' of https://github.com/amirsaam/donutvpn-next. [Amir
  Mohammadi]
- Delete pubspec.lock. [DonutVpn]
- Revert some changers. [Amir Mohammadi]

  This reverts some of commit dde7c2419a12e0f36976f4c0ff8e66d882ec62ac.
- Update pod, getting libcore from spm, adding landscape mode. [Amir
  Mohammadi]
- Change profile options modal. [problematicconsumer]
- Fix infinite sub expire date (#334) [DonutVpn, Pavel Volkov,
  problematicconsumer]

  * Fix infinite sub expire date

  * fix expire

  * fix build

  * refactor

  * make it better readable

  * Fix infinite sub

  * Add test for infinite sub

  ---------
- Fix fragment bugs. [problematicconsumer]
- Merge pull request #340 from amirsaam/main. [DonutVpn]

  fix ui not showing connected after relaunch
- Fix ui not showing connected after relaunch. [Amir Mohammadi]
- Merge pull request #338 from amirsaam/main. [DonutVpn]

  add push notification entitlement
- Merge branch 'main' of https://github.com/amirsaam/donutvpn-next. [Amir
  Mohammadi]
- Add config export on ios. [problematicconsumer]
- Fix ios connection status on app restart. [problematicconsumer]
- Add push notification entitlement. [Amir Mohammadi]
- Fix attempt. [problematicconsumer]
- Fix. [problematicconsumer]
- Fix ci again. [problematicconsumer]
- Fix ci tag matching. [problematicconsumer]
- Fix ci. [problematicconsumer]
- Release: version 0.14.0.dev. [problematicconsumer]
- Change logger to print in console. [problematicconsumer]
- Change routing assets order. [problematicconsumer]
- Change options description. [problematicconsumer]
- Fix ci tag pattern. [problematicconsumer]
- Fix ci. [problematicconsumer]
- Change workflows. [problematicconsumer]
- Add missing es translations. [problematicconsumer]
- Inlang: update translations. [AvatarStark]
- Fix infinite sub traffic. [problematicconsumer]
- Fix bugs. [problematicconsumer]
- Fix modal text alignment. [problematicconsumer]
- Merge pull request #329 from amirsaam/newMain. [DonutVpn]

  revert landscape mode
- Revert landscape mode. [Amir Mohammadi]
- Merge pull request #328 from amirsaam/main. [DonutVpn]

  remove unnecessary force unwrap, libcore linking add landscape mode
- Remove unnecessary force unwrap, libcore linking add landscape mode,
  [Amir Mohammadi]
- Change ios method error handling. [problematicconsumer]
- Fix bugs. [problematicconsumer]
- Remove custom_lint. [problematicconsumer]
- Add vscode launch config macos designed for ipad.
  [problematicconsumer]
- Fix ios flutter plugins. [problematicconsumer]
- Fix ios core name. [problematicconsumer]
- Merge pull request #323 from amirsaam/main. [DonutVpn]

  better fix for vpn connecting
- Better fix for vpn connecting. [Amir Mohammadi]
- Merge pull request #321 from amirsaam/main. [DonutVpn]

  match exportOptions identifiers with project
- Fixing not connecting to vpn. [Amir Mohammadi]
- Match exportOptions identifiers with project. [Amir Mohammadi]
- Merge pull request #319 from amirsaam/main. [DonutVpn]

  fix vpn profile not being created + changing run to release
- Fix vpn profile not being created + changing run to release. [Amir
  Mohammadi]
- Merge pull request #318 from amirsaam/main. [DonutVpn]

  ios fix?
- Ios fix? [Amir Mohammadi]
- Fix build. [problematicconsumer]
- Fix windows deep link. [problematicconsumer]
- Merge pull request #316 from amirsaam/main. [DonutVpn]

  Update pubspec.lock
- Update pubspec.lock. [Amir Mohammadi]
- Merge pull request #315 from amirsaam/main. [DonutVpn]

  add back cupertino_http
- Add back cupertino_http. [Amir Mohammadi]
- Update. [donutvpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [donutvpn]
- Update appcast. [problematicconsumer]


v0.11.1 (2024-01-17)
--------------------
- Change config parser. [problematicconsumer]
- Change gitignore. [problematicconsumer]
- Update Package.swift. [DonutVpn]


v0.11.0 (2024-01-17)
--------------------

New
~~~
- Add support for mux from sublink. [DonutVpn]
- Add sha256 to github action for ios. [DonutVpn]
- Update plist for ios. [DonutVpn]

Fix
~~~
- Disabled padding. [DonutVpn]
- Bug in checksum. [DonutVpn]
- Typo. [DonutVpn]
- Singbox fragment issue. [DonutVpn]
- Bug. [DonutVpn]

Other
~~~~~
- Release: version 0.11.0. [DonutVpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next-core.
  [DonutVpn]
- Merge pull request #14 from amirsaam/main. [DonutVpn]

  fix checksum
- Fix checksum. [Amir Mohammadi]
- Merge pull request #13 from amirsaam/main. [DonutVpn]

  fix typo in spm
- Update Package.swift. [Amir Mohammadi]
- Update. [DonutVpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next-core.
  [DonutVpn]
- Update Package.swift. [DonutVpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next-core.
  [DonutVpn]
- Merge pull request #12 from amirsaam/main. [DonutVpn]

  fix current spm
- Resolve. [Amir Mohammadi]
- Fix current spm. [Amir Mohammadi]
- Better manage configs. [DonutVpn]
- Force json as singbox config. [DonutVpn]
- Merge pull request #11 from amirsaam/main. [DonutVpn]

  Switch to SPM from Pod
- Switch to SPM from Pod. [Amir Mohammadi]
- Update release.yml. [DonutVpn]
- Merge pull request #10 from amirsaam/main. [DonutVpn]

  missed the "v" in generated link
- Update Libcore.podspec. [Amir Mohammadi]
- Merge pull request #9 from amirsaam/main. [DonutVpn]

  another try to fix podspec
- Update Libcore.podspec. [Amir Mohammadi]
- Update Libcore.podspec. [Amir Mohammadi]
- Merge pull request #8 from amirsaam/main. [DonutVpn]

  fix podspec?
- Update Libcore.podspec. [Amir Mohammadi]
- Fix podspec? [Amir Mohammadi]
- Add default padding mode. [problematicconsumer]
- Update Libcore.podspec. [DonutVpn]
- Fix tls tricks. [problematicconsumer]
- Change build cli. [problematicconsumer]
- Add parse cli. [problematicconsumer]
- Refactor shared. [problematicconsumer]
- Possible bubg resolving of fragment. [DonutVpn]
- Add check config cmd. [problematicconsumer]
- Add default config options. [problematicconsumer]
- Bump go version. [problematicconsumer]
- Update podspec. [DonutVpn]
- Update pod spec. [DonutVpn]
- Update podspec. [DonutVpn]
- Update podspec. [DonutVpn]
- Update podspec. [DonutVpn]


v0.13.6 (2024-01-07)
--------------------
- Release: version 0.13.6. [problematicconsumer]
- Fix build. [problematicconsumer]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [donutvpn]
- Release: version 0.13.5.dev. [problematicconsumer]
- Fix qr scanner links. [problematicconsumer]
- Update changelog. [problematicconsumer]
- Update core (sing-box v1.7.8) [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Remove unused dependencies. [problematicconsumer]
- Refactor service wrappers. [problematicconsumer]
- Update. [donutvpn]
- Update ios. [donutvpn]
- Change default connection test url. [problematicconsumer]
- Fix android service mode switch. [problematicconsumer]
- Add profile fetch cancel. [problematicconsumer]
- Release: version v0.13.4.dev. [problematicconsumer]
- Add update all subscriptions. [problematicconsumer]
- Fix profile update service. [problematicconsumer]
- Inlang: update translations. [huajizhige]
- Change app http client. [problematicconsumer]
- Release: version 0.13.3.dev. [problematicconsumer]
- Inlang: update translations. [wujiahao15]
- Fix bugs. [problematicconsumer]
- Add experimental flag to new config options. [problematicconsumer]
- Add connection from lan option. [problematicconsumer]
- Add dns routing option. [problematicconsumer]
- Add bypass lan option. [problematicconsumer]
- Fix profile edit with new url. [problematicconsumer]
- Change tls tricks settings section name. [problematicconsumer]
- Release: version 0.13.2.dev. [problematicconsumer]
- Fix db migration bug. [problematicconsumer]
- Bump version. [problematicconsumer]
- Release: version 0.13.1.dev. [problematicconsumer]
- Fix minor bugs. [problematicconsumer]
- Remove desktop auto connect. [problematicconsumer]
- Add experimental feature notice. [problematicconsumer]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Fix multi instance on windows. [problematicconsumer]
- Inlang: update translations. [huajizhige]
- Inlang: update translations. [huajizhige]
- Update README_cn.md. [wldjdjsks]
- Update README_cn.md. [wldjdjsks]

  Some proxy software such as clash do not have official Chinese names and use English by default.
- Add experimental flag in settings ui. [problematicconsumer]
- Fix inlang setup. [problematicconsumer]
- Release: version 0.13.0 dev. [problematicconsumer]
- Remove robocopy. [problematicconsumer]
- Fix build. [problematicconsumer]
- Fix windows portable again. [problematicconsumer]
- Update changelog. [problematicconsumer]
- Add android high refresh rate support. [problematicconsumer]
- Refactor desktop window management and tray. [problematicconsumer]
- Fix macOS restore from dock. [problematicconsumer]
- Add desktop shortcuts. [problematicconsumer]


v0.10.0 (2024-01-07)
--------------------

New
~~~
- Add podspec in make file. [DonutVpn]

Other
~~~~~
- Release: version 0.10.0. [problematicconsumer]
- Update release.yml. [DonutVpn]
- Update dependencies (sing-box v1.7.8) [problematicconsumer]
- Fix build. [problematicconsumer]
- Add allow connection from lan. [problematicconsumer]
- Add dns routing option. [problematicconsumer]
- Update release.yml. [DonutVpn]
- Create Libcore.podspec. [DonutVpn]


v0.12.3 (2023-12-28)
--------------------

Changes
~~~~~~~
- Fix inconsistent channel naming. [DonutVpn]

Fix
~~~
- Set codec for stream handlers. [donutvpn]
- Ios launch bug. [donutvpn]

Other
~~~~~
- Release: version 0.12.3. [problematicconsumer]
- Add Afghanistan region. [problematicconsumer]
- Fix packaging bug. [problematicconsumer]
- Fix windows packaging. [problematicconsumer]
- Fix bugs. [problematicconsumer]
- Change minimum macos version in docs. [problematicconsumer]
- Add version to desktop window. [problematicconsumer]
- [fix]: SettingsPickerDialog pop exception. [Ray]
- Update ios. [donutvpn]
- Ios: add /platform event channel. [donutvpn]
- Feat: Added stats stream for ios. [donutvpn]
- Feat: Added stats stream for ios. [donutvpn]
- Update ios. [donutvpn]
- Update ios. [donutvpn]
- Update ios. [donutvpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [donutvpn]
- Release: version 0.12.2. [problematicconsumer]
- Remove native http client (temporarily) [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update core (sing-box v1.7.6) [problematicconsumer]
- Fix bugs. [problematicconsumer]
- Ios update. [donutvpn]
- Update core (sing-box v1.7.5) [problematicconsumer]
- Change initialization and logging. [problematicconsumer]
- Release: version 0.12.1. [problematicconsumer]
- Fix preferences initialization error. [problematicconsumer]
- Bump android sdk version. [problematicconsumer]
- Fix android service mode. [problematicconsumer]
- Change privacy policy url. [problematicconsumer]
- Update appcast. [problematicconsumer]


v0.9.2 (2023-12-23)
-------------------
- Update sing-box version 1.7.6. [problematicconsumer]


v0.9.1 (2023-12-22)
-------------------
- Update singbox version 1.7.5. [problematicconsumer]
- Update gomobile. [problematicconsumer]


v0.12.0 (2023-12-20)
--------------------

New
~~~
- Add support for hysteria and wg. [DonutVpn]
- Add user-agent like clash sing-box for better compatibility. [DonutVpn]
- Send all releases to beta by default. [DonutVpn]
- Add ic_launcher for android tv. [DonutVpn]
- Publish dev release to seperate channel. [DonutVpn]
- Add winget. [DonutVpn]
- Add prepare to make file. [DonutVpn]
- Add ios core library to the project. [DonutVpn]

Fix
~~~
- Bug in  windows portable. [DonutVpn]
- Linux parse. [DonutVpn]
- Build-linux-libs. [Aloxaf]
- Build. [DonutVpn]
- Naming issue. [DonutVpn]
- Winget release. [DonutVpn]
- Action version. [DonutVpn]
- Winget publish. [DonutVpn]
- Build error. [DonutVpn]
- No commit message. [DonutVpn]
- Build. [DonutVpn]
- Bug. [DonutVpn]
- Build. [DonutVpn]
- Distutils. [DonutVpn]
- No commit message. [DonutVpn]
- Ipa. [DonutVpn]

Other
~~~~~
- Release: version 0.12.0. [problematicconsumer]
- Add donutvpn/iimport schema. [DonutVpn]
- Fix chinese inlang url. [problematicconsumer]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update changelog. [problematicconsumer]
- Change memory limit on desktop. [problematicconsumer]
- Remove locale name provider. [problematicconsumer]
- Fix service mode. [problematicconsumer]
- Add android dynamic notification. [problematicconsumer]
- Add android stats channel. [problematicconsumer]
- Fix translation generator. [problematicconsumer]
- Add config option reset. [problematicconsumer]
- Add missing routing asset warning. [problematicconsumer]
- Fix text input traversal. [problematicconsumer]
- Add basic d-pad support. [problematicconsumer]
- Fix tray initialization bug. [problematicconsumer]
- Update LICENSE.md. [DonutVpn]
- Update LICENSE.md. [DonutVpn]
- Update LICENSE.md. [DonutVpn]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README.md. [lymanjre]
- Add preferences migration. [problematicconsumer]
- Fix tray behavior. [problematicconsumer]
- Fix intro routing bug. [problematicconsumer]
- Fix macos build. [problematicconsumer]
- Fix macos name. [problematicconsumer]
- Fix macos icon. [problematicconsumer]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update ios. [donutvpn]
- Update ios. [donutvpn]
- Update ios. [donutvpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [donutvpn]
- Fix build. [problematicconsumer]
- Change entrypoint. [problematicconsumer]
- Add TLS trick config options. [problematicconsumer]
- Update ios. [donutvpn]
- Update platform interface. [donutvpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [donutvpn]
- Update dev-i.yml. [DonutVpn]
- Update dev-i.yml. [DonutVpn]
- Update pubspec.yaml. [DonutVpn]
- Create dev-i.yml. [DonutVpn]
- Update changelog. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update core (singbox v1.7) [problematicconsumer]
- Fix status mapper. [problematicconsumer]
- Refactor. [problematicconsumer]
- Refactor logs. [problematicconsumer]
- Remove unused clash api. [problematicconsumer]
- Update core (singbox 1.7.0-rc.2) [problematicconsumer]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Refactor router. [problematicconsumer]
- Refactor profiles. [problematicconsumer]
- Update changelog. [problematicconsumer]
- Fix chinese typography bug. [problematicconsumer]
- Add soffchen geo assets. [problematicconsumer]
- Refactor geo assets. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Bump ci sdk version. [problematicconsumer]
- Bump sdk version. [problematicconsumer]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update appcast. [problematicconsumer]
- Merge pull request #189 from jomertix/main. [DonutVpn]

  Update ru strings
- Nothing changed strings_tr.i18n.json. [jomertix]
- Nothing changed strings_fa.i18n.json. [jomertix]
- Inlang: update translations. [jomertix]
- Inlang: update translations. [jomertix]
- Release: version 0.11.1. [problematicconsumer]
- Fix android manifest. [problematicconsumer]
- Fix documentation. [problematicconsumer]
- Update release_message.md. [lymanjre]
- Release: version 0.11.0. [problematicconsumer]
- Change default direct dns. [problematicconsumer]
- Revert "new: add ic_launcher for android tv" [problematicconsumer]

  This reverts commit 49bb62c8289d0885a762e01b7f0785690a7d3af7.
- Change ir region rules. [problematicconsumer]
- Release: version 0.10.10.dev. [problematicconsumer]
- Merge pull request #179 from betaxab/main. [DonutVpn]

  Update translations
- Inlang: update translations. [betaxab]
- Add recommended geo assets. [problematicconsumer]
- Add error handling for geo assets. [problematicconsumer]
- Fix build. [problematicconsumer]
- Add geo assets settings. [problematicconsumer]
- Fix bug in strings_tr.i18n.json. [DonutVpn]
- Merge pull request #173 from hasankarli/main. [DonutVpn]

  Update Turkish translations
- Inlang: update translations. [hasankarli]
- Merge pull request #169 from solokot/main. [DonutVpn]

  Update Russian translation
- Merge branch 'main' into main. [DonutVpn]
- Merge pull request #168 from jomertix/main. [DonutVpn]

  inlang: update translations
- Inlang: update translations. [jomertix]
- Merge pull request #172 from Locas56227/main. [DonutVpn]

  Fix and improve Chinese README
- Fix and improve Chinese README. [Locas]
- Update build.yml. [DonutVpn]
- Update appcast. [problematicconsumer]
- Release: version 0.10.9.dev. [problematicconsumer]
- Add navigation to system tray. [problematicconsumer]
- Update Russian translation. [solokot]

  and partially revert ugly inlang changes
- Change http adapter. [problematicconsumer]
- Add independent dns cache option. [problematicconsumer]
- Fix bottom navigation bar accessibility. [problematicconsumer]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Merge pull request #165 from huajizhige/main. [DonutVpn]

  Improve translation
- Inlang: update translations. [huajizhige]
- Inlang: update translations. [huajizhige]
- Inlang: update translations. [huajizhige]
- Inlang: update translations. [huajizhige]
- Update appcast. [problematicconsumer]
- Release: version 0.10.8.dev. [problematicconsumer]
- Fix release command. [problematicconsumer]
- Add subscription qr code share. [problematicconsumer]
- Improve qr code scanner ux. [problematicconsumer]
- Fix code typo. [problematicconsumer]
- Ci: ignore appcast. [problematicconsumer]
- Merge pull request #164 from jomertix/main. [DonutVpn]

  Update ru string
- Update strings_ru.i18n.json. [jomertix]

  Fix translation
- Update strings_ru.i18n.json. [jomertix]

  Fix translate
- Merge pull request #162 from solokot/main. [DonutVpn]

  Update strings_ru
- Update strings_ru. [solokot]

  Fix google translate
- Merge pull request #161 from Aloxaf/Aloxaf-patch-1. [DonutVpn]

  fix: build-linux-libs
- Add sub link share. [problematicconsumer]
- Fix bootstrap bug. [problematicconsumer]
- Add proxy tag dialog. [problematicconsumer]
- Add export config to clipboard. [problematicconsumer]
- Update. [donutvpn]
- Add geoassessts to makefile prepare. [donutvpn]
- Merge pull request #155 from solokot/main. [DonutVpn]

  Update strings_ru
- Update strings_ru. [solokot]
- Add service mode and strict route. [problematicconsumer]
- Add basic privilege check. [problematicconsumer]
- Fix ui bugs. [problematicconsumer]
- Remove unnecessary preferences. [problematicconsumer]
- Fix config preferences. [problematicconsumer]
- Update changelog. [problematicconsumer]
- Update core to v0.8.0. [problematicconsumer]
- Fix manual update checker. [problematicconsumer]
- Fix logs page scroll. [problematicconsumer]
- Fix android service restart. [problematicconsumer]
- Fix android service revoke. [problematicconsumer]
- Change toast. [problematicconsumer]
- Fix linting and cleanup. [problematicconsumer]
- Remove execute config as is. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update README_cn.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Release: version 0.10.7.dev. [DonutVpn]
- Release: version 0.10.6.dev. [DonutVpn]
- Release: version 0.10.5.dev. [DonutVpn]
- Release: version 0.10.4.dev. [DonutVpn]
- Release: version 0.10.3.dev. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Merge pull request #138 from huajizhige/main. [DonutVpn]

  inlang: update translations
- Inlang: update translations. [huajizhige]
- Release: version 0.10.2.dev. [DonutVpn]
- Release: version 0.10.1.dev. [DonutVpn]
- Update. [DonutVpn]
- Fix. [DonutVpn]
- Fix. [DonutVpn]
- Test. [DonutVpn]
- Merge pull request #137 from huajizhige/patch-1. [DonutVpn]

  Update README_cn.md
- Update README_cn.md. [wldjdjsks]

  Improve translation
- Update README_cn.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Fix. [DonutVpn]
- Update. [DonutVpn]
- Fix. [DonutVpn]
- Fix build. [DonutVpn]
- Add distutils. [DonutVpn]
- Improve Android TV. [DonutVpn]
- Update. [DonutVpn]
- Add. [DonutVpn]
- Fix appcast url. [problematicconsumer]
- Add appcast. [problematicconsumer]
- Add auto connect on start. [problematicconsumer]
- Merge pull request #123 from Nyar233/main. [DonutVpn]

  Update translations
- Inlang: update translations. [Nyar233]
- Update release_message.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Merge pull request #115 from solokot/main. [DonutVpn]

  Update Russian: fix google translate
- Update Russian: fix google translate. [solokot]
- Merge pull request #113 from Nyar233/main. [DonutVpn]

  inlang: update translations
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Inlang: update translations. [Nyar233]
- Change system tray options. [problematicconsumer]
- Add core mode. [problematicconsumer]
- Add sidebar stats for large screens. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Change router for different screen size. [problematicconsumer]
- Merge pull request #108 from Hiiirad/patch-1. [DonutVpn]

  Update auto_translator.py to fix the Path Traversal Vulnerability
- Update auto_translator.py to fix the Path Traversal Vulnerability.
  [Hirad Rasoolinejad]
- Release: version 0.10.0. [problematicconsumer]
- Remove reconnect on auto profile update. [problematicconsumer]
- Update changelog. [problematicconsumer]
- Add selected tag for selector outbounds. [problematicconsumer]
- Add core debug flag. [problematicconsumer]
- Add url test delay color. [problematicconsumer]
- Change info logs. [problematicconsumer]
- Delete .github/help/linux/آموزش هیدیفای‌نکست فارسی لینوکس.desktop.
  [DonutVpn]
- Delete .github/help/mac-windows/آموزش هیدیفای‌نکست فارسی.url.
  [DonutVpn]
- Update dependencies. [problematicconsumer]
- Change pre release update checking off. [problematicconsumer]
- Add terms and privacy to about page. [problematicconsumer]
- Add memory limit option. [problematicconsumer]
- Change directory management. [problematicconsumer]
- Comment ios. [DonutVpn]
- Change app id. [DonutVpn]
- Merge pull request #98 from GFWFighter/main. [DonutVpn]

  Initial iOS version
- Merge pull request #1 from GFWFighter/ios. [GFWFighter]

  iOS
- Merge branch 'main' into ios. [GFWFighter]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Merge pull request #95 from solokot/main. [DonutVpn]

  Update Russian
- Update Russian. [solokot]

  Mainly fix google translate of  new worlds
- Change logging. [problematicconsumer]
- Change changelog workflow. [problematicconsumer]
- Add russia region. [problematicconsumer]
- Change theme prefs. [problematicconsumer]
- Merge pull request #90 from solokot/main. [DonutVpn]

  Russian translation: fix some mistakes
- Russian translation: fix some mistakes. [solokot]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Add rules. [problematicconsumer]
- Update LICENSE.md. [DonutVpn]
- Inlang: update translations. [donutvpn-com]
- Merge pull request #85 from leic4u/main. [DonutVpn]

  Update README_cn.md
- Update README_cn.md. [leic4u]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Add inlang translations. [DonutVpn]
- Add inlang project and remove localize. [DonutVpn]
- Merge pull request #82 from Iam54r1n4/main. [DonutVpn]

  add inlang translator
- Update: fr_FR.json. [Sarina]
- Add: french language. [Sarina]
- Add: inlang translation files(incomplete) [Sarina]
- Add: setup inlang translation. [Sarina]
- Merge pull request #74 from lifeindarkside/patch-1. [DonutVpn]

  Update strings_ru.i18n.json
- Update strings_ru.i18n.json. [Elshad Guseynov]

  Update some wrong interpretation
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README.md. [lymanjre]
- First working version. [GFWFighter]
- Underlying VPN Logic. [GFWFighter]
- Initialize Packet Tunnel + Config. [GFWFighter]
- Update Tutorial_For_DonutVpnNext_Linux.desktop. [lymanjre]
- Update Tutorial_For_DonutVpnNext.url. [lymanjre]
- Update آموزش هیدیفای‌نکست فارسی.url. [lymanjre]
- Update آموزش هیدیفای‌نکست فارسی لینوکس.desktop. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Release: version 0.9.2. [problematicconsumer]
- Fix ndk setup. [problematicconsumer]
- Fix android arm bug. [problematicconsumer]
- Release: version 0.9.1. [DonutVpn]
- Update README_cn.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Delete docs/file. [lymanjre]
- Google play badge. [lymanjre]
- Delete docs/google-play-badge1.png. [lymanjre]
- Upload google play badge. [lymanjre]
- Create file. [lymanjre]
- Delete docs/google-play-badge.png. [lymanjre]
- Fix ci. [problematicconsumer]
- Change ndk setup. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Change desktop error handling. [problematicconsumer]
- Fix android bugs. [problematicconsumer]
- Update README.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Delete Russian_Flag.png. [lymanjre]
- Add Russian flag. [lymanjre]
- Delete docs/file. [lymanjre]
- Create README_ru.md. [lymanjre]
- Delete REAMME_ru.md. [lymanjre]
- Create REAMME_ru.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_cn.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update file. [lymanjre]
- Uplaod a badge for google play. [lymanjre]
- Delete google-play-badge.png. [lymanjre]
- Add files via upload. [lymanjre]
- Create file. [lymanjre]
- Delete docs/google-play-badge.png. [lymanjre]
- Delete google-play-badge.png. [lymanjre]
- Add files via upload. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Create Chinese README_cn.md. [lymanjre]
- Release: version 0.9.0. [problematicconsumer]


v0.9.0 (2023-12-03)
-------------------
- Add fragment options. [problematicconsumer]
- Fix tls tricks parser. [problematicconsumer]
- Add basic cli. [problematicconsumer]
- Add initial tls tricks. [problematicconsumer]


v0.8.4 (2023-12-01)
-------------------
- Update singbox version 1.7. [problematicconsumer]


v0.8.3 (2023-11-27)
-------------------
- Update dependencies. [problematicconsumer]


v0.8.2 (2023-11-17)
-------------------

New
~~~
- Add turn server. [DonutVpn]

Other
~~~~~
- Add geo asset path. [problematicconsumer]


v0.8.1 (2023-11-14)
-------------------
- Add independent dns cache option. [problematicconsumer]
- Update sing-box to 1.7.0-beta.1. [problematicconsumer]
- Add config export. [problematicconsumer]
- Add strict route option. [problematicconsumer]


v0.8.0 (2023-11-09)
-------------------

Fix
~~~
- Shadow socks links and update to singbox1.7.0.dev11. [DonutVpn]


v0.7.0 (2023-10-27)
-------------------

New
~~~
- Add ios build. [DonutVpn]

Fix
~~~
- Ios release. [DonutVpn]

Other
~~~~~
- Add with_grpc build tag. [problematicconsumer]
- Fix service restart. [problematicconsumer]
- Add debug logs. [problematicconsumer]
- Add start delay. [problematicconsumer]
- Add memory limit. [problematicconsumer]
- Fix build. [problematicconsumer]
- Update to singbox 1.6.rc1. [DonutVpn]
- Merge pull request #2 from GFWFighter/fix-ios-name. [DonutVpn]

  Added PRODUCT_NAME to use with ios build
- Added PRODUCT_NAME to use with ios build. [GFWFighter]
- Change logs and formatting. [problematicconsumer]
- Change logging flags. [problematicconsumer]
- Add fake dns. [problematicconsumer]
- Add routing rules. [problematicconsumer]
- Change android build tags to include debug info. [problematicconsumer]


v0.6.3 (2023-10-15)
-------------------

Fix
~~~
- Issue on arm32. [DonutVpn]


v0.6.2 (2023-10-14)
-------------------
- Add panic recover to ffi. [problematicconsumer]


v0.8.12 (2023-10-13)
--------------------

Fix
~~~
- Typo. [DonutVpn]
- Bug. [DonutVpn]
- Release names. [DonutVpn]

Other
~~~~~
- Update readme. [problematicconsumer]
- Update core. [problematicconsumer]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Merge pull request #56 from solokot/main. [DonutVpn]

  Improvement of Russian translation
- Improvement of Russian translation. [solokot]

  Basically a replacement for machine automatic translation
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Fix bugs. [problematicconsumer]
- Change android signing. [problematicconsumer]
- Update README.md. [lymanjre]
- Update readme. [problematicconsumer]
- Update contribution guide. [problematicconsumer]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update README.md. [DonutVpn]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update release template. [DonutVpn]


v0.6.1 (2023-10-13)
-------------------
- Fix build. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update ray2sing. [DonutVpn]


v0.6.0 (2023-10-12)
-------------------

New
~~~
- Fix support for android 14. [DonutVpn]

Other
~~~~~
- Fix ndk setup action. [problematicconsumer]
- Change android ci. [problematicconsumer]
- Change android ci. [problematicconsumer]
- Bump NDK version. [problematicconsumer]


v0.5.4 (2023-10-11)
-------------------
- Fix parser. [problematicconsumer]


v0.8.11 (2023-10-08)
--------------------

Changes
~~~~~~~
- Remove auto release message. [DonutVpn]

Other
~~~~~
- Release: version 0.8.11. [DonutVpn]


v0.8.10 (2023-10-08)
--------------------

Fix
~~~
- Release changelog. [DonutVpn]

Other
~~~~~
- Release: version 0.8.10. [DonutVpn]


v0.8.9 (2023-10-08)
-------------------

Fix
~~~
- Missing libs. [DonutVpn]

Other
~~~~~
- Release: version 0.8.9. [DonutVpn]


v0.8.8 (2023-10-08)
-------------------

Fix
~~~
- Release bug. [DonutVpn]

Other
~~~~~
- Release: version 0.8.8. [DonutVpn]


v0.8.7 (2023-10-08)
-------------------

Fix
~~~
- Release message. [DonutVpn]

Other
~~~~~
- Release: version 0.8.7. [DonutVpn]


v0.8.6 (2023-10-08)
-------------------

Fix
~~~
- Windows build. [DonutVpn]
- Build issue. [DonutVpn]

Other
~~~~~
- Release: version 0.8.6. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Update release_message.md. [lymanjre]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Update release_message.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Delete docs/note. [lymanjre]
- Add files via upload. [lymanjre]
- Create note. [lymanjre]
- Update contribute.md. [lymanjre]
- Update README.md. [lymanjre]
- Delete assets/images/google-play-badge.png. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Fix build. [DonutVpn]


v0.8.5 (2023-10-07)
-------------------

New
~~~
- Add chinese lang. [DonutVpn]
- Add russian lang. [DonutVpn]

Fix
~~~
- Bug. [DonutVpn]
- Translate. [DonutVpn]
- Hysteria2 and some links. [DonutVpn]
- Chinese translation. [DonutVpn]

Other
~~~~~
- Release: version 0.8.5. [DonutVpn]
- Release: version 0.8.4. [DonutVpn]
- Release: version 0.8.3. [DonutVpn]
- Add release message  and help. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Fix bugs. [problematicconsumer]
- Update release_message.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Update release_message.md. [lymanjre]
- Create release_message.md. [lymanjre]
- Add google play badge to assets. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Create contribute.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Add feature request template. [problematicconsumer]
- Fix issue template. [problematicconsumer]
- Add issue template. [problematicconsumer]
- Release: version 0.8.2. [DonutVpn]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Create README_fa.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]
- Update README.md. [lymanjre]

  Add some features to the readme
- Add debug export to clipboard. [problematicconsumer]
- Release: version 0.8.1. [DonutVpn]
- Update core 0.5.1. [DonutVpn]
- Fix floating number sub info header. [problematicconsumer]
- Add russian. [DonutVpn]
- Add google play descriptions. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Release: version 0.8.0. [problematicconsumer]
- Add proxy tag sanitization. [problematicconsumer]
- Fix bugs. [problematicconsumer]
- Add ignore app update version. [problematicconsumer]
- Add new protocols to link parser. [problematicconsumer]
- Auto update translations on release. [DonutVpn]
- Update translation. [DonutVpn]
- Remove param in ru translation. [DonutVpn]


v0.5.2 (2023-10-07)
-------------------
- Updating ray2sing. [DonutVpn]


v0.5.1 (2023-10-06)
-------------------
- Update hysteria vless http shadowsocks. [DonutVpn]


v0.5.0 (2023-10-05)
-------------------

New
~~~
- Add hidden char. [DonutVpn]
- Add ssh and hysteria2 link support and fix vmess links. [DonutVpn]

Other
~~~~~
- Bump clash2singbox version. [problematicconsumer]
- Update to singbox 1.5.0. [DonutVpn]


v0.7.2 (2023-10-04)
-------------------
- Release: version 0.7.2. [problematicconsumer]
- Fix bugs. [problematicconsumer]


v0.7.1 (2023-10-03)
-------------------

New
~~~
- Add support of some exception panel with zero usage. [DonutVpn]

Other
~~~~~
- Release: version 0.7.1. [problematicconsumer]
- Fix log and analytics bugs. [problematicconsumer]
- Release: version 0.7.0. [problematicconsumer]
- Fix translation bug. [problematicconsumer]
- Improve error handling and presentation. [problematicconsumer]
- Add retry for network ops. [problematicconsumer]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Add local profile. [problematicconsumer]
- Add auto translator. [DonutVpn]
- Add auto translate. [DonutVpn]
- Release: version 0.6.0. [problematicconsumer]
- Fix minor bugs. [problematicconsumer]
- Add scheduled profile update. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Refactor profile details page. [problematicconsumer]


v0.5.11 (2023-09-22)
--------------------
- Release: version 0.5.11. [problematicconsumer]
- Fix android bundle abi. [problematicconsumer]
- Bump core version. [problematicconsumer]


v0.5.10 (2023-09-22)
--------------------
- Release: version 0.5.10. [problematicconsumer]
- Fix minor bugs. [problematicconsumer]
- Fix profile update bug. [problematicconsumer]


v0.4.1 (2023-09-22)
-------------------
- Fix minor bugs. [problematicconsumer]
- Update to latest donutvpn-singbox 1.5.0-rc2. [DonutVpn]


v0.5.9 (2023-09-22)
-------------------
- Release: version 0.5.9. [DonutVpn]
- Add build number. [DonutVpn]


v0.5.8 (2023-09-22)
-------------------

New
~~~
- Add crashlytics. [DonutVpn]
- Add support for base64 sublink for header from content. [DonutVpn]
- Add profile headers from comments in response! good for hosting in
  github and show information. [DonutVpn]
- Automated version release. [DonutVpn]
- Send release versions only to play market add pre-release version.
  [DonutVpn]
- Add signing to android. [DonutVpn]
- Add android universal build also. [DonutVpn]
- Add change in network entitlements. [DonutVpn]
- Change logo icon to next. [DonutVpn]
- Add cache for speed up build process. [DonutVpn]
- Add windows portable version. [DonutVpn]
- Add libclash. [DonutVpn]
- Add as draft release. [DonutVpn]
- Make better ci and building applications. [DonutVpn]

Changes
~~~~~~~
- Change invalid dns 235.5.5.5 to 8.8.8.8. [DonutVpn]
- Change windows logo. [DonutVpn]
- Change macos build  to flutter_distributor. [DonutVpn]
- Remove x86 builds since flutter does not support. [DonutVpn]
- Add x64 to the name. [DonutVpn]
- Update makefile. [DonutVpn]

Fix
~~~
- Improve routing accessibility and logs. [problematicconsumer]
- Minor bugs. [problematicconsumer]
- Prefs persistence. [problematicconsumer]
- Crashlytics. [problematicconsumer]
- App update url. [problematicconsumer]
- Small profiles. [DonutVpn]
- Makefile vars. [problematicconsumer]
- Adaptive icon. [DonutVpn]
- Pre-release. [DonutVpn]
- Typo in adaptive icon. [DonutVpn]
- If .dev is exist in the version do not show update needed. [DonutVpn]
- Keep the link as it is. fix the issue with & [DonutVpn]
- Dependency issue. [DonutVpn]
- Remove extra print. [DonutVpn]
- Bug in get headers from body. [DonutVpn]
- Bug ini ci to google play. [DonutVpn]
- Tag version issue. [DonutVpn]
- Ci bug. [DonutVpn]
- Remove comments. [DonutVpn]
- Bug. [DonutVpn]
- Space bugs in some panels. [DonutVpn]
- Aab. [DonutVpn]
- Aab build. [DonutVpn]
- Universal sign. [DonutVpn]
- KeystoreProperties. [DonutVpn]
- Name issue. [DonutVpn]
- Revert package name change. [DonutVpn]
- Name. [DonutVpn]
- Geosite download. [DonutVpn]
- Windows portable bug. [DonutVpn]
- Change name bug. [DonutVpn]
- Setup exe files. [DonutVpn]
- Libclash.so. [DonutVpn]
- Linux AppImage. [DonutVpn]
- Error. [DonutVpn]
- Category in linux. [DonutVpn]
- Add fuse for linux. [DonutVpn]
- Icon. [DonutVpn]
- Linux logo. [DonutVpn]
- Android. [DonutVpn]
- Makefile error. [DonutVpn]

Other
~~~~~
- Release: version 0.5.8. [DonutVpn]
- Fix ci. [problematicconsumer]
- Fix false-positive error reports. [problematicconsumer]
- Change build setup. [problematicconsumer]
- Release: version 0.5.7.dev ðŸš€ [problematicconsumer]
- Fix minor bugs. [problematicconsumer]
- Refactor app update. [problematicconsumer]
- Fix sentry dart plugin upload. [problematicconsumer]
- Fix ci debug symbols upload. [problematicconsumer]
- Add sentry provider observer. [problematicconsumer]
- Ci: add sentry debug info upload. [problematicconsumer]
- Update dependencies and general fixes. [problematicconsumer]
- Chore: bump agp version. [problematicconsumer]
- Ci: fix env. [problematicconsumer]
- Ci: add dsn env. [problematicconsumer]
- Feat: add region and terms to intro. [problematicconsumer]
- Update ci.yml. [DonutVpn]
- Build: add sentry dsn. [problematicconsumer]
- Feat: add intro screen. [problematicconsumer]
- Feat: add sentry. [problematicconsumer]
- Ci: bump macos version. [problematicconsumer]
- Feat: update profile when adding preexisting url.
  [problematicconsumer]
- Publish draft even with error. [DonutVpn]
- Update version of core. [DonutVpn]
- Merge branch 'main' of donutvpn-github:donutvpn/donutvpn-next. [DonutVpn]
- Add firebase. [DonutVpn]
- Update translation. [DonutVpn]
- Refactor: version presentation. [problematicconsumer]
- Perf: improve header parser. [problematicconsumer]
- Feat: remove check for updates in market releases.
  [problematicconsumer]
- Release: version 0.5.6.dev 🚀 [DonutVpn]
- Release: version 0.5.5.dev 🚀 [DonutVpn]
- Release: version 0.5.4.dev 🚀 [DonutVpn]
- Better manage the market release. [DonutVpn]
- Release: version 0.5.3.dev 🚀 [DonutVpn]
- Release: version 0.5.2.dev 🚀 [DonutVpn]
- Update. [DonutVpn]
- Release: version 0.5.1.dev 🚀 [DonutVpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [DonutVpn]
- Improve accessability. [problematicconsumer]
- Fix per-app proxy selection. [problematicconsumer]
- Add android per-app proxy. [problematicconsumer]
- Add basic flavors. [problematicconsumer]
- Update Makefile. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update Makefile. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update README.md. [DonutVpn]
- Add accessability semantics. [problematicconsumer]
- Release v0.1.0. [problematicconsumer]
- Fix ci build. [problematicconsumer]
- Update ci. [problematicconsumer]
- Add core version. [problematicconsumer]
- Fix barrel file. [problematicconsumer]
- Change proxies lifecycle. [problematicconsumer]
- Add service restart. [problematicconsumer]
- Remove notification service. [problematicconsumer]
- Change android notification permission. [problematicconsumer]
- Add android service restart. [problematicconsumer]
- Change mark new profile active. [problematicconsumer]
- Add support for fragment in the url. [DonutVpn]
- Handle unlimited. [DonutVpn]
- Update upload download link stats. [DonutVpn]
- Update Translations. [DonutVpn]
- Remove // TODO add content disposition parsing. [DonutVpn]
- Add content-disposition for profile title. [DonutVpn]
- Update README.md. [DonutVpn]
- Add hysteria2. [problematicconsumer]
- Fix android build connection. [problematicconsumer]
- Add android connection shortcut. [problematicconsumer]
- Add desktop autostart. [problematicconsumer]
- Add android boot receiver. [problematicconsumer]
- Add android proxy service. [problematicconsumer]
- Add android battery optimizations settings. [problematicconsumer]
- Change sharedpreferences to unify with android. [problematicconsumer]
- Add vclibs. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add submodule. [problematicconsumer]
- Fix translation code gen. [problematicconsumer]
- Change prefs. [problematicconsumer]
- Change default config options. [problematicconsumer]
- Remove string casing. [problematicconsumer]
- Update ci.yml. [DonutVpn]
- Remove caching. [DonutVpn]
- Change core prefs to use code generation. [problematicconsumer]
- Fix custom lint. [problematicconsumer]
- Fix general issues. [problematicconsumer]
- Remove vclibs. [problematicconsumer]
- Update README.md. [DonutVpn]
- Fix blank screen. [problematicconsumer]
- Add proxies sort. [problematicconsumer]
- Refactor preferences. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add more config options to settings. [problematicconsumer]
- Add tun implementation option. [problematicconsumer]
- Add android power manager. [problematicconsumer]
- Add android quick tile. [problematicconsumer]
- Fix macos dependencies. [problematicconsumer]
- Fix mac build. [problematicconsumer]
- Fix build. [problematicconsumer]
- Change system tray icon. [problematicconsumer]
- Fix riverpod code generation. [problematicconsumer]
- Remove unnecessary config options. [problematicconsumer]
- Add accessability semantics. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add config options. [problematicconsumer]
- Add pref utilities. [problematicconsumer]
- Add stats overview. [problematicconsumer]
- Update Translations. [DonutVpn]
- Fix android outbounds view. [problematicconsumer]
- Remove unnecessary options. [problematicconsumer]
- Change proxies flow. [problematicconsumer]
- Add android command client support. [problematicconsumer]
- Update README.md. [DonutVpn]
- Add status command receiver. [problematicconsumer]
- Add donutvpn deeplink. [problematicconsumer]
- Add macos deeplink support. [problematicconsumer]
- Add singbox deeplink. [problematicconsumer]
- Change error prompts. [problematicconsumer]
- Fix url parser. [problematicconsumer]
- Remove unnecessary prefs. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Make AppImage zipped for preserving permission in linux. [DonutVpn]
- Add user agent. [problematicconsumer]
- Fix uri launch. [problematicconsumer]
- Update ci flutter version. [problematicconsumer]
- Fix macos silent start. [problematicconsumer]
- Change proxies page. [problematicconsumer]
- Fix desktop connection error msg. [problematicconsumer]
- Add button tooltips. [problematicconsumer]
- Fix logging. [problematicconsumer]
- Create dependabot.yml. [DonutVpn]
- Create CODE_OF_CONDUCT.md. [DonutVpn]
- Fix aab. [DonutVpn]
- Add aab file. [DonutVpn]
- Change windows-portable name to DonutVpnNext-portable. [DonutVpn]
- Add debug mode. [problematicconsumer]
- Add misc settings ui. [problematicconsumer]
- Remote unnecessary logs. [problematicconsumer]
- Add misc preferences. [problematicconsumer]
- Add directory options. [problematicconsumer]
- Fix linux build. [problematicconsumer]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [problematicconsumer]
- Fix android builds. [DonutVpn]
- Fix android splash screen. [problematicconsumer]
- Update icons. [problematicconsumer]
- Fix windows executable build. [problematicconsumer]
- Remove unnecessary build steps. [problematicconsumer]
- Remove build number from appbar. [problematicconsumer]
- Update readme. [problematicconsumer]
- Change desktop directories. [problematicconsumer]
- Remove desktop notifications. [problematicconsumer]
- Force same version code for all platforms. [DonutVpn]
- Add more logs. [problematicconsumer]
- Add log timestamp. [problematicconsumer]
- Update icons. [problematicconsumer]
- Remove drawer branding. [problematicconsumer]
- Add download section in readme. [DonutVpn]
- Fix name of universal apk. [DonutVpn]
- Fix order. [DonutVpn]
- Update readme. [DonutVpn]
- Update logging. [problematicconsumer]
- Update geo assets url. [problematicconsumer]
- Change linux directories. [problematicconsumer]
- Update icon. [DonutVpn]
- Update icon. [DonutVpn]
- Update logo for all platforms. [DonutVpn]
- Revert name in macos. [DonutVpn]
- Temporary disable app sandbox. [DonutVpn]
- Change name. [DonutVpn]
- Update Release.entitlements to fix binding issue. [DonutVpn]
- Merge pull request #5 from evstegneych/main. [DonutVpn]

  add macos option
- Add macos option. [EvstegneyCh]
- Update ci flutter version. [problematicconsumer]
- Update kotlin version. [problematicconsumer]
- Update other dependencies. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update flutter version. [problematicconsumer]
- Update LICENSE.md. [DonutVpn]
- Add build option for ios but not tested. [DonutVpn]
- Return: build for all. [DonutVpn]
- Fix build for macos. [DonutVpn]
- Fix ci. [problematicconsumer]
- Migrate to singbox. [problematicconsumer]
- Fix routing. [problematicconsumer]
- Fix connection button text casing. [problematicconsumer]
- Add update checking. [problematicconsumer]
- Add separate page for clash overrides. [problematicconsumer]
- Add version number to appbar. [problematicconsumer]
- Add more icons. [problematicconsumer]
- Add sort limited profiles last. [problematicconsumer]
- Fix profile sort icon. [problematicconsumer]
- Fix profile traffic ratio. [problematicconsumer]
- Add profiles sort option. [problematicconsumer]
- Refactor profile addition flow. [problematicconsumer]
- Add extra profile metadata. [problematicconsumer]
- Fix linux deep link service. [problematicconsumer]
- Refactor profile tile. [problematicconsumer]
- Add locale based font. [problematicconsumer]
- Fix linux startup. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add about page. [problematicconsumer]
- Fix linux notifications. [problematicconsumer]
- Build all. [DonutVpn]
- Hard coding. [DonutVpn]
- Add png. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update Makefile. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update Makefile. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update Makefile. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Update Makefile. [DonutVpn]
- Fix makefile. [problematicconsumer]
- Fix actions. [problematicconsumer]
- Fix actions. [problematicconsumer]
- Fix actions. [problematicconsumer]
- Update build setup. [problematicconsumer]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [problematicconsumer]
- Create LICENSE.md. [DonutVpn]
- Add core submodule. [problematicconsumer]
- Remove core libs. [problematicconsumer]
- Update Makefile. [DonutVpn]
- Update Makefile. [DonutVpn]
- Remove c install. [DonutVpn]
- Update make. [DonutVpn]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [DonutVpn]
- Add Persian font. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Update to libclash. [DonutVpn]
- Add cache. [DonutVpn]
- Add: new action. [DonutVpn]
- Add write permission. [DonutVpn]
- Fix. [DonutVpn]
- Modify. [DonutVpn]
- Fix. [DonutVpn]
- Update. [DonutVpn]
- Fix android build. [problematicconsumer]
- Fix code gen bug. [problematicconsumer]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Fix CI. [problematicconsumer]
- Fix CI. [problematicconsumer]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [problematicconsumer]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Fix CI. [problematicconsumer]
- Update ci.yml. [DonutVpn]
- Update ci.yml. [DonutVpn]
- Add CI. [DonutVpn]
- Add silent start for desktop. [problematicconsumer]
- Update readme. [problematicconsumer]
- Add headless mode to desktop. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add distribution setup for windows. [problematicconsumer]
- Add Farsi(fa) language. [problematicconsumer]
- Merge branch 'main' of https://github.com/donutvpn/donutvpn-next.
  [problematicconsumer]
- Initial commit. [DonutVpn]
- Initial. [problematicconsumer]


v0.3.0 (2023-09-16)
-------------------
- Update to latest version of singbox. [DonutVpn]


v0.2.0 (2023-09-11)
-------------------
- Update to newer version of ray2sing. [DonutVpn]


v0.1.1 (2023-09-11)
-------------------
- Release v0.1.1. [problematicconsumer]
- Remove version number. [problematicconsumer]


v0.1.0 (2023-09-11)
-------------------

New
~~~
- Add v2ray parser. [DonutVpn]
- Add dependency for make android. [DonutVpn]
- Add cache to go. [DonutVpn]
- Add go cache for speed up build. [DonutVpn]
- Add ios build. [DonutVpn]
- Fix all archs. [DonutVpn]
- Add macos, x86 build. [DonutVpn]

Changes
~~~~~~~
- Make ios only target to speed up process. [DonutVpn]
- Remove x86 build due to incompatible with flutter. [DonutVpn]

Fix
~~~
- Bug in v2ray to clash. [DonutVpn]
- Remove useless route. [DonutVpn]
- Tree not found. [DonutVpn]
- Issues. [DonutVpn]

Other
~~~~~
- Release v0.1.0. [problematicconsumer]
- Update ci. [problematicconsumer]
- Update ci. [problematicconsumer]
- Update build. [problematicconsumer]
- Update ci. [problematicconsumer]
- Remove cache. [DonutVpn]
- Add service restart. [problematicconsumer]
- Update to singbox to skip fatal error on outbound error. [DonutVpn]
- Fix confilicte with clash. [DonutVpn]
- Add ray2singbox. [DonutVpn]
- Update to singbox v1.5.0.6. [DonutVpn]
- Add tun stack option. [problematicconsumer]
- Update dependencies. [problematicconsumer]
- Add config builder. [problematicconsumer]
- Add groups command client. [problematicconsumer]
- Add status command client. [problematicconsumer]
- Add command server. [problematicconsumer]
- Add dart api bridge. [problematicconsumer]
- Add box service. [problematicconsumer]
- Update. [DonutVpn]
- Update to singbox 1.4.0.2. [DonutVpn]
- Remove: extra command for android. [DonutVpn]
- Update go to 1.21. [DonutVpn]
- Update singbox to version 1.4.0. [DonutVpn]
- Update to singbox v1.4.0 from donutvpn repository. [DonutVpn]
- Fix override logic. [problematicconsumer]
- Fix clash template. [problematicconsumer]
- Update LICENSE.md. [DonutVpn]
- Remove macos-arm64 amd64. [DonutVpn]
- Remove macos seperate build. [DonutVpn]
- Disable ios build. [DonutVpn]
- Add CI. [problematicconsumer]
- Update config flow. [problematicconsumer]
- Add config override. [problematicconsumer]
- Add generalized config parser. [problematicconsumer]
- Bootstrap. [problematicconsumer]
- Create LICENSE.md. [DonutVpn]
- Initial commit. [DonutVpn]


