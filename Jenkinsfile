pipeline {
    agent any

    options {
        ansiColor('xterm')
        skipDefaultCheckout()
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    stages {

        stage('Checkout') {
            steps {
                script {
                    sh "source .env"
                    sh "git rev-parse --short HEAD > .git/commit-id"
                    COMMIT_ID = readFile('.git/commit-id').trim()
                    deleteDir()
                    git url: "https://${GIT_HOST}/${OWNER}/${GIT_REPO}.git", branch: ${GIT_BRANCH}
                }
            }
        }

        stage('BuildDataPipeline') {
            steps {
                sh "docker build -f ${DATA_PIPELINE_NAME_DOCKERFILE} \
                    --build-arg DATA_URL=${DATA_URL} \
                    -t ${GIT_REPO}-${DATA_PIPELINE_NAME}"
            }
        }

        stage('TestDataPipeline') {
            steps {
                script {
                    sh "docker run --rm \
                        --name=${GIT_REPO}-${DATA_PIPELINE_NAME} \
                        -v ${PWD}/data:/app/data \
                        ${GIT_REPO}-${DATA_PIPELINE_NAME}"
                }
            }
        }

        stage('PublishDataPipeline') {
            steps {
                script {
                    // NEW_VERSION = sh(script: 'cat VERSION', returnStdout: true).trim()
                    sh "docker tag \
                        ${COMMIT_ID} \
                        ${OWNER}/${GIT_REPO}/${DATA_PIPELINE_NAME}:${COMMIT_ID}"
                    sh "docker tag \
                        ${COMMIT_ID} \
                        ${OWNER}/${GIT_REPO}/${DATA_PIPELINE_NAME}:latest"
                }

                withCredentials([usernamePassword(credentialsId: 'dockerhub',
                usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh "docker login \
                        --username $USERNAME \
                        --password $PASSWORD \
                        https://index.docker.io/v2/"
                    sh "docker push ${OWNER}/${GIT_REPO}/${DATA_PIPELINE_NAME}:${COMMIT_ID}"
                    sh "docker push ${OWNER}/${GIT_REPO}/${DATA_PIPELINE_NAME}:latest"
                }
            }
        }


    post {
        always {
            script {
                sh "yes | docker container prune"
                sh "yes | docker image prune"
            }
            deleteDir()
        }
    }
}