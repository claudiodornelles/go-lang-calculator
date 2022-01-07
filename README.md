# How to deploy this application
## 1. Required softwares
1. Download Jenkins from (the official website)[https://www.jenkins.io/download/].
2. Download and install Packer from (the official website)[https://learn.hashicorp.com/tutorials/packer/get-started-install-cli].
3. Download and install Docker from (the official website)[https://docs.docker.com/engine/install/].
## 2. Jenkins configuration
1. After installing Jenkins, access the URL http://localhost:8080 to access its dashboard.
	1. If any problem occurs, and you do not have access to Jenkins dashboard, you can check if it is running by typing the following command in a terminal window `service jenkins status`
	2. You can use this command to perform manual tasks with Jenkins `service jenkins { start | stop | status | restart | force-reload }`
	3. Jenkins uses by default the port 8080, you can change that within the file `/etc/default/jenkins`. Just change the property `HTTP_PORT`
2. Once you access Jenkins Dashboard for the first time, it will ask you for a password. This password can be found at `/var/lib/jenkins/secrets/initialAdminPassword`
	1. It's recommended to chance the default admin password at this time.
3. Follow along Jenkins' configuration wizard and install the recommended plugins.
## 3. Create Jenkins jobs
### 1. Job 1 - Bake
1. With Jenkins properly configured, access its main dashboard and click `New Item` on the left-hand menu.
2. Enter the item name as `bake` and select `Pipeline` option.
3. Click `OK` to create the job.
4. Select the tab `Advanced Project Options` and change the Pipeline Definition from `Pipeline script` to `Pipeline script from SCM`.
5. Select `Git` for the `SCM`.
6. Add the repository URL `https://github.com/claudiodornelles/go-lang-calculator.git`.
7. Scroll down to the bottom of the page and change the field `Script Path` from `Jenkinsfile` to `bake/Jenkinsfile`.
8. Click save.
### 2. Job 2 - Launch
1. Get back to Jenkins dashboard and repeat the steps from Job 1.
	1. Click `New Item` on the left-hand menu.
	2. Enter the item name as `launch` and select `Pipeline` option.
	3. Click `OK` to create the job.
	4. Select the tab `Advanced Project Options` and change the Pipeline Definition from `Pipeline script` to `Pipeline script from SCM`.
	5. Select `Git` for the `SCM`.
	6. Add the repository URL `https://github.com/claudiodornelles/go-lang-calculator.git`.
	7. Scroll down to the bottom of the page and change the field `Script Path` from `Jenkinsfile` to `launch/Jenkinsfile`.
	8. Click save.
## 4. Deploy application
1. After creating jobs 1 and 2, you should be able to see them at the main page of Jenkins dashboard.
2. Click at the clock icon on the right-hand side of the screen that corresponds to the `bake` job we have created earlier.
3. After `bake` job build is completed, click at the clock icon on the right-hand side of the screen that corresponds to the `launch` job we have created earlier. This job will launch the Docker container at http://localhost:8080
## 5. How to use the applicaiton
1. You can check if the application is running by running the command `docker ps`.
2. By this time everything should be up and running, and you can now use the application with the following endpoints:
    1. http://localhost:8090/calc/history - This will return a list of the performed operations.
    2. http://localhost:8090/calc/{operation}/{firstValue}/{secondValue}
       1. Available operations: `sum, sub, mul, div`.
3. If you want to stop the container, run the following commands.
   1. Run `docker ps` and copy the container id.
   2. And then `docker stop <CONTAINER ID>` to stop the container.
