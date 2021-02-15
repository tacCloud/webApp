/*
node('jenkins-slave') {
     stage('unit-tests') {
        sh(script: """
            whoami
            docker run --rm alpine /bin/sh -c "echo hello world"
        """)
    }
}
*/
pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
