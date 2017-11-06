@echo off

set ECAgentName=<ecagent-os-sys>
set ECGatewayURL=<gateway-url>
set ECServerID=<server-id>
set ECClientID=<client-id>
set ECZoneID=<ec-zone-id>
set ECDebug=<debug>

set UAAURL=<uaa-url>
set UAAClientID=<uaa-client-id>
set UAAClientSecret=<uaa-client-secret>

set LocalPort=<local-port>
set LocalProxy=<proxy>

@echo on
@echo %ECAgentName% -mod client -aid %ECClientID% -tid %ECServerID% -hst %ECGatewayURL%/agent -lpt %LocalPort% -oa2 %UAAURL% -cid %UAAClientID% -csc %UAAClientSecret%  -dur 300 -zon %ECZoneID% %LocalProxy% %ECDebug%
@echo off
%ECAgentName% -mod client -aid %ECClientID% -tid %ECServerID% -hst %ECGatewayURL%/agent -lpt %LocalPort% -oa2 %UAAURL% -cid %UAAClientID% -csc %UAAClientSecret%  -dur 300 -zon %ECZoneID% %LocalProxy% %ECDebug%
