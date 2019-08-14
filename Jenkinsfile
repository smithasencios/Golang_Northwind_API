pipeline {

  environment {
    PROJECT = "northwind-project"
    APP_NAME = "northwindapi"
    CLUSTER = "northwind"
    CLUSTER_ZONE = "us-central1-f"
    IMAGE_TAG = "gcr.io/${PROJECT}/${APP_NAME}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
    JENKINS_CRED = "${PROJECT}"
  }

  agent {
    kubernetes {
      label 'northwind-api'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: cd-jenkins
  containers:
  - name: alpine
    image: alpine:latest
    command:
    - cat
    tty: true
  - name: gcloud
    image: gcr.io/cloud-builders/gcloud
    command:
    - cat
    tty: true
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
    command:
    - cat
    tty: true
"""
}
  }
  stages {
      stage('Build and push image with Container Builder') {
        steps {
          container('gcloud') {
            sh "PYTHONUNBUFFERED=1 gcloud builds submit -t ${IMAGE_TAG} ."
          }
        }
      }

      stage('Deploy Production') {
        // Production branch
        when { branch 'master' }
        steps{
          container('kubectl') {
            sh("kubectl get ns production || kubectl create ns production")
            sh("sed -i.bak 's#gcr.io/cloud-solutions-images/northwindapi:1.0.0#${IMAGE_TAG}#' ./k8s/production/*.yaml")
            step([$class: 'KubernetesEngineBuilder',namespace:'production', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/production', credentialsId: env.JENKINS_CRED, verifyDeployments: true])                    
          }
        }
      }

      stage('Deploy Dev') {
        // Developer Branches
        when { branch 'development' }
        steps {
          container('kubectl') {
            // Create namespace if it doesn't exist
            sh("kubectl get ns ${env.BRANCH_NAME} || kubectl create ns ${env.BRANCH_NAME}")
            // Don't use public load balancing for development branches
            sh("sed -i.bak 's#LoadBalancer#ClusterIP#' ./k8s/services/frontend.yaml")
            sh("sed -i.bak 's#gcr.io/cloud-solutions-images/northwindapi:1.0.0#${IMAGE_TAG}#' ./k8s/dev/*.yaml")
            step([$class: 'KubernetesEngineBuilder',namespace:${env.BRANCH_NAME}, projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/dev', credentialsId: env.JENKINS_CRED, verifyDeployments: true])
            step([$class: 'KubernetesEngineBuilder',namespace:${env.BRANCH_NAME}, projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/services', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
          }
        }
  }
}