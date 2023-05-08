while true
do
  ab -c 10 -t 5 http://localhost:8080/callOneFunc &
  ab -c 10 -t 5 http://localhost:8080/callTwoFunc &
  sleep 30
done