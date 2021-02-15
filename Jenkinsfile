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

        def root = tool type: 'go', name: 'Go 1.15'

       // Export environment variables pointing to the directory where Go was installed
       withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
           sh 'go version'
       }
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
