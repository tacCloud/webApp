/*
node('jenkins-slave') {
     stage('unit-tests') {
        sh(script: """
            whoami
            docker run --rm alpine /bin/sh -c "echo hello world"
            go version
        """)
    }
}
*/
node('master') {
     stage('lets-do-this') {
        sh(script: """
            whoami
            go version
        """)
    }
}
/*
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
*/
