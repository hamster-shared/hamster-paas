FROM ubuntu:latest
COPY ./build/hamster-paas /hamster-paas
EXPOSE 9898
CMD [ "/hamster-paas" ]