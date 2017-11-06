## Predix-EC-Configurator - Scenario 2: Step-by-Step Doc

**_This is an auto-generated file based on your inputs_. You can find it and all other auto-generated scripts saved into the `output` folder**

Here is a step-to-step document to setup EC for the selected scenario.

### 1. Diego-Enabler

Cloud Foundry uses the Diego architecture to manage application containers. Diego components assume application scheduling and management responsibility from the Cloud Controller.

Enable Diego support for an app running on Cloud Foundry.

```sh
cf add-plugin-repo CF-Community https://plugins.cloudfoundry.org/
cf install-plugin Diego-Enabler -r CF-Community
```

### 2. EC Agent Gateway

Here is the content for `ec-gateway.sh` file

```sh
<gateway_script_content_here>
```

Here is the content for `manifest.yml` file

```sh
<gateway_manifest_content_here>
```

#### Deploy the Agent Gateway to the Predix cloud

It is time now to push the EC Agent Gateway app to Predix.io

```sh
cf login // or predix login if you use Predix CLI
cd output/gateway
cf push
```

Enable Diego support:

```sh
cf enable-diego <ecagent_gateway_name>
```

Now it is time to map CF Route to the Gateway app with

```sh
cf map-route <ecagent_gateway_name> run.aws-usw02-pr.ice.predix.io -n <ecagent_gateway_name>
```

and start the EC Agent Gateway

```sh
cf start <ecagent_gateway_name>
```

Check if it works opening a browser windows at `https://<ecagent_gateway_name>.run.aws-usw02-pr.ice.predix.io/health`

### 3. EC Agent Server

Here is the content for `ec-server.sh` file

```sh
<server_script_content_here>
```

Here is the content for `manifest.yml` file

```sh
<server_manifest_content_here>
```

#### Deploy the EC Agent Server to the Predix cloud

It is time now to push the EC Agent Server app to Predix.io

```sh
cf login // or predix login if you use Predix CLI
cd output/server
cf push
```

Enable Diego support:

```sh
cf enable-diego <ecagent_server_name>
```

Now, it is time to map CF Route to the Gateway app

```sh
cf map-route <ecagent_server_name> run.aws-usw02-pr.ice.predix.io -n <ecagent_server_name>
```

and start the EC Agent Server

```sh
cf start <ecagent_server_name>
```

Check if it works opening a browser windows at `https://<ecagent_server_name>.run.aws-usw02-pr.ice.predix.io/health`

**NOTE:** Verify the server appears as "SupperConns" belongs to the gateway: ``https://<ecagent_gateway_name>.run.aws-usw02-pr.ice.predix.io/health` (it may take a minute)

### 4. EC Agent Client

Here is the content for `ec-client` file

```sh
<client_script_content_here>
```

### 5. Connect to the Predix data source from you local machine

You should now be able to use a local client for Predix resource and connect to it.

E.g. If you want to connect to PostgreSQL on Predix, you could download and install on your local machine [PGAdmin](https://www.pgadmin.org/) and create a new server configuration as below:
 
- Hostname: localhost
- Port: <local_port>
- Maintenance Database: **postgresql-database-name-from-cf-vcap**
- Username: **postgresql-user-from-cf-vcap**
- Password: **postgresql-password-from-cf-vcap**
