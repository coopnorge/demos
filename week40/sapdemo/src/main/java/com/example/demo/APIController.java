package com.example.demo;

import com.sap.conn.jco.JCoException;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class APIController {

	@GetMapping("/hello")
	public String hello(@RequestParam(value = "name", defaultValue = "World") String name) throws JCoException {
		return new HelloService().sayHello(name);
	}

	@GetMapping("/po")
	public String getPurchaseOrder() throws JCoException {
		return new PurchaseOrderService().getPurchaseOrder("0031604590");
	}

	@GetMapping("/rfc")
	public String getBapiFunctions() throws JCoException {
		return new QueryService().listRfc();
	}
}