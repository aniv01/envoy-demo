FROM centos
COPY demo-app /app/bin/demo-app
RUN chmod +x /app/bin/demo-app
CMD ["/bin/bash", "-c","/app/bin/demo-app"]
