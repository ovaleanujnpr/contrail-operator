FROM golang:1.14

RUN apt update && apt install -y docker.io && apt install -y jq

RUN git clone --branch v0.13.x https://github.com/operator-framework/operator-sdk $GOPATH/src/operator-framework/operator-sdk
RUN cd $GOPATH/src/operator-framework/operator-sdk; make install

RUN GO111MODULE="on" go get sigs.k8s.io/kind@v0.7.0

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl && chmod +x kubectl && cp kubectl /usr/bin/