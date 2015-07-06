// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const SwitchPowerServiceDescription = "" +
	"<scpd>" +
	"  <serviceStateTable>" +
	"    <stateVariable>" +
	"      <name>Target</name>" +
	"      <sendEventsAttribute>no</sendEventsAttribute> " +
	"      <dataType>boolean</dataType>" +
	"      <defaultValue>0</defaultValue>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>Status</name>" +
	"      <dataType>boolean</dataType>" +
	"      <defaultValue>0</defaultValue>" +
	"    </stateVariable>" +
	"  </serviceStateTable>" +
	"  <actionList>" +
	"    <action>" +
	"    <name>SetTarget</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>newTargetValue</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>Target</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetTarget</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>RetTargetValue</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>Target</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetStatus</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ResultStatus</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>Status</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"  </actionList>" +
	"</scpd>"