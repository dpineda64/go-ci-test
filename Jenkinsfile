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
    stage('test go'){
      steps {
        sh 'go --version'
      }
    }
  }
}