#!/bin/bash
camera=Mac.local

echo -n "Getting health .. "
curl http://localhost:8000/redeye/health
sleep 1
echo -n "Getting config .. "
curl http://localhost:8000/redeye/config
sleep 1
echo -n "Getting messanger .. "
curl http://localhost:8000/redeye/messanger
sleep 1

echo -n "Get video player status"
curl http://localhost:8000/redeye/video
sleep 1

echo "Turning video on for 10 seconds .. "
mosquitto_pub -t camera/${camera} -m on
sleep 1
curl http://localhost:8000/redeye/video
sleep 9

echo "Turning video off"
mosquitto_pub -t camera/${camera} -m off
sleep 1

curl http://localhost:8000/redeye/video
sleep 1

echo "All done."
