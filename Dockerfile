FROM scratch
MAINTAINER Ivan Krutov <vania-pooh@aerokube.com>

COPY zephyr /usr/bin

ENTRYPOINT ["/usr/bin/zephyr", "-config", "/etc/zephyr/zephyr.json"]
