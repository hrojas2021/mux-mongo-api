FROM golang:alpine AS build

WORKDIR /app/

#Install Dependencies
ADD go.mod go.sum ./
RUN go mod download

ADD . .

#Build application
RUN CGO_ENABLED=0 GOOS=linux go build -v main.go

#Step 2: Copy artifacts on clean image
FROM alpine
COPY --from=build /app/main /app/main
#COPY .env .env

EXPOSE 7000
CMD /app/main


#kubectl apply -f deploy/mongo-config.yaml
#kubectl apply -f deploy/mongo-secret.yaml
#kubectl apply -f deploy/mongo.yaml
#kubectl apply -f deploy/k8s.yaml

#docker compose up -d --build
#docker compose down --remove-orphans


# kubectl delete service go-mongo-svc
# kubectl delete deployment go-mongo-deploy