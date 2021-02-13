node('jenkins-slave') {
     stage('unit-tests') {
        sh(script: """
            whoami
            docker run --rm alpine /bin/sh -c "echo hello world"
        """)
    }
}
