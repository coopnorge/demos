package com.example.demo;

import com.sap.conn.jco.*;

public class PurchaseOrderService {

    public String getPurchaseOrder(String purchaseOrder) throws JCoException {
        // JCoDestination is the logic address of an ABAP system and ...
        JCoDestination destination = JCoDestinationManager.getDestination("demo");
        // ... it always has a reference to a metadata repository
        JCoFunction function = destination.getRepository().getFunction("BAPI_PO_GETDETAIL1");
        if (function == null)
            throw new RuntimeException("BAPI_PO_GETDETAIL1 not found in SAP.");

        // JCoFunction is the container for function values. Each function consists of separate
        // containers for import, export, changing and table parameters.
        // To set or get the parameters use the APIS setXXX() and getXXX().
        function.getImportParameterList().setValue("PURCHASEORDER", purchaseOrder);

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

        JCoStructure poheader = function.getExportParameterList().getStructure("POHEADER");
        String poNumber = poheader.getString("PO_NUMBER");
        String status = poheader.getString("STATUS");
        return poNumber + " : " + status;
    }
}

