go build -o integration/mockserver
cd integration

./mockserver -c conf.toml files &

APP_PID=$!

while ! nc -z localhost 8080; do
  sleep 0.1 # wait for 1/10 of the second before check again
done

echo "server pid is $APP_PID"

go test -v

kill $APP_PID