package com.example.demo;

import com.sap.conn.jco.*;

public class HelloService {

    public String sayHello(String message) throws JCoException {
        // JCoDestination is the logic address of an ABAP system and ...
        JCoDestination destination = JCoDestinationManager.getDestination("demo");
        // ... it always has a reference to a metadata repository
        JCoFunction function = destination.getRepository().getFunction("STFC_CONNECTION");
        if (function == null)
            throw new RuntimeException("STFC_CONNECTION not found in SAP.");

        // JCoFunction is the container for function values. Each function consists of separate
        // containers for import, export, changing and table parameters.
        // To set or get the parameters use the APIS setXXX() and getXXX().
        function.getImportParameterList().setValue("REQUTEXT", message);

        try {
            // execute, i.e. send the function representation to the ABAP system addressed
            // by the specified destination, invoke it there, and retrieve the function
            // result sent back by the system
            // All necessary conversions between Java and ABAP data types are done automatically.
            function.execute(destination);
        } catch (AbapException e) {
            System.out.println(e);
            return e.getMessage();
        }

        System.out.println("STFC_CONNECTION finished:");
        System.out.println(" Echo: " + function.getExportParameterList().getString("ECHOTEXT"));
        System.out.println(" Response: " + function.getExportParameterList().getString("RESPTEXT"));
        System.out.println();
        return function.getExportParameterList().getString("RESPTEXT");
    }
}

