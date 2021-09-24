act=$(go test -v -coverpkg=./... -coverprofile=cover.out ./...\
| grep coverage\
| awk 'NR==1 {print substr($2, 0, length($2)-1)}'\
&& go tool cover -html=cover.out -o cover.html)

cvg=${cvg:=90}

if [ 1 -eq "$(echo "${act} < ${cvg}" | bc)" ]
then
  echo "Code coverage test failed"
  echo "Expected minimum $cvg%, got $act%"
  exit 1
fi
