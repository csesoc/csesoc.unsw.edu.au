pipeline {
    agent none

    stages {
        stage('Database') {
            agent {
                docker { image 'mongo:latest'}
            }
            steps {
                echo 'Mongo Database Setup.'
            }
        }
        stage('Build') {
            agent {
                dockerfile true
            }
            steps {
                echo 'Building...'
            }
        }
        stage('Test') {
            agent any
            steps {
                echo 'Testing...'
                sh 'newman run sponsors.postman_collection.json --suppress-exit-code 1'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}