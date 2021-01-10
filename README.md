# Task challange

You own a coffee shop. You procure coffee beans from a number of suppliers and have them delivered directly to you. You source 3 types of coffee beans: Arabica, Robusta and Liberica.
Each supplier supplies one or more types of beans.
Drivers work for a carrier, and each carrier can haul one or more types of beans.

Write a SQL query that produces the invalid deliveries. (Invalid deliveries are deliveries that a carrier cannot perform due to carrier bean constraints)

Write a simple Golang server that has an endpoint that will return the results.

You will also need to provide the manifests for deploying the server onto a Kubernetes cluster. (e.g. deployment, service, ingress)

## Requirements

You will need to deploy your project using Docker/Kubernetes

## Bonus

- Use gRPC for the server and provide a postman endpoint. Please provide the proto files.
- Given that each delivery contains a random bean from the supplier's stock, what is the probability that the delivery is valid. Write an endpoint that will return this result.

## Task completion status

- ### I have written a SQL query that produces the invalid deliveries, here it is

```sql
WITH all_compination AS (
		SELECT suppliers.id AS supplier_id,
			   drivers.id   AS driver_id
		FROM drivers
				 CROSS JOIN suppliers),
		 valid_delivery AS (
			 SELECT DISTINCT suppliers.id AS supplier_id,
							 drivers.id   AS driver_id
			 FROM bean_types
					  LEFT JOIN carrier_bean_types ON bean_types.id = carrier_bean_types.bean_type_id
					  LEFT JOIN carriers ON carriers.id = carrier_bean_types.carrier_id
					  LEFT JOIN drivers ON drivers.carrier_id = carriers.id
					  LEFT JOIN supplier_bean_types ON supplier_bean_types.bean_type_id = carrier_bean_types.bean_type_id
					  LEFT JOIN suppliers ON suppliers.id = supplier_bean_types.supplier_id AND drivers.id IS NOT NULL
		 )
	SELECT all_compination.*
	FROM all_compination
			 LEFT JOIN valid_delivery ON valid_delivery.supplier_id = all_compination.supplier_id AND
										 valid_delivery.driver_id = all_compination.driver_id
	WHERE valid_delivery.supplier_id IS NULL
	  AND valid_delivery.driver_id IS NULL
```

- ### I have written a Golang gRPC server with REST proxies that you can build and deploy running the below commands

## Dependencies

- [protoc compiler](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [minikube](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-minikube/)
- docker
- kubectl

## Building

Strating minikube

```bash
start minikube
```

Switch to Docker daemon inside Minikube VM:

```bash
eval $(minikube docker-env)
```

Build images and deploy to Kubernetes cluster:

```bash
./build.sh
```

Try it out.

```bash
minikube service api-service --url
```

```bash
curl <output_url_from_previous_command>/deliveries/invalid_combination
```

Example

```bash
curl --request GET 'http://127.0.0.1:51650/deliveries/invalid_combination'
```
