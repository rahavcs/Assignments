pipeline {
    agent any

    // Environment variables
    environment {
        NODE_VERSION = '18.x'  // Specify Node.js version
        DEPLOY_DIR = 'D:\react'  // Deployment directory
        GITHUB_REPO = 'https://github.com/rahavcs/Assignments.git'  // Replace with your repo
    }

    // Tool configurations
    tools {
        nodejs 'Node-18'  // Make sure this matches your Jenkins NodeJS installation name
    }

    stages {
        stage('Checkout') {
            steps {
                // Clean workspace before checking out code
                cleanWs()
                echo 'Cloning repository...'
                git branch: 'main', url: "${GITHUB_REPO}"
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies...'
                sh 'npm install'
            }
        }

        stage('Lint') {
            steps {
                echo 'Running ESLint...'
                sh 'npm run lint || true'  // Add '|| true' to prevent pipeline failure on lint warnings
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
                sh 'npm test -- --watchAll=false'  // Run tests in CI mode
            }
        }

        stage('Build') {
            steps {
                echo 'Building application...'
                sh 'npm run build'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying application...'
                // Create deployment directory if it doesn't exist
                sh 'sudo mkdir -p ${DEPLOY_DIR}'
                
                // Remove old deployment files
                sh 'sudo rm -rf ${DEPLOY_DIR}/*'
                
                // Copy new build files
                sh 'sudo cp -r build/* ${DEPLOY_DIR}/'
                
                // Set proper permissions
                sh 'sudo chown -R jenkins:jenkins ${DEPLOY_DIR}'
                sh 'sudo chmod -R 755 ${DEPLOY_DIR}'
            }
        }

        stage('Post-Deployment Test') {
            steps {
                echo 'Testing deployed application...'
                // Wait for application to be ready
                sh 'sleep 10'
                
                // Test if the application is accessible
                script {
                    def response = sh(
                        script: 'curl -s -o /dev/null -w "%{http_code}" http://localhost',
                        returnStdout: true
                    ).trim()
                    
                    if (response != "200") {
                        error "Deployment verification failed. HTTP status code: ${response}"
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed! Check the logs for details.'
        }
        always {
            // Clean workspace after pipeline completion
            cleanWs()
        }
    }
}