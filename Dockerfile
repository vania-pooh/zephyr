FROM scratch
MAINTAINER Ivan Krutov <vania-pooh@aerokube.com>

COPY zephyr /

ENTRYPOINT ["/zephyr", "-config", "/zephyr.json"]
