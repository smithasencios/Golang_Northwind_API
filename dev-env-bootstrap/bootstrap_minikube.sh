#!/usr/bin/env bash

cluster_status=$(minikube status --format {{.APIServer}})

if [[ "$cluster_status" == "Running" ]]; then
  echo "Minikube running"
else
  echo "You might need this: "
  echo "minikube start --memory 4024 --cpus 4"
  echo "helm init"
  exit 1
fi

if [[ -f "mysql-external-service.yaml" ]]; then    
	helm install --name api-db --namespace mysql --set mysqlRootPassword=apipassword,mysqlUser=demouser,mysqlPassword=userpassword,mysqlDatabase=northwind stable/mysql  
	kubectl apply -f mysql-external-service.yaml -n mysql

	echo "Check env status with `kubectl get all` for ready and then run skaffold dev commands"
	exit 0
else
  echo "Bootstrap script must be run from within same folder."
  exit 2
fi