ginAccounting
==================
description
------------------
  * example project to improve current accounting system
  * HTTP communication between accounting system and other systems
  * orienting business-expandable, shorter and reusable server codes
  * preventing heavy join between journal tables
  * simple journal aggregtion
  * resilient, continuos delivery
  
technical stack
------------------
	1. Go/Gin framework for HTTP service
	2. mongoDB for storing, aggregating accounting journals and other information
	3. Docker/Kubernetes for management, deployment and HA
 
directory description
------------------
  * ginAccounting/
		1. *.go : server program codes
		2. curl_scripts.txt : to test local build
		3. Dockerfile.old : building intermediate result container for last version of 979156:idk
		4. Dockerfile : returns kubernetes deployment container

  * ginAccounting/kubernetes
		kubernetes YAMLs
   
  * ginAccounting/mongo_scripts
		data generation, MR code for mongoDB
   
