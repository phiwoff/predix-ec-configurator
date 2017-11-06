#!/bin/bash

ECAgentName=<ecagent-os-sys>
ECClientID=<client-id>
ECServerID=<server-id>
ECGatewayURL=<gateway-url>

UAAURL=<uaa-url>
UAAClientID=<uaa-client-id>
UAAClientSecret=<uaa-client-secret>

LocalPort=<local-port>

./$ECAgentName -mod client -aid $ECClientID -tid $ECServerID -hst $ECGatewayURL/agent -lpt $LocalPort  -oa2 $UAAURL/oauth/token -cid $UAAClientID -csc $UAAClientSecret -dur 300