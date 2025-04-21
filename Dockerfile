FROM golang:1.23-alpine3.21 AS buildstage

WORKDIR  /app 

COPY . .   

WORKDIR /app/cmd

RUN go build -o  main .



FROM  alpine:3.21

WORKDIR /app 

COPY --from=buildstage  /app/cmd/main .

COPY  --from=buildstage /app/design-ff9af-firebase-adminsdk-fbsvc-7decff48f7.json  .  

CMD [ "/app/main" ]