pipeline {
    agent any

    environment {
        GOPATH = "${WORKSPACE}"
    }
    stages {
        stage('Build') {
            steps {
                script {
                    sh 'mkdir -p $GOPATH/src/KitchenHelper-backend'
                    sh 'ln -sf $WORKSPACE $GOPATH/src/KitchenHelper-backend'
                    sh 'go get -d -v ./...'
                    sh 'go install -v ./...'
                }
            }
        }
        stage('UnitTest') {
            steps {
                script {
                    sh 'go test'
                }
            }
        }
        stage('Build-Image') {
            steps {
                script {
                    def siteImage = docker.build("kitchenhelper-backend:${env.BUILD_ID}")
                    siteImage.inside {
                        sh 'echo "Inside the container"'
                    }
                    siteImage.push('latest')
                } 
            }
        }
        stage('Test-Image') {
            steps {
                script {
                    '''
                        Todo:
                            Write integration tests.
                            Testing front end.

                            Further automate by creating a test dynamoDB table using a TF module
                            TF module will create TF table with and return the name
                            The name will be pushed into the environment values;
                            the rest of the values will be taken from SSM. 
                            Will also provision ACCESS and Secret keys for the container for testing.
                            Permissions will be secured with a permission boundary for Jenkins to make
                            sure  that it can't create any users with unwanted permissions.
                    '''
                }
            }
        }
    }
}