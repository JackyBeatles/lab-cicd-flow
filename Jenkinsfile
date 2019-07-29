import groovy.json.JsonSlurperClassic

def commitSha1() {
  sh 'git rev-parse HEAD > commitSha1'
  def commit = readFile('commitSha1').trim()
  sh 'rm commitSha1'
  commit.substring(0, 8)
}

def helmDeploy(Map args){
  withCredentials([
    file(
      credentialsId: 'dev-cluster-kubeconfig', 
      variable: 'KUBECONFIG'
    )
  ]) {
    if(args.dry_run){
      echo "Running dry-run deployment"
      sh "/usr/local/bin/helm upgrade --dry-run --wait --kubeconfig $KUBECONFIG -i ${args.deploy_name} --namespace ${args.namespace} --set-string image.tag=${args.image_tag} -f ${args.values_file} ${args.chart_dir}"
    } else if(args.environment == "develop") {
      echo "Running standalone deployment"
      
      echo "Deploy ${args.deploy_name}"
      sh "/usr/local/bin/helm upgrade --wait --kubeconfig $KUBECONFIG -i ${args.deploy_name} --namespace ${args.namespace} --set-string image.tag=${args.image_tag} -f ${args.values_file} ${args.chart_dir}"
    }
  }
}

node('jenkins-node8') {
  env.APP_NAME = "http-go"
  env.DOCKER_REGISTRY = "https://registry.hub.docker.com/"
  env.REPO = "dukecyber/${env.APP_NAME}"
  env.IMAGE_NAME = "${REPO}"
  env.NAMESPACE = "app"
  env.COMMIT_SHA1 = commitSha1()

  stage('Clone repository') {
    echo "Cloning code repository"
    checkout scm
    
    echo "Raw branch name=${env.BRANCH_NAME}"
    env.DEPLOY_NAME = null

    if (env.BRANCH_NAME in ["master"]) {
      env.DEPLOY_NAME = "${env.APP_NAME}-${env.BRANCH_NAME}"
    }

    echo "Parsed deploy_name=${env.DEPLOY_NAME}"

    if(env.DEPLOY_NAME == null) {
      currentBuild.result = 'ABORTED'
      error('Error parsing name or pattern not match.')
    }

  }

  stage('Install test dependencies') {
    // sh "rm package-lock.json"
    // sh "rm -Rf node_modules/"
    // sh "npm install"
  }

  stage('Run unit tests') {
    // sh "go test -v"
  }

  stage('Run OWASP checks') {
    try {

      sh "mkdir -p build/owasp"
      dependencycheck(
        additionalArguments: "--project ${env.APP_NAME} --exclude '**/*.exe' --out build/owasp/dependency-check-report.xml --format XML --noupdate", 
        odcInstallation: "owasp-dep-check"
      )
      dependencyCheckPublisher(
        pattern: 'build/owasp/dependency-check-report.xml'
      )
    } catch (Exception e) {
      echo "Result = " + owasp_output
      echo "Result error = " + e.toString()
      def owasp_output = input(message: 'Fail to scan dependency', ok: 'Continue')
    }
  }

  stage('Run SonarQube scanner') {
    def scannerHome = tool 'sonarqube-scanner';
    withSonarQubeEnv('sonarqube-devcen') {
      sh "${scannerHome}/bin/sonar-scanner"
    }
  } /* stage sonarqube */

  stage('Check Quality gate') {
    timeout(time: 1, unit: 'HOURS') { // Just in case something goes wrong, pipeline will be killed after a timeout
      def qg = waitForQualityGate() // Reuse taskId previously collected by withSonarQubeEnv
      if (qg.status != 'OK') {
        error "Pipeline aborted due to quality gate failure: ${qg.status}"
      }
    }
  }

  stage('Build Docker image') {
    echo "Start building [Build No.:${env.BUILD_NUMBER}]"

    ansiColor('xterm') {
      docker.withRegistry("${env.DOCKER_REGISTRY}", "reg-docker") {
        def img = docker.build("${env.IMAGE_NAME}:${env.BRANCH_NAME}-${env.COMMIT_SHA1}", " .")
      }
    }
  }

  stage('Container image analysis') {
    // placeholder
  }

  stage('Static application security testing') {
    // placeholder
  }

  stage('Push Docker image to registry') {
    docker.withRegistry("${env.DOCKER_REGISTRY}", "reg-docker") {
      def img = docker.build("${env.IMAGE_NAME}:${env.BRANCH_NAME}-${env.COMMIT_SHA1}", " .")
      img.push()
      img.push('v1.0-simple')
    }
  }

  stage('Deploy to K8S') {
    // echo "Deployment [Build No.:${env.BUILD_NUMBER}]"
    // checkout(
    //   [
    //     $class: 'GitSCM',
    //     branches: [[name: 'refs/heads/master']],
    //     doGenerateSubmoduleConfigurations: false,
    //     extensions: [[
    //       $class: 'RelativeTargetDirectory', 
    //       relativeTargetDir: 'app-helm-cluster'
    //     ]],
    //     submoduleCfg: [],
    //     userRemoteConfigs: [[
    //       url: 'ssh://git@git-eco.bigc.co.th/kube/bigc-devcen-infra-cluster.git', 
    //       credentialsId: 'jenkins-bitbucket-ssh'
    //     ]]
    //   ]
    // )
    // def values_file = pwd() + "/bigc-devcen-infra-cluster/helm-nginx/values-oaf-frontend.yaml"
    // def chart_dir = pwd() + "/bigc-devcen-infra-cluster/helm-nginx"
    // helmDeploy(
    //   dry_run     : false,
    //   environment : "develop",
    //   deploy_name : env.DEPLOY_NAME,
    //   namespace   : env.NAMESPACE,
    //   image_tag   : "${env.BRANCH_NAME}-${env.COMMIT_SHA1}",
    //   values_file : values_file,
    //   chart_dir   : chart_dir
    // )
  }

  stage('UI testing') {
    // placeholder
  }

  stage('Load testing') {
    // placeholder
  }
}
