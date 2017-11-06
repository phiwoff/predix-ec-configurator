#!/bin/bash

ECAgentName=ecagent_linux_sys
ECZoneID=<ec-zone-id>
ECServiceURI=<ec-cf-service-url>
ECAdminToken=<ec-admin-token-from-vcap>

./$ECAgentName -mod gateway -lpt ${PORT} -zon $ECZoneID -sst $ECServiceURI -tkn $ECAdminToken