@echo off
set ECAgentName=<ecagent-os-sys>
set ECServerID=<server-id>
set ECGatewayURL=<gateway-url>
set ECZoneID=<ec-zone-id>
set ECServiceURI=<ec-cf-service-url>
set ECDebug=<debug>

set UAAURL=<uaa-url>
set UAAClientID=<uaa-client-id>
set UAAClientSecret=<uaa-client-secret>

set ResourceHost=<amazon-postgres-url, schema excluded>
set ResourcePort=<amazon-postgres-port>

set LocalProxy=<proxy>
@echo on
@echo %ECAgentName% -mod server -aid %ECServerID% -hst %ECGatewayURL%/agent -rht %ResourceHost% -rpt %ResourcePort%  -oa2 %UAAURL%/oauth/token -cid %UAAClientID% -csc %UAAClientSecret% -dur 300 -hca 7990 -zon %ECZoneID% -sst %ECServiceURI% %LocalProxy% %ECDebug%
@echo off

%ECAgentName% -mod server -aid %ECServerID% -hst %ECGatewayURL%/agent -rht %ResourceHost% -rpt %ResourcePort%  -oa2 %UAAURL%/oauth/token -cid %UAAClientID% -csc %UAAClientSecret% -dur 300 -hca 7990 -zon %ECZoneID% -sst %ECServiceURI% %LocalProxy% %ECDebug%