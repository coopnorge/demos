FROM openjdk@sha256:d50104b0b5811b60c198a3e5ced23c783a49e0d4f78b99762fca4bfdbbc59dd2

COPY ldlib/sapjco3.jar .
COPY target/demo-0.0.1-SNAPSHOT.jar .
COPY demo.jcoDestination .
COPY ldlib/libsapjco3.so .

ENV LD_LIBRARY_PATH=.

ENTRYPOINT ["java", "-jar", "-Dloader.path=.", "demo-0.0.1-SNAPSHOT.jar"]


