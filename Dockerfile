FROM scratch
LABEL org.opencontainers.image.source=https://github.com/FrangipaneTeam/terraform-analytic-tool
ENTRYPOINT ["/terraform-analytic-tool"]
COPY terraform-analytic-tool /