FROM openjdk:17

COPY sapjco3.jar .
COPY target/demo-0.0.1-SNAPSHOT.jar .
COPY demo.jcoDestination .
COPY ldlib/libsapjco3.so .

ENV LD_LIBRARY_PATH=.

ENTRYPOINT ["java", "-jar", "-Dloader.path=.", "demo-0.0.1-SNAPSHOT.jar"]


