#!/bin/bash

ECAgentName=<ecagent-os-sys>
ECClientID=<client-id>
ECServerID=<server-id>
ECGatewayURL=<gateway-url>
ECZoneID=<ec-zone-id>
ECDebug=<debug>

UAAURL=<uaa-url>
UAAClientID=<uaa-client-id>
UAAClientSecret=<uaa-client-secret>

LocalPort=<local-port>
LocalProxy=<proxy>

./$ECAgentName -mod client -aid $ECClientID -tid $ECServerID -hst $ECGatewayURL/agent -lpt $LocalPort  -oa2 $UAAURL/oauth/token -cid $UAAClientID -csc $UAAClientSecret -dur 300 -zon $ECZoneID $LocalProxy $ECDebug