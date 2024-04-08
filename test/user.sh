




API_URL="http://localhost:8080/api/v1/user"
reqBody='{"name": "name001","email":"email"}'
URL=$API_URL""

resultConfig=$(curl -X POST $URL  -H "request-id: $RANDOM" -d "$reqBody")
if [ $? -ne 0 ]; then
  echo "Error: Failed to PUT the HTTP request"
  exit 1
fi


curl $API_URL"/111"