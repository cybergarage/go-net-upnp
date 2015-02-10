/******************************************************************
 *
 * mUPnP for C
 *
 * Copyright (C) The go-net-upnp Authors 2005
 * Copyright (C) 2006 Nokia Corporation. All rights reserved.
 *
 * This is licensed under BSD-style license, see file COPYING.
 *
 ******************************************************************/

#include <mupnp/soap/soap.h>
#include <mupnp/util/log.h>

/****************************************
* mupnp_soap_createenvelopebodynode
****************************************/

mUpnpXmlNode *mupnp_soap_createenvelopebodynode()
{
	mUpnpXmlNode *envNode;
	mUpnpXmlNode *bodyNode;
	
	mupnp_log_debug_l4("Entering...\n");

	envNode = mupnp_xml_node_new();
	mupnp_xml_node_setname(envNode, CG_SOAP_XMLNS CG_SOAP_DELIM CG_SOAP_ENVELOPE);
	mupnp_xml_node_setattribute(envNode, CG_SOAP_ATTRIBUTE_XMLNS CG_SOAP_DELIM CG_SOAP_XMLNS, CG_SOAP_XMLNS_URL);
	mupnp_xml_node_setattribute(envNode, CG_SOAP_XMLNS CG_SOAP_DELIM CG_SOAP_ENCORDING, CG_SOAP_ENCSTYLE_URL);

	bodyNode = mupnp_xml_node_new();
	mupnp_xml_node_setname(bodyNode, CG_SOAP_XMLNS CG_SOAP_DELIM CG_SOAP_BODY);
	mupnp_xml_node_addchildnode(envNode, bodyNode);

	return envNode;

	mupnp_log_debug_l4("Leaving...\n");
}
