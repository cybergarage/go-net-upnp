// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"testing"
)

func TestNewService(t *testing.T) {
	NewService()
}

func TestServiceLoadDescription(t *testing.T) {
	srv := NewService()

	err := srv.LoadDescriptionString(TestServiceContentDirectory1)
	if err != nil {
		t.Error(err)
	}

	// Check actionList

	for n, action := range srv.Description.ActionList.Actions {
		var expectedActionName string
		switch n {
		case 0:
			expectedActionName = "GetSearchCapabilities"
		case 1:
			expectedActionName = "GetSortCapabilities"
		case 2:
			expectedActionName = "GetSystemUpdateID"
		}
		if 2 < n {
			break
		}
		if action.Name != expectedActionName {
			t.Errorf("action name = %s, expected %s", action.Name, expectedActionName)
		}
	}

	// Check ServiceStateTable

	for n, stateVar := range srv.Description.ServiceStateTable.StateVariables {
		var expectedStatName string
		var expectedStatData string
		switch n {
		case 0:
			expectedStatName = "TransferIDs"
			expectedStatData = "string"
		case 1:
			expectedStatName = "A_ARG_TYPE_ObjectID"
			expectedStatData = "string"
		case 2:
			expectedStatName = "A_ARG_TYPE_Result"
			expectedStatData = "string"
		}
		if 2 < n {
			break
		}
		if stateVar.Name != expectedStatName {
			t.Errorf("state variable name = %s, expected %s", stateVar.Name, expectedStatName)
		}
		if stateVar.DataType != expectedStatData {
			t.Errorf("state variable data = %s, expected %s", stateVar.DataType, expectedStatData)
		}
	}
}

const TestServiceContentDirectory1 = xml.Header +
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
