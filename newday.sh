day=$1
mkdir day$day
cd day$day
touch day$day.go
touch day$day_test.go  
echo "package day$day" > day$day.go
echo "package day$day" > day$day_test.go

touch input.txt
touch example.txt