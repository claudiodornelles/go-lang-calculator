buildApplication() {
  echo 'Installing required dependencies...'
  go mod init calculator
  go get -d -v ./... && echo 'Dependencies installed...'
  echo 'Building application...'
  go build && echo 'Application built'
}

buildApplication