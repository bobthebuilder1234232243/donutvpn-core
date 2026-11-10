FROM alpine:latest
ENV CONFIG='https://raw.githubusercontent.com/ircfspace/warpsub/main/export/warp#WARP%20(IRCF)'
ENV VERSION=v5.5.5
WORKDIR /donutvpn
RUN apk add --no-cache curl tar gzip # iptables ip6tables

# The release's Linux builds: the musl one where there is one (amd64, arm64, 386, armv7), statically
# linked for Alpine; the others are statically linked as they are. Each archive holds a folder
# (donutvpn-core-linux-<build>/) with the donutvpn-core binary; DonutVpnCli, its old name, links to it
# for a compose file that still runs that.
RUN echo "architecture: $(apk --print-arch)" && \
    case "$(apk --print-arch)" in \
            x86_64) BUILD=amd64-musl ;; \
            x86) BUILD=386-musl ;; \
            aarch64) BUILD=arm64-musl ;; \
            armv7) BUILD=armv7-musl ;; \
            armhf) BUILD=armv6 ;; \
            s390x) BUILD=s390x ;; \
            *) echo "Unsupported architecture: $(apk --print-arch) $(uname -m)" && exit 1 ;; \
        esac && \
        URL="https://github.com/bobthebuilder1234232243/donutvpn-core/releases/download/${VERSION}/donutvpn-core-linux-$BUILD.tar.gz" && \
        echo "Downloading $URL" && \
        curl -fL -o donutvpn-core.tar.gz "$URL" && \
        tar -xzf donutvpn-core.tar.gz --strip-components=1 && rm donutvpn-core.tar.gz && \
        chmod +x donutvpn-core && \
        ln -s donutvpn-core DonutVpnCli
COPY donutvpn.sh .
RUN chmod +x donutvpn.sh

EXPOSE 12334
EXPOSE 12335
EXPOSE 12336
EXPOSE 12337
EXPOSE 12338
EXPOSE 16756


ENTRYPOINT [ "/donutvpn/donutvpn.sh" ]
