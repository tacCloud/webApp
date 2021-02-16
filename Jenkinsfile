pipeline {
    agent {
        kubernetes{
            label 'jenkins-slave'
        }

    }
    environment{
        DOCKER_USERNAME = credentials('DOCKER_USERNAME')
        DOCKER_PASSWORD = credentials('DOCKER_PASSWORD')
        TARGETOS='linux'
        TARGETARCH='amd64'
    }
    stages {
        stage('docker login') {
            steps{
                sh(script: """
                    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
                """, returnStdout: true)
            }
        }

        stage('go test') {
            steps{
                sh(script: """
                    export PATH=$PATH:/usr/local/go/bin
                    go get -u golang.org/x/lint/golint
                    ~/go/bin/golint ./...
                    go test -v ./pkg/inventoryMgr
                """, returnStdout: true)
            }
        }


        stage('docker build') {
            steps{
                sh script: '''
                #!/bin/bash
                docker build . \
                 --build-arg TARGETOS=${TARGETOS} \
                 --build-arg TARGETARCH=${TARGETARCH} \
                 --tag ${DOCKER_USERNAME}/inventory-web-app:${BUILD_NUMBER}
                '''
            }
        }

        stage('docker push') {
            steps{
                sh(script: """
                    docker push ${DOCKER_USERNAME}/inventory-web-app:${BUILD_NUMBER}
                """)
            }
        }

/*
This will be a simple test cluster.  I will do more complete tests in argocd.
*/
        stage('deploy') {
            steps{
                sh script: '''
                #!/bin/bash
                #get kubectl for this demo
                curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
                chmod +x ./kubectl
                cat ./testYaml/test.yaml | sed s/0.0.1/1.0.${BUILD_NUMBER}/g | ./kubectl apply -f -
                '''
        }
    }
}
}
