FROM centos:latest

COPY agent agent

CMD ["./agent"]

