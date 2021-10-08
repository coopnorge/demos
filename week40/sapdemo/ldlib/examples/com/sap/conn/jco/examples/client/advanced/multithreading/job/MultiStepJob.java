package com.sap.conn.jco.examples.client.advanced.multithreading.job;

public interface MultiStepJob
{

    boolean isFinished();

    void runNextStep();

    String getName();

    void cleanUp();
}
