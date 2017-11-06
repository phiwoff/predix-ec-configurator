#!/bin/bash

ECAgentName=<ecagent-os-sys>
ECServerID=<server-id>
ECGatewayURL=<gateway-url>
ECZoneID=<ec-zone-id>
ECServiceURI=<ec-cf-service-url>

UAAURL=<uaa-url>
UAAClientID=<uaa-client-id>
UAAClientSecret=<uaa-client-secret>

ResourceHost=<resource-url>
ResourcePort=<resource-port>

./$ECAgentName -mod server -aid $ECServerID -hst $ECGatewayURL/agent -rht $ResourceHost -rpt $ResourcePort  -oa2 $UAAURL/oauth/token -cid $UAAClientID -csc $UAAClientSecret -dur 300 -hca ${PORT} -zon $ECZoneID -sst $ECServiceURI