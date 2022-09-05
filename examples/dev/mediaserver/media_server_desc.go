// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"encoding/xml"
)

const mediaServerVerOneDeviceDescription = xml.Header +
	"<device>" +
	"    <deviceType>urn:schemas-upnp-org:device:MediaServer:1</deviceType>" +
	"    <serviceList>" +
	"        <service>" +
	"            <serviceType>urn:schemas-upnp-org:service:ContentDirectory:1</serviceType>" +
	"			<serviceId>ContentDirectory</serviceId>" +
	"        </service>" +
	"        <service>" +
	"            <serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>" +
	"			<serviceId>ConnectionManager</serviceId>" +
	"        </service>" +
	"    </serviceList>" +
	"</device>"

const connectionManagerOneServiceDescription = xml.Header +
	"<scpd>" +
	"    <serviceStateTable>" +
	"        <stateVariable>" +
	"            <name>SourceProtocolInfo</name>" +
	"            <sendEventsAttribute>yes</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>SinkProtocolInfo</name>" +
	"            <sendEventsAttribute>yes</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>CurrentConnectionIDs</name>" +
	"            <sendEventsAttribute>yes</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_ConnectionStatus</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"            <allowedValueList>" +
	"                <allowedValue>OK</allowedValue>" +
	"                <allowedValue>ContentFormatMismatch</allowedValue>" +
	"                <allowedValue>InsufficientBandwidth</allowedValue>" +
	"                <allowedValue>UnreliableChannel</allowedValue>" +
	"                <allowedValue>Unknown</allowedValue>" +
	"            </allowedValueList>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_ConnectionManager</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_Direction</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"            <allowedValueList>" +
	"                <allowedValue>Input</allowedValue>" +
	"                <allowedValue>Output</allowedValue>" +
	"            </allowedValueList>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_ProtocolInfo</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>string</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_ConnectionID</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>i4</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_AVTransportID</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>i4</dataType>" +
	"        </stateVariable>" +
	"        <stateVariable>" +
	"            <name>A_ARG_TYPE_RcsID</name>" +
	"            <sendEventsAttribute>no</sendEventsAttribute>" +
	"            <dataType>i4</dataType>" +
	"        </stateVariable>" +
	"    </serviceStateTable>" +
	"<actionList>" +
	"        <action>" +
	"            <name>GetProtocolInfo</name>" +
	"            <argumentList>" +
	"                <argument>" +
	"                    <name>Source</name>" +
	"                    <direction>out</direction>               <relatedStateVariable>SourceProtocolInfo</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>Sink</name>" +
	"                    <direction>out</direction>               <relatedStateVariable>SinkProtocolInfo</relatedStateVariable>" +
	"                </argument>" +
	"            </argumentList>" +
	"        </action>" +
	"	  <action>" +
	"	  <Optional/>" +
	"            <name>PrepareForConnection</name>" +
	"            <argumentList>" +
	"                <argument>" +
	"                    <name>RemoteProtocolInfo</name>" +
	"                    <direction>in</direction>                    <relatedStateVariable>A_ARG_TYPE_ProtocolInfo</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>PeerConnectionManager</name>" +
	"                    <direction>in</direction>                  <relatedStateVariable>A_ARG_TYPE_ConnectionManager</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>PeerConnectionID</name>" +
	"                    <direction>in</direction>                    <relatedStateVariable>A_ARG_TYPE_ConnectionID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>Direction</name>" +
	"                    <direction>in</direction>                    <relatedStateVariable>A_ARG_TYPE_Direction</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>ConnectionID</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_ConnectionID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>AVTransportID</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_AVTransportID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>RcsID</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_RcsID</relatedStateVariable>" +
	"                </argument>" +
	"            </argumentList>" +
	"        </action>" +
	"        <action>" +
	"	  <Optional/>" +
	"            <name>ConnectionComplete</name>" +
	"            <argumentList>" +
	"                <argument>" +
	"                    <name>ConnectionID</name>" +
	"                    <direction>in</direction>                    <relatedStateVariable>A_ARG_TYPE_ConnectionID</relatedStateVariable>" +
	"                </argument>" +
	"            </argumentList>" +
	"        </action>" +
	"        <action>" +
	"            <name>GetCurrentConnectionIDs</name>" +
	"            <argumentList>" +
	"                <argument>" +
	"                    <name>ConnectionIDs</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>CurrentConnectionIDs</relatedStateVariable>" +
	"                </argument>" +
	"            </argumentList>" +
	"        </action>" +
	"        <action>" +
	"            <name>GetCurrentConnectionInfo</name>" +
	"            <argumentList>" +
	"                <argument>" +
	"                    <name>ConnectionID</name>" +
	"                    <direction>in</direction>                   <relatedStateVariable>A_ARG_TYPE_ConnectionID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>RcsID</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_RcsID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>AVTransportID</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_AVTransportID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>ProtocolInfo</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_ProtocolInfo</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>PeerConnectionManager</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_ConnectionManager</relatedStateVariable>" +
	"                </argument>" +
	"             <argument>" +
	"                    <name>PeerConnectionID</name>" +
	"                    <direction>out</direction>                   <relatedStateVariable>A_ARG_TYPE_ConnectionID</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>Direction</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_Direction</relatedStateVariable>" +
	"                </argument>" +
	"                <argument>" +
	"                    <name>Status</name>" +
	"                    <direction>out</direction>                    <relatedStateVariable>A_ARG_TYPE_ConnectionStatus</relatedStateVariable>" +
	"                </argument>" +
	"            </argumentList>" +
	"        </action>" +
	"    </actionList>" +
	"</scpd>"

const contentDirectoryOneServiceDescription = xml.Header +
	"<scpd>" +
	"  <serviceStateTable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>TransferIDs</name>" +
	"<sendEventsAttribute>yes</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_ObjectID</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_Result</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_SearchCriteria</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_BrowseFlag</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"<allowedValueList>" +
	"        <allowedValue>BrowseMetadata</allowedValue>" +
	"        <allowedValue>BrowseDirectChildren</allowedValue>" +
	"      </allowedValueList>" +
	"    </stateVariable>" +
	"    <stateVariable> " +
	"      <name>A_ARG_TYPE_Filter</name>" +
	"<sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_SortCriteria</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_Index</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>ui4</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_Count</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>ui4</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>A_ARG_TYPE_UpdateID</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>ui4</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_TransferID</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>ui4</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_TransferStatus</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"      <allowedValueList>" +
	"        <allowedValue>COMPLETED</allowedValue>" +
	"        <allowedValue>ERROR</allowedValue>" +
	"        <allowedValue>IN_PROGRESS</allowedValue>" +
	"        <allowedValue>STOPPED</allowedValue>" +
	"      </allowedValueList>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_TransferLength</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_TransferTotal</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_TagValueList</name> <sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>A_ARG_TYPE_URI</name> " +
	"<sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>uri</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>SearchCapabilities</name>" +
	"<sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>SortCapabilities</name>" +
	"<sendEventsAttribute>no</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>SystemUpdateID</name>" +
	"<sendEventsAttribute>yes</sendEventsAttribute>" +
	"      <dataType>ui4</dataType>" +
	"    </stateVariable>" +
	"    <stateVariable> <Optional/>" +
	"      <name>ContainerUpdateIDs</name>" +
	"<sendEventsAttribute>yes</sendEventsAttribute>" +
	"      <dataType>string</dataType>" +
	"    </stateVariable>" +
	"  </serviceStateTable>" +
	"  <actionList>" +
	"    <action>" +
	"    <name>GetSearchCapabilities</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>SearchCaps</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>SearchCapabilities </relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetSortCapabilities</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>SortCaps</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>SortCapabilities</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetSystemUpdateID</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>Id</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>SystemUpdateID</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>Browse</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ObjectID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>BrowseFlag</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_BrowseFlag</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Filter</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Filter</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>StartingIndex</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Index</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>RequestedCount</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>SortCriteria</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_SortCriteria</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Result</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Result</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>NumberReturned</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TotalMatches</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>UpdateID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_UpdateID</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>Search</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ContainerID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>SearchCriteria</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_SearchCriteria </relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Filter</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Filter</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>StartingIndex</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Index</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>RequestedCount</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>SortCriteria</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_SortCriteria</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Result</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Result</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>NumberReturned</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TotalMatches</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Count</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>UpdateID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_UpdateID</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>CreateObject</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ContainerID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Elements</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Result</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>ObjectID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>Result</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_Result</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>DestroyObject</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ObjectID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>UpdateObject</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ObjectID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>CurrentTagValue</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TagValueList </relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>NewTagValue</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TagValueList </relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>ImportResource</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>SourceURI</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_URI</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>DestinationURI</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_URI</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TransferID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferID </relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>ExportResource</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>SourceURI</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_URI</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>DestinationURI</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_URI</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TransferID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferID </relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>StopTransferResource</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>TransferID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferID </relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>GetTransferProgress</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>TransferID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferID </relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TransferStatus</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferStatus </relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TransferLength</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferLength </relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>TransferTotal</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_TransferTotal</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>DeleteResource</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ResourceURI</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_URI</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action><Optional/>" +
	"    <name>CreateReference</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ContainerID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"        <argument>" +
	"          <name>ObjectID</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"  	<argument>" +
	"          <name>NewID</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>A_ARG_TYPE_ObjectID</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"  </actionList>" +
	"</scpd>"
