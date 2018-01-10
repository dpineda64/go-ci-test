#!/usr/bin/env groovy

node('docker'){
    String applicationName = "go-ci-test"
    String buildNumber = "0.1.${env.BUILD_NUMBER}"
    String goPath = "go/src/github.com/dpineda64/${applicationName}"

    stage('Checkout from Github'){
        checkout scm
    }

    stage("Run") {
        docker.image("golang:1.8.0-alpine").inside("-v ${pwd()}:${goPath}") {
            for (command in binaryBuildCommands){
                sh "cd ${goPath} && go run main.go"
            }
        }
    }
}