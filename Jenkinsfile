pipeline {
    agent any

    environment {
        OWNER = 'gerardovitale'
        GIT_BRANCH = 'main'
        GIT_HOST = 'github.com'
        GIT_REPO = 'covid-19-project'
        DATA_URL = 'https://covid.ourworldindata.org/data/owid-covid-data.csv'
        DATA_PIPELINE_NAME = 'data-pipeline'
        DATA_PIPELINE_DOCKERFILE = 'data_pipeline/Dockerfile'
        MONGO_PASS = credentials('MONGO_PASS')
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

        stage('PrepareDataPipeline') {
            steps {
                script {
                    DOCKER_CONTAINER_NAME = env.GIT_REPO + '-' + env.DATA_PIPELINE_NAME
                    DOCKER_IMAGE_NAME = DOCKER_CONTAINER_NAME + ':' + COMMIT_ID
                    sh "docker build -f ${DATA_PIPELINE_DOCKERFILE} \
                        --build-arg DATA_URL=${DATA_URL} \
                        -t ${DOCKER_IMAGE_NAME} ."
                }
            }
        }

        stage('TestDataPipeline') {
            steps {
                script {
                    sh "docker run --rm \
                        --name=${DOCKER_CONTAINER_NAME} \
                        -v ${PWD}/data:/app/data \
                        ${DOCKER_IMAGE_NAME}"
                }
            }
        }

        stage('PublishDataPipeline') {
            steps {
                script {
                    sh "docker tag ${DOCKER_IMAGE_NAME} \
                        ${OWNER}/${GIT_REPO}-${DATA_PIPELINE_NAME}:${COMMIT_ID}"
                    sh "docker tag ${DOCKER_IMAGE_NAME} \
                        ${OWNER}/${GIT_REPO}-${DATA_PIPELINE_NAME}:latest"
                }

                withCredentials([usernamePassword(credentialsId: 'dockerhub',
                usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh "docker login \
                        --username $USERNAME \
                        --password $PASSWORD \
                        https://index.docker.io/v2/"
                    sh "docker push ${OWNER}/${GIT_REPO}-${DATA_PIPELINE_NAME}:${COMMIT_ID}"
                    sh "docker push ${OWNER}/${GIT_REPO}-${DATA_PIPELINE_NAME}:latest"
                }
            }
        }
    }

    post {
        always {
            script {
                sh "docker image rm ${DOCKER_IMAGE_NAME}"
            }
            deleteDir()
        }
    }
}