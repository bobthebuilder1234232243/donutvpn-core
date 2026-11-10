#!/usr/bin/env bash
# Shared by the OpenWrt package builders: the payload every package carries, and the package
# metadata, in one place so the .ipk (24.10 and earlier, opkg) and the .apk (25.12 and later, apk)
# never drift apart. Sourced, not run.

set -e -o pipefail

PKG_NAME="donutvpn-core"
PKG_DESCRIPTION="Hiddify universal proxy platform: a sing-box core with Hiddify's protocols and subscription support."
PKG_LICENSE="GPL-3.0-or-later"
PKG_URL="https://donutvpn.com/"
PKG_MAINTAINER="Hiddify <wrt@donutvpn.com>"
PKG_SECTION="net"
# kmod-tun for TUN mode, kmod-inet-diag for the connection tracker, ca-bundle for TLS, firewall4
# and kmod-nft-queue for the nftables rules the core installs in TUN mode (as sing-box's package).
PKG_DEPENDS="ca-bundle kmod-inet-diag kmod-tun firewall4 kmod-nft-queue"
PKG_CONFFILES="/etc/config/donutvpn-core /etc/donutvpn-core/config.json"

# The repository root, from this file's place.
PROJECT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WRT_FILES="$PROJECT/platform/wrt/files"

# usage prints the builders' common usage and exits.
usage() {
  echo "Usage: $0 <openwrt architecture> <version> <binary path> <output path or directory>" >&2
  exit 1
}

# sha256 prints the SHA-256 of a file, with the coreutils or the BSD tool.
sha256() {
  if command -v sha256sum >/dev/null; then sha256sum "$1" | cut -d' ' -f1; else shasum -a 256 "$1" | cut -d' ' -f1; fi
}

# stage_payload <root> installs everything a package carries under root, as it lands on the router.
stage_payload() {
  local root="$1"
  install -d -m755 "$root/usr/bin" "$root/etc/donutvpn-core" "$root/etc/config" "$root/etc/init.d" "$root/lib/upgrade/keep.d" "$root/usr/share/licenses/$PKG_NAME"
  install -m755 "$BINARY_PATH" "$root/usr/bin/$PKG_NAME"
  install -m644 "$PROJECT/release/config/config.json" "$root/etc/donutvpn-core/config.json"
  install -m644 "$WRT_FILES/donutvpn-core.conf" "$root/etc/config/donutvpn-core"
  install -m755 "$WRT_FILES/donutvpn-core.init" "$root/etc/init.d/donutvpn-core"
  install -m644 "$WRT_FILES/donutvpn-core.keep" "$root/lib/upgrade/keep.d/donutvpn-core"
  install -m644 "$PROJECT/LICENSE.md" "$root/usr/share/licenses/$PKG_NAME/LICENSE.md"
}

# tar_root <root> <output.tar.gz> packs a staged root as OpenWrt expects: ./ paths, root owned,
# no timestamps that change between builds. GNU tar in CI, BSD tar on a Mac.
tar_root() {
  local root="$1" out="$2"
  if tar --version 2>/dev/null | grep -q GNU; then
    tar -C "$root" --owner=0 --group=0 --numeric-owner --sort=name --mtime="@${SOURCE_DATE_EPOCH:-0}" -cf - . | gzip -9n > "$out"
  else
    tar -C "$root" --uid 0 --gid 0 --uname root --gname root -cf - . | gzip -9n > "$out"
  fi
}
