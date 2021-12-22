*** Notes ***
  - Architech: please have a look file architect.jpg

  - Set limit task per day: in file app.env
      LIMIT_TASK_PER_DAY 
  - Run server:
    docker-compose up
  
  - Test:
    make test

*** Complete Featured ***
  - all features is completed and match with front-end, please check front-end

  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user_id":"1","content":"content 1"}' \
  http://localhost:8080/tasks
  

  curl -v http://localhost:8080/tasks


*** Missing ***
I haven't enough time to finish all test case, so i only write some test case for Controller, Service, Repository.
