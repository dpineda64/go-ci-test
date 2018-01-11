pipeline{
  agent{
    docker{
      image 'golang:1.8.0-alpine'
    }
  }
  stages{
    stage('change to repository'){
      steps {
        checkout scm
      }
    }
    stage('install deps'){
        steps {
            sh 'go get -t -v ./...'
        }
    }
    stage('test go'){
      steps {
        sh 'go test -v ./...'
      }
    }
  }
}