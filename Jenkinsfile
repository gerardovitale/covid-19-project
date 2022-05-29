pipeline {
    agent any

    environment {
        OWNER = 'gerardovitale'
        GIT_BRANCH = 'main'
        GIT_HOST = 'github.com'
        GIT_REPO = 'covid-19-project'
        DATA_PIPELINE_NAME = 'data-pipeline'
        DATA_PIPELINE_DOCKERFILE = 'data_pipeline/Dockerfile'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    sh "git rev-parse --short HEAD > .git/commit-id"
                    COMMIT_ID = readFile('.git/commit-id').trim()
                    deleteDir()
                    git url: "https://${GIT_HOST}/${OWNER}/${GIT_REPO}.git", branch: "${GIT_BRANCH}"
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
                    sh "docker tag ${COMMIT_ID} \
                        ${OWNER}/${GIT_REPO}/${DATA_PIPELINE_NAME}:${COMMIT_ID}"
                    sh "docker tag ${COMMIT_ID} \
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
    }

    post {
        always {
            script {
                sh "docker container rm ${GIT_REPO}-${DATA_PIPELINE_NAME}"
                sh "docker image rm ${GIT_REPO}-${DATA_PIPELINE_NAME}"
            }
            deleteDir()
        }
    }
}