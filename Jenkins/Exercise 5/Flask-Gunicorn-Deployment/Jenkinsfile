pipeline {
    agent any

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/rahavcs/Go-Nexturn-Assignments.git'
            }
        }

        stage('Set up Python Environment') {
            steps {
                bat 'python -m venv venv'
                bat 'venv\\Scripts\\activate && pip install -r requirements.txt'
            }
        }

        stage('Run Unit Tests') {
            steps {
                bat 'venv\\Scripts\\activate && pytest'
            }
        }

        stage('Start Gunicorn') {
            steps {
                bat 'venv\\Scripts\\activate && gunicorn -b 127.0.0.1:8000 app:app'
            }
        }

        stage('Verify Deployment') {
            steps {
                bat 'curl http://127.0.0.1:8000'
            }
        }
    }
}
