ansible() {
  ansible-playbook calculator.yml -i \$HOME/ansible_hosts
}

collectd() {
  systemctl enable collectd
  systemctl start collectd
  systemctl status collectd
}

buildApplication() {
  echo 'Installing required dependencies...'
  go get -d -v ./... && go install -v ./...
  echo 'Dependencies installed...'
  echo 'Building application...'
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calculator .
  echo 'Application built'
}

"$@"