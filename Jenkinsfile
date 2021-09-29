pipeline {
    agent any
    stages {
        stage('Stage 1'){
            steps {
                echo 'Hello world!'
            }
        }
        stage('Stage 2'){
            steps {
                sh 'python main.py'
            }
        }
        stage('Stage 3'){
            steps {
                sh 'python main.py'
            }
        }
    }
}
