FROM ubuntu:latest
COPY ./build/hamster-paas /root/hamster-paas
EXPOSE 9898
CMD [ "/root/hamster-paas" ]