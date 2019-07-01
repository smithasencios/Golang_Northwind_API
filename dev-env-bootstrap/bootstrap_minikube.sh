#!/usr/bin/env bash

cluster_status=$(minikube status --format {{.APIServer}})

if [[ "$cluster_status" == "Running" ]]; then
  echo "Minikube running"
else
  echo "You might need this: "
  echo "minikube config set vm-driver kvm2"
  echo "minikube start --memory 4024 --cpus 4"
  echo "helm init"
  exit 1
fi

if [[ -f "mysql-external-service.yaml" ]]; then      
  
	# helm install --name api-db-mysql --namespace mysql stable/mysql --set mysqlRootPassword=lfda,mysqlUser=demouser,mysqlPassword=userpassword,mysqlDatabase=northwind
	# kubectl apply -f mysql-external-service.yaml -n mysql
  
  helm install --name api-db-mariadb --namespace mariadb stable/mariadb --set rootUser.password=lfda,db.user=judas,db.password=lfda,db.name=northwind,replication.enabled=false
  kubectl apply -f mariadb-external-service.yaml -n mariadb
  kubectl apply -f dev-northwind-api-external-service.yaml
	echo "Check env status with `kubectl get all` for ready "
  echo "Create the tables in northwind database: mysql -u root -p --host $(minikube ip) --port 30002 < database_creation_script.sql"
  echo "Create the tables in northwind database: mysql -u root -p --host $(minikube ip) --port 30002 < database_data_script.sql"
  
  echo "run skaffold dev commands"
	exit 0
else
  echo "Bootstrap script must be run from within same folder."
  exit 2
fi