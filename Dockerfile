FROM scratch

COPY infosvc /

ENTRYPOINT ["/infosvc"]
