package com.example.demo;

import com.sap.conn.jco.*;

public class QueryService {

    public String listRfc() throws JCoException {
        // JCoDestination is the logic address of an ABAP system and ...
        JCoDestination destination = JCoDestinationManager.getDestination("demo");
        // ... it always has a reference to a metadata repository
        JCoFunction function = destination.getRepository().getFunction("SWO_QUERY_API_METHODS");
        if (function == null)
            throw new RuntimeException("SWO_QUERY_API_METHODS not found in SAP.");

        // JCoFunction is the container for function values. Each function consists of separate
        // containers for import, export, changing and table parameters.
        // To set or get the parameters use the APIS setXXX() and getXXX().
        //function.getImportParameterList().setValue("FUNCNAME", "BAPI_*");

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

        StringBuilder builder = new StringBuilder();
        JCoTable api_methods = function.getTableParameterList().getTable("API_METHODS");

        do {
            JCoField field = api_methods.getField("FUNCTION");
            builder.append(field.getString());
            builder.append("</br>");
        } while(api_methods.nextRow());

        return builder.toString();
    }
}

