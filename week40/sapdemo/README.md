# SAP Java Connector Demo

This demo shows how to invoke a BAPI function on SAP ERP.

# Getting started

You will need to have installed the following

- [Java 17](https://jdk.java.net/17/)
- [Maven](https://maven.apache.org/)
- [Docker](https://www.docker.com/)

You might like to use [sdkman](https://sdkman.io/) to manage the maven and java environments

# How to run

You can run the app as a standalone jar file by running the following commands (run.sh)

```
mvn clean package
java -jar -Dloader.path=ldlib target/demo-0.0.1-SNAPSHOT.jar 
```

or in a docker container (run-docker.sh)

```
docker build -t demo .
docker run -p 8080:8080 demo:latest
```

# Notes 

## Packaging the sapjco3.jar file for a SpringBoot app

It was not possible to install the sapjco3.jar file in a maven repository because that renames the jar
file ( appends a hyphen and version number).  The code in the sapjco3.jar file is able to detect the 
name of the enclosing jar file and will throw an exception if it has been renamed.
```
java.lang.ExceptionInInitializerError: JCo initialization failed with java.lang.ExceptionInInitializerError: Illegal JCo archive "sapjco3-3.1.4.jar". It is not allowed to rename or repackage the original archive "sapjco3.jar".
	at com.sap.conn.jco.rt.Middleware.<clinit>(Middleware.java:87) ~[sapjco3-3.1.4.jar!/:20210330 1216 [3.1.4 (2021-03-30)]]
	at com.sap.conn.jco.rt.JCoRuntime.setMiddlewarePropertyValue(JCoRuntime.java:1726) ~[sapjco3-3.1.4.jar!/:20210330 1216 [3.1.4 (2021-03-30)]]
	at com.sap.conn.jco.rt.DefaultJCoRuntime.initialize(DefaultJCoRuntime.java:94) ~[sapjco3-3.1.4.jar!/:20210330 1216 [3.1.4 (2021-03-30)]]
	at com.sap.conn.jco.rt.JCoRuntimeFactory.<clinit>(JCoRuntimeFactory.java:23) ~[sapjco3-3.1.4.jar!/:20210330 1216 [3.1.4 (2021-03-30)]]
	at java.base/java.lang.Class.forName0(Native Method) ~[na:na]
```

It is not possible to just copy the sapjco3.jar file to the BOOTINF/lib directory as this jar file is compressed.
The following exception will be thrown if you do this : 

```
Caused by: java.lang.IllegalStateException: Unable to open nested entry 'BOOT-INF/lib/sapjco3.jar'. It has been compressed and nested jar files must be stored without compression. Please check the mechanism used to create your executable jar file

```